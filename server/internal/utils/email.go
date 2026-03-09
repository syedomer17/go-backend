package utils

import (
	"task-manager/internal/config"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, otp string) error {

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EMAIL_FROM)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify your email")

	body := "Your OTP for email verification is: " + otp
	m.SetBody("text/plain", body)
	
	d := gomail.NewDialer(
		cfg.SMTP_HOST,
		cfg.SMTP_PORT,
		cfg.SMTP_USER,
		cfg.SMTP_PASS,
	)

	return d.DialAndSend(m)
}