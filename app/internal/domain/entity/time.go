package entity

import (
	"strconv"
	"time"
)

type Time time.Time

const TimeLayout = `2006-01-02 15:04:05 MST`

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(
		strconv.Quote(time.Time(t).Format(TimeLayout)),
	), nil
}
