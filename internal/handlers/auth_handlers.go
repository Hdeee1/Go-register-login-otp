package handlers

import (
	"net/http"

	"github.com/Hdeee1/go-register-login-otp/internal/models"
	"github.com/Hdeee1/go-register-login-otp/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
	otpService *services.OTPService
}

func NewAuthHandler(authSrv *services.AuthService, otpSrv *services.OTPService) *AuthHandler {
	return &AuthHandler{
		authService: authSrv,
		otpService: otpSrv,
	}
}

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *AuthHandler) RequestOTP(c *gin.Context) {
	var req OTPRequest
	
	// Input JSON validation
	if err := c.ShouldBindJSON(&req); err != nil {
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

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return user response
	response := models.UserResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully, please check your email for verification",
		"user": response,
	})
}

// Login
type LoginRequest struct {
	Email		string `json:"email" binding:"required,email"`
	Password 	string `json:"password", binding:"required,min=6"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := models.UserResponse {
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"user": response,
	})
}