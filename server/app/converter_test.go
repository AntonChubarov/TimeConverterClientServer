package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ConverterTestCase struct {
	Hours string
	Minutes string
	ConvertedTime string
	Error error
}

func TestConvert(t *testing.T) {
	cases := []ConverterTestCase{
		{"6", "04", "6:04 AM", nil},
		{"15", "28", "3:28 PM", nil},
		{"12", "00", "12:00 AM", nil},
		{"12", "01", "12:01 PM", nil},
		{"12", "15", "12:15 PM", nil},
		{"36", "25", "", fmt.Errorf("hours value should be an integer between 0 and 23")},
		{"ab", "cd", "", fmt.Errorf("hours value should be an integer between 0 and 23")},
		{"12", "72", "", fmt.Errorf("minutes value should be an integer between 0 and 59")},
		{"14", "ab", "", fmt.Errorf("minutes value should be an integer between 0 and 59")},
	}


	for i := range cases {
		convertedTime, err := ConvertEuTimeToUs(cases[i].Hours, cases[i].Minutes)
		assert.Equal(t, cases[i].ConvertedTime, convertedTime)
		assert.Equal(t, err, cases[i].Error)
	}
}