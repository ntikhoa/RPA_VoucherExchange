package custom_error

type ConflictError struct {
	message string `default:"Conflict"`
}

func NewConflictError(message string) error {
	return &ConflictError{
		message: message,
	}
}

func (e *ConflictError) Error() string {
	return e.message
}
