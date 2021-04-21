package assert

import "reflect"

func isNil(obj interface{}) bool {
	if obj == nil {
		return true
	}

	value := reflect.ValueOf(obj)

	if value.IsNil() {
		return true
	}
	return false
}
