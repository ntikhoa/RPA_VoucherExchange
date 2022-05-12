package custom_error

type AuthorizedError struct{}

func (e *AuthorizedError) Error() string {
	return "permission denied"
}
