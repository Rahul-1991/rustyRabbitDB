package validations

// IsSlice checks if an interface{} value is of type slice.
func IsSlice(v interface{}) bool {
	_, ok := v.([]interface{})
	return ok
}
