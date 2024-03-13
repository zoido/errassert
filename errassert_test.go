package errassert_test

import (
	"testing"

	"github.com/zoido/errassert"
)

func NewMockT() *MockT {
	return &MockT{}
}

type MockT struct {
	failCalled    bool
	failNowCalled bool
	logCalled     bool
	logfMsg       string
}

var _ errassert.TestingT = (*MockT)(nil)

func (m *MockT) Helper()  {}
func (m *MockT) FailNow() { m.failNowCalled = true }
func (m *MockT) Fail()    { m.failCalled = true }

func (m *MockT) Log(args ...interface{}) {
	m.logCalled = true
	if len(args) == 1 {
		m.logfMsg = args[0].(string)
	}
}

func (m MockT) AssertFailed(t *testing.T) {
	t.Helper()
	if !m.failCalled {
		t.Error("expected Fail to be called but it was not")
	}
}

func (m MockT) AssertFailedNow(t *testing.T) {
	t.Helper()
	if !m.failNowCalled {
		t.Error("expected FailNow to be called but it was not")
	}
}

func (m MockT) AssertNotFailed(t *testing.T) {
	t.Helper()
	if m.failCalled {
		t.Error("expected FailNow to not be called but it was")
	}
}

func (m MockT) AssertNotFailedNow(t *testing.T) {
	t.Helper()
	if m.failNowCalled {
		t.Error("expected FailNow to not be called but it was")
	}
}

func (m MockT) AssertLogfCalled(t *testing.T) {
	t.Helper()
	if !m.logCalled {
		t.Error("expected Logf to be called but it was not")
	}
}

func (m MockT) AssertLogfNotCalled(t *testing.T) {
	t.Helper()
	if m.logCalled {
		t.Error("expected Logf to not be called but it was")
	}
}

func (m MockT) AssertLogfCalledWith(t *testing.T, expected string) {
	t.Helper()
	m.AssertLogfCalled(t)
	if m.logfMsg != expected {
		t.Errorf(
			"expected Logf to be called with %q but it was called with %q",
			expected,
			m.logfMsg,
		)
	}
}
