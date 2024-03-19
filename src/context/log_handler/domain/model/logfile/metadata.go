package logfile

import "errors"

var (
	RequiredLogFileExtension  = ".log"
	ErrMinSizeRequiredMessage = errors.New("Min size required")
	ErrMissingFileExtension   = errors.New("Missing file extension")
)

type Metadata struct {
	Size      int64
	Extension string
}

func Validate(m *Metadata) error {

	if m.Size < 1 {
		return ErrMinSizeRequiredMessage
	}

	if m.Extension != RequiredLogFileExtension {
		return ErrMissingFileExtension
	}

	return nil
}

func (m *Metadata) GetSize() int64 {
	return m.Size
}

func (m *Metadata) GetExtension() string {
	return m.Extension
}
