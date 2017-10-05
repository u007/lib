package tools

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

var DateLayout = "2006-01-02"

var nilTime = (time.Time{}).UnixNano()

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return
	}
	d.Time, err = time.Parse(DateLayout, s)
	return
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(DateLayout))), nil
}

func (d *Date) IsSet() bool {
	return d.UnixNano() != nilTime
}

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
