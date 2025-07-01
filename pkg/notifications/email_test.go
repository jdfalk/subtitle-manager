// file: pkg/notifications/email_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174012

package notifications

import (
	"context"
	"errors"
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMTPNotifier_Notify_Success(t *testing.T) {
	var capturedAddr string
	var capturedAuth smtp.Auth
	var capturedFrom string
	var capturedTo []string
	var capturedMsg []byte
	
	mockSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		capturedAddr = addr
		capturedAuth = a
		capturedFrom = from
		capturedTo = to
		capturedMsg = msg
		return nil
	}
	
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   []string{"recipient1@example.com", "recipient2@example.com"},
		Send: mockSend,
	}
	
	testMessage := "Test email notification"
	err := notifier.Notify(context.Background(), testMessage)
	
	assert.NoError(t, err)
	assert.Equal(t, "smtp.example.com:587", capturedAddr)
	assert.Equal(t, notifier.Auth, capturedAuth)
	assert.Equal(t, "sender@example.com", capturedFrom)
	assert.Equal(t, []string{"recipient1@example.com", "recipient2@example.com"}, capturedTo)
	assert.Contains(t, string(capturedMsg), testMessage)
	assert.Contains(t, string(capturedMsg), "Subject: Subtitle Manager")
}

func TestSMTPNotifier_Notify_MissingAddr(t *testing.T) {
	notifier := SMTPNotifier{
		Addr: "", // Missing
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smtp configuration incomplete")
}

func TestSMTPNotifier_Notify_MissingFrom(t *testing.T) {
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "", // Missing
		To:   []string{"recipient@example.com"},
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smtp configuration incomplete")
}

func TestSMTPNotifier_Notify_MissingTo(t *testing.T) {
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   []string{}, // Empty
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smtp configuration incomplete")
}

func TestSMTPNotifier_Notify_NilTo(t *testing.T) {
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   nil, // Nil
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smtp configuration incomplete")
}

func TestSMTPNotifier_Notify_SendError(t *testing.T) {
	mockSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return errors.New("SMTP send failed")
	}
	
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
		Send: mockSend,
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "SMTP send failed")
}

func TestSMTPNotifier_Notify_NoAuth(t *testing.T) {
	var capturedAuth smtp.Auth
	
	mockSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		capturedAuth = a
		return nil
	}
	
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: nil, // No authentication
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
		Send: mockSend,
	}
	
	err := notifier.Notify(context.Background(), "test message")
	
	assert.NoError(t, err)
	assert.Nil(t, capturedAuth)
}

func TestSMTPNotifier_Notify_DefaultSend(t *testing.T) {
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		Auth: smtp.PlainAuth("", "user", "pass", "smtp.example.com"),
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
		Send: nil, // Use default smtp.SendMail
	}
	
	// This will likely fail with a network error, but shouldn't panic
	err := notifier.Notify(context.Background(), "test message")
	
	// We expect an error (network failure), but it should not panic
	assert.NotPanics(t, func() {
		notifier.Notify(context.Background(), "test message")
	})
	
	// The error should be a network-related error, not a configuration error
	if err != nil {
		assert.NotContains(t, err.Error(), "smtp configuration incomplete")
	}
}

func TestSMTPNotifier_Notify_MessageFormatting(t *testing.T) {
	var capturedMsg []byte
	
	mockSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		capturedMsg = msg
		return nil
	}
	
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
		Send: mockSend,
	}
	
	testMessage := "Hello, this is a test message\nwith multiple lines"
	err := notifier.Notify(context.Background(), testMessage)
	
	assert.NoError(t, err)
	
	msgStr := string(capturedMsg)
	assert.Contains(t, msgStr, "Subject: Subtitle Manager")
	assert.Contains(t, msgStr, testMessage)
	assert.Contains(t, msgStr, "\r\n\r\n") // Headers should be separated from body
}

func TestSMTPNotifier_Notify_EmptyMessage(t *testing.T) {
	var capturedMsg []byte
	
	mockSend := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		capturedMsg = msg
		return nil
	}
	
	notifier := SMTPNotifier{
		Addr: "smtp.example.com:587",
		From: "sender@example.com",
		To:   []string{"recipient@example.com"},
		Send: mockSend,
	}
	
	err := notifier.Notify(context.Background(), "")
	
	assert.NoError(t, err)
	
	msgStr := string(capturedMsg)
	assert.Contains(t, msgStr, "Subject: Subtitle Manager")
	assert.Contains(t, msgStr, "\r\n\r\n") // Empty body should still have headers
}