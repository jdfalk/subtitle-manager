// file: pkg/notifications/notifier.go
package notifications

import "context"

// Notifier sends messages to a notification service.
type Notifier interface {
	// Notify sends msg to the underlying service.
	// It returns an error if delivery fails.
	Notify(ctx context.Context, msg string) error
}

// Func adapts a function to the Notifier interface.
type Func func(ctx context.Context, msg string) error

// Notify calls f(ctx, msg).
func (f Func) Notify(ctx context.Context, msg string) error { return f(ctx, msg) }

// Nop is a Notifier that performs no action.
type Nop struct{}

// Notify implements Notifier for Nop.
func (Nop) Notify(context.Context, string) error { return nil }
