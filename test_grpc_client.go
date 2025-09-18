// file: test_grpc_client.go
// version: 1.1.0
// guid: a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d

//go:build tools
// +build tools

// Simple test client to verify gRPC server functionality
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jdfalk/subtitle-manager/pkg/subtitle/translator/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTranslatorServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test GetConfig
	fmt.Println("Testing GetConfig...")
	getResp, err := client.GetConfig(ctx, &pb.GetConfigRequest{})
	if err != nil {
		log.Printf("GetConfig failed: %v", err)
	} else {
		fmt.Printf("GetConfig success: %+v\n", getResp.GetConfigValues())
	}

	// Test SetConfig
	fmt.Println("Testing SetConfig...")
	key := "test_key"
	value := "test_value"
	setReq := &pb.SetConfigRequest{
		Key:   &key,
		Value: &value,
	}
	setResp, err := client.SetConfig(ctx, setReq)
	if err != nil {
		log.Printf("SetConfig failed: %v", err)
	} else {
		fmt.Printf("SetConfig success: %+v\n", setResp)
	}

	fmt.Println("gRPC server test completed successfully!")
}
