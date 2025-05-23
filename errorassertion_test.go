package errassert_test

import (
	"errors"
	"testing"

	"github.com/zoido/errassert"
)

var anError = errors.New("an error") //nolint:staticcheck // So far this is better name for our purpose then errSomething.

func TestErrorAssertion_Assert_Pass(t *testing.T) {
	// Given
	mockT := NewMockT()
	var assertion errassert.ErrorAssertion = passAssertion

	// When
	assertion.Assert(mockT, anError)

	// Then
	mockT.AssertNotFailed(t)
	mockT.AssertNotFailedNow(t)
	mockT.AssertLogfNotCalled(t)
}

func TestErrorAssertion_Assert_Fail(t *testing.T) {
	// Given
	mockT := NewMockT()
	var assertion errassert.ErrorAssertion = failAssertion

	// When
	assertion.Assert(mockT, anError)

	// Then
	mockT.AssertFailed(t)
	mockT.AssertNotFailedNow(t)
	mockT.AssertLogfCalledWith(t, failMsg)
}

func TestErrorAssertion_Require_Pass(t *testing.T) {
	// Given
	mockT := NewMockT()
	var assertion errassert.ErrorAssertion = passAssertion

	// When
	assertion.Require(mockT, anError)

	// Then
	mockT.AssertNotFailed(t)
	mockT.AssertNotFailedNow(t)
	mockT.AssertLogfNotCalled(t)
}

func TestErrorAssertion_Require_Fail(t *testing.T) {
	// Given
	mockT := NewMockT()
	var assertion errassert.ErrorAssertion = failAssertion

	// When
	assertion.Require(mockT, anError)

	// Then
	mockT.AssertFailedNow(t)
	mockT.AssertNotFailed(t)
	mockT.AssertLogfCalled(t)
	mockT.AssertLogfCalledWith(t, failMsg)
}

func TestErrorAssertion_NilFunctional(t *testing.T) {
	// Given
	mockT := NewMockT()
	var assertion errassert.ErrorAssertion

	// When
	assertion.Assert(mockT, anError)
	assertion.Require(mockT, anError)

	// Then
	mockT.AssertNotFailed(t)
	mockT.AssertNotFailedNow(t)
	mockT.AssertLogfNotCalled(t)
}

func passAssertion(error) error {
	return nil
}

const failMsg = "fail assertion error"

func failAssertion(error) error {
	return errors.New(failMsg)
}
