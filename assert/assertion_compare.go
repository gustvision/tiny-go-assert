package assert

import "fmt"

func Equal(expected interface{}, actual interface{}) error {
	if expected == actual {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected: %s	Actual: %s`, expected, actual)
}

func Nil(actual interface{}) error {
	if isNil(actual) {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected to be nil but got %s"`, actual)
}

func NotNil(actual interface{}) error {
	if isNil(actual) {
		return fmt.Errorf(`Test failed: Expected to be not nil, but got %s"`, actual)
	}
	return nil
}

func Empty(actual []interface{}) error {
	if len(actual) == 0 {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected to be empty, but got %s"`, actual)
}
