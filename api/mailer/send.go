package mailer

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/mailersend/mailersend-go"
)

func (mail *Mailer) send(ctx context.Context, templateId, subject string, recipients []mailersend.Recipient, personalizations []mailersend.Personalization) (string, error) {
	if templateId == "" {
		return "", fmt.Errorf("template id is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	message := mail.ms.Email.NewMessage()

	message.SetFrom(mail.from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID(templateId)
	message.SetPersonalization(personalizations)

	res, err := mail.ms.Email.Send(ctx, message)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	messageId := res.Header.Get("x-message-id")

	return messageId, nil
}
