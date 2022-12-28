//nolint:testpackage // This is not a test package per-se
package tests

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"gitlab.com/alexandrevinet/leasing/tests/contexts"
)

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
	}

	c := contexts.New()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: c.InitializeTestSuite,
		ScenarioInitializer:  c.InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
