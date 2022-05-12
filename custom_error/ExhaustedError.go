package custom_error

type ExhaustedError struct{}

func (e *ExhaustedError) Error() string {
	return "data exhausted"
}
