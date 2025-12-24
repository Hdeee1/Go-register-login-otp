package handlers

import (
	"net/http"

	"github.com/Hdeee1/go-register-login-otp/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	otpService *services.OTPService
}

func NewAuthHandler(otpSvc *services.OTPService) *AuthHandler {
	return &AuthHandler{otpService: otpSvc}
}

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *AuthHandler) RequestOTP(c *gin.Context) {
	var req OTPRequest
	
	// Input JSON validation
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service
	if err := h.otpService.SendOTP(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The OTP Code has ben successfully sent to your email"})
}