package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		purpose,
		filename,
		wantResult string
		wantError error
	}{
		{
			purpose:    "ParseFile() /Success Case: Valid file data, should parse and return.",
			filename:   "1.txt",
			wantResult: "1 3 N\r\n5 1 E\r\n",
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, invalid plateau coordinates.",
			filename:  "2.txt",
			wantError: errInvalidPlateauData,
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, inconsistent plateau data.",
			filename:  "3.txt",
			wantError: errInconsistentPlateauData,
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, inconsistent rover data.",
			filename:  "4.txt",
			wantError: errInconsistentRoverData,
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, invalid rover coordinates.",
			filename:  "5.txt",
			wantError: errInvalidRoverData,
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, invalid compass direction.",
			filename:  "6.txt",
			wantError: errInvalidCompassDirection,
		},
		{
			purpose:   "ParseFile() /Error Case: Invalid file data, invalid rover movement command string.",
			filename:  "7.txt",
			wantError: errMovementCommand,
		},
	}

	for _, test := range tt {
		t.Run(test.purpose, func(t *testing.T) {
			res, err := ParseFile(fmt.Sprintf("./testdata/%s", test.filename))
			if err == nil {
				assert.Equal(res, test.wantResult)
			}
			assert.Equal(err, test.wantError)
		})
	}
}
