package custom_error

type BadRequestError struct {
	message string `default:"Bad Request"`
}

func NewBadRequestError(message string) error {
	return &BadRequestError{
		message: message,
	}
}

func (e *BadRequestError) Error() string {
	return e.message
}
