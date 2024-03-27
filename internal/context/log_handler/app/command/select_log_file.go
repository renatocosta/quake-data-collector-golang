package command

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
	"github.com/ddd/pkg/building_blocks/app"
	"github.com/ddd/pkg/building_blocks/infra/bus"
)

type SelectLogFileCommand struct {
	Path string
}

type SelectLogFileHandler app.CommandHandler[SelectLogFileCommand, []string]

type selectLogFileHandler struct {
	eventBus *bus.EventBus
}

func NewSelectLogFileHandler(eventBus *bus.EventBus) app.CommandHandler[SelectLogFileCommand, []string] {
	return selectLogFileHandler{
		eventBus: eventBus,
	}
}

func (h selectLogFileHandler) Handle(ctx context.Context, cmd SelectLogFileCommand) ([]string, error) {

	lgfile := logfile.NewLogFile(cmd.Path)

	logfile, err := lgfile.ReadFrom(cmd.Path)

	if err != nil {
		return []string{}, err
	}

	lines, err := logfile.ExtractFrom(logfile.File)

	if err != nil {
		return []string{}, err
	}

	logFile := logfile.Select(lines)

	for _, event := range logFile.Events {
		h.eventBus.Publish(event)
	}

	return lines, nil
}
