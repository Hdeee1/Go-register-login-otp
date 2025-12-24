package main

import (
	"log"

	"github.com/Hdeee1/go-register-login-otp/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()

	r := gin.Default()

	log.Fatal(r.Run(":8080"))
}
