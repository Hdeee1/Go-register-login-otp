package main

import (
	"log"

	"github.com/Hdeee1/go-register-login-otp/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Starting the application...")
	
	log.Println("Initializing database...")
	config.InitDatabase()

	r := gin.Default()

	r.Run(":8080")

}
