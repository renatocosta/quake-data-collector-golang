package persistence

import (
	"context"

	"github.com/ddd/internal/context/log_handler/domain/model/logfile"
)

type Repo struct {
}

type LogFileRepository interface {
	GetAll(ctx context.Context) []logfile.LogFile
}

func (r *Repo) GetAll(ctx context.Context) []logfile.LogFile {
	return nil
}

func NewLogFileRepository() LogFileRepository {
	return &Repo{}
}
