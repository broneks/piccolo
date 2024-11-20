package mailer

import (
	"context"
	"fmt"
	"log/slog"
)

const RESET_PASSWORD_TEMPLATE_ID = "pxkjn416wzqgz781"

func (mail *Mailer) SendResetPassword(ctx context.Context, email, baseUrl, token string) error {
	if email == "" {
		return fmt.Errorf("Email is required")
	}
	if baseUrl == "" {
		return fmt.Errorf("baseUrl is required")
	}
	if token == "" {
		return fmt.Errorf("token is required")
	}

	// recipient := []mailersend.Recipient{
	// 	{
	// 		Email: email,
	// 		Name:  "",
	// 	},
	// }

	resetPasswordLink := fmt.Sprintf("%s?token=%s", baseUrl, token)

	// personalization := []mailersend.Personalization{
	// 	{
	// 		Email: email,
	// 		Data: map[string]any{
	// 			"email":             email,
	// 			"resetPasswordLink": resetPasswordLink,
	// 		},
	// 	},
	// }

	slog.Debug(resetPasswordLink)

	// _, err := mail.send(ctx, RESET_PASSWORD_TEMPLATE_ID, "Piccolo - Reset Password", recipient, personalization)
	// if err != nil {
	// 	return err
	// }

	return nil
}
