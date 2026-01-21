package email

import "context"

// EmailSender defines the interface for sending emails
type EmailSender interface {
	Send(ctx context.Context, to, subject, body string) error
}
