package mailer

import (
	"os"

	"github.com/mailersend/mailersend-go"
)

type Mailer struct {
	ms   *mailersend.Mailersend
	from mailersend.Recipient
}

func New() *Mailer {
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	from := mailersend.From{
		Name:  "Piccolo",
		Email: os.Getenv("MAILERSEND_FROM_EMAIL"),
	}

	return &Mailer{
		ms,
		from,
	}
}
