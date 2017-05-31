package tools

import (
	"fmt"
	"time"
)

//https://golang.org/src/time/format.go

func TimeNowString() string {
	return time.Now().Format(time.RFC3339)
}

func TimeNowSQLString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeToSQLString(t time.Time) (string, error) {
	return t.Format("2006-01-02 15:04:05"), nil
}

func TimeToISOString(t time.Time) (string, error) {
	return t.Format("2006-01-02T15:04:05.000Z"), nil
}

func TimeFromSQLString(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, fmt.Errorf("Empty time")
	}
	layout := "2006-01-02 15:04:05"
	// str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeFromSQLStringWithZone(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, fmt.Errorf("Empty time")
	}
	layout := "2006-01-02 15:04:05 -0700"
	// str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeFromISOString(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, fmt.Errorf("Empty time")
	}
	//layout := "2006-01-02T15:04:05.000Z"
	layout := "2006-01-02T15:04:05-07:00" //moment js format
	if len(str) <= 10 {
		layout = "2006-01-02"
	}
	// str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeFromStringFormat(str, layout string) (time.Time, error) {
	if str == "" {
		return time.Time{}, fmt.Errorf("Empty time")
	}
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
