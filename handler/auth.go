package handler

import (
	"api_gateaway/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Login(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type BodyPayloadAuth struct {
	Username string
	Password string
}

func (a *authImplement) Login(g *gin.Context) {

	bodyPayloadAuth := BodyPayloadAuth{}
	err := g.BindJSON(&bodyPayloadAuth)
	if err != nil {
		log.Panic(err)
	}

	if usecase.NewLogin().Autentifikasi(bodyPayloadAuth.Username, bodyPayloadAuth.Password) {
		g.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayloadAuth,
		})

	} else {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Anda gagal login",
			"data":    err,
		})
	}
}
