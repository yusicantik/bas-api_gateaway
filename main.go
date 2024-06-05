package main

import (
	"api_gateaway/usecase"
	"fmt"
)

func main() {
	login := usecase.NewLogin()
	username := "admin"
	password := "admin123"

	if login.Autentifikasi(username, password) {
		fmt.Println("Login berhasil!")
	} else {
		fmt.Println("Login gagal!")
	}
}
