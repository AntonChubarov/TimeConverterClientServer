package infrastructure

import (
	"fmt"
	"strings"
)

func TryToParse(time string) (string, string, error) {

	time = strings.ReplaceAll(time, " ", "")

	if !strings.Contains(time, ":") {
		return "0", "0", fmt.Errorf("Hours and minutes should be divided with \":\"")
	}

	hm := strings.Split(time, ":")

	return hm[0], hm[1], nil
}