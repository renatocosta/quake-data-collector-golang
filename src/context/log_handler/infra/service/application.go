package service

import (
	"context"
	"sync"

	"github.com/ddd/crosscutting/building_blocks/domain"
	"github.com/ddd/crosscutting/building_blocks/infra/bus"
	"github.com/ddd/src/context/log_handler/app"
	"github.com/ddd/src/context/log_handler/app/command"
	"github.com/ddd/src/context/log_handler/app/event_handler"
	"github.com/ddd/src/context/log_handler/app/query"
	"github.com/ddd/src/context/log_handler/domain/model/logfile/events"
	_ "github.com/go-sql-driver/mysql"
)

func NewApplication(ctx context.Context, selectLogFileWg *sync.WaitGroup) app.Application {

	logFileSelectedChan := make(chan domain.Event)

	eventBus := bus.NewEventBus()
	subscriberLogFileSelected := events.LogFileSelectedEvent
	eventBus.Subscribe(subscriberLogFileSelected, logFileSelectedChan)
	subscriberLogFileHandlers := []bus.EventHandlerFunc{event_handler.SelectLogFileEventHandler}
	selectLogFileWg.Add(len(subscriberLogFileHandlers))

	eventHandlers := map[string][]bus.EventHandlerFunc{
		subscriberLogFileSelected: subscriberLogFileHandlers,
	}

	go bus.HandleEvent(ctx, selectLogFileWg, logFileSelectedChan, eventHandlers)

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
