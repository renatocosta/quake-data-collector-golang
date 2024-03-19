package query

import (
	"context"

	"github.com/ddd/crosscutting/building_blocks/app"
	"github.com/ddd/src/context/log_handler/domain/model/logfile"
)

type AvailableLogFiles struct {
}

type AvailableFilesHandler app.QueryHandler[AvailableLogFiles, *[]logfile.LogFileEntity]

type availableLogFilesHandler struct {
}

func NewAvailableLogFilesHandler() AvailableFilesHandler {
	return availableLogFilesHandler{}
}

func (h availableLogFilesHandler) Handle(ctx context.Context, query AvailableLogFiles) (*[]logfile.LogFileEntity, error) {
	return nil, nil
}
