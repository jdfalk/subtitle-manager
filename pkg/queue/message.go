// file: pkg/queue/message.go
// version: 1.0.0
// guid: 2f76c9a4-8f24-4d5f-9aaa-97db45af4c61
package queue

import "google.golang.org/protobuf/types/known/anypb"

// QueueMessage mirrors the gcommon queue message for basic use.
type QueueMessage struct {
	Id   string
	Body *anypb.Any
}
