package mailer

import (
	"log"
	"net/smtp"
	"os"
)

type Mailer struct {
	From      string
	smtpUrl   string
	plainAuth smtp.Auth
}

func NewMailer() *Mailer {
	m := &Mailer{
		From: os.Getenv("SMTP_FROM"),
	}
	m.smtpAuthentication()

	return m
}

func (m *Mailer) smtpAuthentication() {
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	m.smtpUrl = smtpHost + ":" + smtpPort
	m.plainAuth = smtp.PlainAuth("", username, password, smtpHost)
}

func (m *Mailer) SendEmail(r *Request) {
	r.SetHeaders()
	message := r.Headers + r.Body
	err := smtp.SendMail(m.smtpUrl, m.plainAuth, m.From, r.To, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Email successfully sent to %s", r.To)
}