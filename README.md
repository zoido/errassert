# `errassert`: Simple table driven Go tests error assertions

[![go reference](https://pkg.go.dev/badge/github.com/zoido/errassert)](https://pkg.go.dev/github.com/zoido/errassert)
[![licence](https://img.shields.io/github/license/zoido/errassert?style=flat-square)](https://github.com/zoido/errassert/blob/master/LICENSE)
![CI](https://img.shields.io/github/actions/workflow/status/zoido/errassert/go.yaml?style=flat-square&logoColor=white&logo=github)
[![coverage](https://img.shields.io/codecov/c/github/zoido/errassert?style=flat-square&logoColor=white&logo=codecov)](https://codecov.io/gh/zoido/errassert)
[![go report](https://goreportcard.com/badge/github.com/zoido/errassert?style=flat-square)](https://goreportcard.com/report/github.com/zoido/errassert)

## Overview

- set error assertion in a table driven test declaration
- provides similar experience as Testify's [`ErrorAssertionFunc`](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorAssertionFunc)
- define custom error assertions

## Example

<!-- markdownlint-disable MD010 MD013 -->

```go
func Test(t *testing.T) {
	type testCase struct {
		in        string
		want      int
		errassert errassert.ErrorAssertion
	}

	run := func(t *testing.T, tc testCase) {
		_, err := strconv.Atoi(tc.in)

		tc.errassert.Require(t, err)
	}

	testCases := map[string]testCase{
		"ok": {
			in:        "42",
			want:      42,
			errassert: errassert.NilError(),
		},
		"invalid input fails": {
			in:        "invalid input",
			errassert: errassert.SomeError(),
		},
		"empty input cause invalid syntax error": {
			in:        "",
			// Common basic assertions.
			errassert: errassert.ErrorEndsWith("invalid syntax"),
		},
		"invalid input fails with input": {
			in: "input",
			// Combine basic assertions.
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


```
<!-- markdownlint-enable MD010 MD013 -->
