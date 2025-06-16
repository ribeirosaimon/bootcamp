package util

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func NormalizeTime(s string) (time.Time, error) {
	splitedTime := strings.Split(s, ":")
	if len(splitedTime) != 2 {
		return time.Time{}, errors.New("<UNK>")
	}
	hour, err := strconv.Atoi(splitedTime[0])
	if err != nil {
		return time.Time{}, err
	}
	minute, err := strconv.Atoi(splitedTime[1])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(0, 1, 1, hour, minute, 0, 0, time.Local), nil
}
