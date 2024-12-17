package utils

import (
	"errors"
	"fmt"
	"net/smtp"
	"regexp"
	"time"
)

func IsEmailValid(email string) error {
	// Regular expression for validating an email address
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func IsValidDOB(dob string) error {
	// Định dạng mà chúng ta mong muốn
	layout := "2006-01-02" // YYYY-MM-DD

	// Phân tích chuỗi theo định dạng
	_, err := time.Parse(layout, dob)
	if err != nil {
		return errors.New("invalid date format, must be YYYY-MM-DD")
	}

	return nil
}

type EmailConfig struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}

func SendVerificationEmail(toEmail string, verificationCode string, config *EmailConfig) error {
	// Cấu hình SMTP
	auth := smtp.PlainAuth("", config.SenderEmail, config.SenderPass, config.SMTPHost)

	// Tạo nội dung email
	subject := "Email Verification"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Email Verification</h2>
			<p>Please click the link below to verify your email:</p>
			<a href="http://localhost:3000/verify/%s">Verify Email</a>
			<p>Or enter this verification code: %s</p>
			<p>This code will expire in 15 minutes.</p>
		</body>
		</html>
	`, verificationCode, verificationCode)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := fmt.Sprintf("Subject: %s\n%s\n%s", subject, mime, body)

	// Gửi email
	err := smtp.SendMail(
		config.SMTPHost+":"+config.SMTPPort,
		auth,
		config.SenderEmail,
		[]string{toEmail},
		[]byte(msg),
	)

	return err
}

type ResetPasswordEmailData struct {
	Username  string
	ResetLink string
}

func SendResetPasswordEmail(toEmail string, data ResetPasswordEmailData, config *EmailConfig) error {
	// Cấu hình SMTP
	auth := smtp.PlainAuth("", config.SenderEmail, config.SenderPass, config.SMTPHost)

	// Tạo nội dung email
	subject := "Reset Password"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Reset Your Password</h2>
			<p>Dear %s,</p>
			<p>We received a request to reset your password. Click the link below to set a new password:</p>
			<p><a href="%s" style="padding: 10px 20px; background-color: #1976d2; color: white; text-decoration: none; border-radius: 5px;">Reset Password</a></p>
			<p>If you didn't request this, you can safely ignore this email.</p>
			<p>This link will expire in 15 minutes.</p>
			<p>Best regards,<br>Shop It Team</p>
		</body>
		</html>
	`, data.Username, data.ResetLink)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := fmt.Sprintf("Subject: %s\n%s\n%s", subject, mime, body)

	// Gửi email
	err := smtp.SendMail(
		config.SMTPHost+":"+config.SMTPPort,
		auth,
		config.SenderEmail,
		[]string{toEmail},
		[]byte(msg),
	)

	return err
}
