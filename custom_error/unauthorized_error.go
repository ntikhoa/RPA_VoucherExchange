package custom_error

type UnauthorizedError struct {
	message string `default:"Unauthorized"`
}

func NewUnauthorizedError(message string) error {
	return &UnauthorizedError{
		message: message,
	}
}

func (e UnauthorizedError) Error() string {
	return e.message
}
