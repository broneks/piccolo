package mailer

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/mailersend/mailersend-go"
)

const RESET_PASSWORD_TEMPLATE_ID = "pxkjn416wzqgz781"
const RESET_PASSWORD_ROUTE = "/reset-password"

func (mail *Mailer) SendResetPassword(ctx context.Context, email, token string) error {
	appUrl := os.Getenv("APP_BASE_URL")

	if email == "" {
		return fmt.Errorf("Email is required")
	}
	if token == "" {
		return fmt.Errorf("token is required")
	}

	recipient := []mailersend.Recipient{
		{
			Email: email,
			Name:  "",
		},
	}

	resetPasswordLink := fmt.Sprintf("%s?token=%s", path.Join(appUrl, RESET_PASSWORD_ROUTE), token)

	personalization := []mailersend.Personalization{
		{
			Email: email,
			Data: map[string]any{
				"email":             email,
				"resetPasswordLink": resetPasswordLink,
			},
		},
	}

	_, err := mail.send(ctx, RESET_PASSWORD_TEMPLATE_ID, "Reset Password", recipient, personalization)
	if err != nil {
		return err
	}

	return nil
}
