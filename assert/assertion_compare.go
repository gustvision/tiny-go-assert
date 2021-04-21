package assert

import "fmt"

func Equal(e1 interface{}, e2 interface{}) error {
	if e1 == e2 {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected "%s" to be equals to "%s, but got %s"`, e1, e2, e1)
}

func Nil(e1 interface{}) error {
	if isNil(e1) {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected "%s" to be nil, but got %s"`, e1, e1)
}

func NotNil(e1 interface{}) error {
	if isNil(e1) {
		return fmt.Errorf(`Test failed: Expected "%s" to be nil, but got %s"`, e1, e1)
	}
	return nil
}

func Empty(e1 []interface{}) error {
	if len(e1) == 0 {
		return nil
	}
	return fmt.Errorf(`Test failed: Expected "%s" to be empty, but got %s"`, e1, e1)
}
