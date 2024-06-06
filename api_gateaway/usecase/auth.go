package usecase

type Login struct{}

type LoginInterface interface {
	Autentifikasi(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (a *Login) Autentifikasi(Username, Password string) bool {
	if Username == "admin" && Password == "admin123" {
		return true
	}
	return false
}
