<!-- file: docs/tasks/00-ARCHITECTURE-OVERVIEW.md -->
<!-- version: 1.0.0 -->
<!-- guid: 00000000-1111-2222-3333-444444444444 -->

# Subtitle Manager: 3-Service Active-Active Architecture

## 🎯 Architecture Overview

This document defines the complete architecture for subtitle-manager's transition to a 3-service active-active distributed system with simplified deployment options.

## 🏗️ Service Architecture

### **3 Core Services**

```
┌─────────────────────────────────────────┐
│            Internal Network             │
├─────────────────────────────────────────┤
│  Web Service     │  Web Service          │
│  (Client API)    │  (Load Balanced)      │
│  - HTTP/API      │  - WebSocket          │
│  - File Uploads  │  - File Downloads     │
│  - Authentication│  - Request Routing    │
└─────────┬────────┴─────────┬─────────────┘
          │ gRPC             │ gRPC
┌─────────▼──────────────────▼─────────────┐
│         Engine Services                  │
│  ┌─────────────┬─────────────────────┐   │
│  │Engine S1    │ Engine S2           │   │
│  │ACTIVE       │ ACTIVE              │   │
│  │- Translation│ - Translation       │   │
│  │- Monitoring │ - Monitoring        │   │
│  │- Coord(🔴)  │ - Coord(LEADER 👑)  │   │
│  │  Standby    │   Active            │   │
│  └─────────────┴─────────────────────┘   │
│  ┌─────────────┬─────────────────────┐   │
│  │Engine S3    │ Engine S4           │   │
│  │ACTIVE       │ ACTIVE              │   │
│  │- Translation│ - Translation       │   │
│  │- Monitoring │ - Monitoring        │   │
│  │- Coord(🔴)  │ - Coord(🔴)         │   │
│  │  Standby    │   Standby           │   │
│  └─────────────┴─────────────────────┘   │
└─────────┬──────────────────┬─────────────┘
          │ gRPC             │ gRPC  
┌─────────▼──────────────────▼─────────────┐
│           File Services                  │
│  ┌─────────────┬─────────────────────┐   │
│  │File S1      │ File S2             │   │
│  │/media/movies│ /media/tv           │   │
│  │- Watching   │ - Scanning          │   │
│  │- Extraction │ - I/O Operations    │   │
│  └─────────────┴─────────────────────┘   │
└─────────────────────────────────────────┘
```

### **Service Responsibilities**

#### **1. Web Service (Pure Client Interface)**
- **Purpose**: HTTP/WebSocket API for all clients (web UI, CLI, mobile apps)
- **Responsibilities**:
  - HTTP API endpoints and WebSocket connections
  - File upload/download handling (validates, streams to/from File Services)
  - Authentication and authorization
  - Request routing to Engine Services
  - Client session management
- **What it DOESN'T do**:
  - File monitoring or watching
  - Translation processing
  - File system operations
  - Coordination logic

#### **2. Engine Service (Active-Active + Leader Election)**
- **Purpose**: Core processing with distributed translation and coordinated monitoring
- **Active-Active Responsibilities** (ALL instances):
  - ✅ Translation processing (Whisper, AI, Google Translate)
  - ✅ Job queue processing
  - ✅ Monitoring task execution
  - ✅ Health checking
- **Leader Election Responsibilities** (ONE instance):
  - 👑 Work distribution and scheduling
  - 👑 Resource coordination
  - 👑 System-wide monitoring orchestration
  - 👑 Health status aggregation
- **Scaling**: Add more instances for more translation/processing capacity

#### **3. File Service (Pure File Operations)**
- **Purpose**: All file system operations and media processing
- **Responsibilities**:
  - File system watching (inotify, polling)
  - File I/O operations (read/write subtitles, media files)
  - Media extraction (subtitle extraction from video files)
  - Storage management (cleanup, organization)
  - Directory scanning and indexing
- **Scaling**: Deploy per storage volume or geographic location

## 🔄 Service Communication Patterns

### **Translation Workload Distribution**
```
Client Request → Web Service → Engine Load Balancer
                              ├─→ Engine S1 (Whisper Instance 1)
                              ├─→ Engine S2 (Whisper Instance 2) 👑 
                              ├─→ Engine S3 (AI Translation)
                              └─→ Engine S4 (Google Translate)
```

