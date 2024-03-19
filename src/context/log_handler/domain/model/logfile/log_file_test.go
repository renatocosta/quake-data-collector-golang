package logfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFailToFileWhileReading(t *testing.T) {

	var tests = []struct {
		path string
		want error
	}{
		{"", ErrNameEmpty},
		{"/usr/xx/stts.csv", ErrFileNotFound},
		//{noContentPathFile, ErrFileContentSize},
	}

	for _, test := range tests {
		_, err := ReadFrom(test.path)
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}

}
