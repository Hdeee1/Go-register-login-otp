package controller

import (
	"net/http"
	"time"

	"github.com/Hdeee1/go-register-login-otp/internal/config"
	"github.com/Hdeee1/go-register-login-otp/internal/models"
	"github.com/Hdeee1/go-register-login-otp/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func RequestOTP(c *gin.Context) {
	var input OTPRequest
	
	// Input JSON validation
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search user in database
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The User with that email was not found"})
		return
	}

	// Generate new OTP & exp time
	otpCode := utils.GenerateOTP(6)
	expirationTime := time.Now().Add(5 * time.Minute)

	// Update Database (Save OTP & ExpAt)
	user.OTPCode = otpCode
	user.OTPExpiresAt = &expirationTime

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save OTP"})
		return
	}

	// Send Email
	err := utils.SendOTPEmail(user.Email, otpCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The OTP Code has ben successfully sent to your email"})
}