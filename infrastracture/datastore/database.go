package datastore

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// New Open database connection, create bun DB and return it.
func New(opts ...pgdriver.Option) *bun.DB {
	pgcon := pgdriver.NewConnector(opts...)

	sqlDB := sql.OpenDB(pgcon)

	return bun.NewDB(sqlDB, pgdialect.New())
}
