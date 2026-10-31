# `grpcerrassert`: gRPC status error assertions for `errassert`

[![go reference](https://pkg.go.dev/badge/github.com/zoido/errassert/grpcerrassert)](https://pkg.go.dev/github.com/zoido/errassert/grpcerrassert)

This is a separate Go module, so the core [`errassert`](..) module does not
depend on gRPC.

## Example

<!-- markdownlint-disable MD010 MD013 -->

```go
package example_test

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zoido/errassert"
	"github.com/zoido/errassert/grpcerrassert"
)

func GetUser(id string) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "empty user ID")
	}
	if id != "42" {
		return status.Errorf(codes.NotFound, "user %q not found", id)
	}
	return nil
}

func TestGetUser(t *testing.T) {
	type testCase struct {
		id        string
		errassert errassert.ErrorAssertion
	}

	run := func(t *testing.T, tc testCase) {
		err := GetUser(tc.id)

		tc.errassert.Require(t, err)
	}

	testCases := map[string]testCase{
		"ok": {
			id:        "42",
			errassert: errassert.NilError(),
		},
		"empty ID is invalid argument": {
			id:        "",
			errassert: grpcerrassert.StatusCode(codes.InvalidArgument),
		},
		"unknown user is not found": {
			id: "7",
			errassert: errassert.Want(
				grpcerrassert.StatusCode(codes.NotFound),
				errassert.ErrorContains(`"7"`),
			),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) { run(t, tc) })
	}
}
```

<!-- markdownlint-enable MD010 MD013 -->
