package persistence

import (
	"context"
	"database/sql"

	"github.com/ddd/src/context/log_handler/domain/model/logfile"
)

type Repo struct {
	DB *sql.DB
}

type LogFileRepository interface {
	Add(entity logfile.LogFile, ctx context.Context) error
	GetAll(ctx context.Context) []logfile.LogFile
}

func (r *Repo) Add(entity logfile.LogFile, ctx context.Context) error {
	query := "INSERT INTO `log_file` (`path`) VALUES (?)"
	_, err := r.DB.ExecContext(ctx, query, entity.GetPath())
	return err
}

func (r *Repo) GetAll(ctx context.Context) []logfile.LogFile {
	return nil
}

func NewLogFileRepository(db *sql.DB) LogFileRepository {
	return &Repo{DB: db}
}
