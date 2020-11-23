package response

import (
	"time"
)

type Timestamp time.Time

type BaseResponse struct {
	Status     string `json:"status"`
	StatusCode int32  `json:"status code"`
	Message    string `json:"message"`
	Time       string `json:"timestamp"`
}

type Tabler interface {
	TableName() string
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return nil
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	layout := "2006-01-02 15:04:05"

	ts, err := time.Parse(layout, timeStr)
	*t = Timestamp(ts)

	return err
}
