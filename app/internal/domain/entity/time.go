package entity

import (
	"fmt"
	"strconv"
	"time"

	"balance-service/app/pkg/unsafe"
)

const TimeLayout = time.RFC850

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	s := strconv.Quote(t.Format(TimeLayout))

	return unsafe.StringToBytes(&s), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := unsafe.BytesToString(data)

	// unquote data: "time" -> time
	var err error
	s, err = strconv.Unquote(s)
	if err != nil {
		return fmt.Errorf("unquote: %w", err)
	}

	// parse time
	t.Time, err = time.Parse(TimeLayout, s)
	if err != nil {
		return fmt.Errorf("parse time: %w", err)
	}

	return nil
}
