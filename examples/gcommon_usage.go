// file: examples/gcommon_usage.go
// version: 1.0.0
// guid: 12345678-9abc-def0-1234-567890abcdef

package main

import (
	"fmt"
	"log"

	"github.com/jdfalk/gcommon/sdks/go/gcommon/v1/common"
	"github.com/jdfalk/gcommon/sdks/go/gcommon/v1/queue"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Example 1: Using common protobuf types from gcommon
	retryConfig := &common.ConfigRetrySettings{
		MaxRetries:        3,
		BackoffMultiplier: 2.0,
		InitialInterval:   "1s",
		MaxInterval:       "30s",
	}

	fmt.Printf("Created retry config: %+v\n", retryConfig)

	// Example 2: Using queue protobuf types from gcommon
	queueMessage := &queue.QueueMessage{
		Id:       "msg-123",
		Payload:  []byte("Hello from subtitle-manager!"),
		Priority: 1,
	}

	fmt.Printf("Created queue message: %+v\n", queueMessage)

	// Example 3: Marshaling to protobuf binary format
	data, err := proto.Marshal(queueMessage)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	fmt.Printf("Marshaled message size: %d bytes\n", len(data))

	// Example 4: Unmarshaling from protobuf binary format
	var unmarshaled queue.QueueMessage
	err = proto.Unmarshal(data, &unmarshaled)
	if err != nil {
		log.Fatalf("Failed to unmarshal message: %v", err)
	}

	fmt.Printf("Unmarshaled message: %+v\n", &unmarshaled)
	fmt.Println("✅ Successfully used gcommon protobuf packages in subtitle-manager!")
}
