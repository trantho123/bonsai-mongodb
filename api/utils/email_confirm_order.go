package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type OrderEmailData struct {
	CustomerName    string
	OrderID         string
	Items           []OrderItem
	TotalAmount     int32
	ShippingAddress string
}

type OrderItem struct {
	Name     string
	Quantity int32
	Price    int32
	Total    int32
}

func SendOrderConfirmationEmail(to string, data OrderEmailData, config *EmailConfig) error {
	// Đọc template email
	tmpl, err := template.ParseFiles("templates/order_confirmation.html")
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	// Config SMTP
	from := config.SenderEmail
	password := config.SenderPass
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Tạo email message
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Order Confirmation\n"
	msg := []byte(subject + mime + body.String())

	// Gửi email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
