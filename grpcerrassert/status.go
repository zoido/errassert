// Package grpcerrassert provides assertions for gRPC status errors.
package grpcerrassert

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zoido/errassert"
)

// SomeStatus returns an assertion that checks if the error is a gRPC status.
func SomeStatus() errassert.ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return errors.New("expected a status but got nil")
		}
		if _, isStatus := status.FromError(err); !isStatus {
			return errors.New("expected a status but got a non-status error")
		}
		return nil
	}
}

// NonStatus returns an assertion that checks if the error is not a gRPC status.
func NonStatus() errassert.ErrorAssertion {
	return func(err error) error {
		if err == nil {
			return errors.New("expected a non-status error but got nil")
		}
		if _, isStatus := status.FromError(err); isStatus {
			return errors.New("expected a non-status error but got a status")
		}
		return nil
	}
}

// StatusCode returns an assertion that checks if the status code matches the expected one.
func StatusCode(expectedCode codes.Code) errassert.ErrorAssertion {
	return func(err error) error {
		st, isStatus := status.FromError(err)
		if !isStatus {
			return errors.New("expected a status but got a non-status error")
		}
		if st.Code() != expectedCode {
			return fmt.Errorf(wrongCodeFmt, expectedCode, st.Code())
		}
		return nil
	}
}

const wrongCodeFmt = "expected status code to be '%s' but got '%s'"
