package command

import (
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/urfave/cli/v2"
	"gitlab.com/alexandrevinet/leasing/business/usecases"
	"gitlab.com/alexandrevinet/leasing/config"
)

// NewCommands create urface command collection by calling all new command in current package.
func NewCommands(cfg config.Config, l *zerolog.Logger, uc *usecases.UseCases, db *bun.DB) []*cli.Command {
	return []*cli.Command{
		newDBCommand(l, db),
		newHTTPCommand(cfg, l, uc),
	}
}
