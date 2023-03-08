package utilities

import (
	"reflect"
)

// IsNil In Go, nil is the zero value for pointers, interfaces, maps, slices,
// channels and function types, representing an uninitialized value.
func IsNil(o interface{}) bool {
	if o == nil {
		return true
	}
	return false
}

func TypeOf(o interface{}) string {
	return reflect.TypeOf(o).String()
}
