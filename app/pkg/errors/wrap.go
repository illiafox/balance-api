package errors

import "fmt"

func (i InternalError) Wrap(scope string) error {
	if i.Scope == "" {
		i.Scope = scope
	} else {
		i.Scope = fmt.Sprintf("%s: %s", scope, i.Scope)
	}
	return i
}

func Wrap(err error, scope string) error {
	return fmt.Errorf("%s: %w", scope, err)
}
