package event_handler

import (
	"context"

	"github.com/ddd/crosscutting/building_blocks/domain"
	"github.com/ddd/crosscutting/building_blocks/infra/bus"
	"github.com/ddd/src/context/log_handler/domain/model/logfile"
	"github.com/ddd/src/context/log_handler/domain/model/logfile/events"
)

// UserRegisteredHandler handles the user registered event
func SelectLogFileEventHandler(ctx context.Context, event domain.Event, dependencies bus.AdditionalDependencies) error {

	return dependencies.LogFileRepo.Add(
		logfile.NewLogFile(event.Data.(events.LogFileSelected).Path),
		ctx,
	)

}
