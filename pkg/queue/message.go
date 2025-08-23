// file: pkg/queue/message.go
// version: 1.3.0
// guid: 2f76c9a4-8f24-4d5f-9aaa-97db45af4c61
package queue

import (
	"time"
)

// QueueMessage represents a message in the queue system
// This is a temporary native Go implementation to replace gcommon protobuf dependency
// TODO: Replace with gcommon.v1.queue.QueueMessage when protobuf integration is complete
type QueueMessage struct {
	ID          string            `json:"id"`
	Topic       string            `json:"topic"`
	Payload     []byte            `json:"payload"`
	Headers     map[string]string `json:"headers"`
	CreatedAt   time.Time         `json:"created_at"`
	ProcessedAt *time.Time        `json:"processed_at,omitempty"`
	RetryCount  int32             `json:"retry_count"`
	Priority    int32             `json:"priority"`
}
