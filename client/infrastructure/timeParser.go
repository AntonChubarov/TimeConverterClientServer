package infrastructure

import (
	"fmt"
	"strconv"
	"strings"
)

func TryToParse(time string) (int, int, error) {

	time = strings.ReplaceAll(time, " ", "")

	if !strings.Contains(time, ":") {
		return 0, 0, fmt.Errorf("Hours and minutes should be divided with \":\"")
	}

	hm := strings.Split(time, ":")

	if len(hm) > 2 {
		return 0, 0, fmt.Errorf("Only one \":\" divider should be entered")
	}

	numHours, err := strconv.Atoi(hm[0])
	if err != nil || numHours < 0 || numHours > 23 {
		return 0, 0, fmt.Errorf("Hours value should be an integer between 0 and 23")
	}

	numMinutes, err := strconv.Atoi(hm[1])
	if err != nil || numMinutes < 0 || numMinutes > 59 {
		return 0, 0, fmt.Errorf("Minutes value should be an integer between 0 and 59")
	}

	return numHours, numMinutes, nil
}