package services

type AuthService interface {
	Register()
	Login()
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Register() {

}

func (s *authService) Login() {

}
