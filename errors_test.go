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
		if !errorEq(got, tc.want) {
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
		if !errorEq(got, tc.want) {
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

func TestErrorIs(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	target := errors.New("expected error")

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.ErrorIs(target)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("ErrorIs(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			in:   nil,
			want: errors.New("expected error to be 'expected error' but got '<nil>'"),
		},
		"not matching fails": {
			in:   errors.New("another error"),
			want: errors.New("expected error to be 'expected error' but got 'another error'"),
		},
		"matching passes": {
			in:   target,
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestError(t *testing.T) {
	type testCase struct {
		msg  string
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.Error(tc.msg)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("Error(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			msg:  "expected error",
			in:   nil,
			want: errors.New("expected error to be 'expected error' but got '<nil>'"),
		},
		"not matching fails": {
			msg:  "expected error",
			in:   errors.New("another error"),
			want: errors.New("expected error to be 'expected error' but got 'another error'"),
		},
		"matching passes": {
			msg:  "expected error",
			in:   errors.New("expected error"),
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestErrorAs(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.ErrorAs(new(testError))(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("ErrorAs(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			in:   nil,
			want: errors.New("expected error to be '*errassert_test.testError' but got '<nil>'"),
		},
		"not matching fails": {
			in: errors.New("another error"),
			want: errors.New(
				"expected error to be '*errassert_test.testError' but got '*errors.errorString'",
			),
		},
		"matching passes": {
			in:   testError{},
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

type testError struct{}

func (testError) Error() string { return "test error" }

func TestErrorContains(t *testing.T) {
	type testCase struct {
		substring string
		in        error
		want      error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.ErrorContains(tc.substring)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("ErrorContains(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			substring: "some",
			in:        nil,
			want:      errors.New("expected an error containing 'some' but got nil"),
		},
		"not matching fails": {
			substring: "some",
			in:        errors.New("another error"),
			want:      errors.New("expected an error containing 'some' but got 'another error'"),
		},
		"matching passes": {
			substring: "contains",
			in:        errors.New("error that contains substring"),
			want:      nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestErrorStartsWith(t *testing.T) {
	type testCase struct {
		prefix string
		in     error
		want   error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.ErrorStartsWith(tc.prefix)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("ErrorStartsWith(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			prefix: "some",
			in:     nil,
			want:   errors.New("expected an error starting with 'some' but got nil"),
		},
		"not matching fails": {
			prefix: "some",
			in:     errors.New("another: error"),
			want:   errors.New("expected an error starting with 'some' but got 'another: error'"),
		},
		"matching passes": {
			prefix: "error: ",
			in:     errors.New("error: that starts with prefix"),
			want:   nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestErrorEndsWith(t *testing.T) {
	type testCase struct {
		suffix string
		in     error
		want   error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := errassert.ErrorEndsWith(tc.suffix)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("ErrorEndsWith(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"nil fails": {
			suffix: "some",
			in:     nil,
			want:   errors.New("expected an error ending with 'some' but got nil"),
		},
		"not matching fails": {
			suffix: "some",
			in:     errors.New("error: another"),
			want:   errors.New("expected an error ending with 'some' but got 'error: another'"),
		},
		"matching passes": {
			suffix: "ends with suffix",
			in:     errors.New("error: that ends with suffix"),
			want:   nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func errorEq(a, b error) bool {
	if b == nil {
		return a == b
	}
	return a.Error() == b.Error()
}
