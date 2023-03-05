package utilities

// IsNil In Go, nil is the zero value for pointers, interfaces, maps, slices,
//channels and function types, representing an uninitialized value.
func IsNil(o interface{}) bool {
	if o == nil {
		return true
	}
	return false
}
