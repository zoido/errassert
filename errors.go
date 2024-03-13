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

// ErrorIs returns an assertion that checks if the error passes errors.Is check.
func ErrorIs(expected error) ErrorAssertion {
	return func(err error) error {
		if !errors.Is(err, expected) {
			return fmt.Errorf("expected error to be '%v' but got '%v'", expected, err)
		}
		return nil
	}
}

// ErrorAs returns an assertion that checks if the error passes errors.As check.
func ErrorAs(target interface{}) ErrorAssertion {
	return func(err error) error {
		if !errors.As(err, target) {
			return fmt.Errorf("expected error to be '%T' but got '%T'", target, err)
		}
		return nil
	}
}
