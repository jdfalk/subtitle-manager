// file: examples/gcommon_usage.go
// version: 1.0.0
// guid: 12345678-9abc-def0-1234-567890abcdef

// file: examples/gcommon_usage.go
// version: 1.0.0
// guid: 12345678-1234-1234-1234-123456789abc

package main

import (
	"fmt"
	"log"

	"github.com/jdfalk/gcommon/sdks/go/v1/common"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Example 1: Create a retry settings configuration using the builder
	enabled := true
	maxRetries := int32(3)
	delaySeconds := int32(1)

	retrySettings := (&common.ConfigRetrySettings_builder{
		Enabled:      &enabled,
		MaxRetries:   &maxRetries,
		DelaySeconds: &delaySeconds,
	}).Build()

	fmt.Printf("Created retry settings: enabled=%v, max_retries=%d, delay=%d\n",
		retrySettings.GetEnabled(), retrySettings.GetMaxRetries(), retrySettings.GetDelaySeconds())

	// Example 2: Using setter methods
	retrySettings2 := &common.ConfigRetrySettings{}
	retrySettings2.SetEnabled(true)
	retrySettings2.SetMaxRetries(5)
	retrySettings2.SetDelaySeconds(2)

	fmt.Printf("Created retry settings 2: enabled=%v, max_retries=%d, delay=%d\n",
		retrySettings2.GetEnabled(), retrySettings2.GetMaxRetries(), retrySettings2.GetDelaySeconds())

	// Example 3: Serialize to bytes
	data, err := proto.Marshal(retrySettings)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	fmt.Printf("Serialized size: %d bytes\n", len(data))

	// Example 4: Deserialize from bytes
	retrySettings3 := &common.ConfigRetrySettings{}
	err = proto.Unmarshal(data, retrySettings3)
	if err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}
	fmt.Printf("Deserialized: enabled=%v, max_retries=%d, delay=%d\n",
		retrySettings3.GetEnabled(), retrySettings3.GetMaxRetries(), retrySettings3.GetDelaySeconds())
}
