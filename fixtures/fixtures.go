package fixtures

import (
	"context"
	"os"

	"github.com/uptrace/bun/dbfixture"
)

func Load(ctx context.Context, fixture *dbfixture.Fixture) error {
	err := fixture.Load(
		ctx,
		os.DirFS("fixtures"),
		"fixtures.yaml",
	)
	if err != nil {
		return err
	}

	return nil
}
