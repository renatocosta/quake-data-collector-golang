package logfile

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/ddd/crosscutting/building_blocks/domain"
	"github.com/ddd/src/context/log_handler/domain/model/logfile/events"
)

var (
	ErrNameEmpty        = errors.New("Path file can not be empty")
	ErrFileNotFound     = errors.New("Path file not found")
	ErrUnableToRead     = errors.New("Unable to read file")
	ErrFileContentEmpty = errors.New("Content file can not be empty")
	ErrFileContentSize  = errors.New("Error getting file information")
)

type LogFile interface {
	ReadFrom(path string) (*LogFileEntity, error)
	ExtractFrom(*os.File) ([]string, error)
	Select(rows []string) *LogFileEntity
	GetPath() string
}

type LogFileEntity struct {
	domain.AggregateRoot
	Path string
	File *os.File
}

func NewLogFile(p string) LogFile {
	return &LogFileEntity{
		Path: p,
	}
}

func (l *LogFileEntity) ReadFrom(path string) (*LogFileEntity, error) {

	if path == "" {
		return nil, ErrNameEmpty
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, ErrFileNotFound
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, ErrFileContentSize
	}

	metaData := &Metadata{
		Size:      fileInfo.Size(),
		Extension: filepath.Ext(path),
	}

	if err := Validate(metaData); err != nil {
		return nil, err
	}
	return &LogFileEntity{
		Path: path,
		File: file,
	}, nil

}

func (l *LogFileEntity) GetPath() string {
	return l.Path
}

func (l *LogFileEntity) ExtractFrom(file *os.File) ([]string, error) {
	defer file.Close()

	var rows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, ErrUnableToRead
	}

	return rows, nil
}

func (l *LogFileEntity) Select(rows []string) *LogFileEntity {
	//Raise event
	event := domain.Event{
		Type:      events.LogFileSelectedEvent,
		Timestamp: time.Now(),
		Data: events.LogFileSelected{
			Content: rows,
			Path:    l.GetPath(),
		},
	}

	aggregateRoot := domain.AggregateRoot{
		Id: 0,
	}

	aggregateRoot.RecordThat(event)

	return &LogFileEntity{
		AggregateRoot: aggregateRoot,
	}
}
