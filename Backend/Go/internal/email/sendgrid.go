package email

import (
	"context"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridClient struct {
	client *sendgrid.Client
	from   string
}

// NewSendGridClient creates a new SendGrid email sender
func NewSendGridClient(apiKey, fromEmail string) *SendGridClient {
	return &SendGridClient{
		client: sendgrid.NewSendClient(apiKey),
		from:   fromEmail,
	}
}

// Send sends an email via SendGrid
func (s *SendGridClient) Send(ctx context.Context, to, subject, body string) error {
	from := mail.NewEmail("Clothing Store", s.from)
	toEmail := mail.NewEmail("", to)
	plainTextContent := body
	htmlContent := body

	m := mail.NewSingleEmail(from, subject, toEmail, plainTextContent, htmlContent)

	_, err := s.client.SendWithContext(ctx, m)
	return err
}
