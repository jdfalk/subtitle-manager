# file: docs/PROTOBUF_REGEN.md
# version: 1.1.0
# guid: 3e4f5a6b-7c8d-9e0f-1a2b-3c4d5e6f7a8b

# Regenerating Protobuf Files

This guide explains how to regenerate the Go code used for gRPC.

1. Ensure `protoc`, `protoc-gen-go` and `protoc-gen-go-grpc` are installed.
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
2. Run the Makefile target which invokes `protoc`:
   ```bash
   make proto-gen
   ```
   This command reads `proto/translator.proto` and writes the generated bindings to `pkg/translatorpb`.
3. Commit the updated files alongside your changes.

The protobuf definitions include a `SetConfig` RPC used to update configuration
values over gRPC. After regeneration ensure any clients or servers implementing
this RPC are recompiled.

# Regenerating Mock Files

Mock files are generated using Mockery v2 with a packages-based configuration:

1. Run the Makefile target:
   ```bash
   make mock-gen
   ```
   Or generate both protobuf and mock files:
   ```bash
   make generate
   ```

2. The configuration in `.mockery.yaml` defines which interfaces to mock and where to place the files.

3. Mock files are generated in their respective package `mocks/` directories.

4. Commit the updated mock files alongside your changes.
