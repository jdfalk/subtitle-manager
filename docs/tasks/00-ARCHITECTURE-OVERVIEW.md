<!-- file: docs/tasks/00-ARCHITECTURE-OVERVIEW.md -->
<!-- version: 1.0.0 -->
<!-- guid: arch0000-1111-2222-3333-444444444444 -->

# 00-ARCHITECTURE-OVERVIEW: 3-Service Active-Active Architecture

## Executive Summary

This document defines the complete architectural redesign from the original 6-service DMZ model to a simplified 3-service active-active architecture. The new design eliminates unnecessary complexity while providing better scalability, security, and maintainability.

## Service Architecture

### Core Services

1. **Web Service** - Client interface and user management
2. **Engine Service** - Translation, monitoring, and coordination
3. **File Service** - Pure file operations and media processing

### Service Communication

- **Internal Network Only**: All services communicate via gRPC on internal network
- **No DMZ Complexity**: Single Web service handles external exposure
- **Service Discovery**: Automatic service registration and discovery
- **Load Balancing**: Built-in support for horizontal scaling

## Active-Active Translation Design

### Translation Workers
- **Multiple Active Instances**: All translation workers are active simultaneously
- **Horizontal Scaling**: Add more translation workers for increased capacity
- **Work Distribution**: Tasks distributed across all available workers
- **No Single Point of Failure**: System continues if workers fail

### Leader Election for Coordination Only
- **Coordination Leader**: Single leader coordinates task assignment
- **Not for Translation**: Leader does NOT perform translation work
- **Automatic Failover**: New leader elected if current leader fails
- **Minimal Coordination**: Leader only assigns tasks and monitors progress

## Service Responsibilities

### Web Service
- **User Authentication**: Login, session management, JWT tokens
- **API Gateway**: REST/GraphQL endpoints for client applications
- **File Upload/Download**: Handle client file transfers
- **Request Routing**: Route requests to appropriate backend services
- **Rate Limiting**: Protect backend services from abuse
- **WebUI Serving**: Serve static web assets and SPA

### Engine Service
- **Translation Processing**: Perform subtitle translation work
- **Worker Coordination**: Distribute translation tasks among workers
- **Progress Monitoring**: Track translation job progress
- **Library Monitoring**: Monitor file system for new media
- **Task Scheduling**: Schedule and manage background tasks
- **Leader Election**: Coordinate work distribution

### File Service
- **File Operations**: Read, write, move, delete files
- **Media Processing**: Extract subtitles, convert formats
- **File Watching**: Monitor directories for changes
- **Storage Management**: Cleanup, validation, backup operations
- **Metadata Extraction**: Extract media file information
- **Format Conversion**: Convert between subtitle formats

## Task Structure

### Implementation Phases

1. **01-core-architecture**: Service interfaces and protobuf definitions
2. **02-web-service**: Web service implementation
3. **03-engine-service**: Engine service with translation and coordination
4. **04-file-service**: File service implementation
5. **05-service-coordination**: Service discovery and communication
6. **06-monitoring-observability**: Logging, metrics, and monitoring
7. **07-integration-testing**: End-to-end testing and validation
8. **08-deployment**: Docker, Kubernetes, and production deployment

### Task Characteristics
- **Extremely Detailed**: Each task 500-1000+ lines with complete specifications
- **Self-Contained**: Each task can be implemented independently
- **Parallel Development**: Multiple AI agents can work simultaneously
- **Clear Dependencies**: Explicit task dependencies and prerequisites

## Technology Stack

### Core Technologies
- **Go 1.24+**: Primary development language
- **gRPC**: Service-to-service communication
- **Protocol Buffers Edition 2023**: Message definitions
- **SQLite/PostgreSQL**: Data persistence
- **Docker**: Containerization
- **Kubernetes**: Orchestration (optional)

### Supporting Technologies
- **Prometheus**: Metrics collection
- **Grafana**: Monitoring dashboards
- **Jaeger**: Distributed tracing
- **Zap**: Structured logging
- **Consul/etcd**: Service discovery (optional)

## Deployment Models

### Development
- **Docker Compose**: Single machine deployment
- **Local File Storage**: Files stored on local filesystem
- **SQLite Database**: Embedded database for simplicity

### Production
- **Kubernetes**: Multi-node deployment
- **Persistent Volumes**: Shared storage for file service
- **PostgreSQL**: Dedicated database instance
- **Load Balancers**: External load balancing for web service

## Security Model

### Network Security
- **Internal Network**: Services communicate on private network
- **Single Entry Point**: Only web service exposed externally
- **TLS Everywhere**: All gRPC communication encrypted
- **API Authentication**: JWT tokens for all API access

### File Security
- **Path Validation**: Prevent directory traversal attacks
- **Permission Checks**: Validate file access permissions
- **Content Scanning**: Scan uploaded files for malware
- **Audit Logging**: Log all file operations

## Scalability Design

### Horizontal Scaling
- **Stateless Services**: All services designed to be stateless
- **Shared Storage**: File service uses shared storage backend
- **Load Distribution**: Work distributed across multiple instances
- **Auto-scaling**: Kubernetes HPA for automatic scaling

### Performance Optimization
- **Caching**: Redis for frequently accessed data
- **Connection Pooling**: Efficient database connections
- **Batch Processing**: Batch similar operations
- **Streaming**: Stream large file transfers

## Migration Strategy

### From Current System
- **Gradual Migration**: Migrate one service at a time
- **Feature Parity**: Maintain all existing functionality
- **Data Migration**: Preserve all existing data
- **Zero Downtime**: Rolling deployment strategy

### Rollback Plan
- **Blue-Green Deployment**: Keep old version running during migration
- **Database Compatibility**: Maintain backward compatibility
- **Configuration Rollback**: Quick configuration reversion
- **Monitoring**: Extensive monitoring during migration

## Success Criteria

### Functional Requirements
- **All Features**: 100% feature parity with current system
- **Performance**: Equal or better performance than current system
- **Reliability**: 99.9% uptime target
- **Scalability**: Support 10x current load

### Non-Functional Requirements
- **Maintainability**: Simplified codebase and architecture
- **Testability**: Comprehensive test coverage
- **Observability**: Complete monitoring and logging
- **Security**: Enhanced security posture

## Next Steps

1. **Review and Approve**: Stakeholder review of architecture
2. **Task Implementation**: Begin parallel task development
3. **Prototype Deployment**: Deploy minimal viable system
4. **Integration Testing**: Comprehensive testing
5. **Production Migration**: Gradual migration to new system

This architecture provides a solid foundation for scalable, maintainable, and secure subtitle management services while eliminating the complexity of the previous design.
