package infrastructure

import (
	"strconv"
)

func EuToUsTimeConverter(hours, minutes string) string {
	numHours, err := strconv.Atoi(hours)
	if err != nil || numHours < 0 || numHours > 23 {
		return "Can't to convert! Invalid hours value"
	}

	numMinutes, err := strconv.Atoi(minutes)
	if err != nil || numMinutes < 0 || numMinutes > 59 {
		return "Can't to convert! Invalid minutes value"
	}

	var zeroPrefix string
	if numMinutes < 10 {
		zeroPrefix = "0"
	}

	if numHours <= 12 {
		return strconv.Itoa(numHours) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " AM"
	} else {
		return strconv.Itoa(numHours % 12) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " PM"
	}
}
