package service

import (
	"context"

	"github.com/ddd/internal/context/log_handler/app"
	"github.com/ddd/internal/context/log_handler/app/command"
	"github.com/ddd/internal/context/log_handler/app/event_handler"
	"github.com/ddd/internal/context/log_handler/app/query"
	"github.com/ddd/internal/context/log_handler/domain/model/logfile/events"
	"github.com/ddd/pkg/building_blocks/domain"
	"github.com/ddd/pkg/building_blocks/infra/bus"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
)

func NewApplication(ctx context.Context, selectLogFileGr *errgroup.Group) app.Application {

	additionalDependencies := bus.NewAdditionalDependencies()

	logFileSelectedChan := make(chan domain.Event, 1)

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
			SelectLogFile:      command.NewSelectLogFileHandler(eventBus),
			CreateHumanLogFile: command.NewCreateHumanLogFileHandler(eventBus),
		},
		Queries: app.Queries{
			LogFiles: query.NewAvailableLogFilesHandler(),
		},
	}
}
