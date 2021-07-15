package infrastructure

import (
	"fmt"
	"strconv"
)

func EuToUsTimeConverter(hours, minutes string) (string, error) {
	numHours, err := strconv.Atoi(hours)
	if err != nil || numHours < 0 || numHours > 23 {
		return "", fmt.Errorf("Server error! Hours value should be an integer between 0 and 23")
	}

	numMinutes, err := strconv.Atoi(minutes)
	if err != nil || numMinutes < 0 || numMinutes > 59 {
		return "", fmt.Errorf("Server error! Minutes value should be an integer between 0 and 59")
	}

	var zeroPrefix string
	if numMinutes < 10 {
		zeroPrefix = "0"
	}

	if numHours <= 12 {
		return strconv.Itoa(numHours) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " AM", nil
	} else {
		return strconv.Itoa(numHours % 12) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " PM", nil
	}
}
