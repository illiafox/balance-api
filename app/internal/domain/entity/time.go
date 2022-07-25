package entity

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

const TimeLayout = time.RFC850

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	s := strconv.Quote(t.Format(TimeLayout))
	ptr := unsafe.Pointer(&s)
	slice := *(*[]byte)(ptr)
	return slice[:(*reflect.SliceHeader)(ptr).Len], nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := *(*string)(unsafe.Pointer(&data))

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
