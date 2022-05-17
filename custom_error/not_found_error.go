package custom_error

type NotFoundError struct {
	message string `default:"Not Found"`
}

func NewNotFoundError(message string) error {
	return &NotFoundError{
		message: message,
	}
}

func (e NotFoundError) Error() string {
	return e.message
}
