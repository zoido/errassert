package errassert

// ErrorAssertion represents a single instance of the error assertion.
// If the error is not as expected, call returns an error.
//
// Zero value (nil) is valid and Assert nor Require will never fail. It can be looked at as
// "I don't care about the error" assertion.
type ErrorAssertion func(error) error

// TestingT defines the subset of testing.TB methods used by errassert.
type TestingT interface {
	Log(...interface{})
	Fail()
	FailNow()
	Helper()
}

// Assert checks if the error is as expected and calls t.Fail if not.
func (assertion ErrorAssertion) Assert(t TestingT, err error) {
	if assertion == nil {
		return
	}

	t.Helper()
	if tErr := assertion(err); tErr != nil {
		t.Log(tErr.Error())
		t.Fail()
	}
}

// Require checks if the error is as expected and calls t.FailNow if not.
func (assertion ErrorAssertion) Require(t TestingT, err error) {
	if assertion == nil {
		return
	}

	t.Helper()
	if tErr := assertion(err); tErr != nil {
		t.Log(tErr.Error())
		t.FailNow()
	}
}
