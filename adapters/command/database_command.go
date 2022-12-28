package command

import (
	"context"
	"fmt"
	"strings"

	"gitlab.com/alexandrevinet/leasing/fixtures"

	"github.com/rs/zerolog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"gitlab.com/alexandrevinet/leasing/business/entities"
	"gitlab.com/alexandrevinet/leasing/migrations"
)

type dbCommand struct {
	l        *zerolog.Logger
	db       *bun.DB
	migrator *migrate.Migrator
}

// Init database migration tables.
func (c dbCommand) Init(ctx *cli.Context) error {
	return c.migrator.Init(ctx.Context)
}

// Migrate all unapplied migrations.
func (c dbCommand) Migrate(ctx *cli.Context) error {
	if err := c.migrator.Lock(ctx.Context); err != nil {
		return err
	}
	defer func(migrator *migrate.Migrator, ctx2 context.Context) {
		err := migrator.Unlock(ctx2)
		if err != nil {
			c.l.Fatal().Msg(err.Error())
		}
	}(c.migrator, ctx.Context)

	group, err := c.migrator.Migrate(ctx.Context)
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")

		return nil
	}

	fmt.Printf("migrated to %s\n", group)

	return nil
}

// Rollback the last migration group.
func (c dbCommand) Rollback(ctx *cli.Context) error {
	if err := c.migrator.Lock(ctx.Context); err != nil {
		return err
	}
	defer func(migrator *migrate.Migrator, ctx2 context.Context) {
		err := migrator.Unlock(ctx2)
		if err != nil {
			c.l.Fatal().Msg(err.Error())
		}
	}(c.migrator, ctx.Context)

	group, err := c.migrator.Rollback(ctx.Context)
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Printf("there are no groups to roll back\n")

		return nil
	}

	fmt.Printf("rolled back %s\n", group)

	return nil
}

// Lock will lock migration table.
func (c dbCommand) Lock(ctx *cli.Context) error {
	return c.migrator.Lock(ctx.Context)
}

// Unlock will unlock migration table.
func (c dbCommand) Unlock(ctx *cli.Context) error {
	return c.migrator.Unlock(ctx.Context)
}

// CreateGoMigration command to create sql migration skeleton in go file.
func (c dbCommand) CreateGoMigration(ctx *cli.Context) error {
	name := strings.Join(ctx.Args().Slice(), "_")

	mf, err := c.migrator.CreateGoMigration(ctx.Context, name)
	if err != nil {
		return err
	}

	fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)

	return nil
}

// CreateSQLMigration command to create sql migration skeleton in sql file.
func (c dbCommand) CreateSQLMigration(ctx *cli.Context) error {
	name := strings.Join(ctx.Args().Slice(), "_")

	files, err := c.migrator.CreateSQLMigrations(ctx.Context, name)
	if err != nil {
		return err
	}

	for _, mf := range files {
		fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
	}

	return nil
}

// Status command to display all unapplied migrations and the last migration.
func (c dbCommand) Status(ctx *cli.Context) error {
	ms, err := c.migrator.MigrationsWithStatus(ctx.Context)
	if err != nil {
		return err
	}

	fmt.Printf("migrations: %s\n", ms)
	fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
	fmt.Printf("last migration group: %s\n", ms.LastGroup())

	return nil
}

// MarkApplied command to mak a migration as applied without running it.
func (c dbCommand) MarkApplied(ctx *cli.Context) error {
	group, err := c.migrator.Migrate(ctx.Context, migrate.WithNopMigration())
	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to mark as applied\n")

		return nil
	}

	fmt.Printf("marked as applied %s\n", group)

	return nil
}

// Fixtures command to load all fixtures.
func (c dbCommand) Fixtures(ctx *cli.Context) error {
	c.db.RegisterModel((*entities.Car)(nil), (*entities.Customer)(nil), (*entities.Schedule)(nil))

	fixture := dbfixture.New(c.db, dbfixture.WithTruncateTables())

	err := fixtures.Load(ctx.Context, fixture)
	if err != nil {
		return err
	}

	fmt.Printf("fixtures loaded\n")

	return nil
}

// newDBCommand create db command with all subcommands.
func newDBCommand(l *zerolog.Logger, db *bun.DB) *cli.Command {
	command := &dbCommand{
		l:        l,
		db:       db,
		migrator: migrate.NewMigrator(db, migrations.NewMigrations()),
	}

	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "create migration tables",
				Action: command.Init,
			},
			{
				Name:   "migrate",
				Usage:  "migrate database",
				Action: command.Migrate,
			},
			{
				Name:   "rollback",
				Usage:  "rollback the last migration group",
				Action: command.Rollback,
			},
			{
				Name:   "lock",
				Usage:  "lock migrations",
				Action: command.Lock,
			},
			{
				Name:   "unlock",
				Usage:  "unlock migrations",
				Action: command.Unlock,
			},
			{
				Name:   "create_go",
				Usage:  "create Go migration",
				Action: command.CreateGoMigration,
			},
			{
				Name:   "create_sql",
				Usage:  "create up and down SQL migrations",
				Action: command.CreateSQLMigration,
			},
			{
				Name:   "status",
				Usage:  "print migrations status",
				Action: command.Status,
			},
			{
				Name:   "mark_applied",
				Usage:  "mark migrations as applied without actually running them",
				Action: command.MarkApplied,
			},
			{
				Name:   "fixtures",
				Usage:  "truncate db and load all fixtures",
				Action: command.Fixtures,
			},
		},
	}
}
