package handlers

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(name, email, message string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	toEmail := os.Getenv("TO_EMAIL")

	smtpServer := smtpHost + ":" + smtpPort

	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPassword,
		smtpHost,
	)

	subject := "Customer Message"

	body := fmt.Sprintf(
		"From: %s\nEmail: %s\n\nMessage:\n%s",
		name, email, message,
	)

	emailMessage := []byte(
		"From: " + smtpUser + "\r\n" +
			"To: " + toEmail + "\r\n" +
			"Reply-To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\"\r\n" + body)

	err := smtp.SendMail(
		smtpServer,
		auth,
		smtpUser,
		[]string{toEmail},
		emailMessage,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
