package services

import (
	"errors"
	"fmt"
	

	"github.com/Hdeee1/go-register-login-otp/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
	otpService *OTPService
}

func NewAuthService(db *gorm.DB, otpSrv *OTPService,) *AuthService {
	return &AuthService{
		DB: db,
		otpService: otpSrv,
	}
}

func (s *AuthService) Register(input models.UserRegister) (*models.User, error) {
	var existingUser models.User
	// Check email exist
	if err := s.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already register")
	}

	// Check username exist
	if err := s.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already used")
	}

	// Check phone exist
	if err := s.DB.Where("phone_number = ?", input.PhoneNumber).First(&existingUser).Error; err == nil {
		return nil, errors.New("phone number already register")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create new user
	user := models.User{
		FullName: input.FullName,
		Username: input.Username,
		Email: input.Email,
		PasswordHash: string(hashedPassword),
		PhoneNumber: input.PhoneNumber,
		Role: "user",
		IsActive: true,
		EmailVerified: false,
	}

	// save to db
	if err := s.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to save user")
	}

	// Auto send otp
	if err := s.otpService.SendOTP(input.Email); err != nil {
		fmt.Printf("failed to send otp, err: %v", err)
	}

	return &user, nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	// search user by email
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("wrong email or password")
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("wrong email or password")
	}

	// check email verified
	if !user.EmailVerified {
		return nil, errors.New("email has not been verified, please re-register")
	}

	if !user.IsActive {
		return nil, errors.New("your account is nonactive")
	}

	return &user, nil
}