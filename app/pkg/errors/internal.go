package errors

import (
	"fmt"
)

type InternalError struct {
	Err   error
	Scope string
}

func (i InternalError) Error() string {
	return fmt.Sprintf("%s: %s", i.Scope, i.Err)
}

func (i InternalError) Unwrap() error {
	return i.Err
}

func (i InternalError) Cause() error {
	return i.Err
}

func ToInternal(err error) (internal InternalError, ok bool) {
	//nolint:errorlint
	internal, ok = err.(InternalError)

	return
}

func NewInternal(err error, scope string) error {
	return InternalError{
		Err:   err,
		Scope: scope,
	}
}
