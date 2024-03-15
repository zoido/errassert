package errassert_test

import (
	"errors"
	"testing"

	"github.com/zoido/errassert"
)

func TestWant(t *testing.T) {
	type testCase struct {
		assertions []errassert.ErrorAssertion
		want       error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.Want(tc.assertions...)(nil)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("Want(%v) = %v; want %v", tc.assertions, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"single pass": {
			assertions: []errassert.ErrorAssertion{passAssertion},
		},
		"single fail": {
			assertions: []errassert.ErrorAssertion{failAssertion},
			want:       errors.New(failMsg),
		},
		"multiple pass": {
			assertions: []errassert.ErrorAssertion{passAssertion, passAssertion},
		},
		"first fail": {
			assertions: []errassert.ErrorAssertion{failAssertion, passAssertion},
			want:       errors.New(failMsg),
		},
		"second fail": {
			assertions: []errassert.ErrorAssertion{passAssertion, failAssertion},
			want:       errors.New(failMsg),
		},
		"both fail": {
			assertions: []errassert.ErrorAssertion{
				func(err error) error { return errors.New("first") },
				func(err error) error { return errors.New("second") },
			},
			want: errors.New("first"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}
