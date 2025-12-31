package main

import (
	"log"

	"github.com/Hdeee1/go-register-login-otp/internal/config"
	"github.com/Hdeee1/go-register-login-otp/internal/handlers"
	"github.com/Hdeee1/go-register-login-otp/internal/routes"
	"github.com/Hdeee1/go-register-login-otp/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()

	// Init service
	otpService := services.NewOTPService(config.DB)
	authService := services.NewAuthService(config.DB, otpService)
	userHandler := handlers.NewUserHandler()

	// Init handlers
	authHandler := handlers.NewAuthHandler(authService, otpService)

	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r, authHandler, userHandler)

	log.Println("server running on :8080")
	log.Fatal(r.Run(":8080"))
}
