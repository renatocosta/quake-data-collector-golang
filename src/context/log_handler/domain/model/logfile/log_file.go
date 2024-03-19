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
	getMetadata() *Metadata
	Select() *LogFileEntity
}

type LogFileEntity struct {
	domain.AggregateRoot
	Path string
	File *os.File
}

type LogFileRows struct {
}

func ReadFrom(path string) (*LogFileEntity, error) {

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

func (l *LogFileEntity) ExtractFrom(f *os.File) ([]string, error) {

	defer f.Close()

	var rows []string
	scanner := bufio.NewScanner(f)
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
	event := events.LogFileSelected{Content: rows /*MetaData: metaData*/}
	eventRaised := domain.Event{
		Type:      events.LogFileSelectedEvent,
		Timestamp: time.Now(),
		Data:      event,
	}

	aggregateRoot := domain.AggregateRoot{
		Id: 0,
	}
	aggregateRoot.RecordThat(eventRaised)

	return &LogFileEntity{
		AggregateRoot: aggregateRoot,
	}
}
