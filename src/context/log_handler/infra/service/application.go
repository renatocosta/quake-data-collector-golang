package service

import (
	"context"
	"database/sql"

	"github.com/ddd/crosscutting/building_blocks/domain"
	"github.com/ddd/crosscutting/building_blocks/infra/bus"
	"github.com/ddd/src/context/log_handler/app"
	"github.com/ddd/src/context/log_handler/app/command"
	"github.com/ddd/src/context/log_handler/app/event_handler"
	"github.com/ddd/src/context/log_handler/app/query"
	"github.com/ddd/src/context/log_handler/domain/model/logfile/events"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
)

func NewApplication(ctx context.Context, selectLogFileGr *errgroup.Group, db *sql.DB) app.Application {

	additionalDependencies := bus.NewAdditionalDependencies(db)

	logFileSelectedChan := make(chan domain.Event)

	eventBus := bus.NewEventBus()
	subscriberLogFileSelected := events.LogFileSelectedEvent
	eventBus.Subscribe(subscriberLogFileSelected, logFileSelectedChan)
	subscriberLogFileHandlers := []bus.EventHandlerFunc{event_handler.SelectLogFileEventHandler}

	eventHandlers := map[string][]bus.EventHandlerFunc{
		subscriberLogFileSelected: subscriberLogFileHandlers,
	}

	go bus.HandleEvent(ctx, selectLogFileGr, logFileSelectedChan, additionalDependencies, eventHandlers)

	return app.Application{
		Commands: app.Commands{
			SelectLogFile:      command.NewSelectLogFileHandler(eventBus, db),
			CreateHumanLogFile: command.NewCreateHumanLogFileHandler(eventBus),
		},
		Queries: app.Queries{
			LogFiles: query.NewAvailableLogFilesHandler(),
		},
	}
}
