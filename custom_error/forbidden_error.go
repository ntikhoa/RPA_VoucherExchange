package custom_error

type ForbiddenError struct {
	message string `default:"Forbidden"`
}

func NewForbiddenError(message string) error {
	return &ForbiddenError{
		message: message,
	}
}

func (e *ForbiddenError) Error() string {
	return e.message
}
