package models

import (
	"time"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID              uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName        string     `json:"full_name" gorm:"size:100;not null"`
	Username        string     `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Email           string     `json:"email" gorm:"uniqueIndex;size:100;not null"`
	PasswordHash    string     `json:"-" gorm:"size:255;not null"` // "-" means don't include in JSON
	PhoneNumber     string     `json:"phone_number" gorm:"size:20;not null"`
	Role            string     `json:"role" gorm:"type:varchar(20);default:'user';check:role IN ('admin','user')"`
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	EmailVerified   bool       `json:"email_verified" gorm:"default:false"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" gorm:"default:null"`

	// OTP fields
	OTPCode      string     `json:"-" gorm:"size:10;default:null"`
	OTPExpiresAt *time.Time `json:"-" gorm:"default:null"`

	// Password reset token fields
	ResetToken          string     `json:"-" gorm:"size:255;default:null"`
	ResetTokenExpiresAt *time.Time `json:"-" gorm:"default:null"`
	CreatedAt           time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserLogin request
type UserLogin struct {
	Phone    string `json:"phone" binding:"required,phone"`
    Email      string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserRegister request
type UserRegister struct {
	FullName    string `json:"full_name" binding:"required,min=3,max=50"`
	Username    string `json:"username" binding:"required,min=3,max=50"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// UserResponse user data in response
type UserResponse struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName returns the table name for the User model
func (User) TableName() string {
	return "users"
}