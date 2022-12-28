package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var migrationsFiles embed.FS

// NewMigrations return struct with all migrations in current directory.
func NewMigrations() *migrate.Migrations {
	migrations := migrate.NewMigrations()
	if errM := migrations.Discover(migrationsFiles); errM != nil {
		panic(errM)
	}

	return migrations
}
