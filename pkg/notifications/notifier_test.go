// file: pkg/notifications/notifier_test.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174010

package notifications

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc_Notify(t *testing.T) {
	tests := []struct {
		name        string
		fn          Func
		msg         string
		expectError bool
	}{
		{
			name: "successful notification",
			fn: func(ctx context.Context, msg string) error {
				return nil
			},
			msg:         "test message",
			expectError: false,
		},
		{
			name: "failed notification",
			fn: func(ctx context.Context, msg string) error {
				return errors.New("notification failed")
			},
			msg:         "test message",
			expectError: true,
		},
		{
			name: "empty message",
			fn: func(ctx context.Context, msg string) error {
				assert.Empty(t, msg)
				return nil
			},
			msg:         "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn.Notify(context.Background(), tt.msg)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFunc_MessagePassing(t *testing.T) {
	receivedMessage := ""
	var receivedCtx context.Context

	fn := Func(func(ctx context.Context, msg string) error {
		receivedCtx = ctx
		receivedMessage = msg
		return nil
	})

	ctx := context.WithValue(context.Background(), "test", "value")
	testMessage := "hello world"

	err := fn.Notify(ctx, testMessage)

	assert.NoError(t, err)
	assert.Equal(t, testMessage, receivedMessage)
	assert.Equal(t, ctx, receivedCtx)
}

func TestNop_Notify(t *testing.T) {
	nop := Nop{}

	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "with message",
			msg:  "test message",
		},
		{
			name: "empty message",
			msg:  "",
		},
		{
			name: "long message",
			msg:  "this is a very long message that should still be handled properly by the nop notifier",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := nop.Notify(context.Background(), tt.msg)
			assert.NoError(t, err)
		})
	}
}

func TestNop_ContextHandling(t *testing.T) {
	nop := Nop{}

	// Test with nil context (should not panic)
	err := nop.Notify(nil, "test message")
	assert.NoError(t, err)

	// Test with canceled context (should still succeed)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err = nop.Notify(ctx, "test message")
	assert.NoError(t, err)
}

func TestNotifierInterface(t *testing.T) {
	// Verify that all types implement the Notifier interface
	var _ Notifier = Func(nil)
	var _ Notifier = Nop{}
	var _ Notifier = DiscordNotifier{}
	var _ Notifier = SMTPNotifier{}
	var _ Notifier = TelegramNotifier{}
}
