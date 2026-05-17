package email

import (
	"context"
	"fmt"

	"github.com/resend/resend-go/v3"
)

type ResendEmailSender struct {
	client   *resend.Client
	fromAddr string
}

func NewResendClient(apiKey, fromAddr string) *ResendEmailSender {
	return &ResendEmailSender{
		client:   resend.NewClient(apiKey),
		fromAddr: fromAddr,
	}
}

func (r *ResendEmailSender) Send(ctx context.Context, to, subject, html string) error {
	params := &resend.SendEmailRequest{
		From:    r.fromAddr,
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}
	_, err := r.client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("resend: %w", err)
	}
	return nil
}
