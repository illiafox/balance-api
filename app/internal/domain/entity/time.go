package entity

import (
	"strconv"
	"time"
)

type Time time.Time

const TimeLayout = time.RFC822

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(
		strconv.Quote(time.Time(t).Format(TimeLayout)),
	), nil
}
