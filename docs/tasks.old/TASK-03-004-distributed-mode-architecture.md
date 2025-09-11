<!-- file: docs/tasks/TASK-03-004-secure-distributed-architecture-CONSOLIDATED.md -->
<!-- version: 1.0.0 -->
<!-- guid: 4d5e6f7a-8b9c-0d1e-2f3a-4b5c6d7e8f9a -->

# TASK-03-004: Secure Distributed Architecture + Function Redistribution (CONSOLIDATED)

## Priority: MEDIUM

## Status: PENDING

## Estimated Time: 12-16 hours

## Overview

**CONSOLIDATED TASK**: Design and implement secure distributed architecture with
proper network segmentation AND redistribute functions for optimal service
boundaries. This combines architecture design with function movement for
efficiency.

## Architecture Vision: DMZ + Backend Security Model

### Network Security Design

```
┌─────────────────────────────────────┐
│              DMZ Zone               │
├─────────────────────────────────────┤
│  Web Server 1    │  Web Server 2   │
│  (Load Balanced) │  (Load Balanced) │
│  HTTP/WebSocket  │  HTTP/WebSocket  │
│  to Clients      │  to Clients      │
└─────────┬────────┴─────────┬────────┘
          │ gRPC only        │ gRPC only
┌─────────▼──────────────────▼────────┐
│            Firewall                 │
│     (Only gRPC 50051-50059)         │
└─────────┬──────────────────┬────────┘
          │                  │
┌─────────▼──────────────────▼────────┐
│          Backend Network            │
├─────────────────────────────────────┤
│  Service Discovery  │  Translator   │
│  (gRPC Registry)    │  Service      │
│                     │  (gRPC only)  │
├─────────────────────┼───────────────┤
│  File Server 1      │  File Server 2│
│  (gRPC only)        │  (gRPC only)  │
│  /media/movies      │  /media/tv    │
└─────────────────────────────────────┘
```

**Security Benefits:**

- **DMZ Isolation**: Web servers in DMZ handle all user interaction
- **No File Server Web Access**: File servers only expose gRPC internally
- **Firewall Protection**: Only specific gRPC ports allowed through firewall
- **Service Isolation**: Backend services can't be reached directly from
  internet
- **Upload Security**: File uploads go through DMZ web server, then gRPC to file
  server

## Part A: Secure Service Architecture

### 1. DMZ Web Service

```bash
subtitle-manager web --service-discovery backend-discovery:50050
```

**Responsibilities:**

- **Client Interface**: HTTP API, WebUI, WebSocket
- **Security**: Authentication, authorization, rate limiting
- **Upload Handling**: Receive file uploads, validate, forward via gRPC
- **Service Coordination**: Route requests to appropriate backend services
- **No File Access**: Never touches file system directly

**Security Features:**

- Input validation and sanitization
- File upload scanning before forwarding
- Rate limiting and DDoS protection
- Authentication token validation
- Request routing and load balancing

### 2. Backend File Services (gRPC Only)

```bash
subtitle-manager file-service --discovery backend-discovery:50050 --paths /media/movies
```

**Responsibilities:**

- **File Operations**: All disk I/O, scanning, monitoring
- **Media Processing**: Subtitle extraction, format conversion
- **Storage Management**: File organization, cleanup
- **Monitoring**: File system watching and periodic scanning

**Security Features:**

- **No Web Interface**: Only gRPC endpoints
- **Path Restrictions**: Can only access configured paths
- **Internal Network Only**: Not accessible from DMZ
- **Authentication**: gRPC TLS with service certificates

### 3. Service Discovery (Backend Only)

```bash
subtitle-manager discovery --port 50050
```

**Responsibilities:**

- **Service Registration**: Backend services register with discovery
- **Health Monitoring**: Track service availability
- **Load Balancing**: Distribute requests across instances
- **Configuration Distribution**: Push config updates to services

**Service Discovery Options:**

1. **Built-in gRPC LoadBalancer**: Use gRPC's native service discovery
2. **Simple Coordinator**: Custom service registry (recommended for simplicity)
3. **etcd**: If we need more advanced features
4. **Consul**: HashiCorp option with more features

**Recommendation: Start with simple coordinator, expand if needed**

## Part B: Function Redistribution

### Functions Moving to File Service

**From Web Service:**

- File upload storage (keep HTTP interface in web, actual storage in file
  service)
- Subtitle file serving (web proxies requests)
- Directory scanning operations
- File validation and format checking
- Media metadata extraction

**From Translator Service:**

- `extract` - Subtitle extraction from media files
- File reading operations for translation input
- File writing operations for translation output
- Subtitle format conversion with file I/O

**Rationale:** All disk operations centralized for security and deployment
flexibility

### Functions Staying in Translator Service

- Text translation logic (Google, OpenAI APIs)
- In-memory subtitle format conversion
- Language detection
- Translation caching (memory/Redis)
- Batch processing coordination

**Rationale:** Pure processing functions with no file system dependencies

### New Service Communication Patterns

#### File Upload Workflow

```
1. Client uploads to DMZ Web Server (HTTP)
2. Web Server validates file (virus scan, type check)
3. Web Server calls FileService.SaveUpload(data) via gRPC
4. File Service stores file and returns metadata
5. Web Server returns success to client
```

