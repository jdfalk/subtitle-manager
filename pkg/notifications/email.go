// file: pkg/notifications/email.go
package notifications

import (
	"context"
	"fmt"
	"net/smtp"
)

// SMTPNotifier sends email using an SMTP server.
type SMTPNotifier struct {
	// Addr is the SMTP server address.
	Addr string
	// Auth provides optional authentication.
	Auth smtp.Auth
	// From is the sender address.
	From string
	// To holds destination email addresses.
	To []string
	// Send overrides the send function for testing.
	Send func(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

// Notify sends msg via SMTP.
func (s SMTPNotifier) Notify(ctx context.Context, msg string) error {
	send := s.Send
	if send == nil {
		send = smtp.SendMail
	}
	if s.Addr == "" || s.From == "" || len(s.To) == 0 {
		return fmt.Errorf("smtp configuration incomplete")
	}
	data := []byte("Subject: Subtitle Manager\r\n\r\n" + msg)
	return send(s.Addr, s.Auth, s.From, s.To, data)
}
