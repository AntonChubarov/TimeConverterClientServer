package app

import (
	"fmt"
	"strings"
)

func TryToParse(time string) (hours, minutes string,  err error) {
	time = strings.ReplaceAll(time, " ", "")
	if !strings.Contains(time, ":") {
		return "0", "0", fmt.Errorf("hours and minutes should be divided with \":\"")
	}
	hm := strings.Split(time, ":")
	return hm[0], hm[1], nil
}