package mailer

import (
	"net/smtp"
	"os"
)

type Mailer struct {
	auth smtp.Auth
	host string
	port string
	from string
}

func New() *Mailer {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM_EMAIL")

	auth := smtp.PlainAuth("", from, password, host)

	return &Mailer{
		auth,
		host,
		port,
		from,
	}
}
