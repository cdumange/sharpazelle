package commonerrors

// Error is a custom error type to allow const error.
type Error string

// Error implements error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNotToBeTreated is returned when a folder does not need a rules to be generated.
	ErrNotToBeTreated = Error("folder should not be treated")
)
