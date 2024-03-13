package errassert

import (
	"errors"
	"fmt"
)

// SomeError returns an assertion that checks if the error is not nil.
func SomeError() ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return errors.New("expected an error but got nil")
		}
		return nil
	}
}

// NilError returns an assertion that checks if the error is nil.
func NilError() ErrorAssertion {
	return func(err error) error {
		if err != nil {
			return fmt.Errorf("expected a nil error but got '%s'", err)
		}
		return nil
	}
}
