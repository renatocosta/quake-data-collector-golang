package service

import (
	"context"

	"github.com/ddd/src/context/match_reporting/app"
	"github.com/ddd/src/context/match_reporting/app/command"
	_ "github.com/go-sql-driver/mysql"
)

func NewApplication(ctx context.Context) app.Application {

	return app.Application{
		Commands: app.Commands{
			FindPlayersKilled: command.NewFindPlayersKilledHandler(),
		},
	}
}
