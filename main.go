package main

import (
	"api_gateaway/usecase"
	"fmt"
)

func main() {
	login := usecase.NewLogin()

	auth := login.Autentifikasi("admin", "admin123")
	fmt.Println(auth)
}
