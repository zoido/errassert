package errassert_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/zoido/errassert"
)

func Example() {
	t := testing.T{} // Provided by the testing package.

	type testCase struct {
		in        string
		errassert errassert.ErrorAssertion
	}

	run := func(t *testing.T, tc testCase) {
		_, err := strconv.Atoi(tc.in)

		tc.errassert.Require(t, err)
	}

	testCases := map[string]testCase{
		"ok": {
			in:        "42",
			errassert: errassert.NilError(),
		},
		"invalid input fails": {
			in:        "invalid",
			errassert: errassert.SomeError(),
		},
		"empty input fails": {
			in:        "",
			errassert: errassert.ErrorEndsWith("invalid syntax"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func ExampleErrorAssertion_custom() {
	t := testing.T{} // Provided by the testing package.

	type testCase struct {
		in        string
		errassert errassert.ErrorAssertion
	}

	run := func(t *testing.T, tc testCase) {
		_, err := strconv.Atoi(tc.in)

		tc.errassert.Require(t, err)
	}

	testCases := map[string]testCase{
		"empty input fails": {
			in: "very specific error input",
			errassert: func(err error) error {
				if err == nil {
					return errors.New("expected error, got nil")
				}
				if err.Error() != "very specific error" {
					return fmt.Errorf("expected very specific error, got: '%v'", err.Error())
				}
				return nil
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func ExampleWant() {
	t := testing.T{} // Provided by the testing package.

	type testCase struct {
		in        string
		errassert errassert.ErrorAssertion
	}

	run := func(t *testing.T, tc testCase) {
		_, err := strconv.Atoi(tc.in)

		tc.errassert.Require(t, err)
	}

	testCases := map[string]testCase{
		"ok": {
			in:        "42",
			errassert: errassert.NilError(),
		},
		"invalid input": {
			in: "input",
			errassert: errassert.Want(
				errassert.ErrorContains("\"input\""),
				errassert.ErrorEndsWith("invalid syntax"),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}
