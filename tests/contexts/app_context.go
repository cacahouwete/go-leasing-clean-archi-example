package contexts

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/cucumber/godog"
	_ "github.com/lib/pq"
)

const NbSecondToWait = 10

type AppContext struct {
	Cmd    *exec.Cmd
	Ctx    context.Context
	Cancel context.CancelFunc
	Done   chan bool
	Logs   []string
}

// Restart will cancel previous run if exist and exec a new one
// This function will store all stdout application in Logs to be used in logs check step.
func (a *AppContext) Restart() error {
	if a.Cmd != nil {
		a.Cancel()

		<-a.Done

		close(a.Done)
		a.Done = nil
	}

	a.Done = make(chan bool)

	a.Ctx, a.Cancel = context.WithCancel(context.Background())
	a.Cmd = exec.CommandContext(a.Ctx, "../goapp", "http", "run")

	cmdReader, errO := a.Cmd.StdoutPipe()
	if errO != nil {
		log.Fatal(errO.Error())
	}

	scanner := bufio.NewScanner(cmdReader)

	go func() {
		defer cmdReader.Close()

		for scanner.Scan() {
			a.Logs = append(a.Logs, scanner.Text())
		}

		a.Done <- true
	}()

	if errC := a.Cmd.Start(); errC != nil {
		return errC
	}

	err := a.WaitUntilToSeeInIog(NbSecondToWait, "[GIN-debug] [WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production.")
	if err != nil {
		fmt.Println(a.Logs)

		return err
	}

	return nil
}

// DbClean will delete all data and load all fixtures.
func (a *AppContext) DbClean() error {
	cmdInit := exec.CommandContext(context.Background(), "../goapp", "db", "init")

	if errI := cmdInit.Run(); errI != nil {
		return errI
	}

	cmdMigrate := exec.CommandContext(context.Background(), "../goapp", "db", "migrate")

	if errM := cmdMigrate.Run(); errM != nil {
		return errM
	}

	cmdFixtures := exec.CommandContext(context.Background(), "../goapp", "db", "fixtures")

	if errF := cmdFixtures.Run(); errF != nil {
		return errF
	}

	return nil
}

// Reset will run DbClean and Restart.
func (a *AppContext) Reset() error {
	if err := a.DbClean(); err != nil {
		return err
	}

	if errR := a.Restart(); errR != nil {
		return errR
	}

	return nil
}

// WaitUntilToSeeInIog will check each second if arg2 match with any logs
// If log not found during time defined in arg1 it will dump all logs to help on debug and return an error.
func (a *AppContext) WaitUntilToSeeInIog(arg1 int, arg2 string) error {
	for arg1 >= 0 {
		for _, lastLog := range a.Logs {
			if lastLog == arg2 {
				return nil
			}
		}

		time.Sleep(time.Second)
		arg1--
	}

	a.DumpLogs()

	return fmt.Errorf("log not found")
}

// WaitUntilToSeePartialInIog will check each second if arg2 is contained in any logs
// If log not found during time defined in arg1 it will dump all logs to help on debug and return an error.
func (a *AppContext) WaitUntilToSeePartialInIog(arg1 int, arg2 string) error {
	for arg1 >= 0 {
		for _, lastLog := range a.Logs {
			if strings.Contains(lastLog, arg2) {
				return nil
			}
		}

		time.Sleep(time.Second)
		arg1--
	}

	a.DumpLogs()

	return fmt.Errorf("log not found")
}

// ShouldNotSeePartialInIog will check if arg1 is not contained in any logs
// If log found it will dump all logs to help on debug and return an error.
func (a *AppContext) ShouldNotSeePartialInIog(arg1 string) error {
	for _, lastLog := range a.Logs {
		if strings.Contains(lastLog, arg1) {
			a.DumpLogs()

			return fmt.Errorf("log found")
		}
	}

	return nil
}

// WaitUntilToSeeXTimesPartialInIog will check each second if arg3 is contained exactly nb defined in arg2 in logs
// If log not found or found to many times during time defined in arg1 it will dump all logs to help on debug and return an error.
func (a *AppContext) WaitUntilToSeeXTimesPartialInIog(arg1, arg2 int, arg3 string) error {
	for arg1 >= 0 {
		i := arg2

		for _, lastLog := range a.Logs {
			if strings.Contains(lastLog, arg3) {
				i--
			}
		}

		if i == 0 {
			return nil
		}

		if i < 0 {
			a.DumpLogs()

			return fmt.Errorf("log found to many times")
		}

		time.Sleep(time.Second)
		arg1--
	}

	a.DumpLogs()

	return fmt.Errorf("log not found")
}

// DumpLogs will dump all logs in stderr to help in debug.
func (a *AppContext) DumpLogs() {
	for _, l := range a.Logs {
		println(l)
	}
}

func (a *AppContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Each scenario must be independent that's why we need to clean logs
		a.Logs = nil

		return ctx, nil
	})

	ctx.Step(`^dump logs$`, a.DumpLogs)
	ctx.Step(`^I restart the application$`, a.Restart)
	ctx.Step(`^should not see partial log '([^']*)'$`, a.ShouldNotSeePartialInIog)
	ctx.Step(`^should not see partial log "([^"]*)"$`, a.ShouldNotSeePartialInIog)
	ctx.Step(`^wait (\d+) seconds until to see log '([^']*)'$`, a.WaitUntilToSeeInIog)
	ctx.Step(`^wait (\d+) seconds until to see log "([^"]*)"$`, a.WaitUntilToSeeInIog)
	ctx.Step(`^wait (\d+) seconds until to see partial log '([^']*)'$`, a.WaitUntilToSeePartialInIog)
	ctx.Step(`^wait (\d+) seconds until to see partial log "([^"]*)"$`, a.WaitUntilToSeePartialInIog)
	ctx.Step(`^wait (\d+) seconds until to see (\d+) times partial log '([^']*)'$`, a.WaitUntilToSeeXTimesPartialInIog)
	ctx.Step(`^wait (\d+) seconds until to see (\d+) times partial log "([^"]*)"$`, a.WaitUntilToSeeXTimesPartialInIog)
}

func NewAPPContext() *AppContext {
	return &AppContext{}
}
