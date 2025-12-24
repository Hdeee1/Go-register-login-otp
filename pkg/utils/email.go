package utils

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(email string, otp string) error {
	host := os.Getenv("EMAIL_HOST")
	portStr := os.Getenv("EMAIL_PORT")
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	from := os.Getenv("EMAIL_FROM")
	fromName := os.Getenv("EMAIL_FROM_NAME")

	// Conversion port to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 587 // default fallback
	}

	// Create Email Message
	m := gomail.NewMessage() 
	m.SetHeader("From", m.FormatAddress(from, fromName))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Kode OTP Login Anda")

	// Fill email to HTML
	htmlBody := "<h3>Hello!</h3>" +
			"This is you OTP for verification :" +
			"<h1>" + otp +"</h1>" +
			"<p>This code will expired after 5 minute</p>"
	m.SetBody("text/html", htmlBody)

	// Send Email
	d := gomail.NewDialer(host, port, user, password)
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
} 


