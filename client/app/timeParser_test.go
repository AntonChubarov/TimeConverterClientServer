package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ConverterTestCase struct {
	Time string
	Hours string
	Minutes string
	Error error
}

func TestConvert(t *testing.T) {
	cases := []ConverterTestCase{
		{"6:04", "6", "04", nil},
		{" 16 : 24", "16", "24", nil},
		{" 75 : 3", "75", "3", nil},
		{"75/03", "0", "0", fmt.Errorf("hours and minutes should be divided with \":\"")},
		{"21 30", "0", "0", fmt.Errorf("hours and minutes should be divided with \":\"")},
	}

	for i := range cases {
		hours, minutes, err := TryToParse(cases[i].Time)
		assert.Equal(t, cases[i].Hours, hours)
		assert.Equal(t, cases[i].Minutes, minutes)
		assert.Equal(t, err, cases[i].Error)
	}
}