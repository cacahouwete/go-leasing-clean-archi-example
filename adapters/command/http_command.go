package command

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"gitlab.com/alexandrevinet/leasing/adapters/controller"
	"gitlab.com/alexandrevinet/leasing/business/usecases"
	"gitlab.com/alexandrevinet/leasing/config"
	"gitlab.com/alexandrevinet/leasing/infrastracture/httpserver"
)

type httpCommand struct {
	cfg config.Config
	l   *zerolog.Logger
	uc  *usecases.UseCases
}

// Run start http server and wait sigterm to shut down the server in the correct way.
func (hc httpCommand) Run(ctx *cli.Context) error {
	// HTTP Server
	handler := gin.New()
	controller.NewRouter(handler, hc.l, hc.uc)
	httpServer := httpserver.New(handler, httpserver.Port(hc.cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		hc.l.Info().Msg("app - Run - signal: " + s.String())
	case notifyErr := <-httpServer.Notify():
		hc.l.Error().Msg("app - Run - httpServer.Notify: " + notifyErr.Error())
	}

	// Shutdown
	if err := httpServer.Shutdown(ctx.Context); err != nil {
		hc.l.Error().Msg("app - Run - httpServer.Shutdown: " + err.Error())
	}

	return nil
}

// newHTTPCommand create http command with all subcommands.
func newHTTPCommand(cfg config.Config, l *zerolog.Logger, uc *usecases.UseCases) *cli.Command {
	command := httpCommand{cfg, l, uc}

	return &cli.Command{
		Name:  "http",
		Usage: "http command",
		Subcommands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run http server",
				Action: command.Run,
			},
		},
	}
}
