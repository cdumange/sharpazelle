package pointer

// To allows to create a pointer for a value
func To[T any](value T) *T {
	return &value
}
