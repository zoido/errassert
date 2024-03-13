package errassert

import (
	"errors"
	"fmt"
	"strings"
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

// ErrorContains returns an assertion that checks if the error contains the given substring.
func ErrorContains(substring string) ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return fmt.Errorf("expected an error containing '%s' but got nil", substring)
		}
		if !strings.Contains(err.Error(), substring) {
			return fmt.Errorf("expected an error containing '%s' but got '%v'", substring, err)
		}
		return nil
	}
}

// ErrorStartsWith returns an assertion that checks if the error starts with the given prefix.
func ErrorStartsWith(prefix string) ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return fmt.Errorf("expected an error starting with '%s' but got nil", prefix)
		}
		if !strings.HasPrefix(err.Error(), prefix) {
			return fmt.Errorf("expected an error starting with '%s' but got '%v'", prefix, err)
		}
		return nil
	}
}

// ErrorEndsWith returns an assertion that checks if the error ends with the given suffix.
func ErrorEndsWith(suffix string) ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return fmt.Errorf("expected an error ending with '%s' but got nil", suffix)
		}
		if !strings.HasSuffix(err.Error(), suffix) {
			return fmt.Errorf("expected an error ending with '%s' but got '%v'", suffix, err)
		}
		return nil
	}
}
