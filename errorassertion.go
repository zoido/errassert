package errassert

// ErrorAssertion represents single single instance of the error assertion.
// If the error is not as expected, call returns an error.
type ErrorAssertion func(error) error

// TestingT defines the subset of testing.TB methods used by errassert.
type TestingT interface {
	Log(...interface{})
	Fail()
	FailNow()
	Helper()
}

// Assert checks if the error is as expected and calls t.Errorf if not.
func (assertion ErrorAssertion) Assert(t TestingT, err error, _ ...interface{}) {
	t.Helper()
	if tErr := assertion(err); tErr != nil {
		t.Log(tErr.Error())
		t.Fail()
	}
}

// Require checks if the error is as expected and calls t.FailNow if not.
func (assertion ErrorAssertion) Require(t TestingT, err error, _ ...interface{}) {
	t.Helper()
	if tErr := assertion(err); tErr != nil {
		t.Log(tErr.Error())
		t.FailNow()
	}
}
