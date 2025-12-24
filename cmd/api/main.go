package main

import (
	"log"

	"github.com/Hdeee1/go-register-login-otp/internal/config"
	"github.com/Hdeee1/go-register-login-otp/internal/services"
	"github.com/Hdeee1/go-register-login-otp/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()

	otpService := services.NewOTPService(config.DB)
	authHandler := handlers.NewAuthHandler(otpService)

	r := gin.Default()

	r.POST("api/auth/request-otp", authHandler.RequestOTP)

	log.Fatal(r.Run(":8080"))
}
