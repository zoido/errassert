package grpcerrassert_test

import (
	"errors"
	"testing"

	"github.com/zoido/errassert/grpcerrassert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSomeStatus(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := grpcerrassert.SomeStatus()(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("SomeStatus(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"status passes": {
			in:   status.Error(codes.Unimplemented, "some error"),
			want: nil,
		},
		"nil fails": {
			in:   nil,
			want: errors.New("expected a status but got nil"),
		},
		"non-status fails": {
			in:   errors.New("some error"),
			want: errors.New("expected a status but got a non-status error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestNonStatus(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := grpcerrassert.NonStatus()(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("NonStatus(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"status fails": {
			in:   status.Error(codes.Unimplemented, "some error"),
			want: errors.New("expected a non-status error but got a status"),
		},
		"nil fails": {
			in:   nil,
			want: errors.New("expected a non-status error but got nil"),
		},
		"non-status passes": {
			in:   errors.New("some error"),
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}

func TestStatusCode(t *testing.T) {
	type testCase struct {
		in   error
		want error
	}

	run := func(t *testing.T, tc testCase) {
		// When
		got := grpcerrassert.StatusCode(codes.Unimplemented)(tc.in)

		// Then
		if !errorEq(got, tc.want) {
			t.Errorf("StatusCode(%v) = %v; want %v", tc.in, got, tc.want)
		}
	}

	testCases := map[string]testCase{
		"status with correct code passes": {
			in:   status.Error(codes.Unimplemented, "some error"),
			want: nil,
		},
		"non-status fails": {
			in:   errors.New("some error"),
			want: errors.New("expected a status but got a non-status error"),
		},
		"status with wrong code fails": {
			in:   status.Error(codes.Unavailable, "some error"),
			want: errors.New("expected status code to be 'Unimplemented' but got 'Unavailable'"),
		},
		"nil fails": {
			in:   nil,
			want: errors.New("expected status code to be 'Unimplemented' but got 'OK'"),
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
