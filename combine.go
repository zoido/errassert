package errassert

// Want combines multiple error assertions into a single assertion. All assertions
// must pass for the combined assertion to pass. If any of the assertions fails, the
// combined assertion fails and returns the first error encountered.
func Want(assertions ...ErrorAssertion) ErrorAssertion {
	return func(err error) error {
		for _, a := range assertions {
			if err := a(err); err != nil {
				return err
			}
		}
		return nil
	}
}
