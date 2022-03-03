package dateRunner

import (
	"errors"
	"strconv"
	"time"
)

type AddDate struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func ParseDateText(timeText string) (time.Time, error) {
	var layout string
	if len(timeText) == 8 {
		layout = "20060102"
	} else if len(timeText) == 14 {
		layout = "20060102150405"
	} else {
		return time.Time{}, errors.New("time text is not correct")
	}

	t, err := time.Parse(layout, timeText)
	return t, err
}

func Add(tm time.Time, addDate AddDate) (time.Time, error) {
	var newTime time.Time
	newTime = tm

	if addDate.Year != 0 || addDate.Month != 0 || addDate.Day != 0 {
		newTime = newTime.AddDate(addDate.Year, addDate.Month, addDate.Day)
	}

	var durationDiff time.Duration
	var duration time.Duration
	duration = 0
	var err error

	durationDiff, err = time.ParseDuration(strconv.Itoa(addDate.Hour) + "h")
	if err != nil {
		return time.Time{}, err
	}
	duration += durationDiff

	durationDiff, err = time.ParseDuration(strconv.Itoa(addDate.Minute) + "m")
	if err != nil {
		return time.Time{}, err
	}
	duration += durationDiff

	durationDiff, err = time.ParseDuration(strconv.Itoa(addDate.Second) + "s")
	if err != nil {
		return time.Time{}, err
	}
	duration += durationDiff

	return newTime.Add(duration), nil
}
