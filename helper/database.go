package helper

import (
	"errors"
	"fmt"
	"time"
)

// Date ...
type Date time.Time

// UnmarshalJSON ...
func (t *Date) UnmarshalJSON(data []byte) error {
	layout := "2006-01-02"

	ts, err := validateDataTime(layout, data)
	*t = Date(ts)

	return err
}

// UnmarshalParam ...
func (t *Date) UnmarshalParam(src string) error {
	layout := "2006-01-02"

	ts, err := time.Parse(layout, src)
	*t = Date(ts)
	return err
}

// MarshalJSON ...
func (t Date) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

// String ...
func (t *Date) String() string {
	tm := time.Time(*t)
	return tm.Format("2006-01-02")
}

func validateDataTime(layout string, data []byte) (time.Time, error) {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return time.Time{}, errors.New("Data is Null")
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	ts, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return ts, nil
}
