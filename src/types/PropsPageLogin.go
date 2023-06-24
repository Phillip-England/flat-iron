package types

type PropsPageLogin struct {
	ErrLoginForm string
}

func NewPropsPageLogin(errLoginForm string) *PropsPageLogin {
	return &PropsPageLogin{
		ErrLoginForm: errLoginForm,
	}
}