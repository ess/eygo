package eygo

// Error is a specific error type for errors returned from the Engine Yard
// API.
type Error struct {
	ErrorString string
}

// Error is the full error string from the API for a given operation.
func (err *Error) Error() string {
	return err.ErrorString
}

// NewError instantiates a Error.
func NewError(errorString string) *Error {
	return &Error{
		ErrorString: errorString,
	}
}