#### Translation Workflow

```
1. Client requests translation via DMZ Web Server
2. Web Server calls FileService.ReadSubtitle(path) via gRPC
3. Web Server calls TranslatorService.Translate(data) via gRPC
4. Web Server calls FileService.WriteSubtitle(path, translated) via gRPC
5. Web Server returns result to client
```

#### Media Scanning Workflow

```
1. File Service monitors file system changes
2. File Service calls TranslatorService.ProcessNewMedia(metadata) via gRPC
3. Translator requests file data via FileService.GetMediaData(path) via gRPC
4. Results stored via FileService.SaveResults() via gRPC
```

## Implementation Steps

### Phase 1: Service Interface Definition (3-4 hours)

1. **Define secure gRPC interfaces**
   - FileService with comprehensive file operations
   - TranslatorService for pure processing
   - DiscoveryService for service registration
   - WebService for DMZ coordination

2. **Implement service discovery**
   - Simple coordinator with health checking
   - Service registration with metadata
   - Load balancing algorithms

### Phase 2: Security Implementation (4-5 hours)

1. **DMZ Web Service Security**
   - Remove all file system access
   - Implement gRPC client connections to backend
   - Add request validation and rate limiting
   - Implement upload security scanning

2. **Backend Service Security**
   - Remove web interfaces from file services
   - Implement gRPC-only communication
   - Add service authentication and TLS
   - Implement path restrictions

### Phase 3: Function Migration (4-5 hours)

1. **Move file operations to file service**
   - Extract subtitle extraction from other services
   - Centralize all file I/O operations
   - Implement streaming for large files

2. **Optimize translator service**
   - Remove file dependencies
   - Focus on pure processing
   - Implement efficient caching

### Phase 4: Testing and Optimization (2-3 hours)

1. **Security testing**
   - Verify DMZ isolation
   - Test firewall configuration
   - Validate authentication

2. **Performance testing**
   - Benchmark gRPC communication
   - Test with high concurrency
   - Optimize bottlenecks

## Service Interface Examples

### File Service API

```go
service FileService {
    // File Operations
    rpc ReadSubtitle(ReadRequest) returns (SubtitleData);
    rpc WriteSubtitle(WriteRequest) returns (WriteResult);
    rpc SaveUpload(UploadRequest) returns (FileMetadata);

    // Media Operations
    rpc ExtractSubtitles(ExtractRequest) returns (SubtitleData);
    rpc GetMediaMetadata(MetadataRequest) returns (MediaMetadata);

    // Directory Operations
    rpc ScanDirectory(ScanRequest) returns (stream FileInfo);
    rpc WatchDirectory(WatchRequest) returns (stream FileEvent);
}
```

### Translator Service API

```go
service TranslatorService {
    // Pure Processing
    rpc TranslateText(TranslateRequest) returns (TranslateResult);
    rpc TranslateSubtitle(SubtitleRequest) returns (SubtitleData);
    rpc DetectLanguage(DetectRequest) returns (LanguageResult);
    rpc BatchTranslate(BatchRequest) returns (stream TranslateResult);
}
```

## Configuration Examples

### DMZ Deployment

```bash
# DMZ Web Servers (Multiple instances)
subtitle-manager web \
  --discovery backend-discovery:50050 \
  --port 8080 \
  --no-file-access \
  --upload-max-size 100MB
```

### Backend Deployment

```bash
# Service Discovery
subtitle-manager discovery --port 50050 --bind-internal

# File Services (Multiple instances for different paths)
subtitle-manager file-service \
  --discovery backend-discovery:50050 \
  --paths /media/movies \
  --no-web-interface

subtitle-manager file-service \
  --discovery backend-discovery:50050 \
  --paths /media/tv \
  --no-web-interface

# Translator Services (Stateless, multiple instances)
subtitle-manager translator \
  --discovery backend-discovery:50050 \
  --workers 4
```

## Success Criteria

**Security:**

- ✅ DMZ web servers have no file system access
- ✅ File servers only accessible via gRPC from internal network
- ✅ All user interaction goes through DMZ web servers
- ✅ Backend services protected by firewall
- ✅ Service authentication and TLS implemented

**Architecture:**

- ✅ Services can be deployed independently
- ✅ Automatic service discovery and registration
- ✅ Efficient load balancing and failover
- ✅ File services handle all disk operations
- ✅ Translator services are stateless and scalable
- ✅ Web service coordinates client requests securely

**Function Distribution:**

- ✅ All file I/O centralized in file service
- ✅ Pure processing functions in translator service
- ✅ No cross-service file system dependencies
- ✅ Efficient gRPC communication patterns

## Dependencies

- TASK-03-001 (Unified monitoring + renaming)
- TASK-03-002 (All-in-one mode)

## Security Benefits

This architecture provides:

1. **Network Segmentation**: DMZ isolation with backend protection
2. **Minimal Attack Surface**: File servers not web-accessible
3. **Defense in Depth**: Multiple security layers (DMZ, firewall, service auth)
4. **Secure File Handling**: All uploads validated before backend processing
5. **Service Isolation**: Backend services can't be reached directly
6. **Scalable Security**: Easy to add more DMZ web servers for load
