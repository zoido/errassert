package errassert_test

import (
	"errors"
	"testing"

	"github.com/zoido/errassert"
)

func TestSomeError(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.SomeError()(tc.in)

		// Then
		if !errEqual(got, tc.want) {
			t.Errorf("SomeError(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			in:   nil,
			want: errors.New("expected an error but got nil"),
		},
		"non-nil passes": {
			in:   errors.New("some error"),
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestNoError(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.NilError()(tc.in)

		// Then
		if !errEqual(got, tc.want) {
			t.Errorf("NoError(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil passes": {
			in:   nil,
			want: nil,
		},
		"non-nil fails": {
			in:   errors.New("some error"),
			want: errors.New("expected a nil error but got 'some error'"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func errEqual(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Error() == b.Error()
}
