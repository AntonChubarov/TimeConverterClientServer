package app

import (
	"fmt"
	"strconv"
)

func ConvertEuTimeToUs(hours, minutes string) (string, error) {
	numHours, err := convertAndValidateByMaxVal(hours, 23)
	if err != nil  {
		return "", fmt.Errorf("hours " + err.Error())
	}

	numMinutes, err := convertAndValidateByMaxVal(minutes, 59)
	if err != nil || numMinutes < 0 || numMinutes > 59 {
		return "", fmt.Errorf("minutes " + err.Error())
	}

	var zeroPrefix string
	if numMinutes < 10 {
		zeroPrefix = "0"
	}

	if numHours <= 12 {
		return strconv.Itoa(numHours) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " AM", nil
	}
	return strconv.Itoa(numHours % 12) + ":" + zeroPrefix + strconv.Itoa(numMinutes) + " PM", nil
}

func convertAndValidateByMaxVal(input string, maxVal int) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil || num < 0 || num > maxVal {
		return 0, fmt.Errorf("value should be an integer between 0 and %d", maxVal)
	}
	return num, nil
}
