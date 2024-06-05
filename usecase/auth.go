package usecase

type Login struct{}

type LoginInterface interface {
	Autentifikasi(username, password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (a *Login) Autentifikasi(username, password string) bool {
	if username == "admin" && password == "admin123" {
		return true
	}
	return false
}
