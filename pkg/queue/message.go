// file: pkg/queue/message.go
// version: 1.1.0
// guid: 2f76c9a4-8f24-4d5f-9aaa-97db45af4c61
package queue

import queuepb "github.com/jdfalk/gcommon/pkg/queue/proto"

// QueueMessage is an alias of the gcommon queue message type.
// This allows the internal queue package to use the centralized
// protobuf definition without rewriting existing code.
type QueueMessage = queuepb.QueueMessage
