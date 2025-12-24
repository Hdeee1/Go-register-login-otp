package services

import (
	"errors"
	"time"

	"github.com/Hdeee1/go-register-login-otp/internal/models"
	"github.com/Hdeee1/go-register-login-otp/pkg/utils"
	"gorm.io/gorm"
)

type OTPService struct {
	DB *gorm.DB
}

func NewOTPService(db *gorm.DB) *OTPService {
	return &OTPService{DB: db}
} 

func (s *OTPService) SendOTP(email string) error {
	// Search user by email
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	
	// Generate OTP
	otpCode := utils.GenerateOTP(6)
	expAt := time.Now().Add(5 * time.Minute)

	// Update user with new otp
	user.OTPCode = otpCode
	user.OTPExpiresAt = &expAt
	
	if err := s.DB.Save(&user).Error; err != nil {
		return errors.New("Failed to save otp")
	}

	// Send email
	if err := utils.SendOTPEmail(user.Email, otpCode); err != nil {
		return errors.New("failed to send")
	}

	return nil
}

func (s *OTPService) VerifyOTP(email, otpCode string) error {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	// check exp OTP
	if user.OTPExpiresAt == nil || time.Now().After(*user.OTPExpiresAt) {
		return errors.New("OTP Expired")
	}

	// Check OTP match 
	if user.OTPCode != otpCode {
		return errors.New("Wrong OTP")
	}

	// Update user: verified & clear OTP
	user.EmailVerified = true
	user.OTPCode = ""
	user.OTPExpiresAt = nil

	if err := s.DB.Save(&user).Error; err != nil {
		return errors.New("Failed to update  user")
	}

	return nil
}