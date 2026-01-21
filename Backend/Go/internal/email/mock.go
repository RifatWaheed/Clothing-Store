package email

import (
	"context"
	"fmt"
)

// MockEmailSender is a mock implementation for testing
type MockEmailSender struct{}

// NewMockEmailSender creates a new mock email sender
func NewMockEmailSender() *MockEmailSender {
	return &MockEmailSender{}
}

// Send logs the email details (mock implementation)
func (m *MockEmailSender) Send(ctx context.Context, to, subject, body string) error {
	fmt.Println("=== MOCK EMAIL SENDER ===")
	fmt.Println("To:", to)
	fmt.Println("Subject:", subject)
	fmt.Println("Body:", body)
	fmt.Println("========================")
	return nil
}
