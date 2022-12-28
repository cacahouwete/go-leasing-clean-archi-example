// Package app configures and runs application.
package app

import (
	"os"

	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/urfave/cli/v2"
	"gitlab.com/alexandrevinet/leasing/adapters/command"
	"gitlab.com/alexandrevinet/leasing/business/usecases"
	"gitlab.com/alexandrevinet/leasing/infrastracture/datastore"

	"gitlab.com/alexandrevinet/leasing/adapters/gateways"
	"gitlab.com/alexandrevinet/leasing/config"
	"gitlab.com/alexandrevinet/leasing/infrastracture/logger"
)

// Run creates objects via constructors.
func Run(cfg config.Config) {
	l := logger.New(cfg.Log.Level)

	db := datastore.New(
		pgdriver.WithDSN(cfg.Db.Dsn),
	)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))

	gw := gateways.New(db)

	// Use case
	uc := usecases.New(gw, l)

	app := &cli.App{
		Name:     "leasing",
		Commands: command.NewCommands(cfg, l, uc, db),
	}

	if errA := app.Run(os.Args); errA != nil {
		l.Fatal().Msg(errA.Error())
	}
}
