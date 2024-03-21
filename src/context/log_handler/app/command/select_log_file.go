package command

import (
	"context"
	"database/sql"

	"github.com/ddd/crosscutting/building_blocks/app"
	"github.com/ddd/crosscutting/building_blocks/infra/bus"
	"github.com/ddd/src/context/log_handler/domain/model/logfile"
)

type SelectLogFileCommand struct {
	Path string
}

type SelectLogFileHandler app.CommandHandler[SelectLogFileCommand, []string]

type selectLogFileHandler struct {
	eventBus *bus.EventBus
	db       *sql.DB
}

func NewSelectLogFileHandler(eventBus *bus.EventBus, db *sql.DB) app.CommandHandler[SelectLogFileCommand, []string] {
	return selectLogFileHandler{
		eventBus: eventBus,
		db:       db,
	}
}

func (h selectLogFileHandler) Handle(ctx context.Context, cmd SelectLogFileCommand) ([]string, error) {

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return lines, nil
}