### **File Operation Flow**
```
Upload Request → Web Service → File Service (validate & store)
                            → Engine Service (process if needed)

File Change → File Service (detect) → Engine Service Leader (coordinate)
                                  → Engine Services (process)
```

### **Monitoring Flow**
```
Engine Leader → File Services (scan requests)
             → Engine Services (distribute work)
             → Web Services (status updates)
```

## 📋 Task Breakdown Structure

### **Phase 1: Core Architecture (01-core-architecture/)**
- **TASK-01-001**: Service Interface Definitions
- **TASK-01-002**: gRPC Protocol Design
- **TASK-01-003**: Service Discovery Framework
- **TASK-01-004**: Configuration Management
- **TASK-01-005**: Error Handling Standards

### **Phase 2: Web Service (02-web-service/)**
- **TASK-02-001**: HTTP API Refactoring
- **TASK-02-002**: File Upload/Download Handlers
- **TASK-02-003**: WebSocket Implementation
- **TASK-02-004**: Authentication Integration
- **TASK-02-005**: Request Routing Logic

### **Phase 3: Engine Service (03-engine-service/)**
- **TASK-03-001**: Translation Engine Core
- **TASK-03-002**: Job Queue System
- **TASK-03-003**: Monitoring Integration
- **TASK-03-004**: Leader Election Framework
- **TASK-03-005**: Load Balancing Logic

### **Phase 4: File Service (04-file-service/)**
- **TASK-04-001**: File System Watching
- **TASK-04-002**: Media Processing Pipeline
- **TASK-04-003**: Storage Management
- **TASK-04-004**: I/O Optimization
- **TASK-04-005**: Cleanup and Maintenance

### **Phase 5: Service Communication (05-service-communication/)**
- **TASK-05-001**: gRPC Service Implementation
- **TASK-05-002**: Message Protocols
- **TASK-05-003**: Error Handling and Retries
- **TASK-05-004**: Health Checking
- **TASK-05-005**: Connection Management

### **Phase 6: Leader Election (06-leader-election/)**
- **TASK-06-001**: Election Algorithm Implementation
- **TASK-06-002**: Failover Mechanisms
- **TASK-06-003**: Split-Brain Prevention
- **TASK-06-004**: State Synchronization
- **TASK-06-005**: Recovery Procedures

### **Phase 7: Testing (07-testing/)**
- **TASK-07-001**: Unit Testing Framework
- **TASK-07-002**: Integration Testing
- **TASK-07-003**: End-to-End Testing
- **TASK-07-004**: Performance Testing
- **TASK-07-005**: Chaos Engineering

### **Phase 8: Deployment (08-deployment/)**
- **TASK-08-001**: All-in-One Mode
- **TASK-08-002**: Distributed Deployment
- **TASK-08-003**: Container Configuration
- **TASK-08-004**: Monitoring and Logging
- **TASK-08-005**: Production Readiness

## 🎯 Key Benefits

### **Scalability**
- **Horizontal Scaling**: Add Engine instances for more translation capacity
- **Specialized Workers**: Different instances can run different translation engines
- **Geographic Distribution**: File services can be deployed per region/datacenter

### **Reliability**
- **Active-Active Translation**: No single point of failure for processing
- **Leader Election**: Coordination continues if leader fails
- **Service Isolation**: Web/File/Engine failures don't cascade

### **Performance**
- **Load Distribution**: Multiple Whisper/AI instances running simultaneously
- **Efficient I/O**: File services optimize disk operations
- **Smart Routing**: Work goes to most appropriate service instance

### **Flexibility**
- **All-in-One Mode**: Single binary for simple deployments
- **Distributed Mode**: Full microservices for enterprise scale
- **Mixed Deployment**: Scale individual components as needed

## 🛠️ Development Standards

### **Each Task Must Include**:
1. **Detailed Implementation Steps** (500-1000+ lines)
2. **Complete Code Examples** with file-by-file changes
3. **Testing Procedures** for each component
4. **Error Handling** strategies
5. **Performance Considerations**
6. **Standard Instructions** integration
7. **Beginner-Friendly Explanations**

### **Code Quality Requirements**:
- **File Headers**: All files must include standard headers
- **Version Management**: Semantic versioning for all components
- **Documentation**: Inline code documentation required
- **Testing**: Minimum 80% test coverage per service
- **Error Handling**: Comprehensive error handling and logging

This architecture provides a clean, scalable foundation that can grow from single-server deployments to enterprise-scale distributed systems while maintaining simplicity and reliability.
