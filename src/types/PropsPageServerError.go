package types

type PropsPageServerError struct {
	ServerError string
}

func NewPropsPageServerError(serverError string) *PropsPageServerError {
	return &PropsPageServerError{
		ServerError: serverError,
	}
}