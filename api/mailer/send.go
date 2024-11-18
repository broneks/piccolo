package mailer

import (
	"fmt"
	"net/smtp"
)

func (m *Mailer) Send(to, message string) error {
	if to == "" {
		return fmt.Errorf("to is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	body := []byte(message)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", m.host, m.port), m.auth, m.from, []string{to}, body)
	if err != nil {
		return err
	}

	return nil
}
