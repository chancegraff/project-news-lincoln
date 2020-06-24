package errors

// NilPointerError ...
type NilPointerError interface {
	Error() string
}

type nilPointerError struct{}

func (e *nilPointerError) Error() string {
	return "Invalid schema"
}

// NilPointer ...
func NilPointer() NilPointerError {
	return &nilPointerError{}
}

// MissingKeyError ...
type MissingKeyError interface {
	Error() string
}

type missingKeyError struct{}

func (e *missingKeyError) Error() string {
	return "Invalid request"
}

// MissingKey ...
func MissingKey() MissingKeyError {
	return &missingKeyError{}
}
