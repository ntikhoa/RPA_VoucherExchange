package custom_error

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "data not found"
}
