package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type Claims struct {
	UserID 	uint	`json:"user_id"`
	Email 	string	`json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string) (string, error) {
	expiresIn := os.Getenv("JWT_EXPIRES_IN")

	duration, err := time.ParseDuration(expiresIn); 
	if err != nil {
		duration = 24 * time.Hour
	}
	
	claims := Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	} 

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	// get and validation secret key
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET not set")
	}

	// Parse the token with jwt.ParseWithClaims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// validation claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// return claims
	return claims, nil
}