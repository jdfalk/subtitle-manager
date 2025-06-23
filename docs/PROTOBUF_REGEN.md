# file: docs/PROTOBUF_REGEN.md

# Regenerating Protobuf Files

This guide explains how to regenerate the Go code used for gRPC.

1. Ensure `protoc`, `protoc-gen-go` and `protoc-gen-go-grpc` are installed.
   \```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   \```
2. Run the Makefile target which invokes `protoc`:
   \```bash
make proto-gen
\```
This command reads `proto/translator.proto`and writes the generated
bindings to`pkg/translatorpb`.
3. Commit the updated files alongside your changes.
