package usecase

import (
	"api_gateaway/models"
	"api_gateaway/utils"
)

type Login struct{}

type LoginInterface interface {
	Autentifikasi(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (a *Login) Autentifikasi(Username, Password string) bool {
	bodyPayloadLogin := models.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	orm.Find(&bodyPayloadLogin, "username = ? AND password = ? ", Username, Password)
	if bodyPayloadLogin.AccountID == "" {
		return false
	}

	return true
}
