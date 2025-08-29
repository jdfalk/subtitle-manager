- [Constants](<#constants>)
- [Variables](<#variables>)
- [func RegisterCacheAdminServiceServer\(s grpc.ServiceRegistrar, srv CacheAdminServiceServer\)](<#RegisterCacheAdminServiceServer>)
- [func RegisterCacheServiceServer\(s grpc.ServiceRegistrar, srv CacheServiceServer\)](<#RegisterCacheServiceServer>)
- [func RegisterDatabaseAdminServiceServer\(s grpc.ServiceRegistrar, srv DatabaseAdminServiceServer\)](<#RegisterDatabaseAdminServiceServer>)
- [func RegisterDatabaseServiceServer\(s grpc.ServiceRegistrar, srv DatabaseServiceServer\)](<#RegisterDatabaseServiceServer>)
- [func RegisterMigrationServiceServer\(s grpc.ServiceRegistrar, srv MigrationServiceServer\)](<#RegisterMigrationServiceServer>)
- [func RegisterTransactionServiceServer\(s grpc.ServiceRegistrar, srv TransactionServiceServer\)](<#RegisterTransactionServiceServer>)
- [type AppendRequest](<#AppendRequest>)
  - [func \(x \*AppendRequest\) ClearKey\(\)](<#AppendRequest.ClearKey>)
  - [func \(x \*AppendRequest\) ClearMetadata\(\)](<#AppendRequest.ClearMetadata>)
  - [func \(x \*AppendRequest\) ClearNamespace\(\)](<#AppendRequest.ClearNamespace>)
  - [func \(x \*AppendRequest\) ClearValue\(\)](<#AppendRequest.ClearValue>)
  - [func \(x \*AppendRequest\) GetKey\(\) string](<#AppendRequest.GetKey>)
  - [func \(x \*AppendRequest\) GetMetadata\(\) \*common.RequestMetadata](<#AppendRequest.GetMetadata>)
  - [func \(x \*AppendRequest\) GetNamespace\(\) string](<#AppendRequest.GetNamespace>)
  - [func \(x \*AppendRequest\) GetValue\(\) \*anypb.Any](<#AppendRequest.GetValue>)
  - [func \(x \*AppendRequest\) HasKey\(\) bool](<#AppendRequest.HasKey>)
  - [func \(x \*AppendRequest\) HasMetadata\(\) bool](<#AppendRequest.HasMetadata>)
  - [func \(x \*AppendRequest\) HasNamespace\(\) bool](<#AppendRequest.HasNamespace>)
  - [func \(x \*AppendRequest\) HasValue\(\) bool](<#AppendRequest.HasValue>)
  - [func \(\*AppendRequest\) ProtoMessage\(\)](<#AppendRequest.ProtoMessage>)
  - [func \(x \*AppendRequest\) ProtoReflect\(\) protoreflect.Message](<#AppendRequest.ProtoReflect>)
  - [func \(x \*AppendRequest\) Reset\(\)](<#AppendRequest.Reset>)
  - [func \(x \*AppendRequest\) SetKey\(v string\)](<#AppendRequest.SetKey>)
  - [func \(x \*AppendRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#AppendRequest.SetMetadata>)
  - [func \(x \*AppendRequest\) SetNamespace\(v string\)](<#AppendRequest.SetNamespace>)
  - [func \(x \*AppendRequest\) SetValue\(v \*anypb.Any\)](<#AppendRequest.SetValue>)
  - [func \(x \*AppendRequest\) String\(\) string](<#AppendRequest.String>)
- [type AppendRequest\_builder](<#AppendRequest_builder>)
  - [func \(b0 AppendRequest\_builder\) Build\(\) \*AppendRequest](<#AppendRequest_builder.Build>)
- [type BackupRequest](<#BackupRequest>)
  - [func \(x \*BackupRequest\) ClearDestination\(\)](<#BackupRequest.ClearDestination>)
  - [func \(x \*BackupRequest\) ClearMetadata\(\)](<#BackupRequest.ClearMetadata>)
  - [func \(x \*BackupRequest\) ClearNamespace\(\)](<#BackupRequest.ClearNamespace>)
  - [func \(x \*BackupRequest\) GetDestination\(\) string](<#BackupRequest.GetDestination>)
  - [func \(x \*BackupRequest\) GetMetadata\(\) \*common.RequestMetadata](<#BackupRequest.GetMetadata>)
  - [func \(x \*BackupRequest\) GetNamespace\(\) string](<#BackupRequest.GetNamespace>)
  - [func \(x \*BackupRequest\) HasDestination\(\) bool](<#BackupRequest.HasDestination>)
  - [func \(x \*BackupRequest\) HasMetadata\(\) bool](<#BackupRequest.HasMetadata>)
  - [func \(x \*BackupRequest\) HasNamespace\(\) bool](<#BackupRequest.HasNamespace>)
  - [func \(\*BackupRequest\) ProtoMessage\(\)](<#BackupRequest.ProtoMessage>)
  - [func \(x \*BackupRequest\) ProtoReflect\(\) protoreflect.Message](<#BackupRequest.ProtoReflect>)
  - [func \(x \*BackupRequest\) Reset\(\)](<#BackupRequest.Reset>)
  - [func \(x \*BackupRequest\) SetDestination\(v string\)](<#BackupRequest.SetDestination>)
  - [func \(x \*BackupRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#BackupRequest.SetMetadata>)
  - [func \(x \*BackupRequest\) SetNamespace\(v string\)](<#BackupRequest.SetNamespace>)
  - [func \(x \*BackupRequest\) String\(\) string](<#BackupRequest.String>)
- [type BackupRequest\_builder](<#BackupRequest_builder>)
  - [func \(b0 BackupRequest\_builder\) Build\(\) \*BackupRequest](<#BackupRequest_builder.Build>)
- [type BatchExecuteOptions](<#BatchExecuteOptions>)
  - [func \(x \*BatchExecuteOptions\) ClearFailFast\(\)](<#BatchExecuteOptions.ClearFailFast>)
  - [func \(x \*BatchExecuteOptions\) ClearMaxParallel\(\)](<#BatchExecuteOptions.ClearMaxParallel>)
  - [func \(x \*BatchExecuteOptions\) ClearTimeout\(\)](<#BatchExecuteOptions.ClearTimeout>)
  - [func \(x \*BatchExecuteOptions\) GetFailFast\(\) bool](<#BatchExecuteOptions.GetFailFast>)
  - [func \(x \*BatchExecuteOptions\) GetMaxParallel\(\) int32](<#BatchExecuteOptions.GetMaxParallel>)
  - [func \(x \*BatchExecuteOptions\) GetTimeout\(\) \*durationpb.Duration](<#BatchExecuteOptions.GetTimeout>)
  - [func \(x \*BatchExecuteOptions\) HasFailFast\(\) bool](<#BatchExecuteOptions.HasFailFast>)
  - [func \(x \*BatchExecuteOptions\) HasMaxParallel\(\) bool](<#BatchExecuteOptions.HasMaxParallel>)
  - [func \(x \*BatchExecuteOptions\) HasTimeout\(\) bool](<#BatchExecuteOptions.HasTimeout>)
  - [func \(\*BatchExecuteOptions\) ProtoMessage\(\)](<#BatchExecuteOptions.ProtoMessage>)
  - [func \(x \*BatchExecuteOptions\) ProtoReflect\(\) protoreflect.Message](<#BatchExecuteOptions.ProtoReflect>)
  - [func \(x \*BatchExecuteOptions\) Reset\(\)](<#BatchExecuteOptions.Reset>)
  - [func \(x \*BatchExecuteOptions\) SetFailFast\(v bool\)](<#BatchExecuteOptions.SetFailFast>)
  - [func \(x \*BatchExecuteOptions\) SetMaxParallel\(v int32\)](<#BatchExecuteOptions.SetMaxParallel>)
  - [func \(x \*BatchExecuteOptions\) SetTimeout\(v \*durationpb.Duration\)](<#BatchExecuteOptions.SetTimeout>)
  - [func \(x \*BatchExecuteOptions\) String\(\) string](<#BatchExecuteOptions.String>)
- [type BatchExecuteOptions\_builder](<#BatchExecuteOptions_builder>)
  - [func \(b0 BatchExecuteOptions\_builder\) Build\(\) \*BatchExecuteOptions](<#BatchExecuteOptions_builder.Build>)
- [type BatchOperationResult](<#BatchOperationResult>)
  - [func \(x \*BatchOperationResult\) ClearAffectedRows\(\)](<#BatchOperationResult.ClearAffectedRows>)
  - [func \(x \*BatchOperationResult\) ClearError\(\)](<#BatchOperationResult.ClearError>)
  - [func \(x \*BatchOperationResult\) ClearSuccess\(\)](<#BatchOperationResult.ClearSuccess>)
  - [func \(x \*BatchOperationResult\) GetAffectedRows\(\) int64](<#BatchOperationResult.GetAffectedRows>)
  - [func \(x \*BatchOperationResult\) GetError\(\) \*common.Error](<#BatchOperationResult.GetError>)
  - [func \(x \*BatchOperationResult\) GetGeneratedKeys\(\) \[\]\*anypb.Any](<#BatchOperationResult.GetGeneratedKeys>)
  - [func \(x \*BatchOperationResult\) GetSuccess\(\) bool](<#BatchOperationResult.GetSuccess>)
  - [func \(x \*BatchOperationResult\) HasAffectedRows\(\) bool](<#BatchOperationResult.HasAffectedRows>)
  - [func \(x \*BatchOperationResult\) HasError\(\) bool](<#BatchOperationResult.HasError>)
  - [func \(x \*BatchOperationResult\) HasSuccess\(\) bool](<#BatchOperationResult.HasSuccess>)
  - [func \(\*BatchOperationResult\) ProtoMessage\(\)](<#BatchOperationResult.ProtoMessage>)
  - [func \(x \*BatchOperationResult\) ProtoReflect\(\) protoreflect.Message](<#BatchOperationResult.ProtoReflect>)
  - [func \(x \*BatchOperationResult\) Reset\(\)](<#BatchOperationResult.Reset>)
  - [func \(x \*BatchOperationResult\) SetAffectedRows\(v int64\)](<#BatchOperationResult.SetAffectedRows>)
  - [func \(x \*BatchOperationResult\) SetError\(v \*common.Error\)](<#BatchOperationResult.SetError>)
  - [func \(x \*BatchOperationResult\) SetGeneratedKeys\(v \[\]\*anypb.Any\)](<#BatchOperationResult.SetGeneratedKeys>)
  - [func \(x \*BatchOperationResult\) SetSuccess\(v bool\)](<#BatchOperationResult.SetSuccess>)
  - [func \(x \*BatchOperationResult\) String\(\) string](<#BatchOperationResult.String>)
- [type BatchOperationResult\_builder](<#BatchOperationResult_builder>)
  - [func \(b0 BatchOperationResult\_builder\) Build\(\) \*BatchOperationResult](<#BatchOperationResult_builder.Build>)
- [type BeginTransactionRequest](<#BeginTransactionRequest>)
  - [func \(x \*BeginTransactionRequest\) ClearDatabase\(\)](<#BeginTransactionRequest.ClearDatabase>)
  - [func \(x \*BeginTransactionRequest\) ClearMetadata\(\)](<#BeginTransactionRequest.ClearMetadata>)
  - [func \(x \*BeginTransactionRequest\) ClearOptions\(\)](<#BeginTransactionRequest.ClearOptions>)
  - [func \(x \*BeginTransactionRequest\) GetDatabase\(\) string](<#BeginTransactionRequest.GetDatabase>)
  - [func \(x \*BeginTransactionRequest\) GetMetadata\(\) \*common.RequestMetadata](<#BeginTransactionRequest.GetMetadata>)
  - [func \(x \*BeginTransactionRequest\) GetOptions\(\) \*TransactionOptions](<#BeginTransactionRequest.GetOptions>)
  - [func \(x \*BeginTransactionRequest\) HasDatabase\(\) bool](<#BeginTransactionRequest.HasDatabase>)
  - [func \(x \*BeginTransactionRequest\) HasMetadata\(\) bool](<#BeginTransactionRequest.HasMetadata>)
  - [func \(x \*BeginTransactionRequest\) HasOptions\(\) bool](<#BeginTransactionRequest.HasOptions>)
  - [func \(\*BeginTransactionRequest\) ProtoMessage\(\)](<#BeginTransactionRequest.ProtoMessage>)
  - [func \(x \*BeginTransactionRequest\) ProtoReflect\(\) protoreflect.Message](<#BeginTransactionRequest.ProtoReflect>)
  - [func \(x \*BeginTransactionRequest\) Reset\(\)](<#BeginTransactionRequest.Reset>)
  - [func \(x \*BeginTransactionRequest\) SetDatabase\(v string\)](<#BeginTransactionRequest.SetDatabase>)
  - [func \(x \*BeginTransactionRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#BeginTransactionRequest.SetMetadata>)
  - [func \(x \*BeginTransactionRequest\) SetOptions\(v \*TransactionOptions\)](<#BeginTransactionRequest.SetOptions>)
  - [func \(x \*BeginTransactionRequest\) String\(\) string](<#BeginTransactionRequest.String>)
- [type BeginTransactionRequest\_builder](<#BeginTransactionRequest_builder>)
  - [func \(b0 BeginTransactionRequest\_builder\) Build\(\) \*BeginTransactionRequest](<#BeginTransactionRequest_builder.Build>)
- [type BeginTransactionResponse](<#BeginTransactionResponse>)
  - [func \(x \*BeginTransactionResponse\) ClearStartedAt\(\)](<#BeginTransactionResponse.ClearStartedAt>)
  - [func \(x \*BeginTransactionResponse\) ClearTransactionId\(\)](<#BeginTransactionResponse.ClearTransactionId>)
  - [func \(x \*BeginTransactionResponse\) GetStartedAt\(\) \*timestamppb.Timestamp](<#BeginTransactionResponse.GetStartedAt>)
  - [func \(x \*BeginTransactionResponse\) GetTransactionId\(\) string](<#BeginTransactionResponse.GetTransactionId>)
  - [func \(x \*BeginTransactionResponse\) HasStartedAt\(\) bool](<#BeginTransactionResponse.HasStartedAt>)
  - [func \(x \*BeginTransactionResponse\) HasTransactionId\(\) bool](<#BeginTransactionResponse.HasTransactionId>)
  - [func \(\*BeginTransactionResponse\) ProtoMessage\(\)](<#BeginTransactionResponse.ProtoMessage>)
  - [func \(x \*BeginTransactionResponse\) ProtoReflect\(\) protoreflect.Message](<#BeginTransactionResponse.ProtoReflect>)
  - [func \(x \*BeginTransactionResponse\) Reset\(\)](<#BeginTransactionResponse.Reset>)
  - [func \(x \*BeginTransactionResponse\) SetStartedAt\(v \*timestamppb.Timestamp\)](<#BeginTransactionResponse.SetStartedAt>)
  - [func \(x \*BeginTransactionResponse\) SetTransactionId\(v string\)](<#BeginTransactionResponse.SetTransactionId>)
  - [func \(x \*BeginTransactionResponse\) String\(\) string](<#BeginTransactionResponse.String>)
- [type BeginTransactionResponse\_builder](<#BeginTransactionResponse_builder>)
  - [func \(b0 BeginTransactionResponse\_builder\) Build\(\) \*BeginTransactionResponse](<#BeginTransactionResponse_builder.Build>)
- [type CacheAdminServiceClient](<#CacheAdminServiceClient>)
  - [func NewCacheAdminServiceClient\(cc grpc.ClientConnInterface\) CacheAdminServiceClient](<#NewCacheAdminServiceClient>)
- [type CacheAdminServiceServer](<#CacheAdminServiceServer>)
- [type CacheCacheConfig](<#CacheCacheConfig>)
  - [func \(x \*CacheCacheConfig\) ClearDefaultTtl\(\)](<#CacheCacheConfig.ClearDefaultTtl>)
  - [func \(x \*CacheCacheConfig\) ClearEnablePersistence\(\)](<#CacheCacheConfig.ClearEnablePersistence>)
  - [func \(x \*CacheCacheConfig\) ClearEnableStats\(\)](<#CacheCacheConfig.ClearEnableStats>)
  - [func \(x \*CacheCacheConfig\) ClearEvictionPolicy\(\)](<#CacheCacheConfig.ClearEvictionPolicy>)
  - [func \(x \*CacheCacheConfig\) ClearMaxEntries\(\)](<#CacheCacheConfig.ClearMaxEntries>)
  - [func \(x \*CacheCacheConfig\) ClearMaxMemoryBytes\(\)](<#CacheCacheConfig.ClearMaxMemoryBytes>)
  - [func \(x \*CacheCacheConfig\) ClearName\(\)](<#CacheCacheConfig.ClearName>)
  - [func \(x \*CacheCacheConfig\) ClearPersistenceFile\(\)](<#CacheCacheConfig.ClearPersistenceFile>)
  - [func \(x \*CacheCacheConfig\) GetDefaultTtl\(\) \*durationpb.Duration](<#CacheCacheConfig.GetDefaultTtl>)
  - [func \(x \*CacheCacheConfig\) GetEnablePersistence\(\) bool](<#CacheCacheConfig.GetEnablePersistence>)
  - [func \(x \*CacheCacheConfig\) GetEnableStats\(\) bool](<#CacheCacheConfig.GetEnableStats>)
  - [func \(x \*CacheCacheConfig\) GetEvictionPolicy\(\) common.EvictionPolicy](<#CacheCacheConfig.GetEvictionPolicy>)
  - [func \(x \*CacheCacheConfig\) GetMaxEntries\(\) int64](<#CacheCacheConfig.GetMaxEntries>)
  - [func \(x \*CacheCacheConfig\) GetMaxMemoryBytes\(\) int64](<#CacheCacheConfig.GetMaxMemoryBytes>)
  - [func \(x \*CacheCacheConfig\) GetName\(\) string](<#CacheCacheConfig.GetName>)
  - [func \(x \*CacheCacheConfig\) GetPersistenceFile\(\) string](<#CacheCacheConfig.GetPersistenceFile>)
  - [func \(x \*CacheCacheConfig\) HasDefaultTtl\(\) bool](<#CacheCacheConfig.HasDefaultTtl>)
  - [func \(x \*CacheCacheConfig\) HasEnablePersistence\(\) bool](<#CacheCacheConfig.HasEnablePersistence>)
  - [func \(x \*CacheCacheConfig\) HasEnableStats\(\) bool](<#CacheCacheConfig.HasEnableStats>)
  - [func \(x \*CacheCacheConfig\) HasEvictionPolicy\(\) bool](<#CacheCacheConfig.HasEvictionPolicy>)
  - [func \(x \*CacheCacheConfig\) HasMaxEntries\(\) bool](<#CacheCacheConfig.HasMaxEntries>)
  - [func \(x \*CacheCacheConfig\) HasMaxMemoryBytes\(\) bool](<#CacheCacheConfig.HasMaxMemoryBytes>)
  - [func \(x \*CacheCacheConfig\) HasName\(\) bool](<#CacheCacheConfig.HasName>)
  - [func \(x \*CacheCacheConfig\) HasPersistenceFile\(\) bool](<#CacheCacheConfig.HasPersistenceFile>)
  - [func \(\*CacheCacheConfig\) ProtoMessage\(\)](<#CacheCacheConfig.ProtoMessage>)
  - [func \(x \*CacheCacheConfig\) ProtoReflect\(\) protoreflect.Message](<#CacheCacheConfig.ProtoReflect>)
  - [func \(x \*CacheCacheConfig\) Reset\(\)](<#CacheCacheConfig.Reset>)
  - [func \(x \*CacheCacheConfig\) SetDefaultTtl\(v \*durationpb.Duration\)](<#CacheCacheConfig.SetDefaultTtl>)
  - [func \(x \*CacheCacheConfig\) SetEnablePersistence\(v bool\)](<#CacheCacheConfig.SetEnablePersistence>)
  - [func \(x \*CacheCacheConfig\) SetEnableStats\(v bool\)](<#CacheCacheConfig.SetEnableStats>)
  - [func \(x \*CacheCacheConfig\) SetEvictionPolicy\(v common.EvictionPolicy\)](<#CacheCacheConfig.SetEvictionPolicy>)
  - [func \(x \*CacheCacheConfig\) SetMaxEntries\(v int64\)](<#CacheCacheConfig.SetMaxEntries>)
  - [func \(x \*CacheCacheConfig\) SetMaxMemoryBytes\(v int64\)](<#CacheCacheConfig.SetMaxMemoryBytes>)
  - [func \(x \*CacheCacheConfig\) SetName\(v string\)](<#CacheCacheConfig.SetName>)
  - [func \(x \*CacheCacheConfig\) SetPersistenceFile\(v string\)](<#CacheCacheConfig.SetPersistenceFile>)
  - [func \(x \*CacheCacheConfig\) String\(\) string](<#CacheCacheConfig.String>)
- [type CacheCacheConfig\_builder](<#CacheCacheConfig_builder>)
  - [func \(b0 CacheCacheConfig\_builder\) Build\(\) \*CacheCacheConfig](<#CacheCacheConfig_builder.Build>)
- [type CacheDeleteRequest](<#CacheDeleteRequest>)
  - [func \(x \*CacheDeleteRequest\) ClearKey\(\)](<#CacheDeleteRequest.ClearKey>)
  - [func \(x \*CacheDeleteRequest\) ClearMetadata\(\)](<#CacheDeleteRequest.ClearMetadata>)
  - [func \(x \*CacheDeleteRequest\) ClearNamespace\(\)](<#CacheDeleteRequest.ClearNamespace>)
  - [func \(x \*CacheDeleteRequest\) GetKey\(\) string](<#CacheDeleteRequest.GetKey>)
  - [func \(x \*CacheDeleteRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheDeleteRequest.GetMetadata>)
  - [func \(x \*CacheDeleteRequest\) GetNamespace\(\) string](<#CacheDeleteRequest.GetNamespace>)
  - [func \(x \*CacheDeleteRequest\) HasKey\(\) bool](<#CacheDeleteRequest.HasKey>)
  - [func \(x \*CacheDeleteRequest\) HasMetadata\(\) bool](<#CacheDeleteRequest.HasMetadata>)
  - [func \(x \*CacheDeleteRequest\) HasNamespace\(\) bool](<#CacheDeleteRequest.HasNamespace>)
  - [func \(\*CacheDeleteRequest\) ProtoMessage\(\)](<#CacheDeleteRequest.ProtoMessage>)
  - [func \(x \*CacheDeleteRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheDeleteRequest.ProtoReflect>)
  - [func \(x \*CacheDeleteRequest\) Reset\(\)](<#CacheDeleteRequest.Reset>)
  - [func \(x \*CacheDeleteRequest\) SetKey\(v string\)](<#CacheDeleteRequest.SetKey>)
  - [func \(x \*CacheDeleteRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheDeleteRequest.SetMetadata>)
  - [func \(x \*CacheDeleteRequest\) SetNamespace\(v string\)](<#CacheDeleteRequest.SetNamespace>)
  - [func \(x \*CacheDeleteRequest\) String\(\) string](<#CacheDeleteRequest.String>)
- [type CacheDeleteRequest\_builder](<#CacheDeleteRequest_builder>)
  - [func \(b0 CacheDeleteRequest\_builder\) Build\(\) \*CacheDeleteRequest](<#CacheDeleteRequest_builder.Build>)
- [type CacheDeleteResponse](<#CacheDeleteResponse>)
  - [func \(x \*CacheDeleteResponse\) ClearDeletedCount\(\)](<#CacheDeleteResponse.ClearDeletedCount>)
  - [func \(x \*CacheDeleteResponse\) ClearError\(\)](<#CacheDeleteResponse.ClearError>)
  - [func \(x \*CacheDeleteResponse\) ClearSuccess\(\)](<#CacheDeleteResponse.ClearSuccess>)
  - [func \(x \*CacheDeleteResponse\) GetDeletedCount\(\) int32](<#CacheDeleteResponse.GetDeletedCount>)
  - [func \(x \*CacheDeleteResponse\) GetError\(\) \*common.Error](<#CacheDeleteResponse.GetError>)
  - [func \(x \*CacheDeleteResponse\) GetSuccess\(\) bool](<#CacheDeleteResponse.GetSuccess>)
  - [func \(x \*CacheDeleteResponse\) HasDeletedCount\(\) bool](<#CacheDeleteResponse.HasDeletedCount>)
  - [func \(x \*CacheDeleteResponse\) HasError\(\) bool](<#CacheDeleteResponse.HasError>)
  - [func \(x \*CacheDeleteResponse\) HasSuccess\(\) bool](<#CacheDeleteResponse.HasSuccess>)
  - [func \(\*CacheDeleteResponse\) ProtoMessage\(\)](<#CacheDeleteResponse.ProtoMessage>)
  - [func \(x \*CacheDeleteResponse\) ProtoReflect\(\) protoreflect.Message](<#CacheDeleteResponse.ProtoReflect>)
  - [func \(x \*CacheDeleteResponse\) Reset\(\)](<#CacheDeleteResponse.Reset>)
  - [func \(x \*CacheDeleteResponse\) SetDeletedCount\(v int32\)](<#CacheDeleteResponse.SetDeletedCount>)
  - [func \(x \*CacheDeleteResponse\) SetError\(v \*common.Error\)](<#CacheDeleteResponse.SetError>)
  - [func \(x \*CacheDeleteResponse\) SetSuccess\(v bool\)](<#CacheDeleteResponse.SetSuccess>)
  - [func \(x \*CacheDeleteResponse\) String\(\) string](<#CacheDeleteResponse.String>)
- [type CacheDeleteResponse\_builder](<#CacheDeleteResponse_builder>)
  - [func \(b0 CacheDeleteResponse\_builder\) Build\(\) \*CacheDeleteResponse](<#CacheDeleteResponse_builder.Build>)
- [type CacheEntry](<#CacheEntry>)
  - [func \(x \*CacheEntry\) ClearAccessCount\(\)](<#CacheEntry.ClearAccessCount>)
  - [func \(x \*CacheEntry\) ClearCreatedAt\(\)](<#CacheEntry.ClearCreatedAt>)
  - [func \(x \*CacheEntry\) ClearExpiresAt\(\)](<#CacheEntry.ClearExpiresAt>)
  - [func \(x \*CacheEntry\) ClearKey\(\)](<#CacheEntry.ClearKey>)
  - [func \(x \*CacheEntry\) ClearLastAccessedAt\(\)](<#CacheEntry.ClearLastAccessedAt>)
  - [func \(x \*CacheEntry\) ClearNamespace\(\)](<#CacheEntry.ClearNamespace>)
  - [func \(x \*CacheEntry\) ClearSizeBytes\(\)](<#CacheEntry.ClearSizeBytes>)
  - [func \(x \*CacheEntry\) ClearValue\(\)](<#CacheEntry.ClearValue>)
  - [func \(x \*CacheEntry\) GetAccessCount\(\) int64](<#CacheEntry.GetAccessCount>)
  - [func \(x \*CacheEntry\) GetCreatedAt\(\) \*timestamppb.Timestamp](<#CacheEntry.GetCreatedAt>)
  - [func \(x \*CacheEntry\) GetExpiresAt\(\) \*timestamppb.Timestamp](<#CacheEntry.GetExpiresAt>)
  - [func \(x \*CacheEntry\) GetKey\(\) string](<#CacheEntry.GetKey>)
  - [func \(x \*CacheEntry\) GetLastAccessedAt\(\) \*timestamppb.Timestamp](<#CacheEntry.GetLastAccessedAt>)
  - [func \(x \*CacheEntry\) GetMetadata\(\) map\[string\]string](<#CacheEntry.GetMetadata>)
  - [func \(x \*CacheEntry\) GetNamespace\(\) string](<#CacheEntry.GetNamespace>)
  - [func \(x \*CacheEntry\) GetSizeBytes\(\) int64](<#CacheEntry.GetSizeBytes>)
  - [func \(x \*CacheEntry\) GetValue\(\) \*anypb.Any](<#CacheEntry.GetValue>)
  - [func \(x \*CacheEntry\) HasAccessCount\(\) bool](<#CacheEntry.HasAccessCount>)
  - [func \(x \*CacheEntry\) HasCreatedAt\(\) bool](<#CacheEntry.HasCreatedAt>)
  - [func \(x \*CacheEntry\) HasExpiresAt\(\) bool](<#CacheEntry.HasExpiresAt>)
  - [func \(x \*CacheEntry\) HasKey\(\) bool](<#CacheEntry.HasKey>)
  - [func \(x \*CacheEntry\) HasLastAccessedAt\(\) bool](<#CacheEntry.HasLastAccessedAt>)
  - [func \(x \*CacheEntry\) HasNamespace\(\) bool](<#CacheEntry.HasNamespace>)
  - [func \(x \*CacheEntry\) HasSizeBytes\(\) bool](<#CacheEntry.HasSizeBytes>)
  - [func \(x \*CacheEntry\) HasValue\(\) bool](<#CacheEntry.HasValue>)
  - [func \(\*CacheEntry\) ProtoMessage\(\)](<#CacheEntry.ProtoMessage>)
  - [func \(x \*CacheEntry\) ProtoReflect\(\) protoreflect.Message](<#CacheEntry.ProtoReflect>)
  - [func \(x \*CacheEntry\) Reset\(\)](<#CacheEntry.Reset>)
  - [func \(x \*CacheEntry\) SetAccessCount\(v int64\)](<#CacheEntry.SetAccessCount>)
  - [func \(x \*CacheEntry\) SetCreatedAt\(v \*timestamppb.Timestamp\)](<#CacheEntry.SetCreatedAt>)
  - [func \(x \*CacheEntry\) SetExpiresAt\(v \*timestamppb.Timestamp\)](<#CacheEntry.SetExpiresAt>)
  - [func \(x \*CacheEntry\) SetKey\(v string\)](<#CacheEntry.SetKey>)
  - [func \(x \*CacheEntry\) SetLastAccessedAt\(v \*timestamppb.Timestamp\)](<#CacheEntry.SetLastAccessedAt>)
  - [func \(x \*CacheEntry\) SetMetadata\(v map\[string\]string\)](<#CacheEntry.SetMetadata>)
  - [func \(x \*CacheEntry\) SetNamespace\(v string\)](<#CacheEntry.SetNamespace>)
  - [func \(x \*CacheEntry\) SetSizeBytes\(v int64\)](<#CacheEntry.SetSizeBytes>)
  - [func \(x \*CacheEntry\) SetValue\(v \*anypb.Any\)](<#CacheEntry.SetValue>)
  - [func \(x \*CacheEntry\) String\(\) string](<#CacheEntry.String>)
- [type CacheEntry\_builder](<#CacheEntry_builder>)
  - [func \(b0 CacheEntry\_builder\) Build\(\) \*CacheEntry](<#CacheEntry_builder.Build>)
- [type CacheGetStatsRequest](<#CacheGetStatsRequest>)
  - [func \(x \*CacheGetStatsRequest\) ClearMetadata\(\)](<#CacheGetStatsRequest.ClearMetadata>)
  - [func \(x \*CacheGetStatsRequest\) ClearNamespace\(\)](<#CacheGetStatsRequest.ClearNamespace>)
  - [func \(x \*CacheGetStatsRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheGetStatsRequest.GetMetadata>)
  - [func \(x \*CacheGetStatsRequest\) GetNamespace\(\) string](<#CacheGetStatsRequest.GetNamespace>)
  - [func \(x \*CacheGetStatsRequest\) HasMetadata\(\) bool](<#CacheGetStatsRequest.HasMetadata>)
  - [func \(x \*CacheGetStatsRequest\) HasNamespace\(\) bool](<#CacheGetStatsRequest.HasNamespace>)
  - [func \(\*CacheGetStatsRequest\) ProtoMessage\(\)](<#CacheGetStatsRequest.ProtoMessage>)
  - [func \(x \*CacheGetStatsRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheGetStatsRequest.ProtoReflect>)
  - [func \(x \*CacheGetStatsRequest\) Reset\(\)](<#CacheGetStatsRequest.Reset>)
  - [func \(x \*CacheGetStatsRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheGetStatsRequest.SetMetadata>)
  - [func \(x \*CacheGetStatsRequest\) SetNamespace\(v string\)](<#CacheGetStatsRequest.SetNamespace>)
  - [func \(x \*CacheGetStatsRequest\) String\(\) string](<#CacheGetStatsRequest.String>)
- [type CacheGetStatsRequest\_builder](<#CacheGetStatsRequest_builder>)
  - [func \(b0 CacheGetStatsRequest\_builder\) Build\(\) \*CacheGetStatsRequest](<#CacheGetStatsRequest_builder.Build>)
- [type CacheGetStatsResponse](<#CacheGetStatsResponse>)
  - [func \(x \*CacheGetStatsResponse\) ClearCacheHits\(\)](<#CacheGetStatsResponse.ClearCacheHits>)
  - [func \(x \*CacheGetStatsResponse\) ClearCacheMisses\(\)](<#CacheGetStatsResponse.ClearCacheMisses>)
  - [func \(x \*CacheGetStatsResponse\) ClearError\(\)](<#CacheGetStatsResponse.ClearError>)
  - [func \(x \*CacheGetStatsResponse\) ClearEvictedItems\(\)](<#CacheGetStatsResponse.ClearEvictedItems>)
  - [func \(x \*CacheGetStatsResponse\) ClearHitRatio\(\)](<#CacheGetStatsResponse.ClearHitRatio>)
  - [func \(x \*CacheGetStatsResponse\) ClearMemoryLimit\(\)](<#CacheGetStatsResponse.ClearMemoryLimit>)
  - [func \(x \*CacheGetStatsResponse\) ClearMemoryUsage\(\)](<#CacheGetStatsResponse.ClearMemoryUsage>)
  - [func \(x \*CacheGetStatsResponse\) ClearSuccess\(\)](<#CacheGetStatsResponse.ClearSuccess>)
  - [func \(x \*CacheGetStatsResponse\) ClearTotalItems\(\)](<#CacheGetStatsResponse.ClearTotalItems>)
  - [func \(x \*CacheGetStatsResponse\) GetCacheHits\(\) int64](<#CacheGetStatsResponse.GetCacheHits>)
  - [func \(x \*CacheGetStatsResponse\) GetCacheMisses\(\) int64](<#CacheGetStatsResponse.GetCacheMisses>)
  - [func \(x \*CacheGetStatsResponse\) GetError\(\) \*common.Error](<#CacheGetStatsResponse.GetError>)
  - [func \(x \*CacheGetStatsResponse\) GetEvictedItems\(\) int64](<#CacheGetStatsResponse.GetEvictedItems>)
  - [func \(x \*CacheGetStatsResponse\) GetHitRatio\(\) float64](<#CacheGetStatsResponse.GetHitRatio>)
  - [func \(x \*CacheGetStatsResponse\) GetMemoryLimit\(\) int64](<#CacheGetStatsResponse.GetMemoryLimit>)
  - [func \(x \*CacheGetStatsResponse\) GetMemoryUsage\(\) int64](<#CacheGetStatsResponse.GetMemoryUsage>)
  - [func \(x \*CacheGetStatsResponse\) GetSuccess\(\) bool](<#CacheGetStatsResponse.GetSuccess>)
  - [func \(x \*CacheGetStatsResponse\) GetTotalItems\(\) int64](<#CacheGetStatsResponse.GetTotalItems>)
  - [func \(x \*CacheGetStatsResponse\) HasCacheHits\(\) bool](<#CacheGetStatsResponse.HasCacheHits>)
  - [func \(x \*CacheGetStatsResponse\) HasCacheMisses\(\) bool](<#CacheGetStatsResponse.HasCacheMisses>)
  - [func \(x \*CacheGetStatsResponse\) HasError\(\) bool](<#CacheGetStatsResponse.HasError>)
  - [func \(x \*CacheGetStatsResponse\) HasEvictedItems\(\) bool](<#CacheGetStatsResponse.HasEvictedItems>)
  - [func \(x \*CacheGetStatsResponse\) HasHitRatio\(\) bool](<#CacheGetStatsResponse.HasHitRatio>)
  - [func \(x \*CacheGetStatsResponse\) HasMemoryLimit\(\) bool](<#CacheGetStatsResponse.HasMemoryLimit>)
  - [func \(x \*CacheGetStatsResponse\) HasMemoryUsage\(\) bool](<#CacheGetStatsResponse.HasMemoryUsage>)
  - [func \(x \*CacheGetStatsResponse\) HasSuccess\(\) bool](<#CacheGetStatsResponse.HasSuccess>)
  - [func \(x \*CacheGetStatsResponse\) HasTotalItems\(\) bool](<#CacheGetStatsResponse.HasTotalItems>)
  - [func \(\*CacheGetStatsResponse\) ProtoMessage\(\)](<#CacheGetStatsResponse.ProtoMessage>)
  - [func \(x \*CacheGetStatsResponse\) ProtoReflect\(\) protoreflect.Message](<#CacheGetStatsResponse.ProtoReflect>)
  - [func \(x \*CacheGetStatsResponse\) Reset\(\)](<#CacheGetStatsResponse.Reset>)
  - [func \(x \*CacheGetStatsResponse\) SetCacheHits\(v int64\)](<#CacheGetStatsResponse.SetCacheHits>)
  - [func \(x \*CacheGetStatsResponse\) SetCacheMisses\(v int64\)](<#CacheGetStatsResponse.SetCacheMisses>)
  - [func \(x \*CacheGetStatsResponse\) SetError\(v \*common.Error\)](<#CacheGetStatsResponse.SetError>)
  - [func \(x \*CacheGetStatsResponse\) SetEvictedItems\(v int64\)](<#CacheGetStatsResponse.SetEvictedItems>)
  - [func \(x \*CacheGetStatsResponse\) SetHitRatio\(v float64\)](<#CacheGetStatsResponse.SetHitRatio>)
  - [func \(x \*CacheGetStatsResponse\) SetMemoryLimit\(v int64\)](<#CacheGetStatsResponse.SetMemoryLimit>)
  - [func \(x \*CacheGetStatsResponse\) SetMemoryUsage\(v int64\)](<#CacheGetStatsResponse.SetMemoryUsage>)
  - [func \(x \*CacheGetStatsResponse\) SetSuccess\(v bool\)](<#CacheGetStatsResponse.SetSuccess>)
  - [func \(x \*CacheGetStatsResponse\) SetTotalItems\(v int64\)](<#CacheGetStatsResponse.SetTotalItems>)
  - [func \(x \*CacheGetStatsResponse\) String\(\) string](<#CacheGetStatsResponse.String>)
- [type CacheGetStatsResponse\_builder](<#CacheGetStatsResponse_builder>)
  - [func \(b0 CacheGetStatsResponse\_builder\) Build\(\) \*CacheGetStatsResponse](<#CacheGetStatsResponse_builder.Build>)
- [type CacheInfo](<#CacheInfo>)
  - [func \(x \*CacheInfo\) ClearCacheType\(\)](<#CacheInfo.ClearCacheType>)
  - [func \(x \*CacheInfo\) ClearCreatedAt\(\)](<#CacheInfo.ClearCreatedAt>)
  - [func \(x \*CacheInfo\) ClearDescription\(\)](<#CacheInfo.ClearDescription>)
  - [func \(x \*CacheInfo\) ClearHealthStatus\(\)](<#CacheInfo.ClearHealthStatus>)
  - [func \(x \*CacheInfo\) ClearInstanceId\(\)](<#CacheInfo.ClearInstanceId>)
  - [func \(x \*CacheInfo\) ClearLastAccessed\(\)](<#CacheInfo.ClearLastAccessed>)
  - [func \(x \*CacheInfo\) ClearName\(\)](<#CacheInfo.ClearName>)
  - [func \(x \*CacheInfo\) ClearVersion\(\)](<#CacheInfo.ClearVersion>)
  - [func \(x \*CacheInfo\) GetCacheType\(\) string](<#CacheInfo.GetCacheType>)
  - [func \(x \*CacheInfo\) GetCreatedAt\(\) \*timestamppb.Timestamp](<#CacheInfo.GetCreatedAt>)
  - [func \(x \*CacheInfo\) GetDescription\(\) string](<#CacheInfo.GetDescription>)
  - [func \(x \*CacheInfo\) GetHealthStatus\(\) common.CommonHealthStatus](<#CacheInfo.GetHealthStatus>)
  - [func \(x \*CacheInfo\) GetInstanceId\(\) string](<#CacheInfo.GetInstanceId>)
  - [func \(x \*CacheInfo\) GetLastAccessed\(\) \*timestamppb.Timestamp](<#CacheInfo.GetLastAccessed>)
  - [func \(x \*CacheInfo\) GetMetadata\(\) map\[string\]string](<#CacheInfo.GetMetadata>)
  - [func \(x \*CacheInfo\) GetName\(\) string](<#CacheInfo.GetName>)
  - [func \(x \*CacheInfo\) GetVersion\(\) string](<#CacheInfo.GetVersion>)
  - [func \(x \*CacheInfo\) HasCacheType\(\) bool](<#CacheInfo.HasCacheType>)
  - [func \(x \*CacheInfo\) HasCreatedAt\(\) bool](<#CacheInfo.HasCreatedAt>)
  - [func \(x \*CacheInfo\) HasDescription\(\) bool](<#CacheInfo.HasDescription>)
  - [func \(x \*CacheInfo\) HasHealthStatus\(\) bool](<#CacheInfo.HasHealthStatus>)
  - [func \(x \*CacheInfo\) HasInstanceId\(\) bool](<#CacheInfo.HasInstanceId>)
  - [func \(x \*CacheInfo\) HasLastAccessed\(\) bool](<#CacheInfo.HasLastAccessed>)
  - [func \(x \*CacheInfo\) HasName\(\) bool](<#CacheInfo.HasName>)
  - [func \(x \*CacheInfo\) HasVersion\(\) bool](<#CacheInfo.HasVersion>)
  - [func \(\*CacheInfo\) ProtoMessage\(\)](<#CacheInfo.ProtoMessage>)
  - [func \(x \*CacheInfo\) ProtoReflect\(\) protoreflect.Message](<#CacheInfo.ProtoReflect>)
  - [func \(x \*CacheInfo\) Reset\(\)](<#CacheInfo.Reset>)
  - [func \(x \*CacheInfo\) SetCacheType\(v string\)](<#CacheInfo.SetCacheType>)
  - [func \(x \*CacheInfo\) SetCreatedAt\(v \*timestamppb.Timestamp\)](<#CacheInfo.SetCreatedAt>)
  - [func \(x \*CacheInfo\) SetDescription\(v string\)](<#CacheInfo.SetDescription>)
  - [func \(x \*CacheInfo\) SetHealthStatus\(v common.CommonHealthStatus\)](<#CacheInfo.SetHealthStatus>)
  - [func \(x \*CacheInfo\) SetInstanceId\(v string\)](<#CacheInfo.SetInstanceId>)
  - [func \(x \*CacheInfo\) SetLastAccessed\(v \*timestamppb.Timestamp\)](<#CacheInfo.SetLastAccessed>)
  - [func \(x \*CacheInfo\) SetMetadata\(v map\[string\]string\)](<#CacheInfo.SetMetadata>)
  - [func \(x \*CacheInfo\) SetName\(v string\)](<#CacheInfo.SetName>)
  - [func \(x \*CacheInfo\) SetVersion\(v string\)](<#CacheInfo.SetVersion>)
  - [func \(x \*CacheInfo\) String\(\) string](<#CacheInfo.String>)
- [type CacheInfo\_builder](<#CacheInfo_builder>)
  - [func \(b0 CacheInfo\_builder\) Build\(\) \*CacheInfo](<#CacheInfo_builder.Build>)
- [type CacheListSubscriptionsRequest](<#CacheListSubscriptionsRequest>)
  - [func \(x \*CacheListSubscriptionsRequest\) ClearMetadata\(\)](<#CacheListSubscriptionsRequest.ClearMetadata>)
  - [func \(x \*CacheListSubscriptionsRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheListSubscriptionsRequest.GetMetadata>)
  - [func \(x \*CacheListSubscriptionsRequest\) HasMetadata\(\) bool](<#CacheListSubscriptionsRequest.HasMetadata>)
  - [func \(\*CacheListSubscriptionsRequest\) ProtoMessage\(\)](<#CacheListSubscriptionsRequest.ProtoMessage>)
  - [func \(x \*CacheListSubscriptionsRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheListSubscriptionsRequest.ProtoReflect>)
  - [func \(x \*CacheListSubscriptionsRequest\) Reset\(\)](<#CacheListSubscriptionsRequest.Reset>)
  - [func \(x \*CacheListSubscriptionsRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheListSubscriptionsRequest.SetMetadata>)
  - [func \(x \*CacheListSubscriptionsRequest\) String\(\) string](<#CacheListSubscriptionsRequest.String>)
- [type CacheListSubscriptionsRequest\_builder](<#CacheListSubscriptionsRequest_builder>)
  - [func \(b0 CacheListSubscriptionsRequest\_builder\) Build\(\) \*CacheListSubscriptionsRequest](<#CacheListSubscriptionsRequest_builder.Build>)
- [type CacheMetrics](<#CacheMetrics>)
  - [func \(x \*CacheMetrics\) ClearActiveConnections\(\)](<#CacheMetrics.ClearActiveConnections>)
  - [func \(x \*CacheMetrics\) ClearAvgResponseTime\(\)](<#CacheMetrics.ClearAvgResponseTime>)
  - [func \(x \*CacheMetrics\) ClearCollectedAt\(\)](<#CacheMetrics.ClearCollectedAt>)
  - [func \(x \*CacheMetrics\) ClearCpuUsagePercent\(\)](<#CacheMetrics.ClearCpuUsagePercent>)
  - [func \(x \*CacheMetrics\) ClearMemoryUsagePercent\(\)](<#CacheMetrics.ClearMemoryUsagePercent>)
  - [func \(x \*CacheMetrics\) ClearNetworkBytesIn\(\)](<#CacheMetrics.ClearNetworkBytesIn>)
  - [func \(x \*CacheMetrics\) ClearNetworkBytesOut\(\)](<#CacheMetrics.ClearNetworkBytesOut>)
  - [func \(x \*CacheMetrics\) ClearOpsPerSecond\(\)](<#CacheMetrics.ClearOpsPerSecond>)
  - [func \(x \*CacheMetrics\) ClearP95ResponseTime\(\)](<#CacheMetrics.ClearP95ResponseTime>)
  - [func \(x \*CacheMetrics\) ClearP99ResponseTime\(\)](<#CacheMetrics.ClearP99ResponseTime>)
  - [func \(x \*CacheMetrics\) ClearReadsPerSecond\(\)](<#CacheMetrics.ClearReadsPerSecond>)
  - [func \(x \*CacheMetrics\) ClearTotalConnections\(\)](<#CacheMetrics.ClearTotalConnections>)
  - [func \(x \*CacheMetrics\) ClearWritesPerSecond\(\)](<#CacheMetrics.ClearWritesPerSecond>)
  - [func \(x \*CacheMetrics\) GetActiveConnections\(\) int64](<#CacheMetrics.GetActiveConnections>)
  - [func \(x \*CacheMetrics\) GetAvgResponseTime\(\) \*durationpb.Duration](<#CacheMetrics.GetAvgResponseTime>)
  - [func \(x \*CacheMetrics\) GetCollectedAt\(\) \*timestamppb.Timestamp](<#CacheMetrics.GetCollectedAt>)
  - [func \(x \*CacheMetrics\) GetCpuUsagePercent\(\) float64](<#CacheMetrics.GetCpuUsagePercent>)
  - [func \(x \*CacheMetrics\) GetMemoryUsagePercent\(\) float64](<#CacheMetrics.GetMemoryUsagePercent>)
  - [func \(x \*CacheMetrics\) GetNetworkBytesIn\(\) int64](<#CacheMetrics.GetNetworkBytesIn>)
  - [func \(x \*CacheMetrics\) GetNetworkBytesOut\(\) int64](<#CacheMetrics.GetNetworkBytesOut>)
  - [func \(x \*CacheMetrics\) GetOpsPerSecond\(\) float64](<#CacheMetrics.GetOpsPerSecond>)
  - [func \(x \*CacheMetrics\) GetP95ResponseTime\(\) \*durationpb.Duration](<#CacheMetrics.GetP95ResponseTime>)
  - [func \(x \*CacheMetrics\) GetP99ResponseTime\(\) \*durationpb.Duration](<#CacheMetrics.GetP99ResponseTime>)
  - [func \(x \*CacheMetrics\) GetReadsPerSecond\(\) float64](<#CacheMetrics.GetReadsPerSecond>)
  - [func \(x \*CacheMetrics\) GetTotalConnections\(\) int64](<#CacheMetrics.GetTotalConnections>)
  - [func \(x \*CacheMetrics\) GetWritesPerSecond\(\) float64](<#CacheMetrics.GetWritesPerSecond>)
  - [func \(x \*CacheMetrics\) HasActiveConnections\(\) bool](<#CacheMetrics.HasActiveConnections>)
  - [func \(x \*CacheMetrics\) HasAvgResponseTime\(\) bool](<#CacheMetrics.HasAvgResponseTime>)
  - [func \(x \*CacheMetrics\) HasCollectedAt\(\) bool](<#CacheMetrics.HasCollectedAt>)
  - [func \(x \*CacheMetrics\) HasCpuUsagePercent\(\) bool](<#CacheMetrics.HasCpuUsagePercent>)
  - [func \(x \*CacheMetrics\) HasMemoryUsagePercent\(\) bool](<#CacheMetrics.HasMemoryUsagePercent>)
  - [func \(x \*CacheMetrics\) HasNetworkBytesIn\(\) bool](<#CacheMetrics.HasNetworkBytesIn>)
  - [func \(x \*CacheMetrics\) HasNetworkBytesOut\(\) bool](<#CacheMetrics.HasNetworkBytesOut>)
  - [func \(x \*CacheMetrics\) HasOpsPerSecond\(\) bool](<#CacheMetrics.HasOpsPerSecond>)
  - [func \(x \*CacheMetrics\) HasP95ResponseTime\(\) bool](<#CacheMetrics.HasP95ResponseTime>)
  - [func \(x \*CacheMetrics\) HasP99ResponseTime\(\) bool](<#CacheMetrics.HasP99ResponseTime>)
  - [func \(x \*CacheMetrics\) HasReadsPerSecond\(\) bool](<#CacheMetrics.HasReadsPerSecond>)
  - [func \(x \*CacheMetrics\) HasTotalConnections\(\) bool](<#CacheMetrics.HasTotalConnections>)
  - [func \(x \*CacheMetrics\) HasWritesPerSecond\(\) bool](<#CacheMetrics.HasWritesPerSecond>)
  - [func \(\*CacheMetrics\) ProtoMessage\(\)](<#CacheMetrics.ProtoMessage>)
  - [func \(x \*CacheMetrics\) ProtoReflect\(\) protoreflect.Message](<#CacheMetrics.ProtoReflect>)
  - [func \(x \*CacheMetrics\) Reset\(\)](<#CacheMetrics.Reset>)
  - [func \(x \*CacheMetrics\) SetActiveConnections\(v int64\)](<#CacheMetrics.SetActiveConnections>)
  - [func \(x \*CacheMetrics\) SetAvgResponseTime\(v \*durationpb.Duration\)](<#CacheMetrics.SetAvgResponseTime>)
  - [func \(x \*CacheMetrics\) SetCollectedAt\(v \*timestamppb.Timestamp\)](<#CacheMetrics.SetCollectedAt>)
  - [func \(x \*CacheMetrics\) SetCpuUsagePercent\(v float64\)](<#CacheMetrics.SetCpuUsagePercent>)
  - [func \(x \*CacheMetrics\) SetMemoryUsagePercent\(v float64\)](<#CacheMetrics.SetMemoryUsagePercent>)
  - [func \(x \*CacheMetrics\) SetNetworkBytesIn\(v int64\)](<#CacheMetrics.SetNetworkBytesIn>)
  - [func \(x \*CacheMetrics\) SetNetworkBytesOut\(v int64\)](<#CacheMetrics.SetNetworkBytesOut>)
  - [func \(x \*CacheMetrics\) SetOpsPerSecond\(v float64\)](<#CacheMetrics.SetOpsPerSecond>)
  - [func \(x \*CacheMetrics\) SetP95ResponseTime\(v \*durationpb.Duration\)](<#CacheMetrics.SetP95ResponseTime>)
  - [func \(x \*CacheMetrics\) SetP99ResponseTime\(v \*durationpb.Duration\)](<#CacheMetrics.SetP99ResponseTime>)
  - [func \(x \*CacheMetrics\) SetReadsPerSecond\(v float64\)](<#CacheMetrics.SetReadsPerSecond>)
  - [func \(x \*CacheMetrics\) SetTotalConnections\(v int64\)](<#CacheMetrics.SetTotalConnections>)
  - [func \(x \*CacheMetrics\) SetWritesPerSecond\(v float64\)](<#CacheMetrics.SetWritesPerSecond>)
  - [func \(x \*CacheMetrics\) String\(\) string](<#CacheMetrics.String>)
- [type CacheMetrics\_builder](<#CacheMetrics_builder>)
  - [func \(b0 CacheMetrics\_builder\) Build\(\) \*CacheMetrics](<#CacheMetrics_builder.Build>)
- [type CacheOperationResult](<#CacheOperationResult>)
  - [func \(x \*CacheOperationResult\) ClearDurationMicroseconds\(\)](<#CacheOperationResult.ClearDurationMicroseconds>)
  - [func \(x \*CacheOperationResult\) ClearError\(\)](<#CacheOperationResult.ClearError>)
  - [func \(x \*CacheOperationResult\) ClearItemsAffected\(\)](<#CacheOperationResult.ClearItemsAffected>)
  - [func \(x \*CacheOperationResult\) ClearKey\(\)](<#CacheOperationResult.ClearKey>)
  - [func \(x \*CacheOperationResult\) ClearNamespace\(\)](<#CacheOperationResult.ClearNamespace>)
  - [func \(x \*CacheOperationResult\) ClearOperationType\(\)](<#CacheOperationResult.ClearOperationType>)
  - [func \(x \*CacheOperationResult\) ClearSuccess\(\)](<#CacheOperationResult.ClearSuccess>)
  - [func \(x \*CacheOperationResult\) ClearTimestamp\(\)](<#CacheOperationResult.ClearTimestamp>)
  - [func \(x \*CacheOperationResult\) GetDurationMicroseconds\(\) int64](<#CacheOperationResult.GetDurationMicroseconds>)
  - [func \(x \*CacheOperationResult\) GetError\(\) \*common.Error](<#CacheOperationResult.GetError>)
  - [func \(x \*CacheOperationResult\) GetItemsAffected\(\) int64](<#CacheOperationResult.GetItemsAffected>)
  - [func \(x \*CacheOperationResult\) GetKey\(\) string](<#CacheOperationResult.GetKey>)
  - [func \(x \*CacheOperationResult\) GetMetadata\(\) map\[string\]string](<#CacheOperationResult.GetMetadata>)
  - [func \(x \*CacheOperationResult\) GetNamespace\(\) string](<#CacheOperationResult.GetNamespace>)
  - [func \(x \*CacheOperationResult\) GetOperationType\(\) string](<#CacheOperationResult.GetOperationType>)
  - [func \(x \*CacheOperationResult\) GetSuccess\(\) bool](<#CacheOperationResult.GetSuccess>)
  - [func \(x \*CacheOperationResult\) GetTimestamp\(\) \*timestamppb.Timestamp](<#CacheOperationResult.GetTimestamp>)
  - [func \(x \*CacheOperationResult\) HasDurationMicroseconds\(\) bool](<#CacheOperationResult.HasDurationMicroseconds>)
  - [func \(x \*CacheOperationResult\) HasError\(\) bool](<#CacheOperationResult.HasError>)
  - [func \(x \*CacheOperationResult\) HasItemsAffected\(\) bool](<#CacheOperationResult.HasItemsAffected>)
  - [func \(x \*CacheOperationResult\) HasKey\(\) bool](<#CacheOperationResult.HasKey>)
  - [func \(x \*CacheOperationResult\) HasNamespace\(\) bool](<#CacheOperationResult.HasNamespace>)
  - [func \(x \*CacheOperationResult\) HasOperationType\(\) bool](<#CacheOperationResult.HasOperationType>)
  - [func \(x \*CacheOperationResult\) HasSuccess\(\) bool](<#CacheOperationResult.HasSuccess>)
  - [func \(x \*CacheOperationResult\) HasTimestamp\(\) bool](<#CacheOperationResult.HasTimestamp>)
  - [func \(\*CacheOperationResult\) ProtoMessage\(\)](<#CacheOperationResult.ProtoMessage>)
  - [func \(x \*CacheOperationResult\) ProtoReflect\(\) protoreflect.Message](<#CacheOperationResult.ProtoReflect>)
  - [func \(x \*CacheOperationResult\) Reset\(\)](<#CacheOperationResult.Reset>)
  - [func \(x \*CacheOperationResult\) SetDurationMicroseconds\(v int64\)](<#CacheOperationResult.SetDurationMicroseconds>)
  - [func \(x \*CacheOperationResult\) SetError\(v \*common.Error\)](<#CacheOperationResult.SetError>)
  - [func \(x \*CacheOperationResult\) SetItemsAffected\(v int64\)](<#CacheOperationResult.SetItemsAffected>)
  - [func \(x \*CacheOperationResult\) SetKey\(v string\)](<#CacheOperationResult.SetKey>)
  - [func \(x \*CacheOperationResult\) SetMetadata\(v map\[string\]string\)](<#CacheOperationResult.SetMetadata>)
  - [func \(x \*CacheOperationResult\) SetNamespace\(v string\)](<#CacheOperationResult.SetNamespace>)
  - [func \(x \*CacheOperationResult\) SetOperationType\(v string\)](<#CacheOperationResult.SetOperationType>)
  - [func \(x \*CacheOperationResult\) SetSuccess\(v bool\)](<#CacheOperationResult.SetSuccess>)
  - [func \(x \*CacheOperationResult\) SetTimestamp\(v \*timestamppb.Timestamp\)](<#CacheOperationResult.SetTimestamp>)
  - [func \(x \*CacheOperationResult\) String\(\) string](<#CacheOperationResult.String>)
- [type CacheOperationResult\_builder](<#CacheOperationResult_builder>)
  - [func \(b0 CacheOperationResult\_builder\) Build\(\) \*CacheOperationResult](<#CacheOperationResult_builder.Build>)
- [type CachePublishRequest](<#CachePublishRequest>)
  - [func \(x \*CachePublishRequest\) ClearMetadata\(\)](<#CachePublishRequest.ClearMetadata>)
  - [func \(x \*CachePublishRequest\) ClearPayload\(\)](<#CachePublishRequest.ClearPayload>)
  - [func \(x \*CachePublishRequest\) ClearTopic\(\)](<#CachePublishRequest.ClearTopic>)
  - [func \(x \*CachePublishRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CachePublishRequest.GetMetadata>)
  - [func \(x \*CachePublishRequest\) GetPayload\(\) \*anypb.Any](<#CachePublishRequest.GetPayload>)
  - [func \(x \*CachePublishRequest\) GetTopic\(\) string](<#CachePublishRequest.GetTopic>)
  - [func \(x \*CachePublishRequest\) HasMetadata\(\) bool](<#CachePublishRequest.HasMetadata>)
  - [func \(x \*CachePublishRequest\) HasPayload\(\) bool](<#CachePublishRequest.HasPayload>)
  - [func \(x \*CachePublishRequest\) HasTopic\(\) bool](<#CachePublishRequest.HasTopic>)
  - [func \(\*CachePublishRequest\) ProtoMessage\(\)](<#CachePublishRequest.ProtoMessage>)
  - [func \(x \*CachePublishRequest\) ProtoReflect\(\) protoreflect.Message](<#CachePublishRequest.ProtoReflect>)
  - [func \(x \*CachePublishRequest\) Reset\(\)](<#CachePublishRequest.Reset>)
  - [func \(x \*CachePublishRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CachePublishRequest.SetMetadata>)
  - [func \(x \*CachePublishRequest\) SetPayload\(v \*anypb.Any\)](<#CachePublishRequest.SetPayload>)
  - [func \(x \*CachePublishRequest\) SetTopic\(v string\)](<#CachePublishRequest.SetTopic>)
  - [func \(x \*CachePublishRequest\) String\(\) string](<#CachePublishRequest.String>)
- [type CachePublishRequest\_builder](<#CachePublishRequest_builder>)
  - [func \(b0 CachePublishRequest\_builder\) Build\(\) \*CachePublishRequest](<#CachePublishRequest_builder.Build>)
- [type CacheServiceClient](<#CacheServiceClient>)
  - [func NewCacheServiceClient\(cc grpc.ClientConnInterface\) CacheServiceClient](<#NewCacheServiceClient>)
- [type CacheServiceServer](<#CacheServiceServer>)
- [type CacheStats](<#CacheStats>)
  - [func \(x \*CacheStats\) ClearAvgAccessTimeMs\(\)](<#CacheStats.ClearAvgAccessTimeMs>)
  - [func \(x \*CacheStats\) ClearCacheHits\(\)](<#CacheStats.ClearCacheHits>)
  - [func \(x \*CacheStats\) ClearCacheMisses\(\)](<#CacheStats.ClearCacheMisses>)
  - [func \(x \*CacheStats\) ClearEvictedItems\(\)](<#CacheStats.ClearEvictedItems>)
  - [func \(x \*CacheStats\) ClearExpiredItems\(\)](<#CacheStats.ClearExpiredItems>)
  - [func \(x \*CacheStats\) ClearHitRatio\(\)](<#CacheStats.ClearHitRatio>)
  - [func \(x \*CacheStats\) ClearLastReset\(\)](<#CacheStats.ClearLastReset>)
  - [func \(x \*CacheStats\) ClearMemoryLimit\(\)](<#CacheStats.ClearMemoryLimit>)
  - [func \(x \*CacheStats\) ClearMemoryUsage\(\)](<#CacheStats.ClearMemoryUsage>)
  - [func \(x \*CacheStats\) ClearTotalItems\(\)](<#CacheStats.ClearTotalItems>)
  - [func \(x \*CacheStats\) ClearUptimeSeconds\(\)](<#CacheStats.ClearUptimeSeconds>)
  - [func \(x \*CacheStats\) GetAvgAccessTimeMs\(\) float64](<#CacheStats.GetAvgAccessTimeMs>)
  - [func \(x \*CacheStats\) GetCacheHits\(\) int64](<#CacheStats.GetCacheHits>)
  - [func \(x \*CacheStats\) GetCacheMisses\(\) int64](<#CacheStats.GetCacheMisses>)
  - [func \(x \*CacheStats\) GetEvictedItems\(\) int64](<#CacheStats.GetEvictedItems>)
  - [func \(x \*CacheStats\) GetExpiredItems\(\) int64](<#CacheStats.GetExpiredItems>)
  - [func \(x \*CacheStats\) GetHitRatio\(\) float64](<#CacheStats.GetHitRatio>)
  - [func \(x \*CacheStats\) GetLastReset\(\) \*timestamppb.Timestamp](<#CacheStats.GetLastReset>)
  - [func \(x \*CacheStats\) GetMemoryLimit\(\) int64](<#CacheStats.GetMemoryLimit>)
  - [func \(x \*CacheStats\) GetMemoryUsage\(\) int64](<#CacheStats.GetMemoryUsage>)
  - [func \(x \*CacheStats\) GetTotalItems\(\) int64](<#CacheStats.GetTotalItems>)
  - [func \(x \*CacheStats\) GetUptimeSeconds\(\) int64](<#CacheStats.GetUptimeSeconds>)
  - [func \(x \*CacheStats\) HasAvgAccessTimeMs\(\) bool](<#CacheStats.HasAvgAccessTimeMs>)
  - [func \(x \*CacheStats\) HasCacheHits\(\) bool](<#CacheStats.HasCacheHits>)
  - [func \(x \*CacheStats\) HasCacheMisses\(\) bool](<#CacheStats.HasCacheMisses>)
  - [func \(x \*CacheStats\) HasEvictedItems\(\) bool](<#CacheStats.HasEvictedItems>)
  - [func \(x \*CacheStats\) HasExpiredItems\(\) bool](<#CacheStats.HasExpiredItems>)
  - [func \(x \*CacheStats\) HasHitRatio\(\) bool](<#CacheStats.HasHitRatio>)
  - [func \(x \*CacheStats\) HasLastReset\(\) bool](<#CacheStats.HasLastReset>)
  - [func \(x \*CacheStats\) HasMemoryLimit\(\) bool](<#CacheStats.HasMemoryLimit>)
  - [func \(x \*CacheStats\) HasMemoryUsage\(\) bool](<#CacheStats.HasMemoryUsage>)
  - [func \(x \*CacheStats\) HasTotalItems\(\) bool](<#CacheStats.HasTotalItems>)
  - [func \(x \*CacheStats\) HasUptimeSeconds\(\) bool](<#CacheStats.HasUptimeSeconds>)
  - [func \(\*CacheStats\) ProtoMessage\(\)](<#CacheStats.ProtoMessage>)
  - [func \(x \*CacheStats\) ProtoReflect\(\) protoreflect.Message](<#CacheStats.ProtoReflect>)
  - [func \(x \*CacheStats\) Reset\(\)](<#CacheStats.Reset>)
  - [func \(x \*CacheStats\) SetAvgAccessTimeMs\(v float64\)](<#CacheStats.SetAvgAccessTimeMs>)
  - [func \(x \*CacheStats\) SetCacheHits\(v int64\)](<#CacheStats.SetCacheHits>)
  - [func \(x \*CacheStats\) SetCacheMisses\(v int64\)](<#CacheStats.SetCacheMisses>)
  - [func \(x \*CacheStats\) SetEvictedItems\(v int64\)](<#CacheStats.SetEvictedItems>)
  - [func \(x \*CacheStats\) SetExpiredItems\(v int64\)](<#CacheStats.SetExpiredItems>)
  - [func \(x \*CacheStats\) SetHitRatio\(v float64\)](<#CacheStats.SetHitRatio>)
  - [func \(x \*CacheStats\) SetLastReset\(v \*timestamppb.Timestamp\)](<#CacheStats.SetLastReset>)
  - [func \(x \*CacheStats\) SetMemoryLimit\(v int64\)](<#CacheStats.SetMemoryLimit>)
  - [func \(x \*CacheStats\) SetMemoryUsage\(v int64\)](<#CacheStats.SetMemoryUsage>)
  - [func \(x \*CacheStats\) SetTotalItems\(v int64\)](<#CacheStats.SetTotalItems>)
  - [func \(x \*CacheStats\) SetUptimeSeconds\(v int64\)](<#CacheStats.SetUptimeSeconds>)
  - [func \(x \*CacheStats\) String\(\) string](<#CacheStats.String>)
- [type CacheStats\_builder](<#CacheStats_builder>)
  - [func \(b0 CacheStats\_builder\) Build\(\) \*CacheStats](<#CacheStats_builder.Build>)
- [type CacheSubscribeRequest](<#CacheSubscribeRequest>)
  - [func \(x \*CacheSubscribeRequest\) ClearMetadata\(\)](<#CacheSubscribeRequest.ClearMetadata>)
  - [func \(x \*CacheSubscribeRequest\) ClearTopic\(\)](<#CacheSubscribeRequest.ClearTopic>)
  - [func \(x \*CacheSubscribeRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheSubscribeRequest.GetMetadata>)
  - [func \(x \*CacheSubscribeRequest\) GetTopic\(\) string](<#CacheSubscribeRequest.GetTopic>)
  - [func \(x \*CacheSubscribeRequest\) HasMetadata\(\) bool](<#CacheSubscribeRequest.HasMetadata>)
  - [func \(x \*CacheSubscribeRequest\) HasTopic\(\) bool](<#CacheSubscribeRequest.HasTopic>)
  - [func \(\*CacheSubscribeRequest\) ProtoMessage\(\)](<#CacheSubscribeRequest.ProtoMessage>)
  - [func \(x \*CacheSubscribeRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheSubscribeRequest.ProtoReflect>)
  - [func \(x \*CacheSubscribeRequest\) Reset\(\)](<#CacheSubscribeRequest.Reset>)
  - [func \(x \*CacheSubscribeRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheSubscribeRequest.SetMetadata>)
  - [func \(x \*CacheSubscribeRequest\) SetTopic\(v string\)](<#CacheSubscribeRequest.SetTopic>)
  - [func \(x \*CacheSubscribeRequest\) String\(\) string](<#CacheSubscribeRequest.String>)
- [type CacheSubscribeRequest\_builder](<#CacheSubscribeRequest_builder>)
  - [func \(b0 CacheSubscribeRequest\_builder\) Build\(\) \*CacheSubscribeRequest](<#CacheSubscribeRequest_builder.Build>)
- [type CacheUnsubscribeRequest](<#CacheUnsubscribeRequest>)
  - [func \(x \*CacheUnsubscribeRequest\) ClearMetadata\(\)](<#CacheUnsubscribeRequest.ClearMetadata>)
  - [func \(x \*CacheUnsubscribeRequest\) ClearTopic\(\)](<#CacheUnsubscribeRequest.ClearTopic>)
  - [func \(x \*CacheUnsubscribeRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheUnsubscribeRequest.GetMetadata>)
  - [func \(x \*CacheUnsubscribeRequest\) GetTopic\(\) string](<#CacheUnsubscribeRequest.GetTopic>)
  - [func \(x \*CacheUnsubscribeRequest\) HasMetadata\(\) bool](<#CacheUnsubscribeRequest.HasMetadata>)
  - [func \(x \*CacheUnsubscribeRequest\) HasTopic\(\) bool](<#CacheUnsubscribeRequest.HasTopic>)
  - [func \(\*CacheUnsubscribeRequest\) ProtoMessage\(\)](<#CacheUnsubscribeRequest.ProtoMessage>)
  - [func \(x \*CacheUnsubscribeRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheUnsubscribeRequest.ProtoReflect>)
  - [func \(x \*CacheUnsubscribeRequest\) Reset\(\)](<#CacheUnsubscribeRequest.Reset>)
  - [func \(x \*CacheUnsubscribeRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheUnsubscribeRequest.SetMetadata>)
  - [func \(x \*CacheUnsubscribeRequest\) SetTopic\(v string\)](<#CacheUnsubscribeRequest.SetTopic>)
  - [func \(x \*CacheUnsubscribeRequest\) String\(\) string](<#CacheUnsubscribeRequest.String>)
- [type CacheUnsubscribeRequest\_builder](<#CacheUnsubscribeRequest_builder>)
  - [func \(b0 CacheUnsubscribeRequest\_builder\) Build\(\) \*CacheUnsubscribeRequest](<#CacheUnsubscribeRequest_builder.Build>)
- [type CacheWatchRequest](<#CacheWatchRequest>)
  - [func \(x \*CacheWatchRequest\) ClearKey\(\)](<#CacheWatchRequest.ClearKey>)
  - [func \(x \*CacheWatchRequest\) ClearMetadata\(\)](<#CacheWatchRequest.ClearMetadata>)
  - [func \(x \*CacheWatchRequest\) GetKey\(\) string](<#CacheWatchRequest.GetKey>)
  - [func \(x \*CacheWatchRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CacheWatchRequest.GetMetadata>)
  - [func \(x \*CacheWatchRequest\) HasKey\(\) bool](<#CacheWatchRequest.HasKey>)
  - [func \(x \*CacheWatchRequest\) HasMetadata\(\) bool](<#CacheWatchRequest.HasMetadata>)
  - [func \(\*CacheWatchRequest\) ProtoMessage\(\)](<#CacheWatchRequest.ProtoMessage>)
  - [func \(x \*CacheWatchRequest\) ProtoReflect\(\) protoreflect.Message](<#CacheWatchRequest.ProtoReflect>)
  - [func \(x \*CacheWatchRequest\) Reset\(\)](<#CacheWatchRequest.Reset>)
  - [func \(x \*CacheWatchRequest\) SetKey\(v string\)](<#CacheWatchRequest.SetKey>)
  - [func \(x \*CacheWatchRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CacheWatchRequest.SetMetadata>)
  - [func \(x \*CacheWatchRequest\) String\(\) string](<#CacheWatchRequest.String>)
- [type CacheWatchRequest\_builder](<#CacheWatchRequest_builder>)
  - [func \(b0 CacheWatchRequest\_builder\) Build\(\) \*CacheWatchRequest](<#CacheWatchRequest_builder.Build>)
- [type ClearRequest](<#ClearRequest>)
  - [func \(x \*ClearRequest\) ClearMetadata\(\)](<#ClearRequest.ClearMetadata>)
  - [func \(x \*ClearRequest\) ClearNamespace\(\)](<#ClearRequest.ClearNamespace>)
  - [func \(x \*ClearRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ClearRequest.GetMetadata>)
  - [func \(x \*ClearRequest\) GetNamespace\(\) string](<#ClearRequest.GetNamespace>)
  - [func \(x \*ClearRequest\) HasMetadata\(\) bool](<#ClearRequest.HasMetadata>)
  - [func \(x \*ClearRequest\) HasNamespace\(\) bool](<#ClearRequest.HasNamespace>)
  - [func \(\*ClearRequest\) ProtoMessage\(\)](<#ClearRequest.ProtoMessage>)
  - [func \(x \*ClearRequest\) ProtoReflect\(\) protoreflect.Message](<#ClearRequest.ProtoReflect>)
  - [func \(x \*ClearRequest\) Reset\(\)](<#ClearRequest.Reset>)
  - [func \(x \*ClearRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ClearRequest.SetMetadata>)
  - [func \(x \*ClearRequest\) SetNamespace\(v string\)](<#ClearRequest.SetNamespace>)
  - [func \(x \*ClearRequest\) String\(\) string](<#ClearRequest.String>)
- [type ClearRequest\_builder](<#ClearRequest_builder>)
  - [func \(b0 ClearRequest\_builder\) Build\(\) \*ClearRequest](<#ClearRequest_builder.Build>)
- [type ClearResponse](<#ClearResponse>)
  - [func \(x \*ClearResponse\) ClearClearedCount\(\)](<#ClearResponse.ClearClearedCount>)
  - [func \(x \*ClearResponse\) ClearError\(\)](<#ClearResponse.ClearError>)
  - [func \(x \*ClearResponse\) ClearSuccess\(\)](<#ClearResponse.ClearSuccess>)
  - [func \(x \*ClearResponse\) GetClearedCount\(\) int64](<#ClearResponse.GetClearedCount>)
  - [func \(x \*ClearResponse\) GetError\(\) \*common.Error](<#ClearResponse.GetError>)
  - [func \(x \*ClearResponse\) GetSuccess\(\) bool](<#ClearResponse.GetSuccess>)
  - [func \(x \*ClearResponse\) HasClearedCount\(\) bool](<#ClearResponse.HasClearedCount>)
  - [func \(x \*ClearResponse\) HasError\(\) bool](<#ClearResponse.HasError>)
  - [func \(x \*ClearResponse\) HasSuccess\(\) bool](<#ClearResponse.HasSuccess>)
  - [func \(\*ClearResponse\) ProtoMessage\(\)](<#ClearResponse.ProtoMessage>)
  - [func \(x \*ClearResponse\) ProtoReflect\(\) protoreflect.Message](<#ClearResponse.ProtoReflect>)
  - [func \(x \*ClearResponse\) Reset\(\)](<#ClearResponse.Reset>)
  - [func \(x \*ClearResponse\) SetClearedCount\(v int64\)](<#ClearResponse.SetClearedCount>)
  - [func \(x \*ClearResponse\) SetError\(v \*common.Error\)](<#ClearResponse.SetError>)
  - [func \(x \*ClearResponse\) SetSuccess\(v bool\)](<#ClearResponse.SetSuccess>)
  - [func \(x \*ClearResponse\) String\(\) string](<#ClearResponse.String>)
- [type ClearResponse\_builder](<#ClearResponse_builder>)
  - [func \(b0 ClearResponse\_builder\) Build\(\) \*ClearResponse](<#ClearResponse_builder.Build>)
- [type CockroachConfig](<#CockroachConfig>)
  - [func \(x \*CockroachConfig\) ClearApplicationName\(\)](<#CockroachConfig.ClearApplicationName>)
  - [func \(x \*CockroachConfig\) ClearDatabase\(\)](<#CockroachConfig.ClearDatabase>)
  - [func \(x \*CockroachConfig\) ClearHost\(\)](<#CockroachConfig.ClearHost>)
  - [func \(x \*CockroachConfig\) ClearMaxRetries\(\)](<#CockroachConfig.ClearMaxRetries>)
  - [func \(x \*CockroachConfig\) ClearPassword\(\)](<#CockroachConfig.ClearPassword>)
  - [func \(x \*CockroachConfig\) ClearPort\(\)](<#CockroachConfig.ClearPort>)
  - [func \(x \*CockroachConfig\) ClearRetryBackoffFactor\(\)](<#CockroachConfig.ClearRetryBackoffFactor>)
  - [func \(x \*CockroachConfig\) ClearSslMode\(\)](<#CockroachConfig.ClearSslMode>)
  - [func \(x \*CockroachConfig\) ClearUser\(\)](<#CockroachConfig.ClearUser>)
  - [func \(x \*CockroachConfig\) GetApplicationName\(\) string](<#CockroachConfig.GetApplicationName>)
  - [func \(x \*CockroachConfig\) GetDatabase\(\) string](<#CockroachConfig.GetDatabase>)
  - [func \(x \*CockroachConfig\) GetHost\(\) string](<#CockroachConfig.GetHost>)
  - [func \(x \*CockroachConfig\) GetMaxRetries\(\) int32](<#CockroachConfig.GetMaxRetries>)
  - [func \(x \*CockroachConfig\) GetPassword\(\) string](<#CockroachConfig.GetPassword>)
  - [func \(x \*CockroachConfig\) GetPort\(\) int32](<#CockroachConfig.GetPort>)
  - [func \(x \*CockroachConfig\) GetRetryBackoffFactor\(\) float32](<#CockroachConfig.GetRetryBackoffFactor>)
  - [func \(x \*CockroachConfig\) GetSslMode\(\) string](<#CockroachConfig.GetSslMode>)
  - [func \(x \*CockroachConfig\) GetUser\(\) string](<#CockroachConfig.GetUser>)
  - [func \(x \*CockroachConfig\) HasApplicationName\(\) bool](<#CockroachConfig.HasApplicationName>)
  - [func \(x \*CockroachConfig\) HasDatabase\(\) bool](<#CockroachConfig.HasDatabase>)
  - [func \(x \*CockroachConfig\) HasHost\(\) bool](<#CockroachConfig.HasHost>)
  - [func \(x \*CockroachConfig\) HasMaxRetries\(\) bool](<#CockroachConfig.HasMaxRetries>)
  - [func \(x \*CockroachConfig\) HasPassword\(\) bool](<#CockroachConfig.HasPassword>)
  - [func \(x \*CockroachConfig\) HasPort\(\) bool](<#CockroachConfig.HasPort>)
  - [func \(x \*CockroachConfig\) HasRetryBackoffFactor\(\) bool](<#CockroachConfig.HasRetryBackoffFactor>)
  - [func \(x \*CockroachConfig\) HasSslMode\(\) bool](<#CockroachConfig.HasSslMode>)
  - [func \(x \*CockroachConfig\) HasUser\(\) bool](<#CockroachConfig.HasUser>)
  - [func \(\*CockroachConfig\) ProtoMessage\(\)](<#CockroachConfig.ProtoMessage>)
  - [func \(x \*CockroachConfig\) ProtoReflect\(\) protoreflect.Message](<#CockroachConfig.ProtoReflect>)
  - [func \(x \*CockroachConfig\) Reset\(\)](<#CockroachConfig.Reset>)
  - [func \(x \*CockroachConfig\) SetApplicationName\(v string\)](<#CockroachConfig.SetApplicationName>)
  - [func \(x \*CockroachConfig\) SetDatabase\(v string\)](<#CockroachConfig.SetDatabase>)
  - [func \(x \*CockroachConfig\) SetHost\(v string\)](<#CockroachConfig.SetHost>)
  - [func \(x \*CockroachConfig\) SetMaxRetries\(v int32\)](<#CockroachConfig.SetMaxRetries>)
  - [func \(x \*CockroachConfig\) SetPassword\(v string\)](<#CockroachConfig.SetPassword>)
  - [func \(x \*CockroachConfig\) SetPort\(v int32\)](<#CockroachConfig.SetPort>)
  - [func \(x \*CockroachConfig\) SetRetryBackoffFactor\(v float32\)](<#CockroachConfig.SetRetryBackoffFactor>)
  - [func \(x \*CockroachConfig\) SetSslMode\(v string\)](<#CockroachConfig.SetSslMode>)
  - [func \(x \*CockroachConfig\) SetUser\(v string\)](<#CockroachConfig.SetUser>)
  - [func \(x \*CockroachConfig\) String\(\) string](<#CockroachConfig.String>)
- [type CockroachConfig\_builder](<#CockroachConfig_builder>)
  - [func \(b0 CockroachConfig\_builder\) Build\(\) \*CockroachConfig](<#CockroachConfig_builder.Build>)
- [type ColumnMetadata](<#ColumnMetadata>)
  - [func \(x \*ColumnMetadata\) ClearName\(\)](<#ColumnMetadata.ClearName>)
  - [func \(x \*ColumnMetadata\) ClearNullable\(\)](<#ColumnMetadata.ClearNullable>)
  - [func \(x \*ColumnMetadata\) ClearScale\(\)](<#ColumnMetadata.ClearScale>)
  - [func \(x \*ColumnMetadata\) ClearSize\(\)](<#ColumnMetadata.ClearSize>)
  - [func \(x \*ColumnMetadata\) ClearType\(\)](<#ColumnMetadata.ClearType>)
  - [func \(x \*ColumnMetadata\) GetMetadata\(\) map\[string\]string](<#ColumnMetadata.GetMetadata>)
  - [func \(x \*ColumnMetadata\) GetName\(\) string](<#ColumnMetadata.GetName>)
  - [func \(x \*ColumnMetadata\) GetNullable\(\) bool](<#ColumnMetadata.GetNullable>)
  - [func \(x \*ColumnMetadata\) GetScale\(\) int32](<#ColumnMetadata.GetScale>)
  - [func \(x \*ColumnMetadata\) GetSize\(\) int32](<#ColumnMetadata.GetSize>)
  - [func \(x \*ColumnMetadata\) GetType\(\) string](<#ColumnMetadata.GetType>)
  - [func \(x \*ColumnMetadata\) HasName\(\) bool](<#ColumnMetadata.HasName>)
  - [func \(x \*ColumnMetadata\) HasNullable\(\) bool](<#ColumnMetadata.HasNullable>)
  - [func \(x \*ColumnMetadata\) HasScale\(\) bool](<#ColumnMetadata.HasScale>)
  - [func \(x \*ColumnMetadata\) HasSize\(\) bool](<#ColumnMetadata.HasSize>)
  - [func \(x \*ColumnMetadata\) HasType\(\) bool](<#ColumnMetadata.HasType>)
  - [func \(\*ColumnMetadata\) ProtoMessage\(\)](<#ColumnMetadata.ProtoMessage>)
  - [func \(x \*ColumnMetadata\) ProtoReflect\(\) protoreflect.Message](<#ColumnMetadata.ProtoReflect>)
  - [func \(x \*ColumnMetadata\) Reset\(\)](<#ColumnMetadata.Reset>)
  - [func \(x \*ColumnMetadata\) SetMetadata\(v map\[string\]string\)](<#ColumnMetadata.SetMetadata>)
  - [func \(x \*ColumnMetadata\) SetName\(v string\)](<#ColumnMetadata.SetName>)
  - [func \(x \*ColumnMetadata\) SetNullable\(v bool\)](<#ColumnMetadata.SetNullable>)
  - [func \(x \*ColumnMetadata\) SetScale\(v int32\)](<#ColumnMetadata.SetScale>)
  - [func \(x \*ColumnMetadata\) SetSize\(v int32\)](<#ColumnMetadata.SetSize>)
  - [func \(x \*ColumnMetadata\) SetType\(v string\)](<#ColumnMetadata.SetType>)
  - [func \(x \*ColumnMetadata\) String\(\) string](<#ColumnMetadata.String>)
- [type ColumnMetadata\_builder](<#ColumnMetadata_builder>)
  - [func \(b0 ColumnMetadata\_builder\) Build\(\) \*ColumnMetadata](<#ColumnMetadata_builder.Build>)
- [type CommitTransactionRequest](<#CommitTransactionRequest>)
  - [func \(x \*CommitTransactionRequest\) ClearMetadata\(\)](<#CommitTransactionRequest.ClearMetadata>)
  - [func \(x \*CommitTransactionRequest\) ClearTransactionId\(\)](<#CommitTransactionRequest.ClearTransactionId>)
  - [func \(x \*CommitTransactionRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CommitTransactionRequest.GetMetadata>)
  - [func \(x \*CommitTransactionRequest\) GetTransactionId\(\) string](<#CommitTransactionRequest.GetTransactionId>)
  - [func \(x \*CommitTransactionRequest\) HasMetadata\(\) bool](<#CommitTransactionRequest.HasMetadata>)
  - [func \(x \*CommitTransactionRequest\) HasTransactionId\(\) bool](<#CommitTransactionRequest.HasTransactionId>)
  - [func \(\*CommitTransactionRequest\) ProtoMessage\(\)](<#CommitTransactionRequest.ProtoMessage>)
  - [func \(x \*CommitTransactionRequest\) ProtoReflect\(\) protoreflect.Message](<#CommitTransactionRequest.ProtoReflect>)
  - [func \(x \*CommitTransactionRequest\) Reset\(\)](<#CommitTransactionRequest.Reset>)
  - [func \(x \*CommitTransactionRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CommitTransactionRequest.SetMetadata>)
  - [func \(x \*CommitTransactionRequest\) SetTransactionId\(v string\)](<#CommitTransactionRequest.SetTransactionId>)
  - [func \(x \*CommitTransactionRequest\) String\(\) string](<#CommitTransactionRequest.String>)
- [type CommitTransactionRequest\_builder](<#CommitTransactionRequest_builder>)
  - [func \(b0 CommitTransactionRequest\_builder\) Build\(\) \*CommitTransactionRequest](<#CommitTransactionRequest_builder.Build>)
- [type ConfigurePolicyRequest](<#ConfigurePolicyRequest>)
  - [func \(x \*ConfigurePolicyRequest\) ClearEvictionPolicy\(\)](<#ConfigurePolicyRequest.ClearEvictionPolicy>)
  - [func \(x \*ConfigurePolicyRequest\) ClearMaxTtlSeconds\(\)](<#ConfigurePolicyRequest.ClearMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyRequest\) ClearMemoryThresholdPercent\(\)](<#ConfigurePolicyRequest.ClearMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyRequest\) ClearNamespaceId\(\)](<#ConfigurePolicyRequest.ClearNamespaceId>)
  - [func \(x \*ConfigurePolicyRequest\) GetEvictionPolicy\(\) string](<#ConfigurePolicyRequest.GetEvictionPolicy>)
  - [func \(x \*ConfigurePolicyRequest\) GetMaxTtlSeconds\(\) int32](<#ConfigurePolicyRequest.GetMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyRequest\) GetMemoryThresholdPercent\(\) float64](<#ConfigurePolicyRequest.GetMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyRequest\) GetNamespaceId\(\) string](<#ConfigurePolicyRequest.GetNamespaceId>)
  - [func \(x \*ConfigurePolicyRequest\) GetPolicyConfig\(\) map\[string\]string](<#ConfigurePolicyRequest.GetPolicyConfig>)
  - [func \(x \*ConfigurePolicyRequest\) HasEvictionPolicy\(\) bool](<#ConfigurePolicyRequest.HasEvictionPolicy>)
  - [func \(x \*ConfigurePolicyRequest\) HasMaxTtlSeconds\(\) bool](<#ConfigurePolicyRequest.HasMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyRequest\) HasMemoryThresholdPercent\(\) bool](<#ConfigurePolicyRequest.HasMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyRequest\) HasNamespaceId\(\) bool](<#ConfigurePolicyRequest.HasNamespaceId>)
  - [func \(\*ConfigurePolicyRequest\) ProtoMessage\(\)](<#ConfigurePolicyRequest.ProtoMessage>)
  - [func \(x \*ConfigurePolicyRequest\) ProtoReflect\(\) protoreflect.Message](<#ConfigurePolicyRequest.ProtoReflect>)
  - [func \(x \*ConfigurePolicyRequest\) Reset\(\)](<#ConfigurePolicyRequest.Reset>)
  - [func \(x \*ConfigurePolicyRequest\) SetEvictionPolicy\(v string\)](<#ConfigurePolicyRequest.SetEvictionPolicy>)
  - [func \(x \*ConfigurePolicyRequest\) SetMaxTtlSeconds\(v int32\)](<#ConfigurePolicyRequest.SetMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyRequest\) SetMemoryThresholdPercent\(v float64\)](<#ConfigurePolicyRequest.SetMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyRequest\) SetNamespaceId\(v string\)](<#ConfigurePolicyRequest.SetNamespaceId>)
  - [func \(x \*ConfigurePolicyRequest\) SetPolicyConfig\(v map\[string\]string\)](<#ConfigurePolicyRequest.SetPolicyConfig>)
  - [func \(x \*ConfigurePolicyRequest\) String\(\) string](<#ConfigurePolicyRequest.String>)
- [type ConfigurePolicyRequest\_builder](<#ConfigurePolicyRequest_builder>)
  - [func \(b0 ConfigurePolicyRequest\_builder\) Build\(\) \*ConfigurePolicyRequest](<#ConfigurePolicyRequest_builder.Build>)
- [type ConfigurePolicyResponse](<#ConfigurePolicyResponse>)
  - [func \(x \*ConfigurePolicyResponse\) ClearAppliedAt\(\)](<#ConfigurePolicyResponse.ClearAppliedAt>)
  - [func \(x \*ConfigurePolicyResponse\) ClearEvictionPolicy\(\)](<#ConfigurePolicyResponse.ClearEvictionPolicy>)
  - [func \(x \*ConfigurePolicyResponse\) ClearMaxTtlSeconds\(\)](<#ConfigurePolicyResponse.ClearMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyResponse\) ClearMemoryThresholdPercent\(\)](<#ConfigurePolicyResponse.ClearMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyResponse\) ClearNamespaceId\(\)](<#ConfigurePolicyResponse.ClearNamespaceId>)
  - [func \(x \*ConfigurePolicyResponse\) GetAppliedAt\(\) \*timestamppb.Timestamp](<#ConfigurePolicyResponse.GetAppliedAt>)
  - [func \(x \*ConfigurePolicyResponse\) GetEvictionPolicy\(\) string](<#ConfigurePolicyResponse.GetEvictionPolicy>)
  - [func \(x \*ConfigurePolicyResponse\) GetMaxTtlSeconds\(\) int32](<#ConfigurePolicyResponse.GetMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyResponse\) GetMemoryThresholdPercent\(\) float64](<#ConfigurePolicyResponse.GetMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyResponse\) GetNamespaceId\(\) string](<#ConfigurePolicyResponse.GetNamespaceId>)
  - [func \(x \*ConfigurePolicyResponse\) GetNewConfig\(\) map\[string\]string](<#ConfigurePolicyResponse.GetNewConfig>)
  - [func \(x \*ConfigurePolicyResponse\) GetPreviousConfig\(\) map\[string\]string](<#ConfigurePolicyResponse.GetPreviousConfig>)
  - [func \(x \*ConfigurePolicyResponse\) HasAppliedAt\(\) bool](<#ConfigurePolicyResponse.HasAppliedAt>)
  - [func \(x \*ConfigurePolicyResponse\) HasEvictionPolicy\(\) bool](<#ConfigurePolicyResponse.HasEvictionPolicy>)
  - [func \(x \*ConfigurePolicyResponse\) HasMaxTtlSeconds\(\) bool](<#ConfigurePolicyResponse.HasMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyResponse\) HasMemoryThresholdPercent\(\) bool](<#ConfigurePolicyResponse.HasMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyResponse\) HasNamespaceId\(\) bool](<#ConfigurePolicyResponse.HasNamespaceId>)
  - [func \(\*ConfigurePolicyResponse\) ProtoMessage\(\)](<#ConfigurePolicyResponse.ProtoMessage>)
  - [func \(x \*ConfigurePolicyResponse\) ProtoReflect\(\) protoreflect.Message](<#ConfigurePolicyResponse.ProtoReflect>)
  - [func \(x \*ConfigurePolicyResponse\) Reset\(\)](<#ConfigurePolicyResponse.Reset>)
  - [func \(x \*ConfigurePolicyResponse\) SetAppliedAt\(v \*timestamppb.Timestamp\)](<#ConfigurePolicyResponse.SetAppliedAt>)
  - [func \(x \*ConfigurePolicyResponse\) SetEvictionPolicy\(v string\)](<#ConfigurePolicyResponse.SetEvictionPolicy>)
  - [func \(x \*ConfigurePolicyResponse\) SetMaxTtlSeconds\(v int32\)](<#ConfigurePolicyResponse.SetMaxTtlSeconds>)
  - [func \(x \*ConfigurePolicyResponse\) SetMemoryThresholdPercent\(v float64\)](<#ConfigurePolicyResponse.SetMemoryThresholdPercent>)
  - [func \(x \*ConfigurePolicyResponse\) SetNamespaceId\(v string\)](<#ConfigurePolicyResponse.SetNamespaceId>)
  - [func \(x \*ConfigurePolicyResponse\) SetNewConfig\(v map\[string\]string\)](<#ConfigurePolicyResponse.SetNewConfig>)
  - [func \(x \*ConfigurePolicyResponse\) SetPreviousConfig\(v map\[string\]string\)](<#ConfigurePolicyResponse.SetPreviousConfig>)
  - [func \(x \*ConfigurePolicyResponse\) String\(\) string](<#ConfigurePolicyResponse.String>)
- [type ConfigurePolicyResponse\_builder](<#ConfigurePolicyResponse_builder>)
  - [func \(b0 ConfigurePolicyResponse\_builder\) Build\(\) \*ConfigurePolicyResponse](<#ConfigurePolicyResponse_builder.Build>)
- [type ConnectionPoolInfo](<#ConnectionPoolInfo>)
  - [func \(x \*ConnectionPoolInfo\) ClearActiveConnections\(\)](<#ConnectionPoolInfo.ClearActiveConnections>)
  - [func \(x \*ConnectionPoolInfo\) ClearAvgLifetime\(\)](<#ConnectionPoolInfo.ClearAvgLifetime>)
  - [func \(x \*ConnectionPoolInfo\) ClearIdleConnections\(\)](<#ConnectionPoolInfo.ClearIdleConnections>)
  - [func \(x \*ConnectionPoolInfo\) ClearMaxConnections\(\)](<#ConnectionPoolInfo.ClearMaxConnections>)
  - [func \(x \*ConnectionPoolInfo\) ClearStats\(\)](<#ConnectionPoolInfo.ClearStats>)
  - [func \(x \*ConnectionPoolInfo\) GetActiveConnections\(\) int32](<#ConnectionPoolInfo.GetActiveConnections>)
  - [func \(x \*ConnectionPoolInfo\) GetAvgLifetime\(\) \*durationpb.Duration](<#ConnectionPoolInfo.GetAvgLifetime>)
  - [func \(x \*ConnectionPoolInfo\) GetIdleConnections\(\) int32](<#ConnectionPoolInfo.GetIdleConnections>)
  - [func \(x \*ConnectionPoolInfo\) GetMaxConnections\(\) int32](<#ConnectionPoolInfo.GetMaxConnections>)
  - [func \(x \*ConnectionPoolInfo\) GetStats\(\) \*PoolStats](<#ConnectionPoolInfo.GetStats>)
  - [func \(x \*ConnectionPoolInfo\) HasActiveConnections\(\) bool](<#ConnectionPoolInfo.HasActiveConnections>)
  - [func \(x \*ConnectionPoolInfo\) HasAvgLifetime\(\) bool](<#ConnectionPoolInfo.HasAvgLifetime>)
  - [func \(x \*ConnectionPoolInfo\) HasIdleConnections\(\) bool](<#ConnectionPoolInfo.HasIdleConnections>)
  - [func \(x \*ConnectionPoolInfo\) HasMaxConnections\(\) bool](<#ConnectionPoolInfo.HasMaxConnections>)
  - [func \(x \*ConnectionPoolInfo\) HasStats\(\) bool](<#ConnectionPoolInfo.HasStats>)
  - [func \(\*ConnectionPoolInfo\) ProtoMessage\(\)](<#ConnectionPoolInfo.ProtoMessage>)
  - [func \(x \*ConnectionPoolInfo\) ProtoReflect\(\) protoreflect.Message](<#ConnectionPoolInfo.ProtoReflect>)
  - [func \(x \*ConnectionPoolInfo\) Reset\(\)](<#ConnectionPoolInfo.Reset>)
  - [func \(x \*ConnectionPoolInfo\) SetActiveConnections\(v int32\)](<#ConnectionPoolInfo.SetActiveConnections>)
  - [func \(x \*ConnectionPoolInfo\) SetAvgLifetime\(v \*durationpb.Duration\)](<#ConnectionPoolInfo.SetAvgLifetime>)
  - [func \(x \*ConnectionPoolInfo\) SetIdleConnections\(v int32\)](<#ConnectionPoolInfo.SetIdleConnections>)
  - [func \(x \*ConnectionPoolInfo\) SetMaxConnections\(v int32\)](<#ConnectionPoolInfo.SetMaxConnections>)
  - [func \(x \*ConnectionPoolInfo\) SetStats\(v \*PoolStats\)](<#ConnectionPoolInfo.SetStats>)
  - [func \(x \*ConnectionPoolInfo\) String\(\) string](<#ConnectionPoolInfo.String>)
- [type ConnectionPoolInfo\_builder](<#ConnectionPoolInfo_builder>)
  - [func \(b0 ConnectionPoolInfo\_builder\) Build\(\) \*ConnectionPoolInfo](<#ConnectionPoolInfo_builder.Build>)
- [type CreateDatabaseRequest](<#CreateDatabaseRequest>)
  - [func \(x \*CreateDatabaseRequest\) ClearMetadata\(\)](<#CreateDatabaseRequest.ClearMetadata>)
  - [func \(x \*CreateDatabaseRequest\) ClearName\(\)](<#CreateDatabaseRequest.ClearName>)
  - [func \(x \*CreateDatabaseRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CreateDatabaseRequest.GetMetadata>)
  - [func \(x \*CreateDatabaseRequest\) GetName\(\) string](<#CreateDatabaseRequest.GetName>)
  - [func \(x \*CreateDatabaseRequest\) GetOptions\(\) map\[string\]string](<#CreateDatabaseRequest.GetOptions>)
  - [func \(x \*CreateDatabaseRequest\) HasMetadata\(\) bool](<#CreateDatabaseRequest.HasMetadata>)
  - [func \(x \*CreateDatabaseRequest\) HasName\(\) bool](<#CreateDatabaseRequest.HasName>)
  - [func \(\*CreateDatabaseRequest\) ProtoMessage\(\)](<#CreateDatabaseRequest.ProtoMessage>)
  - [func \(x \*CreateDatabaseRequest\) ProtoReflect\(\) protoreflect.Message](<#CreateDatabaseRequest.ProtoReflect>)
  - [func \(x \*CreateDatabaseRequest\) Reset\(\)](<#CreateDatabaseRequest.Reset>)
  - [func \(x \*CreateDatabaseRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CreateDatabaseRequest.SetMetadata>)
  - [func \(x \*CreateDatabaseRequest\) SetName\(v string\)](<#CreateDatabaseRequest.SetName>)
  - [func \(x \*CreateDatabaseRequest\) SetOptions\(v map\[string\]string\)](<#CreateDatabaseRequest.SetOptions>)
  - [func \(x \*CreateDatabaseRequest\) String\(\) string](<#CreateDatabaseRequest.String>)
- [type CreateDatabaseRequest\_builder](<#CreateDatabaseRequest_builder>)
  - [func \(b0 CreateDatabaseRequest\_builder\) Build\(\) \*CreateDatabaseRequest](<#CreateDatabaseRequest_builder.Build>)
- [type CreateDatabaseResponse](<#CreateDatabaseResponse>)
  - [func \(x \*CreateDatabaseResponse\) ClearError\(\)](<#CreateDatabaseResponse.ClearError>)
  - [func \(x \*CreateDatabaseResponse\) ClearSuccess\(\)](<#CreateDatabaseResponse.ClearSuccess>)
  - [func \(x \*CreateDatabaseResponse\) GetError\(\) \*common.Error](<#CreateDatabaseResponse.GetError>)
  - [func \(x \*CreateDatabaseResponse\) GetSuccess\(\) bool](<#CreateDatabaseResponse.GetSuccess>)
  - [func \(x \*CreateDatabaseResponse\) HasError\(\) bool](<#CreateDatabaseResponse.HasError>)
  - [func \(x \*CreateDatabaseResponse\) HasSuccess\(\) bool](<#CreateDatabaseResponse.HasSuccess>)
  - [func \(\*CreateDatabaseResponse\) ProtoMessage\(\)](<#CreateDatabaseResponse.ProtoMessage>)
  - [func \(x \*CreateDatabaseResponse\) ProtoReflect\(\) protoreflect.Message](<#CreateDatabaseResponse.ProtoReflect>)
  - [func \(x \*CreateDatabaseResponse\) Reset\(\)](<#CreateDatabaseResponse.Reset>)
  - [func \(x \*CreateDatabaseResponse\) SetError\(v \*common.Error\)](<#CreateDatabaseResponse.SetError>)
  - [func \(x \*CreateDatabaseResponse\) SetSuccess\(v bool\)](<#CreateDatabaseResponse.SetSuccess>)
  - [func \(x \*CreateDatabaseResponse\) String\(\) string](<#CreateDatabaseResponse.String>)
- [type CreateDatabaseResponse\_builder](<#CreateDatabaseResponse_builder>)
  - [func \(b0 CreateDatabaseResponse\_builder\) Build\(\) \*CreateDatabaseResponse](<#CreateDatabaseResponse_builder.Build>)
- [type CreateNamespaceRequest](<#CreateNamespaceRequest>)
  - [func \(x \*CreateNamespaceRequest\) ClearDefaultTtlSeconds\(\)](<#CreateNamespaceRequest.ClearDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceRequest\) ClearDescription\(\)](<#CreateNamespaceRequest.ClearDescription>)
  - [func \(x \*CreateNamespaceRequest\) ClearMaxKeys\(\)](<#CreateNamespaceRequest.ClearMaxKeys>)
  - [func \(x \*CreateNamespaceRequest\) ClearMaxMemoryBytes\(\)](<#CreateNamespaceRequest.ClearMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceRequest\) ClearName\(\)](<#CreateNamespaceRequest.ClearName>)
  - [func \(x \*CreateNamespaceRequest\) GetConfig\(\) map\[string\]string](<#CreateNamespaceRequest.GetConfig>)
  - [func \(x \*CreateNamespaceRequest\) GetDefaultTtlSeconds\(\) int32](<#CreateNamespaceRequest.GetDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceRequest\) GetDescription\(\) string](<#CreateNamespaceRequest.GetDescription>)
  - [func \(x \*CreateNamespaceRequest\) GetMaxKeys\(\) int64](<#CreateNamespaceRequest.GetMaxKeys>)
  - [func \(x \*CreateNamespaceRequest\) GetMaxMemoryBytes\(\) int64](<#CreateNamespaceRequest.GetMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceRequest\) GetName\(\) string](<#CreateNamespaceRequest.GetName>)
  - [func \(x \*CreateNamespaceRequest\) HasDefaultTtlSeconds\(\) bool](<#CreateNamespaceRequest.HasDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceRequest\) HasDescription\(\) bool](<#CreateNamespaceRequest.HasDescription>)
  - [func \(x \*CreateNamespaceRequest\) HasMaxKeys\(\) bool](<#CreateNamespaceRequest.HasMaxKeys>)
  - [func \(x \*CreateNamespaceRequest\) HasMaxMemoryBytes\(\) bool](<#CreateNamespaceRequest.HasMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceRequest\) HasName\(\) bool](<#CreateNamespaceRequest.HasName>)
  - [func \(\*CreateNamespaceRequest\) ProtoMessage\(\)](<#CreateNamespaceRequest.ProtoMessage>)
  - [func \(x \*CreateNamespaceRequest\) ProtoReflect\(\) protoreflect.Message](<#CreateNamespaceRequest.ProtoReflect>)
  - [func \(x \*CreateNamespaceRequest\) Reset\(\)](<#CreateNamespaceRequest.Reset>)
  - [func \(x \*CreateNamespaceRequest\) SetConfig\(v map\[string\]string\)](<#CreateNamespaceRequest.SetConfig>)
  - [func \(x \*CreateNamespaceRequest\) SetDefaultTtlSeconds\(v int32\)](<#CreateNamespaceRequest.SetDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceRequest\) SetDescription\(v string\)](<#CreateNamespaceRequest.SetDescription>)
  - [func \(x \*CreateNamespaceRequest\) SetMaxKeys\(v int64\)](<#CreateNamespaceRequest.SetMaxKeys>)
  - [func \(x \*CreateNamespaceRequest\) SetMaxMemoryBytes\(v int64\)](<#CreateNamespaceRequest.SetMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceRequest\) SetName\(v string\)](<#CreateNamespaceRequest.SetName>)
  - [func \(x \*CreateNamespaceRequest\) String\(\) string](<#CreateNamespaceRequest.String>)
- [type CreateNamespaceRequest\_builder](<#CreateNamespaceRequest_builder>)
  - [func \(b0 CreateNamespaceRequest\_builder\) Build\(\) \*CreateNamespaceRequest](<#CreateNamespaceRequest_builder.Build>)
- [type CreateNamespaceResponse](<#CreateNamespaceResponse>)
  - [func \(x \*CreateNamespaceResponse\) ClearCreatedAt\(\)](<#CreateNamespaceResponse.ClearCreatedAt>)
  - [func \(x \*CreateNamespaceResponse\) ClearDefaultTtlSeconds\(\)](<#CreateNamespaceResponse.ClearDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceResponse\) ClearDescription\(\)](<#CreateNamespaceResponse.ClearDescription>)
  - [func \(x \*CreateNamespaceResponse\) ClearMaxKeys\(\)](<#CreateNamespaceResponse.ClearMaxKeys>)
  - [func \(x \*CreateNamespaceResponse\) ClearMaxMemoryBytes\(\)](<#CreateNamespaceResponse.ClearMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceResponse\) ClearName\(\)](<#CreateNamespaceResponse.ClearName>)
  - [func \(x \*CreateNamespaceResponse\) ClearNamespaceId\(\)](<#CreateNamespaceResponse.ClearNamespaceId>)
  - [func \(x \*CreateNamespaceResponse\) GetConfig\(\) map\[string\]string](<#CreateNamespaceResponse.GetConfig>)
  - [func \(x \*CreateNamespaceResponse\) GetCreatedAt\(\) \*timestamppb.Timestamp](<#CreateNamespaceResponse.GetCreatedAt>)
  - [func \(x \*CreateNamespaceResponse\) GetDefaultTtlSeconds\(\) int32](<#CreateNamespaceResponse.GetDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceResponse\) GetDescription\(\) string](<#CreateNamespaceResponse.GetDescription>)
  - [func \(x \*CreateNamespaceResponse\) GetMaxKeys\(\) int64](<#CreateNamespaceResponse.GetMaxKeys>)
  - [func \(x \*CreateNamespaceResponse\) GetMaxMemoryBytes\(\) int64](<#CreateNamespaceResponse.GetMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceResponse\) GetName\(\) string](<#CreateNamespaceResponse.GetName>)
  - [func \(x \*CreateNamespaceResponse\) GetNamespaceId\(\) string](<#CreateNamespaceResponse.GetNamespaceId>)
  - [func \(x \*CreateNamespaceResponse\) HasCreatedAt\(\) bool](<#CreateNamespaceResponse.HasCreatedAt>)
  - [func \(x \*CreateNamespaceResponse\) HasDefaultTtlSeconds\(\) bool](<#CreateNamespaceResponse.HasDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceResponse\) HasDescription\(\) bool](<#CreateNamespaceResponse.HasDescription>)
  - [func \(x \*CreateNamespaceResponse\) HasMaxKeys\(\) bool](<#CreateNamespaceResponse.HasMaxKeys>)
  - [func \(x \*CreateNamespaceResponse\) HasMaxMemoryBytes\(\) bool](<#CreateNamespaceResponse.HasMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceResponse\) HasName\(\) bool](<#CreateNamespaceResponse.HasName>)
  - [func \(x \*CreateNamespaceResponse\) HasNamespaceId\(\) bool](<#CreateNamespaceResponse.HasNamespaceId>)
  - [func \(\*CreateNamespaceResponse\) ProtoMessage\(\)](<#CreateNamespaceResponse.ProtoMessage>)
  - [func \(x \*CreateNamespaceResponse\) ProtoReflect\(\) protoreflect.Message](<#CreateNamespaceResponse.ProtoReflect>)
  - [func \(x \*CreateNamespaceResponse\) Reset\(\)](<#CreateNamespaceResponse.Reset>)
  - [func \(x \*CreateNamespaceResponse\) SetConfig\(v map\[string\]string\)](<#CreateNamespaceResponse.SetConfig>)
  - [func \(x \*CreateNamespaceResponse\) SetCreatedAt\(v \*timestamppb.Timestamp\)](<#CreateNamespaceResponse.SetCreatedAt>)
  - [func \(x \*CreateNamespaceResponse\) SetDefaultTtlSeconds\(v int32\)](<#CreateNamespaceResponse.SetDefaultTtlSeconds>)
  - [func \(x \*CreateNamespaceResponse\) SetDescription\(v string\)](<#CreateNamespaceResponse.SetDescription>)
  - [func \(x \*CreateNamespaceResponse\) SetMaxKeys\(v int64\)](<#CreateNamespaceResponse.SetMaxKeys>)
  - [func \(x \*CreateNamespaceResponse\) SetMaxMemoryBytes\(v int64\)](<#CreateNamespaceResponse.SetMaxMemoryBytes>)
  - [func \(x \*CreateNamespaceResponse\) SetName\(v string\)](<#CreateNamespaceResponse.SetName>)
  - [func \(x \*CreateNamespaceResponse\) SetNamespaceId\(v string\)](<#CreateNamespaceResponse.SetNamespaceId>)
  - [func \(x \*CreateNamespaceResponse\) String\(\) string](<#CreateNamespaceResponse.String>)
- [type CreateNamespaceResponse\_builder](<#CreateNamespaceResponse_builder>)
  - [func \(b0 CreateNamespaceResponse\_builder\) Build\(\) \*CreateNamespaceResponse](<#CreateNamespaceResponse_builder.Build>)
- [type CreateSchemaRequest](<#CreateSchemaRequest>)
  - [func \(x \*CreateSchemaRequest\) ClearDatabase\(\)](<#CreateSchemaRequest.ClearDatabase>)
  - [func \(x \*CreateSchemaRequest\) ClearMetadata\(\)](<#CreateSchemaRequest.ClearMetadata>)
  - [func \(x \*CreateSchemaRequest\) ClearSchema\(\)](<#CreateSchemaRequest.ClearSchema>)
  - [func \(x \*CreateSchemaRequest\) GetDatabase\(\) string](<#CreateSchemaRequest.GetDatabase>)
  - [func \(x \*CreateSchemaRequest\) GetMetadata\(\) \*common.RequestMetadata](<#CreateSchemaRequest.GetMetadata>)
  - [func \(x \*CreateSchemaRequest\) GetSchema\(\) string](<#CreateSchemaRequest.GetSchema>)
  - [func \(x \*CreateSchemaRequest\) HasDatabase\(\) bool](<#CreateSchemaRequest.HasDatabase>)
  - [func \(x \*CreateSchemaRequest\) HasMetadata\(\) bool](<#CreateSchemaRequest.HasMetadata>)
  - [func \(x \*CreateSchemaRequest\) HasSchema\(\) bool](<#CreateSchemaRequest.HasSchema>)
  - [func \(\*CreateSchemaRequest\) ProtoMessage\(\)](<#CreateSchemaRequest.ProtoMessage>)
  - [func \(x \*CreateSchemaRequest\) ProtoReflect\(\) protoreflect.Message](<#CreateSchemaRequest.ProtoReflect>)
  - [func \(x \*CreateSchemaRequest\) Reset\(\)](<#CreateSchemaRequest.Reset>)
  - [func \(x \*CreateSchemaRequest\) SetDatabase\(v string\)](<#CreateSchemaRequest.SetDatabase>)
  - [func \(x \*CreateSchemaRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#CreateSchemaRequest.SetMetadata>)
  - [func \(x \*CreateSchemaRequest\) SetSchema\(v string\)](<#CreateSchemaRequest.SetSchema>)
  - [func \(x \*CreateSchemaRequest\) String\(\) string](<#CreateSchemaRequest.String>)
- [type CreateSchemaRequest\_builder](<#CreateSchemaRequest_builder>)
  - [func \(b0 CreateSchemaRequest\_builder\) Build\(\) \*CreateSchemaRequest](<#CreateSchemaRequest_builder.Build>)
- [type CreateSchemaResponse](<#CreateSchemaResponse>)
  - [func \(x \*CreateSchemaResponse\) ClearError\(\)](<#CreateSchemaResponse.ClearError>)
  - [func \(x \*CreateSchemaResponse\) ClearSuccess\(\)](<#CreateSchemaResponse.ClearSuccess>)
  - [func \(x \*CreateSchemaResponse\) GetError\(\) \*common.Error](<#CreateSchemaResponse.GetError>)
  - [func \(x \*CreateSchemaResponse\) GetSuccess\(\) bool](<#CreateSchemaResponse.GetSuccess>)
  - [func \(x \*CreateSchemaResponse\) HasError\(\) bool](<#CreateSchemaResponse.HasError>)
  - [func \(x \*CreateSchemaResponse\) HasSuccess\(\) bool](<#CreateSchemaResponse.HasSuccess>)
  - [func \(\*CreateSchemaResponse\) ProtoMessage\(\)](<#CreateSchemaResponse.ProtoMessage>)
  - [func \(x \*CreateSchemaResponse\) ProtoReflect\(\) protoreflect.Message](<#CreateSchemaResponse.ProtoReflect>)
  - [func \(x \*CreateSchemaResponse\) Reset\(\)](<#CreateSchemaResponse.Reset>)
  - [func \(x \*CreateSchemaResponse\) SetError\(v \*common.Error\)](<#CreateSchemaResponse.SetError>)
  - [func \(x \*CreateSchemaResponse\) SetSuccess\(v bool\)](<#CreateSchemaResponse.SetSuccess>)
  - [func \(x \*CreateSchemaResponse\) String\(\) string](<#CreateSchemaResponse.String>)
- [type CreateSchemaResponse\_builder](<#CreateSchemaResponse_builder>)
  - [func \(b0 CreateSchemaResponse\_builder\) Build\(\) \*CreateSchemaResponse](<#CreateSchemaResponse_builder.Build>)
- [type DatabaseAdminServiceClient](<#DatabaseAdminServiceClient>)
  - [func NewDatabaseAdminServiceClient\(cc grpc.ClientConnInterface\) DatabaseAdminServiceClient](<#NewDatabaseAdminServiceClient>)
- [type DatabaseAdminServiceServer](<#DatabaseAdminServiceServer>)
- [type DatabaseBatchOperation](<#DatabaseBatchOperation>)
  - [func \(x \*DatabaseBatchOperation\) ClearStatement\(\)](<#DatabaseBatchOperation.ClearStatement>)
  - [func \(x \*DatabaseBatchOperation\) GetParameters\(\) \[\]\*QueryParameter](<#DatabaseBatchOperation.GetParameters>)
  - [func \(x \*DatabaseBatchOperation\) GetStatement\(\) string](<#DatabaseBatchOperation.GetStatement>)
  - [func \(x \*DatabaseBatchOperation\) HasStatement\(\) bool](<#DatabaseBatchOperation.HasStatement>)
  - [func \(\*DatabaseBatchOperation\) ProtoMessage\(\)](<#DatabaseBatchOperation.ProtoMessage>)
  - [func \(x \*DatabaseBatchOperation\) ProtoReflect\(\) protoreflect.Message](<#DatabaseBatchOperation.ProtoReflect>)
  - [func \(x \*DatabaseBatchOperation\) Reset\(\)](<#DatabaseBatchOperation.Reset>)
  - [func \(x \*DatabaseBatchOperation\) SetParameters\(v \[\]\*QueryParameter\)](<#DatabaseBatchOperation.SetParameters>)
  - [func \(x \*DatabaseBatchOperation\) SetStatement\(v string\)](<#DatabaseBatchOperation.SetStatement>)
  - [func \(x \*DatabaseBatchOperation\) String\(\) string](<#DatabaseBatchOperation.String>)
- [type DatabaseBatchOperation\_builder](<#DatabaseBatchOperation_builder>)
  - [func \(b0 DatabaseBatchOperation\_builder\) Build\(\) \*DatabaseBatchOperation](<#DatabaseBatchOperation_builder.Build>)
- [type DatabaseBatchStats](<#DatabaseBatchStats>)
  - [func \(x \*DatabaseBatchStats\) ClearFailedOperations\(\)](<#DatabaseBatchStats.ClearFailedOperations>)
  - [func \(x \*DatabaseBatchStats\) ClearSuccessfulOperations\(\)](<#DatabaseBatchStats.ClearSuccessfulOperations>)
  - [func \(x \*DatabaseBatchStats\) ClearTotalAffectedRows\(\)](<#DatabaseBatchStats.ClearTotalAffectedRows>)
  - [func \(x \*DatabaseBatchStats\) ClearTotalTime\(\)](<#DatabaseBatchStats.ClearTotalTime>)
  - [func \(x \*DatabaseBatchStats\) GetFailedOperations\(\) int32](<#DatabaseBatchStats.GetFailedOperations>)
  - [func \(x \*DatabaseBatchStats\) GetSuccessfulOperations\(\) int32](<#DatabaseBatchStats.GetSuccessfulOperations>)
  - [func \(x \*DatabaseBatchStats\) GetTotalAffectedRows\(\) int64](<#DatabaseBatchStats.GetTotalAffectedRows>)
  - [func \(x \*DatabaseBatchStats\) GetTotalTime\(\) \*durationpb.Duration](<#DatabaseBatchStats.GetTotalTime>)
  - [func \(x \*DatabaseBatchStats\) HasFailedOperations\(\) bool](<#DatabaseBatchStats.HasFailedOperations>)
  - [func \(x \*DatabaseBatchStats\) HasSuccessfulOperations\(\) bool](<#DatabaseBatchStats.HasSuccessfulOperations>)
  - [func \(x \*DatabaseBatchStats\) HasTotalAffectedRows\(\) bool](<#DatabaseBatchStats.HasTotalAffectedRows>)
  - [func \(x \*DatabaseBatchStats\) HasTotalTime\(\) bool](<#DatabaseBatchStats.HasTotalTime>)
  - [func \(\*DatabaseBatchStats\) ProtoMessage\(\)](<#DatabaseBatchStats.ProtoMessage>)
  - [func \(x \*DatabaseBatchStats\) ProtoReflect\(\) protoreflect.Message](<#DatabaseBatchStats.ProtoReflect>)
  - [func \(x \*DatabaseBatchStats\) Reset\(\)](<#DatabaseBatchStats.Reset>)
  - [func \(x \*DatabaseBatchStats\) SetFailedOperations\(v int32\)](<#DatabaseBatchStats.SetFailedOperations>)
  - [func \(x \*DatabaseBatchStats\) SetSuccessfulOperations\(v int32\)](<#DatabaseBatchStats.SetSuccessfulOperations>)
  - [func \(x \*DatabaseBatchStats\) SetTotalAffectedRows\(v int64\)](<#DatabaseBatchStats.SetTotalAffectedRows>)
  - [func \(x \*DatabaseBatchStats\) SetTotalTime\(v \*durationpb.Duration\)](<#DatabaseBatchStats.SetTotalTime>)
  - [func \(x \*DatabaseBatchStats\) String\(\) string](<#DatabaseBatchStats.String>)
- [type DatabaseBatchStats\_builder](<#DatabaseBatchStats_builder>)
  - [func \(b0 DatabaseBatchStats\_builder\) Build\(\) \*DatabaseBatchStats](<#DatabaseBatchStats_builder.Build>)
- [type DatabaseHealthCheckRequest](<#DatabaseHealthCheckRequest>)
  - [func \(x \*DatabaseHealthCheckRequest\) ClearMetadata\(\)](<#DatabaseHealthCheckRequest.ClearMetadata>)
  - [func \(x \*DatabaseHealthCheckRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DatabaseHealthCheckRequest.GetMetadata>)
  - [func \(x \*DatabaseHealthCheckRequest\) HasMetadata\(\) bool](<#DatabaseHealthCheckRequest.HasMetadata>)
  - [func \(\*DatabaseHealthCheckRequest\) ProtoMessage\(\)](<#DatabaseHealthCheckRequest.ProtoMessage>)
  - [func \(x \*DatabaseHealthCheckRequest\) ProtoReflect\(\) protoreflect.Message](<#DatabaseHealthCheckRequest.ProtoReflect>)
  - [func \(x \*DatabaseHealthCheckRequest\) Reset\(\)](<#DatabaseHealthCheckRequest.Reset>)
  - [func \(x \*DatabaseHealthCheckRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DatabaseHealthCheckRequest.SetMetadata>)
  - [func \(x \*DatabaseHealthCheckRequest\) String\(\) string](<#DatabaseHealthCheckRequest.String>)
- [type DatabaseHealthCheckRequest\_builder](<#DatabaseHealthCheckRequest_builder>)
  - [func \(b0 DatabaseHealthCheckRequest\_builder\) Build\(\) \*DatabaseHealthCheckRequest](<#DatabaseHealthCheckRequest_builder.Build>)
- [type DatabaseHealthCheckResponse](<#DatabaseHealthCheckResponse>)
  - [func \(x \*DatabaseHealthCheckResponse\) ClearConnectionOk\(\)](<#DatabaseHealthCheckResponse.ClearConnectionOk>)
  - [func \(x \*DatabaseHealthCheckResponse\) ClearError\(\)](<#DatabaseHealthCheckResponse.ClearError>)
  - [func \(x \*DatabaseHealthCheckResponse\) ClearResponseTime\(\)](<#DatabaseHealthCheckResponse.ClearResponseTime>)
  - [func \(x \*DatabaseHealthCheckResponse\) ClearStatus\(\)](<#DatabaseHealthCheckResponse.ClearStatus>)
  - [func \(x \*DatabaseHealthCheckResponse\) GetConnectionOk\(\) bool](<#DatabaseHealthCheckResponse.GetConnectionOk>)
  - [func \(x \*DatabaseHealthCheckResponse\) GetError\(\) \*common.Error](<#DatabaseHealthCheckResponse.GetError>)
  - [func \(x \*DatabaseHealthCheckResponse\) GetResponseTime\(\) \*durationpb.Duration](<#DatabaseHealthCheckResponse.GetResponseTime>)
  - [func \(x \*DatabaseHealthCheckResponse\) GetStatus\(\) common.CommonHealthStatus](<#DatabaseHealthCheckResponse.GetStatus>)
  - [func \(x \*DatabaseHealthCheckResponse\) HasConnectionOk\(\) bool](<#DatabaseHealthCheckResponse.HasConnectionOk>)
  - [func \(x \*DatabaseHealthCheckResponse\) HasError\(\) bool](<#DatabaseHealthCheckResponse.HasError>)
  - [func \(x \*DatabaseHealthCheckResponse\) HasResponseTime\(\) bool](<#DatabaseHealthCheckResponse.HasResponseTime>)
  - [func \(x \*DatabaseHealthCheckResponse\) HasStatus\(\) bool](<#DatabaseHealthCheckResponse.HasStatus>)
  - [func \(\*DatabaseHealthCheckResponse\) ProtoMessage\(\)](<#DatabaseHealthCheckResponse.ProtoMessage>)
  - [func \(x \*DatabaseHealthCheckResponse\) ProtoReflect\(\) protoreflect.Message](<#DatabaseHealthCheckResponse.ProtoReflect>)
  - [func \(x \*DatabaseHealthCheckResponse\) Reset\(\)](<#DatabaseHealthCheckResponse.Reset>)
  - [func \(x \*DatabaseHealthCheckResponse\) SetConnectionOk\(v bool\)](<#DatabaseHealthCheckResponse.SetConnectionOk>)
  - [func \(x \*DatabaseHealthCheckResponse\) SetError\(v \*common.Error\)](<#DatabaseHealthCheckResponse.SetError>)
  - [func \(x \*DatabaseHealthCheckResponse\) SetResponseTime\(v \*durationpb.Duration\)](<#DatabaseHealthCheckResponse.SetResponseTime>)
  - [func \(x \*DatabaseHealthCheckResponse\) SetStatus\(v common.CommonHealthStatus\)](<#DatabaseHealthCheckResponse.SetStatus>)
  - [func \(x \*DatabaseHealthCheckResponse\) String\(\) string](<#DatabaseHealthCheckResponse.String>)
- [type DatabaseHealthCheckResponse\_builder](<#DatabaseHealthCheckResponse_builder>)
  - [func \(b0 DatabaseHealthCheckResponse\_builder\) Build\(\) \*DatabaseHealthCheckResponse](<#DatabaseHealthCheckResponse_builder.Build>)
- [type DatabaseInfo](<#DatabaseInfo>)
  - [func \(x \*DatabaseInfo\) ClearConnectionString\(\)](<#DatabaseInfo.ClearConnectionString>)
  - [func \(x \*DatabaseInfo\) ClearName\(\)](<#DatabaseInfo.ClearName>)
  - [func \(x \*DatabaseInfo\) ClearType\(\)](<#DatabaseInfo.ClearType>)
  - [func \(x \*DatabaseInfo\) ClearVersion\(\)](<#DatabaseInfo.ClearVersion>)
  - [func \(x \*DatabaseInfo\) GetConnectionString\(\) string](<#DatabaseInfo.GetConnectionString>)
  - [func \(x \*DatabaseInfo\) GetFeatures\(\) \[\]string](<#DatabaseInfo.GetFeatures>)
  - [func \(x \*DatabaseInfo\) GetName\(\) string](<#DatabaseInfo.GetName>)
  - [func \(x \*DatabaseInfo\) GetType\(\) string](<#DatabaseInfo.GetType>)
  - [func \(x \*DatabaseInfo\) GetVersion\(\) string](<#DatabaseInfo.GetVersion>)
  - [func \(x \*DatabaseInfo\) HasConnectionString\(\) bool](<#DatabaseInfo.HasConnectionString>)
  - [func \(x \*DatabaseInfo\) HasName\(\) bool](<#DatabaseInfo.HasName>)
  - [func \(x \*DatabaseInfo\) HasType\(\) bool](<#DatabaseInfo.HasType>)
  - [func \(x \*DatabaseInfo\) HasVersion\(\) bool](<#DatabaseInfo.HasVersion>)
  - [func \(\*DatabaseInfo\) ProtoMessage\(\)](<#DatabaseInfo.ProtoMessage>)
  - [func \(x \*DatabaseInfo\) ProtoReflect\(\) protoreflect.Message](<#DatabaseInfo.ProtoReflect>)
  - [func \(x \*DatabaseInfo\) Reset\(\)](<#DatabaseInfo.Reset>)
  - [func \(x \*DatabaseInfo\) SetConnectionString\(v string\)](<#DatabaseInfo.SetConnectionString>)
  - [func \(x \*DatabaseInfo\) SetFeatures\(v \[\]string\)](<#DatabaseInfo.SetFeatures>)
  - [func \(x \*DatabaseInfo\) SetName\(v string\)](<#DatabaseInfo.SetName>)
  - [func \(x \*DatabaseInfo\) SetType\(v string\)](<#DatabaseInfo.SetType>)
  - [func \(x \*DatabaseInfo\) SetVersion\(v string\)](<#DatabaseInfo.SetVersion>)
  - [func \(x \*DatabaseInfo\) String\(\) string](<#DatabaseInfo.String>)
- [type DatabaseInfo\_builder](<#DatabaseInfo_builder>)
  - [func \(b0 DatabaseInfo\_builder\) Build\(\) \*DatabaseInfo](<#DatabaseInfo_builder.Build>)
- [type DatabaseQueryStats](<#DatabaseQueryStats>)
  - [func \(x \*DatabaseQueryStats\) ClearColumnCount\(\)](<#DatabaseQueryStats.ClearColumnCount>)
  - [func \(x \*DatabaseQueryStats\) ClearCostEstimate\(\)](<#DatabaseQueryStats.ClearCostEstimate>)
  - [func \(x \*DatabaseQueryStats\) ClearExecutionTime\(\)](<#DatabaseQueryStats.ClearExecutionTime>)
  - [func \(x \*DatabaseQueryStats\) ClearQueryPlan\(\)](<#DatabaseQueryStats.ClearQueryPlan>)
  - [func \(x \*DatabaseQueryStats\) ClearRowCount\(\)](<#DatabaseQueryStats.ClearRowCount>)
  - [func \(x \*DatabaseQueryStats\) GetColumnCount\(\) int32](<#DatabaseQueryStats.GetColumnCount>)
  - [func \(x \*DatabaseQueryStats\) GetCostEstimate\(\) float64](<#DatabaseQueryStats.GetCostEstimate>)
  - [func \(x \*DatabaseQueryStats\) GetExecutionTime\(\) \*durationpb.Duration](<#DatabaseQueryStats.GetExecutionTime>)
  - [func \(x \*DatabaseQueryStats\) GetQueryPlan\(\) string](<#DatabaseQueryStats.GetQueryPlan>)
  - [func \(x \*DatabaseQueryStats\) GetRowCount\(\) int64](<#DatabaseQueryStats.GetRowCount>)
  - [func \(x \*DatabaseQueryStats\) HasColumnCount\(\) bool](<#DatabaseQueryStats.HasColumnCount>)
  - [func \(x \*DatabaseQueryStats\) HasCostEstimate\(\) bool](<#DatabaseQueryStats.HasCostEstimate>)
  - [func \(x \*DatabaseQueryStats\) HasExecutionTime\(\) bool](<#DatabaseQueryStats.HasExecutionTime>)
  - [func \(x \*DatabaseQueryStats\) HasQueryPlan\(\) bool](<#DatabaseQueryStats.HasQueryPlan>)
  - [func \(x \*DatabaseQueryStats\) HasRowCount\(\) bool](<#DatabaseQueryStats.HasRowCount>)
  - [func \(\*DatabaseQueryStats\) ProtoMessage\(\)](<#DatabaseQueryStats.ProtoMessage>)
  - [func \(x \*DatabaseQueryStats\) ProtoReflect\(\) protoreflect.Message](<#DatabaseQueryStats.ProtoReflect>)
  - [func \(x \*DatabaseQueryStats\) Reset\(\)](<#DatabaseQueryStats.Reset>)
  - [func \(x \*DatabaseQueryStats\) SetColumnCount\(v int32\)](<#DatabaseQueryStats.SetColumnCount>)
  - [func \(x \*DatabaseQueryStats\) SetCostEstimate\(v float64\)](<#DatabaseQueryStats.SetCostEstimate>)
  - [func \(x \*DatabaseQueryStats\) SetExecutionTime\(v \*durationpb.Duration\)](<#DatabaseQueryStats.SetExecutionTime>)
  - [func \(x \*DatabaseQueryStats\) SetQueryPlan\(v string\)](<#DatabaseQueryStats.SetQueryPlan>)
  - [func \(x \*DatabaseQueryStats\) SetRowCount\(v int64\)](<#DatabaseQueryStats.SetRowCount>)
  - [func \(x \*DatabaseQueryStats\) String\(\) string](<#DatabaseQueryStats.String>)
- [type DatabaseQueryStats\_builder](<#DatabaseQueryStats_builder>)
  - [func \(b0 DatabaseQueryStats\_builder\) Build\(\) \*DatabaseQueryStats](<#DatabaseQueryStats_builder.Build>)
- [type DatabaseServiceClient](<#DatabaseServiceClient>)
  - [func NewDatabaseServiceClient\(cc grpc.ClientConnInterface\) DatabaseServiceClient](<#NewDatabaseServiceClient>)
- [type DatabaseServiceServer](<#DatabaseServiceServer>)
- [type DatabaseStatus](<#DatabaseStatus>)
  - [func \(x \*DatabaseStatus\) ClearCode\(\)](<#DatabaseStatus.ClearCode>)
  - [func \(x \*DatabaseStatus\) ClearMessage\(\)](<#DatabaseStatus.ClearMessage>)
  - [func \(x \*DatabaseStatus\) GetCode\(\) common.DatabaseStatusCode](<#DatabaseStatus.GetCode>)
  - [func \(x \*DatabaseStatus\) GetMessage\(\) string](<#DatabaseStatus.GetMessage>)
  - [func \(x \*DatabaseStatus\) HasCode\(\) bool](<#DatabaseStatus.HasCode>)
  - [func \(x \*DatabaseStatus\) HasMessage\(\) bool](<#DatabaseStatus.HasMessage>)
  - [func \(\*DatabaseStatus\) ProtoMessage\(\)](<#DatabaseStatus.ProtoMessage>)
  - [func \(x \*DatabaseStatus\) ProtoReflect\(\) protoreflect.Message](<#DatabaseStatus.ProtoReflect>)
  - [func \(x \*DatabaseStatus\) Reset\(\)](<#DatabaseStatus.Reset>)
  - [func \(x \*DatabaseStatus\) SetCode\(v common.DatabaseStatusCode\)](<#DatabaseStatus.SetCode>)
  - [func \(x \*DatabaseStatus\) SetMessage\(v string\)](<#DatabaseStatus.SetMessage>)
  - [func \(x \*DatabaseStatus\) String\(\) string](<#DatabaseStatus.String>)
- [type DatabaseStatus\_builder](<#DatabaseStatus_builder>)
  - [func \(b0 DatabaseStatus\_builder\) Build\(\) \*DatabaseStatus](<#DatabaseStatus_builder.Build>)
- [type DecrementRequest](<#DecrementRequest>)
  - [func \(x \*DecrementRequest\) ClearDelta\(\)](<#DecrementRequest.ClearDelta>)
  - [func \(x \*DecrementRequest\) ClearInitialValue\(\)](<#DecrementRequest.ClearInitialValue>)
  - [func \(x \*DecrementRequest\) ClearKey\(\)](<#DecrementRequest.ClearKey>)
  - [func \(x \*DecrementRequest\) ClearMetadata\(\)](<#DecrementRequest.ClearMetadata>)
  - [func \(x \*DecrementRequest\) ClearNamespace\(\)](<#DecrementRequest.ClearNamespace>)
  - [func \(x \*DecrementRequest\) ClearTtl\(\)](<#DecrementRequest.ClearTtl>)
  - [func \(x \*DecrementRequest\) GetDelta\(\) int64](<#DecrementRequest.GetDelta>)
  - [func \(x \*DecrementRequest\) GetInitialValue\(\) int64](<#DecrementRequest.GetInitialValue>)
  - [func \(x \*DecrementRequest\) GetKey\(\) string](<#DecrementRequest.GetKey>)
  - [func \(x \*DecrementRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DecrementRequest.GetMetadata>)
  - [func \(x \*DecrementRequest\) GetNamespace\(\) string](<#DecrementRequest.GetNamespace>)
  - [func \(x \*DecrementRequest\) GetTtl\(\) \*durationpb.Duration](<#DecrementRequest.GetTtl>)
  - [func \(x \*DecrementRequest\) HasDelta\(\) bool](<#DecrementRequest.HasDelta>)
  - [func \(x \*DecrementRequest\) HasInitialValue\(\) bool](<#DecrementRequest.HasInitialValue>)
  - [func \(x \*DecrementRequest\) HasKey\(\) bool](<#DecrementRequest.HasKey>)
  - [func \(x \*DecrementRequest\) HasMetadata\(\) bool](<#DecrementRequest.HasMetadata>)
  - [func \(x \*DecrementRequest\) HasNamespace\(\) bool](<#DecrementRequest.HasNamespace>)
  - [func \(x \*DecrementRequest\) HasTtl\(\) bool](<#DecrementRequest.HasTtl>)
  - [func \(\*DecrementRequest\) ProtoMessage\(\)](<#DecrementRequest.ProtoMessage>)
  - [func \(x \*DecrementRequest\) ProtoReflect\(\) protoreflect.Message](<#DecrementRequest.ProtoReflect>)
  - [func \(x \*DecrementRequest\) Reset\(\)](<#DecrementRequest.Reset>)
  - [func \(x \*DecrementRequest\) SetDelta\(v int64\)](<#DecrementRequest.SetDelta>)
  - [func \(x \*DecrementRequest\) SetInitialValue\(v int64\)](<#DecrementRequest.SetInitialValue>)
  - [func \(x \*DecrementRequest\) SetKey\(v string\)](<#DecrementRequest.SetKey>)
  - [func \(x \*DecrementRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DecrementRequest.SetMetadata>)
  - [func \(x \*DecrementRequest\) SetNamespace\(v string\)](<#DecrementRequest.SetNamespace>)
  - [func \(x \*DecrementRequest\) SetTtl\(v \*durationpb.Duration\)](<#DecrementRequest.SetTtl>)
  - [func \(x \*DecrementRequest\) String\(\) string](<#DecrementRequest.String>)
- [type DecrementRequest\_builder](<#DecrementRequest_builder>)
  - [func \(b0 DecrementRequest\_builder\) Build\(\) \*DecrementRequest](<#DecrementRequest_builder.Build>)
- [type DecrementResponse](<#DecrementResponse>)
  - [func \(x \*DecrementResponse\) ClearError\(\)](<#DecrementResponse.ClearError>)
  - [func \(x \*DecrementResponse\) ClearNewValue\(\)](<#DecrementResponse.ClearNewValue>)
  - [func \(x \*DecrementResponse\) ClearSuccess\(\)](<#DecrementResponse.ClearSuccess>)
  - [func \(x \*DecrementResponse\) GetError\(\) \*common.Error](<#DecrementResponse.GetError>)
  - [func \(x \*DecrementResponse\) GetNewValue\(\) int64](<#DecrementResponse.GetNewValue>)
  - [func \(x \*DecrementResponse\) GetSuccess\(\) bool](<#DecrementResponse.GetSuccess>)
  - [func \(x \*DecrementResponse\) HasError\(\) bool](<#DecrementResponse.HasError>)
  - [func \(x \*DecrementResponse\) HasNewValue\(\) bool](<#DecrementResponse.HasNewValue>)
  - [func \(x \*DecrementResponse\) HasSuccess\(\) bool](<#DecrementResponse.HasSuccess>)
  - [func \(\*DecrementResponse\) ProtoMessage\(\)](<#DecrementResponse.ProtoMessage>)
  - [func \(x \*DecrementResponse\) ProtoReflect\(\) protoreflect.Message](<#DecrementResponse.ProtoReflect>)
  - [func \(x \*DecrementResponse\) Reset\(\)](<#DecrementResponse.Reset>)
  - [func \(x \*DecrementResponse\) SetError\(v \*common.Error\)](<#DecrementResponse.SetError>)
  - [func \(x \*DecrementResponse\) SetNewValue\(v int64\)](<#DecrementResponse.SetNewValue>)
  - [func \(x \*DecrementResponse\) SetSuccess\(v bool\)](<#DecrementResponse.SetSuccess>)
  - [func \(x \*DecrementResponse\) String\(\) string](<#DecrementResponse.String>)
- [type DecrementResponse\_builder](<#DecrementResponse_builder>)
  - [func \(b0 DecrementResponse\_builder\) Build\(\) \*DecrementResponse](<#DecrementResponse_builder.Build>)
- [type DefragRequest](<#DefragRequest>)
  - [func \(x \*DefragRequest\) ClearMetadata\(\)](<#DefragRequest.ClearMetadata>)
  - [func \(x \*DefragRequest\) ClearNamespace\(\)](<#DefragRequest.ClearNamespace>)
  - [func \(x \*DefragRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DefragRequest.GetMetadata>)
  - [func \(x \*DefragRequest\) GetNamespace\(\) string](<#DefragRequest.GetNamespace>)
  - [func \(x \*DefragRequest\) HasMetadata\(\) bool](<#DefragRequest.HasMetadata>)
  - [func \(x \*DefragRequest\) HasNamespace\(\) bool](<#DefragRequest.HasNamespace>)
  - [func \(\*DefragRequest\) ProtoMessage\(\)](<#DefragRequest.ProtoMessage>)
  - [func \(x \*DefragRequest\) ProtoReflect\(\) protoreflect.Message](<#DefragRequest.ProtoReflect>)
  - [func \(x \*DefragRequest\) Reset\(\)](<#DefragRequest.Reset>)
  - [func \(x \*DefragRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DefragRequest.SetMetadata>)
  - [func \(x \*DefragRequest\) SetNamespace\(v string\)](<#DefragRequest.SetNamespace>)
  - [func \(x \*DefragRequest\) String\(\) string](<#DefragRequest.String>)
- [type DefragRequest\_builder](<#DefragRequest_builder>)
  - [func \(b0 DefragRequest\_builder\) Build\(\) \*DefragRequest](<#DefragRequest_builder.Build>)
- [type DeleteMultipleRequest](<#DeleteMultipleRequest>)
  - [func \(x \*DeleteMultipleRequest\) ClearMetadata\(\)](<#DeleteMultipleRequest.ClearMetadata>)
  - [func \(x \*DeleteMultipleRequest\) ClearNamespace\(\)](<#DeleteMultipleRequest.ClearNamespace>)
  - [func \(x \*DeleteMultipleRequest\) GetKeys\(\) \[\]string](<#DeleteMultipleRequest.GetKeys>)
  - [func \(x \*DeleteMultipleRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DeleteMultipleRequest.GetMetadata>)
  - [func \(x \*DeleteMultipleRequest\) GetNamespace\(\) string](<#DeleteMultipleRequest.GetNamespace>)
  - [func \(x \*DeleteMultipleRequest\) HasMetadata\(\) bool](<#DeleteMultipleRequest.HasMetadata>)
  - [func \(x \*DeleteMultipleRequest\) HasNamespace\(\) bool](<#DeleteMultipleRequest.HasNamespace>)
  - [func \(\*DeleteMultipleRequest\) ProtoMessage\(\)](<#DeleteMultipleRequest.ProtoMessage>)
  - [func \(x \*DeleteMultipleRequest\) ProtoReflect\(\) protoreflect.Message](<#DeleteMultipleRequest.ProtoReflect>)
  - [func \(x \*DeleteMultipleRequest\) Reset\(\)](<#DeleteMultipleRequest.Reset>)
  - [func \(x \*DeleteMultipleRequest\) SetKeys\(v \[\]string\)](<#DeleteMultipleRequest.SetKeys>)
  - [func \(x \*DeleteMultipleRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DeleteMultipleRequest.SetMetadata>)
  - [func \(x \*DeleteMultipleRequest\) SetNamespace\(v string\)](<#DeleteMultipleRequest.SetNamespace>)
  - [func \(x \*DeleteMultipleRequest\) String\(\) string](<#DeleteMultipleRequest.String>)
- [type DeleteMultipleRequest\_builder](<#DeleteMultipleRequest_builder>)
  - [func \(b0 DeleteMultipleRequest\_builder\) Build\(\) \*DeleteMultipleRequest](<#DeleteMultipleRequest_builder.Build>)
- [type DeleteMultipleResponse](<#DeleteMultipleResponse>)
  - [func \(x \*DeleteMultipleResponse\) ClearDeletedCount\(\)](<#DeleteMultipleResponse.ClearDeletedCount>)
  - [func \(x \*DeleteMultipleResponse\) ClearError\(\)](<#DeleteMultipleResponse.ClearError>)
  - [func \(x \*DeleteMultipleResponse\) ClearFailedCount\(\)](<#DeleteMultipleResponse.ClearFailedCount>)
  - [func \(x \*DeleteMultipleResponse\) GetDeletedCount\(\) int32](<#DeleteMultipleResponse.GetDeletedCount>)
  - [func \(x \*DeleteMultipleResponse\) GetError\(\) \*common.Error](<#DeleteMultipleResponse.GetError>)
  - [func \(x \*DeleteMultipleResponse\) GetFailedCount\(\) int32](<#DeleteMultipleResponse.GetFailedCount>)
  - [func \(x \*DeleteMultipleResponse\) GetFailedKeys\(\) \[\]string](<#DeleteMultipleResponse.GetFailedKeys>)
  - [func \(x \*DeleteMultipleResponse\) HasDeletedCount\(\) bool](<#DeleteMultipleResponse.HasDeletedCount>)
  - [func \(x \*DeleteMultipleResponse\) HasError\(\) bool](<#DeleteMultipleResponse.HasError>)
  - [func \(x \*DeleteMultipleResponse\) HasFailedCount\(\) bool](<#DeleteMultipleResponse.HasFailedCount>)
  - [func \(\*DeleteMultipleResponse\) ProtoMessage\(\)](<#DeleteMultipleResponse.ProtoMessage>)
  - [func \(x \*DeleteMultipleResponse\) ProtoReflect\(\) protoreflect.Message](<#DeleteMultipleResponse.ProtoReflect>)
  - [func \(x \*DeleteMultipleResponse\) Reset\(\)](<#DeleteMultipleResponse.Reset>)
  - [func \(x \*DeleteMultipleResponse\) SetDeletedCount\(v int32\)](<#DeleteMultipleResponse.SetDeletedCount>)
  - [func \(x \*DeleteMultipleResponse\) SetError\(v \*common.Error\)](<#DeleteMultipleResponse.SetError>)
  - [func \(x \*DeleteMultipleResponse\) SetFailedCount\(v int32\)](<#DeleteMultipleResponse.SetFailedCount>)
  - [func \(x \*DeleteMultipleResponse\) SetFailedKeys\(v \[\]string\)](<#DeleteMultipleResponse.SetFailedKeys>)
  - [func \(x \*DeleteMultipleResponse\) String\(\) string](<#DeleteMultipleResponse.String>)
- [type DeleteMultipleResponse\_builder](<#DeleteMultipleResponse_builder>)
  - [func \(b0 DeleteMultipleResponse\_builder\) Build\(\) \*DeleteMultipleResponse](<#DeleteMultipleResponse_builder.Build>)
- [type DeleteNamespaceRequest](<#DeleteNamespaceRequest>)
  - [func \(x \*DeleteNamespaceRequest\) ClearBackup\(\)](<#DeleteNamespaceRequest.ClearBackup>)
  - [func \(x \*DeleteNamespaceRequest\) ClearForce\(\)](<#DeleteNamespaceRequest.ClearForce>)
  - [func \(x \*DeleteNamespaceRequest\) ClearNamespaceId\(\)](<#DeleteNamespaceRequest.ClearNamespaceId>)
  - [func \(x \*DeleteNamespaceRequest\) GetBackup\(\) bool](<#DeleteNamespaceRequest.GetBackup>)
  - [func \(x \*DeleteNamespaceRequest\) GetForce\(\) bool](<#DeleteNamespaceRequest.GetForce>)
  - [func \(x \*DeleteNamespaceRequest\) GetNamespaceId\(\) string](<#DeleteNamespaceRequest.GetNamespaceId>)
  - [func \(x \*DeleteNamespaceRequest\) HasBackup\(\) bool](<#DeleteNamespaceRequest.HasBackup>)
  - [func \(x \*DeleteNamespaceRequest\) HasForce\(\) bool](<#DeleteNamespaceRequest.HasForce>)
  - [func \(x \*DeleteNamespaceRequest\) HasNamespaceId\(\) bool](<#DeleteNamespaceRequest.HasNamespaceId>)
  - [func \(\*DeleteNamespaceRequest\) ProtoMessage\(\)](<#DeleteNamespaceRequest.ProtoMessage>)
  - [func \(x \*DeleteNamespaceRequest\) ProtoReflect\(\) protoreflect.Message](<#DeleteNamespaceRequest.ProtoReflect>)
  - [func \(x \*DeleteNamespaceRequest\) Reset\(\)](<#DeleteNamespaceRequest.Reset>)
  - [func \(x \*DeleteNamespaceRequest\) SetBackup\(v bool\)](<#DeleteNamespaceRequest.SetBackup>)
  - [func \(x \*DeleteNamespaceRequest\) SetForce\(v bool\)](<#DeleteNamespaceRequest.SetForce>)
  - [func \(x \*DeleteNamespaceRequest\) SetNamespaceId\(v string\)](<#DeleteNamespaceRequest.SetNamespaceId>)
  - [func \(x \*DeleteNamespaceRequest\) String\(\) string](<#DeleteNamespaceRequest.String>)
- [type DeleteNamespaceRequest\_builder](<#DeleteNamespaceRequest_builder>)
  - [func \(b0 DeleteNamespaceRequest\_builder\) Build\(\) \*DeleteNamespaceRequest](<#DeleteNamespaceRequest_builder.Build>)
- [type DropDatabaseRequest](<#DropDatabaseRequest>)
  - [func \(x \*DropDatabaseRequest\) ClearMetadata\(\)](<#DropDatabaseRequest.ClearMetadata>)
  - [func \(x \*DropDatabaseRequest\) ClearName\(\)](<#DropDatabaseRequest.ClearName>)
  - [func \(x \*DropDatabaseRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DropDatabaseRequest.GetMetadata>)
  - [func \(x \*DropDatabaseRequest\) GetName\(\) string](<#DropDatabaseRequest.GetName>)
  - [func \(x \*DropDatabaseRequest\) HasMetadata\(\) bool](<#DropDatabaseRequest.HasMetadata>)
  - [func \(x \*DropDatabaseRequest\) HasName\(\) bool](<#DropDatabaseRequest.HasName>)
  - [func \(\*DropDatabaseRequest\) ProtoMessage\(\)](<#DropDatabaseRequest.ProtoMessage>)
  - [func \(x \*DropDatabaseRequest\) ProtoReflect\(\) protoreflect.Message](<#DropDatabaseRequest.ProtoReflect>)
  - [func \(x \*DropDatabaseRequest\) Reset\(\)](<#DropDatabaseRequest.Reset>)
  - [func \(x \*DropDatabaseRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DropDatabaseRequest.SetMetadata>)
  - [func \(x \*DropDatabaseRequest\) SetName\(v string\)](<#DropDatabaseRequest.SetName>)
  - [func \(x \*DropDatabaseRequest\) String\(\) string](<#DropDatabaseRequest.String>)
- [type DropDatabaseRequest\_builder](<#DropDatabaseRequest_builder>)
  - [func \(b0 DropDatabaseRequest\_builder\) Build\(\) \*DropDatabaseRequest](<#DropDatabaseRequest_builder.Build>)
- [type DropSchemaRequest](<#DropSchemaRequest>)
  - [func \(x \*DropSchemaRequest\) ClearDatabase\(\)](<#DropSchemaRequest.ClearDatabase>)
  - [func \(x \*DropSchemaRequest\) ClearMetadata\(\)](<#DropSchemaRequest.ClearMetadata>)
  - [func \(x \*DropSchemaRequest\) ClearSchema\(\)](<#DropSchemaRequest.ClearSchema>)
  - [func \(x \*DropSchemaRequest\) GetDatabase\(\) string](<#DropSchemaRequest.GetDatabase>)
  - [func \(x \*DropSchemaRequest\) GetMetadata\(\) \*common.RequestMetadata](<#DropSchemaRequest.GetMetadata>)
  - [func \(x \*DropSchemaRequest\) GetSchema\(\) string](<#DropSchemaRequest.GetSchema>)
  - [func \(x \*DropSchemaRequest\) HasDatabase\(\) bool](<#DropSchemaRequest.HasDatabase>)
  - [func \(x \*DropSchemaRequest\) HasMetadata\(\) bool](<#DropSchemaRequest.HasMetadata>)
  - [func \(x \*DropSchemaRequest\) HasSchema\(\) bool](<#DropSchemaRequest.HasSchema>)
  - [func \(\*DropSchemaRequest\) ProtoMessage\(\)](<#DropSchemaRequest.ProtoMessage>)
  - [func \(x \*DropSchemaRequest\) ProtoReflect\(\) protoreflect.Message](<#DropSchemaRequest.ProtoReflect>)
  - [func \(x \*DropSchemaRequest\) Reset\(\)](<#DropSchemaRequest.Reset>)
  - [func \(x \*DropSchemaRequest\) SetDatabase\(v string\)](<#DropSchemaRequest.SetDatabase>)
  - [func \(x \*DropSchemaRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#DropSchemaRequest.SetMetadata>)
  - [func \(x \*DropSchemaRequest\) SetSchema\(v string\)](<#DropSchemaRequest.SetSchema>)
  - [func \(x \*DropSchemaRequest\) String\(\) string](<#DropSchemaRequest.String>)
- [type DropSchemaRequest\_builder](<#DropSchemaRequest_builder>)
  - [func \(b0 DropSchemaRequest\_builder\) Build\(\) \*DropSchemaRequest](<#DropSchemaRequest_builder.Build>)
- [type EvictionResult](<#EvictionResult>)
  - [func \(x \*EvictionResult\) ClearEvictedAt\(\)](<#EvictionResult.ClearEvictedAt>)
  - [func \(x \*EvictionResult\) ClearEvictedCount\(\)](<#EvictionResult.ClearEvictedCount>)
  - [func \(x \*EvictionResult\) ClearEvictionReason\(\)](<#EvictionResult.ClearEvictionReason>)
  - [func \(x \*EvictionResult\) ClearMemoryFreed\(\)](<#EvictionResult.ClearMemoryFreed>)
  - [func \(x \*EvictionResult\) ClearPolicyUsed\(\)](<#EvictionResult.ClearPolicyUsed>)
  - [func \(x \*EvictionResult\) ClearSuccess\(\)](<#EvictionResult.ClearSuccess>)
  - [func \(x \*EvictionResult\) GetEvictedAt\(\) \*timestamppb.Timestamp](<#EvictionResult.GetEvictedAt>)
  - [func \(x \*EvictionResult\) GetEvictedCount\(\) int64](<#EvictionResult.GetEvictedCount>)
  - [func \(x \*EvictionResult\) GetEvictedKeys\(\) \[\]string](<#EvictionResult.GetEvictedKeys>)
  - [func \(x \*EvictionResult\) GetEvictionReason\(\) string](<#EvictionResult.GetEvictionReason>)
  - [func \(x \*EvictionResult\) GetMemoryFreed\(\) int64](<#EvictionResult.GetMemoryFreed>)
  - [func \(x \*EvictionResult\) GetPolicyUsed\(\) common.EvictionPolicy](<#EvictionResult.GetPolicyUsed>)
  - [func \(x \*EvictionResult\) GetSuccess\(\) bool](<#EvictionResult.GetSuccess>)
  - [func \(x \*EvictionResult\) HasEvictedAt\(\) bool](<#EvictionResult.HasEvictedAt>)
  - [func \(x \*EvictionResult\) HasEvictedCount\(\) bool](<#EvictionResult.HasEvictedCount>)
  - [func \(x \*EvictionResult\) HasEvictionReason\(\) bool](<#EvictionResult.HasEvictionReason>)
  - [func \(x \*EvictionResult\) HasMemoryFreed\(\) bool](<#EvictionResult.HasMemoryFreed>)
  - [func \(x \*EvictionResult\) HasPolicyUsed\(\) bool](<#EvictionResult.HasPolicyUsed>)
  - [func \(x \*EvictionResult\) HasSuccess\(\) bool](<#EvictionResult.HasSuccess>)
  - [func \(\*EvictionResult\) ProtoMessage\(\)](<#EvictionResult.ProtoMessage>)
  - [func \(x \*EvictionResult\) ProtoReflect\(\) protoreflect.Message](<#EvictionResult.ProtoReflect>)
  - [func \(x \*EvictionResult\) Reset\(\)](<#EvictionResult.Reset>)
  - [func \(x \*EvictionResult\) SetEvictedAt\(v \*timestamppb.Timestamp\)](<#EvictionResult.SetEvictedAt>)
  - [func \(x \*EvictionResult\) SetEvictedCount\(v int64\)](<#EvictionResult.SetEvictedCount>)
  - [func \(x \*EvictionResult\) SetEvictedKeys\(v \[\]string\)](<#EvictionResult.SetEvictedKeys>)
  - [func \(x \*EvictionResult\) SetEvictionReason\(v string\)](<#EvictionResult.SetEvictionReason>)
  - [func \(x \*EvictionResult\) SetMemoryFreed\(v int64\)](<#EvictionResult.SetMemoryFreed>)
  - [func \(x \*EvictionResult\) SetPolicyUsed\(v common.EvictionPolicy\)](<#EvictionResult.SetPolicyUsed>)
  - [func \(x \*EvictionResult\) SetSuccess\(v bool\)](<#EvictionResult.SetSuccess>)
  - [func \(x \*EvictionResult\) String\(\) string](<#EvictionResult.String>)
- [type EvictionResult\_builder](<#EvictionResult_builder>)
  - [func \(b0 EvictionResult\_builder\) Build\(\) \*EvictionResult](<#EvictionResult_builder.Build>)
- [type ExecuteBatchRequest](<#ExecuteBatchRequest>)
  - [func \(x \*ExecuteBatchRequest\) ClearDatabase\(\)](<#ExecuteBatchRequest.ClearDatabase>)
  - [func \(x \*ExecuteBatchRequest\) ClearMetadata\(\)](<#ExecuteBatchRequest.ClearMetadata>)
  - [func \(x \*ExecuteBatchRequest\) ClearOptions\(\)](<#ExecuteBatchRequest.ClearOptions>)
  - [func \(x \*ExecuteBatchRequest\) ClearTransactionId\(\)](<#ExecuteBatchRequest.ClearTransactionId>)
  - [func \(x \*ExecuteBatchRequest\) GetDatabase\(\) string](<#ExecuteBatchRequest.GetDatabase>)
  - [func \(x \*ExecuteBatchRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ExecuteBatchRequest.GetMetadata>)
  - [func \(x \*ExecuteBatchRequest\) GetOperations\(\) \[\]\*DatabaseBatchOperation](<#ExecuteBatchRequest.GetOperations>)
  - [func \(x \*ExecuteBatchRequest\) GetOptions\(\) \*BatchExecuteOptions](<#ExecuteBatchRequest.GetOptions>)
  - [func \(x \*ExecuteBatchRequest\) GetTransactionId\(\) string](<#ExecuteBatchRequest.GetTransactionId>)
  - [func \(x \*ExecuteBatchRequest\) HasDatabase\(\) bool](<#ExecuteBatchRequest.HasDatabase>)
  - [func \(x \*ExecuteBatchRequest\) HasMetadata\(\) bool](<#ExecuteBatchRequest.HasMetadata>)
  - [func \(x \*ExecuteBatchRequest\) HasOptions\(\) bool](<#ExecuteBatchRequest.HasOptions>)
  - [func \(x \*ExecuteBatchRequest\) HasTransactionId\(\) bool](<#ExecuteBatchRequest.HasTransactionId>)
  - [func \(\*ExecuteBatchRequest\) ProtoMessage\(\)](<#ExecuteBatchRequest.ProtoMessage>)
  - [func \(x \*ExecuteBatchRequest\) ProtoReflect\(\) protoreflect.Message](<#ExecuteBatchRequest.ProtoReflect>)
  - [func \(x \*ExecuteBatchRequest\) Reset\(\)](<#ExecuteBatchRequest.Reset>)
  - [func \(x \*ExecuteBatchRequest\) SetDatabase\(v string\)](<#ExecuteBatchRequest.SetDatabase>)
  - [func \(x \*ExecuteBatchRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ExecuteBatchRequest.SetMetadata>)
  - [func \(x \*ExecuteBatchRequest\) SetOperations\(v \[\]\*DatabaseBatchOperation\)](<#ExecuteBatchRequest.SetOperations>)
  - [func \(x \*ExecuteBatchRequest\) SetOptions\(v \*BatchExecuteOptions\)](<#ExecuteBatchRequest.SetOptions>)
  - [func \(x \*ExecuteBatchRequest\) SetTransactionId\(v string\)](<#ExecuteBatchRequest.SetTransactionId>)
  - [func \(x \*ExecuteBatchRequest\) String\(\) string](<#ExecuteBatchRequest.String>)
- [type ExecuteBatchRequest\_builder](<#ExecuteBatchRequest_builder>)
  - [func \(b0 ExecuteBatchRequest\_builder\) Build\(\) \*ExecuteBatchRequest](<#ExecuteBatchRequest_builder.Build>)
- [type ExecuteBatchResponse](<#ExecuteBatchResponse>)
  - [func \(x \*ExecuteBatchResponse\) ClearError\(\)](<#ExecuteBatchResponse.ClearError>)
  - [func \(x \*ExecuteBatchResponse\) ClearStats\(\)](<#ExecuteBatchResponse.ClearStats>)
  - [func \(x \*ExecuteBatchResponse\) GetError\(\) \*common.Error](<#ExecuteBatchResponse.GetError>)
  - [func \(x \*ExecuteBatchResponse\) GetResults\(\) \[\]\*BatchOperationResult](<#ExecuteBatchResponse.GetResults>)
  - [func \(x \*ExecuteBatchResponse\) GetStats\(\) \*DatabaseBatchStats](<#ExecuteBatchResponse.GetStats>)
  - [func \(x \*ExecuteBatchResponse\) HasError\(\) bool](<#ExecuteBatchResponse.HasError>)
  - [func \(x \*ExecuteBatchResponse\) HasStats\(\) bool](<#ExecuteBatchResponse.HasStats>)
  - [func \(\*ExecuteBatchResponse\) ProtoMessage\(\)](<#ExecuteBatchResponse.ProtoMessage>)
  - [func \(x \*ExecuteBatchResponse\) ProtoReflect\(\) protoreflect.Message](<#ExecuteBatchResponse.ProtoReflect>)
  - [func \(x \*ExecuteBatchResponse\) Reset\(\)](<#ExecuteBatchResponse.Reset>)
  - [func \(x \*ExecuteBatchResponse\) SetError\(v \*common.Error\)](<#ExecuteBatchResponse.SetError>)
  - [func \(x \*ExecuteBatchResponse\) SetResults\(v \[\]\*BatchOperationResult\)](<#ExecuteBatchResponse.SetResults>)
  - [func \(x \*ExecuteBatchResponse\) SetStats\(v \*DatabaseBatchStats\)](<#ExecuteBatchResponse.SetStats>)
  - [func \(x \*ExecuteBatchResponse\) String\(\) string](<#ExecuteBatchResponse.String>)
- [type ExecuteBatchResponse\_builder](<#ExecuteBatchResponse_builder>)
  - [func \(b0 ExecuteBatchResponse\_builder\) Build\(\) \*ExecuteBatchResponse](<#ExecuteBatchResponse_builder.Build>)
- [type ExecuteOptions](<#ExecuteOptions>)
  - [func \(x \*ExecuteOptions\) ClearIsolation\(\)](<#ExecuteOptions.ClearIsolation>)
  - [func \(x \*ExecuteOptions\) ClearReturnGeneratedKeys\(\)](<#ExecuteOptions.ClearReturnGeneratedKeys>)
  - [func \(x \*ExecuteOptions\) ClearTimeout\(\)](<#ExecuteOptions.ClearTimeout>)
  - [func \(x \*ExecuteOptions\) GetIsolation\(\) common.DatabaseIsolationLevel](<#ExecuteOptions.GetIsolation>)
  - [func \(x \*ExecuteOptions\) GetReturnGeneratedKeys\(\) bool](<#ExecuteOptions.GetReturnGeneratedKeys>)
  - [func \(x \*ExecuteOptions\) GetTimeout\(\) \*durationpb.Duration](<#ExecuteOptions.GetTimeout>)
  - [func \(x \*ExecuteOptions\) HasIsolation\(\) bool](<#ExecuteOptions.HasIsolation>)
  - [func \(x \*ExecuteOptions\) HasReturnGeneratedKeys\(\) bool](<#ExecuteOptions.HasReturnGeneratedKeys>)
  - [func \(x \*ExecuteOptions\) HasTimeout\(\) bool](<#ExecuteOptions.HasTimeout>)
  - [func \(\*ExecuteOptions\) ProtoMessage\(\)](<#ExecuteOptions.ProtoMessage>)
  - [func \(x \*ExecuteOptions\) ProtoReflect\(\) protoreflect.Message](<#ExecuteOptions.ProtoReflect>)
  - [func \(x \*ExecuteOptions\) Reset\(\)](<#ExecuteOptions.Reset>)
  - [func \(x \*ExecuteOptions\) SetIsolation\(v common.DatabaseIsolationLevel\)](<#ExecuteOptions.SetIsolation>)
  - [func \(x \*ExecuteOptions\) SetReturnGeneratedKeys\(v bool\)](<#ExecuteOptions.SetReturnGeneratedKeys>)
  - [func \(x \*ExecuteOptions\) SetTimeout\(v \*durationpb.Duration\)](<#ExecuteOptions.SetTimeout>)
  - [func \(x \*ExecuteOptions\) String\(\) string](<#ExecuteOptions.String>)
- [type ExecuteOptions\_builder](<#ExecuteOptions_builder>)
  - [func \(b0 ExecuteOptions\_builder\) Build\(\) \*ExecuteOptions](<#ExecuteOptions_builder.Build>)
- [type ExecuteRequest](<#ExecuteRequest>)
  - [func \(x \*ExecuteRequest\) ClearDatabase\(\)](<#ExecuteRequest.ClearDatabase>)
  - [func \(x \*ExecuteRequest\) ClearMetadata\(\)](<#ExecuteRequest.ClearMetadata>)
  - [func \(x \*ExecuteRequest\) ClearOptions\(\)](<#ExecuteRequest.ClearOptions>)
  - [func \(x \*ExecuteRequest\) ClearStatement\(\)](<#ExecuteRequest.ClearStatement>)
  - [func \(x \*ExecuteRequest\) ClearTransactionId\(\)](<#ExecuteRequest.ClearTransactionId>)
  - [func \(x \*ExecuteRequest\) GetDatabase\(\) string](<#ExecuteRequest.GetDatabase>)
  - [func \(x \*ExecuteRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ExecuteRequest.GetMetadata>)
  - [func \(x \*ExecuteRequest\) GetOptions\(\) \*ExecuteOptions](<#ExecuteRequest.GetOptions>)
  - [func \(x \*ExecuteRequest\) GetParameters\(\) \[\]\*QueryParameter](<#ExecuteRequest.GetParameters>)
  - [func \(x \*ExecuteRequest\) GetStatement\(\) string](<#ExecuteRequest.GetStatement>)
  - [func \(x \*ExecuteRequest\) GetTransactionId\(\) string](<#ExecuteRequest.GetTransactionId>)
  - [func \(x \*ExecuteRequest\) HasDatabase\(\) bool](<#ExecuteRequest.HasDatabase>)
  - [func \(x \*ExecuteRequest\) HasMetadata\(\) bool](<#ExecuteRequest.HasMetadata>)
  - [func \(x \*ExecuteRequest\) HasOptions\(\) bool](<#ExecuteRequest.HasOptions>)
  - [func \(x \*ExecuteRequest\) HasStatement\(\) bool](<#ExecuteRequest.HasStatement>)
  - [func \(x \*ExecuteRequest\) HasTransactionId\(\) bool](<#ExecuteRequest.HasTransactionId>)
  - [func \(\*ExecuteRequest\) ProtoMessage\(\)](<#ExecuteRequest.ProtoMessage>)
  - [func \(x \*ExecuteRequest\) ProtoReflect\(\) protoreflect.Message](<#ExecuteRequest.ProtoReflect>)
  - [func \(x \*ExecuteRequest\) Reset\(\)](<#ExecuteRequest.Reset>)
  - [func \(x \*ExecuteRequest\) SetDatabase\(v string\)](<#ExecuteRequest.SetDatabase>)
  - [func \(x \*ExecuteRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ExecuteRequest.SetMetadata>)
  - [func \(x \*ExecuteRequest\) SetOptions\(v \*ExecuteOptions\)](<#ExecuteRequest.SetOptions>)
  - [func \(x \*ExecuteRequest\) SetParameters\(v \[\]\*QueryParameter\)](<#ExecuteRequest.SetParameters>)
  - [func \(x \*ExecuteRequest\) SetStatement\(v string\)](<#ExecuteRequest.SetStatement>)
  - [func \(x \*ExecuteRequest\) SetTransactionId\(v string\)](<#ExecuteRequest.SetTransactionId>)
  - [func \(x \*ExecuteRequest\) String\(\) string](<#ExecuteRequest.String>)
- [type ExecuteRequest\_builder](<#ExecuteRequest_builder>)
  - [func \(b0 ExecuteRequest\_builder\) Build\(\) \*ExecuteRequest](<#ExecuteRequest_builder.Build>)
- [type ExecuteResponse](<#ExecuteResponse>)
  - [func \(x \*ExecuteResponse\) ClearAffectedRows\(\)](<#ExecuteResponse.ClearAffectedRows>)
  - [func \(x \*ExecuteResponse\) ClearError\(\)](<#ExecuteResponse.ClearError>)
  - [func \(x \*ExecuteResponse\) ClearStats\(\)](<#ExecuteResponse.ClearStats>)
  - [func \(x \*ExecuteResponse\) GetAffectedRows\(\) int64](<#ExecuteResponse.GetAffectedRows>)
  - [func \(x \*ExecuteResponse\) GetError\(\) \*common.Error](<#ExecuteResponse.GetError>)
  - [func \(x \*ExecuteResponse\) GetGeneratedKeys\(\) \[\]\*anypb.Any](<#ExecuteResponse.GetGeneratedKeys>)
  - [func \(x \*ExecuteResponse\) GetStats\(\) \*ExecuteStats](<#ExecuteResponse.GetStats>)
  - [func \(x \*ExecuteResponse\) HasAffectedRows\(\) bool](<#ExecuteResponse.HasAffectedRows>)
  - [func \(x \*ExecuteResponse\) HasError\(\) bool](<#ExecuteResponse.HasError>)
  - [func \(x \*ExecuteResponse\) HasStats\(\) bool](<#ExecuteResponse.HasStats>)
  - [func \(\*ExecuteResponse\) ProtoMessage\(\)](<#ExecuteResponse.ProtoMessage>)
  - [func \(x \*ExecuteResponse\) ProtoReflect\(\) protoreflect.Message](<#ExecuteResponse.ProtoReflect>)
  - [func \(x \*ExecuteResponse\) Reset\(\)](<#ExecuteResponse.Reset>)
  - [func \(x \*ExecuteResponse\) SetAffectedRows\(v int64\)](<#ExecuteResponse.SetAffectedRows>)
  - [func \(x \*ExecuteResponse\) SetError\(v \*common.Error\)](<#ExecuteResponse.SetError>)
  - [func \(x \*ExecuteResponse\) SetGeneratedKeys\(v \[\]\*anypb.Any\)](<#ExecuteResponse.SetGeneratedKeys>)
  - [func \(x \*ExecuteResponse\) SetStats\(v \*ExecuteStats\)](<#ExecuteResponse.SetStats>)
  - [func \(x \*ExecuteResponse\) String\(\) string](<#ExecuteResponse.String>)
- [type ExecuteResponse\_builder](<#ExecuteResponse_builder>)
  - [func \(b0 ExecuteResponse\_builder\) Build\(\) \*ExecuteResponse](<#ExecuteResponse_builder.Build>)
- [type ExecuteStats](<#ExecuteStats>)
  - [func \(x \*ExecuteStats\) ClearAffectedRows\(\)](<#ExecuteStats.ClearAffectedRows>)
  - [func \(x \*ExecuteStats\) ClearCostEstimate\(\)](<#ExecuteStats.ClearCostEstimate>)
  - [func \(x \*ExecuteStats\) ClearExecutionTime\(\)](<#ExecuteStats.ClearExecutionTime>)
  - [func \(x \*ExecuteStats\) GetAffectedRows\(\) int64](<#ExecuteStats.GetAffectedRows>)
  - [func \(x \*ExecuteStats\) GetCostEstimate\(\) float64](<#ExecuteStats.GetCostEstimate>)
  - [func \(x \*ExecuteStats\) GetExecutionTime\(\) \*durationpb.Duration](<#ExecuteStats.GetExecutionTime>)
  - [func \(x \*ExecuteStats\) HasAffectedRows\(\) bool](<#ExecuteStats.HasAffectedRows>)
  - [func \(x \*ExecuteStats\) HasCostEstimate\(\) bool](<#ExecuteStats.HasCostEstimate>)
  - [func \(x \*ExecuteStats\) HasExecutionTime\(\) bool](<#ExecuteStats.HasExecutionTime>)
  - [func \(\*ExecuteStats\) ProtoMessage\(\)](<#ExecuteStats.ProtoMessage>)
  - [func \(x \*ExecuteStats\) ProtoReflect\(\) protoreflect.Message](<#ExecuteStats.ProtoReflect>)
  - [func \(x \*ExecuteStats\) Reset\(\)](<#ExecuteStats.Reset>)
  - [func \(x \*ExecuteStats\) SetAffectedRows\(v int64\)](<#ExecuteStats.SetAffectedRows>)
  - [func \(x \*ExecuteStats\) SetCostEstimate\(v float64\)](<#ExecuteStats.SetCostEstimate>)
  - [func \(x \*ExecuteStats\) SetExecutionTime\(v \*durationpb.Duration\)](<#ExecuteStats.SetExecutionTime>)
  - [func \(x \*ExecuteStats\) String\(\) string](<#ExecuteStats.String>)
- [type ExecuteStats\_builder](<#ExecuteStats_builder>)
  - [func \(b0 ExecuteStats\_builder\) Build\(\) \*ExecuteStats](<#ExecuteStats_builder.Build>)
- [type ExistsRequest](<#ExistsRequest>)
  - [func \(x \*ExistsRequest\) ClearKey\(\)](<#ExistsRequest.ClearKey>)
  - [func \(x \*ExistsRequest\) ClearMetadata\(\)](<#ExistsRequest.ClearMetadata>)
  - [func \(x \*ExistsRequest\) ClearNamespace\(\)](<#ExistsRequest.ClearNamespace>)
  - [func \(x \*ExistsRequest\) GetKey\(\) string](<#ExistsRequest.GetKey>)
  - [func \(x \*ExistsRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ExistsRequest.GetMetadata>)
  - [func \(x \*ExistsRequest\) GetNamespace\(\) string](<#ExistsRequest.GetNamespace>)
  - [func \(x \*ExistsRequest\) HasKey\(\) bool](<#ExistsRequest.HasKey>)
  - [func \(x \*ExistsRequest\) HasMetadata\(\) bool](<#ExistsRequest.HasMetadata>)
  - [func \(x \*ExistsRequest\) HasNamespace\(\) bool](<#ExistsRequest.HasNamespace>)
  - [func \(\*ExistsRequest\) ProtoMessage\(\)](<#ExistsRequest.ProtoMessage>)
  - [func \(x \*ExistsRequest\) ProtoReflect\(\) protoreflect.Message](<#ExistsRequest.ProtoReflect>)
  - [func \(x \*ExistsRequest\) Reset\(\)](<#ExistsRequest.Reset>)
  - [func \(x \*ExistsRequest\) SetKey\(v string\)](<#ExistsRequest.SetKey>)
  - [func \(x \*ExistsRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ExistsRequest.SetMetadata>)
  - [func \(x \*ExistsRequest\) SetNamespace\(v string\)](<#ExistsRequest.SetNamespace>)
  - [func \(x \*ExistsRequest\) String\(\) string](<#ExistsRequest.String>)
- [type ExistsRequest\_builder](<#ExistsRequest_builder>)
  - [func \(b0 ExistsRequest\_builder\) Build\(\) \*ExistsRequest](<#ExistsRequest_builder.Build>)
- [type ExistsResponse](<#ExistsResponse>)
  - [func \(x \*ExistsResponse\) ClearError\(\)](<#ExistsResponse.ClearError>)
  - [func \(x \*ExistsResponse\) ClearExists\(\)](<#ExistsResponse.ClearExists>)
  - [func \(x \*ExistsResponse\) GetError\(\) \*common.Error](<#ExistsResponse.GetError>)
  - [func \(x \*ExistsResponse\) GetExists\(\) bool](<#ExistsResponse.GetExists>)
  - [func \(x \*ExistsResponse\) HasError\(\) bool](<#ExistsResponse.HasError>)
  - [func \(x \*ExistsResponse\) HasExists\(\) bool](<#ExistsResponse.HasExists>)
  - [func \(\*ExistsResponse\) ProtoMessage\(\)](<#ExistsResponse.ProtoMessage>)
  - [func \(x \*ExistsResponse\) ProtoReflect\(\) protoreflect.Message](<#ExistsResponse.ProtoReflect>)
  - [func \(x \*ExistsResponse\) Reset\(\)](<#ExistsResponse.Reset>)
  - [func \(x \*ExistsResponse\) SetError\(v \*common.Error\)](<#ExistsResponse.SetError>)
  - [func \(x \*ExistsResponse\) SetExists\(v bool\)](<#ExistsResponse.SetExists>)
  - [func \(x \*ExistsResponse\) String\(\) string](<#ExistsResponse.String>)
- [type ExistsResponse\_builder](<#ExistsResponse_builder>)
  - [func \(b0 ExistsResponse\_builder\) Build\(\) \*ExistsResponse](<#ExistsResponse_builder.Build>)
- [type ExpireRequest](<#ExpireRequest>)
  - [func \(x \*ExpireRequest\) ClearKey\(\)](<#ExpireRequest.ClearKey>)
  - [func \(x \*ExpireRequest\) ClearMetadata\(\)](<#ExpireRequest.ClearMetadata>)
  - [func \(x \*ExpireRequest\) ClearNamespace\(\)](<#ExpireRequest.ClearNamespace>)
  - [func \(x \*ExpireRequest\) ClearTtl\(\)](<#ExpireRequest.ClearTtl>)
  - [func \(x \*ExpireRequest\) GetKey\(\) string](<#ExpireRequest.GetKey>)
  - [func \(x \*ExpireRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ExpireRequest.GetMetadata>)
  - [func \(x \*ExpireRequest\) GetNamespace\(\) string](<#ExpireRequest.GetNamespace>)
  - [func \(x \*ExpireRequest\) GetTtl\(\) \*durationpb.Duration](<#ExpireRequest.GetTtl>)
  - [func \(x \*ExpireRequest\) HasKey\(\) bool](<#ExpireRequest.HasKey>)
  - [func \(x \*ExpireRequest\) HasMetadata\(\) bool](<#ExpireRequest.HasMetadata>)
  - [func \(x \*ExpireRequest\) HasNamespace\(\) bool](<#ExpireRequest.HasNamespace>)
  - [func \(x \*ExpireRequest\) HasTtl\(\) bool](<#ExpireRequest.HasTtl>)
  - [func \(\*ExpireRequest\) ProtoMessage\(\)](<#ExpireRequest.ProtoMessage>)
  - [func \(x \*ExpireRequest\) ProtoReflect\(\) protoreflect.Message](<#ExpireRequest.ProtoReflect>)
  - [func \(x \*ExpireRequest\) Reset\(\)](<#ExpireRequest.Reset>)
  - [func \(x \*ExpireRequest\) SetKey\(v string\)](<#ExpireRequest.SetKey>)
  - [func \(x \*ExpireRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ExpireRequest.SetMetadata>)
  - [func \(x \*ExpireRequest\) SetNamespace\(v string\)](<#ExpireRequest.SetNamespace>)
  - [func \(x \*ExpireRequest\) SetTtl\(v \*durationpb.Duration\)](<#ExpireRequest.SetTtl>)
  - [func \(x \*ExpireRequest\) String\(\) string](<#ExpireRequest.String>)
- [type ExpireRequest\_builder](<#ExpireRequest_builder>)
  - [func \(b0 ExpireRequest\_builder\) Build\(\) \*ExpireRequest](<#ExpireRequest_builder.Build>)
- [type ExportRequest](<#ExportRequest>)
  - [func \(x \*ExportRequest\) ClearDestination\(\)](<#ExportRequest.ClearDestination>)
  - [func \(x \*ExportRequest\) ClearMetadata\(\)](<#ExportRequest.ClearMetadata>)
  - [func \(x \*ExportRequest\) ClearNamespace\(\)](<#ExportRequest.ClearNamespace>)
  - [func \(x \*ExportRequest\) GetDestination\(\) string](<#ExportRequest.GetDestination>)
  - [func \(x \*ExportRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ExportRequest.GetMetadata>)
  - [func \(x \*ExportRequest\) GetNamespace\(\) string](<#ExportRequest.GetNamespace>)
  - [func \(x \*ExportRequest\) HasDestination\(\) bool](<#ExportRequest.HasDestination>)
  - [func \(x \*ExportRequest\) HasMetadata\(\) bool](<#ExportRequest.HasMetadata>)
  - [func \(x \*ExportRequest\) HasNamespace\(\) bool](<#ExportRequest.HasNamespace>)
  - [func \(\*ExportRequest\) ProtoMessage\(\)](<#ExportRequest.ProtoMessage>)
  - [func \(x \*ExportRequest\) ProtoReflect\(\) protoreflect.Message](<#ExportRequest.ProtoReflect>)
  - [func \(x \*ExportRequest\) Reset\(\)](<#ExportRequest.Reset>)
  - [func \(x \*ExportRequest\) SetDestination\(v string\)](<#ExportRequest.SetDestination>)
  - [func \(x \*ExportRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ExportRequest.SetMetadata>)
  - [func \(x \*ExportRequest\) SetNamespace\(v string\)](<#ExportRequest.SetNamespace>)
  - [func \(x \*ExportRequest\) String\(\) string](<#ExportRequest.String>)
- [type ExportRequest\_builder](<#ExportRequest_builder>)
  - [func \(b0 ExportRequest\_builder\) Build\(\) \*ExportRequest](<#ExportRequest_builder.Build>)
- [type FlushRequest](<#FlushRequest>)
  - [func \(x \*FlushRequest\) ClearMetadata\(\)](<#FlushRequest.ClearMetadata>)
  - [func \(x \*FlushRequest\) ClearNamespace\(\)](<#FlushRequest.ClearNamespace>)
  - [func \(x \*FlushRequest\) GetMetadata\(\) \*common.RequestMetadata](<#FlushRequest.GetMetadata>)
  - [func \(x \*FlushRequest\) GetNamespace\(\) string](<#FlushRequest.GetNamespace>)
  - [func \(x \*FlushRequest\) HasMetadata\(\) bool](<#FlushRequest.HasMetadata>)
  - [func \(x \*FlushRequest\) HasNamespace\(\) bool](<#FlushRequest.HasNamespace>)
  - [func \(\*FlushRequest\) ProtoMessage\(\)](<#FlushRequest.ProtoMessage>)
  - [func \(x \*FlushRequest\) ProtoReflect\(\) protoreflect.Message](<#FlushRequest.ProtoReflect>)
  - [func \(x \*FlushRequest\) Reset\(\)](<#FlushRequest.Reset>)
  - [func \(x \*FlushRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#FlushRequest.SetMetadata>)
  - [func \(x \*FlushRequest\) SetNamespace\(v string\)](<#FlushRequest.SetNamespace>)
  - [func \(x \*FlushRequest\) String\(\) string](<#FlushRequest.String>)
- [type FlushRequest\_builder](<#FlushRequest_builder>)
  - [func \(b0 FlushRequest\_builder\) Build\(\) \*FlushRequest](<#FlushRequest_builder.Build>)
- [type FlushResponse](<#FlushResponse>)
  - [func \(x \*FlushResponse\) ClearError\(\)](<#FlushResponse.ClearError>)
  - [func \(x \*FlushResponse\) ClearFlushedCount\(\)](<#FlushResponse.ClearFlushedCount>)
  - [func \(x \*FlushResponse\) ClearSuccess\(\)](<#FlushResponse.ClearSuccess>)
  - [func \(x \*FlushResponse\) GetError\(\) \*common.Error](<#FlushResponse.GetError>)
  - [func \(x \*FlushResponse\) GetFlushedCount\(\) int64](<#FlushResponse.GetFlushedCount>)
  - [func \(x \*FlushResponse\) GetSuccess\(\) bool](<#FlushResponse.GetSuccess>)
  - [func \(x \*FlushResponse\) HasError\(\) bool](<#FlushResponse.HasError>)
  - [func \(x \*FlushResponse\) HasFlushedCount\(\) bool](<#FlushResponse.HasFlushedCount>)
  - [func \(x \*FlushResponse\) HasSuccess\(\) bool](<#FlushResponse.HasSuccess>)
  - [func \(\*FlushResponse\) ProtoMessage\(\)](<#FlushResponse.ProtoMessage>)
  - [func \(x \*FlushResponse\) ProtoReflect\(\) protoreflect.Message](<#FlushResponse.ProtoReflect>)
  - [func \(x \*FlushResponse\) Reset\(\)](<#FlushResponse.Reset>)
  - [func \(x \*FlushResponse\) SetError\(v \*common.Error\)](<#FlushResponse.SetError>)
  - [func \(x \*FlushResponse\) SetFlushedCount\(v int64\)](<#FlushResponse.SetFlushedCount>)
  - [func \(x \*FlushResponse\) SetSuccess\(v bool\)](<#FlushResponse.SetSuccess>)
  - [func \(x \*FlushResponse\) String\(\) string](<#FlushResponse.String>)
- [type FlushResponse\_builder](<#FlushResponse_builder>)
  - [func \(b0 FlushResponse\_builder\) Build\(\) \*FlushResponse](<#FlushResponse_builder.Build>)
- [type GcRequest](<#GcRequest>)
  - [func \(x \*GcRequest\) ClearMetadata\(\)](<#GcRequest.ClearMetadata>)
  - [func \(x \*GcRequest\) ClearNamespace\(\)](<#GcRequest.ClearNamespace>)
  - [func \(x \*GcRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GcRequest.GetMetadata>)
  - [func \(x \*GcRequest\) GetNamespace\(\) string](<#GcRequest.GetNamespace>)
  - [func \(x \*GcRequest\) HasMetadata\(\) bool](<#GcRequest.HasMetadata>)
  - [func \(x \*GcRequest\) HasNamespace\(\) bool](<#GcRequest.HasNamespace>)
  - [func \(\*GcRequest\) ProtoMessage\(\)](<#GcRequest.ProtoMessage>)
  - [func \(x \*GcRequest\) ProtoReflect\(\) protoreflect.Message](<#GcRequest.ProtoReflect>)
  - [func \(x \*GcRequest\) Reset\(\)](<#GcRequest.Reset>)
  - [func \(x \*GcRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GcRequest.SetMetadata>)
  - [func \(x \*GcRequest\) SetNamespace\(v string\)](<#GcRequest.SetNamespace>)
  - [func \(x \*GcRequest\) String\(\) string](<#GcRequest.String>)
- [type GcRequest\_builder](<#GcRequest_builder>)
  - [func \(b0 GcRequest\_builder\) Build\(\) \*GcRequest](<#GcRequest_builder.Build>)
- [type GetConnectionInfoRequest](<#GetConnectionInfoRequest>)
  - [func \(x \*GetConnectionInfoRequest\) ClearMetadata\(\)](<#GetConnectionInfoRequest.ClearMetadata>)
  - [func \(x \*GetConnectionInfoRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetConnectionInfoRequest.GetMetadata>)
  - [func \(x \*GetConnectionInfoRequest\) HasMetadata\(\) bool](<#GetConnectionInfoRequest.HasMetadata>)
  - [func \(\*GetConnectionInfoRequest\) ProtoMessage\(\)](<#GetConnectionInfoRequest.ProtoMessage>)
  - [func \(x \*GetConnectionInfoRequest\) ProtoReflect\(\) protoreflect.Message](<#GetConnectionInfoRequest.ProtoReflect>)
  - [func \(x \*GetConnectionInfoRequest\) Reset\(\)](<#GetConnectionInfoRequest.Reset>)
  - [func \(x \*GetConnectionInfoRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetConnectionInfoRequest.SetMetadata>)
  - [func \(x \*GetConnectionInfoRequest\) String\(\) string](<#GetConnectionInfoRequest.String>)
- [type GetConnectionInfoRequest\_builder](<#GetConnectionInfoRequest_builder>)
  - [func \(b0 GetConnectionInfoRequest\_builder\) Build\(\) \*GetConnectionInfoRequest](<#GetConnectionInfoRequest_builder.Build>)
- [type GetConnectionInfoResponse](<#GetConnectionInfoResponse>)
  - [func \(x \*GetConnectionInfoResponse\) ClearDatabaseInfo\(\)](<#GetConnectionInfoResponse.ClearDatabaseInfo>)
  - [func \(x \*GetConnectionInfoResponse\) ClearPoolInfo\(\)](<#GetConnectionInfoResponse.ClearPoolInfo>)
  - [func \(x \*GetConnectionInfoResponse\) GetDatabaseInfo\(\) \*DatabaseInfo](<#GetConnectionInfoResponse.GetDatabaseInfo>)
  - [func \(x \*GetConnectionInfoResponse\) GetPoolInfo\(\) \*ConnectionPoolInfo](<#GetConnectionInfoResponse.GetPoolInfo>)
  - [func \(x \*GetConnectionInfoResponse\) HasDatabaseInfo\(\) bool](<#GetConnectionInfoResponse.HasDatabaseInfo>)
  - [func \(x \*GetConnectionInfoResponse\) HasPoolInfo\(\) bool](<#GetConnectionInfoResponse.HasPoolInfo>)
  - [func \(\*GetConnectionInfoResponse\) ProtoMessage\(\)](<#GetConnectionInfoResponse.ProtoMessage>)
  - [func \(x \*GetConnectionInfoResponse\) ProtoReflect\(\) protoreflect.Message](<#GetConnectionInfoResponse.ProtoReflect>)
  - [func \(x \*GetConnectionInfoResponse\) Reset\(\)](<#GetConnectionInfoResponse.Reset>)
  - [func \(x \*GetConnectionInfoResponse\) SetDatabaseInfo\(v \*DatabaseInfo\)](<#GetConnectionInfoResponse.SetDatabaseInfo>)
  - [func \(x \*GetConnectionInfoResponse\) SetPoolInfo\(v \*ConnectionPoolInfo\)](<#GetConnectionInfoResponse.SetPoolInfo>)
  - [func \(x \*GetConnectionInfoResponse\) String\(\) string](<#GetConnectionInfoResponse.String>)
- [type GetConnectionInfoResponse\_builder](<#GetConnectionInfoResponse_builder>)
  - [func \(b0 GetConnectionInfoResponse\_builder\) Build\(\) \*GetConnectionInfoResponse](<#GetConnectionInfoResponse_builder.Build>)
- [type GetDatabaseInfoRequest](<#GetDatabaseInfoRequest>)
  - [func \(x \*GetDatabaseInfoRequest\) ClearMetadata\(\)](<#GetDatabaseInfoRequest.ClearMetadata>)
  - [func \(x \*GetDatabaseInfoRequest\) ClearName\(\)](<#GetDatabaseInfoRequest.ClearName>)
  - [func \(x \*GetDatabaseInfoRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetDatabaseInfoRequest.GetMetadata>)
  - [func \(x \*GetDatabaseInfoRequest\) GetName\(\) string](<#GetDatabaseInfoRequest.GetName>)
  - [func \(x \*GetDatabaseInfoRequest\) HasMetadata\(\) bool](<#GetDatabaseInfoRequest.HasMetadata>)
  - [func \(x \*GetDatabaseInfoRequest\) HasName\(\) bool](<#GetDatabaseInfoRequest.HasName>)
  - [func \(\*GetDatabaseInfoRequest\) ProtoMessage\(\)](<#GetDatabaseInfoRequest.ProtoMessage>)
  - [func \(x \*GetDatabaseInfoRequest\) ProtoReflect\(\) protoreflect.Message](<#GetDatabaseInfoRequest.ProtoReflect>)
  - [func \(x \*GetDatabaseInfoRequest\) Reset\(\)](<#GetDatabaseInfoRequest.Reset>)
  - [func \(x \*GetDatabaseInfoRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetDatabaseInfoRequest.SetMetadata>)
  - [func \(x \*GetDatabaseInfoRequest\) SetName\(v string\)](<#GetDatabaseInfoRequest.SetName>)
  - [func \(x \*GetDatabaseInfoRequest\) String\(\) string](<#GetDatabaseInfoRequest.String>)
- [type GetDatabaseInfoRequest\_builder](<#GetDatabaseInfoRequest_builder>)
  - [func \(b0 GetDatabaseInfoRequest\_builder\) Build\(\) \*GetDatabaseInfoRequest](<#GetDatabaseInfoRequest_builder.Build>)
- [type GetDatabaseInfoResponse](<#GetDatabaseInfoResponse>)
  - [func \(x \*GetDatabaseInfoResponse\) ClearInfo\(\)](<#GetDatabaseInfoResponse.ClearInfo>)
  - [func \(x \*GetDatabaseInfoResponse\) GetInfo\(\) \*DatabaseInfo](<#GetDatabaseInfoResponse.GetInfo>)
  - [func \(x \*GetDatabaseInfoResponse\) HasInfo\(\) bool](<#GetDatabaseInfoResponse.HasInfo>)
  - [func \(\*GetDatabaseInfoResponse\) ProtoMessage\(\)](<#GetDatabaseInfoResponse.ProtoMessage>)
  - [func \(x \*GetDatabaseInfoResponse\) ProtoReflect\(\) protoreflect.Message](<#GetDatabaseInfoResponse.ProtoReflect>)
  - [func \(x \*GetDatabaseInfoResponse\) Reset\(\)](<#GetDatabaseInfoResponse.Reset>)
  - [func \(x \*GetDatabaseInfoResponse\) SetInfo\(v \*DatabaseInfo\)](<#GetDatabaseInfoResponse.SetInfo>)
  - [func \(x \*GetDatabaseInfoResponse\) String\(\) string](<#GetDatabaseInfoResponse.String>)
- [type GetDatabaseInfoResponse\_builder](<#GetDatabaseInfoResponse_builder>)
  - [func \(b0 GetDatabaseInfoResponse\_builder\) Build\(\) \*GetDatabaseInfoResponse](<#GetDatabaseInfoResponse_builder.Build>)
- [type GetMemoryUsageRequest](<#GetMemoryUsageRequest>)
  - [func \(x \*GetMemoryUsageRequest\) ClearMetadata\(\)](<#GetMemoryUsageRequest.ClearMetadata>)
  - [func \(x \*GetMemoryUsageRequest\) ClearNamespace\(\)](<#GetMemoryUsageRequest.ClearNamespace>)
  - [func \(x \*GetMemoryUsageRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetMemoryUsageRequest.GetMetadata>)
  - [func \(x \*GetMemoryUsageRequest\) GetNamespace\(\) string](<#GetMemoryUsageRequest.GetNamespace>)
  - [func \(x \*GetMemoryUsageRequest\) HasMetadata\(\) bool](<#GetMemoryUsageRequest.HasMetadata>)
  - [func \(x \*GetMemoryUsageRequest\) HasNamespace\(\) bool](<#GetMemoryUsageRequest.HasNamespace>)
  - [func \(\*GetMemoryUsageRequest\) ProtoMessage\(\)](<#GetMemoryUsageRequest.ProtoMessage>)
  - [func \(x \*GetMemoryUsageRequest\) ProtoReflect\(\) protoreflect.Message](<#GetMemoryUsageRequest.ProtoReflect>)
  - [func \(x \*GetMemoryUsageRequest\) Reset\(\)](<#GetMemoryUsageRequest.Reset>)
  - [func \(x \*GetMemoryUsageRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetMemoryUsageRequest.SetMetadata>)
  - [func \(x \*GetMemoryUsageRequest\) SetNamespace\(v string\)](<#GetMemoryUsageRequest.SetNamespace>)
  - [func \(x \*GetMemoryUsageRequest\) String\(\) string](<#GetMemoryUsageRequest.String>)
- [type GetMemoryUsageRequest\_builder](<#GetMemoryUsageRequest_builder>)
  - [func \(b0 GetMemoryUsageRequest\_builder\) Build\(\) \*GetMemoryUsageRequest](<#GetMemoryUsageRequest_builder.Build>)
- [type GetMemoryUsageResponse](<#GetMemoryUsageResponse>)
  - [func \(x \*GetMemoryUsageResponse\) ClearError\(\)](<#GetMemoryUsageResponse.ClearError>)
  - [func \(x \*GetMemoryUsageResponse\) ClearMemoryUsageBytes\(\)](<#GetMemoryUsageResponse.ClearMemoryUsageBytes>)
  - [func \(x \*GetMemoryUsageResponse\) ClearMemoryUsagePercent\(\)](<#GetMemoryUsageResponse.ClearMemoryUsagePercent>)
  - [func \(x \*GetMemoryUsageResponse\) GetError\(\) \*common.Error](<#GetMemoryUsageResponse.GetError>)
  - [func \(x \*GetMemoryUsageResponse\) GetMemoryUsageBytes\(\) int64](<#GetMemoryUsageResponse.GetMemoryUsageBytes>)
  - [func \(x \*GetMemoryUsageResponse\) GetMemoryUsagePercent\(\) float64](<#GetMemoryUsageResponse.GetMemoryUsagePercent>)
  - [func \(x \*GetMemoryUsageResponse\) HasError\(\) bool](<#GetMemoryUsageResponse.HasError>)
  - [func \(x \*GetMemoryUsageResponse\) HasMemoryUsageBytes\(\) bool](<#GetMemoryUsageResponse.HasMemoryUsageBytes>)
  - [func \(x \*GetMemoryUsageResponse\) HasMemoryUsagePercent\(\) bool](<#GetMemoryUsageResponse.HasMemoryUsagePercent>)
  - [func \(\*GetMemoryUsageResponse\) ProtoMessage\(\)](<#GetMemoryUsageResponse.ProtoMessage>)
  - [func \(x \*GetMemoryUsageResponse\) ProtoReflect\(\) protoreflect.Message](<#GetMemoryUsageResponse.ProtoReflect>)
  - [func \(x \*GetMemoryUsageResponse\) Reset\(\)](<#GetMemoryUsageResponse.Reset>)
  - [func \(x \*GetMemoryUsageResponse\) SetError\(v \*common.Error\)](<#GetMemoryUsageResponse.SetError>)
  - [func \(x \*GetMemoryUsageResponse\) SetMemoryUsageBytes\(v int64\)](<#GetMemoryUsageResponse.SetMemoryUsageBytes>)
  - [func \(x \*GetMemoryUsageResponse\) SetMemoryUsagePercent\(v float64\)](<#GetMemoryUsageResponse.SetMemoryUsagePercent>)
  - [func \(x \*GetMemoryUsageResponse\) String\(\) string](<#GetMemoryUsageResponse.String>)
- [type GetMemoryUsageResponse\_builder](<#GetMemoryUsageResponse_builder>)
  - [func \(b0 GetMemoryUsageResponse\_builder\) Build\(\) \*GetMemoryUsageResponse](<#GetMemoryUsageResponse_builder.Build>)
- [type GetMigrationStatusRequest](<#GetMigrationStatusRequest>)
  - [func \(x \*GetMigrationStatusRequest\) ClearDatabase\(\)](<#GetMigrationStatusRequest.ClearDatabase>)
  - [func \(x \*GetMigrationStatusRequest\) ClearMetadata\(\)](<#GetMigrationStatusRequest.ClearMetadata>)
  - [func \(x \*GetMigrationStatusRequest\) GetDatabase\(\) string](<#GetMigrationStatusRequest.GetDatabase>)
  - [func \(x \*GetMigrationStatusRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetMigrationStatusRequest.GetMetadata>)
  - [func \(x \*GetMigrationStatusRequest\) HasDatabase\(\) bool](<#GetMigrationStatusRequest.HasDatabase>)
  - [func \(x \*GetMigrationStatusRequest\) HasMetadata\(\) bool](<#GetMigrationStatusRequest.HasMetadata>)
  - [func \(\*GetMigrationStatusRequest\) ProtoMessage\(\)](<#GetMigrationStatusRequest.ProtoMessage>)
  - [func \(x \*GetMigrationStatusRequest\) ProtoReflect\(\) protoreflect.Message](<#GetMigrationStatusRequest.ProtoReflect>)
  - [func \(x \*GetMigrationStatusRequest\) Reset\(\)](<#GetMigrationStatusRequest.Reset>)
  - [func \(x \*GetMigrationStatusRequest\) SetDatabase\(v string\)](<#GetMigrationStatusRequest.SetDatabase>)
  - [func \(x \*GetMigrationStatusRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetMigrationStatusRequest.SetMetadata>)
  - [func \(x \*GetMigrationStatusRequest\) String\(\) string](<#GetMigrationStatusRequest.String>)
- [type GetMigrationStatusRequest\_builder](<#GetMigrationStatusRequest_builder>)
  - [func \(b0 GetMigrationStatusRequest\_builder\) Build\(\) \*GetMigrationStatusRequest](<#GetMigrationStatusRequest_builder.Build>)
- [type GetMigrationStatusResponse](<#GetMigrationStatusResponse>)
  - [func \(x \*GetMigrationStatusResponse\) ClearCurrentVersion\(\)](<#GetMigrationStatusResponse.ClearCurrentVersion>)
  - [func \(x \*GetMigrationStatusResponse\) GetAppliedVersions\(\) \[\]string](<#GetMigrationStatusResponse.GetAppliedVersions>)
  - [func \(x \*GetMigrationStatusResponse\) GetCurrentVersion\(\) string](<#GetMigrationStatusResponse.GetCurrentVersion>)
  - [func \(x \*GetMigrationStatusResponse\) GetPendingVersions\(\) \[\]string](<#GetMigrationStatusResponse.GetPendingVersions>)
  - [func \(x \*GetMigrationStatusResponse\) HasCurrentVersion\(\) bool](<#GetMigrationStatusResponse.HasCurrentVersion>)
  - [func \(\*GetMigrationStatusResponse\) ProtoMessage\(\)](<#GetMigrationStatusResponse.ProtoMessage>)
  - [func \(x \*GetMigrationStatusResponse\) ProtoReflect\(\) protoreflect.Message](<#GetMigrationStatusResponse.ProtoReflect>)
  - [func \(x \*GetMigrationStatusResponse\) Reset\(\)](<#GetMigrationStatusResponse.Reset>)
  - [func \(x \*GetMigrationStatusResponse\) SetAppliedVersions\(v \[\]string\)](<#GetMigrationStatusResponse.SetAppliedVersions>)
  - [func \(x \*GetMigrationStatusResponse\) SetCurrentVersion\(v string\)](<#GetMigrationStatusResponse.SetCurrentVersion>)
  - [func \(x \*GetMigrationStatusResponse\) SetPendingVersions\(v \[\]string\)](<#GetMigrationStatusResponse.SetPendingVersions>)
  - [func \(x \*GetMigrationStatusResponse\) String\(\) string](<#GetMigrationStatusResponse.String>)
- [type GetMigrationStatusResponse\_builder](<#GetMigrationStatusResponse_builder>)
  - [func \(b0 GetMigrationStatusResponse\_builder\) Build\(\) \*GetMigrationStatusResponse](<#GetMigrationStatusResponse_builder.Build>)
- [type GetMultipleRequest](<#GetMultipleRequest>)
  - [func \(x \*GetMultipleRequest\) ClearMetadata\(\)](<#GetMultipleRequest.ClearMetadata>)
  - [func \(x \*GetMultipleRequest\) GetKeys\(\) \[\]string](<#GetMultipleRequest.GetKeys>)
  - [func \(x \*GetMultipleRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetMultipleRequest.GetMetadata>)
  - [func \(x \*GetMultipleRequest\) HasMetadata\(\) bool](<#GetMultipleRequest.HasMetadata>)
  - [func \(\*GetMultipleRequest\) ProtoMessage\(\)](<#GetMultipleRequest.ProtoMessage>)
  - [func \(x \*GetMultipleRequest\) ProtoReflect\(\) protoreflect.Message](<#GetMultipleRequest.ProtoReflect>)
  - [func \(x \*GetMultipleRequest\) Reset\(\)](<#GetMultipleRequest.Reset>)
  - [func \(x \*GetMultipleRequest\) SetKeys\(v \[\]string\)](<#GetMultipleRequest.SetKeys>)
  - [func \(x \*GetMultipleRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetMultipleRequest.SetMetadata>)
  - [func \(x \*GetMultipleRequest\) String\(\) string](<#GetMultipleRequest.String>)
- [type GetMultipleRequest\_builder](<#GetMultipleRequest_builder>)
  - [func \(b0 GetMultipleRequest\_builder\) Build\(\) \*GetMultipleRequest](<#GetMultipleRequest_builder.Build>)
- [type GetMultipleResponse](<#GetMultipleResponse>)
  - [func \(x \*GetMultipleResponse\) ClearError\(\)](<#GetMultipleResponse.ClearError>)
  - [func \(x \*GetMultipleResponse\) GetError\(\) \*common.Error](<#GetMultipleResponse.GetError>)
  - [func \(x \*GetMultipleResponse\) GetMissingKeys\(\) \[\]string](<#GetMultipleResponse.GetMissingKeys>)
  - [func \(x \*GetMultipleResponse\) GetValues\(\) map\[string\]\[\]byte](<#GetMultipleResponse.GetValues>)
  - [func \(x \*GetMultipleResponse\) HasError\(\) bool](<#GetMultipleResponse.HasError>)
  - [func \(\*GetMultipleResponse\) ProtoMessage\(\)](<#GetMultipleResponse.ProtoMessage>)
  - [func \(x \*GetMultipleResponse\) ProtoReflect\(\) protoreflect.Message](<#GetMultipleResponse.ProtoReflect>)
  - [func \(x \*GetMultipleResponse\) Reset\(\)](<#GetMultipleResponse.Reset>)
  - [func \(x \*GetMultipleResponse\) SetError\(v \*common.Error\)](<#GetMultipleResponse.SetError>)
  - [func \(x \*GetMultipleResponse\) SetMissingKeys\(v \[\]string\)](<#GetMultipleResponse.SetMissingKeys>)
  - [func \(x \*GetMultipleResponse\) SetValues\(v map\[string\]\[\]byte\)](<#GetMultipleResponse.SetValues>)
  - [func \(x \*GetMultipleResponse\) String\(\) string](<#GetMultipleResponse.String>)
- [type GetMultipleResponse\_builder](<#GetMultipleResponse_builder>)
  - [func \(b0 GetMultipleResponse\_builder\) Build\(\) \*GetMultipleResponse](<#GetMultipleResponse_builder.Build>)
- [type GetNamespaceStatsRequest](<#GetNamespaceStatsRequest>)
  - [func \(x \*GetNamespaceStatsRequest\) ClearIncludeDetailedMetrics\(\)](<#GetNamespaceStatsRequest.ClearIncludeDetailedMetrics>)
  - [func \(x \*GetNamespaceStatsRequest\) ClearIncludeKeyDistribution\(\)](<#GetNamespaceStatsRequest.ClearIncludeKeyDistribution>)
  - [func \(x \*GetNamespaceStatsRequest\) ClearNamespaceId\(\)](<#GetNamespaceStatsRequest.ClearNamespaceId>)
  - [func \(x \*GetNamespaceStatsRequest\) GetIncludeDetailedMetrics\(\) bool](<#GetNamespaceStatsRequest.GetIncludeDetailedMetrics>)
  - [func \(x \*GetNamespaceStatsRequest\) GetIncludeKeyDistribution\(\) bool](<#GetNamespaceStatsRequest.GetIncludeKeyDistribution>)
  - [func \(x \*GetNamespaceStatsRequest\) GetNamespaceId\(\) string](<#GetNamespaceStatsRequest.GetNamespaceId>)
  - [func \(x \*GetNamespaceStatsRequest\) HasIncludeDetailedMetrics\(\) bool](<#GetNamespaceStatsRequest.HasIncludeDetailedMetrics>)
  - [func \(x \*GetNamespaceStatsRequest\) HasIncludeKeyDistribution\(\) bool](<#GetNamespaceStatsRequest.HasIncludeKeyDistribution>)
  - [func \(x \*GetNamespaceStatsRequest\) HasNamespaceId\(\) bool](<#GetNamespaceStatsRequest.HasNamespaceId>)
  - [func \(\*GetNamespaceStatsRequest\) ProtoMessage\(\)](<#GetNamespaceStatsRequest.ProtoMessage>)
  - [func \(x \*GetNamespaceStatsRequest\) ProtoReflect\(\) protoreflect.Message](<#GetNamespaceStatsRequest.ProtoReflect>)
  - [func \(x \*GetNamespaceStatsRequest\) Reset\(\)](<#GetNamespaceStatsRequest.Reset>)
  - [func \(x \*GetNamespaceStatsRequest\) SetIncludeDetailedMetrics\(v bool\)](<#GetNamespaceStatsRequest.SetIncludeDetailedMetrics>)
  - [func \(x \*GetNamespaceStatsRequest\) SetIncludeKeyDistribution\(v bool\)](<#GetNamespaceStatsRequest.SetIncludeKeyDistribution>)
  - [func \(x \*GetNamespaceStatsRequest\) SetNamespaceId\(v string\)](<#GetNamespaceStatsRequest.SetNamespaceId>)
  - [func \(x \*GetNamespaceStatsRequest\) String\(\) string](<#GetNamespaceStatsRequest.String>)
- [type GetNamespaceStatsRequest\_builder](<#GetNamespaceStatsRequest_builder>)
  - [func \(b0 GetNamespaceStatsRequest\_builder\) Build\(\) \*GetNamespaceStatsRequest](<#GetNamespaceStatsRequest_builder.Build>)
- [type GetNamespaceStatsResponse](<#GetNamespaceStatsResponse>)
  - [func \(x \*GetNamespaceStatsResponse\) ClearCollectedAt\(\)](<#GetNamespaceStatsResponse.ClearCollectedAt>)
  - [func \(x \*GetNamespaceStatsResponse\) ClearNamespaceId\(\)](<#GetNamespaceStatsResponse.ClearNamespaceId>)
  - [func \(x \*GetNamespaceStatsResponse\) ClearStats\(\)](<#GetNamespaceStatsResponse.ClearStats>)
  - [func \(x \*GetNamespaceStatsResponse\) GetCollectedAt\(\) \*timestamppb.Timestamp](<#GetNamespaceStatsResponse.GetCollectedAt>)
  - [func \(x \*GetNamespaceStatsResponse\) GetNamespaceId\(\) string](<#GetNamespaceStatsResponse.GetNamespaceId>)
  - [func \(x \*GetNamespaceStatsResponse\) GetStats\(\) \*NamespaceStats](<#GetNamespaceStatsResponse.GetStats>)
  - [func \(x \*GetNamespaceStatsResponse\) HasCollectedAt\(\) bool](<#GetNamespaceStatsResponse.HasCollectedAt>)
  - [func \(x \*GetNamespaceStatsResponse\) HasNamespaceId\(\) bool](<#GetNamespaceStatsResponse.HasNamespaceId>)
  - [func \(x \*GetNamespaceStatsResponse\) HasStats\(\) bool](<#GetNamespaceStatsResponse.HasStats>)
  - [func \(\*GetNamespaceStatsResponse\) ProtoMessage\(\)](<#GetNamespaceStatsResponse.ProtoMessage>)
  - [func \(x \*GetNamespaceStatsResponse\) ProtoReflect\(\) protoreflect.Message](<#GetNamespaceStatsResponse.ProtoReflect>)
  - [func \(x \*GetNamespaceStatsResponse\) Reset\(\)](<#GetNamespaceStatsResponse.Reset>)
  - [func \(x \*GetNamespaceStatsResponse\) SetCollectedAt\(v \*timestamppb.Timestamp\)](<#GetNamespaceStatsResponse.SetCollectedAt>)
  - [func \(x \*GetNamespaceStatsResponse\) SetNamespaceId\(v string\)](<#GetNamespaceStatsResponse.SetNamespaceId>)
  - [func \(x \*GetNamespaceStatsResponse\) SetStats\(v \*NamespaceStats\)](<#GetNamespaceStatsResponse.SetStats>)
  - [func \(x \*GetNamespaceStatsResponse\) String\(\) string](<#GetNamespaceStatsResponse.String>)
- [type GetNamespaceStatsResponse\_builder](<#GetNamespaceStatsResponse_builder>)
  - [func \(b0 GetNamespaceStatsResponse\_builder\) Build\(\) \*GetNamespaceStatsResponse](<#GetNamespaceStatsResponse_builder.Build>)
- [type GetRequest](<#GetRequest>)
  - [func \(x \*GetRequest\) ClearKey\(\)](<#GetRequest.ClearKey>)
  - [func \(x \*GetRequest\) ClearMetadata\(\)](<#GetRequest.ClearMetadata>)
  - [func \(x \*GetRequest\) ClearNamespace\(\)](<#GetRequest.ClearNamespace>)
  - [func \(x \*GetRequest\) ClearUpdateAccessTime\(\)](<#GetRequest.ClearUpdateAccessTime>)
  - [func \(x \*GetRequest\) GetKey\(\) string](<#GetRequest.GetKey>)
  - [func \(x \*GetRequest\) GetMetadata\(\) \*common.RequestMetadata](<#GetRequest.GetMetadata>)
  - [func \(x \*GetRequest\) GetNamespace\(\) string](<#GetRequest.GetNamespace>)
  - [func \(x \*GetRequest\) GetUpdateAccessTime\(\) bool](<#GetRequest.GetUpdateAccessTime>)
  - [func \(x \*GetRequest\) HasKey\(\) bool](<#GetRequest.HasKey>)
  - [func \(x \*GetRequest\) HasMetadata\(\) bool](<#GetRequest.HasMetadata>)
  - [func \(x \*GetRequest\) HasNamespace\(\) bool](<#GetRequest.HasNamespace>)
  - [func \(x \*GetRequest\) HasUpdateAccessTime\(\) bool](<#GetRequest.HasUpdateAccessTime>)
  - [func \(\*GetRequest\) ProtoMessage\(\)](<#GetRequest.ProtoMessage>)
  - [func \(x \*GetRequest\) ProtoReflect\(\) protoreflect.Message](<#GetRequest.ProtoReflect>)
  - [func \(x \*GetRequest\) Reset\(\)](<#GetRequest.Reset>)
  - [func \(x \*GetRequest\) SetKey\(v string\)](<#GetRequest.SetKey>)
  - [func \(x \*GetRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#GetRequest.SetMetadata>)
  - [func \(x \*GetRequest\) SetNamespace\(v string\)](<#GetRequest.SetNamespace>)
  - [func \(x \*GetRequest\) SetUpdateAccessTime\(v bool\)](<#GetRequest.SetUpdateAccessTime>)
  - [func \(x \*GetRequest\) String\(\) string](<#GetRequest.String>)
- [type GetRequest\_builder](<#GetRequest_builder>)
  - [func \(b0 GetRequest\_builder\) Build\(\) \*GetRequest](<#GetRequest_builder.Build>)
- [type GetResponse](<#GetResponse>)
  - [func \(x \*GetResponse\) ClearCacheHit\(\)](<#GetResponse.ClearCacheHit>)
  - [func \(x \*GetResponse\) ClearEntry\(\)](<#GetResponse.ClearEntry>)
  - [func \(x \*GetResponse\) ClearFound\(\)](<#GetResponse.ClearFound>)
  - [func \(x \*GetResponse\) GetCacheHit\(\) bool](<#GetResponse.GetCacheHit>)
  - [func \(x \*GetResponse\) GetEntry\(\) \*CacheEntry](<#GetResponse.GetEntry>)
  - [func \(x \*GetResponse\) GetFound\(\) bool](<#GetResponse.GetFound>)
  - [func \(x \*GetResponse\) HasCacheHit\(\) bool](<#GetResponse.HasCacheHit>)
  - [func \(x \*GetResponse\) HasEntry\(\) bool](<#GetResponse.HasEntry>)
  - [func \(x \*GetResponse\) HasFound\(\) bool](<#GetResponse.HasFound>)
  - [func \(\*GetResponse\) ProtoMessage\(\)](<#GetResponse.ProtoMessage>)
  - [func \(x \*GetResponse\) ProtoReflect\(\) protoreflect.Message](<#GetResponse.ProtoReflect>)
  - [func \(x \*GetResponse\) Reset\(\)](<#GetResponse.Reset>)
  - [func \(x \*GetResponse\) SetCacheHit\(v bool\)](<#GetResponse.SetCacheHit>)
  - [func \(x \*GetResponse\) SetEntry\(v \*CacheEntry\)](<#GetResponse.SetEntry>)
  - [func \(x \*GetResponse\) SetFound\(v bool\)](<#GetResponse.SetFound>)
  - [func \(x \*GetResponse\) String\(\) string](<#GetResponse.String>)
- [type GetResponse\_builder](<#GetResponse_builder>)
  - [func \(b0 GetResponse\_builder\) Build\(\) \*GetResponse](<#GetResponse_builder.Build>)
- [type ImportRequest](<#ImportRequest>)
  - [func \(x \*ImportRequest\) ClearMetadata\(\)](<#ImportRequest.ClearMetadata>)
  - [func \(x \*ImportRequest\) ClearNamespace\(\)](<#ImportRequest.ClearNamespace>)
  - [func \(x \*ImportRequest\) ClearSource\(\)](<#ImportRequest.ClearSource>)
  - [func \(x \*ImportRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ImportRequest.GetMetadata>)
  - [func \(x \*ImportRequest\) GetNamespace\(\) string](<#ImportRequest.GetNamespace>)
  - [func \(x \*ImportRequest\) GetSource\(\) string](<#ImportRequest.GetSource>)
  - [func \(x \*ImportRequest\) HasMetadata\(\) bool](<#ImportRequest.HasMetadata>)
  - [func \(x \*ImportRequest\) HasNamespace\(\) bool](<#ImportRequest.HasNamespace>)
  - [func \(x \*ImportRequest\) HasSource\(\) bool](<#ImportRequest.HasSource>)
  - [func \(\*ImportRequest\) ProtoMessage\(\)](<#ImportRequest.ProtoMessage>)
  - [func \(x \*ImportRequest\) ProtoReflect\(\) protoreflect.Message](<#ImportRequest.ProtoReflect>)
  - [func \(x \*ImportRequest\) Reset\(\)](<#ImportRequest.Reset>)
  - [func \(x \*ImportRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ImportRequest.SetMetadata>)
  - [func \(x \*ImportRequest\) SetNamespace\(v string\)](<#ImportRequest.SetNamespace>)
  - [func \(x \*ImportRequest\) SetSource\(v string\)](<#ImportRequest.SetSource>)
  - [func \(x \*ImportRequest\) String\(\) string](<#ImportRequest.String>)
- [type ImportRequest\_builder](<#ImportRequest_builder>)
  - [func \(b0 ImportRequest\_builder\) Build\(\) \*ImportRequest](<#ImportRequest_builder.Build>)
- [type IncrementRequest](<#IncrementRequest>)
  - [func \(x \*IncrementRequest\) ClearDelta\(\)](<#IncrementRequest.ClearDelta>)
  - [func \(x \*IncrementRequest\) ClearInitialValue\(\)](<#IncrementRequest.ClearInitialValue>)
  - [func \(x \*IncrementRequest\) ClearKey\(\)](<#IncrementRequest.ClearKey>)
  - [func \(x \*IncrementRequest\) ClearMetadata\(\)](<#IncrementRequest.ClearMetadata>)
  - [func \(x \*IncrementRequest\) ClearNamespace\(\)](<#IncrementRequest.ClearNamespace>)
  - [func \(x \*IncrementRequest\) ClearTtl\(\)](<#IncrementRequest.ClearTtl>)
  - [func \(x \*IncrementRequest\) GetDelta\(\) int64](<#IncrementRequest.GetDelta>)
  - [func \(x \*IncrementRequest\) GetInitialValue\(\) int64](<#IncrementRequest.GetInitialValue>)
  - [func \(x \*IncrementRequest\) GetKey\(\) string](<#IncrementRequest.GetKey>)
  - [func \(x \*IncrementRequest\) GetMetadata\(\) \*common.RequestMetadata](<#IncrementRequest.GetMetadata>)
  - [func \(x \*IncrementRequest\) GetNamespace\(\) string](<#IncrementRequest.GetNamespace>)
  - [func \(x \*IncrementRequest\) GetTtl\(\) \*durationpb.Duration](<#IncrementRequest.GetTtl>)
  - [func \(x \*IncrementRequest\) HasDelta\(\) bool](<#IncrementRequest.HasDelta>)
  - [func \(x \*IncrementRequest\) HasInitialValue\(\) bool](<#IncrementRequest.HasInitialValue>)
  - [func \(x \*IncrementRequest\) HasKey\(\) bool](<#IncrementRequest.HasKey>)
  - [func \(x \*IncrementRequest\) HasMetadata\(\) bool](<#IncrementRequest.HasMetadata>)
  - [func \(x \*IncrementRequest\) HasNamespace\(\) bool](<#IncrementRequest.HasNamespace>)
  - [func \(x \*IncrementRequest\) HasTtl\(\) bool](<#IncrementRequest.HasTtl>)
  - [func \(\*IncrementRequest\) ProtoMessage\(\)](<#IncrementRequest.ProtoMessage>)
  - [func \(x \*IncrementRequest\) ProtoReflect\(\) protoreflect.Message](<#IncrementRequest.ProtoReflect>)
  - [func \(x \*IncrementRequest\) Reset\(\)](<#IncrementRequest.Reset>)
  - [func \(x \*IncrementRequest\) SetDelta\(v int64\)](<#IncrementRequest.SetDelta>)
  - [func \(x \*IncrementRequest\) SetInitialValue\(v int64\)](<#IncrementRequest.SetInitialValue>)
  - [func \(x \*IncrementRequest\) SetKey\(v string\)](<#IncrementRequest.SetKey>)
  - [func \(x \*IncrementRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#IncrementRequest.SetMetadata>)
  - [func \(x \*IncrementRequest\) SetNamespace\(v string\)](<#IncrementRequest.SetNamespace>)
  - [func \(x \*IncrementRequest\) SetTtl\(v \*durationpb.Duration\)](<#IncrementRequest.SetTtl>)
  - [func \(x \*IncrementRequest\) String\(\) string](<#IncrementRequest.String>)
- [type IncrementRequest\_builder](<#IncrementRequest_builder>)
  - [func \(b0 IncrementRequest\_builder\) Build\(\) \*IncrementRequest](<#IncrementRequest_builder.Build>)
- [type IncrementResponse](<#IncrementResponse>)
  - [func \(x \*IncrementResponse\) ClearError\(\)](<#IncrementResponse.ClearError>)
  - [func \(x \*IncrementResponse\) ClearNewValue\(\)](<#IncrementResponse.ClearNewValue>)
  - [func \(x \*IncrementResponse\) ClearSuccess\(\)](<#IncrementResponse.ClearSuccess>)
  - [func \(x \*IncrementResponse\) GetError\(\) \*common.Error](<#IncrementResponse.GetError>)
  - [func \(x \*IncrementResponse\) GetNewValue\(\) int64](<#IncrementResponse.GetNewValue>)
  - [func \(x \*IncrementResponse\) GetSuccess\(\) bool](<#IncrementResponse.GetSuccess>)
  - [func \(x \*IncrementResponse\) HasError\(\) bool](<#IncrementResponse.HasError>)
  - [func \(x \*IncrementResponse\) HasNewValue\(\) bool](<#IncrementResponse.HasNewValue>)
  - [func \(x \*IncrementResponse\) HasSuccess\(\) bool](<#IncrementResponse.HasSuccess>)
  - [func \(\*IncrementResponse\) ProtoMessage\(\)](<#IncrementResponse.ProtoMessage>)
  - [func \(x \*IncrementResponse\) ProtoReflect\(\) protoreflect.Message](<#IncrementResponse.ProtoReflect>)
  - [func \(x \*IncrementResponse\) Reset\(\)](<#IncrementResponse.Reset>)
  - [func \(x \*IncrementResponse\) SetError\(v \*common.Error\)](<#IncrementResponse.SetError>)
  - [func \(x \*IncrementResponse\) SetNewValue\(v int64\)](<#IncrementResponse.SetNewValue>)
  - [func \(x \*IncrementResponse\) SetSuccess\(v bool\)](<#IncrementResponse.SetSuccess>)
  - [func \(x \*IncrementResponse\) String\(\) string](<#IncrementResponse.String>)
- [type IncrementResponse\_builder](<#IncrementResponse_builder>)
  - [func \(b0 IncrementResponse\_builder\) Build\(\) \*IncrementResponse](<#IncrementResponse_builder.Build>)
- [type InfoRequest](<#InfoRequest>)
  - [func \(x \*InfoRequest\) ClearMetadata\(\)](<#InfoRequest.ClearMetadata>)
  - [func \(x \*InfoRequest\) GetMetadata\(\) \*common.RequestMetadata](<#InfoRequest.GetMetadata>)
  - [func \(x \*InfoRequest\) HasMetadata\(\) bool](<#InfoRequest.HasMetadata>)
  - [func \(\*InfoRequest\) ProtoMessage\(\)](<#InfoRequest.ProtoMessage>)
  - [func \(x \*InfoRequest\) ProtoReflect\(\) protoreflect.Message](<#InfoRequest.ProtoReflect>)
  - [func \(x \*InfoRequest\) Reset\(\)](<#InfoRequest.Reset>)
  - [func \(x \*InfoRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#InfoRequest.SetMetadata>)
  - [func \(x \*InfoRequest\) String\(\) string](<#InfoRequest.String>)
- [type InfoRequest\_builder](<#InfoRequest_builder>)
  - [func \(b0 InfoRequest\_builder\) Build\(\) \*InfoRequest](<#InfoRequest_builder.Build>)
- [type KeysRequest](<#KeysRequest>)
  - [func \(x \*KeysRequest\) ClearMetadata\(\)](<#KeysRequest.ClearMetadata>)
  - [func \(x \*KeysRequest\) ClearNamespace\(\)](<#KeysRequest.ClearNamespace>)
  - [func \(x \*KeysRequest\) ClearPagination\(\)](<#KeysRequest.ClearPagination>)
  - [func \(x \*KeysRequest\) ClearPattern\(\)](<#KeysRequest.ClearPattern>)
  - [func \(x \*KeysRequest\) GetMetadata\(\) \*common.RequestMetadata](<#KeysRequest.GetMetadata>)
  - [func \(x \*KeysRequest\) GetNamespace\(\) string](<#KeysRequest.GetNamespace>)
  - [func \(x \*KeysRequest\) GetPagination\(\) \*common.Pagination](<#KeysRequest.GetPagination>)
  - [func \(x \*KeysRequest\) GetPattern\(\) string](<#KeysRequest.GetPattern>)
  - [func \(x \*KeysRequest\) HasMetadata\(\) bool](<#KeysRequest.HasMetadata>)
  - [func \(x \*KeysRequest\) HasNamespace\(\) bool](<#KeysRequest.HasNamespace>)
  - [func \(x \*KeysRequest\) HasPagination\(\) bool](<#KeysRequest.HasPagination>)
  - [func \(x \*KeysRequest\) HasPattern\(\) bool](<#KeysRequest.HasPattern>)
  - [func \(\*KeysRequest\) ProtoMessage\(\)](<#KeysRequest.ProtoMessage>)
  - [func \(x \*KeysRequest\) ProtoReflect\(\) protoreflect.Message](<#KeysRequest.ProtoReflect>)
  - [func \(x \*KeysRequest\) Reset\(\)](<#KeysRequest.Reset>)
  - [func \(x \*KeysRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#KeysRequest.SetMetadata>)
  - [func \(x \*KeysRequest\) SetNamespace\(v string\)](<#KeysRequest.SetNamespace>)
  - [func \(x \*KeysRequest\) SetPagination\(v \*common.Pagination\)](<#KeysRequest.SetPagination>)
  - [func \(x \*KeysRequest\) SetPattern\(v string\)](<#KeysRequest.SetPattern>)
  - [func \(x \*KeysRequest\) String\(\) string](<#KeysRequest.String>)
- [type KeysRequest\_builder](<#KeysRequest_builder>)
  - [func \(b0 KeysRequest\_builder\) Build\(\) \*KeysRequest](<#KeysRequest_builder.Build>)
- [type KeysResponse](<#KeysResponse>)
  - [func \(x \*KeysResponse\) ClearError\(\)](<#KeysResponse.ClearError>)
  - [func \(x \*KeysResponse\) ClearSuccess\(\)](<#KeysResponse.ClearSuccess>)
  - [func \(x \*KeysResponse\) ClearTotalCount\(\)](<#KeysResponse.ClearTotalCount>)
  - [func \(x \*KeysResponse\) GetError\(\) \*common.Error](<#KeysResponse.GetError>)
  - [func \(x \*KeysResponse\) GetKeys\(\) \[\]string](<#KeysResponse.GetKeys>)
  - [func \(x \*KeysResponse\) GetSuccess\(\) bool](<#KeysResponse.GetSuccess>)
  - [func \(x \*KeysResponse\) GetTotalCount\(\) int64](<#KeysResponse.GetTotalCount>)
  - [func \(x \*KeysResponse\) HasError\(\) bool](<#KeysResponse.HasError>)
  - [func \(x \*KeysResponse\) HasSuccess\(\) bool](<#KeysResponse.HasSuccess>)
  - [func \(x \*KeysResponse\) HasTotalCount\(\) bool](<#KeysResponse.HasTotalCount>)
  - [func \(\*KeysResponse\) ProtoMessage\(\)](<#KeysResponse.ProtoMessage>)
  - [func \(x \*KeysResponse\) ProtoReflect\(\) protoreflect.Message](<#KeysResponse.ProtoReflect>)
  - [func \(x \*KeysResponse\) Reset\(\)](<#KeysResponse.Reset>)
  - [func \(x \*KeysResponse\) SetError\(v \*common.Error\)](<#KeysResponse.SetError>)
  - [func \(x \*KeysResponse\) SetKeys\(v \[\]string\)](<#KeysResponse.SetKeys>)
  - [func \(x \*KeysResponse\) SetSuccess\(v bool\)](<#KeysResponse.SetSuccess>)
  - [func \(x \*KeysResponse\) SetTotalCount\(v int64\)](<#KeysResponse.SetTotalCount>)
  - [func \(x \*KeysResponse\) String\(\) string](<#KeysResponse.String>)
- [type KeysResponse\_builder](<#KeysResponse_builder>)
  - [func \(b0 KeysResponse\_builder\) Build\(\) \*KeysResponse](<#KeysResponse_builder.Build>)
- [type ListDatabasesRequest](<#ListDatabasesRequest>)
  - [func \(x \*ListDatabasesRequest\) ClearMetadata\(\)](<#ListDatabasesRequest.ClearMetadata>)
  - [func \(x \*ListDatabasesRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ListDatabasesRequest.GetMetadata>)
  - [func \(x \*ListDatabasesRequest\) HasMetadata\(\) bool](<#ListDatabasesRequest.HasMetadata>)
  - [func \(\*ListDatabasesRequest\) ProtoMessage\(\)](<#ListDatabasesRequest.ProtoMessage>)
  - [func \(x \*ListDatabasesRequest\) ProtoReflect\(\) protoreflect.Message](<#ListDatabasesRequest.ProtoReflect>)
  - [func \(x \*ListDatabasesRequest\) Reset\(\)](<#ListDatabasesRequest.Reset>)
  - [func \(x \*ListDatabasesRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ListDatabasesRequest.SetMetadata>)
  - [func \(x \*ListDatabasesRequest\) String\(\) string](<#ListDatabasesRequest.String>)
- [type ListDatabasesRequest\_builder](<#ListDatabasesRequest_builder>)
  - [func \(b0 ListDatabasesRequest\_builder\) Build\(\) \*ListDatabasesRequest](<#ListDatabasesRequest_builder.Build>)
- [type ListDatabasesResponse](<#ListDatabasesResponse>)
  - [func \(x \*ListDatabasesResponse\) GetDatabases\(\) \[\]string](<#ListDatabasesResponse.GetDatabases>)
  - [func \(\*ListDatabasesResponse\) ProtoMessage\(\)](<#ListDatabasesResponse.ProtoMessage>)
  - [func \(x \*ListDatabasesResponse\) ProtoReflect\(\) protoreflect.Message](<#ListDatabasesResponse.ProtoReflect>)
  - [func \(x \*ListDatabasesResponse\) Reset\(\)](<#ListDatabasesResponse.Reset>)
  - [func \(x \*ListDatabasesResponse\) SetDatabases\(v \[\]string\)](<#ListDatabasesResponse.SetDatabases>)
  - [func \(x \*ListDatabasesResponse\) String\(\) string](<#ListDatabasesResponse.String>)
- [type ListDatabasesResponse\_builder](<#ListDatabasesResponse_builder>)
  - [func \(b0 ListDatabasesResponse\_builder\) Build\(\) \*ListDatabasesResponse](<#ListDatabasesResponse_builder.Build>)
- [type ListMigrationsRequest](<#ListMigrationsRequest>)
  - [func \(x \*ListMigrationsRequest\) ClearDatabase\(\)](<#ListMigrationsRequest.ClearDatabase>)
  - [func \(x \*ListMigrationsRequest\) ClearMetadata\(\)](<#ListMigrationsRequest.ClearMetadata>)
  - [func \(x \*ListMigrationsRequest\) ClearStatusFilter\(\)](<#ListMigrationsRequest.ClearStatusFilter>)
  - [func \(x \*ListMigrationsRequest\) GetDatabase\(\) string](<#ListMigrationsRequest.GetDatabase>)
  - [func \(x \*ListMigrationsRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ListMigrationsRequest.GetMetadata>)
  - [func \(x \*ListMigrationsRequest\) GetStatusFilter\(\) string](<#ListMigrationsRequest.GetStatusFilter>)
  - [func \(x \*ListMigrationsRequest\) HasDatabase\(\) bool](<#ListMigrationsRequest.HasDatabase>)
  - [func \(x \*ListMigrationsRequest\) HasMetadata\(\) bool](<#ListMigrationsRequest.HasMetadata>)
  - [func \(x \*ListMigrationsRequest\) HasStatusFilter\(\) bool](<#ListMigrationsRequest.HasStatusFilter>)
  - [func \(\*ListMigrationsRequest\) ProtoMessage\(\)](<#ListMigrationsRequest.ProtoMessage>)
  - [func \(x \*ListMigrationsRequest\) ProtoReflect\(\) protoreflect.Message](<#ListMigrationsRequest.ProtoReflect>)
  - [func \(x \*ListMigrationsRequest\) Reset\(\)](<#ListMigrationsRequest.Reset>)
  - [func \(x \*ListMigrationsRequest\) SetDatabase\(v string\)](<#ListMigrationsRequest.SetDatabase>)
  - [func \(x \*ListMigrationsRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ListMigrationsRequest.SetMetadata>)
  - [func \(x \*ListMigrationsRequest\) SetStatusFilter\(v string\)](<#ListMigrationsRequest.SetStatusFilter>)
  - [func \(x \*ListMigrationsRequest\) String\(\) string](<#ListMigrationsRequest.String>)
- [type ListMigrationsRequest\_builder](<#ListMigrationsRequest_builder>)
  - [func \(b0 ListMigrationsRequest\_builder\) Build\(\) \*ListMigrationsRequest](<#ListMigrationsRequest_builder.Build>)
- [type ListMigrationsResponse](<#ListMigrationsResponse>)
  - [func \(x \*ListMigrationsResponse\) GetMigrations\(\) \[\]\*MigrationInfo](<#ListMigrationsResponse.GetMigrations>)
  - [func \(\*ListMigrationsResponse\) ProtoMessage\(\)](<#ListMigrationsResponse.ProtoMessage>)
  - [func \(x \*ListMigrationsResponse\) ProtoReflect\(\) protoreflect.Message](<#ListMigrationsResponse.ProtoReflect>)
  - [func \(x \*ListMigrationsResponse\) Reset\(\)](<#ListMigrationsResponse.Reset>)
  - [func \(x \*ListMigrationsResponse\) SetMigrations\(v \[\]\*MigrationInfo\)](<#ListMigrationsResponse.SetMigrations>)
  - [func \(x \*ListMigrationsResponse\) String\(\) string](<#ListMigrationsResponse.String>)
- [type ListMigrationsResponse\_builder](<#ListMigrationsResponse_builder>)
  - [func \(b0 ListMigrationsResponse\_builder\) Build\(\) \*ListMigrationsResponse](<#ListMigrationsResponse_builder.Build>)
- [type ListNamespacesRequest](<#ListNamespacesRequest>)
  - [func \(x \*ListNamespacesRequest\) ClearIncludeStats\(\)](<#ListNamespacesRequest.ClearIncludeStats>)
  - [func \(x \*ListNamespacesRequest\) ClearNameFilter\(\)](<#ListNamespacesRequest.ClearNameFilter>)
  - [func \(x \*ListNamespacesRequest\) ClearPage\(\)](<#ListNamespacesRequest.ClearPage>)
  - [func \(x \*ListNamespacesRequest\) ClearPageSize\(\)](<#ListNamespacesRequest.ClearPageSize>)
  - [func \(x \*ListNamespacesRequest\) GetIncludeStats\(\) bool](<#ListNamespacesRequest.GetIncludeStats>)
  - [func \(x \*ListNamespacesRequest\) GetNameFilter\(\) string](<#ListNamespacesRequest.GetNameFilter>)
  - [func \(x \*ListNamespacesRequest\) GetPage\(\) int32](<#ListNamespacesRequest.GetPage>)
  - [func \(x \*ListNamespacesRequest\) GetPageSize\(\) int32](<#ListNamespacesRequest.GetPageSize>)
  - [func \(x \*ListNamespacesRequest\) HasIncludeStats\(\) bool](<#ListNamespacesRequest.HasIncludeStats>)
  - [func \(x \*ListNamespacesRequest\) HasNameFilter\(\) bool](<#ListNamespacesRequest.HasNameFilter>)
  - [func \(x \*ListNamespacesRequest\) HasPage\(\) bool](<#ListNamespacesRequest.HasPage>)
  - [func \(x \*ListNamespacesRequest\) HasPageSize\(\) bool](<#ListNamespacesRequest.HasPageSize>)
  - [func \(\*ListNamespacesRequest\) ProtoMessage\(\)](<#ListNamespacesRequest.ProtoMessage>)
  - [func \(x \*ListNamespacesRequest\) ProtoReflect\(\) protoreflect.Message](<#ListNamespacesRequest.ProtoReflect>)
  - [func \(x \*ListNamespacesRequest\) Reset\(\)](<#ListNamespacesRequest.Reset>)
  - [func \(x \*ListNamespacesRequest\) SetIncludeStats\(v bool\)](<#ListNamespacesRequest.SetIncludeStats>)
  - [func \(x \*ListNamespacesRequest\) SetNameFilter\(v string\)](<#ListNamespacesRequest.SetNameFilter>)
  - [func \(x \*ListNamespacesRequest\) SetPage\(v int32\)](<#ListNamespacesRequest.SetPage>)
  - [func \(x \*ListNamespacesRequest\) SetPageSize\(v int32\)](<#ListNamespacesRequest.SetPageSize>)
  - [func \(x \*ListNamespacesRequest\) String\(\) string](<#ListNamespacesRequest.String>)
- [type ListNamespacesRequest\_builder](<#ListNamespacesRequest_builder>)
  - [func \(b0 ListNamespacesRequest\_builder\) Build\(\) \*ListNamespacesRequest](<#ListNamespacesRequest_builder.Build>)
- [type ListNamespacesResponse](<#ListNamespacesResponse>)
  - [func \(x \*ListNamespacesResponse\) ClearPage\(\)](<#ListNamespacesResponse.ClearPage>)
  - [func \(x \*ListNamespacesResponse\) ClearPageSize\(\)](<#ListNamespacesResponse.ClearPageSize>)
  - [func \(x \*ListNamespacesResponse\) ClearTotalCount\(\)](<#ListNamespacesResponse.ClearTotalCount>)
  - [func \(x \*ListNamespacesResponse\) ClearTotalPages\(\)](<#ListNamespacesResponse.ClearTotalPages>)
  - [func \(x \*ListNamespacesResponse\) GetNamespaces\(\) \[\]\*NamespaceInfo](<#ListNamespacesResponse.GetNamespaces>)
  - [func \(x \*ListNamespacesResponse\) GetPage\(\) int32](<#ListNamespacesResponse.GetPage>)
  - [func \(x \*ListNamespacesResponse\) GetPageSize\(\) int32](<#ListNamespacesResponse.GetPageSize>)
  - [func \(x \*ListNamespacesResponse\) GetTotalCount\(\) int32](<#ListNamespacesResponse.GetTotalCount>)
  - [func \(x \*ListNamespacesResponse\) GetTotalPages\(\) int32](<#ListNamespacesResponse.GetTotalPages>)
  - [func \(x \*ListNamespacesResponse\) HasPage\(\) bool](<#ListNamespacesResponse.HasPage>)
  - [func \(x \*ListNamespacesResponse\) HasPageSize\(\) bool](<#ListNamespacesResponse.HasPageSize>)
  - [func \(x \*ListNamespacesResponse\) HasTotalCount\(\) bool](<#ListNamespacesResponse.HasTotalCount>)
  - [func \(x \*ListNamespacesResponse\) HasTotalPages\(\) bool](<#ListNamespacesResponse.HasTotalPages>)
  - [func \(\*ListNamespacesResponse\) ProtoMessage\(\)](<#ListNamespacesResponse.ProtoMessage>)
  - [func \(x \*ListNamespacesResponse\) ProtoReflect\(\) protoreflect.Message](<#ListNamespacesResponse.ProtoReflect>)
  - [func \(x \*ListNamespacesResponse\) Reset\(\)](<#ListNamespacesResponse.Reset>)
  - [func \(x \*ListNamespacesResponse\) SetNamespaces\(v \[\]\*NamespaceInfo\)](<#ListNamespacesResponse.SetNamespaces>)
  - [func \(x \*ListNamespacesResponse\) SetPage\(v int32\)](<#ListNamespacesResponse.SetPage>)
  - [func \(x \*ListNamespacesResponse\) SetPageSize\(v int32\)](<#ListNamespacesResponse.SetPageSize>)
  - [func \(x \*ListNamespacesResponse\) SetTotalCount\(v int32\)](<#ListNamespacesResponse.SetTotalCount>)
  - [func \(x \*ListNamespacesResponse\) SetTotalPages\(v int32\)](<#ListNamespacesResponse.SetTotalPages>)
  - [func \(x \*ListNamespacesResponse\) String\(\) string](<#ListNamespacesResponse.String>)
- [type ListNamespacesResponse\_builder](<#ListNamespacesResponse_builder>)
  - [func \(b0 ListNamespacesResponse\_builder\) Build\(\) \*ListNamespacesResponse](<#ListNamespacesResponse_builder.Build>)
- [type ListSchemasRequest](<#ListSchemasRequest>)
  - [func \(x \*ListSchemasRequest\) ClearDatabase\(\)](<#ListSchemasRequest.ClearDatabase>)
  - [func \(x \*ListSchemasRequest\) ClearMetadata\(\)](<#ListSchemasRequest.ClearMetadata>)
  - [func \(x \*ListSchemasRequest\) GetDatabase\(\) string](<#ListSchemasRequest.GetDatabase>)
  - [func \(x \*ListSchemasRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ListSchemasRequest.GetMetadata>)
  - [func \(x \*ListSchemasRequest\) HasDatabase\(\) bool](<#ListSchemasRequest.HasDatabase>)
  - [func \(x \*ListSchemasRequest\) HasMetadata\(\) bool](<#ListSchemasRequest.HasMetadata>)
  - [func \(\*ListSchemasRequest\) ProtoMessage\(\)](<#ListSchemasRequest.ProtoMessage>)
  - [func \(x \*ListSchemasRequest\) ProtoReflect\(\) protoreflect.Message](<#ListSchemasRequest.ProtoReflect>)
  - [func \(x \*ListSchemasRequest\) Reset\(\)](<#ListSchemasRequest.Reset>)
  - [func \(x \*ListSchemasRequest\) SetDatabase\(v string\)](<#ListSchemasRequest.SetDatabase>)
  - [func \(x \*ListSchemasRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ListSchemasRequest.SetMetadata>)
  - [func \(x \*ListSchemasRequest\) String\(\) string](<#ListSchemasRequest.String>)
- [type ListSchemasRequest\_builder](<#ListSchemasRequest_builder>)
  - [func \(b0 ListSchemasRequest\_builder\) Build\(\) \*ListSchemasRequest](<#ListSchemasRequest_builder.Build>)
- [type ListSchemasResponse](<#ListSchemasResponse>)
  - [func \(x \*ListSchemasResponse\) GetSchemas\(\) \[\]string](<#ListSchemasResponse.GetSchemas>)
  - [func \(\*ListSchemasResponse\) ProtoMessage\(\)](<#ListSchemasResponse.ProtoMessage>)
  - [func \(x \*ListSchemasResponse\) ProtoReflect\(\) protoreflect.Message](<#ListSchemasResponse.ProtoReflect>)
  - [func \(x \*ListSchemasResponse\) Reset\(\)](<#ListSchemasResponse.Reset>)
  - [func \(x \*ListSchemasResponse\) SetSchemas\(v \[\]string\)](<#ListSchemasResponse.SetSchemas>)
  - [func \(x \*ListSchemasResponse\) String\(\) string](<#ListSchemasResponse.String>)
- [type ListSchemasResponse\_builder](<#ListSchemasResponse_builder>)
  - [func \(b0 ListSchemasResponse\_builder\) Build\(\) \*ListSchemasResponse](<#ListSchemasResponse_builder.Build>)
- [type LockRequest](<#LockRequest>)
  - [func \(x \*LockRequest\) ClearKey\(\)](<#LockRequest.ClearKey>)
  - [func \(x \*LockRequest\) ClearMetadata\(\)](<#LockRequest.ClearMetadata>)
  - [func \(x \*LockRequest\) ClearNamespace\(\)](<#LockRequest.ClearNamespace>)
  - [func \(x \*LockRequest\) ClearTtl\(\)](<#LockRequest.ClearTtl>)
  - [func \(x \*LockRequest\) GetKey\(\) string](<#LockRequest.GetKey>)
  - [func \(x \*LockRequest\) GetMetadata\(\) \*common.RequestMetadata](<#LockRequest.GetMetadata>)
  - [func \(x \*LockRequest\) GetNamespace\(\) string](<#LockRequest.GetNamespace>)
  - [func \(x \*LockRequest\) GetTtl\(\) \*durationpb.Duration](<#LockRequest.GetTtl>)
  - [func \(x \*LockRequest\) HasKey\(\) bool](<#LockRequest.HasKey>)
  - [func \(x \*LockRequest\) HasMetadata\(\) bool](<#LockRequest.HasMetadata>)
  - [func \(x \*LockRequest\) HasNamespace\(\) bool](<#LockRequest.HasNamespace>)
  - [func \(x \*LockRequest\) HasTtl\(\) bool](<#LockRequest.HasTtl>)
  - [func \(\*LockRequest\) ProtoMessage\(\)](<#LockRequest.ProtoMessage>)
  - [func \(x \*LockRequest\) ProtoReflect\(\) protoreflect.Message](<#LockRequest.ProtoReflect>)
  - [func \(x \*LockRequest\) Reset\(\)](<#LockRequest.Reset>)
  - [func \(x \*LockRequest\) SetKey\(v string\)](<#LockRequest.SetKey>)
  - [func \(x \*LockRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#LockRequest.SetMetadata>)
  - [func \(x \*LockRequest\) SetNamespace\(v string\)](<#LockRequest.SetNamespace>)
  - [func \(x \*LockRequest\) SetTtl\(v \*durationpb.Duration\)](<#LockRequest.SetTtl>)
  - [func \(x \*LockRequest\) String\(\) string](<#LockRequest.String>)
- [type LockRequest\_builder](<#LockRequest_builder>)
  - [func \(b0 LockRequest\_builder\) Build\(\) \*LockRequest](<#LockRequest_builder.Build>)
- [type MGetRequest](<#MGetRequest>)
  - [func \(x \*MGetRequest\) ClearIncludeExpired\(\)](<#MGetRequest.ClearIncludeExpired>)
  - [func \(x \*MGetRequest\) ClearMetadata\(\)](<#MGetRequest.ClearMetadata>)
  - [func \(x \*MGetRequest\) ClearNamespace\(\)](<#MGetRequest.ClearNamespace>)
  - [func \(x \*MGetRequest\) ClearUpdateAccessTime\(\)](<#MGetRequest.ClearUpdateAccessTime>)
  - [func \(x \*MGetRequest\) GetIncludeExpired\(\) bool](<#MGetRequest.GetIncludeExpired>)
  - [func \(x \*MGetRequest\) GetKeys\(\) \[\]string](<#MGetRequest.GetKeys>)
  - [func \(x \*MGetRequest\) GetMetadata\(\) \*common.RequestMetadata](<#MGetRequest.GetMetadata>)
  - [func \(x \*MGetRequest\) GetNamespace\(\) string](<#MGetRequest.GetNamespace>)
  - [func \(x \*MGetRequest\) GetUpdateAccessTime\(\) bool](<#MGetRequest.GetUpdateAccessTime>)
  - [func \(x \*MGetRequest\) HasIncludeExpired\(\) bool](<#MGetRequest.HasIncludeExpired>)
  - [func \(x \*MGetRequest\) HasMetadata\(\) bool](<#MGetRequest.HasMetadata>)
  - [func \(x \*MGetRequest\) HasNamespace\(\) bool](<#MGetRequest.HasNamespace>)
  - [func \(x \*MGetRequest\) HasUpdateAccessTime\(\) bool](<#MGetRequest.HasUpdateAccessTime>)
  - [func \(\*MGetRequest\) ProtoMessage\(\)](<#MGetRequest.ProtoMessage>)
  - [func \(x \*MGetRequest\) ProtoReflect\(\) protoreflect.Message](<#MGetRequest.ProtoReflect>)
  - [func \(x \*MGetRequest\) Reset\(\)](<#MGetRequest.Reset>)
  - [func \(x \*MGetRequest\) SetIncludeExpired\(v bool\)](<#MGetRequest.SetIncludeExpired>)
  - [func \(x \*MGetRequest\) SetKeys\(v \[\]string\)](<#MGetRequest.SetKeys>)
  - [func \(x \*MGetRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#MGetRequest.SetMetadata>)
  - [func \(x \*MGetRequest\) SetNamespace\(v string\)](<#MGetRequest.SetNamespace>)
  - [func \(x \*MGetRequest\) SetUpdateAccessTime\(v bool\)](<#MGetRequest.SetUpdateAccessTime>)
  - [func \(x \*MGetRequest\) String\(\) string](<#MGetRequest.String>)
- [type MGetRequest\_builder](<#MGetRequest_builder>)
  - [func \(b0 MGetRequest\_builder\) Build\(\) \*MGetRequest](<#MGetRequest_builder.Build>)
- [type MigrationInfo](<#MigrationInfo>)
  - [func \(x \*MigrationInfo\) ClearAppliedAt\(\)](<#MigrationInfo.ClearAppliedAt>)
  - [func \(x \*MigrationInfo\) ClearDescription\(\)](<#MigrationInfo.ClearDescription>)
  - [func \(x \*MigrationInfo\) ClearVersion\(\)](<#MigrationInfo.ClearVersion>)
  - [func \(x \*MigrationInfo\) GetAppliedAt\(\) \*timestamppb.Timestamp](<#MigrationInfo.GetAppliedAt>)
  - [func \(x \*MigrationInfo\) GetDescription\(\) string](<#MigrationInfo.GetDescription>)
  - [func \(x \*MigrationInfo\) GetVersion\(\) string](<#MigrationInfo.GetVersion>)
  - [func \(x \*MigrationInfo\) HasAppliedAt\(\) bool](<#MigrationInfo.HasAppliedAt>)
  - [func \(x \*MigrationInfo\) HasDescription\(\) bool](<#MigrationInfo.HasDescription>)
  - [func \(x \*MigrationInfo\) HasVersion\(\) bool](<#MigrationInfo.HasVersion>)
  - [func \(\*MigrationInfo\) ProtoMessage\(\)](<#MigrationInfo.ProtoMessage>)
  - [func \(x \*MigrationInfo\) ProtoReflect\(\) protoreflect.Message](<#MigrationInfo.ProtoReflect>)
  - [func \(x \*MigrationInfo\) Reset\(\)](<#MigrationInfo.Reset>)
  - [func \(x \*MigrationInfo\) SetAppliedAt\(v \*timestamppb.Timestamp\)](<#MigrationInfo.SetAppliedAt>)
  - [func \(x \*MigrationInfo\) SetDescription\(v string\)](<#MigrationInfo.SetDescription>)
  - [func \(x \*MigrationInfo\) SetVersion\(v string\)](<#MigrationInfo.SetVersion>)
  - [func \(x \*MigrationInfo\) String\(\) string](<#MigrationInfo.String>)
- [type MigrationInfo\_builder](<#MigrationInfo_builder>)
  - [func \(b0 MigrationInfo\_builder\) Build\(\) \*MigrationInfo](<#MigrationInfo_builder.Build>)
- [type MigrationScript](<#MigrationScript>)
  - [func \(x \*MigrationScript\) ClearDescription\(\)](<#MigrationScript.ClearDescription>)
  - [func \(x \*MigrationScript\) ClearScript\(\)](<#MigrationScript.ClearScript>)
  - [func \(x \*MigrationScript\) ClearVersion\(\)](<#MigrationScript.ClearVersion>)
  - [func \(x \*MigrationScript\) GetDescription\(\) string](<#MigrationScript.GetDescription>)
  - [func \(x \*MigrationScript\) GetScript\(\) string](<#MigrationScript.GetScript>)
  - [func \(x \*MigrationScript\) GetVersion\(\) string](<#MigrationScript.GetVersion>)
  - [func \(x \*MigrationScript\) HasDescription\(\) bool](<#MigrationScript.HasDescription>)
  - [func \(x \*MigrationScript\) HasScript\(\) bool](<#MigrationScript.HasScript>)
  - [func \(x \*MigrationScript\) HasVersion\(\) bool](<#MigrationScript.HasVersion>)
  - [func \(\*MigrationScript\) ProtoMessage\(\)](<#MigrationScript.ProtoMessage>)
  - [func \(x \*MigrationScript\) ProtoReflect\(\) protoreflect.Message](<#MigrationScript.ProtoReflect>)
  - [func \(x \*MigrationScript\) Reset\(\)](<#MigrationScript.Reset>)
  - [func \(x \*MigrationScript\) SetDescription\(v string\)](<#MigrationScript.SetDescription>)
  - [func \(x \*MigrationScript\) SetScript\(v string\)](<#MigrationScript.SetScript>)
  - [func \(x \*MigrationScript\) SetVersion\(v string\)](<#MigrationScript.SetVersion>)
  - [func \(x \*MigrationScript\) String\(\) string](<#MigrationScript.String>)
- [type MigrationScript\_builder](<#MigrationScript_builder>)
  - [func \(b0 MigrationScript\_builder\) Build\(\) \*MigrationScript](<#MigrationScript_builder.Build>)
- [type MigrationServiceClient](<#MigrationServiceClient>)
  - [func NewMigrationServiceClient\(cc grpc.ClientConnInterface\) MigrationServiceClient](<#NewMigrationServiceClient>)
- [type MigrationServiceServer](<#MigrationServiceServer>)
- [type MySQLConfig](<#MySQLConfig>)
  - [func \(x \*MySQLConfig\) ClearConnectTimeoutSeconds\(\)](<#MySQLConfig.ClearConnectTimeoutSeconds>)
  - [func \(x \*MySQLConfig\) ClearDsn\(\)](<#MySQLConfig.ClearDsn>)
  - [func \(x \*MySQLConfig\) ClearMaxIdleConns\(\)](<#MySQLConfig.ClearMaxIdleConns>)
  - [func \(x \*MySQLConfig\) ClearMaxOpenConns\(\)](<#MySQLConfig.ClearMaxOpenConns>)
  - [func \(x \*MySQLConfig\) GetConnectTimeoutSeconds\(\) int32](<#MySQLConfig.GetConnectTimeoutSeconds>)
  - [func \(x \*MySQLConfig\) GetDsn\(\) string](<#MySQLConfig.GetDsn>)
  - [func \(x \*MySQLConfig\) GetMaxIdleConns\(\) int32](<#MySQLConfig.GetMaxIdleConns>)
  - [func \(x \*MySQLConfig\) GetMaxOpenConns\(\) int32](<#MySQLConfig.GetMaxOpenConns>)
  - [func \(x \*MySQLConfig\) HasConnectTimeoutSeconds\(\) bool](<#MySQLConfig.HasConnectTimeoutSeconds>)
  - [func \(x \*MySQLConfig\) HasDsn\(\) bool](<#MySQLConfig.HasDsn>)
  - [func \(x \*MySQLConfig\) HasMaxIdleConns\(\) bool](<#MySQLConfig.HasMaxIdleConns>)
  - [func \(x \*MySQLConfig\) HasMaxOpenConns\(\) bool](<#MySQLConfig.HasMaxOpenConns>)
  - [func \(\*MySQLConfig\) ProtoMessage\(\)](<#MySQLConfig.ProtoMessage>)
  - [func \(x \*MySQLConfig\) ProtoReflect\(\) protoreflect.Message](<#MySQLConfig.ProtoReflect>)
  - [func \(x \*MySQLConfig\) Reset\(\)](<#MySQLConfig.Reset>)
  - [func \(x \*MySQLConfig\) SetConnectTimeoutSeconds\(v int32\)](<#MySQLConfig.SetConnectTimeoutSeconds>)
  - [func \(x \*MySQLConfig\) SetDsn\(v string\)](<#MySQLConfig.SetDsn>)
  - [func \(x \*MySQLConfig\) SetMaxIdleConns\(v int32\)](<#MySQLConfig.SetMaxIdleConns>)
  - [func \(x \*MySQLConfig\) SetMaxOpenConns\(v int32\)](<#MySQLConfig.SetMaxOpenConns>)
  - [func \(x \*MySQLConfig\) String\(\) string](<#MySQLConfig.String>)
- [type MySQLConfig\_builder](<#MySQLConfig_builder>)
  - [func \(b0 MySQLConfig\_builder\) Build\(\) \*MySQLConfig](<#MySQLConfig_builder.Build>)
- [type MySQLStatus](<#MySQLStatus>)
  - [func \(x \*MySQLStatus\) ClearOpenConnections\(\)](<#MySQLStatus.ClearOpenConnections>)
  - [func \(x \*MySQLStatus\) ClearRole\(\)](<#MySQLStatus.ClearRole>)
  - [func \(x \*MySQLStatus\) ClearStartedAt\(\)](<#MySQLStatus.ClearStartedAt>)
  - [func \(x \*MySQLStatus\) ClearVersion\(\)](<#MySQLStatus.ClearVersion>)
  - [func \(x \*MySQLStatus\) GetOpenConnections\(\) int32](<#MySQLStatus.GetOpenConnections>)
  - [func \(x \*MySQLStatus\) GetRole\(\) string](<#MySQLStatus.GetRole>)
  - [func \(x \*MySQLStatus\) GetStartedAt\(\) \*timestamppb.Timestamp](<#MySQLStatus.GetStartedAt>)
  - [func \(x \*MySQLStatus\) GetVersion\(\) string](<#MySQLStatus.GetVersion>)
  - [func \(x \*MySQLStatus\) HasOpenConnections\(\) bool](<#MySQLStatus.HasOpenConnections>)
  - [func \(x \*MySQLStatus\) HasRole\(\) bool](<#MySQLStatus.HasRole>)
  - [func \(x \*MySQLStatus\) HasStartedAt\(\) bool](<#MySQLStatus.HasStartedAt>)
  - [func \(x \*MySQLStatus\) HasVersion\(\) bool](<#MySQLStatus.HasVersion>)
  - [func \(\*MySQLStatus\) ProtoMessage\(\)](<#MySQLStatus.ProtoMessage>)
  - [func \(x \*MySQLStatus\) ProtoReflect\(\) protoreflect.Message](<#MySQLStatus.ProtoReflect>)
  - [func \(x \*MySQLStatus\) Reset\(\)](<#MySQLStatus.Reset>)
  - [func \(x \*MySQLStatus\) SetOpenConnections\(v int32\)](<#MySQLStatus.SetOpenConnections>)
  - [func \(x \*MySQLStatus\) SetRole\(v string\)](<#MySQLStatus.SetRole>)
  - [func \(x \*MySQLStatus\) SetStartedAt\(v \*timestamppb.Timestamp\)](<#MySQLStatus.SetStartedAt>)
  - [func \(x \*MySQLStatus\) SetVersion\(v string\)](<#MySQLStatus.SetVersion>)
  - [func \(x \*MySQLStatus\) String\(\) string](<#MySQLStatus.String>)
- [type MySQLStatus\_builder](<#MySQLStatus_builder>)
  - [func \(b0 MySQLStatus\_builder\) Build\(\) \*MySQLStatus](<#MySQLStatus_builder.Build>)
- [type NamespaceInfo](<#NamespaceInfo>)
  - [func \(x \*NamespaceInfo\) ClearCreatedAt\(\)](<#NamespaceInfo.ClearCreatedAt>)
  - [func \(x \*NamespaceInfo\) ClearCurrentKeys\(\)](<#NamespaceInfo.ClearCurrentKeys>)
  - [func \(x \*NamespaceInfo\) ClearCurrentMemoryBytes\(\)](<#NamespaceInfo.ClearCurrentMemoryBytes>)
  - [func \(x \*NamespaceInfo\) ClearDescription\(\)](<#NamespaceInfo.ClearDescription>)
  - [func \(x \*NamespaceInfo\) ClearName\(\)](<#NamespaceInfo.ClearName>)
  - [func \(x \*NamespaceInfo\) ClearNamespaceId\(\)](<#NamespaceInfo.ClearNamespaceId>)
  - [func \(x \*NamespaceInfo\) GetConfig\(\) map\[string\]string](<#NamespaceInfo.GetConfig>)
  - [func \(x \*NamespaceInfo\) GetCreatedAt\(\) \*timestamppb.Timestamp](<#NamespaceInfo.GetCreatedAt>)
  - [func \(x \*NamespaceInfo\) GetCurrentKeys\(\) int64](<#NamespaceInfo.GetCurrentKeys>)
  - [func \(x \*NamespaceInfo\) GetCurrentMemoryBytes\(\) int64](<#NamespaceInfo.GetCurrentMemoryBytes>)
  - [func \(x \*NamespaceInfo\) GetDescription\(\) string](<#NamespaceInfo.GetDescription>)
  - [func \(x \*NamespaceInfo\) GetName\(\) string](<#NamespaceInfo.GetName>)
  - [func \(x \*NamespaceInfo\) GetNamespaceId\(\) string](<#NamespaceInfo.GetNamespaceId>)
  - [func \(x \*NamespaceInfo\) HasCreatedAt\(\) bool](<#NamespaceInfo.HasCreatedAt>)
  - [func \(x \*NamespaceInfo\) HasCurrentKeys\(\) bool](<#NamespaceInfo.HasCurrentKeys>)
  - [func \(x \*NamespaceInfo\) HasCurrentMemoryBytes\(\) bool](<#NamespaceInfo.HasCurrentMemoryBytes>)
  - [func \(x \*NamespaceInfo\) HasDescription\(\) bool](<#NamespaceInfo.HasDescription>)
  - [func \(x \*NamespaceInfo\) HasName\(\) bool](<#NamespaceInfo.HasName>)
  - [func \(x \*NamespaceInfo\) HasNamespaceId\(\) bool](<#NamespaceInfo.HasNamespaceId>)
  - [func \(\*NamespaceInfo\) ProtoMessage\(\)](<#NamespaceInfo.ProtoMessage>)
  - [func \(x \*NamespaceInfo\) ProtoReflect\(\) protoreflect.Message](<#NamespaceInfo.ProtoReflect>)
  - [func \(x \*NamespaceInfo\) Reset\(\)](<#NamespaceInfo.Reset>)
  - [func \(x \*NamespaceInfo\) SetConfig\(v map\[string\]string\)](<#NamespaceInfo.SetConfig>)
  - [func \(x \*NamespaceInfo\) SetCreatedAt\(v \*timestamppb.Timestamp\)](<#NamespaceInfo.SetCreatedAt>)
  - [func \(x \*NamespaceInfo\) SetCurrentKeys\(v int64\)](<#NamespaceInfo.SetCurrentKeys>)
  - [func \(x \*NamespaceInfo\) SetCurrentMemoryBytes\(v int64\)](<#NamespaceInfo.SetCurrentMemoryBytes>)
  - [func \(x \*NamespaceInfo\) SetDescription\(v string\)](<#NamespaceInfo.SetDescription>)
  - [func \(x \*NamespaceInfo\) SetName\(v string\)](<#NamespaceInfo.SetName>)
  - [func \(x \*NamespaceInfo\) SetNamespaceId\(v string\)](<#NamespaceInfo.SetNamespaceId>)
  - [func \(x \*NamespaceInfo\) String\(\) string](<#NamespaceInfo.String>)
- [type NamespaceInfo\_builder](<#NamespaceInfo_builder>)
  - [func \(b0 NamespaceInfo\_builder\) Build\(\) \*NamespaceInfo](<#NamespaceInfo_builder.Build>)
- [type NamespaceStats](<#NamespaceStats>)
  - [func \(x \*NamespaceStats\) ClearAvgKeySizeBytes\(\)](<#NamespaceStats.ClearAvgKeySizeBytes>)
  - [func \(x \*NamespaceStats\) ClearAvgValueSizeBytes\(\)](<#NamespaceStats.ClearAvgValueSizeBytes>)
  - [func \(x \*NamespaceStats\) ClearCacheHits\(\)](<#NamespaceStats.ClearCacheHits>)
  - [func \(x \*NamespaceStats\) ClearCacheMisses\(\)](<#NamespaceStats.ClearCacheMisses>)
  - [func \(x \*NamespaceStats\) ClearEvictions\(\)](<#NamespaceStats.ClearEvictions>)
  - [func \(x \*NamespaceStats\) ClearHitRatePercent\(\)](<#NamespaceStats.ClearHitRatePercent>)
  - [func \(x \*NamespaceStats\) ClearLastAccessTime\(\)](<#NamespaceStats.ClearLastAccessTime>)
  - [func \(x \*NamespaceStats\) ClearMemoryUsageBytes\(\)](<#NamespaceStats.ClearMemoryUsageBytes>)
  - [func \(x \*NamespaceStats\) ClearTotalKeys\(\)](<#NamespaceStats.ClearTotalKeys>)
  - [func \(x \*NamespaceStats\) GetAvgKeySizeBytes\(\) float64](<#NamespaceStats.GetAvgKeySizeBytes>)
  - [func \(x \*NamespaceStats\) GetAvgValueSizeBytes\(\) float64](<#NamespaceStats.GetAvgValueSizeBytes>)
  - [func \(x \*NamespaceStats\) GetCacheHits\(\) int64](<#NamespaceStats.GetCacheHits>)
  - [func \(x \*NamespaceStats\) GetCacheMisses\(\) int64](<#NamespaceStats.GetCacheMisses>)
  - [func \(x \*NamespaceStats\) GetEvictions\(\) int64](<#NamespaceStats.GetEvictions>)
  - [func \(x \*NamespaceStats\) GetHitRatePercent\(\) float64](<#NamespaceStats.GetHitRatePercent>)
  - [func \(x \*NamespaceStats\) GetLastAccessTime\(\) \*timestamppb.Timestamp](<#NamespaceStats.GetLastAccessTime>)
  - [func \(x \*NamespaceStats\) GetMemoryUsageBytes\(\) int64](<#NamespaceStats.GetMemoryUsageBytes>)
  - [func \(x \*NamespaceStats\) GetTotalKeys\(\) int64](<#NamespaceStats.GetTotalKeys>)
  - [func \(x \*NamespaceStats\) HasAvgKeySizeBytes\(\) bool](<#NamespaceStats.HasAvgKeySizeBytes>)
  - [func \(x \*NamespaceStats\) HasAvgValueSizeBytes\(\) bool](<#NamespaceStats.HasAvgValueSizeBytes>)
  - [func \(x \*NamespaceStats\) HasCacheHits\(\) bool](<#NamespaceStats.HasCacheHits>)
  - [func \(x \*NamespaceStats\) HasCacheMisses\(\) bool](<#NamespaceStats.HasCacheMisses>)
  - [func \(x \*NamespaceStats\) HasEvictions\(\) bool](<#NamespaceStats.HasEvictions>)
  - [func \(x \*NamespaceStats\) HasHitRatePercent\(\) bool](<#NamespaceStats.HasHitRatePercent>)
  - [func \(x \*NamespaceStats\) HasLastAccessTime\(\) bool](<#NamespaceStats.HasLastAccessTime>)
  - [func \(x \*NamespaceStats\) HasMemoryUsageBytes\(\) bool](<#NamespaceStats.HasMemoryUsageBytes>)
  - [func \(x \*NamespaceStats\) HasTotalKeys\(\) bool](<#NamespaceStats.HasTotalKeys>)
  - [func \(\*NamespaceStats\) ProtoMessage\(\)](<#NamespaceStats.ProtoMessage>)
  - [func \(x \*NamespaceStats\) ProtoReflect\(\) protoreflect.Message](<#NamespaceStats.ProtoReflect>)
  - [func \(x \*NamespaceStats\) Reset\(\)](<#NamespaceStats.Reset>)
  - [func \(x \*NamespaceStats\) SetAvgKeySizeBytes\(v float64\)](<#NamespaceStats.SetAvgKeySizeBytes>)
  - [func \(x \*NamespaceStats\) SetAvgValueSizeBytes\(v float64\)](<#NamespaceStats.SetAvgValueSizeBytes>)
  - [func \(x \*NamespaceStats\) SetCacheHits\(v int64\)](<#NamespaceStats.SetCacheHits>)
  - [func \(x \*NamespaceStats\) SetCacheMisses\(v int64\)](<#NamespaceStats.SetCacheMisses>)
  - [func \(x \*NamespaceStats\) SetEvictions\(v int64\)](<#NamespaceStats.SetEvictions>)
  - [func \(x \*NamespaceStats\) SetHitRatePercent\(v float64\)](<#NamespaceStats.SetHitRatePercent>)
  - [func \(x \*NamespaceStats\) SetLastAccessTime\(v \*timestamppb.Timestamp\)](<#NamespaceStats.SetLastAccessTime>)
  - [func \(x \*NamespaceStats\) SetMemoryUsageBytes\(v int64\)](<#NamespaceStats.SetMemoryUsageBytes>)
  - [func \(x \*NamespaceStats\) SetTotalKeys\(v int64\)](<#NamespaceStats.SetTotalKeys>)
  - [func \(x \*NamespaceStats\) String\(\) string](<#NamespaceStats.String>)
- [type NamespaceStats\_builder](<#NamespaceStats_builder>)
  - [func \(b0 NamespaceStats\_builder\) Build\(\) \*NamespaceStats](<#NamespaceStats_builder.Build>)
- [type OptimizeRequest](<#OptimizeRequest>)
  - [func \(x \*OptimizeRequest\) ClearMetadata\(\)](<#OptimizeRequest.ClearMetadata>)
  - [func \(x \*OptimizeRequest\) ClearNamespace\(\)](<#OptimizeRequest.ClearNamespace>)
  - [func \(x \*OptimizeRequest\) GetMetadata\(\) \*common.RequestMetadata](<#OptimizeRequest.GetMetadata>)
  - [func \(x \*OptimizeRequest\) GetNamespace\(\) string](<#OptimizeRequest.GetNamespace>)
  - [func \(x \*OptimizeRequest\) HasMetadata\(\) bool](<#OptimizeRequest.HasMetadata>)
  - [func \(x \*OptimizeRequest\) HasNamespace\(\) bool](<#OptimizeRequest.HasNamespace>)
  - [func \(\*OptimizeRequest\) ProtoMessage\(\)](<#OptimizeRequest.ProtoMessage>)
  - [func \(x \*OptimizeRequest\) ProtoReflect\(\) protoreflect.Message](<#OptimizeRequest.ProtoReflect>)
  - [func \(x \*OptimizeRequest\) Reset\(\)](<#OptimizeRequest.Reset>)
  - [func \(x \*OptimizeRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#OptimizeRequest.SetMetadata>)
  - [func \(x \*OptimizeRequest\) SetNamespace\(v string\)](<#OptimizeRequest.SetNamespace>)
  - [func \(x \*OptimizeRequest\) String\(\) string](<#OptimizeRequest.String>)
- [type OptimizeRequest\_builder](<#OptimizeRequest_builder>)
  - [func \(b0 OptimizeRequest\_builder\) Build\(\) \*OptimizeRequest](<#OptimizeRequest_builder.Build>)
- [type PebbleConfig](<#PebbleConfig>)
  - [func \(x \*PebbleConfig\) ClearCacheSize\(\)](<#PebbleConfig.ClearCacheSize>)
  - [func \(x \*PebbleConfig\) ClearCompression\(\)](<#PebbleConfig.ClearCompression>)
  - [func \(x \*PebbleConfig\) ClearMaxOpenFiles\(\)](<#PebbleConfig.ClearMaxOpenFiles>)
  - [func \(x \*PebbleConfig\) ClearMemtableSize\(\)](<#PebbleConfig.ClearMemtableSize>)
  - [func \(x \*PebbleConfig\) ClearPath\(\)](<#PebbleConfig.ClearPath>)
  - [func \(x \*PebbleConfig\) GetCacheSize\(\) int64](<#PebbleConfig.GetCacheSize>)
  - [func \(x \*PebbleConfig\) GetCompression\(\) bool](<#PebbleConfig.GetCompression>)
  - [func \(x \*PebbleConfig\) GetMaxOpenFiles\(\) int32](<#PebbleConfig.GetMaxOpenFiles>)
  - [func \(x \*PebbleConfig\) GetMemtableSize\(\) int64](<#PebbleConfig.GetMemtableSize>)
  - [func \(x \*PebbleConfig\) GetPath\(\) string](<#PebbleConfig.GetPath>)
  - [func \(x \*PebbleConfig\) HasCacheSize\(\) bool](<#PebbleConfig.HasCacheSize>)
  - [func \(x \*PebbleConfig\) HasCompression\(\) bool](<#PebbleConfig.HasCompression>)
  - [func \(x \*PebbleConfig\) HasMaxOpenFiles\(\) bool](<#PebbleConfig.HasMaxOpenFiles>)
  - [func \(x \*PebbleConfig\) HasMemtableSize\(\) bool](<#PebbleConfig.HasMemtableSize>)
  - [func \(x \*PebbleConfig\) HasPath\(\) bool](<#PebbleConfig.HasPath>)
  - [func \(\*PebbleConfig\) ProtoMessage\(\)](<#PebbleConfig.ProtoMessage>)
  - [func \(x \*PebbleConfig\) ProtoReflect\(\) protoreflect.Message](<#PebbleConfig.ProtoReflect>)
  - [func \(x \*PebbleConfig\) Reset\(\)](<#PebbleConfig.Reset>)
  - [func \(x \*PebbleConfig\) SetCacheSize\(v int64\)](<#PebbleConfig.SetCacheSize>)
  - [func \(x \*PebbleConfig\) SetCompression\(v bool\)](<#PebbleConfig.SetCompression>)
  - [func \(x \*PebbleConfig\) SetMaxOpenFiles\(v int32\)](<#PebbleConfig.SetMaxOpenFiles>)
  - [func \(x \*PebbleConfig\) SetMemtableSize\(v int64\)](<#PebbleConfig.SetMemtableSize>)
  - [func \(x \*PebbleConfig\) SetPath\(v string\)](<#PebbleConfig.SetPath>)
  - [func \(x \*PebbleConfig\) String\(\) string](<#PebbleConfig.String>)
- [type PebbleConfig\_builder](<#PebbleConfig_builder>)
  - [func \(b0 PebbleConfig\_builder\) Build\(\) \*PebbleConfig](<#PebbleConfig_builder.Build>)
- [type PipelineRequest](<#PipelineRequest>)
  - [func \(x \*PipelineRequest\) ClearMetadata\(\)](<#PipelineRequest.ClearMetadata>)
  - [func \(x \*PipelineRequest\) ClearNamespace\(\)](<#PipelineRequest.ClearNamespace>)
  - [func \(x \*PipelineRequest\) ClearOperations\(\)](<#PipelineRequest.ClearOperations>)
  - [func \(x \*PipelineRequest\) GetMetadata\(\) \*common.RequestMetadata](<#PipelineRequest.GetMetadata>)
  - [func \(x \*PipelineRequest\) GetNamespace\(\) string](<#PipelineRequest.GetNamespace>)
  - [func \(x \*PipelineRequest\) GetOperations\(\) \[\]byte](<#PipelineRequest.GetOperations>)
  - [func \(x \*PipelineRequest\) HasMetadata\(\) bool](<#PipelineRequest.HasMetadata>)
  - [func \(x \*PipelineRequest\) HasNamespace\(\) bool](<#PipelineRequest.HasNamespace>)
  - [func \(x \*PipelineRequest\) HasOperations\(\) bool](<#PipelineRequest.HasOperations>)
  - [func \(\*PipelineRequest\) ProtoMessage\(\)](<#PipelineRequest.ProtoMessage>)
  - [func \(x \*PipelineRequest\) ProtoReflect\(\) protoreflect.Message](<#PipelineRequest.ProtoReflect>)
  - [func \(x \*PipelineRequest\) Reset\(\)](<#PipelineRequest.Reset>)
  - [func \(x \*PipelineRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#PipelineRequest.SetMetadata>)
  - [func \(x \*PipelineRequest\) SetNamespace\(v string\)](<#PipelineRequest.SetNamespace>)
  - [func \(x \*PipelineRequest\) SetOperations\(v \[\]byte\)](<#PipelineRequest.SetOperations>)
  - [func \(x \*PipelineRequest\) String\(\) string](<#PipelineRequest.String>)
- [type PipelineRequest\_builder](<#PipelineRequest_builder>)
  - [func \(b0 PipelineRequest\_builder\) Build\(\) \*PipelineRequest](<#PipelineRequest_builder.Build>)
- [type PoolStats](<#PoolStats>)
  - [func \(x \*PoolStats\) ClearAcquisitionFailures\(\)](<#PoolStats.ClearAcquisitionFailures>)
  - [func \(x \*PoolStats\) ClearAvgAcquisitionTime\(\)](<#PoolStats.ClearAvgAcquisitionTime>)
  - [func \(x \*PoolStats\) ClearTotalClosed\(\)](<#PoolStats.ClearTotalClosed>)
  - [func \(x \*PoolStats\) ClearTotalCreated\(\)](<#PoolStats.ClearTotalCreated>)
  - [func \(x \*PoolStats\) GetAcquisitionFailures\(\) int64](<#PoolStats.GetAcquisitionFailures>)
  - [func \(x \*PoolStats\) GetAvgAcquisitionTime\(\) \*durationpb.Duration](<#PoolStats.GetAvgAcquisitionTime>)
  - [func \(x \*PoolStats\) GetTotalClosed\(\) int64](<#PoolStats.GetTotalClosed>)
  - [func \(x \*PoolStats\) GetTotalCreated\(\) int64](<#PoolStats.GetTotalCreated>)
  - [func \(x \*PoolStats\) HasAcquisitionFailures\(\) bool](<#PoolStats.HasAcquisitionFailures>)
  - [func \(x \*PoolStats\) HasAvgAcquisitionTime\(\) bool](<#PoolStats.HasAvgAcquisitionTime>)
  - [func \(x \*PoolStats\) HasTotalClosed\(\) bool](<#PoolStats.HasTotalClosed>)
  - [func \(x \*PoolStats\) HasTotalCreated\(\) bool](<#PoolStats.HasTotalCreated>)
  - [func \(\*PoolStats\) ProtoMessage\(\)](<#PoolStats.ProtoMessage>)
  - [func \(x \*PoolStats\) ProtoReflect\(\) protoreflect.Message](<#PoolStats.ProtoReflect>)
  - [func \(x \*PoolStats\) Reset\(\)](<#PoolStats.Reset>)
  - [func \(x \*PoolStats\) SetAcquisitionFailures\(v int64\)](<#PoolStats.SetAcquisitionFailures>)
  - [func \(x \*PoolStats\) SetAvgAcquisitionTime\(v \*durationpb.Duration\)](<#PoolStats.SetAvgAcquisitionTime>)
  - [func \(x \*PoolStats\) SetTotalClosed\(v int64\)](<#PoolStats.SetTotalClosed>)
  - [func \(x \*PoolStats\) SetTotalCreated\(v int64\)](<#PoolStats.SetTotalCreated>)
  - [func \(x \*PoolStats\) String\(\) string](<#PoolStats.String>)
- [type PoolStats\_builder](<#PoolStats_builder>)
  - [func \(b0 PoolStats\_builder\) Build\(\) \*PoolStats](<#PoolStats_builder.Build>)
- [type PrependRequest](<#PrependRequest>)
  - [func \(x \*PrependRequest\) ClearKey\(\)](<#PrependRequest.ClearKey>)
  - [func \(x \*PrependRequest\) ClearMetadata\(\)](<#PrependRequest.ClearMetadata>)
  - [func \(x \*PrependRequest\) ClearNamespace\(\)](<#PrependRequest.ClearNamespace>)
  - [func \(x \*PrependRequest\) ClearValue\(\)](<#PrependRequest.ClearValue>)
  - [func \(x \*PrependRequest\) GetKey\(\) string](<#PrependRequest.GetKey>)
  - [func \(x \*PrependRequest\) GetMetadata\(\) \*common.RequestMetadata](<#PrependRequest.GetMetadata>)
  - [func \(x \*PrependRequest\) GetNamespace\(\) string](<#PrependRequest.GetNamespace>)
  - [func \(x \*PrependRequest\) GetValue\(\) \*anypb.Any](<#PrependRequest.GetValue>)
  - [func \(x \*PrependRequest\) HasKey\(\) bool](<#PrependRequest.HasKey>)
  - [func \(x \*PrependRequest\) HasMetadata\(\) bool](<#PrependRequest.HasMetadata>)
  - [func \(x \*PrependRequest\) HasNamespace\(\) bool](<#PrependRequest.HasNamespace>)
  - [func \(x \*PrependRequest\) HasValue\(\) bool](<#PrependRequest.HasValue>)
  - [func \(\*PrependRequest\) ProtoMessage\(\)](<#PrependRequest.ProtoMessage>)
  - [func \(x \*PrependRequest\) ProtoReflect\(\) protoreflect.Message](<#PrependRequest.ProtoReflect>)
  - [func \(x \*PrependRequest\) Reset\(\)](<#PrependRequest.Reset>)
  - [func \(x \*PrependRequest\) SetKey\(v string\)](<#PrependRequest.SetKey>)
  - [func \(x \*PrependRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#PrependRequest.SetMetadata>)
  - [func \(x \*PrependRequest\) SetNamespace\(v string\)](<#PrependRequest.SetNamespace>)
  - [func \(x \*PrependRequest\) SetValue\(v \*anypb.Any\)](<#PrependRequest.SetValue>)
  - [func \(x \*PrependRequest\) String\(\) string](<#PrependRequest.String>)
- [type PrependRequest\_builder](<#PrependRequest_builder>)
  - [func \(b0 PrependRequest\_builder\) Build\(\) \*PrependRequest](<#PrependRequest_builder.Build>)
- [type QueryOptions](<#QueryOptions>)
  - [func \(x \*QueryOptions\) ClearConsistency\(\)](<#QueryOptions.ClearConsistency>)
  - [func \(x \*QueryOptions\) ClearIncludeMetadata\(\)](<#QueryOptions.ClearIncludeMetadata>)
  - [func \(x \*QueryOptions\) ClearLimit\(\)](<#QueryOptions.ClearLimit>)
  - [func \(x \*QueryOptions\) ClearOffset\(\)](<#QueryOptions.ClearOffset>)
  - [func \(x \*QueryOptions\) ClearTimeout\(\)](<#QueryOptions.ClearTimeout>)
  - [func \(x \*QueryOptions\) GetConsistency\(\) common.DatabaseConsistencyLevel](<#QueryOptions.GetConsistency>)
  - [func \(x \*QueryOptions\) GetIncludeMetadata\(\) bool](<#QueryOptions.GetIncludeMetadata>)
  - [func \(x \*QueryOptions\) GetLimit\(\) int32](<#QueryOptions.GetLimit>)
  - [func \(x \*QueryOptions\) GetOffset\(\) int32](<#QueryOptions.GetOffset>)
  - [func \(x \*QueryOptions\) GetTimeout\(\) \*durationpb.Duration](<#QueryOptions.GetTimeout>)
  - [func \(x \*QueryOptions\) HasConsistency\(\) bool](<#QueryOptions.HasConsistency>)
  - [func \(x \*QueryOptions\) HasIncludeMetadata\(\) bool](<#QueryOptions.HasIncludeMetadata>)
  - [func \(x \*QueryOptions\) HasLimit\(\) bool](<#QueryOptions.HasLimit>)
  - [func \(x \*QueryOptions\) HasOffset\(\) bool](<#QueryOptions.HasOffset>)
  - [func \(x \*QueryOptions\) HasTimeout\(\) bool](<#QueryOptions.HasTimeout>)
  - [func \(\*QueryOptions\) ProtoMessage\(\)](<#QueryOptions.ProtoMessage>)
  - [func \(x \*QueryOptions\) ProtoReflect\(\) protoreflect.Message](<#QueryOptions.ProtoReflect>)
  - [func \(x \*QueryOptions\) Reset\(\)](<#QueryOptions.Reset>)
  - [func \(x \*QueryOptions\) SetConsistency\(v common.DatabaseConsistencyLevel\)](<#QueryOptions.SetConsistency>)
  - [func \(x \*QueryOptions\) SetIncludeMetadata\(v bool\)](<#QueryOptions.SetIncludeMetadata>)
  - [func \(x \*QueryOptions\) SetLimit\(v int32\)](<#QueryOptions.SetLimit>)
  - [func \(x \*QueryOptions\) SetOffset\(v int32\)](<#QueryOptions.SetOffset>)
  - [func \(x \*QueryOptions\) SetTimeout\(v \*durationpb.Duration\)](<#QueryOptions.SetTimeout>)
  - [func \(x \*QueryOptions\) String\(\) string](<#QueryOptions.String>)
- [type QueryOptions\_builder](<#QueryOptions_builder>)
  - [func \(b0 QueryOptions\_builder\) Build\(\) \*QueryOptions](<#QueryOptions_builder.Build>)
- [type QueryParameter](<#QueryParameter>)
  - [func \(x \*QueryParameter\) ClearName\(\)](<#QueryParameter.ClearName>)
  - [func \(x \*QueryParameter\) ClearType\(\)](<#QueryParameter.ClearType>)
  - [func \(x \*QueryParameter\) ClearValue\(\)](<#QueryParameter.ClearValue>)
  - [func \(x \*QueryParameter\) GetName\(\) string](<#QueryParameter.GetName>)
  - [func \(x \*QueryParameter\) GetType\(\) string](<#QueryParameter.GetType>)
  - [func \(x \*QueryParameter\) GetValue\(\) \*anypb.Any](<#QueryParameter.GetValue>)
  - [func \(x \*QueryParameter\) HasName\(\) bool](<#QueryParameter.HasName>)
  - [func \(x \*QueryParameter\) HasType\(\) bool](<#QueryParameter.HasType>)
  - [func \(x \*QueryParameter\) HasValue\(\) bool](<#QueryParameter.HasValue>)
  - [func \(\*QueryParameter\) ProtoMessage\(\)](<#QueryParameter.ProtoMessage>)
  - [func \(x \*QueryParameter\) ProtoReflect\(\) protoreflect.Message](<#QueryParameter.ProtoReflect>)
  - [func \(x \*QueryParameter\) Reset\(\)](<#QueryParameter.Reset>)
  - [func \(x \*QueryParameter\) SetName\(v string\)](<#QueryParameter.SetName>)
  - [func \(x \*QueryParameter\) SetType\(v string\)](<#QueryParameter.SetType>)
  - [func \(x \*QueryParameter\) SetValue\(v \*anypb.Any\)](<#QueryParameter.SetValue>)
  - [func \(x \*QueryParameter\) String\(\) string](<#QueryParameter.String>)
- [type QueryParameter\_builder](<#QueryParameter_builder>)
  - [func \(b0 QueryParameter\_builder\) Build\(\) \*QueryParameter](<#QueryParameter_builder.Build>)
- [type QueryRequest](<#QueryRequest>)
  - [func \(x \*QueryRequest\) ClearDatabase\(\)](<#QueryRequest.ClearDatabase>)
  - [func \(x \*QueryRequest\) ClearMetadata\(\)](<#QueryRequest.ClearMetadata>)
  - [func \(x \*QueryRequest\) ClearOptions\(\)](<#QueryRequest.ClearOptions>)
  - [func \(x \*QueryRequest\) ClearQuery\(\)](<#QueryRequest.ClearQuery>)
  - [func \(x \*QueryRequest\) ClearTransactionId\(\)](<#QueryRequest.ClearTransactionId>)
  - [func \(x \*QueryRequest\) GetDatabase\(\) string](<#QueryRequest.GetDatabase>)
  - [func \(x \*QueryRequest\) GetMetadata\(\) \*common.RequestMetadata](<#QueryRequest.GetMetadata>)
  - [func \(x \*QueryRequest\) GetOptions\(\) \*QueryOptions](<#QueryRequest.GetOptions>)
  - [func \(x \*QueryRequest\) GetParameters\(\) \[\]\*QueryParameter](<#QueryRequest.GetParameters>)
  - [func \(x \*QueryRequest\) GetQuery\(\) string](<#QueryRequest.GetQuery>)
  - [func \(x \*QueryRequest\) GetTransactionId\(\) string](<#QueryRequest.GetTransactionId>)
  - [func \(x \*QueryRequest\) HasDatabase\(\) bool](<#QueryRequest.HasDatabase>)
  - [func \(x \*QueryRequest\) HasMetadata\(\) bool](<#QueryRequest.HasMetadata>)
  - [func \(x \*QueryRequest\) HasOptions\(\) bool](<#QueryRequest.HasOptions>)
  - [func \(x \*QueryRequest\) HasQuery\(\) bool](<#QueryRequest.HasQuery>)
  - [func \(x \*QueryRequest\) HasTransactionId\(\) bool](<#QueryRequest.HasTransactionId>)
  - [func \(\*QueryRequest\) ProtoMessage\(\)](<#QueryRequest.ProtoMessage>)
  - [func \(x \*QueryRequest\) ProtoReflect\(\) protoreflect.Message](<#QueryRequest.ProtoReflect>)
  - [func \(x \*QueryRequest\) Reset\(\)](<#QueryRequest.Reset>)
  - [func \(x \*QueryRequest\) SetDatabase\(v string\)](<#QueryRequest.SetDatabase>)
  - [func \(x \*QueryRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#QueryRequest.SetMetadata>)
  - [func \(x \*QueryRequest\) SetOptions\(v \*QueryOptions\)](<#QueryRequest.SetOptions>)
  - [func \(x \*QueryRequest\) SetParameters\(v \[\]\*QueryParameter\)](<#QueryRequest.SetParameters>)
  - [func \(x \*QueryRequest\) SetQuery\(v string\)](<#QueryRequest.SetQuery>)
  - [func \(x \*QueryRequest\) SetTransactionId\(v string\)](<#QueryRequest.SetTransactionId>)
  - [func \(x \*QueryRequest\) String\(\) string](<#QueryRequest.String>)
- [type QueryRequest\_builder](<#QueryRequest_builder>)
  - [func \(b0 QueryRequest\_builder\) Build\(\) \*QueryRequest](<#QueryRequest_builder.Build>)
- [type QueryResponse](<#QueryResponse>)
  - [func \(x \*QueryResponse\) ClearError\(\)](<#QueryResponse.ClearError>)
  - [func \(x \*QueryResponse\) ClearResultSet\(\)](<#QueryResponse.ClearResultSet>)
  - [func \(x \*QueryResponse\) ClearStats\(\)](<#QueryResponse.ClearStats>)
  - [func \(x \*QueryResponse\) GetError\(\) \*common.Error](<#QueryResponse.GetError>)
  - [func \(x \*QueryResponse\) GetResultSet\(\) \*ResultSet](<#QueryResponse.GetResultSet>)
  - [func \(x \*QueryResponse\) GetStats\(\) \*DatabaseQueryStats](<#QueryResponse.GetStats>)
  - [func \(x \*QueryResponse\) HasError\(\) bool](<#QueryResponse.HasError>)
  - [func \(x \*QueryResponse\) HasResultSet\(\) bool](<#QueryResponse.HasResultSet>)
  - [func \(x \*QueryResponse\) HasStats\(\) bool](<#QueryResponse.HasStats>)
  - [func \(\*QueryResponse\) ProtoMessage\(\)](<#QueryResponse.ProtoMessage>)
  - [func \(x \*QueryResponse\) ProtoReflect\(\) protoreflect.Message](<#QueryResponse.ProtoReflect>)
  - [func \(x \*QueryResponse\) Reset\(\)](<#QueryResponse.Reset>)
  - [func \(x \*QueryResponse\) SetError\(v \*common.Error\)](<#QueryResponse.SetError>)
  - [func \(x \*QueryResponse\) SetResultSet\(v \*ResultSet\)](<#QueryResponse.SetResultSet>)
  - [func \(x \*QueryResponse\) SetStats\(v \*DatabaseQueryStats\)](<#QueryResponse.SetStats>)
  - [func \(x \*QueryResponse\) String\(\) string](<#QueryResponse.String>)
- [type QueryResponse\_builder](<#QueryResponse_builder>)
  - [func \(b0 QueryResponse\_builder\) Build\(\) \*QueryResponse](<#QueryResponse_builder.Build>)
- [type QueryRowRequest](<#QueryRowRequest>)
  - [func \(x \*QueryRowRequest\) ClearDatabase\(\)](<#QueryRowRequest.ClearDatabase>)
  - [func \(x \*QueryRowRequest\) ClearMetadata\(\)](<#QueryRowRequest.ClearMetadata>)
  - [func \(x \*QueryRowRequest\) ClearOptions\(\)](<#QueryRowRequest.ClearOptions>)
  - [func \(x \*QueryRowRequest\) ClearQuery\(\)](<#QueryRowRequest.ClearQuery>)
  - [func \(x \*QueryRowRequest\) ClearTransactionId\(\)](<#QueryRowRequest.ClearTransactionId>)
  - [func \(x \*QueryRowRequest\) GetDatabase\(\) string](<#QueryRowRequest.GetDatabase>)
  - [func \(x \*QueryRowRequest\) GetMetadata\(\) \*common.RequestMetadata](<#QueryRowRequest.GetMetadata>)
  - [func \(x \*QueryRowRequest\) GetOptions\(\) \*QueryOptions](<#QueryRowRequest.GetOptions>)
  - [func \(x \*QueryRowRequest\) GetParameters\(\) \[\]\*QueryParameter](<#QueryRowRequest.GetParameters>)
  - [func \(x \*QueryRowRequest\) GetQuery\(\) string](<#QueryRowRequest.GetQuery>)
  - [func \(x \*QueryRowRequest\) GetTransactionId\(\) string](<#QueryRowRequest.GetTransactionId>)
  - [func \(x \*QueryRowRequest\) HasDatabase\(\) bool](<#QueryRowRequest.HasDatabase>)
  - [func \(x \*QueryRowRequest\) HasMetadata\(\) bool](<#QueryRowRequest.HasMetadata>)
  - [func \(x \*QueryRowRequest\) HasOptions\(\) bool](<#QueryRowRequest.HasOptions>)
  - [func \(x \*QueryRowRequest\) HasQuery\(\) bool](<#QueryRowRequest.HasQuery>)
  - [func \(x \*QueryRowRequest\) HasTransactionId\(\) bool](<#QueryRowRequest.HasTransactionId>)
  - [func \(\*QueryRowRequest\) ProtoMessage\(\)](<#QueryRowRequest.ProtoMessage>)
  - [func \(x \*QueryRowRequest\) ProtoReflect\(\) protoreflect.Message](<#QueryRowRequest.ProtoReflect>)
  - [func \(x \*QueryRowRequest\) Reset\(\)](<#QueryRowRequest.Reset>)
  - [func \(x \*QueryRowRequest\) SetDatabase\(v string\)](<#QueryRowRequest.SetDatabase>)
  - [func \(x \*QueryRowRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#QueryRowRequest.SetMetadata>)
  - [func \(x \*QueryRowRequest\) SetOptions\(v \*QueryOptions\)](<#QueryRowRequest.SetOptions>)
  - [func \(x \*QueryRowRequest\) SetParameters\(v \[\]\*QueryParameter\)](<#QueryRowRequest.SetParameters>)
  - [func \(x \*QueryRowRequest\) SetQuery\(v string\)](<#QueryRowRequest.SetQuery>)
  - [func \(x \*QueryRowRequest\) SetTransactionId\(v string\)](<#QueryRowRequest.SetTransactionId>)
  - [func \(x \*QueryRowRequest\) String\(\) string](<#QueryRowRequest.String>)
- [type QueryRowRequest\_builder](<#QueryRowRequest_builder>)
  - [func \(b0 QueryRowRequest\_builder\) Build\(\) \*QueryRowRequest](<#QueryRowRequest_builder.Build>)
- [type QueryRowResponse](<#QueryRowResponse>)
  - [func \(x \*QueryRowResponse\) ClearError\(\)](<#QueryRowResponse.ClearError>)
  - [func \(x \*QueryRowResponse\) ClearFound\(\)](<#QueryRowResponse.ClearFound>)
  - [func \(x \*QueryRowResponse\) ClearStats\(\)](<#QueryRowResponse.ClearStats>)
  - [func \(x \*QueryRowResponse\) GetColumns\(\) \[\]string](<#QueryRowResponse.GetColumns>)
  - [func \(x \*QueryRowResponse\) GetError\(\) \*common.Error](<#QueryRowResponse.GetError>)
  - [func \(x \*QueryRowResponse\) GetFound\(\) bool](<#QueryRowResponse.GetFound>)
  - [func \(x \*QueryRowResponse\) GetStats\(\) \*DatabaseQueryStats](<#QueryRowResponse.GetStats>)
  - [func \(x \*QueryRowResponse\) GetValues\(\) \[\]\*anypb.Any](<#QueryRowResponse.GetValues>)
  - [func \(x \*QueryRowResponse\) HasError\(\) bool](<#QueryRowResponse.HasError>)
  - [func \(x \*QueryRowResponse\) HasFound\(\) bool](<#QueryRowResponse.HasFound>)
  - [func \(x \*QueryRowResponse\) HasStats\(\) bool](<#QueryRowResponse.HasStats>)
  - [func \(\*QueryRowResponse\) ProtoMessage\(\)](<#QueryRowResponse.ProtoMessage>)
  - [func \(x \*QueryRowResponse\) ProtoReflect\(\) protoreflect.Message](<#QueryRowResponse.ProtoReflect>)
  - [func \(x \*QueryRowResponse\) Reset\(\)](<#QueryRowResponse.Reset>)
  - [func \(x \*QueryRowResponse\) SetColumns\(v \[\]string\)](<#QueryRowResponse.SetColumns>)
  - [func \(x \*QueryRowResponse\) SetError\(v \*common.Error\)](<#QueryRowResponse.SetError>)
  - [func \(x \*QueryRowResponse\) SetFound\(v bool\)](<#QueryRowResponse.SetFound>)
  - [func \(x \*QueryRowResponse\) SetStats\(v \*DatabaseQueryStats\)](<#QueryRowResponse.SetStats>)
  - [func \(x \*QueryRowResponse\) SetValues\(v \[\]\*anypb.Any\)](<#QueryRowResponse.SetValues>)
  - [func \(x \*QueryRowResponse\) String\(\) string](<#QueryRowResponse.String>)
- [type QueryRowResponse\_builder](<#QueryRowResponse_builder>)
  - [func \(b0 QueryRowResponse\_builder\) Build\(\) \*QueryRowResponse](<#QueryRowResponse_builder.Build>)
- [type RestoreRequest](<#RestoreRequest>)
  - [func \(x \*RestoreRequest\) ClearMetadata\(\)](<#RestoreRequest.ClearMetadata>)
  - [func \(x \*RestoreRequest\) ClearNamespace\(\)](<#RestoreRequest.ClearNamespace>)
  - [func \(x \*RestoreRequest\) ClearSource\(\)](<#RestoreRequest.ClearSource>)
  - [func \(x \*RestoreRequest\) GetMetadata\(\) \*common.RequestMetadata](<#RestoreRequest.GetMetadata>)
  - [func \(x \*RestoreRequest\) GetNamespace\(\) string](<#RestoreRequest.GetNamespace>)
  - [func \(x \*RestoreRequest\) GetSource\(\) string](<#RestoreRequest.GetSource>)
  - [func \(x \*RestoreRequest\) HasMetadata\(\) bool](<#RestoreRequest.HasMetadata>)
  - [func \(x \*RestoreRequest\) HasNamespace\(\) bool](<#RestoreRequest.HasNamespace>)
  - [func \(x \*RestoreRequest\) HasSource\(\) bool](<#RestoreRequest.HasSource>)
  - [func \(\*RestoreRequest\) ProtoMessage\(\)](<#RestoreRequest.ProtoMessage>)
  - [func \(x \*RestoreRequest\) ProtoReflect\(\) protoreflect.Message](<#RestoreRequest.ProtoReflect>)
  - [func \(x \*RestoreRequest\) Reset\(\)](<#RestoreRequest.Reset>)
  - [func \(x \*RestoreRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#RestoreRequest.SetMetadata>)
  - [func \(x \*RestoreRequest\) SetNamespace\(v string\)](<#RestoreRequest.SetNamespace>)
  - [func \(x \*RestoreRequest\) SetSource\(v string\)](<#RestoreRequest.SetSource>)
  - [func \(x \*RestoreRequest\) String\(\) string](<#RestoreRequest.String>)
- [type RestoreRequest\_builder](<#RestoreRequest_builder>)
  - [func \(b0 RestoreRequest\_builder\) Build\(\) \*RestoreRequest](<#RestoreRequest_builder.Build>)
- [type ResultSet](<#ResultSet>)
  - [func \(x \*ResultSet\) ClearHasMore\(\)](<#ResultSet.ClearHasMore>)
  - [func \(x \*ResultSet\) ClearTotalCount\(\)](<#ResultSet.ClearTotalCount>)
  - [func \(x \*ResultSet\) GetColumns\(\) \[\]\*ColumnMetadata](<#ResultSet.GetColumns>)
  - [func \(x \*ResultSet\) GetHasMore\(\) bool](<#ResultSet.GetHasMore>)
  - [func \(x \*ResultSet\) GetRows\(\) \[\]\*Row](<#ResultSet.GetRows>)
  - [func \(x \*ResultSet\) GetTotalCount\(\) int64](<#ResultSet.GetTotalCount>)
  - [func \(x \*ResultSet\) HasHasMore\(\) bool](<#ResultSet.HasHasMore>)
  - [func \(x \*ResultSet\) HasTotalCount\(\) bool](<#ResultSet.HasTotalCount>)
  - [func \(\*ResultSet\) ProtoMessage\(\)](<#ResultSet.ProtoMessage>)
  - [func \(x \*ResultSet\) ProtoReflect\(\) protoreflect.Message](<#ResultSet.ProtoReflect>)
  - [func \(x \*ResultSet\) Reset\(\)](<#ResultSet.Reset>)
  - [func \(x \*ResultSet\) SetColumns\(v \[\]\*ColumnMetadata\)](<#ResultSet.SetColumns>)
  - [func \(x \*ResultSet\) SetHasMore\(v bool\)](<#ResultSet.SetHasMore>)
  - [func \(x \*ResultSet\) SetRows\(v \[\]\*Row\)](<#ResultSet.SetRows>)
  - [func \(x \*ResultSet\) SetTotalCount\(v int64\)](<#ResultSet.SetTotalCount>)
  - [func \(x \*ResultSet\) String\(\) string](<#ResultSet.String>)
- [type ResultSet\_builder](<#ResultSet_builder>)
  - [func \(b0 ResultSet\_builder\) Build\(\) \*ResultSet](<#ResultSet_builder.Build>)
- [type RevertMigrationRequest](<#RevertMigrationRequest>)
  - [func \(x \*RevertMigrationRequest\) ClearDatabase\(\)](<#RevertMigrationRequest.ClearDatabase>)
  - [func \(x \*RevertMigrationRequest\) ClearMetadata\(\)](<#RevertMigrationRequest.ClearMetadata>)
  - [func \(x \*RevertMigrationRequest\) ClearTargetVersion\(\)](<#RevertMigrationRequest.ClearTargetVersion>)
  - [func \(x \*RevertMigrationRequest\) GetDatabase\(\) string](<#RevertMigrationRequest.GetDatabase>)
  - [func \(x \*RevertMigrationRequest\) GetMetadata\(\) \*common.RequestMetadata](<#RevertMigrationRequest.GetMetadata>)
  - [func \(x \*RevertMigrationRequest\) GetTargetVersion\(\) string](<#RevertMigrationRequest.GetTargetVersion>)
  - [func \(x \*RevertMigrationRequest\) HasDatabase\(\) bool](<#RevertMigrationRequest.HasDatabase>)
  - [func \(x \*RevertMigrationRequest\) HasMetadata\(\) bool](<#RevertMigrationRequest.HasMetadata>)
  - [func \(x \*RevertMigrationRequest\) HasTargetVersion\(\) bool](<#RevertMigrationRequest.HasTargetVersion>)
  - [func \(\*RevertMigrationRequest\) ProtoMessage\(\)](<#RevertMigrationRequest.ProtoMessage>)
  - [func \(x \*RevertMigrationRequest\) ProtoReflect\(\) protoreflect.Message](<#RevertMigrationRequest.ProtoReflect>)
  - [func \(x \*RevertMigrationRequest\) Reset\(\)](<#RevertMigrationRequest.Reset>)
  - [func \(x \*RevertMigrationRequest\) SetDatabase\(v string\)](<#RevertMigrationRequest.SetDatabase>)
  - [func \(x \*RevertMigrationRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#RevertMigrationRequest.SetMetadata>)
  - [func \(x \*RevertMigrationRequest\) SetTargetVersion\(v string\)](<#RevertMigrationRequest.SetTargetVersion>)
  - [func \(x \*RevertMigrationRequest\) String\(\) string](<#RevertMigrationRequest.String>)
- [type RevertMigrationRequest\_builder](<#RevertMigrationRequest_builder>)
  - [func \(b0 RevertMigrationRequest\_builder\) Build\(\) \*RevertMigrationRequest](<#RevertMigrationRequest_builder.Build>)
- [type RevertMigrationResponse](<#RevertMigrationResponse>)
  - [func \(x \*RevertMigrationResponse\) ClearError\(\)](<#RevertMigrationResponse.ClearError>)
  - [func \(x \*RevertMigrationResponse\) ClearRevertedTo\(\)](<#RevertMigrationResponse.ClearRevertedTo>)
  - [func \(x \*RevertMigrationResponse\) ClearSuccess\(\)](<#RevertMigrationResponse.ClearSuccess>)
  - [func \(x \*RevertMigrationResponse\) GetError\(\) \*common.Error](<#RevertMigrationResponse.GetError>)
  - [func \(x \*RevertMigrationResponse\) GetRevertedTo\(\) string](<#RevertMigrationResponse.GetRevertedTo>)
  - [func \(x \*RevertMigrationResponse\) GetSuccess\(\) bool](<#RevertMigrationResponse.GetSuccess>)
  - [func \(x \*RevertMigrationResponse\) HasError\(\) bool](<#RevertMigrationResponse.HasError>)
  - [func \(x \*RevertMigrationResponse\) HasRevertedTo\(\) bool](<#RevertMigrationResponse.HasRevertedTo>)
  - [func \(x \*RevertMigrationResponse\) HasSuccess\(\) bool](<#RevertMigrationResponse.HasSuccess>)
  - [func \(\*RevertMigrationResponse\) ProtoMessage\(\)](<#RevertMigrationResponse.ProtoMessage>)
  - [func \(x \*RevertMigrationResponse\) ProtoReflect\(\) protoreflect.Message](<#RevertMigrationResponse.ProtoReflect>)
  - [func \(x \*RevertMigrationResponse\) Reset\(\)](<#RevertMigrationResponse.Reset>)
  - [func \(x \*RevertMigrationResponse\) SetError\(v \*common.Error\)](<#RevertMigrationResponse.SetError>)
  - [func \(x \*RevertMigrationResponse\) SetRevertedTo\(v string\)](<#RevertMigrationResponse.SetRevertedTo>)
  - [func \(x \*RevertMigrationResponse\) SetSuccess\(v bool\)](<#RevertMigrationResponse.SetSuccess>)
  - [func \(x \*RevertMigrationResponse\) String\(\) string](<#RevertMigrationResponse.String>)
- [type RevertMigrationResponse\_builder](<#RevertMigrationResponse_builder>)
  - [func \(b0 RevertMigrationResponse\_builder\) Build\(\) \*RevertMigrationResponse](<#RevertMigrationResponse_builder.Build>)
- [type RollbackTransactionRequest](<#RollbackTransactionRequest>)
  - [func \(x \*RollbackTransactionRequest\) ClearMetadata\(\)](<#RollbackTransactionRequest.ClearMetadata>)
  - [func \(x \*RollbackTransactionRequest\) ClearTransactionId\(\)](<#RollbackTransactionRequest.ClearTransactionId>)
  - [func \(x \*RollbackTransactionRequest\) GetMetadata\(\) \*common.RequestMetadata](<#RollbackTransactionRequest.GetMetadata>)
  - [func \(x \*RollbackTransactionRequest\) GetTransactionId\(\) string](<#RollbackTransactionRequest.GetTransactionId>)
  - [func \(x \*RollbackTransactionRequest\) HasMetadata\(\) bool](<#RollbackTransactionRequest.HasMetadata>)
  - [func \(x \*RollbackTransactionRequest\) HasTransactionId\(\) bool](<#RollbackTransactionRequest.HasTransactionId>)
  - [func \(\*RollbackTransactionRequest\) ProtoMessage\(\)](<#RollbackTransactionRequest.ProtoMessage>)
  - [func \(x \*RollbackTransactionRequest\) ProtoReflect\(\) protoreflect.Message](<#RollbackTransactionRequest.ProtoReflect>)
  - [func \(x \*RollbackTransactionRequest\) Reset\(\)](<#RollbackTransactionRequest.Reset>)
  - [func \(x \*RollbackTransactionRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#RollbackTransactionRequest.SetMetadata>)
  - [func \(x \*RollbackTransactionRequest\) SetTransactionId\(v string\)](<#RollbackTransactionRequest.SetTransactionId>)
  - [func \(x \*RollbackTransactionRequest\) String\(\) string](<#RollbackTransactionRequest.String>)
- [type RollbackTransactionRequest\_builder](<#RollbackTransactionRequest_builder>)
  - [func \(b0 RollbackTransactionRequest\_builder\) Build\(\) \*RollbackTransactionRequest](<#RollbackTransactionRequest_builder.Build>)
- [type Row](<#Row>)
  - [func \(x \*Row\) GetValues\(\) \[\]\*anypb.Any](<#Row.GetValues>)
  - [func \(\*Row\) ProtoMessage\(\)](<#Row.ProtoMessage>)
  - [func \(x \*Row\) ProtoReflect\(\) protoreflect.Message](<#Row.ProtoReflect>)
  - [func \(x \*Row\) Reset\(\)](<#Row.Reset>)
  - [func \(x \*Row\) SetValues\(v \[\]\*anypb.Any\)](<#Row.SetValues>)
  - [func \(x \*Row\) String\(\) string](<#Row.String>)
- [type Row\_builder](<#Row_builder>)
  - [func \(b0 Row\_builder\) Build\(\) \*Row](<#Row_builder.Build>)
- [type RunMigrationRequest](<#RunMigrationRequest>)
  - [func \(x \*RunMigrationRequest\) ClearDatabase\(\)](<#RunMigrationRequest.ClearDatabase>)
  - [func \(x \*RunMigrationRequest\) ClearMetadata\(\)](<#RunMigrationRequest.ClearMetadata>)
  - [func \(x \*RunMigrationRequest\) GetDatabase\(\) string](<#RunMigrationRequest.GetDatabase>)
  - [func \(x \*RunMigrationRequest\) GetMetadata\(\) \*common.RequestMetadata](<#RunMigrationRequest.GetMetadata>)
  - [func \(x \*RunMigrationRequest\) GetScripts\(\) \[\]\*MigrationScript](<#RunMigrationRequest.GetScripts>)
  - [func \(x \*RunMigrationRequest\) HasDatabase\(\) bool](<#RunMigrationRequest.HasDatabase>)
  - [func \(x \*RunMigrationRequest\) HasMetadata\(\) bool](<#RunMigrationRequest.HasMetadata>)
  - [func \(\*RunMigrationRequest\) ProtoMessage\(\)](<#RunMigrationRequest.ProtoMessage>)
  - [func \(x \*RunMigrationRequest\) ProtoReflect\(\) protoreflect.Message](<#RunMigrationRequest.ProtoReflect>)
  - [func \(x \*RunMigrationRequest\) Reset\(\)](<#RunMigrationRequest.Reset>)
  - [func \(x \*RunMigrationRequest\) SetDatabase\(v string\)](<#RunMigrationRequest.SetDatabase>)
  - [func \(x \*RunMigrationRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#RunMigrationRequest.SetMetadata>)
  - [func \(x \*RunMigrationRequest\) SetScripts\(v \[\]\*MigrationScript\)](<#RunMigrationRequest.SetScripts>)
  - [func \(x \*RunMigrationRequest\) String\(\) string](<#RunMigrationRequest.String>)
- [type RunMigrationRequest\_builder](<#RunMigrationRequest_builder>)
  - [func \(b0 RunMigrationRequest\_builder\) Build\(\) \*RunMigrationRequest](<#RunMigrationRequest_builder.Build>)
- [type RunMigrationResponse](<#RunMigrationResponse>)
  - [func \(x \*RunMigrationResponse\) ClearError\(\)](<#RunMigrationResponse.ClearError>)
  - [func \(x \*RunMigrationResponse\) ClearSuccess\(\)](<#RunMigrationResponse.ClearSuccess>)
  - [func \(x \*RunMigrationResponse\) GetAppliedVersions\(\) \[\]string](<#RunMigrationResponse.GetAppliedVersions>)
  - [func \(x \*RunMigrationResponse\) GetError\(\) \*common.Error](<#RunMigrationResponse.GetError>)
  - [func \(x \*RunMigrationResponse\) GetSuccess\(\) bool](<#RunMigrationResponse.GetSuccess>)
  - [func \(x \*RunMigrationResponse\) HasError\(\) bool](<#RunMigrationResponse.HasError>)
  - [func \(x \*RunMigrationResponse\) HasSuccess\(\) bool](<#RunMigrationResponse.HasSuccess>)
  - [func \(\*RunMigrationResponse\) ProtoMessage\(\)](<#RunMigrationResponse.ProtoMessage>)
  - [func \(x \*RunMigrationResponse\) ProtoReflect\(\) protoreflect.Message](<#RunMigrationResponse.ProtoReflect>)
  - [func \(x \*RunMigrationResponse\) Reset\(\)](<#RunMigrationResponse.Reset>)
  - [func \(x \*RunMigrationResponse\) SetAppliedVersions\(v \[\]string\)](<#RunMigrationResponse.SetAppliedVersions>)
  - [func \(x \*RunMigrationResponse\) SetError\(v \*common.Error\)](<#RunMigrationResponse.SetError>)
  - [func \(x \*RunMigrationResponse\) SetSuccess\(v bool\)](<#RunMigrationResponse.SetSuccess>)
  - [func \(x \*RunMigrationResponse\) String\(\) string](<#RunMigrationResponse.String>)
- [type RunMigrationResponse\_builder](<#RunMigrationResponse_builder>)
  - [func \(b0 RunMigrationResponse\_builder\) Build\(\) \*RunMigrationResponse](<#RunMigrationResponse_builder.Build>)
- [type ScanRequest](<#ScanRequest>)
  - [func \(x \*ScanRequest\) ClearCursor\(\)](<#ScanRequest.ClearCursor>)
  - [func \(x \*ScanRequest\) ClearMetadata\(\)](<#ScanRequest.ClearMetadata>)
  - [func \(x \*ScanRequest\) ClearNamespace\(\)](<#ScanRequest.ClearNamespace>)
  - [func \(x \*ScanRequest\) ClearPattern\(\)](<#ScanRequest.ClearPattern>)
  - [func \(x \*ScanRequest\) GetCursor\(\) string](<#ScanRequest.GetCursor>)
  - [func \(x \*ScanRequest\) GetMetadata\(\) \*common.RequestMetadata](<#ScanRequest.GetMetadata>)
  - [func \(x \*ScanRequest\) GetNamespace\(\) string](<#ScanRequest.GetNamespace>)
  - [func \(x \*ScanRequest\) GetPattern\(\) string](<#ScanRequest.GetPattern>)
  - [func \(x \*ScanRequest\) HasCursor\(\) bool](<#ScanRequest.HasCursor>)
  - [func \(x \*ScanRequest\) HasMetadata\(\) bool](<#ScanRequest.HasMetadata>)
  - [func \(x \*ScanRequest\) HasNamespace\(\) bool](<#ScanRequest.HasNamespace>)
  - [func \(x \*ScanRequest\) HasPattern\(\) bool](<#ScanRequest.HasPattern>)
  - [func \(\*ScanRequest\) ProtoMessage\(\)](<#ScanRequest.ProtoMessage>)
  - [func \(x \*ScanRequest\) ProtoReflect\(\) protoreflect.Message](<#ScanRequest.ProtoReflect>)
  - [func \(x \*ScanRequest\) Reset\(\)](<#ScanRequest.Reset>)
  - [func \(x \*ScanRequest\) SetCursor\(v string\)](<#ScanRequest.SetCursor>)
  - [func \(x \*ScanRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#ScanRequest.SetMetadata>)
  - [func \(x \*ScanRequest\) SetNamespace\(v string\)](<#ScanRequest.SetNamespace>)
  - [func \(x \*ScanRequest\) SetPattern\(v string\)](<#ScanRequest.SetPattern>)
  - [func \(x \*ScanRequest\) String\(\) string](<#ScanRequest.String>)
- [type ScanRequest\_builder](<#ScanRequest_builder>)
  - [func \(b0 ScanRequest\_builder\) Build\(\) \*ScanRequest](<#ScanRequest_builder.Build>)
- [type SetMultipleRequest](<#SetMultipleRequest>)
  - [func \(x \*SetMultipleRequest\) ClearMetadata\(\)](<#SetMultipleRequest.ClearMetadata>)
  - [func \(x \*SetMultipleRequest\) ClearTtl\(\)](<#SetMultipleRequest.ClearTtl>)
  - [func \(x \*SetMultipleRequest\) GetMetadata\(\) \*common.RequestMetadata](<#SetMultipleRequest.GetMetadata>)
  - [func \(x \*SetMultipleRequest\) GetTtl\(\) \*durationpb.Duration](<#SetMultipleRequest.GetTtl>)
  - [func \(x \*SetMultipleRequest\) GetValues\(\) map\[string\]\[\]byte](<#SetMultipleRequest.GetValues>)
  - [func \(x \*SetMultipleRequest\) HasMetadata\(\) bool](<#SetMultipleRequest.HasMetadata>)
  - [func \(x \*SetMultipleRequest\) HasTtl\(\) bool](<#SetMultipleRequest.HasTtl>)
  - [func \(\*SetMultipleRequest\) ProtoMessage\(\)](<#SetMultipleRequest.ProtoMessage>)
  - [func \(x \*SetMultipleRequest\) ProtoReflect\(\) protoreflect.Message](<#SetMultipleRequest.ProtoReflect>)
  - [func \(x \*SetMultipleRequest\) Reset\(\)](<#SetMultipleRequest.Reset>)
  - [func \(x \*SetMultipleRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#SetMultipleRequest.SetMetadata>)
  - [func \(x \*SetMultipleRequest\) SetTtl\(v \*durationpb.Duration\)](<#SetMultipleRequest.SetTtl>)
  - [func \(x \*SetMultipleRequest\) SetValues\(v map\[string\]\[\]byte\)](<#SetMultipleRequest.SetValues>)
  - [func \(x \*SetMultipleRequest\) String\(\) string](<#SetMultipleRequest.String>)
- [type SetMultipleRequest\_builder](<#SetMultipleRequest_builder>)
  - [func \(b0 SetMultipleRequest\_builder\) Build\(\) \*SetMultipleRequest](<#SetMultipleRequest_builder.Build>)
- [type SetMultipleResponse](<#SetMultipleResponse>)
  - [func \(x \*SetMultipleResponse\) ClearError\(\)](<#SetMultipleResponse.ClearError>)
  - [func \(x \*SetMultipleResponse\) ClearSetCount\(\)](<#SetMultipleResponse.ClearSetCount>)
  - [func \(x \*SetMultipleResponse\) ClearSuccess\(\)](<#SetMultipleResponse.ClearSuccess>)
  - [func \(x \*SetMultipleResponse\) GetError\(\) \*common.Error](<#SetMultipleResponse.GetError>)
  - [func \(x \*SetMultipleResponse\) GetFailedKeys\(\) \[\]string](<#SetMultipleResponse.GetFailedKeys>)
  - [func \(x \*SetMultipleResponse\) GetSetCount\(\) int32](<#SetMultipleResponse.GetSetCount>)
  - [func \(x \*SetMultipleResponse\) GetSuccess\(\) bool](<#SetMultipleResponse.GetSuccess>)
  - [func \(x \*SetMultipleResponse\) HasError\(\) bool](<#SetMultipleResponse.HasError>)
  - [func \(x \*SetMultipleResponse\) HasSetCount\(\) bool](<#SetMultipleResponse.HasSetCount>)
  - [func \(x \*SetMultipleResponse\) HasSuccess\(\) bool](<#SetMultipleResponse.HasSuccess>)
  - [func \(\*SetMultipleResponse\) ProtoMessage\(\)](<#SetMultipleResponse.ProtoMessage>)
  - [func \(x \*SetMultipleResponse\) ProtoReflect\(\) protoreflect.Message](<#SetMultipleResponse.ProtoReflect>)
  - [func \(x \*SetMultipleResponse\) Reset\(\)](<#SetMultipleResponse.Reset>)
  - [func \(x \*SetMultipleResponse\) SetError\(v \*common.Error\)](<#SetMultipleResponse.SetError>)
  - [func \(x \*SetMultipleResponse\) SetFailedKeys\(v \[\]string\)](<#SetMultipleResponse.SetFailedKeys>)
  - [func \(x \*SetMultipleResponse\) SetSetCount\(v int32\)](<#SetMultipleResponse.SetSetCount>)
  - [func \(x \*SetMultipleResponse\) SetSuccess\(v bool\)](<#SetMultipleResponse.SetSuccess>)
  - [func \(x \*SetMultipleResponse\) String\(\) string](<#SetMultipleResponse.String>)
- [type SetMultipleResponse\_builder](<#SetMultipleResponse_builder>)
  - [func \(b0 SetMultipleResponse\_builder\) Build\(\) \*SetMultipleResponse](<#SetMultipleResponse_builder.Build>)
- [type SetOptions](<#SetOptions>)
  - [func \(x \*SetOptions\) ClearOnlyIfAbsent\(\)](<#SetOptions.ClearOnlyIfAbsent>)
  - [func \(x \*SetOptions\) ClearOnlyIfPresent\(\)](<#SetOptions.ClearOnlyIfPresent>)
  - [func \(x \*SetOptions\) ClearReturnPrevious\(\)](<#SetOptions.ClearReturnPrevious>)
  - [func \(x \*SetOptions\) ClearTtl\(\)](<#SetOptions.ClearTtl>)
  - [func \(x \*SetOptions\) GetOnlyIfAbsent\(\) bool](<#SetOptions.GetOnlyIfAbsent>)
  - [func \(x \*SetOptions\) GetOnlyIfPresent\(\) bool](<#SetOptions.GetOnlyIfPresent>)
  - [func \(x \*SetOptions\) GetReturnPrevious\(\) bool](<#SetOptions.GetReturnPrevious>)
  - [func \(x \*SetOptions\) GetTtl\(\) \*durationpb.Duration](<#SetOptions.GetTtl>)
  - [func \(x \*SetOptions\) HasOnlyIfAbsent\(\) bool](<#SetOptions.HasOnlyIfAbsent>)
  - [func \(x \*SetOptions\) HasOnlyIfPresent\(\) bool](<#SetOptions.HasOnlyIfPresent>)
  - [func \(x \*SetOptions\) HasReturnPrevious\(\) bool](<#SetOptions.HasReturnPrevious>)
  - [func \(x \*SetOptions\) HasTtl\(\) bool](<#SetOptions.HasTtl>)
  - [func \(\*SetOptions\) ProtoMessage\(\)](<#SetOptions.ProtoMessage>)
  - [func \(x \*SetOptions\) ProtoReflect\(\) protoreflect.Message](<#SetOptions.ProtoReflect>)
  - [func \(x \*SetOptions\) Reset\(\)](<#SetOptions.Reset>)
  - [func \(x \*SetOptions\) SetOnlyIfAbsent\(v bool\)](<#SetOptions.SetOnlyIfAbsent>)
  - [func \(x \*SetOptions\) SetOnlyIfPresent\(v bool\)](<#SetOptions.SetOnlyIfPresent>)
  - [func \(x \*SetOptions\) SetReturnPrevious\(v bool\)](<#SetOptions.SetReturnPrevious>)
  - [func \(x \*SetOptions\) SetTtl\(v \*durationpb.Duration\)](<#SetOptions.SetTtl>)
  - [func \(x \*SetOptions\) String\(\) string](<#SetOptions.String>)
- [type SetOptions\_builder](<#SetOptions_builder>)
  - [func \(b0 SetOptions\_builder\) Build\(\) \*SetOptions](<#SetOptions_builder.Build>)
- [type SetRequest](<#SetRequest>)
  - [func \(x \*SetRequest\) ClearKey\(\)](<#SetRequest.ClearKey>)
  - [func \(x \*SetRequest\) ClearMetadata\(\)](<#SetRequest.ClearMetadata>)
  - [func \(x \*SetRequest\) ClearNamespace\(\)](<#SetRequest.ClearNamespace>)
  - [func \(x \*SetRequest\) ClearOverwrite\(\)](<#SetRequest.ClearOverwrite>)
  - [func \(x \*SetRequest\) ClearTtl\(\)](<#SetRequest.ClearTtl>)
  - [func \(x \*SetRequest\) ClearValue\(\)](<#SetRequest.ClearValue>)
  - [func \(x \*SetRequest\) GetEntryMetadata\(\) map\[string\]string](<#SetRequest.GetEntryMetadata>)
  - [func \(x \*SetRequest\) GetKey\(\) string](<#SetRequest.GetKey>)
  - [func \(x \*SetRequest\) GetMetadata\(\) \*common.RequestMetadata](<#SetRequest.GetMetadata>)
  - [func \(x \*SetRequest\) GetNamespace\(\) string](<#SetRequest.GetNamespace>)
  - [func \(x \*SetRequest\) GetOverwrite\(\) bool](<#SetRequest.GetOverwrite>)
  - [func \(x \*SetRequest\) GetTtl\(\) \*durationpb.Duration](<#SetRequest.GetTtl>)
  - [func \(x \*SetRequest\) GetValue\(\) \*anypb.Any](<#SetRequest.GetValue>)
  - [func \(x \*SetRequest\) HasKey\(\) bool](<#SetRequest.HasKey>)
  - [func \(x \*SetRequest\) HasMetadata\(\) bool](<#SetRequest.HasMetadata>)
  - [func \(x \*SetRequest\) HasNamespace\(\) bool](<#SetRequest.HasNamespace>)
  - [func \(x \*SetRequest\) HasOverwrite\(\) bool](<#SetRequest.HasOverwrite>)
  - [func \(x \*SetRequest\) HasTtl\(\) bool](<#SetRequest.HasTtl>)
  - [func \(x \*SetRequest\) HasValue\(\) bool](<#SetRequest.HasValue>)
  - [func \(\*SetRequest\) ProtoMessage\(\)](<#SetRequest.ProtoMessage>)
  - [func \(x \*SetRequest\) ProtoReflect\(\) protoreflect.Message](<#SetRequest.ProtoReflect>)
  - [func \(x \*SetRequest\) Reset\(\)](<#SetRequest.Reset>)
  - [func \(x \*SetRequest\) SetEntryMetadata\(v map\[string\]string\)](<#SetRequest.SetEntryMetadata>)
  - [func \(x \*SetRequest\) SetKey\(v string\)](<#SetRequest.SetKey>)
  - [func \(x \*SetRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#SetRequest.SetMetadata>)
  - [func \(x \*SetRequest\) SetNamespace\(v string\)](<#SetRequest.SetNamespace>)
  - [func \(x \*SetRequest\) SetOverwrite\(v bool\)](<#SetRequest.SetOverwrite>)
  - [func \(x \*SetRequest\) SetTtl\(v \*durationpb.Duration\)](<#SetRequest.SetTtl>)
  - [func \(x \*SetRequest\) SetValue\(v \*anypb.Any\)](<#SetRequest.SetValue>)
  - [func \(x \*SetRequest\) String\(\) string](<#SetRequest.String>)
- [type SetRequest\_builder](<#SetRequest_builder>)
  - [func \(b0 SetRequest\_builder\) Build\(\) \*SetRequest](<#SetRequest_builder.Build>)
- [type SetResponse](<#SetResponse>)
  - [func \(x \*SetResponse\) ClearOverwritten\(\)](<#SetResponse.ClearOverwritten>)
  - [func \(x \*SetResponse\) ClearSizeBytes\(\)](<#SetResponse.ClearSizeBytes>)
  - [func \(x \*SetResponse\) ClearSuccess\(\)](<#SetResponse.ClearSuccess>)
  - [func \(x \*SetResponse\) GetOverwritten\(\) bool](<#SetResponse.GetOverwritten>)
  - [func \(x \*SetResponse\) GetSizeBytes\(\) int64](<#SetResponse.GetSizeBytes>)
  - [func \(x \*SetResponse\) GetSuccess\(\) bool](<#SetResponse.GetSuccess>)
  - [func \(x \*SetResponse\) HasOverwritten\(\) bool](<#SetResponse.HasOverwritten>)
  - [func \(x \*SetResponse\) HasSizeBytes\(\) bool](<#SetResponse.HasSizeBytes>)
  - [func \(x \*SetResponse\) HasSuccess\(\) bool](<#SetResponse.HasSuccess>)
  - [func \(\*SetResponse\) ProtoMessage\(\)](<#SetResponse.ProtoMessage>)
  - [func \(x \*SetResponse\) ProtoReflect\(\) protoreflect.Message](<#SetResponse.ProtoReflect>)
  - [func \(x \*SetResponse\) Reset\(\)](<#SetResponse.Reset>)
  - [func \(x \*SetResponse\) SetOverwritten\(v bool\)](<#SetResponse.SetOverwritten>)
  - [func \(x \*SetResponse\) SetSizeBytes\(v int64\)](<#SetResponse.SetSizeBytes>)
  - [func \(x \*SetResponse\) SetSuccess\(v bool\)](<#SetResponse.SetSuccess>)
  - [func \(x \*SetResponse\) String\(\) string](<#SetResponse.String>)
- [type SetResponse\_builder](<#SetResponse_builder>)
  - [func \(b0 SetResponse\_builder\) Build\(\) \*SetResponse](<#SetResponse_builder.Build>)
- [type TouchExpirationRequest](<#TouchExpirationRequest>)
  - [func \(x \*TouchExpirationRequest\) ClearKey\(\)](<#TouchExpirationRequest.ClearKey>)
  - [func \(x \*TouchExpirationRequest\) ClearMetadata\(\)](<#TouchExpirationRequest.ClearMetadata>)
  - [func \(x \*TouchExpirationRequest\) ClearNamespace\(\)](<#TouchExpirationRequest.ClearNamespace>)
  - [func \(x \*TouchExpirationRequest\) ClearTtl\(\)](<#TouchExpirationRequest.ClearTtl>)
  - [func \(x \*TouchExpirationRequest\) GetKey\(\) string](<#TouchExpirationRequest.GetKey>)
  - [func \(x \*TouchExpirationRequest\) GetMetadata\(\) \*common.RequestMetadata](<#TouchExpirationRequest.GetMetadata>)
  - [func \(x \*TouchExpirationRequest\) GetNamespace\(\) string](<#TouchExpirationRequest.GetNamespace>)
  - [func \(x \*TouchExpirationRequest\) GetTtl\(\) \*durationpb.Duration](<#TouchExpirationRequest.GetTtl>)
  - [func \(x \*TouchExpirationRequest\) HasKey\(\) bool](<#TouchExpirationRequest.HasKey>)
  - [func \(x \*TouchExpirationRequest\) HasMetadata\(\) bool](<#TouchExpirationRequest.HasMetadata>)
  - [func \(x \*TouchExpirationRequest\) HasNamespace\(\) bool](<#TouchExpirationRequest.HasNamespace>)
  - [func \(x \*TouchExpirationRequest\) HasTtl\(\) bool](<#TouchExpirationRequest.HasTtl>)
  - [func \(\*TouchExpirationRequest\) ProtoMessage\(\)](<#TouchExpirationRequest.ProtoMessage>)
  - [func \(x \*TouchExpirationRequest\) ProtoReflect\(\) protoreflect.Message](<#TouchExpirationRequest.ProtoReflect>)
  - [func \(x \*TouchExpirationRequest\) Reset\(\)](<#TouchExpirationRequest.Reset>)
  - [func \(x \*TouchExpirationRequest\) SetKey\(v string\)](<#TouchExpirationRequest.SetKey>)
  - [func \(x \*TouchExpirationRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#TouchExpirationRequest.SetMetadata>)
  - [func \(x \*TouchExpirationRequest\) SetNamespace\(v string\)](<#TouchExpirationRequest.SetNamespace>)
  - [func \(x \*TouchExpirationRequest\) SetTtl\(v \*durationpb.Duration\)](<#TouchExpirationRequest.SetTtl>)
  - [func \(x \*TouchExpirationRequest\) String\(\) string](<#TouchExpirationRequest.String>)
- [type TouchExpirationRequest\_builder](<#TouchExpirationRequest_builder>)
  - [func \(b0 TouchExpirationRequest\_builder\) Build\(\) \*TouchExpirationRequest](<#TouchExpirationRequest_builder.Build>)
- [type TouchExpirationResponse](<#TouchExpirationResponse>)
  - [func \(x \*TouchExpirationResponse\) ClearError\(\)](<#TouchExpirationResponse.ClearError>)
  - [func \(x \*TouchExpirationResponse\) ClearKeyExisted\(\)](<#TouchExpirationResponse.ClearKeyExisted>)
  - [func \(x \*TouchExpirationResponse\) ClearSuccess\(\)](<#TouchExpirationResponse.ClearSuccess>)
  - [func \(x \*TouchExpirationResponse\) GetError\(\) \*common.Error](<#TouchExpirationResponse.GetError>)
  - [func \(x \*TouchExpirationResponse\) GetKeyExisted\(\) bool](<#TouchExpirationResponse.GetKeyExisted>)
  - [func \(x \*TouchExpirationResponse\) GetSuccess\(\) bool](<#TouchExpirationResponse.GetSuccess>)
  - [func \(x \*TouchExpirationResponse\) HasError\(\) bool](<#TouchExpirationResponse.HasError>)
  - [func \(x \*TouchExpirationResponse\) HasKeyExisted\(\) bool](<#TouchExpirationResponse.HasKeyExisted>)
  - [func \(x \*TouchExpirationResponse\) HasSuccess\(\) bool](<#TouchExpirationResponse.HasSuccess>)
  - [func \(\*TouchExpirationResponse\) ProtoMessage\(\)](<#TouchExpirationResponse.ProtoMessage>)
  - [func \(x \*TouchExpirationResponse\) ProtoReflect\(\) protoreflect.Message](<#TouchExpirationResponse.ProtoReflect>)
  - [func \(x \*TouchExpirationResponse\) Reset\(\)](<#TouchExpirationResponse.Reset>)
  - [func \(x \*TouchExpirationResponse\) SetError\(v \*common.Error\)](<#TouchExpirationResponse.SetError>)
  - [func \(x \*TouchExpirationResponse\) SetKeyExisted\(v bool\)](<#TouchExpirationResponse.SetKeyExisted>)
  - [func \(x \*TouchExpirationResponse\) SetSuccess\(v bool\)](<#TouchExpirationResponse.SetSuccess>)
  - [func \(x \*TouchExpirationResponse\) String\(\) string](<#TouchExpirationResponse.String>)
- [type TouchExpirationResponse\_builder](<#TouchExpirationResponse_builder>)
  - [func \(b0 TouchExpirationResponse\_builder\) Build\(\) \*TouchExpirationResponse](<#TouchExpirationResponse_builder.Build>)
- [type TransactionOptions](<#TransactionOptions>)
  - [func \(x \*TransactionOptions\) ClearIsolation\(\)](<#TransactionOptions.ClearIsolation>)
  - [func \(x \*TransactionOptions\) ClearReadOnly\(\)](<#TransactionOptions.ClearReadOnly>)
  - [func \(x \*TransactionOptions\) ClearTimeout\(\)](<#TransactionOptions.ClearTimeout>)
  - [func \(x \*TransactionOptions\) GetIsolation\(\) common.DatabaseIsolationLevel](<#TransactionOptions.GetIsolation>)
  - [func \(x \*TransactionOptions\) GetReadOnly\(\) bool](<#TransactionOptions.GetReadOnly>)
  - [func \(x \*TransactionOptions\) GetTimeout\(\) \*durationpb.Duration](<#TransactionOptions.GetTimeout>)
  - [func \(x \*TransactionOptions\) HasIsolation\(\) bool](<#TransactionOptions.HasIsolation>)
  - [func \(x \*TransactionOptions\) HasReadOnly\(\) bool](<#TransactionOptions.HasReadOnly>)
  - [func \(x \*TransactionOptions\) HasTimeout\(\) bool](<#TransactionOptions.HasTimeout>)
  - [func \(\*TransactionOptions\) ProtoMessage\(\)](<#TransactionOptions.ProtoMessage>)
  - [func \(x \*TransactionOptions\) ProtoReflect\(\) protoreflect.Message](<#TransactionOptions.ProtoReflect>)
  - [func \(x \*TransactionOptions\) Reset\(\)](<#TransactionOptions.Reset>)
  - [func \(x \*TransactionOptions\) SetIsolation\(v common.DatabaseIsolationLevel\)](<#TransactionOptions.SetIsolation>)
  - [func \(x \*TransactionOptions\) SetReadOnly\(v bool\)](<#TransactionOptions.SetReadOnly>)
  - [func \(x \*TransactionOptions\) SetTimeout\(v \*durationpb.Duration\)](<#TransactionOptions.SetTimeout>)
  - [func \(x \*TransactionOptions\) String\(\) string](<#TransactionOptions.String>)
- [type TransactionOptions\_builder](<#TransactionOptions_builder>)
  - [func \(b0 TransactionOptions\_builder\) Build\(\) \*TransactionOptions](<#TransactionOptions_builder.Build>)
- [type TransactionRequest](<#TransactionRequest>)
  - [func \(x \*TransactionRequest\) ClearMetadata\(\)](<#TransactionRequest.ClearMetadata>)
  - [func \(x \*TransactionRequest\) ClearOperations\(\)](<#TransactionRequest.ClearOperations>)
  - [func \(x \*TransactionRequest\) GetMetadata\(\) \*common.RequestMetadata](<#TransactionRequest.GetMetadata>)
  - [func \(x \*TransactionRequest\) GetOperations\(\) \[\]byte](<#TransactionRequest.GetOperations>)
  - [func \(x \*TransactionRequest\) HasMetadata\(\) bool](<#TransactionRequest.HasMetadata>)
  - [func \(x \*TransactionRequest\) HasOperations\(\) bool](<#TransactionRequest.HasOperations>)
  - [func \(\*TransactionRequest\) ProtoMessage\(\)](<#TransactionRequest.ProtoMessage>)
  - [func \(x \*TransactionRequest\) ProtoReflect\(\) protoreflect.Message](<#TransactionRequest.ProtoReflect>)
  - [func \(x \*TransactionRequest\) Reset\(\)](<#TransactionRequest.Reset>)
  - [func \(x \*TransactionRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#TransactionRequest.SetMetadata>)
  - [func \(x \*TransactionRequest\) SetOperations\(v \[\]byte\)](<#TransactionRequest.SetOperations>)
  - [func \(x \*TransactionRequest\) String\(\) string](<#TransactionRequest.String>)
- [type TransactionRequest\_builder](<#TransactionRequest_builder>)
  - [func \(b0 TransactionRequest\_builder\) Build\(\) \*TransactionRequest](<#TransactionRequest_builder.Build>)
- [type TransactionServiceClient](<#TransactionServiceClient>)
  - [func NewTransactionServiceClient\(cc grpc.ClientConnInterface\) TransactionServiceClient](<#NewTransactionServiceClient>)
- [type TransactionServiceServer](<#TransactionServiceServer>)
- [type TransactionStatusRequest](<#TransactionStatusRequest>)
  - [func \(x \*TransactionStatusRequest\) ClearMetadata\(\)](<#TransactionStatusRequest.ClearMetadata>)
  - [func \(x \*TransactionStatusRequest\) ClearTransactionId\(\)](<#TransactionStatusRequest.ClearTransactionId>)
  - [func \(x \*TransactionStatusRequest\) GetMetadata\(\) \*common.RequestMetadata](<#TransactionStatusRequest.GetMetadata>)
  - [func \(x \*TransactionStatusRequest\) GetTransactionId\(\) string](<#TransactionStatusRequest.GetTransactionId>)
  - [func \(x \*TransactionStatusRequest\) HasMetadata\(\) bool](<#TransactionStatusRequest.HasMetadata>)
  - [func \(x \*TransactionStatusRequest\) HasTransactionId\(\) bool](<#TransactionStatusRequest.HasTransactionId>)
  - [func \(\*TransactionStatusRequest\) ProtoMessage\(\)](<#TransactionStatusRequest.ProtoMessage>)
  - [func \(x \*TransactionStatusRequest\) ProtoReflect\(\) protoreflect.Message](<#TransactionStatusRequest.ProtoReflect>)
  - [func \(x \*TransactionStatusRequest\) Reset\(\)](<#TransactionStatusRequest.Reset>)
  - [func \(x \*TransactionStatusRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#TransactionStatusRequest.SetMetadata>)
  - [func \(x \*TransactionStatusRequest\) SetTransactionId\(v string\)](<#TransactionStatusRequest.SetTransactionId>)
  - [func \(x \*TransactionStatusRequest\) String\(\) string](<#TransactionStatusRequest.String>)
- [type TransactionStatusRequest\_builder](<#TransactionStatusRequest_builder>)
  - [func \(b0 TransactionStatusRequest\_builder\) Build\(\) \*TransactionStatusRequest](<#TransactionStatusRequest_builder.Build>)
- [type TransactionStatusResponse](<#TransactionStatusResponse>)
  - [func \(x \*TransactionStatusResponse\) ClearError\(\)](<#TransactionStatusResponse.ClearError>)
  - [func \(x \*TransactionStatusResponse\) ClearStatus\(\)](<#TransactionStatusResponse.ClearStatus>)
  - [func \(x \*TransactionStatusResponse\) GetError\(\) \*common.Error](<#TransactionStatusResponse.GetError>)
  - [func \(x \*TransactionStatusResponse\) GetStatus\(\) string](<#TransactionStatusResponse.GetStatus>)
  - [func \(x \*TransactionStatusResponse\) HasError\(\) bool](<#TransactionStatusResponse.HasError>)
  - [func \(x \*TransactionStatusResponse\) HasStatus\(\) bool](<#TransactionStatusResponse.HasStatus>)
  - [func \(\*TransactionStatusResponse\) ProtoMessage\(\)](<#TransactionStatusResponse.ProtoMessage>)
  - [func \(x \*TransactionStatusResponse\) ProtoReflect\(\) protoreflect.Message](<#TransactionStatusResponse.ProtoReflect>)
  - [func \(x \*TransactionStatusResponse\) Reset\(\)](<#TransactionStatusResponse.Reset>)
  - [func \(x \*TransactionStatusResponse\) SetError\(v \*common.Error\)](<#TransactionStatusResponse.SetError>)
  - [func \(x \*TransactionStatusResponse\) SetStatus\(v string\)](<#TransactionStatusResponse.SetStatus>)
  - [func \(x \*TransactionStatusResponse\) String\(\) string](<#TransactionStatusResponse.String>)
- [type TransactionStatusResponse\_builder](<#TransactionStatusResponse_builder>)
  - [func \(b0 TransactionStatusResponse\_builder\) Build\(\) \*TransactionStatusResponse](<#TransactionStatusResponse_builder.Build>)
- [type UnimplementedCacheAdminServiceServer](<#UnimplementedCacheAdminServiceServer>)
  - [func \(UnimplementedCacheAdminServiceServer\) ConfigurePolicy\(context.Context, \*ConfigurePolicyRequest\) \(\*ConfigurePolicyResponse, error\)](<#UnimplementedCacheAdminServiceServer.ConfigurePolicy>)
  - [func \(UnimplementedCacheAdminServiceServer\) CreateNamespace\(context.Context, \*CreateNamespaceRequest\) \(\*CreateNamespaceResponse, error\)](<#UnimplementedCacheAdminServiceServer.CreateNamespace>)
  - [func \(UnimplementedCacheAdminServiceServer\) DeleteNamespace\(context.Context, \*DeleteNamespaceRequest\) \(\*emptypb.Empty, error\)](<#UnimplementedCacheAdminServiceServer.DeleteNamespace>)
  - [func \(UnimplementedCacheAdminServiceServer\) GetNamespaceStats\(context.Context, \*GetNamespaceStatsRequest\) \(\*GetNamespaceStatsResponse, error\)](<#UnimplementedCacheAdminServiceServer.GetNamespaceStats>)
  - [func \(UnimplementedCacheAdminServiceServer\) ListNamespaces\(context.Context, \*ListNamespacesRequest\) \(\*ListNamespacesResponse, error\)](<#UnimplementedCacheAdminServiceServer.ListNamespaces>)
- [type UnimplementedCacheServiceServer](<#UnimplementedCacheServiceServer>)
  - [func \(UnimplementedCacheServiceServer\) Clear\(context.Context, \*ClearRequest\) \(\*ClearResponse, error\)](<#UnimplementedCacheServiceServer.Clear>)
  - [func \(UnimplementedCacheServiceServer\) Decrement\(context.Context, \*DecrementRequest\) \(\*DecrementResponse, error\)](<#UnimplementedCacheServiceServer.Decrement>)
  - [func \(UnimplementedCacheServiceServer\) Delete\(context.Context, \*CacheDeleteRequest\) \(\*CacheDeleteResponse, error\)](<#UnimplementedCacheServiceServer.Delete>)
  - [func \(UnimplementedCacheServiceServer\) DeleteMultiple\(context.Context, \*DeleteMultipleRequest\) \(\*DeleteMultipleResponse, error\)](<#UnimplementedCacheServiceServer.DeleteMultiple>)
  - [func \(UnimplementedCacheServiceServer\) Exists\(context.Context, \*ExistsRequest\) \(\*ExistsResponse, error\)](<#UnimplementedCacheServiceServer.Exists>)
  - [func \(UnimplementedCacheServiceServer\) Flush\(context.Context, \*FlushRequest\) \(\*FlushResponse, error\)](<#UnimplementedCacheServiceServer.Flush>)
  - [func \(UnimplementedCacheServiceServer\) Get\(context.Context, \*GetRequest\) \(\*GetResponse, error\)](<#UnimplementedCacheServiceServer.Get>)
  - [func \(UnimplementedCacheServiceServer\) GetMultiple\(context.Context, \*GetMultipleRequest\) \(\*GetMultipleResponse, error\)](<#UnimplementedCacheServiceServer.GetMultiple>)
  - [func \(UnimplementedCacheServiceServer\) GetStats\(context.Context, \*CacheGetStatsRequest\) \(\*CacheGetStatsResponse, error\)](<#UnimplementedCacheServiceServer.GetStats>)
  - [func \(UnimplementedCacheServiceServer\) Increment\(context.Context, \*IncrementRequest\) \(\*IncrementResponse, error\)](<#UnimplementedCacheServiceServer.Increment>)
  - [func \(UnimplementedCacheServiceServer\) Keys\(context.Context, \*KeysRequest\) \(\*KeysResponse, error\)](<#UnimplementedCacheServiceServer.Keys>)
  - [func \(UnimplementedCacheServiceServer\) Set\(context.Context, \*SetRequest\) \(\*SetResponse, error\)](<#UnimplementedCacheServiceServer.Set>)
  - [func \(UnimplementedCacheServiceServer\) SetMultiple\(context.Context, \*SetMultipleRequest\) \(\*SetMultipleResponse, error\)](<#UnimplementedCacheServiceServer.SetMultiple>)
  - [func \(UnimplementedCacheServiceServer\) TouchExpiration\(context.Context, \*TouchExpirationRequest\) \(\*TouchExpirationResponse, error\)](<#UnimplementedCacheServiceServer.TouchExpiration>)
- [type UnimplementedDatabaseAdminServiceServer](<#UnimplementedDatabaseAdminServiceServer>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) CreateDatabase\(context.Context, \*CreateDatabaseRequest\) \(\*CreateDatabaseResponse, error\)](<#UnimplementedDatabaseAdminServiceServer.CreateDatabase>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) CreateSchema\(context.Context, \*CreateSchemaRequest\) \(\*CreateSchemaResponse, error\)](<#UnimplementedDatabaseAdminServiceServer.CreateSchema>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) DropDatabase\(context.Context, \*DropDatabaseRequest\) \(\*emptypb.Empty, error\)](<#UnimplementedDatabaseAdminServiceServer.DropDatabase>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) DropSchema\(context.Context, \*DropSchemaRequest\) \(\*emptypb.Empty, error\)](<#UnimplementedDatabaseAdminServiceServer.DropSchema>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) GetDatabaseInfo\(context.Context, \*GetDatabaseInfoRequest\) \(\*GetDatabaseInfoResponse, error\)](<#UnimplementedDatabaseAdminServiceServer.GetDatabaseInfo>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) ListDatabases\(context.Context, \*ListDatabasesRequest\) \(\*ListDatabasesResponse, error\)](<#UnimplementedDatabaseAdminServiceServer.ListDatabases>)
  - [func \(UnimplementedDatabaseAdminServiceServer\) ListSchemas\(context.Context, \*ListSchemasRequest\) \(\*ListSchemasResponse, error\)](<#UnimplementedDatabaseAdminServiceServer.ListSchemas>)
- [type UnimplementedDatabaseServiceServer](<#UnimplementedDatabaseServiceServer>)
  - [func \(UnimplementedDatabaseServiceServer\) Execute\(context.Context, \*ExecuteRequest\) \(\*ExecuteResponse, error\)](<#UnimplementedDatabaseServiceServer.Execute>)
  - [func \(UnimplementedDatabaseServiceServer\) ExecuteBatch\(context.Context, \*ExecuteBatchRequest\) \(\*ExecuteBatchResponse, error\)](<#UnimplementedDatabaseServiceServer.ExecuteBatch>)
  - [func \(UnimplementedDatabaseServiceServer\) GetConnectionInfo\(context.Context, \*GetConnectionInfoRequest\) \(\*GetConnectionInfoResponse, error\)](<#UnimplementedDatabaseServiceServer.GetConnectionInfo>)
  - [func \(UnimplementedDatabaseServiceServer\) HealthCheck\(context.Context, \*DatabaseHealthCheckRequest\) \(\*DatabaseHealthCheckResponse, error\)](<#UnimplementedDatabaseServiceServer.HealthCheck>)
  - [func \(UnimplementedDatabaseServiceServer\) Query\(context.Context, \*QueryRequest\) \(\*QueryResponse, error\)](<#UnimplementedDatabaseServiceServer.Query>)
  - [func \(UnimplementedDatabaseServiceServer\) QueryRow\(context.Context, \*QueryRowRequest\) \(\*QueryRowResponse, error\)](<#UnimplementedDatabaseServiceServer.QueryRow>)
- [type UnimplementedMigrationServiceServer](<#UnimplementedMigrationServiceServer>)
  - [func \(UnimplementedMigrationServiceServer\) ApplyMigration\(context.Context, \*RunMigrationRequest\) \(\*RunMigrationResponse, error\)](<#UnimplementedMigrationServiceServer.ApplyMigration>)
  - [func \(UnimplementedMigrationServiceServer\) GetMigrationStatus\(context.Context, \*GetMigrationStatusRequest\) \(\*GetMigrationStatusResponse, error\)](<#UnimplementedMigrationServiceServer.GetMigrationStatus>)
  - [func \(UnimplementedMigrationServiceServer\) ListMigrations\(context.Context, \*ListMigrationsRequest\) \(\*ListMigrationsResponse, error\)](<#UnimplementedMigrationServiceServer.ListMigrations>)
  - [func \(UnimplementedMigrationServiceServer\) RevertMigration\(context.Context, \*RevertMigrationRequest\) \(\*RevertMigrationResponse, error\)](<#UnimplementedMigrationServiceServer.RevertMigration>)
- [type UnimplementedTransactionServiceServer](<#UnimplementedTransactionServiceServer>)
  - [func \(UnimplementedTransactionServiceServer\) BeginTransaction\(context.Context, \*BeginTransactionRequest\) \(\*BeginTransactionResponse, error\)](<#UnimplementedTransactionServiceServer.BeginTransaction>)
  - [func \(UnimplementedTransactionServiceServer\) CommitTransaction\(context.Context, \*CommitTransactionRequest\) \(\*emptypb.Empty, error\)](<#UnimplementedTransactionServiceServer.CommitTransaction>)
  - [func \(UnimplementedTransactionServiceServer\) GetTransactionStatus\(context.Context, \*TransactionStatusRequest\) \(\*TransactionStatusResponse, error\)](<#UnimplementedTransactionServiceServer.GetTransactionStatus>)
  - [func \(UnimplementedTransactionServiceServer\) RollbackTransaction\(context.Context, \*RollbackTransactionRequest\) \(\*emptypb.Empty, error\)](<#UnimplementedTransactionServiceServer.RollbackTransaction>)
- [type UnlockRequest](<#UnlockRequest>)
  - [func \(x \*UnlockRequest\) ClearKey\(\)](<#UnlockRequest.ClearKey>)
  - [func \(x \*UnlockRequest\) ClearMetadata\(\)](<#UnlockRequest.ClearMetadata>)
  - [func \(x \*UnlockRequest\) ClearNamespace\(\)](<#UnlockRequest.ClearNamespace>)
  - [func \(x \*UnlockRequest\) GetKey\(\) string](<#UnlockRequest.GetKey>)
  - [func \(x \*UnlockRequest\) GetMetadata\(\) \*common.RequestMetadata](<#UnlockRequest.GetMetadata>)
  - [func \(x \*UnlockRequest\) GetNamespace\(\) string](<#UnlockRequest.GetNamespace>)
  - [func \(x \*UnlockRequest\) HasKey\(\) bool](<#UnlockRequest.HasKey>)
  - [func \(x \*UnlockRequest\) HasMetadata\(\) bool](<#UnlockRequest.HasMetadata>)
  - [func \(x \*UnlockRequest\) HasNamespace\(\) bool](<#UnlockRequest.HasNamespace>)
  - [func \(\*UnlockRequest\) ProtoMessage\(\)](<#UnlockRequest.ProtoMessage>)
  - [func \(x \*UnlockRequest\) ProtoReflect\(\) protoreflect.Message](<#UnlockRequest.ProtoReflect>)
  - [func \(x \*UnlockRequest\) Reset\(\)](<#UnlockRequest.Reset>)
  - [func \(x \*UnlockRequest\) SetKey\(v string\)](<#UnlockRequest.SetKey>)
  - [func \(x \*UnlockRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#UnlockRequest.SetMetadata>)
  - [func \(x \*UnlockRequest\) SetNamespace\(v string\)](<#UnlockRequest.SetNamespace>)
  - [func \(x \*UnlockRequest\) String\(\) string](<#UnlockRequest.String>)
- [type UnlockRequest\_builder](<#UnlockRequest_builder>)
  - [func \(b0 UnlockRequest\_builder\) Build\(\) \*UnlockRequest](<#UnlockRequest_builder.Build>)
- [type UnsafeCacheAdminServiceServer](<#UnsafeCacheAdminServiceServer>)
- [type UnsafeCacheServiceServer](<#UnsafeCacheServiceServer>)
- [type UnsafeDatabaseAdminServiceServer](<#UnsafeDatabaseAdminServiceServer>)
- [type UnsafeDatabaseServiceServer](<#UnsafeDatabaseServiceServer>)
- [type UnsafeMigrationServiceServer](<#UnsafeMigrationServiceServer>)
- [type UnsafeTransactionServiceServer](<#UnsafeTransactionServiceServer>)
- [type UnwatchRequest](<#UnwatchRequest>)
  - [func \(x \*UnwatchRequest\) ClearKey\(\)](<#UnwatchRequest.ClearKey>)
  - [func \(x \*UnwatchRequest\) ClearMetadata\(\)](<#UnwatchRequest.ClearMetadata>)
  - [func \(x \*UnwatchRequest\) GetKey\(\) string](<#UnwatchRequest.GetKey>)
  - [func \(x \*UnwatchRequest\) GetMetadata\(\) \*common.RequestMetadata](<#UnwatchRequest.GetMetadata>)
  - [func \(x \*UnwatchRequest\) HasKey\(\) bool](<#UnwatchRequest.HasKey>)
  - [func \(x \*UnwatchRequest\) HasMetadata\(\) bool](<#UnwatchRequest.HasMetadata>)
  - [func \(\*UnwatchRequest\) ProtoMessage\(\)](<#UnwatchRequest.ProtoMessage>)
  - [func \(x \*UnwatchRequest\) ProtoReflect\(\) protoreflect.Message](<#UnwatchRequest.ProtoReflect>)
  - [func \(x \*UnwatchRequest\) Reset\(\)](<#UnwatchRequest.Reset>)
  - [func \(x \*UnwatchRequest\) SetKey\(v string\)](<#UnwatchRequest.SetKey>)
  - [func \(x \*UnwatchRequest\) SetMetadata\(v \*common.RequestMetadata\)](<#UnwatchRequest.SetMetadata>)
  - [func \(x \*UnwatchRequest\) String\(\) string](<#UnwatchRequest.String>)
- [type UnwatchRequest\_builder](<#UnwatchRequest_builder>)
  - [func \(b0 UnwatchRequest\_builder\) Build\(\) \*UnwatchRequest](<#UnwatchRequest_builder.Build>)
## Constants
<a name="MigrationService_ApplyMigration_FullMethodName"></a>
    MigrationService_ApplyMigration_FullMethodName     = "/gcommon.v1.database.MigrationService/ApplyMigration"
    MigrationService_RevertMigration_FullMethodName    = "/gcommon.v1.database.MigrationService/RevertMigration"
    MigrationService_GetMigrationStatus_FullMethodName = "/gcommon.v1.database.MigrationService/GetMigrationStatus"
    MigrationService_ListMigrations_FullMethodName     = "/gcommon.v1.database.MigrationService/ListMigrations"
<a name="TransactionService_BeginTransaction_FullMethodName"></a>
    TransactionService_BeginTransaction_FullMethodName     = "/gcommon.v1.database.TransactionService/BeginTransaction"
    TransactionService_CommitTransaction_FullMethodName    = "/gcommon.v1.database.TransactionService/CommitTransaction"
    TransactionService_RollbackTransaction_FullMethodName  = "/gcommon.v1.database.TransactionService/RollbackTransaction"
    TransactionService_GetTransactionStatus_FullMethodName = "/gcommon.v1.database.TransactionService/GetTransactionStatus"
## Variables
<a name="CacheAdminService_ServiceDesc"></a>CacheAdminService\_ServiceDesc is the grpc.ServiceDesc for CacheAdminService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var CacheAdminService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.CacheAdminService",
    HandlerType: (*CacheAdminServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "CreateNamespace",
            Handler:    _CacheAdminService_CreateNamespace_Handler,
        },
        {
            MethodName: "DeleteNamespace",
            Handler:    _CacheAdminService_DeleteNamespace_Handler,
        },
        {
            MethodName: "ListNamespaces",
            Handler:    _CacheAdminService_ListNamespaces_Handler,
        },
        {
            MethodName: "GetNamespaceStats",
            Handler:    _CacheAdminService_GetNamespaceStats_Handler,
        },
        {
            MethodName: "ConfigurePolicy",
            Handler:    _CacheAdminService_ConfigurePolicy_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/cache_admin_service.proto",
}
<a name="CacheService_ServiceDesc"></a>CacheService\_ServiceDesc is the grpc.ServiceDesc for CacheService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var CacheService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.CacheService",
    HandlerType: (*CacheServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "Get",
            Handler:    _CacheService_Get_Handler,
        },
        {
            MethodName: "Set",
            Handler:    _CacheService_Set_Handler,
        },
        {
            MethodName: "Delete",
            Handler:    _CacheService_Delete_Handler,
        },
        {
            MethodName: "Exists",
            Handler:    _CacheService_Exists_Handler,
        },
        {
            MethodName: "GetMultiple",
            Handler:    _CacheService_GetMultiple_Handler,
        },
        {
            MethodName: "SetMultiple",
            Handler:    _CacheService_SetMultiple_Handler,
        },
        {
            MethodName: "DeleteMultiple",
            Handler:    _CacheService_DeleteMultiple_Handler,
        },
        {
            MethodName: "Increment",
            Handler:    _CacheService_Increment_Handler,
        },
        {
            MethodName: "Decrement",
            Handler:    _CacheService_Decrement_Handler,
        },
        {
            MethodName: "Clear",
            Handler:    _CacheService_Clear_Handler,
        },
        {
            MethodName: "Keys",
            Handler:    _CacheService_Keys_Handler,
        },
        {
            MethodName: "GetStats",
            Handler:    _CacheService_GetStats_Handler,
        },
        {
            MethodName: "Flush",
            Handler:    _CacheService_Flush_Handler,
        },
        {
            MethodName: "TouchExpiration",
            Handler:    _CacheService_TouchExpiration_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/cache_service.proto",
<a name="DatabaseAdminService_ServiceDesc"></a>DatabaseAdminService\_ServiceDesc is the grpc.ServiceDesc for DatabaseAdminService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var DatabaseAdminService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.DatabaseAdminService",
    HandlerType: (*DatabaseAdminServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "CreateDatabase",
            Handler:    _DatabaseAdminService_CreateDatabase_Handler,
        },
        {
            MethodName: "DropDatabase",
            Handler:    _DatabaseAdminService_DropDatabase_Handler,
        },
        {
            MethodName: "ListDatabases",
            Handler:    _DatabaseAdminService_ListDatabases_Handler,
        },
        {
            MethodName: "GetDatabaseInfo",
            Handler:    _DatabaseAdminService_GetDatabaseInfo_Handler,
        },
        {
            MethodName: "CreateSchema",
            Handler:    _DatabaseAdminService_CreateSchema_Handler,
        },
        {
            MethodName: "DropSchema",
            Handler:    _DatabaseAdminService_DropSchema_Handler,
        },
        {
            MethodName: "ListSchemas",
            Handler:    _DatabaseAdminService_ListSchemas_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/database_admin_service.proto",
<a name="DatabaseService_ServiceDesc"></a>DatabaseService\_ServiceDesc is the grpc.ServiceDesc for DatabaseService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.DatabaseService",
    HandlerType: (*DatabaseServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "Query",
            Handler:    _DatabaseService_Query_Handler,
        },
        {
            MethodName: "QueryRow",
            Handler:    _DatabaseService_QueryRow_Handler,
        },
        {
            MethodName: "Execute",
            Handler:    _DatabaseService_Execute_Handler,
        },
        {
            MethodName: "ExecuteBatch",
            Handler:    _DatabaseService_ExecuteBatch_Handler,
        },
        {
            MethodName: "GetConnectionInfo",
            Handler:    _DatabaseService_GetConnectionInfo_Handler,
        },
        {
            MethodName: "HealthCheck",
            Handler:    _DatabaseService_HealthCheck_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/database_service.proto",
}
<a name="File_gcommon_v1_database_append_request_proto"></a>
var File_gcommon_v1_database_append_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_backup_request_proto"></a>
var File_gcommon_v1_database_backup_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_batch_execute_options_proto"></a>
var File_gcommon_v1_database_batch_execute_options_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_batch_operation_proto"></a>
var File_gcommon_v1_database_batch_operation_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_batch_operation_result_proto"></a>
var File_gcommon_v1_database_batch_operation_result_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_batch_stats_proto"></a>
var File_gcommon_v1_database_batch_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_begin_transaction_request_proto"></a>
var File_gcommon_v1_database_begin_transaction_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_begin_transaction_response_proto"></a>
var File_gcommon_v1_database_begin_transaction_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_admin_service_proto"></a>
var File_gcommon_v1_database_cache_admin_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_config_proto"></a>
var File_gcommon_v1_database_cache_config_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_entry_proto"></a>
var File_gcommon_v1_database_cache_entry_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_info_proto"></a>
var File_gcommon_v1_database_cache_info_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_metrics_proto"></a>
var File_gcommon_v1_database_cache_metrics_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_operation_result_proto"></a>
var File_gcommon_v1_database_cache_operation_result_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_service_proto"></a>
var File_gcommon_v1_database_cache_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cache_stats_proto"></a>
var File_gcommon_v1_database_cache_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_clear_request_proto"></a>
var File_gcommon_v1_database_clear_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_clear_response_proto"></a>
var File_gcommon_v1_database_clear_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_cockroach_config_proto"></a>
var File_gcommon_v1_database_cockroach_config_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_column_metadata_proto"></a>
var File_gcommon_v1_database_column_metadata_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_commit_transaction_request_proto"></a>
var File_gcommon_v1_database_commit_transaction_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_configure_policy_request_proto"></a>
var File_gcommon_v1_database_configure_policy_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_configure_policy_response_proto"></a>
var File_gcommon_v1_database_configure_policy_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_connection_pool_info_proto"></a>
var File_gcommon_v1_database_connection_pool_info_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_database_request_proto"></a>
var File_gcommon_v1_database_create_database_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_database_response_proto"></a>
var File_gcommon_v1_database_create_database_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_namespace_request_proto"></a>
var File_gcommon_v1_database_create_namespace_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_namespace_response_proto"></a>
var File_gcommon_v1_database_create_namespace_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_schema_request_proto"></a>
var File_gcommon_v1_database_create_schema_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_create_schema_response_proto"></a>
var File_gcommon_v1_database_create_schema_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_database_admin_service_proto"></a>
var File_gcommon_v1_database_database_admin_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_database_info_proto"></a>
var File_gcommon_v1_database_database_info_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_database_service_proto"></a>
var File_gcommon_v1_database_database_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_database_status_proto"></a>
var File_gcommon_v1_database_database_status_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_decrement_request_proto"></a>
var File_gcommon_v1_database_decrement_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_decrement_response_proto"></a>
var File_gcommon_v1_database_decrement_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_defrag_request_proto"></a>
var File_gcommon_v1_database_defrag_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_delete_multiple_request_proto"></a>
var File_gcommon_v1_database_delete_multiple_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_delete_multiple_response_proto"></a>
var File_gcommon_v1_database_delete_multiple_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_delete_namespace_request_proto"></a>
var File_gcommon_v1_database_delete_namespace_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_delete_request_proto"></a>
var File_gcommon_v1_database_delete_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_delete_response_proto"></a>
var File_gcommon_v1_database_delete_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_drop_database_request_proto"></a>
var File_gcommon_v1_database_drop_database_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_drop_schema_request_proto"></a>
var File_gcommon_v1_database_drop_schema_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_eviction_result_proto"></a>
var File_gcommon_v1_database_eviction_result_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_batch_request_proto"></a>
var File_gcommon_v1_database_execute_batch_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_batch_response_proto"></a>
var File_gcommon_v1_database_execute_batch_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_options_proto"></a>
var File_gcommon_v1_database_execute_options_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_request_proto"></a>
var File_gcommon_v1_database_execute_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_response_proto"></a>
var File_gcommon_v1_database_execute_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_execute_stats_proto"></a>
var File_gcommon_v1_database_execute_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_exists_request_proto"></a>
var File_gcommon_v1_database_exists_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_exists_response_proto"></a>
var File_gcommon_v1_database_exists_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_expire_request_proto"></a>
var File_gcommon_v1_database_expire_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_export_request_proto"></a>
var File_gcommon_v1_database_export_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_flush_request_proto"></a>
var File_gcommon_v1_database_flush_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_flush_response_proto"></a>
var File_gcommon_v1_database_flush_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_gc_request_proto"></a>
var File_gcommon_v1_database_gc_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_connection_info_request_proto"></a>
var File_gcommon_v1_database_get_connection_info_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_connection_info_response_proto"></a>
var File_gcommon_v1_database_get_connection_info_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_database_info_request_proto"></a>
var File_gcommon_v1_database_get_database_info_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_database_info_response_proto"></a>
var File_gcommon_v1_database_get_database_info_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_memory_usage_request_proto"></a>
var File_gcommon_v1_database_get_memory_usage_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_memory_usage_response_proto"></a>
var File_gcommon_v1_database_get_memory_usage_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_migration_status_request_proto"></a>
var File_gcommon_v1_database_get_migration_status_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_migration_status_response_proto"></a>
var File_gcommon_v1_database_get_migration_status_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_multiple_request_proto"></a>
var File_gcommon_v1_database_get_multiple_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_multiple_response_proto"></a>
var File_gcommon_v1_database_get_multiple_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_namespace_stats_request_proto"></a>
var File_gcommon_v1_database_get_namespace_stats_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_namespace_stats_response_proto"></a>
var File_gcommon_v1_database_get_namespace_stats_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_request_proto"></a>
var File_gcommon_v1_database_get_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_response_proto"></a>
var File_gcommon_v1_database_get_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_stats_request_proto"></a>
var File_gcommon_v1_database_get_stats_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_get_stats_response_proto"></a>
var File_gcommon_v1_database_get_stats_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_health_check_request_proto"></a>
var File_gcommon_v1_database_health_check_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_health_check_response_proto"></a>
var File_gcommon_v1_database_health_check_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_import_request_proto"></a>
var File_gcommon_v1_database_import_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_increment_request_proto"></a>
var File_gcommon_v1_database_increment_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_increment_response_proto"></a>
var File_gcommon_v1_database_increment_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_info_request_proto"></a>
var File_gcommon_v1_database_info_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_keys_request_proto"></a>
var File_gcommon_v1_database_keys_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_keys_response_proto"></a>
var File_gcommon_v1_database_keys_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_databases_request_proto"></a>
var File_gcommon_v1_database_list_databases_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_databases_response_proto"></a>
var File_gcommon_v1_database_list_databases_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_migrations_request_proto"></a>
var File_gcommon_v1_database_list_migrations_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_migrations_response_proto"></a>
var File_gcommon_v1_database_list_migrations_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_namespaces_request_proto"></a>
var File_gcommon_v1_database_list_namespaces_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_namespaces_response_proto"></a>
var File_gcommon_v1_database_list_namespaces_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_schemas_request_proto"></a>
var File_gcommon_v1_database_list_schemas_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_schemas_response_proto"></a>
var File_gcommon_v1_database_list_schemas_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_list_subscriptions_request_proto"></a>
var File_gcommon_v1_database_list_subscriptions_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_lock_request_proto"></a>
var File_gcommon_v1_database_lock_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_m_get_request_proto"></a>
var File_gcommon_v1_database_m_get_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_migration_info_proto"></a>
var File_gcommon_v1_database_migration_info_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_migration_script_proto"></a>
var File_gcommon_v1_database_migration_script_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_migration_service_proto"></a>
var File_gcommon_v1_database_migration_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_my_sql_config_proto"></a>
var File_gcommon_v1_database_my_sql_config_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_my_sql_status_proto"></a>
var File_gcommon_v1_database_my_sql_status_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_namespace_info_proto"></a>
var File_gcommon_v1_database_namespace_info_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_namespace_stats_proto"></a>
var File_gcommon_v1_database_namespace_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_optimize_request_proto"></a>
var File_gcommon_v1_database_optimize_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_pebble_config_proto"></a>
var File_gcommon_v1_database_pebble_config_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_pipeline_request_proto"></a>
var File_gcommon_v1_database_pipeline_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_pool_stats_proto"></a>
var File_gcommon_v1_database_pool_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_prepend_request_proto"></a>
var File_gcommon_v1_database_prepend_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_publish_request_proto"></a>
var File_gcommon_v1_database_publish_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_options_proto"></a>
var File_gcommon_v1_database_query_options_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_parameter_proto"></a>
var File_gcommon_v1_database_query_parameter_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_request_proto"></a>
var File_gcommon_v1_database_query_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_response_proto"></a>
var File_gcommon_v1_database_query_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_row_request_proto"></a>
var File_gcommon_v1_database_query_row_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_row_response_proto"></a>
var File_gcommon_v1_database_query_row_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_query_stats_proto"></a>
var File_gcommon_v1_database_query_stats_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_restore_request_proto"></a>
var File_gcommon_v1_database_restore_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_result_set_proto"></a>
var File_gcommon_v1_database_result_set_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_revert_migration_request_proto"></a>
var File_gcommon_v1_database_revert_migration_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_revert_migration_response_proto"></a>
var File_gcommon_v1_database_revert_migration_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_rollback_transaction_request_proto"></a>
var File_gcommon_v1_database_rollback_transaction_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_row_proto"></a>
var File_gcommon_v1_database_row_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_run_migration_request_proto"></a>
var File_gcommon_v1_database_run_migration_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_run_migration_response_proto"></a>
var File_gcommon_v1_database_run_migration_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_scan_request_proto"></a>
var File_gcommon_v1_database_scan_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_set_multiple_request_proto"></a>
var File_gcommon_v1_database_set_multiple_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_set_multiple_response_proto"></a>
var File_gcommon_v1_database_set_multiple_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_set_options_proto"></a>
var File_gcommon_v1_database_set_options_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_set_request_proto"></a>
var File_gcommon_v1_database_set_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_set_response_proto"></a>
var File_gcommon_v1_database_set_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_subscribe_request_proto"></a>
var File_gcommon_v1_database_subscribe_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_touch_expiration_request_proto"></a>
var File_gcommon_v1_database_touch_expiration_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_touch_expiration_response_proto"></a>
var File_gcommon_v1_database_touch_expiration_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_transaction_options_proto"></a>
var File_gcommon_v1_database_transaction_options_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_transaction_request_proto"></a>
var File_gcommon_v1_database_transaction_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_transaction_service_proto"></a>
var File_gcommon_v1_database_transaction_service_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_transaction_status_request_proto"></a>
var File_gcommon_v1_database_transaction_status_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_transaction_status_response_proto"></a>
var File_gcommon_v1_database_transaction_status_response_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_unlock_request_proto"></a>
var File_gcommon_v1_database_unlock_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_unsubscribe_request_proto"></a>
var File_gcommon_v1_database_unsubscribe_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_unwatch_request_proto"></a>
var File_gcommon_v1_database_unwatch_request_proto protoreflect.FileDescriptor
<a name="File_gcommon_v1_database_watch_request_proto"></a>
var File_gcommon_v1_database_watch_request_proto protoreflect.FileDescriptor
<a name="MigrationService_ServiceDesc"></a>MigrationService\_ServiceDesc is the grpc.ServiceDesc for MigrationService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var MigrationService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.MigrationService",
    HandlerType: (*MigrationServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "ApplyMigration",
            Handler:    _MigrationService_ApplyMigration_Handler,
        },
        {
            MethodName: "RevertMigration",
            Handler:    _MigrationService_RevertMigration_Handler,
        },
        {
            MethodName: "GetMigrationStatus",
            Handler:    _MigrationService_GetMigrationStatus_Handler,
        },
        {
            MethodName: "ListMigrations",
            Handler:    _MigrationService_ListMigrations_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/migration_service.proto",
}
<a name="TransactionService_ServiceDesc"></a>TransactionService\_ServiceDesc is the grpc.ServiceDesc for TransactionService service. It's only intended for direct use with grpc.RegisterService, and not to be introspected or modified \(even as a copy\)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "gcommon.v1.database.TransactionService",
    HandlerType: (*TransactionServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "BeginTransaction",
            Handler:    _TransactionService_BeginTransaction_Handler,
        },
        {
            MethodName: "CommitTransaction",
            Handler:    _TransactionService_CommitTransaction_Handler,
        },
        {
            MethodName: "RollbackTransaction",
            Handler:    _TransactionService_RollbackTransaction_Handler,
        },
        {
            MethodName: "GetTransactionStatus",
            Handler:    _TransactionService_GetTransactionStatus_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "gcommon/v1/database/transaction_service.proto",
}
<a name="RegisterCacheAdminServiceServer"></a>
## func [RegisterCacheAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L159>)
func RegisterCacheAdminServiceServer(s grpc.ServiceRegistrar, srv CacheAdminServiceServer)

<a name="RegisterCacheServiceServer"></a>
## func [RegisterCacheServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L324>)
func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer)

<a name="RegisterDatabaseAdminServiceServer"></a>
## func [RegisterDatabaseAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L197>)
func RegisterDatabaseAdminServiceServer(s grpc.ServiceRegistrar, srv DatabaseAdminServiceServer)

<a name="RegisterDatabaseServiceServer"></a>
## func [RegisterDatabaseServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L178>)
func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer)

<a name="RegisterMigrationServiceServer"></a>
## func [RegisterMigrationServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L140>)
func RegisterMigrationServiceServer(s grpc.ServiceRegistrar, srv MigrationServiceServer)

<a name="RegisterTransactionServiceServer"></a>
## func [RegisterTransactionServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L141>)
func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer)
## type [AppendRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L29-L41>)

    // contains filtered or unexported fields
### func \(\*AppendRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L172>)

<a name="AppendRequest.ClearMetadata"></a>
### func \(\*AppendRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L187>)

<a name="AppendRequest.ClearNamespace"></a>
### func \(\*AppendRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L182>)

<a name="AppendRequest.ClearValue"></a>
### func \(\*AppendRequest\) [ClearValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L177>)

<a name="AppendRequest.GetKey"></a>
### func \(\*AppendRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L68>)

<a name="AppendRequest.GetMetadata"></a>
### func \(\*AppendRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L102>)

<a name="AppendRequest.GetNamespace"></a>
### func \(\*AppendRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L92>)

<a name="AppendRequest.GetValue"></a>
### func \(\*AppendRequest\) [GetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L78>)

<a name="AppendRequest.HasKey"></a>
### func \(\*AppendRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L144>)

<a name="AppendRequest.HasMetadata"></a>
### func \(\*AppendRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L165>)

<a name="AppendRequest.HasNamespace"></a>
### func \(\*AppendRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L158>)

<a name="AppendRequest.HasValue"></a>
### func \(\*AppendRequest\) [HasValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L151>)

<a name="AppendRequest.ProtoMessage"></a>
### func \(\*AppendRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L54>)

<a name="AppendRequest.ProtoReflect"></a>
### func \(\*AppendRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L56>)

<a name="AppendRequest.Reset"></a>
### func \(\*AppendRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L43>)

<a name="AppendRequest.SetKey"></a>
### func \(\*AppendRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L116>)

<a name="AppendRequest.SetMetadata"></a>
### func \(\*AppendRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L135>)

<a name="AppendRequest.SetNamespace"></a>
### func \(\*AppendRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L130>)

<a name="AppendRequest.SetValue"></a>
### func \(\*AppendRequest\) [SetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L121>)

<a name="AppendRequest.String"></a>
### func \(\*AppendRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L50>)


## type [AppendRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L192-L203>)

    // contains filtered or unexported fields
### func \(AppendRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/append_request.pb.go#L205>)

<a name="BackupRequest"></a>
## type [BackupRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*BackupRequest\) [ClearDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L140>)

<a name="BackupRequest.ClearMetadata"></a>
### func \(\*BackupRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L150>)

<a name="BackupRequest.ClearNamespace"></a>
### func \(\*BackupRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L145>)

<a name="BackupRequest.GetDestination"></a>
### func \(\*BackupRequest\) [GetDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L66>)

<a name="BackupRequest.GetMetadata"></a>
### func \(\*BackupRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L86>)

<a name="BackupRequest.GetNamespace"></a>
### func \(\*BackupRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L76>)

<a name="BackupRequest.HasDestination"></a>
### func \(\*BackupRequest\) [HasDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L119>)

<a name="BackupRequest.HasMetadata"></a>
### func \(\*BackupRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L133>)

<a name="BackupRequest.HasNamespace"></a>
### func \(\*BackupRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L126>)

<a name="BackupRequest.ProtoMessage"></a>
### func \(\*BackupRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L52>)

<a name="BackupRequest.ProtoReflect"></a>
### func \(\*BackupRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L54>)

<a name="BackupRequest.Reset"></a>
### func \(\*BackupRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L41>)

<a name="BackupRequest.SetDestination"></a>
### func \(\*BackupRequest\) [SetDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L100>)

<a name="BackupRequest.SetMetadata"></a>
### func \(\*BackupRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L110>)

<a name="BackupRequest.SetNamespace"></a>
### func \(\*BackupRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L105>)

<a name="BackupRequest.String"></a>
### func \(\*BackupRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L48>)


## type [BackupRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L155-L164>)

    // contains filtered or unexported fields
### func \(BackupRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/backup_request.pb.go#L166>)
<a name="BatchExecuteOptions"></a>
## type [BatchExecuteOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L29-L40>)

\* BatchExecuteOptions configures behavior for batch database operations. Controls error handling, timeouts, and parallelism for batch execution.

    // contains filtered or unexported fields
### func \(\*BatchExecuteOptions\) [ClearFailFast](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L135>)

<a name="BatchExecuteOptions.ClearMaxParallel"></a>
### func \(\*BatchExecuteOptions\) [ClearMaxParallel](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L145>)

<a name="BatchExecuteOptions.ClearTimeout"></a>
### func \(\*BatchExecuteOptions\) [ClearTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L140>)

<a name="BatchExecuteOptions.GetFailFast"></a>
### func \(\*BatchExecuteOptions\) [GetFailFast](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L67>)

<a name="BatchExecuteOptions.GetMaxParallel"></a>
### func \(\*BatchExecuteOptions\) [GetMaxParallel](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L88>)

<a name="BatchExecuteOptions.GetTimeout"></a>
### func \(\*BatchExecuteOptions\) [GetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L74>)

<a name="BatchExecuteOptions.HasFailFast"></a>
### func \(\*BatchExecuteOptions\) [HasFailFast](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L114>)

<a name="BatchExecuteOptions.HasMaxParallel"></a>
### func \(\*BatchExecuteOptions\) [HasMaxParallel](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L128>)

<a name="BatchExecuteOptions.HasTimeout"></a>
### func \(\*BatchExecuteOptions\) [HasTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L121>)

<a name="BatchExecuteOptions.ProtoMessage"></a>
### func \(\*BatchExecuteOptions\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L53>)

<a name="BatchExecuteOptions.ProtoReflect"></a>
### func \(\*BatchExecuteOptions\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L55>)

<a name="BatchExecuteOptions.Reset"></a>
### func \(\*BatchExecuteOptions\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L42>)

<a name="BatchExecuteOptions.SetFailFast"></a>
### func \(\*BatchExecuteOptions\) [SetFailFast](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L95>)

<a name="BatchExecuteOptions.SetMaxParallel"></a>
### func \(\*BatchExecuteOptions\) [SetMaxParallel](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L109>)

<a name="BatchExecuteOptions.SetTimeout"></a>
### func \(\*BatchExecuteOptions\) [SetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L100>)

<a name="BatchExecuteOptions.String"></a>
### func \(\*BatchExecuteOptions\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L49>)


## type [BatchExecuteOptions\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L150-L159>)

    // contains filtered or unexported fields
### func \(BatchExecuteOptions\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_execute_options.pb.go#L161>)
<a name="BatchOperationResult"></a>
## type [BatchOperationResult](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L30-L42>)

\* BatchOperationResult contains the result of a single operation in a batch. Provides success status, affected rows, and error information.

    // contains filtered or unexported fields
### func \(\*BatchOperationResult\) [ClearAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L167>)

<a name="BatchOperationResult.ClearError"></a>
### func \(\*BatchOperationResult\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L172>)

<a name="BatchOperationResult.ClearSuccess"></a>
### func \(\*BatchOperationResult\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L162>)

<a name="BatchOperationResult.GetAffectedRows"></a>
### func \(\*BatchOperationResult\) [GetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L76>)

<a name="BatchOperationResult.GetError"></a>
### func \(\*BatchOperationResult\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L97>)

<a name="BatchOperationResult.GetGeneratedKeys"></a>
### func \(\*BatchOperationResult\) [GetGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L83>)

<a name="BatchOperationResult.GetSuccess"></a>
### func \(\*BatchOperationResult\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L69>)

<a name="BatchOperationResult.HasAffectedRows"></a>
### func \(\*BatchOperationResult\) [HasAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L148>)

<a name="BatchOperationResult.HasError"></a>
### func \(\*BatchOperationResult\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L155>)

<a name="BatchOperationResult.HasSuccess"></a>
### func \(\*BatchOperationResult\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L141>)

<a name="BatchOperationResult.ProtoMessage"></a>
### func \(\*BatchOperationResult\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L55>)

<a name="BatchOperationResult.ProtoReflect"></a>
### func \(\*BatchOperationResult\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L57>)

<a name="BatchOperationResult.Reset"></a>
### func \(\*BatchOperationResult\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L44>)

<a name="BatchOperationResult.SetAffectedRows"></a>
### func \(\*BatchOperationResult\) [SetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L116>)

<a name="BatchOperationResult.SetError"></a>
### func \(\*BatchOperationResult\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L132>)

<a name="BatchOperationResult.SetGeneratedKeys"></a>
### func \(\*BatchOperationResult\) [SetGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L121>)

<a name="BatchOperationResult.SetSuccess"></a>
### func \(\*BatchOperationResult\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L111>)

<a name="BatchOperationResult.String"></a>
### func \(\*BatchOperationResult\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L51>)


## type [BatchOperationResult\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L177-L188>)

    // contains filtered or unexported fields
### func \(BatchOperationResult\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation_result.pb.go#L190>)


## type [BeginTransactionRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L26-L37>)


    // contains filtered or unexported fields
### func \(\*BeginTransactionRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L146>)

<a name="BeginTransactionRequest.ClearMetadata"></a>
### func \(\*BeginTransactionRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L156>)

<a name="BeginTransactionRequest.ClearOptions"></a>
### func \(\*BeginTransactionRequest\) [ClearOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L151>)

<a name="BeginTransactionRequest.GetDatabase"></a>
### func \(\*BeginTransactionRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L64>)

<a name="BeginTransactionRequest.GetMetadata"></a>
### func \(\*BeginTransactionRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L88>)

<a name="BeginTransactionRequest.GetOptions"></a>
### func \(\*BeginTransactionRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L74>)

<a name="BeginTransactionRequest.HasDatabase"></a>
### func \(\*BeginTransactionRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L125>)

<a name="BeginTransactionRequest.HasMetadata"></a>
### func \(\*BeginTransactionRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L139>)

<a name="BeginTransactionRequest.HasOptions"></a>
### func \(\*BeginTransactionRequest\) [HasOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L132>)

<a name="BeginTransactionRequest.ProtoMessage"></a>
### func \(\*BeginTransactionRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L50>)


### func \(\*BeginTransactionRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L52>)

<a name="BeginTransactionRequest.Reset"></a>
### func \(\*BeginTransactionRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L39>)

<a name="BeginTransactionRequest.SetDatabase"></a>
### func \(\*BeginTransactionRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L102>)

<a name="BeginTransactionRequest.SetMetadata"></a>
### func \(\*BeginTransactionRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L116>)

<a name="BeginTransactionRequest.SetOptions"></a>
### func \(\*BeginTransactionRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L107>)

<a name="BeginTransactionRequest.String"></a>
### func \(\*BeginTransactionRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L46>)


## type [BeginTransactionRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L161-L170>)

    // contains filtered or unexported fields
### func \(BeginTransactionRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_request.pb.go#L172>)
<a name="BeginTransactionResponse"></a>
## type [BeginTransactionResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L29-L39>)

\* BeginTransactionResponse contains the result of starting a new transaction. Provides the transaction ID and timestamp for tracking.

    // contains filtered or unexported fields
### func \(\*BeginTransactionResponse\) [ClearStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L123>)

<a name="BeginTransactionResponse.ClearTransactionId"></a>
### func \(\*BeginTransactionResponse\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L118>)

<a name="BeginTransactionResponse.GetStartedAt"></a>
### func \(\*BeginTransactionResponse\) [GetStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L76>)

<a name="BeginTransactionResponse.GetTransactionId"></a>
### func \(\*BeginTransactionResponse\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L66>)

<a name="BeginTransactionResponse.HasStartedAt"></a>
### func \(\*BeginTransactionResponse\) [HasStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L111>)

<a name="BeginTransactionResponse.HasTransactionId"></a>
### func \(\*BeginTransactionResponse\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L104>)

<a name="BeginTransactionResponse.ProtoMessage"></a>
### func \(\*BeginTransactionResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L52>)

<a name="BeginTransactionResponse.ProtoReflect"></a>
### func \(\*BeginTransactionResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L54>)

<a name="BeginTransactionResponse.Reset"></a>
### func \(\*BeginTransactionResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L41>)

<a name="BeginTransactionResponse.SetStartedAt"></a>
### func \(\*BeginTransactionResponse\) [SetStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L95>)

<a name="BeginTransactionResponse.SetTransactionId"></a>
### func \(\*BeginTransactionResponse\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L90>)

<a name="BeginTransactionResponse.String"></a>
### func \(\*BeginTransactionResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L48>)


## type [BeginTransactionResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L128-L135>)

    // contains filtered or unexported fields
### func \(BeginTransactionResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/begin_transaction_response.pb.go#L137>)

<a name="CacheAdminServiceClient"></a>
## type [CacheAdminServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L36-L47>)
For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
### func [NewCacheAdminServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L53>)
<a name="CacheAdminServiceServer"></a>
## type [CacheAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L113-L125>)

CacheAdminServiceServer is the server API for CacheAdminService service. All implementations must embed UnimplementedCacheAdminServiceServer for forward compatibility.
    // contains filtered or unexported methods
## type [CacheCacheConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L30-L44>)
\* Configuration settings for cache behavior. Defines cache policies, limits, and operational parameters.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheCacheConfig\) [ClearDefaultTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L240>)

<a name="CacheCacheConfig.ClearEnablePersistence"></a>
### func \(\*CacheCacheConfig\) [ClearEnablePersistence](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L254>)

<a name="CacheCacheConfig.ClearEnableStats"></a>
### func \(\*CacheCacheConfig\) [ClearEnableStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L249>)

<a name="CacheCacheConfig.ClearEvictionPolicy"></a>
### func \(\*CacheCacheConfig\) [ClearEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L244>)

<a name="CacheCacheConfig.ClearMaxEntries"></a>
### func \(\*CacheCacheConfig\) [ClearMaxEntries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L230>)

<a name="CacheCacheConfig.ClearMaxMemoryBytes"></a>
### func \(\*CacheCacheConfig\) [ClearMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L235>)

<a name="CacheCacheConfig.ClearName"></a>
### func \(\*CacheCacheConfig\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L264>)

<a name="CacheCacheConfig.ClearPersistenceFile"></a>
### func \(\*CacheCacheConfig\) [ClearPersistenceFile](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L259>)

<a name="CacheCacheConfig.GetDefaultTtl"></a>
### func \(\*CacheCacheConfig\) [GetDefaultTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L85>)

<a name="CacheCacheConfig.GetEnablePersistence"></a>
### func \(\*CacheCacheConfig\) [GetEnablePersistence](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L108>)

<a name="CacheCacheConfig.GetEnableStats"></a>
### func \(\*CacheCacheConfig\) [GetEnableStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L101>)

<a name="CacheCacheConfig.GetEvictionPolicy"></a>
### func \(\*CacheCacheConfig\) [GetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L92>)

<a name="CacheCacheConfig.GetMaxEntries"></a>
### func \(\*CacheCacheConfig\) [GetMaxEntries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L71>)

<a name="CacheCacheConfig.GetMaxMemoryBytes"></a>
### func \(\*CacheCacheConfig\) [GetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L78>)

<a name="CacheCacheConfig.GetName"></a>
### func \(\*CacheCacheConfig\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L125>)

<a name="CacheCacheConfig.GetPersistenceFile"></a>
### func \(\*CacheCacheConfig\) [GetPersistenceFile](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L115>)

<a name="CacheCacheConfig.HasDefaultTtl"></a>
### func \(\*CacheCacheConfig\) [HasDefaultTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L188>)

<a name="CacheCacheConfig.HasEnablePersistence"></a>
### func \(\*CacheCacheConfig\) [HasEnablePersistence](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L209>)

<a name="CacheCacheConfig.HasEnableStats"></a>
### func \(\*CacheCacheConfig\) [HasEnableStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L202>)

<a name="CacheCacheConfig.HasEvictionPolicy"></a>
### func \(\*CacheCacheConfig\) [HasEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L195>)

<a name="CacheCacheConfig.HasMaxEntries"></a>
### func \(\*CacheCacheConfig\) [HasMaxEntries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L174>)

<a name="CacheCacheConfig.HasMaxMemoryBytes"></a>
### func \(\*CacheCacheConfig\) [HasMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L181>)

<a name="CacheCacheConfig.HasName"></a>
### func \(\*CacheCacheConfig\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L223>)

<a name="CacheCacheConfig.HasPersistenceFile"></a>
### func \(\*CacheCacheConfig\) [HasPersistenceFile](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L216>)

<a name="CacheCacheConfig.ProtoMessage"></a>
### func \(\*CacheCacheConfig\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L57>)

<a name="CacheCacheConfig.ProtoReflect"></a>
### func \(\*CacheCacheConfig\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L59>)

<a name="CacheCacheConfig.Reset"></a>
### func \(\*CacheCacheConfig\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L46>)

<a name="CacheCacheConfig.SetDefaultTtl"></a>
### func \(\*CacheCacheConfig\) [SetDefaultTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L145>)

<a name="CacheCacheConfig.SetEnablePersistence"></a>
### func \(\*CacheCacheConfig\) [SetEnablePersistence](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L159>)

<a name="CacheCacheConfig.SetEnableStats"></a>
### func \(\*CacheCacheConfig\) [SetEnableStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L154>)

<a name="CacheCacheConfig.SetEvictionPolicy"></a>
### func \(\*CacheCacheConfig\) [SetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L149>)

<a name="CacheCacheConfig.SetMaxEntries"></a>
### func \(\*CacheCacheConfig\) [SetMaxEntries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L135>)

<a name="CacheCacheConfig.SetMaxMemoryBytes"></a>
### func \(\*CacheCacheConfig\) [SetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L140>)

<a name="CacheCacheConfig.SetName"></a>
### func \(\*CacheCacheConfig\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L169>)

<a name="CacheCacheConfig.SetPersistenceFile"></a>
### func \(\*CacheCacheConfig\) [SetPersistenceFile](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L164>)

<a name="CacheCacheConfig.String"></a>
### func \(\*CacheCacheConfig\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L53>)


## type [CacheCacheConfig\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L269-L288>)

    // contains filtered or unexported fields
### func \(CacheCacheConfig\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_config.pb.go#L290>)

<a name="CacheDeleteRequest"></a>
## type [CacheDeleteRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*CacheDeleteRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L140>)

<a name="CacheDeleteRequest.ClearMetadata"></a>
### func \(\*CacheDeleteRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L150>)

<a name="CacheDeleteRequest.ClearNamespace"></a>
### func \(\*CacheDeleteRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L145>)

<a name="CacheDeleteRequest.GetKey"></a>
### func \(\*CacheDeleteRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L66>)

<a name="CacheDeleteRequest.GetMetadata"></a>
### func \(\*CacheDeleteRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L86>)

<a name="CacheDeleteRequest.GetNamespace"></a>
### func \(\*CacheDeleteRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L76>)

<a name="CacheDeleteRequest.HasKey"></a>
### func \(\*CacheDeleteRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L119>)

<a name="CacheDeleteRequest.HasMetadata"></a>
### func \(\*CacheDeleteRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L133>)

<a name="CacheDeleteRequest.HasNamespace"></a>
### func \(\*CacheDeleteRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L126>)

<a name="CacheDeleteRequest.ProtoMessage"></a>
### func \(\*CacheDeleteRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L52>)

<a name="CacheDeleteRequest.ProtoReflect"></a>
### func \(\*CacheDeleteRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L54>)

<a name="CacheDeleteRequest.Reset"></a>
### func \(\*CacheDeleteRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L41>)

<a name="CacheDeleteRequest.SetKey"></a>
### func \(\*CacheDeleteRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L100>)

<a name="CacheDeleteRequest.SetMetadata"></a>
### func \(\*CacheDeleteRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L110>)

<a name="CacheDeleteRequest.SetNamespace"></a>
### func \(\*CacheDeleteRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L105>)

<a name="CacheDeleteRequest.String"></a>
### func \(\*CacheDeleteRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L48>)


## type [CacheDeleteRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L155-L164>)

    // contains filtered or unexported fields
### func \(CacheDeleteRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_request.pb.go#L166>)
<a name="CacheDeleteResponse"></a>
## type [CacheDeleteResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L29-L38>)

\* Response for cache delete operations. Indicates success/failure of key deletion.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheDeleteResponse\) [ClearDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L130>)

<a name="CacheDeleteResponse.ClearError"></a>
### func \(\*CacheDeleteResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L126>)

<a name="CacheDeleteResponse.ClearSuccess"></a>
### func \(\*CacheDeleteResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L121>)

<a name="CacheDeleteResponse.GetDeletedCount"></a>
### func \(\*CacheDeleteResponse\) [GetDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L79>)

<a name="CacheDeleteResponse.GetError"></a>
### func \(\*CacheDeleteResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L72>)


### func \(\*CacheDeleteResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L65>)

<a name="CacheDeleteResponse.HasDeletedCount"></a>
### func \(\*CacheDeleteResponse\) [HasDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L114>)

<a name="CacheDeleteResponse.HasError"></a>
### func \(\*CacheDeleteResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L107>)

<a name="CacheDeleteResponse.HasSuccess"></a>
### func \(\*CacheDeleteResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L100>)

<a name="CacheDeleteResponse.ProtoMessage"></a>
### func \(\*CacheDeleteResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L51>)

<a name="CacheDeleteResponse.ProtoReflect"></a>
### func \(\*CacheDeleteResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L53>)

<a name="CacheDeleteResponse.Reset"></a>
### func \(\*CacheDeleteResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L40>)

<a name="CacheDeleteResponse.SetDeletedCount"></a>
### func \(\*CacheDeleteResponse\) [SetDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L95>)

<a name="CacheDeleteResponse.SetError"></a>
### func \(\*CacheDeleteResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L91>)

<a name="CacheDeleteResponse.SetSuccess"></a>
### func \(\*CacheDeleteResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L86>)

<a name="CacheDeleteResponse.String"></a>
### func \(\*CacheDeleteResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L47>)


## type [CacheDeleteResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L135-L144>)

    // contains filtered or unexported fields
### func \(CacheDeleteResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_response.pb.go#L146>)
<a name="CacheEntry"></a>
## type [CacheEntry](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L31-L48>)

\* Cache entry containing the cached value and metadata. Supports multiple value types with comprehensive expiration and access tracking for cache policies.

    // contains filtered or unexported fields
### func \(\*CacheEntry\) [ClearAccessCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L313>)

<a name="CacheEntry.ClearCreatedAt"></a>
### func \(\*CacheEntry\) [ClearCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L298>)

<a name="CacheEntry.ClearExpiresAt"></a>
### func \(\*CacheEntry\) [ClearExpiresAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L308>)

<a name="CacheEntry.ClearKey"></a>
### func \(\*CacheEntry\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L288>)

<a name="CacheEntry.ClearLastAccessedAt"></a>
### func \(\*CacheEntry\) [ClearLastAccessedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L303>)

<a name="CacheEntry.ClearNamespace"></a>
### func \(\*CacheEntry\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L323>)

<a name="CacheEntry.ClearSizeBytes"></a>
### func \(\*CacheEntry\) [ClearSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L318>)

<a name="CacheEntry.ClearValue"></a>
### func \(\*CacheEntry\) [ClearValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L293>)

<a name="CacheEntry.GetAccessCount"></a>
### func \(\*CacheEntry\) [GetAccessCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L141>)

<a name="CacheEntry.GetCreatedAt"></a>
### func \(\*CacheEntry\) [GetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L99>)

<a name="CacheEntry.GetExpiresAt"></a>
### func \(\*CacheEntry\) [GetExpiresAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L127>)

<a name="CacheEntry.GetKey"></a>
### func \(\*CacheEntry\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L75>)

<a name="CacheEntry.GetLastAccessedAt"></a>
### func \(\*CacheEntry\) [GetLastAccessedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L113>)

<a name="CacheEntry.GetMetadata"></a>
### func \(\*CacheEntry\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L155>)

<a name="CacheEntry.GetNamespace"></a>
### func \(\*CacheEntry\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L162>)

<a name="CacheEntry.GetSizeBytes"></a>
### func \(\*CacheEntry\) [GetSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L148>)

<a name="CacheEntry.GetValue"></a>
### func \(\*CacheEntry\) [GetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L85>)

<a name="CacheEntry.HasAccessCount"></a>
### func \(\*CacheEntry\) [HasAccessCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L267>)

<a name="CacheEntry.HasCreatedAt"></a>
### func \(\*CacheEntry\) [HasCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L246>)

<a name="CacheEntry.HasExpiresAt"></a>
### func \(\*CacheEntry\) [HasExpiresAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L260>)

<a name="CacheEntry.HasKey"></a>
### func \(\*CacheEntry\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L232>)

<a name="CacheEntry.HasLastAccessedAt"></a>
### func \(\*CacheEntry\) [HasLastAccessedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L253>)

<a name="CacheEntry.HasNamespace"></a>
### func \(\*CacheEntry\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L281>)

<a name="CacheEntry.HasSizeBytes"></a>
### func \(\*CacheEntry\) [HasSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L274>)

<a name="CacheEntry.HasValue"></a>
### func \(\*CacheEntry\) [HasValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L239>)

<a name="CacheEntry.ProtoMessage"></a>
### func \(\*CacheEntry\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L61>)

<a name="CacheEntry.ProtoReflect"></a>
### func \(\*CacheEntry\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L63>)

<a name="CacheEntry.Reset"></a>
### func \(\*CacheEntry\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L50>)

<a name="CacheEntry.SetAccessCount"></a>
### func \(\*CacheEntry\) [SetAccessCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L213>)

<a name="CacheEntry.SetCreatedAt"></a>
### func \(\*CacheEntry\) [SetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L186>)

<a name="CacheEntry.SetExpiresAt"></a>
### func \(\*CacheEntry\) [SetExpiresAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L204>)

<a name="CacheEntry.SetKey"></a>
### func \(\*CacheEntry\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L172>)

<a name="CacheEntry.SetLastAccessedAt"></a>
### func \(\*CacheEntry\) [SetLastAccessedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L195>)

<a name="CacheEntry.SetMetadata"></a>
### func \(\*CacheEntry\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L223>)

<a name="CacheEntry.SetNamespace"></a>
### func \(\*CacheEntry\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L227>)

<a name="CacheEntry.SetSizeBytes"></a>
### func \(\*CacheEntry\) [SetSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L218>)

<a name="CacheEntry.SetValue"></a>
### func \(\*CacheEntry\) [SetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L177>)

<a name="CacheEntry.String"></a>
### func \(\*CacheEntry\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L57>)


## type [CacheEntry\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L328-L349>)

    // contains filtered or unexported fields
### func \(CacheEntry\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_entry.pb.go#L351>)

<a name="CacheGetStatsRequest"></a>
## type [CacheGetStatsRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*CacheGetStatsRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L122>)

<a name="CacheGetStatsRequest.ClearNamespace"></a>
### func \(\*CacheGetStatsRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L117>)

<a name="CacheGetStatsRequest.GetMetadata"></a>
### func \(\*CacheGetStatsRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L75>)

<a name="CacheGetStatsRequest.GetNamespace"></a>
### func \(\*CacheGetStatsRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L65>)

<a name="CacheGetStatsRequest.HasMetadata"></a>
### func \(\*CacheGetStatsRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L110>)

<a name="CacheGetStatsRequest.HasNamespace"></a>
### func \(\*CacheGetStatsRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L103>)

<a name="CacheGetStatsRequest.ProtoMessage"></a>
### func \(\*CacheGetStatsRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L51>)

<a name="CacheGetStatsRequest.ProtoReflect"></a>
### func \(\*CacheGetStatsRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L53>)

<a name="CacheGetStatsRequest.Reset"></a>
### func \(\*CacheGetStatsRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L40>)

<a name="CacheGetStatsRequest.SetMetadata"></a>
### func \(\*CacheGetStatsRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L94>)

<a name="CacheGetStatsRequest.SetNamespace"></a>
### func \(\*CacheGetStatsRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L89>)

<a name="CacheGetStatsRequest.String"></a>
### func \(\*CacheGetStatsRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L47>)


## type [CacheGetStatsRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(CacheGetStatsRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_request.pb.go#L136>)
<a name="CacheGetStatsResponse"></a>
## type [CacheGetStatsResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L29-L44>)

\* Response for cache statistics operations. Provides detailed cache performance metrics.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheGetStatsResponse\) [ClearCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L246>)

<a name="CacheGetStatsResponse.ClearCacheMisses"></a>
### func \(\*CacheGetStatsResponse\) [ClearCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L251>)

<a name="CacheGetStatsResponse.ClearError"></a>
### func \(\*CacheGetStatsResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L281>)

<a name="CacheGetStatsResponse.ClearEvictedItems"></a>
### func \(\*CacheGetStatsResponse\) [ClearEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L271>)

<a name="CacheGetStatsResponse.ClearHitRatio"></a>
### func \(\*CacheGetStatsResponse\) [ClearHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L256>)

<a name="CacheGetStatsResponse.ClearMemoryLimit"></a>
### func \(\*CacheGetStatsResponse\) [ClearMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L266>)

<a name="CacheGetStatsResponse.ClearMemoryUsage"></a>
### func \(\*CacheGetStatsResponse\) [ClearMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L261>)

<a name="CacheGetStatsResponse.ClearSuccess"></a>
### func \(\*CacheGetStatsResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L276>)

<a name="CacheGetStatsResponse.ClearTotalItems"></a>
### func \(\*CacheGetStatsResponse\) [ClearTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L241>)

<a name="CacheGetStatsResponse.GetCacheHits"></a>
### func \(\*CacheGetStatsResponse\) [GetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L78>)

<a name="CacheGetStatsResponse.GetCacheMisses"></a>
### func \(\*CacheGetStatsResponse\) [GetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L85>)

<a name="CacheGetStatsResponse.GetError"></a>
### func \(\*CacheGetStatsResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L127>)

<a name="CacheGetStatsResponse.GetEvictedItems"></a>
### func \(\*CacheGetStatsResponse\) [GetEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L113>)

<a name="CacheGetStatsResponse.GetHitRatio"></a>
### func \(\*CacheGetStatsResponse\) [GetHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L92>)

<a name="CacheGetStatsResponse.GetMemoryLimit"></a>
### func \(\*CacheGetStatsResponse\) [GetMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L106>)

<a name="CacheGetStatsResponse.GetMemoryUsage"></a>
### func \(\*CacheGetStatsResponse\) [GetMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L99>)

<a name="CacheGetStatsResponse.GetSuccess"></a>
### func \(\*CacheGetStatsResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L120>)

<a name="CacheGetStatsResponse.GetTotalItems"></a>
### func \(\*CacheGetStatsResponse\) [GetTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L71>)


### func \(\*CacheGetStatsResponse\) [HasCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L185>)

<a name="CacheGetStatsResponse.HasCacheMisses"></a>
### func \(\*CacheGetStatsResponse\) [HasCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L192>)

<a name="CacheGetStatsResponse.HasError"></a>
### func \(\*CacheGetStatsResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L234>)

<a name="CacheGetStatsResponse.HasEvictedItems"></a>
### func \(\*CacheGetStatsResponse\) [HasEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L220>)

<a name="CacheGetStatsResponse.HasHitRatio"></a>
### func \(\*CacheGetStatsResponse\) [HasHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L199>)

<a name="CacheGetStatsResponse.HasMemoryLimit"></a>
### func \(\*CacheGetStatsResponse\) [HasMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L213>)

<a name="CacheGetStatsResponse.HasMemoryUsage"></a>
### func \(\*CacheGetStatsResponse\) [HasMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L206>)

<a name="CacheGetStatsResponse.HasSuccess"></a>
### func \(\*CacheGetStatsResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L227>)

<a name="CacheGetStatsResponse.HasTotalItems"></a>
### func \(\*CacheGetStatsResponse\) [HasTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L178>)

<a name="CacheGetStatsResponse.ProtoMessage"></a>
### func \(\*CacheGetStatsResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L57>)

<a name="CacheGetStatsResponse.ProtoReflect"></a>
### func \(\*CacheGetStatsResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L59>)

<a name="CacheGetStatsResponse.Reset"></a>
### func \(\*CacheGetStatsResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L46>)

<a name="CacheGetStatsResponse.SetCacheHits"></a>
### func \(\*CacheGetStatsResponse\) [SetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L139>)

<a name="CacheGetStatsResponse.SetCacheMisses"></a>
### func \(\*CacheGetStatsResponse\) [SetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L144>)

<a name="CacheGetStatsResponse.SetError"></a>
### func \(\*CacheGetStatsResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L174>)

<a name="CacheGetStatsResponse.SetEvictedItems"></a>
### func \(\*CacheGetStatsResponse\) [SetEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L164>)

<a name="CacheGetStatsResponse.SetHitRatio"></a>
### func \(\*CacheGetStatsResponse\) [SetHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L149>)

<a name="CacheGetStatsResponse.SetMemoryLimit"></a>
### func \(\*CacheGetStatsResponse\) [SetMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L159>)

<a name="CacheGetStatsResponse.SetMemoryUsage"></a>
### func \(\*CacheGetStatsResponse\) [SetMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L154>)

<a name="CacheGetStatsResponse.SetSuccess"></a>
### func \(\*CacheGetStatsResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L169>)

<a name="CacheGetStatsResponse.SetTotalItems"></a>
### func \(\*CacheGetStatsResponse\) [SetTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L134>)

<a name="CacheGetStatsResponse.String"></a>
### func \(\*CacheGetStatsResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L53>)


## type [CacheGetStatsResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L285-L306>)

    // contains filtered or unexported fields
### func \(CacheGetStatsResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_stats_response.pb.go#L308>)
<a name="CacheInfo"></a>
## type [CacheInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L30-L45>)

\* General cache information and metadata. Provides cache instance details and operational status.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheInfo\) [ClearCacheType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L260>)

<a name="CacheInfo.ClearCreatedAt"></a>
### func \(\*CacheInfo\) [ClearCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L270>)

<a name="CacheInfo.ClearDescription"></a>
### func \(\*CacheInfo\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L283>)

<a name="CacheInfo.ClearHealthStatus"></a>
### func \(\*CacheInfo\) [ClearHealthStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L265>)

<a name="CacheInfo.ClearInstanceId"></a>
### func \(\*CacheInfo\) [ClearInstanceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L278>)

<a name="CacheInfo.ClearLastAccessed"></a>
### func \(\*CacheInfo\) [ClearLastAccessed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L274>)

<a name="CacheInfo.ClearName"></a>
### func \(\*CacheInfo\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L250>)

<a name="CacheInfo.ClearVersion"></a>
### func \(\*CacheInfo\) [ClearVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L255>)

<a name="CacheInfo.GetCacheType"></a>
### func \(\*CacheInfo\) [GetCacheType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L92>)

<a name="CacheInfo.GetCreatedAt"></a>
### func \(\*CacheInfo\) [GetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L111>)

<a name="CacheInfo.GetDescription"></a>
### func \(\*CacheInfo\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L135>)

<a name="CacheInfo.GetHealthStatus"></a>
### func \(\*CacheInfo\) [GetHealthStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L102>)

<a name="CacheInfo.GetInstanceId"></a>
### func \(\*CacheInfo\) [GetInstanceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L125>)

<a name="CacheInfo.GetLastAccessed"></a>
### func \(\*CacheInfo\) [GetLastAccessed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L118>)

<a name="CacheInfo.GetMetadata"></a>
### func \(\*CacheInfo\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L145>)

<a name="CacheInfo.GetName"></a>
### func \(\*CacheInfo\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L72>)

<a name="CacheInfo.GetVersion"></a>
### func \(\*CacheInfo\) [GetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L82>)

<a name="CacheInfo.HasCacheType"></a>
### func \(\*CacheInfo\) [HasCacheType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L208>)

<a name="CacheInfo.HasCreatedAt"></a>
### func \(\*CacheInfo\) [HasCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L222>)

<a name="CacheInfo.HasDescription"></a>
### func \(\*CacheInfo\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L243>)

<a name="CacheInfo.HasHealthStatus"></a>
### func \(\*CacheInfo\) [HasHealthStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L215>)

<a name="CacheInfo.HasInstanceId"></a>
### func \(\*CacheInfo\) [HasInstanceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L236>)

<a name="CacheInfo.HasLastAccessed"></a>
### func \(\*CacheInfo\) [HasLastAccessed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L229>)

<a name="CacheInfo.HasName"></a>
### func \(\*CacheInfo\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L194>)

<a name="CacheInfo.HasVersion"></a>
### func \(\*CacheInfo\) [HasVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L201>)

<a name="CacheInfo.ProtoMessage"></a>
### func \(\*CacheInfo\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L58>)

<a name="CacheInfo.ProtoReflect"></a>
### func \(\*CacheInfo\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L60>)

<a name="CacheInfo.Reset"></a>
### func \(\*CacheInfo\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L47>)

<a name="CacheInfo.SetCacheType"></a>
### func \(\*CacheInfo\) [SetCacheType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L162>)

<a name="CacheInfo.SetCreatedAt"></a>
### func \(\*CacheInfo\) [SetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L172>)

<a name="CacheInfo.SetDescription"></a>
### func \(\*CacheInfo\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L185>)

<a name="CacheInfo.SetHealthStatus"></a>
### func \(\*CacheInfo\) [SetHealthStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L167>)

<a name="CacheInfo.SetInstanceId"></a>
### func \(\*CacheInfo\) [SetInstanceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L180>)

<a name="CacheInfo.SetLastAccessed"></a>
### func \(\*CacheInfo\) [SetLastAccessed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L176>)

<a name="CacheInfo.SetMetadata"></a>
### func \(\*CacheInfo\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L190>)

<a name="CacheInfo.SetName"></a>
### func \(\*CacheInfo\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L152>)

<a name="CacheInfo.SetVersion"></a>
### func \(\*CacheInfo\) [SetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L157>)

<a name="CacheInfo.String"></a>
### func \(\*CacheInfo\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L54>)


## type [CacheInfo\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L288-L309>)

    // contains filtered or unexported fields
### func \(CacheInfo\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_info.pb.go#L311>)

<a name="CacheListSubscriptionsRequest"></a>
## type [CacheListSubscriptionsRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L27-L36>)

    // contains filtered or unexported fields
### func \(\*CacheListSubscriptionsRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L93>)

<a name="CacheListSubscriptionsRequest.GetMetadata"></a>
### func \(\*CacheListSubscriptionsRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L63>)

<a name="CacheListSubscriptionsRequest.HasMetadata"></a>
### func \(\*CacheListSubscriptionsRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L86>)

<a name="CacheListSubscriptionsRequest.ProtoMessage"></a>
### func \(\*CacheListSubscriptionsRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L49>)

<a name="CacheListSubscriptionsRequest.ProtoReflect"></a>
### func \(\*CacheListSubscriptionsRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L51>)

<a name="CacheListSubscriptionsRequest.Reset"></a>
### func \(\*CacheListSubscriptionsRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L38>)

<a name="CacheListSubscriptionsRequest.SetMetadata"></a>
### func \(\*CacheListSubscriptionsRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L77>)

<a name="CacheListSubscriptionsRequest.String"></a>
### func \(\*CacheListSubscriptionsRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L45>)


## type [CacheListSubscriptionsRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L98-L103>)

    // contains filtered or unexported fields
### func \(CacheListSubscriptionsRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_subscriptions_request.pb.go#L105>)
<a name="CacheMetrics"></a>
## type [CacheMetrics](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L30-L49>)

\* Detailed cache performance metrics. Provides comprehensive metrics for cache monitoring and optimization.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheMetrics\) [ClearActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L351>)

<a name="CacheMetrics.ClearAvgResponseTime"></a>
### func \(\*CacheMetrics\) [ClearAvgResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L334>)

<a name="CacheMetrics.ClearCollectedAt"></a>
### func \(\*CacheMetrics\) [ClearCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L376>)

<a name="CacheMetrics.ClearCpuUsagePercent"></a>
### func \(\*CacheMetrics\) [ClearCpuUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L366>)

<a name="CacheMetrics.ClearMemoryUsagePercent"></a>
### func \(\*CacheMetrics\) [ClearMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L371>)

<a name="CacheMetrics.ClearNetworkBytesIn"></a>
### func \(\*CacheMetrics\) [ClearNetworkBytesIn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L356>)

<a name="CacheMetrics.ClearNetworkBytesOut"></a>
### func \(\*CacheMetrics\) [ClearNetworkBytesOut](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L361>)

<a name="CacheMetrics.ClearOpsPerSecond"></a>
### func \(\*CacheMetrics\) [ClearOpsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L319>)

<a name="CacheMetrics.ClearP95ResponseTime"></a>
### func \(\*CacheMetrics\) [ClearP95ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L338>)

<a name="CacheMetrics.ClearP99ResponseTime"></a>
### func \(\*CacheMetrics\) [ClearP99ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L342>)

<a name="CacheMetrics.ClearReadsPerSecond"></a>
### func \(\*CacheMetrics\) [ClearReadsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L324>)

<a name="CacheMetrics.ClearTotalConnections"></a>
### func \(\*CacheMetrics\) [ClearTotalConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L346>)

<a name="CacheMetrics.ClearWritesPerSecond"></a>
### func \(\*CacheMetrics\) [ClearWritesPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L329>)

<a name="CacheMetrics.GetActiveConnections"></a>
### func \(\*CacheMetrics\) [GetActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L125>)

<a name="CacheMetrics.GetAvgResponseTime"></a>
### func \(\*CacheMetrics\) [GetAvgResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L97>)

<a name="CacheMetrics.GetCollectedAt"></a>
### func \(\*CacheMetrics\) [GetCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L160>)

<a name="CacheMetrics.GetCpuUsagePercent"></a>
### func \(\*CacheMetrics\) [GetCpuUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L146>)

<a name="CacheMetrics.GetMemoryUsagePercent"></a>
### func \(\*CacheMetrics\) [GetMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L153>)

<a name="CacheMetrics.GetNetworkBytesIn"></a>
### func \(\*CacheMetrics\) [GetNetworkBytesIn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L132>)

<a name="CacheMetrics.GetNetworkBytesOut"></a>
### func \(\*CacheMetrics\) [GetNetworkBytesOut](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L139>)

<a name="CacheMetrics.GetOpsPerSecond"></a>
### func \(\*CacheMetrics\) [GetOpsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L76>)

<a name="CacheMetrics.GetP95ResponseTime"></a>
### func \(\*CacheMetrics\) [GetP95ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L104>)

<a name="CacheMetrics.GetP99ResponseTime"></a>
### func \(\*CacheMetrics\) [GetP99ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L111>)

<a name="CacheMetrics.GetReadsPerSecond"></a>
### func \(\*CacheMetrics\) [GetReadsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L83>)

<a name="CacheMetrics.GetTotalConnections"></a>
### func \(\*CacheMetrics\) [GetTotalConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L118>)

<a name="CacheMetrics.GetWritesPerSecond"></a>
### func \(\*CacheMetrics\) [GetWritesPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L90>)

<a name="CacheMetrics.HasActiveConnections"></a>
### func \(\*CacheMetrics\) [HasActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L277>)

<a name="CacheMetrics.HasAvgResponseTime"></a>
### func \(\*CacheMetrics\) [HasAvgResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L249>)

<a name="CacheMetrics.HasCollectedAt"></a>
### func \(\*CacheMetrics\) [HasCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L312>)

<a name="CacheMetrics.HasCpuUsagePercent"></a>
### func \(\*CacheMetrics\) [HasCpuUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L298>)

<a name="CacheMetrics.HasMemoryUsagePercent"></a>
### func \(\*CacheMetrics\) [HasMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L305>)

<a name="CacheMetrics.HasNetworkBytesIn"></a>
### func \(\*CacheMetrics\) [HasNetworkBytesIn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L284>)

<a name="CacheMetrics.HasNetworkBytesOut"></a>
### func \(\*CacheMetrics\) [HasNetworkBytesOut](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L291>)

<a name="CacheMetrics.HasOpsPerSecond"></a>
### func \(\*CacheMetrics\) [HasOpsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L228>)

<a name="CacheMetrics.HasP95ResponseTime"></a>
### func \(\*CacheMetrics\) [HasP95ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L256>)

<a name="CacheMetrics.HasP99ResponseTime"></a>
### func \(\*CacheMetrics\) [HasP99ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L263>)

<a name="CacheMetrics.HasReadsPerSecond"></a>
### func \(\*CacheMetrics\) [HasReadsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L235>)

<a name="CacheMetrics.HasTotalConnections"></a>
### func \(\*CacheMetrics\) [HasTotalConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L270>)

<a name="CacheMetrics.HasWritesPerSecond"></a>
### func \(\*CacheMetrics\) [HasWritesPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L242>)

<a name="CacheMetrics.ProtoMessage"></a>
### func \(\*CacheMetrics\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L62>)

<a name="CacheMetrics.ProtoReflect"></a>
### func \(\*CacheMetrics\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L64>)

<a name="CacheMetrics.Reset"></a>
### func \(\*CacheMetrics\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L51>)

<a name="CacheMetrics.SetActiveConnections"></a>
### func \(\*CacheMetrics\) [SetActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L199>)

<a name="CacheMetrics.SetAvgResponseTime"></a>
### func \(\*CacheMetrics\) [SetAvgResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L182>)

<a name="CacheMetrics.SetCollectedAt"></a>
### func \(\*CacheMetrics\) [SetCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L224>)

<a name="CacheMetrics.SetCpuUsagePercent"></a>
### func \(\*CacheMetrics\) [SetCpuUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L214>)

<a name="CacheMetrics.SetMemoryUsagePercent"></a>
### func \(\*CacheMetrics\) [SetMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L219>)

<a name="CacheMetrics.SetNetworkBytesIn"></a>
### func \(\*CacheMetrics\) [SetNetworkBytesIn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L204>)

<a name="CacheMetrics.SetNetworkBytesOut"></a>
### func \(\*CacheMetrics\) [SetNetworkBytesOut](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L209>)

<a name="CacheMetrics.SetOpsPerSecond"></a>
### func \(\*CacheMetrics\) [SetOpsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L167>)

<a name="CacheMetrics.SetP95ResponseTime"></a>
### func \(\*CacheMetrics\) [SetP95ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L186>)

<a name="CacheMetrics.SetP99ResponseTime"></a>
### func \(\*CacheMetrics\) [SetP99ResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L190>)

<a name="CacheMetrics.SetReadsPerSecond"></a>
### func \(\*CacheMetrics\) [SetReadsPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L172>)

<a name="CacheMetrics.SetTotalConnections"></a>
### func \(\*CacheMetrics\) [SetTotalConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L194>)

<a name="CacheMetrics.SetWritesPerSecond"></a>
### func \(\*CacheMetrics\) [SetWritesPerSecond](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L177>)

<a name="CacheMetrics.String"></a>
### func \(\*CacheMetrics\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L58>)


## type [CacheMetrics\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L380-L409>)

    // contains filtered or unexported fields
### func \(CacheMetrics\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_metrics.pb.go#L411>)
<a name="CacheOperationResult"></a>
## type [CacheOperationResult](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L30-L45>)

\* Result of cache operations. Provides detailed outcome information for cache operations.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheOperationResult\) [ClearDurationMicroseconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L262>)

<a name="CacheOperationResult.ClearError"></a>
### func \(\*CacheOperationResult\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L276>)

<a name="CacheOperationResult.ClearItemsAffected"></a>
### func \(\*CacheOperationResult\) [ClearItemsAffected](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L271>)

<a name="CacheOperationResult.ClearKey"></a>
### func \(\*CacheOperationResult\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L252>)

<a name="CacheOperationResult.ClearNamespace"></a>
### func \(\*CacheOperationResult\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L257>)

<a name="CacheOperationResult.ClearOperationType"></a>
### func \(\*CacheOperationResult\) [ClearOperationType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L247>)

<a name="CacheOperationResult.ClearSuccess"></a>
### func \(\*CacheOperationResult\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L242>)

<a name="CacheOperationResult.ClearTimestamp"></a>
### func \(\*CacheOperationResult\) [ClearTimestamp](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L267>)

<a name="CacheOperationResult.GetDurationMicroseconds"></a>
### func \(\*CacheOperationResult\) [GetDurationMicroseconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L109>)

<a name="CacheOperationResult.GetError"></a>
### func \(\*CacheOperationResult\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L130>)

<a name="CacheOperationResult.GetItemsAffected"></a>
### func \(\*CacheOperationResult\) [GetItemsAffected](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L123>)

<a name="CacheOperationResult.GetKey"></a>
### func \(\*CacheOperationResult\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L89>)

<a name="CacheOperationResult.GetMetadata"></a>
### func \(\*CacheOperationResult\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L137>)

<a name="CacheOperationResult.GetNamespace"></a>
### func \(\*CacheOperationResult\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L99>)

<a name="CacheOperationResult.GetOperationType"></a>
### func \(\*CacheOperationResult\) [GetOperationType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L79>)

<a name="CacheOperationResult.GetSuccess"></a>
### func \(\*CacheOperationResult\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L72>)

<a name="CacheOperationResult.GetTimestamp"></a>
### func \(\*CacheOperationResult\) [GetTimestamp](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L116>)

<a name="CacheOperationResult.HasDurationMicroseconds"></a>
### func \(\*CacheOperationResult\) [HasDurationMicroseconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L214>)

<a name="CacheOperationResult.HasError"></a>
### func \(\*CacheOperationResult\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L235>)

<a name="CacheOperationResult.HasItemsAffected"></a>
### func \(\*CacheOperationResult\) [HasItemsAffected](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L228>)

<a name="CacheOperationResult.HasKey"></a>
### func \(\*CacheOperationResult\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L200>)

<a name="CacheOperationResult.HasNamespace"></a>
### func \(\*CacheOperationResult\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L207>)

<a name="CacheOperationResult.HasOperationType"></a>
### func \(\*CacheOperationResult\) [HasOperationType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L193>)

<a name="CacheOperationResult.HasSuccess"></a>
### func \(\*CacheOperationResult\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L186>)

<a name="CacheOperationResult.HasTimestamp"></a>
### func \(\*CacheOperationResult\) [HasTimestamp](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L221>)

<a name="CacheOperationResult.ProtoMessage"></a>
### func \(\*CacheOperationResult\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L58>)

<a name="CacheOperationResult.ProtoReflect"></a>
### func \(\*CacheOperationResult\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L60>)

<a name="CacheOperationResult.Reset"></a>
### func \(\*CacheOperationResult\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L47>)

<a name="CacheOperationResult.SetDurationMicroseconds"></a>
### func \(\*CacheOperationResult\) [SetDurationMicroseconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L164>)

<a name="CacheOperationResult.SetError"></a>
### func \(\*CacheOperationResult\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L178>)

<a name="CacheOperationResult.SetItemsAffected"></a>
### func \(\*CacheOperationResult\) [SetItemsAffected](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L173>)

<a name="CacheOperationResult.SetKey"></a>
### func \(\*CacheOperationResult\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L154>)

<a name="CacheOperationResult.SetMetadata"></a>
### func \(\*CacheOperationResult\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L182>)

<a name="CacheOperationResult.SetNamespace"></a>
### func \(\*CacheOperationResult\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L159>)

<a name="CacheOperationResult.SetOperationType"></a>
### func \(\*CacheOperationResult\) [SetOperationType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L149>)

<a name="CacheOperationResult.SetSuccess"></a>
### func \(\*CacheOperationResult\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L144>)

<a name="CacheOperationResult.SetTimestamp"></a>
### func \(\*CacheOperationResult\) [SetTimestamp](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L169>)

<a name="CacheOperationResult.String"></a>
### func \(\*CacheOperationResult\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L54>)


## type [CacheOperationResult\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L280-L301>)

    // contains filtered or unexported fields
### func \(CacheOperationResult\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_operation_result.pb.go#L303>)

<a name="CachePublishRequest"></a>
## type [CachePublishRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L29-L40>)

    // contains filtered or unexported fields
### func \(\*CachePublishRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L159>)

<a name="CachePublishRequest.ClearPayload"></a>
### func \(\*CachePublishRequest\) [ClearPayload](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L154>)

<a name="CachePublishRequest.ClearTopic"></a>
### func \(\*CachePublishRequest\) [ClearTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L149>)

<a name="CachePublishRequest.GetMetadata"></a>
### func \(\*CachePublishRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L91>)

<a name="CachePublishRequest.GetPayload"></a>
### func \(\*CachePublishRequest\) [GetPayload](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L77>)

<a name="CachePublishRequest.GetTopic"></a>
### func \(\*CachePublishRequest\) [GetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L67>)

<a name="CachePublishRequest.HasMetadata"></a>
### func \(\*CachePublishRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L142>)

<a name="CachePublishRequest.HasPayload"></a>
### func \(\*CachePublishRequest\) [HasPayload](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L135>)

<a name="CachePublishRequest.HasTopic"></a>
### func \(\*CachePublishRequest\) [HasTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L128>)

<a name="CachePublishRequest.ProtoMessage"></a>
### func \(\*CachePublishRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L53>)

<a name="CachePublishRequest.ProtoReflect"></a>
### func \(\*CachePublishRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L55>)

<a name="CachePublishRequest.Reset"></a>
### func \(\*CachePublishRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L42>)

<a name="CachePublishRequest.SetMetadata"></a>
### func \(\*CachePublishRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L119>)

<a name="CachePublishRequest.SetPayload"></a>
### func \(\*CachePublishRequest\) [SetPayload](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L110>)

<a name="CachePublishRequest.SetTopic"></a>
### func \(\*CachePublishRequest\) [SetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L105>)

<a name="CachePublishRequest.String"></a>
### func \(\*CachePublishRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L49>)


## type [CachePublishRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L164-L173>)

    // contains filtered or unexported fields
### func \(CachePublishRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/publish_request.pb.go#L175>)

<a name="CacheServiceClient"></a>
## type [CacheServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L46-L75>)
For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
\* CacheService provides comprehensive caching capabilities. Supports CRUD operations, batch operations, atomic operations, and cache management with flexible expiration policies.
### func [NewCacheServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L81>)
<a name="CacheServiceServer"></a>
## type [CacheServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L233-L263>)

CacheServiceServer is the server API for CacheService service. All implementations must embed UnimplementedCacheServiceServer for forward compatibility.
\* CacheService provides comprehensive caching capabilities. Supports CRUD operations, batch operations, atomic operations, and cache management with flexible expiration policies.
    // contains filtered or unexported methods
## type [CacheStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L29-L46>)
\* Cache statistics and performance metrics. Provides detailed information about cache usage and performance.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CacheStats\) [ClearAvgAccessTimeMs](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L321>)

<a name="CacheStats.ClearCacheHits"></a>
### func \(\*CacheStats\) [ClearCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L286>)

<a name="CacheStats.ClearCacheMisses"></a>
### func \(\*CacheStats\) [ClearCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L291>)

<a name="CacheStats.ClearEvictedItems"></a>
### func \(\*CacheStats\) [ClearEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L311>)

<a name="CacheStats.ClearExpiredItems"></a>
### func \(\*CacheStats\) [ClearExpiredItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L316>)

<a name="CacheStats.ClearHitRatio"></a>
### func \(\*CacheStats\) [ClearHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L296>)

<a name="CacheStats.ClearLastReset"></a>
### func \(\*CacheStats\) [ClearLastReset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L326>)

<a name="CacheStats.ClearMemoryLimit"></a>
### func \(\*CacheStats\) [ClearMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L306>)

<a name="CacheStats.ClearMemoryUsage"></a>
### func \(\*CacheStats\) [ClearMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L301>)

<a name="CacheStats.ClearTotalItems"></a>
### func \(\*CacheStats\) [ClearTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L281>)

<a name="CacheStats.ClearUptimeSeconds"></a>
### func \(\*CacheStats\) [ClearUptimeSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L330>)

<a name="CacheStats.GetAvgAccessTimeMs"></a>
### func \(\*CacheStats\) [GetAvgAccessTimeMs](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L129>)

<a name="CacheStats.GetCacheHits"></a>
### func \(\*CacheStats\) [GetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L80>)

<a name="CacheStats.GetCacheMisses"></a>
### func \(\*CacheStats\) [GetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L87>)

<a name="CacheStats.GetEvictedItems"></a>
### func \(\*CacheStats\) [GetEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L115>)

<a name="CacheStats.GetExpiredItems"></a>
### func \(\*CacheStats\) [GetExpiredItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L122>)

<a name="CacheStats.GetHitRatio"></a>
### func \(\*CacheStats\) [GetHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L94>)

<a name="CacheStats.GetLastReset"></a>
### func \(\*CacheStats\) [GetLastReset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L136>)

<a name="CacheStats.GetMemoryLimit"></a>
### func \(\*CacheStats\) [GetMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L108>)

<a name="CacheStats.GetMemoryUsage"></a>
### func \(\*CacheStats\) [GetMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L101>)

<a name="CacheStats.GetTotalItems"></a>
### func \(\*CacheStats\) [GetTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L73>)

<a name="CacheStats.GetUptimeSeconds"></a>
### func \(\*CacheStats\) [GetUptimeSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L143>)

<a name="CacheStats.HasAvgAccessTimeMs"></a>
### func \(\*CacheStats\) [HasAvgAccessTimeMs](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L260>)

<a name="CacheStats.HasCacheHits"></a>
### func \(\*CacheStats\) [HasCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L211>)

<a name="CacheStats.HasCacheMisses"></a>
### func \(\*CacheStats\) [HasCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L218>)

<a name="CacheStats.HasEvictedItems"></a>
### func \(\*CacheStats\) [HasEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L246>)

<a name="CacheStats.HasExpiredItems"></a>
### func \(\*CacheStats\) [HasExpiredItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L253>)

<a name="CacheStats.HasHitRatio"></a>
### func \(\*CacheStats\) [HasHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L225>)

<a name="CacheStats.HasLastReset"></a>
### func \(\*CacheStats\) [HasLastReset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L267>)

<a name="CacheStats.HasMemoryLimit"></a>
### func \(\*CacheStats\) [HasMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L239>)

<a name="CacheStats.HasMemoryUsage"></a>
### func \(\*CacheStats\) [HasMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L232>)

<a name="CacheStats.HasTotalItems"></a>
### func \(\*CacheStats\) [HasTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L204>)

<a name="CacheStats.HasUptimeSeconds"></a>
### func \(\*CacheStats\) [HasUptimeSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L274>)

<a name="CacheStats.ProtoMessage"></a>
### func \(\*CacheStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L59>)

<a name="CacheStats.ProtoReflect"></a>
### func \(\*CacheStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L61>)

<a name="CacheStats.Reset"></a>
### func \(\*CacheStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L48>)

<a name="CacheStats.SetAvgAccessTimeMs"></a>
### func \(\*CacheStats\) [SetAvgAccessTimeMs](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L190>)

<a name="CacheStats.SetCacheHits"></a>
### func \(\*CacheStats\) [SetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L155>)

<a name="CacheStats.SetCacheMisses"></a>
### func \(\*CacheStats\) [SetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L160>)

<a name="CacheStats.SetEvictedItems"></a>
### func \(\*CacheStats\) [SetEvictedItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L180>)

<a name="CacheStats.SetExpiredItems"></a>
### func \(\*CacheStats\) [SetExpiredItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L185>)

<a name="CacheStats.SetHitRatio"></a>
### func \(\*CacheStats\) [SetHitRatio](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L165>)

<a name="CacheStats.SetLastReset"></a>
### func \(\*CacheStats\) [SetLastReset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L195>)

<a name="CacheStats.SetMemoryLimit"></a>
### func \(\*CacheStats\) [SetMemoryLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L175>)

<a name="CacheStats.SetMemoryUsage"></a>
### func \(\*CacheStats\) [SetMemoryUsage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L170>)

<a name="CacheStats.SetTotalItems"></a>
### func \(\*CacheStats\) [SetTotalItems](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L150>)

<a name="CacheStats.SetUptimeSeconds"></a>
### func \(\*CacheStats\) [SetUptimeSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L199>)

<a name="CacheStats.String"></a>
### func \(\*CacheStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L55>)


## type [CacheStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L335-L360>)

    // contains filtered or unexported fields
### func \(CacheStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_stats.pb.go#L362>)

<a name="CacheSubscribeRequest"></a>
## type [CacheSubscribeRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*CacheSubscribeRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L122>)

<a name="CacheSubscribeRequest.ClearTopic"></a>
### func \(\*CacheSubscribeRequest\) [ClearTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L117>)

<a name="CacheSubscribeRequest.GetMetadata"></a>
### func \(\*CacheSubscribeRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L75>)

<a name="CacheSubscribeRequest.GetTopic"></a>
### func \(\*CacheSubscribeRequest\) [GetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L65>)

<a name="CacheSubscribeRequest.HasMetadata"></a>
### func \(\*CacheSubscribeRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L110>)

<a name="CacheSubscribeRequest.HasTopic"></a>
### func \(\*CacheSubscribeRequest\) [HasTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L103>)

<a name="CacheSubscribeRequest.ProtoMessage"></a>
### func \(\*CacheSubscribeRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L51>)

<a name="CacheSubscribeRequest.ProtoReflect"></a>
### func \(\*CacheSubscribeRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L53>)

<a name="CacheSubscribeRequest.Reset"></a>
### func \(\*CacheSubscribeRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L40>)

<a name="CacheSubscribeRequest.SetMetadata"></a>
### func \(\*CacheSubscribeRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L94>)

<a name="CacheSubscribeRequest.SetTopic"></a>
### func \(\*CacheSubscribeRequest\) [SetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L89>)

<a name="CacheSubscribeRequest.String"></a>
### func \(\*CacheSubscribeRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L47>)


## type [CacheSubscribeRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(CacheSubscribeRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/subscribe_request.pb.go#L136>)

<a name="CacheUnsubscribeRequest"></a>
## type [CacheUnsubscribeRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*CacheUnsubscribeRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L122>)

<a name="CacheUnsubscribeRequest.ClearTopic"></a>
### func \(\*CacheUnsubscribeRequest\) [ClearTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L117>)

<a name="CacheUnsubscribeRequest.GetMetadata"></a>
### func \(\*CacheUnsubscribeRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L75>)

<a name="CacheUnsubscribeRequest.GetTopic"></a>
### func \(\*CacheUnsubscribeRequest\) [GetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L65>)

<a name="CacheUnsubscribeRequest.HasMetadata"></a>
### func \(\*CacheUnsubscribeRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L110>)

<a name="CacheUnsubscribeRequest.HasTopic"></a>
### func \(\*CacheUnsubscribeRequest\) [HasTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L103>)

<a name="CacheUnsubscribeRequest.ProtoMessage"></a>
### func \(\*CacheUnsubscribeRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L51>)

<a name="CacheUnsubscribeRequest.ProtoReflect"></a>
### func \(\*CacheUnsubscribeRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L53>)

<a name="CacheUnsubscribeRequest.Reset"></a>
### func \(\*CacheUnsubscribeRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L40>)

<a name="CacheUnsubscribeRequest.SetMetadata"></a>
### func \(\*CacheUnsubscribeRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L94>)

<a name="CacheUnsubscribeRequest.SetTopic"></a>
### func \(\*CacheUnsubscribeRequest\) [SetTopic](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L89>)

<a name="CacheUnsubscribeRequest.String"></a>
### func \(\*CacheUnsubscribeRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L47>)


## type [CacheUnsubscribeRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(CacheUnsubscribeRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unsubscribe_request.pb.go#L136>)

<a name="CacheWatchRequest"></a>
## type [CacheWatchRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*CacheWatchRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L117>)

<a name="CacheWatchRequest.ClearMetadata"></a>
### func \(\*CacheWatchRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L122>)

<a name="CacheWatchRequest.GetKey"></a>
### func \(\*CacheWatchRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L65>)

<a name="CacheWatchRequest.GetMetadata"></a>
### func \(\*CacheWatchRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L75>)

<a name="CacheWatchRequest.HasKey"></a>
### func \(\*CacheWatchRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L103>)

<a name="CacheWatchRequest.HasMetadata"></a>
### func \(\*CacheWatchRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L110>)

<a name="CacheWatchRequest.ProtoMessage"></a>
### func \(\*CacheWatchRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L51>)

<a name="CacheWatchRequest.ProtoReflect"></a>
### func \(\*CacheWatchRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L53>)

<a name="CacheWatchRequest.Reset"></a>
### func \(\*CacheWatchRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L40>)

<a name="CacheWatchRequest.SetKey"></a>
### func \(\*CacheWatchRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L89>)

<a name="CacheWatchRequest.SetMetadata"></a>
### func \(\*CacheWatchRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L94>)

<a name="CacheWatchRequest.String"></a>
### func \(\*CacheWatchRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L47>)


## type [CacheWatchRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(CacheWatchRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/watch_request.pb.go#L136>)
<a name="ClearRequest"></a>
## type [ClearRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L29-L37>)

\* Request to clear all cache entries. Optionally clear only a specific namespace.
    // contains filtered or unexported fields
### func \(\*ClearRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L109>)

<a name="ClearRequest.ClearNamespace"></a>
### func \(\*ClearRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L104>)

<a name="ClearRequest.GetMetadata"></a>
### func \(\*ClearRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L74>)

<a name="ClearRequest.GetNamespace"></a>
### func \(\*ClearRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L64>)

<a name="ClearRequest.HasMetadata"></a>
### func \(\*ClearRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L97>)

<a name="ClearRequest.HasNamespace"></a>
### func \(\*ClearRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L90>)

<a name="ClearRequest.ProtoMessage"></a>
### func \(\*ClearRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L50>)

<a name="ClearRequest.ProtoReflect"></a>
### func \(\*ClearRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L52>)

<a name="ClearRequest.Reset"></a>
### func \(\*ClearRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L39>)

<a name="ClearRequest.SetMetadata"></a>
### func \(\*ClearRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L86>)

<a name="ClearRequest.SetNamespace"></a>
### func \(\*ClearRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L81>)

<a name="ClearRequest.String"></a>
### func \(\*ClearRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L46>)


## type [ClearRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L113-L120>)

    // contains filtered or unexported fields
### func \(ClearRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_request.pb.go#L122>)
<a name="ClearResponse"></a>
## type [ClearResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L29-L38>)

\* Response for cache clear operations. Indicates success/failure of clearing cache entries.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*ClearResponse\) [ClearClearedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L121>)

<a name="ClearResponse.ClearError"></a>
### func \(\*ClearResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L131>)

<a name="ClearResponse.ClearSuccess"></a>
### func \(\*ClearResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L126>)

<a name="ClearResponse.GetClearedCount"></a>
### func \(\*ClearResponse\) [GetClearedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L65>)

<a name="ClearResponse.GetError"></a>
### func \(\*ClearResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L79>)

<a name="ClearResponse.GetSuccess"></a>
### func \(\*ClearResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L72>)

<a name="ClearResponse.HasClearedCount"></a>
### func \(\*ClearResponse\) [HasClearedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L100>)

<a name="ClearResponse.HasError"></a>
### func \(\*ClearResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L114>)

<a name="ClearResponse.HasSuccess"></a>
### func \(\*ClearResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L107>)

<a name="ClearResponse.ProtoMessage"></a>
### func \(\*ClearResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L51>)

<a name="ClearResponse.ProtoReflect"></a>
### func \(\*ClearResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L53>)

<a name="ClearResponse.Reset"></a>
### func \(\*ClearResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L40>)

<a name="ClearResponse.SetClearedCount"></a>
### func \(\*ClearResponse\) [SetClearedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L86>)

<a name="ClearResponse.SetError"></a>
### func \(\*ClearResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L96>)

<a name="ClearResponse.SetSuccess"></a>
### func \(\*ClearResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L91>)

<a name="ClearResponse.String"></a>
### func \(\*ClearResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L47>)


## type [ClearResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L135-L144>)

    // contains filtered or unexported fields
### func \(ClearResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/clear_response.pb.go#L146>)
<a name="CockroachConfig"></a>
## type [CockroachConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L28-L43>)

\* CockroachConfig provides CockroachDB\-specific connection configuration. Includes retry behavior and identification options for robust connections.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CockroachConfig\) [ClearApplicationName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L289>)

<a name="CockroachConfig.ClearDatabase"></a>
### func \(\*CockroachConfig\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L279>)

<a name="CockroachConfig.ClearHost"></a>
### func \(\*CockroachConfig\) [ClearHost](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L259>)

<a name="CockroachConfig.ClearMaxRetries"></a>
### func \(\*CockroachConfig\) [ClearMaxRetries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L299>)

<a name="CockroachConfig.ClearPassword"></a>
### func \(\*CockroachConfig\) [ClearPassword](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L274>)

<a name="CockroachConfig.ClearPort"></a>
### func \(\*CockroachConfig\) [ClearPort](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L264>)

<a name="CockroachConfig.ClearRetryBackoffFactor"></a>
### func \(\*CockroachConfig\) [ClearRetryBackoffFactor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L294>)

<a name="CockroachConfig.ClearSslMode"></a>
### func \(\*CockroachConfig\) [ClearSslMode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L284>)

<a name="CockroachConfig.ClearUser"></a>
### func \(\*CockroachConfig\) [ClearUser](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L269>)

<a name="CockroachConfig.GetApplicationName"></a>
### func \(\*CockroachConfig\) [GetApplicationName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L127>)

<a name="CockroachConfig.GetDatabase"></a>
### func \(\*CockroachConfig\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L107>)

<a name="CockroachConfig.GetHost"></a>
### func \(\*CockroachConfig\) [GetHost](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L70>)

<a name="CockroachConfig.GetMaxRetries"></a>
### func \(\*CockroachConfig\) [GetMaxRetries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L144>)

<a name="CockroachConfig.GetPassword"></a>
### func \(\*CockroachConfig\) [GetPassword](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L97>)

<a name="CockroachConfig.GetPort"></a>
### func \(\*CockroachConfig\) [GetPort](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L80>)

<a name="CockroachConfig.GetRetryBackoffFactor"></a>
### func \(\*CockroachConfig\) [GetRetryBackoffFactor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L137>)

<a name="CockroachConfig.GetSslMode"></a>
### func \(\*CockroachConfig\) [GetSslMode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L117>)

<a name="CockroachConfig.GetUser"></a>
### func \(\*CockroachConfig\) [GetUser](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L87>)

<a name="CockroachConfig.HasApplicationName"></a>
### func \(\*CockroachConfig\) [HasApplicationName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L238>)

<a name="CockroachConfig.HasDatabase"></a>
### func \(\*CockroachConfig\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L224>)

<a name="CockroachConfig.HasHost"></a>
### func \(\*CockroachConfig\) [HasHost](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L196>)

<a name="CockroachConfig.HasMaxRetries"></a>
### func \(\*CockroachConfig\) [HasMaxRetries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L252>)

<a name="CockroachConfig.HasPassword"></a>
### func \(\*CockroachConfig\) [HasPassword](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L217>)

<a name="CockroachConfig.HasPort"></a>
### func \(\*CockroachConfig\) [HasPort](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L203>)


### func \(\*CockroachConfig\) [HasRetryBackoffFactor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L245>)

<a name="CockroachConfig.HasSslMode"></a>
### func \(\*CockroachConfig\) [HasSslMode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L231>)

<a name="CockroachConfig.HasUser"></a>
### func \(\*CockroachConfig\) [HasUser](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L210>)

<a name="CockroachConfig.ProtoMessage"></a>
### func \(\*CockroachConfig\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L56>)

<a name="CockroachConfig.ProtoReflect"></a>
### func \(\*CockroachConfig\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L58>)

<a name="CockroachConfig.Reset"></a>
### func \(\*CockroachConfig\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L45>)

<a name="CockroachConfig.SetApplicationName"></a>
### func \(\*CockroachConfig\) [SetApplicationName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L181>)

<a name="CockroachConfig.SetDatabase"></a>
### func \(\*CockroachConfig\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L171>)

<a name="CockroachConfig.SetHost"></a>
### func \(\*CockroachConfig\) [SetHost](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L151>)

<a name="CockroachConfig.SetMaxRetries"></a>
### func \(\*CockroachConfig\) [SetMaxRetries](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L191>)

<a name="CockroachConfig.SetPassword"></a>
### func \(\*CockroachConfig\) [SetPassword](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L166>)

<a name="CockroachConfig.SetPort"></a>
### func \(\*CockroachConfig\) [SetPort](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L156>)

<a name="CockroachConfig.SetRetryBackoffFactor"></a>
### func \(\*CockroachConfig\) [SetRetryBackoffFactor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L186>)

<a name="CockroachConfig.SetSslMode"></a>
### func \(\*CockroachConfig\) [SetSslMode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L176>)

<a name="CockroachConfig.SetUser"></a>
### func \(\*CockroachConfig\) [SetUser](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L161>)

<a name="CockroachConfig.String"></a>
### func \(\*CockroachConfig\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L52>)


## type [CockroachConfig\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L304-L325>)

    // contains filtered or unexported fields
### func \(CockroachConfig\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cockroach_config.pb.go#L327>)
<a name="ColumnMetadata"></a>
## type [ColumnMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L28-L42>)

\* ColumnMetadata describes the structure and properties of a database column. Used in result sets to provide type information for proper data handling.

    // contains filtered or unexported fields
### func \(\*ColumnMetadata\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L181>)

<a name="ColumnMetadata.ClearNullable"></a>
### func \(\*ColumnMetadata\) [ClearNullable](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L191>)

<a name="ColumnMetadata.ClearScale"></a>
### func \(\*ColumnMetadata\) [ClearScale](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L201>)

<a name="ColumnMetadata.ClearSize"></a>
### func \(\*ColumnMetadata\) [ClearSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L196>)

<a name="ColumnMetadata.ClearType"></a>
### func \(\*ColumnMetadata\) [ClearType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L186>)

<a name="ColumnMetadata.GetMetadata"></a>
### func \(\*ColumnMetadata\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L110>)

<a name="ColumnMetadata.GetName"></a>
### func \(\*ColumnMetadata\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L69>)

<a name="ColumnMetadata.GetNullable"></a>
### func \(\*ColumnMetadata\) [GetNullable](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L89>)

<a name="ColumnMetadata.GetScale"></a>
### func \(\*ColumnMetadata\) [GetScale](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L103>)

<a name="ColumnMetadata.GetSize"></a>
### func \(\*ColumnMetadata\) [GetSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L96>)

<a name="ColumnMetadata.GetType"></a>
### func \(\*ColumnMetadata\) [GetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L79>)

<a name="ColumnMetadata.HasName"></a>
### func \(\*ColumnMetadata\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L146>)

<a name="ColumnMetadata.HasNullable"></a>
### func \(\*ColumnMetadata\) [HasNullable](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L160>)

<a name="ColumnMetadata.HasScale"></a>
### func \(\*ColumnMetadata\) [HasScale](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L174>)

<a name="ColumnMetadata.HasSize"></a>
### func \(\*ColumnMetadata\) [HasSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L167>)

<a name="ColumnMetadata.HasType"></a>
### func \(\*ColumnMetadata\) [HasType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L153>)

<a name="ColumnMetadata.ProtoMessage"></a>
### func \(\*ColumnMetadata\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L55>)

<a name="ColumnMetadata.ProtoReflect"></a>
### func \(\*ColumnMetadata\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L57>)

<a name="ColumnMetadata.Reset"></a>
### func \(\*ColumnMetadata\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L44>)

<a name="ColumnMetadata.SetMetadata"></a>
### func \(\*ColumnMetadata\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L142>)

<a name="ColumnMetadata.SetName"></a>
### func \(\*ColumnMetadata\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L117>)

<a name="ColumnMetadata.SetNullable"></a>
### func \(\*ColumnMetadata\) [SetNullable](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L127>)

<a name="ColumnMetadata.SetScale"></a>
### func \(\*ColumnMetadata\) [SetScale](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L137>)

<a name="ColumnMetadata.SetSize"></a>
### func \(\*ColumnMetadata\) [SetSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L132>)

<a name="ColumnMetadata.SetType"></a>
### func \(\*ColumnMetadata\) [SetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L122>)

<a name="ColumnMetadata.String"></a>
### func \(\*ColumnMetadata\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L51>)


## type [ColumnMetadata\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L206-L221>)

    // contains filtered or unexported fields
### func \(ColumnMetadata\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/column_metadata.pb.go#L223>)


## type [CommitTransactionRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L26-L36>)


    // contains filtered or unexported fields
### func \(\*CommitTransactionRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L120>)

<a name="CommitTransactionRequest.ClearTransactionId"></a>
### func \(\*CommitTransactionRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L115>)

<a name="CommitTransactionRequest.GetMetadata"></a>
### func \(\*CommitTransactionRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L73>)

<a name="CommitTransactionRequest.GetTransactionId"></a>
### func \(\*CommitTransactionRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L63>)

<a name="CommitTransactionRequest.HasMetadata"></a>
### func \(\*CommitTransactionRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L108>)

<a name="CommitTransactionRequest.HasTransactionId"></a>
### func \(\*CommitTransactionRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L101>)

<a name="CommitTransactionRequest.ProtoMessage"></a>
### func \(\*CommitTransactionRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L49>)

<a name="CommitTransactionRequest.ProtoReflect"></a>
### func \(\*CommitTransactionRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L51>)

<a name="CommitTransactionRequest.Reset"></a>
### func \(\*CommitTransactionRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L38>)

<a name="CommitTransactionRequest.SetMetadata"></a>
### func \(\*CommitTransactionRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L92>)

<a name="CommitTransactionRequest.SetTransactionId"></a>
### func \(\*CommitTransactionRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L87>)

<a name="CommitTransactionRequest.String"></a>
### func \(\*CommitTransactionRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L45>)


## type [CommitTransactionRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L125-L132>)

    // contains filtered or unexported fields
### func \(CommitTransactionRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/commit_transaction_request.pb.go#L134>)

<a name="ConfigurePolicyRequest"></a>
## type [ConfigurePolicyRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L27-L38>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*ConfigurePolicyRequest\) [ClearEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L163>)

<a name="ConfigurePolicyRequest.ClearMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyRequest\) [ClearMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L168>)

<a name="ConfigurePolicyRequest.ClearMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyRequest\) [ClearMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L173>)

<a name="ConfigurePolicyRequest.ClearNamespaceId"></a>
### func \(\*ConfigurePolicyRequest\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L158>)

<a name="ConfigurePolicyRequest.GetEvictionPolicy"></a>
### func \(\*ConfigurePolicyRequest\) [GetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L75>)

<a name="ConfigurePolicyRequest.GetMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyRequest\) [GetMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L85>)

<a name="ConfigurePolicyRequest.GetMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyRequest\) [GetMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L92>)

<a name="ConfigurePolicyRequest.GetNamespaceId"></a>
### func \(\*ConfigurePolicyRequest\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L65>)

<a name="ConfigurePolicyRequest.GetPolicyConfig"></a>
### func \(\*ConfigurePolicyRequest\) [GetPolicyConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L99>)

<a name="ConfigurePolicyRequest.HasEvictionPolicy"></a>
### func \(\*ConfigurePolicyRequest\) [HasEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L137>)

<a name="ConfigurePolicyRequest.HasMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyRequest\) [HasMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L144>)

<a name="ConfigurePolicyRequest.HasMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyRequest\) [HasMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L151>)

<a name="ConfigurePolicyRequest.HasNamespaceId"></a>
### func \(\*ConfigurePolicyRequest\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L130>)

<a name="ConfigurePolicyRequest.ProtoMessage"></a>
### func \(\*ConfigurePolicyRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L51>)

<a name="ConfigurePolicyRequest.ProtoReflect"></a>
### func \(\*ConfigurePolicyRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L53>)

<a name="ConfigurePolicyRequest.Reset"></a>
### func \(\*ConfigurePolicyRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L40>)

<a name="ConfigurePolicyRequest.SetEvictionPolicy"></a>
### func \(\*ConfigurePolicyRequest\) [SetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L111>)

<a name="ConfigurePolicyRequest.SetMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyRequest\) [SetMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L116>)

<a name="ConfigurePolicyRequest.SetMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyRequest\) [SetMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L121>)

<a name="ConfigurePolicyRequest.SetNamespaceId"></a>
### func \(\*ConfigurePolicyRequest\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L106>)

<a name="ConfigurePolicyRequest.SetPolicyConfig"></a>
### func \(\*ConfigurePolicyRequest\) [SetPolicyConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L126>)

<a name="ConfigurePolicyRequest.String"></a>
### func \(\*ConfigurePolicyRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L47>)


## type [ConfigurePolicyRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L178-L191>)

    // contains filtered or unexported fields
### func \(ConfigurePolicyRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_request.pb.go#L193>)

<a name="ConfigurePolicyResponse"></a>
## type [ConfigurePolicyResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L28-L41>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*ConfigurePolicyResponse\) [ClearAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L210>)

<a name="ConfigurePolicyResponse.ClearEvictionPolicy"></a>
### func \(\*ConfigurePolicyResponse\) [ClearEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L195>)

<a name="ConfigurePolicyResponse.ClearMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyResponse\) [ClearMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L200>)

<a name="ConfigurePolicyResponse.ClearMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyResponse\) [ClearMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L205>)

<a name="ConfigurePolicyResponse.ClearNamespaceId"></a>
### func \(\*ConfigurePolicyResponse\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L190>)

<a name="ConfigurePolicyResponse.GetAppliedAt"></a>
### func \(\*ConfigurePolicyResponse\) [GetAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L102>)

<a name="ConfigurePolicyResponse.GetEvictionPolicy"></a>
### func \(\*ConfigurePolicyResponse\) [GetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L78>)

<a name="ConfigurePolicyResponse.GetMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyResponse\) [GetMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L88>)

<a name="ConfigurePolicyResponse.GetMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyResponse\) [GetMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L95>)

<a name="ConfigurePolicyResponse.GetNamespaceId"></a>
### func \(\*ConfigurePolicyResponse\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L68>)

<a name="ConfigurePolicyResponse.GetNewConfig"></a>
### func \(\*ConfigurePolicyResponse\) [GetNewConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L116>)

<a name="ConfigurePolicyResponse.GetPreviousConfig"></a>
### func \(\*ConfigurePolicyResponse\) [GetPreviousConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L109>)

<a name="ConfigurePolicyResponse.HasAppliedAt"></a>
### func \(\*ConfigurePolicyResponse\) [HasAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L183>)

<a name="ConfigurePolicyResponse.HasEvictionPolicy"></a>
### func \(\*ConfigurePolicyResponse\) [HasEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L162>)

<a name="ConfigurePolicyResponse.HasMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyResponse\) [HasMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L169>)

<a name="ConfigurePolicyResponse.HasMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyResponse\) [HasMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L176>)

<a name="ConfigurePolicyResponse.HasNamespaceId"></a>
### func \(\*ConfigurePolicyResponse\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L155>)

<a name="ConfigurePolicyResponse.ProtoMessage"></a>
### func \(\*ConfigurePolicyResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L54>)

<a name="ConfigurePolicyResponse.ProtoReflect"></a>
### func \(\*ConfigurePolicyResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L56>)

<a name="ConfigurePolicyResponse.Reset"></a>
### func \(\*ConfigurePolicyResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L43>)

<a name="ConfigurePolicyResponse.SetAppliedAt"></a>
### func \(\*ConfigurePolicyResponse\) [SetAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L143>)

<a name="ConfigurePolicyResponse.SetEvictionPolicy"></a>
### func \(\*ConfigurePolicyResponse\) [SetEvictionPolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L128>)

<a name="ConfigurePolicyResponse.SetMaxTtlSeconds"></a>
### func \(\*ConfigurePolicyResponse\) [SetMaxTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L133>)

<a name="ConfigurePolicyResponse.SetMemoryThresholdPercent"></a>
### func \(\*ConfigurePolicyResponse\) [SetMemoryThresholdPercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L138>)

<a name="ConfigurePolicyResponse.SetNamespaceId"></a>
### func \(\*ConfigurePolicyResponse\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L123>)

<a name="ConfigurePolicyResponse.SetNewConfig"></a>
### func \(\*ConfigurePolicyResponse\) [SetNewConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L151>)

<a name="ConfigurePolicyResponse.SetPreviousConfig"></a>
### func \(\*ConfigurePolicyResponse\) [SetPreviousConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L147>)

<a name="ConfigurePolicyResponse.String"></a>
### func \(\*ConfigurePolicyResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L50>)


## type [ConfigurePolicyResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L214-L231>)

    // contains filtered or unexported fields
### func \(ConfigurePolicyResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/configure_policy_response.pb.go#L233>)
<a name="ConnectionPoolInfo"></a>
## type [ConnectionPoolInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L29-L42>)

\* ConnectionPoolInfo provides information about database connection pool status. Used for monitoring connection health and pool performance.

    // contains filtered or unexported fields
### func \(\*ConnectionPoolInfo\) [ClearActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L191>)

<a name="ConnectionPoolInfo.ClearAvgLifetime"></a>
### func \(\*ConnectionPoolInfo\) [ClearAvgLifetime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L201>)

<a name="ConnectionPoolInfo.ClearIdleConnections"></a>
### func \(\*ConnectionPoolInfo\) [ClearIdleConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L196>)

<a name="ConnectionPoolInfo.ClearMaxConnections"></a>
### func \(\*ConnectionPoolInfo\) [ClearMaxConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L186>)

<a name="ConnectionPoolInfo.ClearStats"></a>
### func \(\*ConnectionPoolInfo\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L206>)

<a name="ConnectionPoolInfo.GetActiveConnections"></a>
### func \(\*ConnectionPoolInfo\) [GetActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L76>)

<a name="ConnectionPoolInfo.GetAvgLifetime"></a>
### func \(\*ConnectionPoolInfo\) [GetAvgLifetime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L90>)

<a name="ConnectionPoolInfo.GetIdleConnections"></a>
### func \(\*ConnectionPoolInfo\) [GetIdleConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L83>)

<a name="ConnectionPoolInfo.GetMaxConnections"></a>
### func \(\*ConnectionPoolInfo\) [GetMaxConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L69>)

<a name="ConnectionPoolInfo.GetStats"></a>
### func \(\*ConnectionPoolInfo\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L104>)

<a name="ConnectionPoolInfo.HasActiveConnections"></a>
### func \(\*ConnectionPoolInfo\) [HasActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L158>)

<a name="ConnectionPoolInfo.HasAvgLifetime"></a>
### func \(\*ConnectionPoolInfo\) [HasAvgLifetime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L172>)

<a name="ConnectionPoolInfo.HasIdleConnections"></a>
### func \(\*ConnectionPoolInfo\) [HasIdleConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L165>)

<a name="ConnectionPoolInfo.HasMaxConnections"></a>
### func \(\*ConnectionPoolInfo\) [HasMaxConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L151>)

<a name="ConnectionPoolInfo.HasStats"></a>
### func \(\*ConnectionPoolInfo\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L179>)

<a name="ConnectionPoolInfo.ProtoMessage"></a>
### func \(\*ConnectionPoolInfo\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L55>)

<a name="ConnectionPoolInfo.ProtoReflect"></a>
### func \(\*ConnectionPoolInfo\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L57>)

<a name="ConnectionPoolInfo.Reset"></a>
### func \(\*ConnectionPoolInfo\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L44>)

<a name="ConnectionPoolInfo.SetActiveConnections"></a>
### func \(\*ConnectionPoolInfo\) [SetActiveConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L123>)

<a name="ConnectionPoolInfo.SetAvgLifetime"></a>
### func \(\*ConnectionPoolInfo\) [SetAvgLifetime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L133>)

<a name="ConnectionPoolInfo.SetIdleConnections"></a>
### func \(\*ConnectionPoolInfo\) [SetIdleConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L128>)

<a name="ConnectionPoolInfo.SetMaxConnections"></a>
### func \(\*ConnectionPoolInfo\) [SetMaxConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L118>)

<a name="ConnectionPoolInfo.SetStats"></a>
### func \(\*ConnectionPoolInfo\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L142>)

<a name="ConnectionPoolInfo.String"></a>
### func \(\*ConnectionPoolInfo\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L51>)


## type [ConnectionPoolInfo\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L211-L224>)

    // contains filtered or unexported fields
### func \(ConnectionPoolInfo\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/connection_pool_info.pb.go#L226>)


## type [CreateDatabaseRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L26-L37>)


    // contains filtered or unexported fields
### func \(\*CreateDatabaseRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L132>)

<a name="CreateDatabaseRequest.ClearName"></a>
### func \(\*CreateDatabaseRequest\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L127>)

<a name="CreateDatabaseRequest.GetMetadata"></a>
### func \(\*CreateDatabaseRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L81>)

<a name="CreateDatabaseRequest.GetName"></a>
### func \(\*CreateDatabaseRequest\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L64>)

<a name="CreateDatabaseRequest.GetOptions"></a>
### func \(\*CreateDatabaseRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L74>)

<a name="CreateDatabaseRequest.HasMetadata"></a>
### func \(\*CreateDatabaseRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L120>)

<a name="CreateDatabaseRequest.HasName"></a>
### func \(\*CreateDatabaseRequest\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L113>)

<a name="CreateDatabaseRequest.ProtoMessage"></a>
### func \(\*CreateDatabaseRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L50>)

<a name="CreateDatabaseRequest.ProtoReflect"></a>
### func \(\*CreateDatabaseRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L52>)

<a name="CreateDatabaseRequest.Reset"></a>
### func \(\*CreateDatabaseRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L39>)

<a name="CreateDatabaseRequest.SetMetadata"></a>
### func \(\*CreateDatabaseRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L104>)

<a name="CreateDatabaseRequest.SetName"></a>
### func \(\*CreateDatabaseRequest\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L95>)

<a name="CreateDatabaseRequest.SetOptions"></a>
### func \(\*CreateDatabaseRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L100>)

<a name="CreateDatabaseRequest.String"></a>
### func \(\*CreateDatabaseRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L46>)


## type [CreateDatabaseRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L137-L146>)

    // contains filtered or unexported fields
### func \(CreateDatabaseRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_request.pb.go#L148>)
<a name="CreateDatabaseResponse"></a>
## type [CreateDatabaseResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L28-L38>)

\* CreateDatabaseResponse contains the result of a database creation operation. Indicates success status and provides error details if creation failed.

    // contains filtered or unexported fields
### func \(\*CreateDatabaseResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L119>)

<a name="CreateDatabaseResponse.ClearSuccess"></a>
### func \(\*CreateDatabaseResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L114>)

<a name="CreateDatabaseResponse.GetError"></a>
### func \(\*CreateDatabaseResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L72>)

<a name="CreateDatabaseResponse.GetSuccess"></a>
### func \(\*CreateDatabaseResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L65>)

<a name="CreateDatabaseResponse.HasError"></a>
### func \(\*CreateDatabaseResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L107>)

<a name="CreateDatabaseResponse.HasSuccess"></a>
### func \(\*CreateDatabaseResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L100>)

<a name="CreateDatabaseResponse.ProtoMessage"></a>
### func \(\*CreateDatabaseResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L51>)

<a name="CreateDatabaseResponse.ProtoReflect"></a>
### func \(\*CreateDatabaseResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L53>)

<a name="CreateDatabaseResponse.Reset"></a>
### func \(\*CreateDatabaseResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L40>)

<a name="CreateDatabaseResponse.SetError"></a>
### func \(\*CreateDatabaseResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L91>)

<a name="CreateDatabaseResponse.SetSuccess"></a>
### func \(\*CreateDatabaseResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L86>)

<a name="CreateDatabaseResponse.String"></a>
### func \(\*CreateDatabaseResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L47>)


## type [CreateDatabaseResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L124-L131>)

    // contains filtered or unexported fields
### func \(CreateDatabaseResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_database_response.pb.go#L133>)

<a name="CreateNamespaceRequest"></a>
## type [CreateNamespaceRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L27-L39>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CreateNamespaceRequest\) [ClearDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L198>)

<a name="CreateNamespaceRequest.ClearDescription"></a>
### func \(\*CreateNamespaceRequest\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L183>)

<a name="CreateNamespaceRequest.ClearMaxKeys"></a>
### func \(\*CreateNamespaceRequest\) [ClearMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L188>)

<a name="CreateNamespaceRequest.ClearMaxMemoryBytes"></a>
### func \(\*CreateNamespaceRequest\) [ClearMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L193>)

<a name="CreateNamespaceRequest.ClearName"></a>
### func \(\*CreateNamespaceRequest\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L178>)

<a name="CreateNamespaceRequest.GetConfig"></a>
### func \(\*CreateNamespaceRequest\) [GetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L107>)

<a name="CreateNamespaceRequest.GetDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceRequest\) [GetDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L100>)

<a name="CreateNamespaceRequest.GetDescription"></a>
### func \(\*CreateNamespaceRequest\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L76>)

<a name="CreateNamespaceRequest.GetMaxKeys"></a>
### func \(\*CreateNamespaceRequest\) [GetMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L86>)

<a name="CreateNamespaceRequest.GetMaxMemoryBytes"></a>
### func \(\*CreateNamespaceRequest\) [GetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L93>)

<a name="CreateNamespaceRequest.GetName"></a>
### func \(\*CreateNamespaceRequest\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L66>)

<a name="CreateNamespaceRequest.HasDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceRequest\) [HasDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L171>)

<a name="CreateNamespaceRequest.HasDescription"></a>
### func \(\*CreateNamespaceRequest\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L150>)

<a name="CreateNamespaceRequest.HasMaxKeys"></a>
### func \(\*CreateNamespaceRequest\) [HasMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L157>)

<a name="CreateNamespaceRequest.HasMaxMemoryBytes"></a>
### func \(\*CreateNamespaceRequest\) [HasMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L164>)

<a name="CreateNamespaceRequest.HasName"></a>
### func \(\*CreateNamespaceRequest\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L143>)

<a name="CreateNamespaceRequest.ProtoMessage"></a>
### func \(\*CreateNamespaceRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L52>)

<a name="CreateNamespaceRequest.ProtoReflect"></a>
### func \(\*CreateNamespaceRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L54>)

<a name="CreateNamespaceRequest.Reset"></a>
### func \(\*CreateNamespaceRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L41>)

<a name="CreateNamespaceRequest.SetConfig"></a>
### func \(\*CreateNamespaceRequest\) [SetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L139>)

<a name="CreateNamespaceRequest.SetDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceRequest\) [SetDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L134>)

<a name="CreateNamespaceRequest.SetDescription"></a>
### func \(\*CreateNamespaceRequest\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L119>)

<a name="CreateNamespaceRequest.SetMaxKeys"></a>
### func \(\*CreateNamespaceRequest\) [SetMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L124>)

<a name="CreateNamespaceRequest.SetMaxMemoryBytes"></a>
### func \(\*CreateNamespaceRequest\) [SetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L129>)

<a name="CreateNamespaceRequest.SetName"></a>
### func \(\*CreateNamespaceRequest\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L114>)

<a name="CreateNamespaceRequest.String"></a>
### func \(\*CreateNamespaceRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L48>)


## type [CreateNamespaceRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L203-L218>)

    // contains filtered or unexported fields
### func \(CreateNamespaceRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_request.pb.go#L220>)

<a name="CreateNamespaceResponse"></a>
## type [CreateNamespaceResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L28-L42>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*CreateNamespaceResponse\) [ClearCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L236>)

<a name="CreateNamespaceResponse.ClearDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceResponse\) [ClearDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L250>)

<a name="CreateNamespaceResponse.ClearDescription"></a>
### func \(\*CreateNamespaceResponse\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L231>)

<a name="CreateNamespaceResponse.ClearMaxKeys"></a>
### func \(\*CreateNamespaceResponse\) [ClearMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L240>)

<a name="CreateNamespaceResponse.ClearMaxMemoryBytes"></a>
### func \(\*CreateNamespaceResponse\) [ClearMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L245>)

<a name="CreateNamespaceResponse.ClearName"></a>
### func \(\*CreateNamespaceResponse\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L226>)

<a name="CreateNamespaceResponse.ClearNamespaceId"></a>
### func \(\*CreateNamespaceResponse\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L221>)

<a name="CreateNamespaceResponse.GetConfig"></a>
### func \(\*CreateNamespaceResponse\) [GetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L127>)

<a name="CreateNamespaceResponse.GetCreatedAt"></a>
### func \(\*CreateNamespaceResponse\) [GetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L99>)

<a name="CreateNamespaceResponse.GetDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceResponse\) [GetDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L120>)

<a name="CreateNamespaceResponse.GetDescription"></a>
### func \(\*CreateNamespaceResponse\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L89>)

<a name="CreateNamespaceResponse.GetMaxKeys"></a>
### func \(\*CreateNamespaceResponse\) [GetMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L106>)

<a name="CreateNamespaceResponse.GetMaxMemoryBytes"></a>
### func \(\*CreateNamespaceResponse\) [GetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L113>)

<a name="CreateNamespaceResponse.GetName"></a>
### func \(\*CreateNamespaceResponse\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L79>)

<a name="CreateNamespaceResponse.GetNamespaceId"></a>
### func \(\*CreateNamespaceResponse\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L69>)

<a name="CreateNamespaceResponse.HasCreatedAt"></a>
### func \(\*CreateNamespaceResponse\) [HasCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L193>)

<a name="CreateNamespaceResponse.HasDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceResponse\) [HasDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L214>)

<a name="CreateNamespaceResponse.HasDescription"></a>
### func \(\*CreateNamespaceResponse\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L186>)

<a name="CreateNamespaceResponse.HasMaxKeys"></a>
### func \(\*CreateNamespaceResponse\) [HasMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L200>)

<a name="CreateNamespaceResponse.HasMaxMemoryBytes"></a>
### func \(\*CreateNamespaceResponse\) [HasMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L207>)

<a name="CreateNamespaceResponse.HasName"></a>
### func \(\*CreateNamespaceResponse\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L179>)

<a name="CreateNamespaceResponse.HasNamespaceId"></a>
### func \(\*CreateNamespaceResponse\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L172>)

<a name="CreateNamespaceResponse.ProtoMessage"></a>
### func \(\*CreateNamespaceResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L55>)

<a name="CreateNamespaceResponse.ProtoReflect"></a>
### func \(\*CreateNamespaceResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L57>)

<a name="CreateNamespaceResponse.Reset"></a>
### func \(\*CreateNamespaceResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L44>)

<a name="CreateNamespaceResponse.SetConfig"></a>
### func \(\*CreateNamespaceResponse\) [SetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L168>)

<a name="CreateNamespaceResponse.SetCreatedAt"></a>
### func \(\*CreateNamespaceResponse\) [SetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L149>)

<a name="CreateNamespaceResponse.SetDefaultTtlSeconds"></a>
### func \(\*CreateNamespaceResponse\) [SetDefaultTtlSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L163>)

<a name="CreateNamespaceResponse.SetDescription"></a>
### func \(\*CreateNamespaceResponse\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L144>)

<a name="CreateNamespaceResponse.SetMaxKeys"></a>
### func \(\*CreateNamespaceResponse\) [SetMaxKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L153>)

<a name="CreateNamespaceResponse.SetMaxMemoryBytes"></a>
### func \(\*CreateNamespaceResponse\) [SetMaxMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L158>)

<a name="CreateNamespaceResponse.SetName"></a>
### func \(\*CreateNamespaceResponse\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L139>)

<a name="CreateNamespaceResponse.SetNamespaceId"></a>
### func \(\*CreateNamespaceResponse\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L134>)

<a name="CreateNamespaceResponse.String"></a>
### func \(\*CreateNamespaceResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L51>)


## type [CreateNamespaceResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L255-L274>)

    // contains filtered or unexported fields
### func \(CreateNamespaceResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_namespace_response.pb.go#L276>)


## type [CreateSchemaRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L26-L37>)


    // contains filtered or unexported fields
### func \(\*CreateSchemaRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L138>)

<a name="CreateSchemaRequest.ClearMetadata"></a>
### func \(\*CreateSchemaRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L148>)

<a name="CreateSchemaRequest.ClearSchema"></a>
### func \(\*CreateSchemaRequest\) [ClearSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L143>)

<a name="CreateSchemaRequest.GetDatabase"></a>
### func \(\*CreateSchemaRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L64>)

<a name="CreateSchemaRequest.GetMetadata"></a>
### func \(\*CreateSchemaRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L84>)

<a name="CreateSchemaRequest.GetSchema"></a>
### func \(\*CreateSchemaRequest\) [GetSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L74>)

<a name="CreateSchemaRequest.HasDatabase"></a>
### func \(\*CreateSchemaRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L117>)

<a name="CreateSchemaRequest.HasMetadata"></a>
### func \(\*CreateSchemaRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L131>)

<a name="CreateSchemaRequest.HasSchema"></a>
### func \(\*CreateSchemaRequest\) [HasSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L124>)

<a name="CreateSchemaRequest.ProtoMessage"></a>
### func \(\*CreateSchemaRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L50>)

<a name="CreateSchemaRequest.ProtoReflect"></a>
### func \(\*CreateSchemaRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L52>)

<a name="CreateSchemaRequest.Reset"></a>
### func \(\*CreateSchemaRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L39>)

<a name="CreateSchemaRequest.SetDatabase"></a>
### func \(\*CreateSchemaRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L98>)

<a name="CreateSchemaRequest.SetMetadata"></a>
### func \(\*CreateSchemaRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L108>)

<a name="CreateSchemaRequest.SetSchema"></a>
### func \(\*CreateSchemaRequest\) [SetSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L103>)

<a name="CreateSchemaRequest.String"></a>
### func \(\*CreateSchemaRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L46>)


## type [CreateSchemaRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L153-L162>)

    // contains filtered or unexported fields
### func \(CreateSchemaRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_request.pb.go#L164>)
<a name="CreateSchemaResponse"></a>
## type [CreateSchemaResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L28-L38>)

\* CreateSchemaResponse contains the result of a schema creation operation. Indicates success status and provides error details if creation failed.

    // contains filtered or unexported fields
### func \(\*CreateSchemaResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L119>)

<a name="CreateSchemaResponse.ClearSuccess"></a>
### func \(\*CreateSchemaResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L114>)

<a name="CreateSchemaResponse.GetError"></a>
### func \(\*CreateSchemaResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L72>)

<a name="CreateSchemaResponse.GetSuccess"></a>
### func \(\*CreateSchemaResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L65>)

<a name="CreateSchemaResponse.HasError"></a>
### func \(\*CreateSchemaResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L107>)

<a name="CreateSchemaResponse.HasSuccess"></a>
### func \(\*CreateSchemaResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L100>)

<a name="CreateSchemaResponse.ProtoMessage"></a>
### func \(\*CreateSchemaResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L51>)

<a name="CreateSchemaResponse.ProtoReflect"></a>
### func \(\*CreateSchemaResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L53>)

<a name="CreateSchemaResponse.Reset"></a>
### func \(\*CreateSchemaResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L40>)

<a name="CreateSchemaResponse.SetError"></a>
### func \(\*CreateSchemaResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L91>)

<a name="CreateSchemaResponse.SetSuccess"></a>
### func \(\*CreateSchemaResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L86>)

<a name="CreateSchemaResponse.String"></a>
### func \(\*CreateSchemaResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L47>)


## type [CreateSchemaResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L124-L131>)

    // contains filtered or unexported fields
### func \(CreateSchemaResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/create_schema_response.pb.go#L133>)

<a name="DatabaseAdminServiceClient"></a>
## type [DatabaseAdminServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L39-L54>)
For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
\* DatabaseAdminService provides administrative operations for database management including schema operations and migrations.
### func [NewDatabaseAdminServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L60>)
<a name="DatabaseAdminServiceServer"></a>
## type [DatabaseAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L141-L157>)

DatabaseAdminServiceServer is the server API for DatabaseAdminService service. All implementations must embed UnimplementedDatabaseAdminServiceServer for forward compatibility.
\* DatabaseAdminService provides administrative operations for database management including schema operations and migrations.
    // contains filtered or unexported methods
## type [DatabaseBatchOperation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L28-L38>)
\* BatchOperation represents a single operation within a batch execution. Contains the SQL statement and its parameters for batch processing.

    // contains filtered or unexported fields
### func \(\*DatabaseBatchOperation\) [ClearStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L112>)

<a name="DatabaseBatchOperation.GetParameters"></a>
### func \(\*DatabaseBatchOperation\) [GetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L75>)

<a name="DatabaseBatchOperation.GetStatement"></a>
### func \(\*DatabaseBatchOperation\) [GetStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L65>)

<a name="DatabaseBatchOperation.HasStatement"></a>
### func \(\*DatabaseBatchOperation\) [HasStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L105>)

<a name="DatabaseBatchOperation.ProtoMessage"></a>
### func \(\*DatabaseBatchOperation\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L51>)

<a name="DatabaseBatchOperation.ProtoReflect"></a>
### func \(\*DatabaseBatchOperation\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L53>)

<a name="DatabaseBatchOperation.Reset"></a>
### func \(\*DatabaseBatchOperation\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L40>)

<a name="DatabaseBatchOperation.SetParameters"></a>
### func \(\*DatabaseBatchOperation\) [SetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L94>)

<a name="DatabaseBatchOperation.SetStatement"></a>
### func \(\*DatabaseBatchOperation\) [SetStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L89>)

<a name="DatabaseBatchOperation.String"></a>
### func \(\*DatabaseBatchOperation\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L47>)


## type [DatabaseBatchOperation\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L117-L124>)

    // contains filtered or unexported fields
### func \(DatabaseBatchOperation\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_operation.pb.go#L126>)
<a name="DatabaseBatchStats"></a>
## type [DatabaseBatchStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L29-L41>)

\* BatchStats provides execution statistics for batch database operations. Used for monitoring batch performance and operation success rates.

    // contains filtered or unexported fields
### func \(\*DatabaseBatchStats\) [ClearFailedOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L165>)

<a name="DatabaseBatchStats.ClearSuccessfulOperations"></a>
### func \(\*DatabaseBatchStats\) [ClearSuccessfulOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L160>)

<a name="DatabaseBatchStats.ClearTotalAffectedRows"></a>
### func \(\*DatabaseBatchStats\) [ClearTotalAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L170>)

<a name="DatabaseBatchStats.ClearTotalTime"></a>
### func \(\*DatabaseBatchStats\) [ClearTotalTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L155>)

<a name="DatabaseBatchStats.GetFailedOperations"></a>
### func \(\*DatabaseBatchStats\) [GetFailedOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L89>)

<a name="DatabaseBatchStats.GetSuccessfulOperations"></a>
### func \(\*DatabaseBatchStats\) [GetSuccessfulOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L82>)

<a name="DatabaseBatchStats.GetTotalAffectedRows"></a>
### func \(\*DatabaseBatchStats\) [GetTotalAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L96>)

<a name="DatabaseBatchStats.GetTotalTime"></a>
### func \(\*DatabaseBatchStats\) [GetTotalTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L68>)

<a name="DatabaseBatchStats.HasFailedOperations"></a>
### func \(\*DatabaseBatchStats\) [HasFailedOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L141>)

<a name="DatabaseBatchStats.HasSuccessfulOperations"></a>
### func \(\*DatabaseBatchStats\) [HasSuccessfulOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L134>)

<a name="DatabaseBatchStats.HasTotalAffectedRows"></a>
### func \(\*DatabaseBatchStats\) [HasTotalAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L148>)

<a name="DatabaseBatchStats.HasTotalTime"></a>
### func \(\*DatabaseBatchStats\) [HasTotalTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L127>)

<a name="DatabaseBatchStats.ProtoMessage"></a>
### func \(\*DatabaseBatchStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L54>)

<a name="DatabaseBatchStats.ProtoReflect"></a>
### func \(\*DatabaseBatchStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L56>)

<a name="DatabaseBatchStats.Reset"></a>
### func \(\*DatabaseBatchStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L43>)

<a name="DatabaseBatchStats.SetFailedOperations"></a>
### func \(\*DatabaseBatchStats\) [SetFailedOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L117>)

<a name="DatabaseBatchStats.SetSuccessfulOperations"></a>
### func \(\*DatabaseBatchStats\) [SetSuccessfulOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L112>)

<a name="DatabaseBatchStats.SetTotalAffectedRows"></a>
### func \(\*DatabaseBatchStats\) [SetTotalAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L122>)

<a name="DatabaseBatchStats.SetTotalTime"></a>
### func \(\*DatabaseBatchStats\) [SetTotalTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L103>)

<a name="DatabaseBatchStats.String"></a>
### func \(\*DatabaseBatchStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L50>)


## type [DatabaseBatchStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L175-L186>)

    // contains filtered or unexported fields
### func \(DatabaseBatchStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/batch_stats.pb.go#L188>)


## type [DatabaseHealthCheckRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L25-L34>)


    // contains filtered or unexported fields
### func \(\*DatabaseHealthCheckRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L91>)

<a name="DatabaseHealthCheckRequest.GetMetadata"></a>
### func \(\*DatabaseHealthCheckRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L61>)

<a name="DatabaseHealthCheckRequest.HasMetadata"></a>
### func \(\*DatabaseHealthCheckRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L84>)

<a name="DatabaseHealthCheckRequest.ProtoMessage"></a>
### func \(\*DatabaseHealthCheckRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L47>)

<a name="DatabaseHealthCheckRequest.ProtoReflect"></a>
### func \(\*DatabaseHealthCheckRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L49>)

<a name="DatabaseHealthCheckRequest.Reset"></a>
### func \(\*DatabaseHealthCheckRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L36>)

<a name="DatabaseHealthCheckRequest.SetMetadata"></a>
### func \(\*DatabaseHealthCheckRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L75>)

<a name="DatabaseHealthCheckRequest.String"></a>
### func \(\*DatabaseHealthCheckRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L43>)


## type [DatabaseHealthCheckRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L96-L101>)

    // contains filtered or unexported fields
### func \(DatabaseHealthCheckRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_request.pb.go#L103>)
<a name="DatabaseHealthCheckResponse"></a>
## type [DatabaseHealthCheckResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L29-L41>)

\* HealthCheckResponse contains the result of a database health check. Provides connection status, response time, and error details.

    // contains filtered or unexported fields
### func \(\*DatabaseHealthCheckResponse\) [ClearConnectionOk](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L173>)

<a name="DatabaseHealthCheckResponse.ClearError"></a>
### func \(\*DatabaseHealthCheckResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L183>)

<a name="DatabaseHealthCheckResponse.ClearResponseTime"></a>
### func \(\*DatabaseHealthCheckResponse\) [ClearResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L178>)

<a name="DatabaseHealthCheckResponse.ClearStatus"></a>
### func \(\*DatabaseHealthCheckResponse\) [ClearStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L168>)

<a name="DatabaseHealthCheckResponse.GetConnectionOk"></a>
### func \(\*DatabaseHealthCheckResponse\) [GetConnectionOk](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L77>)

<a name="DatabaseHealthCheckResponse.GetError"></a>
### func \(\*DatabaseHealthCheckResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L98>)

<a name="DatabaseHealthCheckResponse.GetResponseTime"></a>
### func \(\*DatabaseHealthCheckResponse\) [GetResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L84>)

<a name="DatabaseHealthCheckResponse.GetStatus"></a>
### func \(\*DatabaseHealthCheckResponse\) [GetStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L68>)

<a name="DatabaseHealthCheckResponse.HasConnectionOk"></a>
### func \(\*DatabaseHealthCheckResponse\) [HasConnectionOk](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L147>)

<a name="DatabaseHealthCheckResponse.HasError"></a>
### func \(\*DatabaseHealthCheckResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L161>)

<a name="DatabaseHealthCheckResponse.HasResponseTime"></a>
### func \(\*DatabaseHealthCheckResponse\) [HasResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L154>)

<a name="DatabaseHealthCheckResponse.HasStatus"></a>
### func \(\*DatabaseHealthCheckResponse\) [HasStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L140>)

<a name="DatabaseHealthCheckResponse.ProtoMessage"></a>
### func \(\*DatabaseHealthCheckResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L54>)

<a name="DatabaseHealthCheckResponse.ProtoReflect"></a>
### func \(\*DatabaseHealthCheckResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L56>)

<a name="DatabaseHealthCheckResponse.Reset"></a>
### func \(\*DatabaseHealthCheckResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L43>)

<a name="DatabaseHealthCheckResponse.SetConnectionOk"></a>
### func \(\*DatabaseHealthCheckResponse\) [SetConnectionOk](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L117>)

<a name="DatabaseHealthCheckResponse.SetError"></a>
### func \(\*DatabaseHealthCheckResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L131>)

<a name="DatabaseHealthCheckResponse.SetResponseTime"></a>
### func \(\*DatabaseHealthCheckResponse\) [SetResponseTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L122>)

<a name="DatabaseHealthCheckResponse.SetStatus"></a>
### func \(\*DatabaseHealthCheckResponse\) [SetStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L112>)

<a name="DatabaseHealthCheckResponse.String"></a>
### func \(\*DatabaseHealthCheckResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L50>)


## type [DatabaseHealthCheckResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L188-L199>)

    // contains filtered or unexported fields
### func \(DatabaseHealthCheckResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/health_check_response.pb.go#L201>)
<a name="DatabaseInfo"></a>
## type [DatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L28-L39>)

\* DatabaseInfo provides metadata about a database instance. Used for identifying database capabilities and connection details.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*DatabaseInfo\) [ClearConnectionString](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L180>)

<a name="DatabaseInfo.ClearName"></a>
### func \(\*DatabaseInfo\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L165>)

<a name="DatabaseInfo.ClearType"></a>
### func \(\*DatabaseInfo\) [ClearType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L175>)

<a name="DatabaseInfo.ClearVersion"></a>
### func \(\*DatabaseInfo\) [ClearVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L170>)

<a name="DatabaseInfo.GetConnectionString"></a>
### func \(\*DatabaseInfo\) [GetConnectionString](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L96>)

<a name="DatabaseInfo.GetFeatures"></a>
### func \(\*DatabaseInfo\) [GetFeatures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L106>)


### func \(\*DatabaseInfo\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L66>)

<a name="DatabaseInfo.GetType"></a>
### func \(\*DatabaseInfo\) [GetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L86>)

<a name="DatabaseInfo.GetVersion"></a>
### func \(\*DatabaseInfo\) [GetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L76>)

<a name="DatabaseInfo.HasConnectionString"></a>
### func \(\*DatabaseInfo\) [HasConnectionString](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L158>)

<a name="DatabaseInfo.HasName"></a>
### func \(\*DatabaseInfo\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L137>)

<a name="DatabaseInfo.HasType"></a>
### func \(\*DatabaseInfo\) [HasType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L151>)

<a name="DatabaseInfo.HasVersion"></a>
### func \(\*DatabaseInfo\) [HasVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L144>)

<a name="DatabaseInfo.ProtoMessage"></a>
### func \(\*DatabaseInfo\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L52>)

<a name="DatabaseInfo.ProtoReflect"></a>
### func \(\*DatabaseInfo\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L54>)

<a name="DatabaseInfo.Reset"></a>
### func \(\*DatabaseInfo\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L41>)

<a name="DatabaseInfo.SetConnectionString"></a>
### func \(\*DatabaseInfo\) [SetConnectionString](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L128>)

<a name="DatabaseInfo.SetFeatures"></a>
### func \(\*DatabaseInfo\) [SetFeatures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L133>)

<a name="DatabaseInfo.SetName"></a>
### func \(\*DatabaseInfo\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L113>)

<a name="DatabaseInfo.SetType"></a>
### func \(\*DatabaseInfo\) [SetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L123>)

<a name="DatabaseInfo.SetVersion"></a>
### func \(\*DatabaseInfo\) [SetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L118>)

<a name="DatabaseInfo.String"></a>
### func \(\*DatabaseInfo\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L48>)


## type [DatabaseInfo\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L185-L198>)

    // contains filtered or unexported fields
### func \(DatabaseInfo\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_info.pb.go#L200>)
<a name="DatabaseQueryStats"></a>
## type [DatabaseQueryStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L29-L42>)

\* QueryStats provides execution statistics for database queries. Used for performance monitoring and query optimization.

    // contains filtered or unexported fields
### func \(\*DatabaseQueryStats\) [ClearColumnCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L188>)

<a name="DatabaseQueryStats.ClearCostEstimate"></a>
### func \(\*DatabaseQueryStats\) [ClearCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L198>)

<a name="DatabaseQueryStats.ClearExecutionTime"></a>
### func \(\*DatabaseQueryStats\) [ClearExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L178>)

<a name="DatabaseQueryStats.ClearQueryPlan"></a>
### func \(\*DatabaseQueryStats\) [ClearQueryPlan](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L193>)

<a name="DatabaseQueryStats.ClearRowCount"></a>
### func \(\*DatabaseQueryStats\) [ClearRowCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L183>)

<a name="DatabaseQueryStats.GetColumnCount"></a>
### func \(\*DatabaseQueryStats\) [GetColumnCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L90>)

<a name="DatabaseQueryStats.GetCostEstimate"></a>
### func \(\*DatabaseQueryStats\) [GetCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L107>)

<a name="DatabaseQueryStats.GetExecutionTime"></a>
### func \(\*DatabaseQueryStats\) [GetExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L69>)

<a name="DatabaseQueryStats.GetQueryPlan"></a>
### func \(\*DatabaseQueryStats\) [GetQueryPlan](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L97>)

<a name="DatabaseQueryStats.GetRowCount"></a>
### func \(\*DatabaseQueryStats\) [GetRowCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L83>)

<a name="DatabaseQueryStats.HasColumnCount"></a>
### func \(\*DatabaseQueryStats\) [HasColumnCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L157>)

<a name="DatabaseQueryStats.HasCostEstimate"></a>
### func \(\*DatabaseQueryStats\) [HasCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L171>)

<a name="DatabaseQueryStats.HasExecutionTime"></a>
### func \(\*DatabaseQueryStats\) [HasExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L143>)

<a name="DatabaseQueryStats.HasQueryPlan"></a>
### func \(\*DatabaseQueryStats\) [HasQueryPlan](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L164>)

<a name="DatabaseQueryStats.HasRowCount"></a>
### func \(\*DatabaseQueryStats\) [HasRowCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L150>)

<a name="DatabaseQueryStats.ProtoMessage"></a>
### func \(\*DatabaseQueryStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L55>)

<a name="DatabaseQueryStats.ProtoReflect"></a>
### func \(\*DatabaseQueryStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L57>)

<a name="DatabaseQueryStats.Reset"></a>
### func \(\*DatabaseQueryStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L44>)

<a name="DatabaseQueryStats.SetColumnCount"></a>
### func \(\*DatabaseQueryStats\) [SetColumnCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L128>)

<a name="DatabaseQueryStats.SetCostEstimate"></a>
### func \(\*DatabaseQueryStats\) [SetCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L138>)

<a name="DatabaseQueryStats.SetExecutionTime"></a>
### func \(\*DatabaseQueryStats\) [SetExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L114>)

<a name="DatabaseQueryStats.SetQueryPlan"></a>
### func \(\*DatabaseQueryStats\) [SetQueryPlan](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L133>)

<a name="DatabaseQueryStats.SetRowCount"></a>
### func \(\*DatabaseQueryStats\) [SetRowCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L123>)

<a name="DatabaseQueryStats.String"></a>
### func \(\*DatabaseQueryStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L51>)


## type [DatabaseQueryStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L203-L216>)

    // contains filtered or unexported fields
### func \(DatabaseQueryStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_stats.pb.go#L218>)

<a name="DatabaseServiceClient"></a>
## type [DatabaseServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L37-L50>)
For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
\* DatabaseService provides comprehensive database operations including queries, transactions, batch operations, and health monitoring.
### func [NewDatabaseServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L56>)
<a name="DatabaseServiceServer"></a>
## type [DatabaseServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L127-L141>)

DatabaseServiceServer is the server API for DatabaseService service. All implementations must embed UnimplementedDatabaseServiceServer for forward compatibility.
\* DatabaseService provides comprehensive database operations including queries, transactions, batch operations, and health monitoring.
    // contains filtered or unexported methods
## type [DatabaseStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L29-L37>)
\* DatabaseStatus reports the current connection status for a database driver or service.
    // contains filtered or unexported fields
### func \(\*DatabaseStatus\) [ClearCode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L107>)

<a name="DatabaseStatus.ClearMessage"></a>
### func \(\*DatabaseStatus\) [ClearMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L112>)

<a name="DatabaseStatus.GetCode"></a>
### func \(\*DatabaseStatus\) [GetCode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L64>)

<a name="DatabaseStatus.GetMessage"></a>
### func \(\*DatabaseStatus\) [GetMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L73>)

<a name="DatabaseStatus.HasCode"></a>
### func \(\*DatabaseStatus\) [HasCode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L93>)

<a name="DatabaseStatus.HasMessage"></a>
### func \(\*DatabaseStatus\) [HasMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L100>)

<a name="DatabaseStatus.ProtoMessage"></a>
### func \(\*DatabaseStatus\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L50>)

<a name="DatabaseStatus.ProtoReflect"></a>
### func \(\*DatabaseStatus\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L52>)

<a name="DatabaseStatus.Reset"></a>
### func \(\*DatabaseStatus\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L39>)

<a name="DatabaseStatus.SetCode"></a>
### func \(\*DatabaseStatus\) [SetCode](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L83>)

<a name="DatabaseStatus.SetMessage"></a>
### func \(\*DatabaseStatus\) [SetMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L88>)

<a name="DatabaseStatus.String"></a>
### func \(\*DatabaseStatus\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L46>)


## type [DatabaseStatus\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L117-L124>)

    // contains filtered or unexported fields
### func \(DatabaseStatus\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_status.pb.go#L126>)

<a name="DecrementRequest"></a>
## type [DecrementRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L29-L43>)

    // contains filtered or unexported fields
### func \(\*DecrementRequest\) [ClearDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L217>)

<a name="DecrementRequest.ClearInitialValue"></a>
### func \(\*DecrementRequest\) [ClearInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L222>)

<a name="DecrementRequest.ClearKey"></a>
### func \(\*DecrementRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L212>)

<a name="DecrementRequest.ClearMetadata"></a>
### func \(\*DecrementRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L237>)

<a name="DecrementRequest.ClearNamespace"></a>
### func \(\*DecrementRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L232>)

<a name="DecrementRequest.ClearTtl"></a>
### func \(\*DecrementRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L227>)

<a name="DecrementRequest.GetDelta"></a>
### func \(\*DecrementRequest\) [GetDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L80>)

<a name="DecrementRequest.GetInitialValue"></a>
### func \(\*DecrementRequest\) [GetInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L87>)

<a name="DecrementRequest.GetKey"></a>
### func \(\*DecrementRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L70>)

<a name="DecrementRequest.GetMetadata"></a>
### func \(\*DecrementRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L118>)

<a name="DecrementRequest.GetNamespace"></a>
### func \(\*DecrementRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L108>)

<a name="DecrementRequest.GetTtl"></a>
### func \(\*DecrementRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L94>)

<a name="DecrementRequest.HasDelta"></a>
### func \(\*DecrementRequest\) [HasDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L177>)

<a name="DecrementRequest.HasInitialValue"></a>
### func \(\*DecrementRequest\) [HasInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L184>)

<a name="DecrementRequest.HasKey"></a>
### func \(\*DecrementRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L170>)

<a name="DecrementRequest.HasMetadata"></a>
### func \(\*DecrementRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L205>)

<a name="DecrementRequest.HasNamespace"></a>
### func \(\*DecrementRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L198>)

<a name="DecrementRequest.HasTtl"></a>
### func \(\*DecrementRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L191>)

<a name="DecrementRequest.ProtoMessage"></a>
### func \(\*DecrementRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L56>)

<a name="DecrementRequest.ProtoReflect"></a>
### func \(\*DecrementRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L58>)

<a name="DecrementRequest.Reset"></a>
### func \(\*DecrementRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L45>)

<a name="DecrementRequest.SetDelta"></a>
### func \(\*DecrementRequest\) [SetDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L137>)

<a name="DecrementRequest.SetInitialValue"></a>
### func \(\*DecrementRequest\) [SetInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L142>)

<a name="DecrementRequest.SetKey"></a>
### func \(\*DecrementRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L132>)

<a name="DecrementRequest.SetMetadata"></a>
### func \(\*DecrementRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L161>)

<a name="DecrementRequest.SetNamespace"></a>
### func \(\*DecrementRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L156>)


### func \(\*DecrementRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L147>)

<a name="DecrementRequest.String"></a>
### func \(\*DecrementRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L52>)


## type [DecrementRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L242-L257>)

    // contains filtered or unexported fields
### func \(DecrementRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_request.pb.go#L259>)
<a name="DecrementResponse"></a>
## type [DecrementResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L29-L38>)

\* Response for cache decrement operations. Returns the new value after decrementing.
    // contains filtered or unexported fields
### func \(\*DecrementResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L131>)

<a name="DecrementResponse.ClearNewValue"></a>
### func \(\*DecrementResponse\) [ClearNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L121>)

<a name="DecrementResponse.ClearSuccess"></a>
### func \(\*DecrementResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L126>)

<a name="DecrementResponse.GetError"></a>
### func \(\*DecrementResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L79>)

<a name="DecrementResponse.GetNewValue"></a>
### func \(\*DecrementResponse\) [GetNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L65>)

<a name="DecrementResponse.GetSuccess"></a>
### func \(\*DecrementResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L72>)

<a name="DecrementResponse.HasError"></a>
### func \(\*DecrementResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L114>)

<a name="DecrementResponse.HasNewValue"></a>
### func \(\*DecrementResponse\) [HasNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L100>)

<a name="DecrementResponse.HasSuccess"></a>
### func \(\*DecrementResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L107>)

<a name="DecrementResponse.ProtoMessage"></a>
### func \(\*DecrementResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L51>)

<a name="DecrementResponse.ProtoReflect"></a>
### func \(\*DecrementResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L53>)

<a name="DecrementResponse.Reset"></a>
### func \(\*DecrementResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L40>)

<a name="DecrementResponse.SetError"></a>
### func \(\*DecrementResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L96>)

<a name="DecrementResponse.SetNewValue"></a>
### func \(\*DecrementResponse\) [SetNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L86>)

<a name="DecrementResponse.SetSuccess"></a>
### func \(\*DecrementResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L91>)

<a name="DecrementResponse.String"></a>
### func \(\*DecrementResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L47>)


## type [DecrementResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L135-L144>)

    // contains filtered or unexported fields
### func \(DecrementResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/decrement_response.pb.go#L146>)

<a name="DefragRequest"></a>
## type [DefragRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*DefragRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L122>)

<a name="DefragRequest.ClearNamespace"></a>
### func \(\*DefragRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L117>)

<a name="DefragRequest.GetMetadata"></a>
### func \(\*DefragRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L75>)

<a name="DefragRequest.GetNamespace"></a>
### func \(\*DefragRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L65>)

<a name="DefragRequest.HasMetadata"></a>
### func \(\*DefragRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L110>)

<a name="DefragRequest.HasNamespace"></a>
### func \(\*DefragRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L103>)

<a name="DefragRequest.ProtoMessage"></a>
### func \(\*DefragRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L51>)

<a name="DefragRequest.ProtoReflect"></a>
### func \(\*DefragRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L53>)

<a name="DefragRequest.Reset"></a>
### func \(\*DefragRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L40>)

<a name="DefragRequest.SetMetadata"></a>
### func \(\*DefragRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L94>)

<a name="DefragRequest.SetNamespace"></a>
### func \(\*DefragRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L89>)

<a name="DefragRequest.String"></a>
### func \(\*DefragRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L47>)


## type [DefragRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(DefragRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/defrag_request.pb.go#L136>)

<a name="DeleteMultipleRequest"></a>
## type [DeleteMultipleRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*DeleteMultipleRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L134>)

<a name="DeleteMultipleRequest.ClearNamespace"></a>
### func \(\*DeleteMultipleRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L129>)

<a name="DeleteMultipleRequest.GetKeys"></a>
### func \(\*DeleteMultipleRequest\) [GetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L66>)

<a name="DeleteMultipleRequest.GetMetadata"></a>
### func \(\*DeleteMultipleRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L83>)

<a name="DeleteMultipleRequest.GetNamespace"></a>
### func \(\*DeleteMultipleRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L73>)

<a name="DeleteMultipleRequest.HasMetadata"></a>
### func \(\*DeleteMultipleRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L122>)

<a name="DeleteMultipleRequest.HasNamespace"></a>
### func \(\*DeleteMultipleRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L115>)

<a name="DeleteMultipleRequest.ProtoMessage"></a>
### func \(\*DeleteMultipleRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L52>)

<a name="DeleteMultipleRequest.ProtoReflect"></a>
### func \(\*DeleteMultipleRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L54>)

<a name="DeleteMultipleRequest.Reset"></a>
### func \(\*DeleteMultipleRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L41>)

<a name="DeleteMultipleRequest.SetKeys"></a>
### func \(\*DeleteMultipleRequest\) [SetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L97>)

<a name="DeleteMultipleRequest.SetMetadata"></a>
### func \(\*DeleteMultipleRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L106>)

<a name="DeleteMultipleRequest.SetNamespace"></a>
### func \(\*DeleteMultipleRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L101>)

<a name="DeleteMultipleRequest.String"></a>
### func \(\*DeleteMultipleRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L48>)


## type [DeleteMultipleRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L139-L148>)

    // contains filtered or unexported fields
### func \(DeleteMultipleRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_request.pb.go#L150>)
<a name="DeleteMultipleResponse"></a>
## type [DeleteMultipleResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L29-L39>)

\* Response for cache delete multiple operations. Indicates success/failure of multiple key deletions.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*DeleteMultipleResponse\) [ClearDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L133>)

<a name="DeleteMultipleResponse.ClearError"></a>
### func \(\*DeleteMultipleResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L143>)

<a name="DeleteMultipleResponse.ClearFailedCount"></a>
### func \(\*DeleteMultipleResponse\) [ClearFailedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L138>)

<a name="DeleteMultipleResponse.GetDeletedCount"></a>
### func \(\*DeleteMultipleResponse\) [GetDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L66>)

<a name="DeleteMultipleResponse.GetError"></a>
### func \(\*DeleteMultipleResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L87>)

<a name="DeleteMultipleResponse.GetFailedCount"></a>
### func \(\*DeleteMultipleResponse\) [GetFailedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L73>)

<a name="DeleteMultipleResponse.GetFailedKeys"></a>
### func \(\*DeleteMultipleResponse\) [GetFailedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L80>)

<a name="DeleteMultipleResponse.HasDeletedCount"></a>
### func \(\*DeleteMultipleResponse\) [HasDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L112>)

<a name="DeleteMultipleResponse.HasError"></a>
### func \(\*DeleteMultipleResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L126>)

<a name="DeleteMultipleResponse.HasFailedCount"></a>
### func \(\*DeleteMultipleResponse\) [HasFailedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L119>)

<a name="DeleteMultipleResponse.ProtoMessage"></a>
### func \(\*DeleteMultipleResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L52>)

<a name="DeleteMultipleResponse.ProtoReflect"></a>
### func \(\*DeleteMultipleResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L54>)

<a name="DeleteMultipleResponse.Reset"></a>
### func \(\*DeleteMultipleResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L41>)

<a name="DeleteMultipleResponse.SetDeletedCount"></a>
### func \(\*DeleteMultipleResponse\) [SetDeletedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L94>)

<a name="DeleteMultipleResponse.SetError"></a>
### func \(\*DeleteMultipleResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L108>)

<a name="DeleteMultipleResponse.SetFailedCount"></a>
### func \(\*DeleteMultipleResponse\) [SetFailedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L99>)

<a name="DeleteMultipleResponse.SetFailedKeys"></a>
### func \(\*DeleteMultipleResponse\) [SetFailedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L104>)

<a name="DeleteMultipleResponse.String"></a>
### func \(\*DeleteMultipleResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L48>)


## type [DeleteMultipleResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L147-L158>)

    // contains filtered or unexported fields
### func \(DeleteMultipleResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_multiple_response.pb.go#L160>)

<a name="DeleteNamespaceRequest"></a>
## type [DeleteNamespaceRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L27-L36>)
    // contains filtered or unexported fields
### func \(\*DeleteNamespaceRequest\) [ClearBackup](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L133>)

<a name="DeleteNamespaceRequest.ClearForce"></a>
### func \(\*DeleteNamespaceRequest\) [ClearForce](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L128>)

<a name="DeleteNamespaceRequest.ClearNamespaceId"></a>
### func \(\*DeleteNamespaceRequest\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L123>)

<a name="DeleteNamespaceRequest.GetBackup"></a>
### func \(\*DeleteNamespaceRequest\) [GetBackup](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L80>)

<a name="DeleteNamespaceRequest.GetForce"></a>
### func \(\*DeleteNamespaceRequest\) [GetForce](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L73>)

<a name="DeleteNamespaceRequest.GetNamespaceId"></a>
### func \(\*DeleteNamespaceRequest\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L63>)

<a name="DeleteNamespaceRequest.HasBackup"></a>
### func \(\*DeleteNamespaceRequest\) [HasBackup](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L116>)

<a name="DeleteNamespaceRequest.HasForce"></a>
### func \(\*DeleteNamespaceRequest\) [HasForce](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L109>)

<a name="DeleteNamespaceRequest.HasNamespaceId"></a>
### func \(\*DeleteNamespaceRequest\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L102>)

<a name="DeleteNamespaceRequest.ProtoMessage"></a>
### func \(\*DeleteNamespaceRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L49>)

<a name="DeleteNamespaceRequest.ProtoReflect"></a>
### func \(\*DeleteNamespaceRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L51>)

<a name="DeleteNamespaceRequest.Reset"></a>
### func \(\*DeleteNamespaceRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L38>)

<a name="DeleteNamespaceRequest.SetBackup"></a>
### func \(\*DeleteNamespaceRequest\) [SetBackup](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L97>)

<a name="DeleteNamespaceRequest.SetForce"></a>
### func \(\*DeleteNamespaceRequest\) [SetForce](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L92>)

<a name="DeleteNamespaceRequest.SetNamespaceId"></a>
### func \(\*DeleteNamespaceRequest\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L87>)

<a name="DeleteNamespaceRequest.String"></a>
### func \(\*DeleteNamespaceRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L45>)


## type [DeleteNamespaceRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L138-L147>)

    // contains filtered or unexported fields
### func \(DeleteNamespaceRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/delete_namespace_request.pb.go#L149>)


## type [DropDatabaseRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L26-L36>)


    // contains filtered or unexported fields
### func \(\*DropDatabaseRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L120>)

<a name="DropDatabaseRequest.ClearName"></a>
### func \(\*DropDatabaseRequest\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L115>)

<a name="DropDatabaseRequest.GetMetadata"></a>
### func \(\*DropDatabaseRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L73>)

<a name="DropDatabaseRequest.GetName"></a>
### func \(\*DropDatabaseRequest\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L63>)

<a name="DropDatabaseRequest.HasMetadata"></a>
### func \(\*DropDatabaseRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L108>)

<a name="DropDatabaseRequest.HasName"></a>
### func \(\*DropDatabaseRequest\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L101>)

<a name="DropDatabaseRequest.ProtoMessage"></a>
### func \(\*DropDatabaseRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L49>)

<a name="DropDatabaseRequest.ProtoReflect"></a>
### func \(\*DropDatabaseRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L51>)

<a name="DropDatabaseRequest.Reset"></a>
### func \(\*DropDatabaseRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L38>)

<a name="DropDatabaseRequest.SetMetadata"></a>
### func \(\*DropDatabaseRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L92>)

<a name="DropDatabaseRequest.SetName"></a>
### func \(\*DropDatabaseRequest\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L87>)

<a name="DropDatabaseRequest.String"></a>
### func \(\*DropDatabaseRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L45>)


## type [DropDatabaseRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L125-L132>)

    // contains filtered or unexported fields
### func \(DropDatabaseRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_database_request.pb.go#L134>)


## type [DropSchemaRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L26-L37>)


    // contains filtered or unexported fields
### func \(\*DropSchemaRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L138>)

<a name="DropSchemaRequest.ClearMetadata"></a>
### func \(\*DropSchemaRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L148>)

<a name="DropSchemaRequest.ClearSchema"></a>
### func \(\*DropSchemaRequest\) [ClearSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L143>)

<a name="DropSchemaRequest.GetDatabase"></a>
### func \(\*DropSchemaRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L64>)

<a name="DropSchemaRequest.GetMetadata"></a>
### func \(\*DropSchemaRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L84>)

<a name="DropSchemaRequest.GetSchema"></a>
### func \(\*DropSchemaRequest\) [GetSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L74>)

<a name="DropSchemaRequest.HasDatabase"></a>
### func \(\*DropSchemaRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L117>)

<a name="DropSchemaRequest.HasMetadata"></a>
### func \(\*DropSchemaRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L131>)

<a name="DropSchemaRequest.HasSchema"></a>
### func \(\*DropSchemaRequest\) [HasSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L124>)

<a name="DropSchemaRequest.ProtoMessage"></a>
### func \(\*DropSchemaRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L50>)

<a name="DropSchemaRequest.ProtoReflect"></a>
### func \(\*DropSchemaRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L52>)

<a name="DropSchemaRequest.Reset"></a>
### func \(\*DropSchemaRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L39>)

<a name="DropSchemaRequest.SetDatabase"></a>
### func \(\*DropSchemaRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L98>)

<a name="DropSchemaRequest.SetMetadata"></a>
### func \(\*DropSchemaRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L108>)

<a name="DropSchemaRequest.SetSchema"></a>
### func \(\*DropSchemaRequest\) [SetSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L103>)

<a name="DropSchemaRequest.String"></a>
### func \(\*DropSchemaRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L46>)


## type [DropSchemaRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L153-L162>)

    // contains filtered or unexported fields
### func \(DropSchemaRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/drop_schema_request.pb.go#L164>)
<a name="EvictionResult"></a>
## type [EvictionResult](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L30-L43>)

\* Result of cache eviction operations. Provides details about items removed from cache.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*EvictionResult\) [ClearEvictedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L214>)

<a name="EvictionResult.ClearEvictedCount"></a>
### func \(\*EvictionResult\) [ClearEvictedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L199>)

<a name="EvictionResult.ClearEvictionReason"></a>
### func \(\*EvictionResult\) [ClearEvictionReason](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L209>)

<a name="EvictionResult.ClearMemoryFreed"></a>
### func \(\*EvictionResult\) [ClearMemoryFreed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L218>)

<a name="EvictionResult.ClearPolicyUsed"></a>
### func \(\*EvictionResult\) [ClearPolicyUsed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L204>)

<a name="EvictionResult.ClearSuccess"></a>
### func \(\*EvictionResult\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L223>)

<a name="EvictionResult.GetEvictedAt"></a>
### func \(\*EvictionResult\) [GetEvictedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L103>)

<a name="EvictionResult.GetEvictedCount"></a>
### func \(\*EvictionResult\) [GetEvictedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L70>)

<a name="EvictionResult.GetEvictedKeys"></a>
### func \(\*EvictionResult\) [GetEvictedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L77>)

<a name="EvictionResult.GetEvictionReason"></a>
### func \(\*EvictionResult\) [GetEvictionReason](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L93>)

<a name="EvictionResult.GetMemoryFreed"></a>
### func \(\*EvictionResult\) [GetMemoryFreed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L110>)

<a name="EvictionResult.GetPolicyUsed"></a>
### func \(\*EvictionResult\) [GetPolicyUsed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L84>)

<a name="EvictionResult.GetSuccess"></a>
### func \(\*EvictionResult\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L117>)

<a name="EvictionResult.HasEvictedAt"></a>
### func \(\*EvictionResult\) [HasEvictedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L178>)

<a name="EvictionResult.HasEvictedCount"></a>
### func \(\*EvictionResult\) [HasEvictedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L157>)

<a name="EvictionResult.HasEvictionReason"></a>
### func \(\*EvictionResult\) [HasEvictionReason](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L171>)

<a name="EvictionResult.HasMemoryFreed"></a>
### func \(\*EvictionResult\) [HasMemoryFreed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L185>)

<a name="EvictionResult.HasPolicyUsed"></a>
### func \(\*EvictionResult\) [HasPolicyUsed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L164>)

<a name="EvictionResult.HasSuccess"></a>
### func \(\*EvictionResult\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L192>)

<a name="EvictionResult.ProtoMessage"></a>
### func \(\*EvictionResult\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L56>)

<a name="EvictionResult.ProtoReflect"></a>
### func \(\*EvictionResult\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L58>)

<a name="EvictionResult.Reset"></a>
### func \(\*EvictionResult\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L45>)

<a name="EvictionResult.SetEvictedAt"></a>
### func \(\*EvictionResult\) [SetEvictedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L143>)

<a name="EvictionResult.SetEvictedCount"></a>
### func \(\*EvictionResult\) [SetEvictedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L124>)

<a name="EvictionResult.SetEvictedKeys"></a>
### func \(\*EvictionResult\) [SetEvictedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L129>)

<a name="EvictionResult.SetEvictionReason"></a>
### func \(\*EvictionResult\) [SetEvictionReason](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L138>)

<a name="EvictionResult.SetMemoryFreed"></a>
### func \(\*EvictionResult\) [SetMemoryFreed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L147>)

<a name="EvictionResult.SetPolicyUsed"></a>
### func \(\*EvictionResult\) [SetPolicyUsed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L133>)

<a name="EvictionResult.SetSuccess"></a>
### func \(\*EvictionResult\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L152>)

<a name="EvictionResult.String"></a>
### func \(\*EvictionResult\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L52>)


## type [EvictionResult\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L228-L245>)

    // contains filtered or unexported fields
### func \(EvictionResult\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/eviction_result.pb.go#L247>)


## type [ExecuteBatchRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L26-L39>)


    // contains filtered or unexported fields
### func \(\*ExecuteBatchRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L195>)

<a name="ExecuteBatchRequest.ClearMetadata"></a>
### func \(\*ExecuteBatchRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L205>)

<a name="ExecuteBatchRequest.ClearOptions"></a>
### func \(\*ExecuteBatchRequest\) [ClearOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L200>)

<a name="ExecuteBatchRequest.ClearTransactionId"></a>
### func \(\*ExecuteBatchRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L210>)

<a name="ExecuteBatchRequest.GetDatabase"></a>
### func \(\*ExecuteBatchRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L80>)

<a name="ExecuteBatchRequest.GetMetadata"></a>
### func \(\*ExecuteBatchRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L104>)

<a name="ExecuteBatchRequest.GetOperations"></a>
### func \(\*ExecuteBatchRequest\) [GetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L66>)

<a name="ExecuteBatchRequest.GetOptions"></a>
### func \(\*ExecuteBatchRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L90>)

<a name="ExecuteBatchRequest.GetTransactionId"></a>
### func \(\*ExecuteBatchRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L118>)

<a name="ExecuteBatchRequest.HasDatabase"></a>
### func \(\*ExecuteBatchRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L167>)

<a name="ExecuteBatchRequest.HasMetadata"></a>
### func \(\*ExecuteBatchRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L181>)

<a name="ExecuteBatchRequest.HasOptions"></a>
### func \(\*ExecuteBatchRequest\) [HasOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L174>)

<a name="ExecuteBatchRequest.HasTransactionId"></a>
### func \(\*ExecuteBatchRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L188>)

<a name="ExecuteBatchRequest.ProtoMessage"></a>
### func \(\*ExecuteBatchRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L52>)

<a name="ExecuteBatchRequest.ProtoReflect"></a>
### func \(\*ExecuteBatchRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L54>)

<a name="ExecuteBatchRequest.Reset"></a>
### func \(\*ExecuteBatchRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L41>)

<a name="ExecuteBatchRequest.SetDatabase"></a>
### func \(\*ExecuteBatchRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L139>)

<a name="ExecuteBatchRequest.SetMetadata"></a>
### func \(\*ExecuteBatchRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L153>)

<a name="ExecuteBatchRequest.SetOperations"></a>
### func \(\*ExecuteBatchRequest\) [SetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L128>)

<a name="ExecuteBatchRequest.SetOptions"></a>
### func \(\*ExecuteBatchRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L144>)

<a name="ExecuteBatchRequest.SetTransactionId"></a>
### func \(\*ExecuteBatchRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L162>)

<a name="ExecuteBatchRequest.String"></a>
### func \(\*ExecuteBatchRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L48>)


## type [ExecuteBatchRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L215-L228>)

    // contains filtered or unexported fields
### func \(ExecuteBatchRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_request.pb.go#L230>)
<a name="ExecuteBatchResponse"></a>
## type [ExecuteBatchResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L29-L40>)

\* ExecuteBatchResponse contains the results of a batch database operation. Includes individual operation results and overall batch statistics.

    // contains filtered or unexported fields
### func \(\*ExecuteBatchResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L157>)

<a name="ExecuteBatchResponse.ClearStats"></a>
### func \(\*ExecuteBatchResponse\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L152>)

<a name="ExecuteBatchResponse.GetError"></a>
### func \(\*ExecuteBatchResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L95>)

<a name="ExecuteBatchResponse.GetResults"></a>
### func \(\*ExecuteBatchResponse\) [GetResults](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L67>)

<a name="ExecuteBatchResponse.GetStats"></a>
### func \(\*ExecuteBatchResponse\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L81>)

<a name="ExecuteBatchResponse.HasError"></a>
### func \(\*ExecuteBatchResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L145>)

<a name="ExecuteBatchResponse.HasStats"></a>
### func \(\*ExecuteBatchResponse\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L138>)

<a name="ExecuteBatchResponse.ProtoMessage"></a>
### func \(\*ExecuteBatchResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L53>)

<a name="ExecuteBatchResponse.ProtoReflect"></a>
### func \(\*ExecuteBatchResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L55>)

<a name="ExecuteBatchResponse.Reset"></a>
### func \(\*ExecuteBatchResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L42>)

<a name="ExecuteBatchResponse.SetError"></a>
### func \(\*ExecuteBatchResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L129>)

<a name="ExecuteBatchResponse.SetResults"></a>
### func \(\*ExecuteBatchResponse\) [SetResults](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L109>)

<a name="ExecuteBatchResponse.SetStats"></a>
### func \(\*ExecuteBatchResponse\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L120>)

<a name="ExecuteBatchResponse.String"></a>
### func \(\*ExecuteBatchResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L49>)


## type [ExecuteBatchResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L162-L171>)

    // contains filtered or unexported fields
### func \(ExecuteBatchResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_batch_response.pb.go#L173>)
<a name="ExecuteOptions"></a>
## type [ExecuteOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L29-L40>)

\* ExecuteOptions configures behavior for database execute operations. Controls timeouts, transaction isolation, and result handling.

    // contains filtered or unexported fields
### func \(\*ExecuteOptions\) [ClearIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L147>)

<a name="ExecuteOptions.ClearReturnGeneratedKeys"></a>
### func \(\*ExecuteOptions\) [ClearReturnGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L142>)

<a name="ExecuteOptions.ClearTimeout"></a>
### func \(\*ExecuteOptions\) [ClearTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L137>)

<a name="ExecuteOptions.GetIsolation"></a>
### func \(\*ExecuteOptions\) [GetIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L88>)

<a name="ExecuteOptions.GetReturnGeneratedKeys"></a>
### func \(\*ExecuteOptions\) [GetReturnGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L81>)

<a name="ExecuteOptions.GetTimeout"></a>
### func \(\*ExecuteOptions\) [GetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L67>)

<a name="ExecuteOptions.HasIsolation"></a>
### func \(\*ExecuteOptions\) [HasIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L130>)

<a name="ExecuteOptions.HasReturnGeneratedKeys"></a>
### func \(\*ExecuteOptions\) [HasReturnGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L123>)

<a name="ExecuteOptions.HasTimeout"></a>
### func \(\*ExecuteOptions\) [HasTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L116>)

<a name="ExecuteOptions.ProtoMessage"></a>
### func \(\*ExecuteOptions\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L53>)

<a name="ExecuteOptions.ProtoReflect"></a>
### func \(\*ExecuteOptions\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L55>)

<a name="ExecuteOptions.Reset"></a>
### func \(\*ExecuteOptions\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L42>)

<a name="ExecuteOptions.SetIsolation"></a>
### func \(\*ExecuteOptions\) [SetIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L111>)

<a name="ExecuteOptions.SetReturnGeneratedKeys"></a>
### func \(\*ExecuteOptions\) [SetReturnGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L106>)

<a name="ExecuteOptions.SetTimeout"></a>
### func \(\*ExecuteOptions\) [SetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L97>)

<a name="ExecuteOptions.String"></a>
### func \(\*ExecuteOptions\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L49>)


## type [ExecuteOptions\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L152-L161>)

    // contains filtered or unexported fields
### func \(ExecuteOptions\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_options.pb.go#L163>)


## type [ExecuteRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L26-L40>)


    // contains filtered or unexported fields
### func \(\*ExecuteRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L223>)

<a name="ExecuteRequest.ClearMetadata"></a>
### func \(\*ExecuteRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L233>)

<a name="ExecuteRequest.ClearOptions"></a>
### func \(\*ExecuteRequest\) [ClearOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L228>)

<a name="ExecuteRequest.ClearStatement"></a>
### func \(\*ExecuteRequest\) [ClearStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L218>)

<a name="ExecuteRequest.ClearTransactionId"></a>
### func \(\*ExecuteRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L238>)

<a name="ExecuteRequest.GetDatabase"></a>
### func \(\*ExecuteRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L91>)

<a name="ExecuteRequest.GetMetadata"></a>
### func \(\*ExecuteRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L115>)

<a name="ExecuteRequest.GetOptions"></a>
### func \(\*ExecuteRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L101>)

<a name="ExecuteRequest.GetParameters"></a>
### func \(\*ExecuteRequest\) [GetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L77>)

<a name="ExecuteRequest.GetStatement"></a>
### func \(\*ExecuteRequest\) [GetStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L67>)

<a name="ExecuteRequest.GetTransactionId"></a>
### func \(\*ExecuteRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L129>)

<a name="ExecuteRequest.HasDatabase"></a>
### func \(\*ExecuteRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L190>)

<a name="ExecuteRequest.HasMetadata"></a>
### func \(\*ExecuteRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L204>)

<a name="ExecuteRequest.HasOptions"></a>
### func \(\*ExecuteRequest\) [HasOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L197>)

<a name="ExecuteRequest.HasStatement"></a>
### func \(\*ExecuteRequest\) [HasStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L183>)

<a name="ExecuteRequest.HasTransactionId"></a>
### func \(\*ExecuteRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L211>)

<a name="ExecuteRequest.ProtoMessage"></a>
### func \(\*ExecuteRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L53>)

<a name="ExecuteRequest.ProtoReflect"></a>
### func \(\*ExecuteRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L55>)

<a name="ExecuteRequest.Reset"></a>
### func \(\*ExecuteRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L42>)

<a name="ExecuteRequest.SetDatabase"></a>
### func \(\*ExecuteRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L155>)

<a name="ExecuteRequest.SetMetadata"></a>
### func \(\*ExecuteRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L169>)

<a name="ExecuteRequest.SetOptions"></a>
### func \(\*ExecuteRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L160>)

<a name="ExecuteRequest.SetParameters"></a>
### func \(\*ExecuteRequest\) [SetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L144>)

<a name="ExecuteRequest.SetStatement"></a>
### func \(\*ExecuteRequest\) [SetStatement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L139>)

<a name="ExecuteRequest.SetTransactionId"></a>
### func \(\*ExecuteRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L178>)

<a name="ExecuteRequest.String"></a>
### func \(\*ExecuteRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L49>)


## type [ExecuteRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L243-L258>)

    // contains filtered or unexported fields
### func \(ExecuteRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_request.pb.go#L260>)
<a name="ExecuteResponse"></a>
## type [ExecuteResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L30-L42>)

\* ExecuteResponse contains the results of a database execute operation. Includes affected row count, generated keys, and execution statistics.

    // contains filtered or unexported fields
### func \(\*ExecuteResponse\) [ClearAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L173>)

<a name="ExecuteResponse.ClearError"></a>
### func \(\*ExecuteResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L183>)

<a name="ExecuteResponse.ClearStats"></a>
### func \(\*ExecuteResponse\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L178>)

<a name="ExecuteResponse.GetAffectedRows"></a>
### func \(\*ExecuteResponse\) [GetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L69>)

<a name="ExecuteResponse.GetError"></a>
### func \(\*ExecuteResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L104>)

<a name="ExecuteResponse.GetGeneratedKeys"></a>
### func \(\*ExecuteResponse\) [GetGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L76>)

<a name="ExecuteResponse.GetStats"></a>
### func \(\*ExecuteResponse\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L90>)

<a name="ExecuteResponse.HasAffectedRows"></a>
### func \(\*ExecuteResponse\) [HasAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L152>)

<a name="ExecuteResponse.HasError"></a>
### func \(\*ExecuteResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L166>)

<a name="ExecuteResponse.HasStats"></a>
### func \(\*ExecuteResponse\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L159>)

<a name="ExecuteResponse.ProtoMessage"></a>
### func \(\*ExecuteResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L55>)

<a name="ExecuteResponse.ProtoReflect"></a>
### func \(\*ExecuteResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L57>)

<a name="ExecuteResponse.Reset"></a>
### func \(\*ExecuteResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L44>)

<a name="ExecuteResponse.SetAffectedRows"></a>
### func \(\*ExecuteResponse\) [SetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L118>)

<a name="ExecuteResponse.SetError"></a>
### func \(\*ExecuteResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L143>)

<a name="ExecuteResponse.SetGeneratedKeys"></a>
### func \(\*ExecuteResponse\) [SetGeneratedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L123>)


### func \(\*ExecuteResponse\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L134>)

<a name="ExecuteResponse.String"></a>
### func \(\*ExecuteResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L51>)


## type [ExecuteResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L188-L199>)

    // contains filtered or unexported fields
### func \(ExecuteResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_response.pb.go#L201>)
<a name="ExecuteStats"></a>
## type [ExecuteStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L29-L40>)

\* ExecuteStats provides execution statistics for database operations. Used for performance monitoring and operation optimization.

    // contains filtered or unexported fields
### func \(\*ExecuteStats\) [ClearAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L140>)

<a name="ExecuteStats.ClearCostEstimate"></a>
### func \(\*ExecuteStats\) [ClearCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L145>)

<a name="ExecuteStats.ClearExecutionTime"></a>
### func \(\*ExecuteStats\) [ClearExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L135>)

<a name="ExecuteStats.GetAffectedRows"></a>
### func \(\*ExecuteStats\) [GetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L81>)

<a name="ExecuteStats.GetCostEstimate"></a>
### func \(\*ExecuteStats\) [GetCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L88>)

<a name="ExecuteStats.GetExecutionTime"></a>
### func \(\*ExecuteStats\) [GetExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L67>)

<a name="ExecuteStats.HasAffectedRows"></a>
### func \(\*ExecuteStats\) [HasAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L121>)

<a name="ExecuteStats.HasCostEstimate"></a>
### func \(\*ExecuteStats\) [HasCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L128>)

<a name="ExecuteStats.HasExecutionTime"></a>
### func \(\*ExecuteStats\) [HasExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L114>)

<a name="ExecuteStats.ProtoMessage"></a>
### func \(\*ExecuteStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L53>)

<a name="ExecuteStats.ProtoReflect"></a>
### func \(\*ExecuteStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L55>)

<a name="ExecuteStats.Reset"></a>
### func \(\*ExecuteStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L42>)

<a name="ExecuteStats.SetAffectedRows"></a>
### func \(\*ExecuteStats\) [SetAffectedRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L104>)

<a name="ExecuteStats.SetCostEstimate"></a>
### func \(\*ExecuteStats\) [SetCostEstimate](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L109>)

<a name="ExecuteStats.SetExecutionTime"></a>
### func \(\*ExecuteStats\) [SetExecutionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L95>)

<a name="ExecuteStats.String"></a>
### func \(\*ExecuteStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L49>)


## type [ExecuteStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L150-L159>)

    // contains filtered or unexported fields
### func \(ExecuteStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/execute_stats.pb.go#L161>)

<a name="ExistsRequest"></a>
## type [ExistsRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*ExistsRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L140>)

<a name="ExistsRequest.ClearMetadata"></a>
### func \(\*ExistsRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L150>)

<a name="ExistsRequest.ClearNamespace"></a>
### func \(\*ExistsRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L145>)

<a name="ExistsRequest.GetKey"></a>
### func \(\*ExistsRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L66>)

<a name="ExistsRequest.GetMetadata"></a>
### func \(\*ExistsRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L86>)

<a name="ExistsRequest.GetNamespace"></a>
### func \(\*ExistsRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L76>)

<a name="ExistsRequest.HasKey"></a>
### func \(\*ExistsRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L119>)

<a name="ExistsRequest.HasMetadata"></a>
### func \(\*ExistsRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L133>)

<a name="ExistsRequest.HasNamespace"></a>
### func \(\*ExistsRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L126>)

<a name="ExistsRequest.ProtoMessage"></a>
### func \(\*ExistsRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L52>)

<a name="ExistsRequest.ProtoReflect"></a>
### func \(\*ExistsRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L54>)

<a name="ExistsRequest.Reset"></a>
### func \(\*ExistsRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L41>)

<a name="ExistsRequest.SetKey"></a>
### func \(\*ExistsRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L100>)

<a name="ExistsRequest.SetMetadata"></a>
### func \(\*ExistsRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L110>)

<a name="ExistsRequest.SetNamespace"></a>
### func \(\*ExistsRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L105>)

<a name="ExistsRequest.String"></a>
### func \(\*ExistsRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L48>)


## type [ExistsRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L155-L164>)

    // contains filtered or unexported fields
### func \(ExistsRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_request.pb.go#L166>)
<a name="ExistsResponse"></a>
## type [ExistsResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L28-L36>)

\* Response for cache key existence checks. Indicates whether a key exists in the cache.
    // contains filtered or unexported fields
### func \(\*ExistsResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L105>)

<a name="ExistsResponse.ClearExists"></a>
### func \(\*ExistsResponse\) [ClearExists](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L100>)

<a name="ExistsResponse.GetError"></a>
### func \(\*ExistsResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L70>)

<a name="ExistsResponse.GetExists"></a>
### func \(\*ExistsResponse\) [GetExists](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L63>)

<a name="ExistsResponse.HasError"></a>
### func \(\*ExistsResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L93>)

<a name="ExistsResponse.HasExists"></a>
### func \(\*ExistsResponse\) [HasExists](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L86>)

<a name="ExistsResponse.ProtoMessage"></a>
### func \(\*ExistsResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L49>)

<a name="ExistsResponse.ProtoReflect"></a>
### func \(\*ExistsResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L51>)

<a name="ExistsResponse.Reset"></a>
### func \(\*ExistsResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L38>)

<a name="ExistsResponse.SetError"></a>
### func \(\*ExistsResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L82>)

<a name="ExistsResponse.SetExists"></a>
### func \(\*ExistsResponse\) [SetExists](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L77>)

<a name="ExistsResponse.String"></a>
### func \(\*ExistsResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L45>)


## type [ExistsResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L109-L116>)

    // contains filtered or unexported fields
### func \(ExistsResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/exists_response.pb.go#L118>)

<a name="ExpireRequest"></a>
## type [ExpireRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L29-L41>)

    // contains filtered or unexported fields
### func \(\*ExpireRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L172>)

<a name="ExpireRequest.ClearMetadata"></a>
### func \(\*ExpireRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L187>)

<a name="ExpireRequest.ClearNamespace"></a>
### func \(\*ExpireRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L182>)

<a name="ExpireRequest.ClearTtl"></a>
### func \(\*ExpireRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L177>)

<a name="ExpireRequest.GetKey"></a>
### func \(\*ExpireRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L68>)

<a name="ExpireRequest.GetMetadata"></a>
### func \(\*ExpireRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L102>)

<a name="ExpireRequest.GetNamespace"></a>
### func \(\*ExpireRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L92>)

<a name="ExpireRequest.GetTtl"></a>
### func \(\*ExpireRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L78>)

<a name="ExpireRequest.HasKey"></a>
### func \(\*ExpireRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L144>)

<a name="ExpireRequest.HasMetadata"></a>
### func \(\*ExpireRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L165>)

<a name="ExpireRequest.HasNamespace"></a>
### func \(\*ExpireRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L158>)

<a name="ExpireRequest.HasTtl"></a>
### func \(\*ExpireRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L151>)

<a name="ExpireRequest.ProtoMessage"></a>
### func \(\*ExpireRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L54>)

<a name="ExpireRequest.ProtoReflect"></a>
### func \(\*ExpireRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L56>)

<a name="ExpireRequest.Reset"></a>
### func \(\*ExpireRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L43>)

<a name="ExpireRequest.SetKey"></a>
### func \(\*ExpireRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L116>)

<a name="ExpireRequest.SetMetadata"></a>
### func \(\*ExpireRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L135>)

<a name="ExpireRequest.SetNamespace"></a>
### func \(\*ExpireRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L130>)

<a name="ExpireRequest.SetTtl"></a>
### func \(\*ExpireRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L121>)

<a name="ExpireRequest.String"></a>
### func \(\*ExpireRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L50>)


## type [ExpireRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L192-L203>)

    // contains filtered or unexported fields
### func \(ExpireRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/expire_request.pb.go#L205>)

<a name="ExportRequest"></a>
## type [ExportRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*ExportRequest\) [ClearDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L140>)

<a name="ExportRequest.ClearMetadata"></a>
### func \(\*ExportRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L150>)

<a name="ExportRequest.ClearNamespace"></a>
### func \(\*ExportRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L145>)

<a name="ExportRequest.GetDestination"></a>
### func \(\*ExportRequest\) [GetDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L66>)

<a name="ExportRequest.GetMetadata"></a>
### func \(\*ExportRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L86>)

<a name="ExportRequest.GetNamespace"></a>
### func \(\*ExportRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L76>)

<a name="ExportRequest.HasDestination"></a>
### func \(\*ExportRequest\) [HasDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L119>)

<a name="ExportRequest.HasMetadata"></a>
### func \(\*ExportRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L133>)

<a name="ExportRequest.HasNamespace"></a>
### func \(\*ExportRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L126>)

<a name="ExportRequest.ProtoMessage"></a>
### func \(\*ExportRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L52>)

<a name="ExportRequest.ProtoReflect"></a>
### func \(\*ExportRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L54>)

<a name="ExportRequest.Reset"></a>
### func \(\*ExportRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L41>)

<a name="ExportRequest.SetDestination"></a>
### func \(\*ExportRequest\) [SetDestination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L100>)

<a name="ExportRequest.SetMetadata"></a>
### func \(\*ExportRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L110>)

<a name="ExportRequest.SetNamespace"></a>
### func \(\*ExportRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L105>)

<a name="ExportRequest.String"></a>
### func \(\*ExportRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L48>)


## type [ExportRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L155-L164>)

    // contains filtered or unexported fields
### func \(ExportRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/export_request.pb.go#L166>)

<a name="FlushRequest"></a>
## type [FlushRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*FlushRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L122>)

<a name="FlushRequest.ClearNamespace"></a>
### func \(\*FlushRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L117>)

<a name="FlushRequest.GetMetadata"></a>
### func \(\*FlushRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L75>)

<a name="FlushRequest.GetNamespace"></a>
### func \(\*FlushRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L65>)

<a name="FlushRequest.HasMetadata"></a>
### func \(\*FlushRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L110>)

<a name="FlushRequest.HasNamespace"></a>
### func \(\*FlushRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L103>)

<a name="FlushRequest.ProtoMessage"></a>
### func \(\*FlushRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L51>)

<a name="FlushRequest.ProtoReflect"></a>
### func \(\*FlushRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L53>)

<a name="FlushRequest.Reset"></a>
### func \(\*FlushRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L40>)

<a name="FlushRequest.SetMetadata"></a>
### func \(\*FlushRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L94>)

<a name="FlushRequest.SetNamespace"></a>
### func \(\*FlushRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L89>)

<a name="FlushRequest.String"></a>
### func \(\*FlushRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L47>)


## type [FlushRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(FlushRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_request.pb.go#L136>)
<a name="FlushResponse"></a>
## type [FlushResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L29-L38>)

\* Response for cache flush operations. Indicates success/failure of cache flushing.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*FlushResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L131>)

<a name="FlushResponse.ClearFlushedCount"></a>
### func \(\*FlushResponse\) [ClearFlushedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L121>)

<a name="FlushResponse.ClearSuccess"></a>
### func \(\*FlushResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L126>)

<a name="FlushResponse.GetError"></a>
### func \(\*FlushResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L79>)

<a name="FlushResponse.GetFlushedCount"></a>
### func \(\*FlushResponse\) [GetFlushedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L65>)

<a name="FlushResponse.GetSuccess"></a>
### func \(\*FlushResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L72>)

<a name="FlushResponse.HasError"></a>
### func \(\*FlushResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L114>)

<a name="FlushResponse.HasFlushedCount"></a>
### func \(\*FlushResponse\) [HasFlushedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L100>)

<a name="FlushResponse.HasSuccess"></a>
### func \(\*FlushResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L107>)

<a name="FlushResponse.ProtoMessage"></a>
### func \(\*FlushResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L51>)

<a name="FlushResponse.ProtoReflect"></a>
### func \(\*FlushResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L53>)

<a name="FlushResponse.Reset"></a>
### func \(\*FlushResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L40>)

<a name="FlushResponse.SetError"></a>
### func \(\*FlushResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L96>)

<a name="FlushResponse.SetFlushedCount"></a>
### func \(\*FlushResponse\) [SetFlushedCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L86>)

<a name="FlushResponse.SetSuccess"></a>
### func \(\*FlushResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L91>)

<a name="FlushResponse.String"></a>
### func \(\*FlushResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L47>)


## type [FlushResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L135-L144>)

    // contains filtered or unexported fields
### func \(FlushResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/flush_response.pb.go#L146>)

<a name="GcRequest"></a>
## type [GcRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*GcRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L122>)

<a name="GcRequest.ClearNamespace"></a>
### func \(\*GcRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L117>)

<a name="GcRequest.GetMetadata"></a>
### func \(\*GcRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L75>)

<a name="GcRequest.GetNamespace"></a>
### func \(\*GcRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L65>)

<a name="GcRequest.HasMetadata"></a>
### func \(\*GcRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L110>)

<a name="GcRequest.HasNamespace"></a>
### func \(\*GcRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L103>)

<a name="GcRequest.ProtoMessage"></a>
### func \(\*GcRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L51>)

<a name="GcRequest.ProtoReflect"></a>
### func \(\*GcRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L53>)

<a name="GcRequest.Reset"></a>
### func \(\*GcRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L40>)

<a name="GcRequest.SetMetadata"></a>
### func \(\*GcRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L94>)

<a name="GcRequest.SetNamespace"></a>
### func \(\*GcRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L89>)

<a name="GcRequest.String"></a>
### func \(\*GcRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L47>)


## type [GcRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(GcRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/gc_request.pb.go#L136>)


## type [GetConnectionInfoRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L25-L34>)


    // contains filtered or unexported fields
### func \(\*GetConnectionInfoRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L91>)

<a name="GetConnectionInfoRequest.GetMetadata"></a>
### func \(\*GetConnectionInfoRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L61>)

<a name="GetConnectionInfoRequest.HasMetadata"></a>
### func \(\*GetConnectionInfoRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L84>)

<a name="GetConnectionInfoRequest.ProtoMessage"></a>
### func \(\*GetConnectionInfoRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L47>)

<a name="GetConnectionInfoRequest.ProtoReflect"></a>
### func \(\*GetConnectionInfoRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L49>)

<a name="GetConnectionInfoRequest.Reset"></a>
### func \(\*GetConnectionInfoRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L36>)

<a name="GetConnectionInfoRequest.SetMetadata"></a>
### func \(\*GetConnectionInfoRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L75>)

<a name="GetConnectionInfoRequest.String"></a>
### func \(\*GetConnectionInfoRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L43>)


## type [GetConnectionInfoRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L96-L101>)

    // contains filtered or unexported fields
### func \(GetConnectionInfoRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_request.pb.go#L103>)
<a name="GetConnectionInfoResponse"></a>
## type [GetConnectionInfoResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L27-L37>)

\* GetConnectionInfoResponse contains database connection and pool information. Provides details about connection health and database capabilities.

    // contains filtered or unexported fields
### func \(\*GetConnectionInfoResponse\) [ClearDatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L129>)

<a name="GetConnectionInfoResponse.ClearPoolInfo"></a>
### func \(\*GetConnectionInfoResponse\) [ClearPoolInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L124>)

<a name="GetConnectionInfoResponse.GetDatabaseInfo"></a>
### func \(\*GetConnectionInfoResponse\) [GetDatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L78>)

<a name="GetConnectionInfoResponse.GetPoolInfo"></a>
### func \(\*GetConnectionInfoResponse\) [GetPoolInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L64>)

<a name="GetConnectionInfoResponse.HasDatabaseInfo"></a>
### func \(\*GetConnectionInfoResponse\) [HasDatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L117>)

<a name="GetConnectionInfoResponse.HasPoolInfo"></a>
### func \(\*GetConnectionInfoResponse\) [HasPoolInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L110>)

<a name="GetConnectionInfoResponse.ProtoMessage"></a>
### func \(\*GetConnectionInfoResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L50>)

<a name="GetConnectionInfoResponse.ProtoReflect"></a>
### func \(\*GetConnectionInfoResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L52>)

<a name="GetConnectionInfoResponse.Reset"></a>
### func \(\*GetConnectionInfoResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L39>)

<a name="GetConnectionInfoResponse.SetDatabaseInfo"></a>
### func \(\*GetConnectionInfoResponse\) [SetDatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L101>)

<a name="GetConnectionInfoResponse.SetPoolInfo"></a>
### func \(\*GetConnectionInfoResponse\) [SetPoolInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L92>)

<a name="GetConnectionInfoResponse.String"></a>
### func \(\*GetConnectionInfoResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L46>)


## type [GetConnectionInfoResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L134-L141>)

    // contains filtered or unexported fields
### func \(GetConnectionInfoResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_connection_info_response.pb.go#L143>)


## type [GetDatabaseInfoRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L26-L36>)


    // contains filtered or unexported fields
### func \(\*GetDatabaseInfoRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L120>)

<a name="GetDatabaseInfoRequest.ClearName"></a>
### func \(\*GetDatabaseInfoRequest\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L115>)

<a name="GetDatabaseInfoRequest.GetMetadata"></a>
### func \(\*GetDatabaseInfoRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L73>)

<a name="GetDatabaseInfoRequest.GetName"></a>
### func \(\*GetDatabaseInfoRequest\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L63>)

<a name="GetDatabaseInfoRequest.HasMetadata"></a>
### func \(\*GetDatabaseInfoRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L108>)

<a name="GetDatabaseInfoRequest.HasName"></a>
### func \(\*GetDatabaseInfoRequest\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L101>)

<a name="GetDatabaseInfoRequest.ProtoMessage"></a>
### func \(\*GetDatabaseInfoRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L49>)

<a name="GetDatabaseInfoRequest.ProtoReflect"></a>
### func \(\*GetDatabaseInfoRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L51>)

<a name="GetDatabaseInfoRequest.Reset"></a>
### func \(\*GetDatabaseInfoRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L38>)

<a name="GetDatabaseInfoRequest.SetMetadata"></a>
### func \(\*GetDatabaseInfoRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L92>)

<a name="GetDatabaseInfoRequest.SetName"></a>
### func \(\*GetDatabaseInfoRequest\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L87>)

<a name="GetDatabaseInfoRequest.String"></a>
### func \(\*GetDatabaseInfoRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L45>)


## type [GetDatabaseInfoRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L125-L132>)

    // contains filtered or unexported fields
### func \(GetDatabaseInfoRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_request.pb.go#L134>)
<a name="GetDatabaseInfoResponse"></a>
## type [GetDatabaseInfoResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L27-L36>)

\* GetDatabaseInfoResponse contains detailed metadata about a database. Includes version, type, capabilities, and connection information.

    // contains filtered or unexported fields
### func \(\*GetDatabaseInfoResponse\) [ClearInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L93>)

<a name="GetDatabaseInfoResponse.GetInfo"></a>
### func \(\*GetDatabaseInfoResponse\) [GetInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L63>)

<a name="GetDatabaseInfoResponse.HasInfo"></a>
### func \(\*GetDatabaseInfoResponse\) [HasInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L86>)

<a name="GetDatabaseInfoResponse.ProtoMessage"></a>
### func \(\*GetDatabaseInfoResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L49>)

<a name="GetDatabaseInfoResponse.ProtoReflect"></a>
### func \(\*GetDatabaseInfoResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L51>)

<a name="GetDatabaseInfoResponse.Reset"></a>
### func \(\*GetDatabaseInfoResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L38>)

<a name="GetDatabaseInfoResponse.SetInfo"></a>
### func \(\*GetDatabaseInfoResponse\) [SetInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L77>)

<a name="GetDatabaseInfoResponse.String"></a>
### func \(\*GetDatabaseInfoResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L45>)


## type [GetDatabaseInfoResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L98-L103>)

    // contains filtered or unexported fields
### func \(GetDatabaseInfoResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_database_info_response.pb.go#L105>)

<a name="GetMemoryUsageRequest"></a>
## type [GetMemoryUsageRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*GetMemoryUsageRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L122>)

<a name="GetMemoryUsageRequest.ClearNamespace"></a>
### func \(\*GetMemoryUsageRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L117>)

<a name="GetMemoryUsageRequest.GetMetadata"></a>
### func \(\*GetMemoryUsageRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L75>)

<a name="GetMemoryUsageRequest.GetNamespace"></a>
### func \(\*GetMemoryUsageRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L65>)

<a name="GetMemoryUsageRequest.HasMetadata"></a>
### func \(\*GetMemoryUsageRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L110>)

<a name="GetMemoryUsageRequest.HasNamespace"></a>
### func \(\*GetMemoryUsageRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L103>)

<a name="GetMemoryUsageRequest.ProtoMessage"></a>
### func \(\*GetMemoryUsageRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L51>)

<a name="GetMemoryUsageRequest.ProtoReflect"></a>
### func \(\*GetMemoryUsageRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L53>)

<a name="GetMemoryUsageRequest.Reset"></a>
### func \(\*GetMemoryUsageRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L40>)

<a name="GetMemoryUsageRequest.SetMetadata"></a>
### func \(\*GetMemoryUsageRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L94>)

<a name="GetMemoryUsageRequest.SetNamespace"></a>
### func \(\*GetMemoryUsageRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L89>)

<a name="GetMemoryUsageRequest.String"></a>
### func \(\*GetMemoryUsageRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L47>)


## type [GetMemoryUsageRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(GetMemoryUsageRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_request.pb.go#L136>)

<a name="GetMemoryUsageResponse"></a>
## type [GetMemoryUsageResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L28-L37>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*GetMemoryUsageResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L130>)

<a name="GetMemoryUsageResponse.ClearMemoryUsageBytes"></a>
### func \(\*GetMemoryUsageResponse\) [ClearMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L120>)

<a name="GetMemoryUsageResponse.ClearMemoryUsagePercent"></a>
### func \(\*GetMemoryUsageResponse\) [ClearMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L125>)

<a name="GetMemoryUsageResponse.GetError"></a>
### func \(\*GetMemoryUsageResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L78>)

<a name="GetMemoryUsageResponse.GetMemoryUsageBytes"></a>
### func \(\*GetMemoryUsageResponse\) [GetMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L64>)

<a name="GetMemoryUsageResponse.GetMemoryUsagePercent"></a>
### func \(\*GetMemoryUsageResponse\) [GetMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L71>)

<a name="GetMemoryUsageResponse.HasError"></a>
### func \(\*GetMemoryUsageResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L113>)

<a name="GetMemoryUsageResponse.HasMemoryUsageBytes"></a>
### func \(\*GetMemoryUsageResponse\) [HasMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L99>)

<a name="GetMemoryUsageResponse.HasMemoryUsagePercent"></a>
### func \(\*GetMemoryUsageResponse\) [HasMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L106>)

<a name="GetMemoryUsageResponse.ProtoMessage"></a>
### func \(\*GetMemoryUsageResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L50>)

<a name="GetMemoryUsageResponse.ProtoReflect"></a>
### func \(\*GetMemoryUsageResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L52>)

<a name="GetMemoryUsageResponse.Reset"></a>
### func \(\*GetMemoryUsageResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L39>)

<a name="GetMemoryUsageResponse.SetError"></a>
### func \(\*GetMemoryUsageResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L95>)

<a name="GetMemoryUsageResponse.SetMemoryUsageBytes"></a>
### func \(\*GetMemoryUsageResponse\) [SetMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L85>)

<a name="GetMemoryUsageResponse.SetMemoryUsagePercent"></a>
### func \(\*GetMemoryUsageResponse\) [SetMemoryUsagePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L90>)

<a name="GetMemoryUsageResponse.String"></a>
### func \(\*GetMemoryUsageResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L46>)


## type [GetMemoryUsageResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L134-L143>)

    // contains filtered or unexported fields
### func \(GetMemoryUsageResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_memory_usage_response.pb.go#L145>)


## type [GetMigrationStatusRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L26-L36>)


    // contains filtered or unexported fields
### func \(\*GetMigrationStatusRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L115>)

<a name="GetMigrationStatusRequest.ClearMetadata"></a>
### func \(\*GetMigrationStatusRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L120>)

<a name="GetMigrationStatusRequest.GetDatabase"></a>
### func \(\*GetMigrationStatusRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L63>)

<a name="GetMigrationStatusRequest.GetMetadata"></a>
### func \(\*GetMigrationStatusRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L73>)

<a name="GetMigrationStatusRequest.HasDatabase"></a>
### func \(\*GetMigrationStatusRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L101>)

<a name="GetMigrationStatusRequest.HasMetadata"></a>
### func \(\*GetMigrationStatusRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L108>)

<a name="GetMigrationStatusRequest.ProtoMessage"></a>
### func \(\*GetMigrationStatusRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L49>)

<a name="GetMigrationStatusRequest.ProtoReflect"></a>
### func \(\*GetMigrationStatusRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L51>)

<a name="GetMigrationStatusRequest.Reset"></a>
### func \(\*GetMigrationStatusRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L38>)

<a name="GetMigrationStatusRequest.SetDatabase"></a>
### func \(\*GetMigrationStatusRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L87>)

<a name="GetMigrationStatusRequest.SetMetadata"></a>
### func \(\*GetMigrationStatusRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L92>)

<a name="GetMigrationStatusRequest.String"></a>
### func \(\*GetMigrationStatusRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L45>)


## type [GetMigrationStatusRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L125-L132>)

    // contains filtered or unexported fields
### func \(GetMigrationStatusRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_request.pb.go#L134>)
<a name="GetMigrationStatusResponse"></a>
## type [GetMigrationStatusResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L28-L37>)

\* GetMigrationStatusResponse contains the current migration status for a database. Shows current version, applied migrations, and pending migrations.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*GetMigrationStatusResponse\) [ClearCurrentVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L108>)

<a name="GetMigrationStatusResponse.GetAppliedVersions"></a>
### func \(\*GetMigrationStatusResponse\) [GetAppliedVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L74>)

<a name="GetMigrationStatusResponse.GetCurrentVersion"></a>
### func \(\*GetMigrationStatusResponse\) [GetCurrentVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L64>)

<a name="GetMigrationStatusResponse.GetPendingVersions"></a>
### func \(\*GetMigrationStatusResponse\) [GetPendingVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L81>)

<a name="GetMigrationStatusResponse.HasCurrentVersion"></a>
### func \(\*GetMigrationStatusResponse\) [HasCurrentVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L101>)

<a name="GetMigrationStatusResponse.ProtoMessage"></a>
### func \(\*GetMigrationStatusResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L50>)

<a name="GetMigrationStatusResponse.ProtoReflect"></a>
### func \(\*GetMigrationStatusResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L52>)

<a name="GetMigrationStatusResponse.Reset"></a>
### func \(\*GetMigrationStatusResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L39>)

<a name="GetMigrationStatusResponse.SetAppliedVersions"></a>
### func \(\*GetMigrationStatusResponse\) [SetAppliedVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L93>)

<a name="GetMigrationStatusResponse.SetCurrentVersion"></a>
### func \(\*GetMigrationStatusResponse\) [SetCurrentVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L88>)

<a name="GetMigrationStatusResponse.SetPendingVersions"></a>
### func \(\*GetMigrationStatusResponse\) [SetPendingVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L97>)

<a name="GetMigrationStatusResponse.String"></a>
### func \(\*GetMigrationStatusResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L46>)


## type [GetMigrationStatusResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L113-L122>)

    // contains filtered or unexported fields
### func \(GetMigrationStatusResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_migration_status_response.pb.go#L124>)
<a name="GetMultipleRequest"></a>
## type [GetMultipleRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L29-L39>)

\* Request to get multiple cache values by keys. Supports batch retrieval for performance optimization.

    // contains filtered or unexported fields
### func \(\*GetMultipleRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L107>)

<a name="GetMultipleRequest.GetKeys"></a>
### func \(\*GetMultipleRequest\) [GetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L66>)

<a name="GetMultipleRequest.GetMetadata"></a>
### func \(\*GetMultipleRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L73>)

<a name="GetMultipleRequest.HasMetadata"></a>
### func \(\*GetMultipleRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L100>)

<a name="GetMultipleRequest.ProtoMessage"></a>
### func \(\*GetMultipleRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L52>)

<a name="GetMultipleRequest.ProtoReflect"></a>
### func \(\*GetMultipleRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L54>)

<a name="GetMultipleRequest.Reset"></a>
### func \(\*GetMultipleRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L41>)

<a name="GetMultipleRequest.SetKeys"></a>
### func \(\*GetMultipleRequest\) [SetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L87>)

<a name="GetMultipleRequest.SetMetadata"></a>
### func \(\*GetMultipleRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L91>)

<a name="GetMultipleRequest.String"></a>
### func \(\*GetMultipleRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L48>)


## type [GetMultipleRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L112-L119>)

    // contains filtered or unexported fields
### func \(GetMultipleRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_request.pb.go#L121>)
<a name="GetMultipleResponse"></a>
## type [GetMultipleResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L29-L36>)

\* Response for multiple cache value retrieval. Contains a map of keys to their values or error information.
    // contains filtered or unexported fields
### func \(\*GetMultipleResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L103>)

<a name="GetMultipleResponse.GetError"></a>
### func \(\*GetMultipleResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L77>)

<a name="GetMultipleResponse.GetMissingKeys"></a>
### func \(\*GetMultipleResponse\) [GetMissingKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L70>)

<a name="GetMultipleResponse.GetValues"></a>
### func \(\*GetMultipleResponse\) [GetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L63>)

<a name="GetMultipleResponse.HasError"></a>
### func \(\*GetMultipleResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L96>)

<a name="GetMultipleResponse.ProtoMessage"></a>
### func \(\*GetMultipleResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L49>)

<a name="GetMultipleResponse.ProtoReflect"></a>
### func \(\*GetMultipleResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L51>)

<a name="GetMultipleResponse.Reset"></a>
### func \(\*GetMultipleResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L38>)

<a name="GetMultipleResponse.SetError"></a>
### func \(\*GetMultipleResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L92>)

<a name="GetMultipleResponse.SetMissingKeys"></a>
### func \(\*GetMultipleResponse\) [SetMissingKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L88>)

<a name="GetMultipleResponse.SetValues"></a>
### func \(\*GetMultipleResponse\) [SetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L84>)

<a name="GetMultipleResponse.String"></a>
### func \(\*GetMultipleResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L45>)


## type [GetMultipleResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L107-L116>)

    // contains filtered or unexported fields
### func \(GetMultipleResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_multiple_response.pb.go#L118>)

<a name="GetNamespaceStatsRequest"></a>
## type [GetNamespaceStatsRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L27-L36>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*GetNamespaceStatsRequest\) [ClearIncludeDetailedMetrics](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L128>)

<a name="GetNamespaceStatsRequest.ClearIncludeKeyDistribution"></a>
### func \(\*GetNamespaceStatsRequest\) [ClearIncludeKeyDistribution](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L133>)

<a name="GetNamespaceStatsRequest.ClearNamespaceId"></a>
### func \(\*GetNamespaceStatsRequest\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L123>)

<a name="GetNamespaceStatsRequest.GetIncludeDetailedMetrics"></a>
### func \(\*GetNamespaceStatsRequest\) [GetIncludeDetailedMetrics](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L73>)

<a name="GetNamespaceStatsRequest.GetIncludeKeyDistribution"></a>
### func \(\*GetNamespaceStatsRequest\) [GetIncludeKeyDistribution](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L80>)

<a name="GetNamespaceStatsRequest.GetNamespaceId"></a>
### func \(\*GetNamespaceStatsRequest\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L63>)

<a name="GetNamespaceStatsRequest.HasIncludeDetailedMetrics"></a>
### func \(\*GetNamespaceStatsRequest\) [HasIncludeDetailedMetrics](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L109>)

<a name="GetNamespaceStatsRequest.HasIncludeKeyDistribution"></a>
### func \(\*GetNamespaceStatsRequest\) [HasIncludeKeyDistribution](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L116>)

<a name="GetNamespaceStatsRequest.HasNamespaceId"></a>
### func \(\*GetNamespaceStatsRequest\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L102>)

<a name="GetNamespaceStatsRequest.ProtoMessage"></a>
### func \(\*GetNamespaceStatsRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L49>)

<a name="GetNamespaceStatsRequest.ProtoReflect"></a>
### func \(\*GetNamespaceStatsRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L51>)

<a name="GetNamespaceStatsRequest.Reset"></a>
### func \(\*GetNamespaceStatsRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L38>)

<a name="GetNamespaceStatsRequest.SetIncludeDetailedMetrics"></a>
### func \(\*GetNamespaceStatsRequest\) [SetIncludeDetailedMetrics](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L92>)

<a name="GetNamespaceStatsRequest.SetIncludeKeyDistribution"></a>
### func \(\*GetNamespaceStatsRequest\) [SetIncludeKeyDistribution](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L97>)

<a name="GetNamespaceStatsRequest.SetNamespaceId"></a>
### func \(\*GetNamespaceStatsRequest\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L87>)

<a name="GetNamespaceStatsRequest.String"></a>
### func \(\*GetNamespaceStatsRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L45>)


## type [GetNamespaceStatsRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L138-L147>)

    // contains filtered or unexported fields
### func \(GetNamespaceStatsRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_request.pb.go#L149>)


## type [GetNamespaceStatsResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L26-L35>)

    // contains filtered or unexported fields
### func \(\*GetNamespaceStatsResponse\) [ClearCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L129>)

<a name="GetNamespaceStatsResponse.ClearNamespaceId"></a>
### func \(\*GetNamespaceStatsResponse\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L120>)

<a name="GetNamespaceStatsResponse.ClearStats"></a>
### func \(\*GetNamespaceStatsResponse\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L125>)

<a name="GetNamespaceStatsResponse.GetCollectedAt"></a>
### func \(\*GetNamespaceStatsResponse\) [GetCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L79>)

<a name="GetNamespaceStatsResponse.GetNamespaceId"></a>
### func \(\*GetNamespaceStatsResponse\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L62>)

<a name="GetNamespaceStatsResponse.GetStats"></a>
### func \(\*GetNamespaceStatsResponse\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L72>)

<a name="GetNamespaceStatsResponse.HasCollectedAt"></a>
### func \(\*GetNamespaceStatsResponse\) [HasCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L113>)

<a name="GetNamespaceStatsResponse.HasNamespaceId"></a>
### func \(\*GetNamespaceStatsResponse\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L99>)

<a name="GetNamespaceStatsResponse.HasStats"></a>
### func \(\*GetNamespaceStatsResponse\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L106>)

<a name="GetNamespaceStatsResponse.ProtoMessage"></a>
### func \(\*GetNamespaceStatsResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L48>)

<a name="GetNamespaceStatsResponse.ProtoReflect"></a>
### func \(\*GetNamespaceStatsResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L50>)

<a name="GetNamespaceStatsResponse.Reset"></a>
### func \(\*GetNamespaceStatsResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L37>)

<a name="GetNamespaceStatsResponse.SetCollectedAt"></a>
### func \(\*GetNamespaceStatsResponse\) [SetCollectedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L95>)

<a name="GetNamespaceStatsResponse.SetNamespaceId"></a>
### func \(\*GetNamespaceStatsResponse\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L86>)

<a name="GetNamespaceStatsResponse.SetStats"></a>
### func \(\*GetNamespaceStatsResponse\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L91>)

<a name="GetNamespaceStatsResponse.String"></a>
### func \(\*GetNamespaceStatsResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L44>)


## type [GetNamespaceStatsResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L133-L143>)

    // contains filtered or unexported fields
### func \(GetNamespaceStatsResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_namespace_stats_response.pb.go#L145>)
<a name="GetRequest"></a>
## type [GetRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L30-L42>)

\* Request to retrieve a cached value by key. Supports namespace isolation and access time tracking for LRU cache policies.

    // contains filtered or unexported fields
### func \(\*GetRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L162>)

<a name="GetRequest.ClearMetadata"></a>
### func \(\*GetRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L172>)

<a name="GetRequest.ClearNamespace"></a>
### func \(\*GetRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L167>)

<a name="GetRequest.ClearUpdateAccessTime"></a>
### func \(\*GetRequest\) [ClearUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L177>)

<a name="GetRequest.GetKey"></a>
### func \(\*GetRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L69>)

<a name="GetRequest.GetMetadata"></a>
### func \(\*GetRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L89>)

<a name="GetRequest.GetNamespace"></a>
### func \(\*GetRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L79>)

<a name="GetRequest.GetUpdateAccessTime"></a>
### func \(\*GetRequest\) [GetUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L103>)

<a name="GetRequest.HasKey"></a>
### func \(\*GetRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L134>)

<a name="GetRequest.HasMetadata"></a>
### func \(\*GetRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L148>)

<a name="GetRequest.HasNamespace"></a>
### func \(\*GetRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L141>)

<a name="GetRequest.HasUpdateAccessTime"></a>
### func \(\*GetRequest\) [HasUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L155>)

<a name="GetRequest.ProtoMessage"></a>
### func \(\*GetRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L55>)

<a name="GetRequest.ProtoReflect"></a>
### func \(\*GetRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L57>)

<a name="GetRequest.Reset"></a>
### func \(\*GetRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L44>)

<a name="GetRequest.SetKey"></a>
### func \(\*GetRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L110>)

<a name="GetRequest.SetMetadata"></a>
### func \(\*GetRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L120>)

<a name="GetRequest.SetNamespace"></a>
### func \(\*GetRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L115>)

<a name="GetRequest.SetUpdateAccessTime"></a>
### func \(\*GetRequest\) [SetUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L129>)

<a name="GetRequest.String"></a>
### func \(\*GetRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L51>)


## type [GetRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L182-L193>)

    // contains filtered or unexported fields
### func \(GetRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_request.pb.go#L195>)
<a name="GetResponse"></a>
## type [GetResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L27-L38>)

\* Response containing a cached value and metadata. Includes cache hit/miss information and entry details.

    // contains filtered or unexported fields
### func \(\*GetResponse\) [ClearCacheHit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L143>)

<a name="GetResponse.ClearEntry"></a>
### func \(\*GetResponse\) [ClearEntry](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L133>)

<a name="GetResponse.ClearFound"></a>
### func \(\*GetResponse\) [ClearFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L138>)

<a name="GetResponse.GetCacheHit"></a>
### func \(\*GetResponse\) [GetCacheHit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L86>)

<a name="GetResponse.GetEntry"></a>
### func \(\*GetResponse\) [GetEntry](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L65>)

<a name="GetResponse.GetFound"></a>
### func \(\*GetResponse\) [GetFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L79>)

<a name="GetResponse.HasCacheHit"></a>
### func \(\*GetResponse\) [HasCacheHit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L126>)

<a name="GetResponse.HasEntry"></a>
### func \(\*GetResponse\) [HasEntry](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L112>)

<a name="GetResponse.HasFound"></a>
### func \(\*GetResponse\) [HasFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L119>)


### func \(\*GetResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L51>)

<a name="GetResponse.ProtoReflect"></a>
### func \(\*GetResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L53>)

<a name="GetResponse.Reset"></a>
### func \(\*GetResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L40>)

<a name="GetResponse.SetCacheHit"></a>
### func \(\*GetResponse\) [SetCacheHit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L107>)

<a name="GetResponse.SetEntry"></a>
### func \(\*GetResponse\) [SetEntry](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L93>)

<a name="GetResponse.SetFound"></a>
### func \(\*GetResponse\) [SetFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L102>)

<a name="GetResponse.String"></a>
### func \(\*GetResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L47>)


## type [GetResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L148-L157>)

    // contains filtered or unexported fields
### func \(GetResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/get_response.pb.go#L159>)

<a name="ImportRequest"></a>
## type [ImportRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L28-L39>)

    // contains filtered or unexported fields
### func \(\*ImportRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L150>)

<a name="ImportRequest.ClearNamespace"></a>
### func \(\*ImportRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L145>)

<a name="ImportRequest.ClearSource"></a>
### func \(\*ImportRequest\) [ClearSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L140>)

<a name="ImportRequest.GetMetadata"></a>
### func \(\*ImportRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L86>)

<a name="ImportRequest.GetNamespace"></a>
### func \(\*ImportRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L76>)

<a name="ImportRequest.GetSource"></a>
### func \(\*ImportRequest\) [GetSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L66>)

<a name="ImportRequest.HasMetadata"></a>
### func \(\*ImportRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L133>)

<a name="ImportRequest.HasNamespace"></a>
### func \(\*ImportRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L126>)

<a name="ImportRequest.HasSource"></a>
### func \(\*ImportRequest\) [HasSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L119>)

<a name="ImportRequest.ProtoMessage"></a>
### func \(\*ImportRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L52>)

<a name="ImportRequest.ProtoReflect"></a>
### func \(\*ImportRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L54>)

<a name="ImportRequest.Reset"></a>
### func \(\*ImportRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L41>)

<a name="ImportRequest.SetMetadata"></a>
### func \(\*ImportRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L110>)

<a name="ImportRequest.SetNamespace"></a>
### func \(\*ImportRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L105>)

<a name="ImportRequest.SetSource"></a>
### func \(\*ImportRequest\) [SetSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L100>)

<a name="ImportRequest.String"></a>
### func \(\*ImportRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L48>)


## type [ImportRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L155-L164>)

    // contains filtered or unexported fields
### func \(ImportRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/import_request.pb.go#L166>)

<a name="IncrementRequest"></a>
## type [IncrementRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L29-L43>)

    // contains filtered or unexported fields
### func \(\*IncrementRequest\) [ClearDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L217>)

<a name="IncrementRequest.ClearInitialValue"></a>
### func \(\*IncrementRequest\) [ClearInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L222>)

<a name="IncrementRequest.ClearKey"></a>
### func \(\*IncrementRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L212>)

<a name="IncrementRequest.ClearMetadata"></a>
### func \(\*IncrementRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L237>)

<a name="IncrementRequest.ClearNamespace"></a>
### func \(\*IncrementRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L232>)

<a name="IncrementRequest.ClearTtl"></a>
### func \(\*IncrementRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L227>)

<a name="IncrementRequest.GetDelta"></a>
### func \(\*IncrementRequest\) [GetDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L80>)

<a name="IncrementRequest.GetInitialValue"></a>
### func \(\*IncrementRequest\) [GetInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L87>)

<a name="IncrementRequest.GetKey"></a>
### func \(\*IncrementRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L70>)

<a name="IncrementRequest.GetMetadata"></a>
### func \(\*IncrementRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L118>)

<a name="IncrementRequest.GetNamespace"></a>
### func \(\*IncrementRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L108>)

<a name="IncrementRequest.GetTtl"></a>
### func \(\*IncrementRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L94>)

<a name="IncrementRequest.HasDelta"></a>
### func \(\*IncrementRequest\) [HasDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L177>)

<a name="IncrementRequest.HasInitialValue"></a>
### func \(\*IncrementRequest\) [HasInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L184>)

<a name="IncrementRequest.HasKey"></a>
### func \(\*IncrementRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L170>)

<a name="IncrementRequest.HasMetadata"></a>
### func \(\*IncrementRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L205>)

<a name="IncrementRequest.HasNamespace"></a>
### func \(\*IncrementRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L198>)

<a name="IncrementRequest.HasTtl"></a>
### func \(\*IncrementRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L191>)

<a name="IncrementRequest.ProtoMessage"></a>
### func \(\*IncrementRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L56>)

<a name="IncrementRequest.ProtoReflect"></a>
### func \(\*IncrementRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L58>)

<a name="IncrementRequest.Reset"></a>
### func \(\*IncrementRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L45>)

<a name="IncrementRequest.SetDelta"></a>
### func \(\*IncrementRequest\) [SetDelta](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L137>)

<a name="IncrementRequest.SetInitialValue"></a>
### func \(\*IncrementRequest\) [SetInitialValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L142>)

<a name="IncrementRequest.SetKey"></a>
### func \(\*IncrementRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L132>)

<a name="IncrementRequest.SetMetadata"></a>
### func \(\*IncrementRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L161>)

<a name="IncrementRequest.SetNamespace"></a>
### func \(\*IncrementRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L156>)

<a name="IncrementRequest.SetTtl"></a>
### func \(\*IncrementRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L147>)

<a name="IncrementRequest.String"></a>
### func \(\*IncrementRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L52>)


## type [IncrementRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L242-L257>)

    // contains filtered or unexported fields
### func \(IncrementRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_request.pb.go#L259>)
<a name="IncrementResponse"></a>
## type [IncrementResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L29-L38>)

\* Response for cache increment operations. Returns the new value after incrementing.
    // contains filtered or unexported fields
### func \(\*IncrementResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L131>)

<a name="IncrementResponse.ClearNewValue"></a>
### func \(\*IncrementResponse\) [ClearNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L121>)

<a name="IncrementResponse.ClearSuccess"></a>
### func \(\*IncrementResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L126>)

<a name="IncrementResponse.GetError"></a>
### func \(\*IncrementResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L79>)

<a name="IncrementResponse.GetNewValue"></a>
### func \(\*IncrementResponse\) [GetNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L65>)

<a name="IncrementResponse.GetSuccess"></a>
### func \(\*IncrementResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L72>)

<a name="IncrementResponse.HasError"></a>
### func \(\*IncrementResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L114>)

<a name="IncrementResponse.HasNewValue"></a>
### func \(\*IncrementResponse\) [HasNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L100>)

<a name="IncrementResponse.HasSuccess"></a>
### func \(\*IncrementResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L107>)

<a name="IncrementResponse.ProtoMessage"></a>
### func \(\*IncrementResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L51>)

<a name="IncrementResponse.ProtoReflect"></a>
### func \(\*IncrementResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L53>)

<a name="IncrementResponse.Reset"></a>
### func \(\*IncrementResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L40>)

<a name="IncrementResponse.SetError"></a>
### func \(\*IncrementResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L96>)

<a name="IncrementResponse.SetNewValue"></a>
### func \(\*IncrementResponse\) [SetNewValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L86>)

<a name="IncrementResponse.SetSuccess"></a>
### func \(\*IncrementResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L91>)

<a name="IncrementResponse.String"></a>
### func \(\*IncrementResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L47>)


## type [IncrementResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L135-L144>)

    // contains filtered or unexported fields
### func \(IncrementResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/increment_response.pb.go#L146>)

<a name="InfoRequest"></a>
## type [InfoRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L27-L36>)

    // contains filtered or unexported fields
### func \(\*InfoRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L93>)

<a name="InfoRequest.GetMetadata"></a>
### func \(\*InfoRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L63>)

<a name="InfoRequest.HasMetadata"></a>
### func \(\*InfoRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L86>)

<a name="InfoRequest.ProtoMessage"></a>
### func \(\*InfoRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L49>)

<a name="InfoRequest.ProtoReflect"></a>
### func \(\*InfoRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L51>)

<a name="InfoRequest.Reset"></a>
### func \(\*InfoRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L38>)

<a name="InfoRequest.SetMetadata"></a>
### func \(\*InfoRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L77>)

<a name="InfoRequest.String"></a>
### func \(\*InfoRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L45>)


## type [InfoRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L98-L103>)

    // contains filtered or unexported fields
### func \(InfoRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/info_request.pb.go#L105>)

<a name="KeysRequest"></a>
## type [KeysRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L28-L40>)

    // contains filtered or unexported fields
### func \(\*KeysRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L186>)

<a name="KeysRequest.ClearNamespace"></a>
### func \(\*KeysRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L176>)

<a name="KeysRequest.ClearPagination"></a>
### func \(\*KeysRequest\) [ClearPagination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L181>)

<a name="KeysRequest.ClearPattern"></a>
### func \(\*KeysRequest\) [ClearPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L171>)

<a name="KeysRequest.GetMetadata"></a>
### func \(\*KeysRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L101>)

<a name="KeysRequest.GetNamespace"></a>
### func \(\*KeysRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L77>)

<a name="KeysRequest.GetPagination"></a>
### func \(\*KeysRequest\) [GetPagination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L87>)

<a name="KeysRequest.GetPattern"></a>
### func \(\*KeysRequest\) [GetPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L67>)

<a name="KeysRequest.HasMetadata"></a>
### func \(\*KeysRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L164>)

<a name="KeysRequest.HasNamespace"></a>
### func \(\*KeysRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L150>)

<a name="KeysRequest.HasPagination"></a>
### func \(\*KeysRequest\) [HasPagination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L157>)

<a name="KeysRequest.HasPattern"></a>
### func \(\*KeysRequest\) [HasPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L143>)

<a name="KeysRequest.ProtoMessage"></a>
### func \(\*KeysRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L53>)

<a name="KeysRequest.ProtoReflect"></a>
### func \(\*KeysRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L55>)

<a name="KeysRequest.Reset"></a>
### func \(\*KeysRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L42>)

<a name="KeysRequest.SetMetadata"></a>
### func \(\*KeysRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L134>)

<a name="KeysRequest.SetNamespace"></a>
### func \(\*KeysRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L120>)

<a name="KeysRequest.SetPagination"></a>
### func \(\*KeysRequest\) [SetPagination](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L125>)

<a name="KeysRequest.SetPattern"></a>
### func \(\*KeysRequest\) [SetPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L115>)

<a name="KeysRequest.String"></a>
### func \(\*KeysRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L49>)


## type [KeysRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L191-L202>)

    // contains filtered or unexported fields
### func \(KeysRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_request.pb.go#L204>)
<a name="KeysResponse"></a>
## type [KeysResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L29-L39>)

\* Response for cache keys listing operations. Returns all matching keys from the cache.
    // contains filtered or unexported fields
### func \(\*KeysResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L143>)

<a name="KeysResponse.ClearSuccess"></a>
### func \(\*KeysResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L138>)

<a name="KeysResponse.ClearTotalCount"></a>
### func \(\*KeysResponse\) [ClearTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L133>)

<a name="KeysResponse.GetError"></a>
### func \(\*KeysResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L87>)

<a name="KeysResponse.GetKeys"></a>
### func \(\*KeysResponse\) [GetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L66>)

<a name="KeysResponse.GetSuccess"></a>
### func \(\*KeysResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L80>)

<a name="KeysResponse.GetTotalCount"></a>
### func \(\*KeysResponse\) [GetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L73>)

<a name="KeysResponse.HasError"></a>
### func \(\*KeysResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L126>)

<a name="KeysResponse.HasSuccess"></a>
### func \(\*KeysResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L119>)

<a name="KeysResponse.HasTotalCount"></a>
### func \(\*KeysResponse\) [HasTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L112>)

<a name="KeysResponse.ProtoMessage"></a>
### func \(\*KeysResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L52>)

<a name="KeysResponse.ProtoReflect"></a>
### func \(\*KeysResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L54>)

<a name="KeysResponse.Reset"></a>
### func \(\*KeysResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L41>)

<a name="KeysResponse.SetError"></a>
### func \(\*KeysResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L108>)

<a name="KeysResponse.SetKeys"></a>
### func \(\*KeysResponse\) [SetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L94>)

<a name="KeysResponse.SetSuccess"></a>
### func \(\*KeysResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L103>)

<a name="KeysResponse.SetTotalCount"></a>
### func \(\*KeysResponse\) [SetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L98>)

<a name="KeysResponse.String"></a>
### func \(\*KeysResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L48>)


## type [KeysResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L147-L158>)

    // contains filtered or unexported fields
### func \(KeysResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/keys_response.pb.go#L160>)


## type [ListDatabasesRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L25-L34>)


    // contains filtered or unexported fields
### func \(\*ListDatabasesRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L91>)

<a name="ListDatabasesRequest.GetMetadata"></a>
### func \(\*ListDatabasesRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L61>)

<a name="ListDatabasesRequest.HasMetadata"></a>
### func \(\*ListDatabasesRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L84>)

<a name="ListDatabasesRequest.ProtoMessage"></a>
### func \(\*ListDatabasesRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L47>)

<a name="ListDatabasesRequest.ProtoReflect"></a>
### func \(\*ListDatabasesRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L49>)

<a name="ListDatabasesRequest.Reset"></a>
### func \(\*ListDatabasesRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L36>)

<a name="ListDatabasesRequest.SetMetadata"></a>
### func \(\*ListDatabasesRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L75>)

<a name="ListDatabasesRequest.String"></a>
### func \(\*ListDatabasesRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L43>)


## type [ListDatabasesRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L96-L101>)

    // contains filtered or unexported fields
### func \(ListDatabasesRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_request.pb.go#L103>)
<a name="ListDatabasesResponse"></a>
## type [ListDatabasesResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L28-L33>)

\* ListDatabasesResponse contains the list of available databases. Provides database names accessible to the authenticated user.
    // contains filtered or unexported fields
### func \(\*ListDatabasesResponse\) [GetDatabases](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L60>)

<a name="ListDatabasesResponse.ProtoMessage"></a>
### func \(\*ListDatabasesResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L46>)

<a name="ListDatabasesResponse.ProtoReflect"></a>
### func \(\*ListDatabasesResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L48>)

<a name="ListDatabasesResponse.Reset"></a>
### func \(\*ListDatabasesResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L35>)

<a name="ListDatabasesResponse.SetDatabases"></a>
### func \(\*ListDatabasesResponse\) [SetDatabases](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L67>)

<a name="ListDatabasesResponse.String"></a>
### func \(\*ListDatabasesResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L42>)


## type [ListDatabasesResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L71-L76>)

    // contains filtered or unexported fields
### func \(ListDatabasesResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_databases_response.pb.go#L78>)


## type [ListMigrationsRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L26-L37>)


    // contains filtered or unexported fields
### func \(\*ListMigrationsRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L138>)

<a name="ListMigrationsRequest.ClearMetadata"></a>
### func \(\*ListMigrationsRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L148>)

<a name="ListMigrationsRequest.ClearStatusFilter"></a>
### func \(\*ListMigrationsRequest\) [ClearStatusFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L143>)

<a name="ListMigrationsRequest.GetDatabase"></a>
### func \(\*ListMigrationsRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L64>)

<a name="ListMigrationsRequest.GetMetadata"></a>
### func \(\*ListMigrationsRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L84>)

<a name="ListMigrationsRequest.GetStatusFilter"></a>
### func \(\*ListMigrationsRequest\) [GetStatusFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L74>)

<a name="ListMigrationsRequest.HasDatabase"></a>
### func \(\*ListMigrationsRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L117>)

<a name="ListMigrationsRequest.HasMetadata"></a>
### func \(\*ListMigrationsRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L131>)

<a name="ListMigrationsRequest.HasStatusFilter"></a>
### func \(\*ListMigrationsRequest\) [HasStatusFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L124>)

<a name="ListMigrationsRequest.ProtoMessage"></a>
### func \(\*ListMigrationsRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L50>)

<a name="ListMigrationsRequest.ProtoReflect"></a>
### func \(\*ListMigrationsRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L52>)

<a name="ListMigrationsRequest.Reset"></a>
### func \(\*ListMigrationsRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L39>)

<a name="ListMigrationsRequest.SetDatabase"></a>
### func \(\*ListMigrationsRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L98>)

<a name="ListMigrationsRequest.SetMetadata"></a>
### func \(\*ListMigrationsRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L108>)

<a name="ListMigrationsRequest.SetStatusFilter"></a>
### func \(\*ListMigrationsRequest\) [SetStatusFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L103>)

<a name="ListMigrationsRequest.String"></a>
### func \(\*ListMigrationsRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L46>)


## type [ListMigrationsRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L153-L162>)

    // contains filtered or unexported fields
### func \(ListMigrationsRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_request.pb.go#L164>)

<a name="ListMigrationsResponse"></a>
## type [ListMigrationsResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L27-L36>)

    // contains filtered or unexported fields
### func \(\*ListMigrationsResponse\) [GetMigrations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L63>)

<a name="ListMigrationsResponse.ProtoMessage"></a>
### func \(\*ListMigrationsResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L49>)

<a name="ListMigrationsResponse.ProtoReflect"></a>
### func \(\*ListMigrationsResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L51>)

<a name="ListMigrationsResponse.Reset"></a>
### func \(\*ListMigrationsResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L38>)

<a name="ListMigrationsResponse.SetMigrations"></a>
### func \(\*ListMigrationsResponse\) [SetMigrations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L77>)

<a name="ListMigrationsResponse.String"></a>
### func \(\*ListMigrationsResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L45>)


## type [ListMigrationsResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L88-L93>)

    // contains filtered or unexported fields
### func \(ListMigrationsResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_migrations_response.pb.go#L95>)

<a name="ListNamespacesRequest"></a>
## type [ListNamespacesRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L27-L37>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*ListNamespacesRequest\) [ClearIncludeStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L158>)

<a name="ListNamespacesRequest.ClearNameFilter"></a>
### func \(\*ListNamespacesRequest\) [ClearNameFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L153>)

<a name="ListNamespacesRequest.ClearPage"></a>
### func \(\*ListNamespacesRequest\) [ClearPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L143>)

<a name="ListNamespacesRequest.ClearPageSize"></a>
### func \(\*ListNamespacesRequest\) [ClearPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L148>)

<a name="ListNamespacesRequest.GetIncludeStats"></a>
### func \(\*ListNamespacesRequest\) [GetIncludeStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L88>)

<a name="ListNamespacesRequest.GetNameFilter"></a>
### func \(\*ListNamespacesRequest\) [GetNameFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L78>)

<a name="ListNamespacesRequest.GetPage"></a>
### func \(\*ListNamespacesRequest\) [GetPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L64>)

<a name="ListNamespacesRequest.GetPageSize"></a>
### func \(\*ListNamespacesRequest\) [GetPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L71>)

<a name="ListNamespacesRequest.HasIncludeStats"></a>
### func \(\*ListNamespacesRequest\) [HasIncludeStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L136>)

<a name="ListNamespacesRequest.HasNameFilter"></a>
### func \(\*ListNamespacesRequest\) [HasNameFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L129>)

<a name="ListNamespacesRequest.HasPage"></a>
### func \(\*ListNamespacesRequest\) [HasPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L115>)

<a name="ListNamespacesRequest.HasPageSize"></a>
### func \(\*ListNamespacesRequest\) [HasPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L122>)

<a name="ListNamespacesRequest.ProtoMessage"></a>
### func \(\*ListNamespacesRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L50>)

<a name="ListNamespacesRequest.ProtoReflect"></a>
### func \(\*ListNamespacesRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L52>)

<a name="ListNamespacesRequest.Reset"></a>
### func \(\*ListNamespacesRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L39>)

<a name="ListNamespacesRequest.SetIncludeStats"></a>
### func \(\*ListNamespacesRequest\) [SetIncludeStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L110>)

<a name="ListNamespacesRequest.SetNameFilter"></a>
### func \(\*ListNamespacesRequest\) [SetNameFilter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L105>)

<a name="ListNamespacesRequest.SetPage"></a>
### func \(\*ListNamespacesRequest\) [SetPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L95>)

<a name="ListNamespacesRequest.SetPageSize"></a>
### func \(\*ListNamespacesRequest\) [SetPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L100>)

<a name="ListNamespacesRequest.String"></a>
### func \(\*ListNamespacesRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L46>)


## type [ListNamespacesRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L163-L174>)

    // contains filtered or unexported fields
### func \(ListNamespacesRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_request.pb.go#L176>)


## type [ListNamespacesResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L25-L36>)

    // contains filtered or unexported fields
### func \(\*ListNamespacesResponse\) [ClearPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L157>)

<a name="ListNamespacesResponse.ClearPageSize"></a>
### func \(\*ListNamespacesResponse\) [ClearPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L162>)

<a name="ListNamespacesResponse.ClearTotalCount"></a>
### func \(\*ListNamespacesResponse\) [ClearTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L152>)

<a name="ListNamespacesResponse.ClearTotalPages"></a>
### func \(\*ListNamespacesResponse\) [ClearTotalPages](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L167>)

<a name="ListNamespacesResponse.GetNamespaces"></a>
### func \(\*ListNamespacesResponse\) [GetNamespaces](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L63>)

<a name="ListNamespacesResponse.GetPage"></a>
### func \(\*ListNamespacesResponse\) [GetPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L79>)

<a name="ListNamespacesResponse.GetPageSize"></a>
### func \(\*ListNamespacesResponse\) [GetPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L86>)

<a name="ListNamespacesResponse.GetTotalCount"></a>
### func \(\*ListNamespacesResponse\) [GetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L72>)

<a name="ListNamespacesResponse.GetTotalPages"></a>
### func \(\*ListNamespacesResponse\) [GetTotalPages](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L93>)

<a name="ListNamespacesResponse.HasPage"></a>
### func \(\*ListNamespacesResponse\) [HasPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L131>)

<a name="ListNamespacesResponse.HasPageSize"></a>
### func \(\*ListNamespacesResponse\) [HasPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L138>)

<a name="ListNamespacesResponse.HasTotalCount"></a>
### func \(\*ListNamespacesResponse\) [HasTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L124>)

<a name="ListNamespacesResponse.HasTotalPages"></a>
### func \(\*ListNamespacesResponse\) [HasTotalPages](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L145>)

<a name="ListNamespacesResponse.ProtoMessage"></a>
### func \(\*ListNamespacesResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L49>)

<a name="ListNamespacesResponse.ProtoReflect"></a>
### func \(\*ListNamespacesResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L51>)

<a name="ListNamespacesResponse.Reset"></a>
### func \(\*ListNamespacesResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L38>)

<a name="ListNamespacesResponse.SetNamespaces"></a>
### func \(\*ListNamespacesResponse\) [SetNamespaces](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L100>)

<a name="ListNamespacesResponse.SetPage"></a>
### func \(\*ListNamespacesResponse\) [SetPage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L109>)

<a name="ListNamespacesResponse.SetPageSize"></a>
### func \(\*ListNamespacesResponse\) [SetPageSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L114>)

<a name="ListNamespacesResponse.SetTotalCount"></a>
### func \(\*ListNamespacesResponse\) [SetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L104>)

<a name="ListNamespacesResponse.SetTotalPages"></a>
### func \(\*ListNamespacesResponse\) [SetTotalPages](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L119>)

<a name="ListNamespacesResponse.String"></a>
### func \(\*ListNamespacesResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L45>)


## type [ListNamespacesResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L172-L185>)

    // contains filtered or unexported fields
### func \(ListNamespacesResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_namespaces_response.pb.go#L187>)


## type [ListSchemasRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L26-L36>)


    // contains filtered or unexported fields
### func \(\*ListSchemasRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L115>)

<a name="ListSchemasRequest.ClearMetadata"></a>
### func \(\*ListSchemasRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L120>)

<a name="ListSchemasRequest.GetDatabase"></a>
### func \(\*ListSchemasRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L63>)

<a name="ListSchemasRequest.GetMetadata"></a>
### func \(\*ListSchemasRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L73>)

<a name="ListSchemasRequest.HasDatabase"></a>
### func \(\*ListSchemasRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L101>)

<a name="ListSchemasRequest.HasMetadata"></a>
### func \(\*ListSchemasRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L108>)

<a name="ListSchemasRequest.ProtoMessage"></a>
### func \(\*ListSchemasRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L49>)

<a name="ListSchemasRequest.ProtoReflect"></a>
### func \(\*ListSchemasRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L51>)

<a name="ListSchemasRequest.Reset"></a>
### func \(\*ListSchemasRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L38>)

<a name="ListSchemasRequest.SetDatabase"></a>
### func \(\*ListSchemasRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L87>)

<a name="ListSchemasRequest.SetMetadata"></a>
### func \(\*ListSchemasRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L92>)

<a name="ListSchemasRequest.String"></a>
### func \(\*ListSchemasRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L45>)


## type [ListSchemasRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L125-L132>)

    // contains filtered or unexported fields
### func \(ListSchemasRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_request.pb.go#L134>)
<a name="ListSchemasResponse"></a>
## type [ListSchemasResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L28-L33>)

\* ListSchemasResponse contains the list of schemas within a database. Provides schema names available in the specified database.
    // contains filtered or unexported fields
### func \(\*ListSchemasResponse\) [GetSchemas](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L60>)

<a name="ListSchemasResponse.ProtoMessage"></a>
### func \(\*ListSchemasResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L46>)

<a name="ListSchemasResponse.ProtoReflect"></a>
### func \(\*ListSchemasResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L48>)

<a name="ListSchemasResponse.Reset"></a>
### func \(\*ListSchemasResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L35>)

<a name="ListSchemasResponse.SetSchemas"></a>
### func \(\*ListSchemasResponse\) [SetSchemas](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L67>)

<a name="ListSchemasResponse.String"></a>
### func \(\*ListSchemasResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L42>)


## type [ListSchemasResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L71-L76>)

    // contains filtered or unexported fields
### func \(ListSchemasResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/list_schemas_response.pb.go#L78>)

<a name="LockRequest"></a>
## type [LockRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L29-L41>)

    // contains filtered or unexported fields
### func \(\*LockRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L172>)

<a name="LockRequest.ClearMetadata"></a>
### func \(\*LockRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L187>)

<a name="LockRequest.ClearNamespace"></a>
### func \(\*LockRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L182>)

<a name="LockRequest.ClearTtl"></a>
### func \(\*LockRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L177>)

<a name="LockRequest.GetKey"></a>
### func \(\*LockRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L68>)

<a name="LockRequest.GetMetadata"></a>
### func \(\*LockRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L102>)

<a name="LockRequest.GetNamespace"></a>
### func \(\*LockRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L92>)

<a name="LockRequest.GetTtl"></a>
### func \(\*LockRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L78>)

<a name="LockRequest.HasKey"></a>
### func \(\*LockRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L144>)

<a name="LockRequest.HasMetadata"></a>
### func \(\*LockRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L165>)

<a name="LockRequest.HasNamespace"></a>
### func \(\*LockRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L158>)

<a name="LockRequest.HasTtl"></a>
### func \(\*LockRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L151>)

<a name="LockRequest.ProtoMessage"></a>
### func \(\*LockRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L54>)

<a name="LockRequest.ProtoReflect"></a>
### func \(\*LockRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L56>)

<a name="LockRequest.Reset"></a>
### func \(\*LockRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L43>)

<a name="LockRequest.SetKey"></a>
### func \(\*LockRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L116>)

<a name="LockRequest.SetMetadata"></a>
### func \(\*LockRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L135>)

<a name="LockRequest.SetNamespace"></a>
### func \(\*LockRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L130>)

<a name="LockRequest.SetTtl"></a>
### func \(\*LockRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L121>)

<a name="LockRequest.String"></a>
### func \(\*LockRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L50>)


## type [LockRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L192-L203>)

    // contains filtered or unexported fields
### func \(LockRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/lock_request.pb.go#L205>)
<a name="MGetRequest"></a>
## type [MGetRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L29-L40>)

\* MGetRequest is used to retrieve multiple cache entries in a single operation. This is more efficient than multiple individual Get operations.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*MGetRequest\) [ClearIncludeExpired](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L161>)

<a name="MGetRequest.ClearMetadata"></a>
### func \(\*MGetRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L171>)

<a name="MGetRequest.ClearNamespace"></a>
### func \(\*MGetRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L156>)

<a name="MGetRequest.ClearUpdateAccessTime"></a>
### func \(\*MGetRequest\) [ClearUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L166>)

<a name="MGetRequest.GetIncludeExpired"></a>
### func \(\*MGetRequest\) [GetIncludeExpired](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L84>)

<a name="MGetRequest.GetKeys"></a>
### func \(\*MGetRequest\) [GetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L67>)

<a name="MGetRequest.GetMetadata"></a>
### func \(\*MGetRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L98>)

<a name="MGetRequest.GetNamespace"></a>
### func \(\*MGetRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L74>)

<a name="MGetRequest.GetUpdateAccessTime"></a>
### func \(\*MGetRequest\) [GetUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L91>)

<a name="MGetRequest.HasIncludeExpired"></a>
### func \(\*MGetRequest\) [HasIncludeExpired](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L135>)

<a name="MGetRequest.HasMetadata"></a>
### func \(\*MGetRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L149>)

<a name="MGetRequest.HasNamespace"></a>
### func \(\*MGetRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L128>)

<a name="MGetRequest.HasUpdateAccessTime"></a>
### func \(\*MGetRequest\) [HasUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L142>)

<a name="MGetRequest.ProtoMessage"></a>
### func \(\*MGetRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L53>)

<a name="MGetRequest.ProtoReflect"></a>
### func \(\*MGetRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L55>)

<a name="MGetRequest.Reset"></a>
### func \(\*MGetRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L42>)

<a name="MGetRequest.SetIncludeExpired"></a>
### func \(\*MGetRequest\) [SetIncludeExpired](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L114>)

<a name="MGetRequest.SetKeys"></a>
### func \(\*MGetRequest\) [SetKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L105>)

<a name="MGetRequest.SetMetadata"></a>
### func \(\*MGetRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L124>)

<a name="MGetRequest.SetNamespace"></a>
### func \(\*MGetRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L109>)

<a name="MGetRequest.SetUpdateAccessTime"></a>
### func \(\*MGetRequest\) [SetUpdateAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L119>)

<a name="MGetRequest.String"></a>
### func \(\*MGetRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L49>)


## type [MGetRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L175-L188>)

    // contains filtered or unexported fields
### func \(MGetRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/m_get_request.pb.go#L190>)

<a name="MigrationInfo"></a>
## type [MigrationInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L28-L37>)
    // contains filtered or unexported fields
### func \(\*MigrationInfo\) [ClearAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L136>)

<a name="MigrationInfo.ClearDescription"></a>
### func \(\*MigrationInfo\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L131>)

<a name="MigrationInfo.ClearVersion"></a>
### func \(\*MigrationInfo\) [ClearVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L126>)

<a name="MigrationInfo.GetAppliedAt"></a>
### func \(\*MigrationInfo\) [GetAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L84>)

<a name="MigrationInfo.GetDescription"></a>
### func \(\*MigrationInfo\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L74>)

<a name="MigrationInfo.GetVersion"></a>
### func \(\*MigrationInfo\) [GetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L64>)

<a name="MigrationInfo.HasAppliedAt"></a>
### func \(\*MigrationInfo\) [HasAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L119>)

<a name="MigrationInfo.HasDescription"></a>
### func \(\*MigrationInfo\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L112>)

<a name="MigrationInfo.HasVersion"></a>
### func \(\*MigrationInfo\) [HasVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L105>)

<a name="MigrationInfo.ProtoMessage"></a>
### func \(\*MigrationInfo\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L50>)

<a name="MigrationInfo.ProtoReflect"></a>
### func \(\*MigrationInfo\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L52>)

<a name="MigrationInfo.Reset"></a>
### func \(\*MigrationInfo\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L39>)

<a name="MigrationInfo.SetAppliedAt"></a>
### func \(\*MigrationInfo\) [SetAppliedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L101>)

<a name="MigrationInfo.SetDescription"></a>
### func \(\*MigrationInfo\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L96>)

<a name="MigrationInfo.SetVersion"></a>
### func \(\*MigrationInfo\) [SetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L91>)

<a name="MigrationInfo.String"></a>
### func \(\*MigrationInfo\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L46>)


## type [MigrationInfo\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L140-L149>)

    // contains filtered or unexported fields
### func \(MigrationInfo\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_info.pb.go#L151>)
<a name="MigrationScript"></a>
## type [MigrationScript](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L28-L37>)

\* MigrationScript represents a database migration script with version control. Used for managing database schema changes and data migrations.
    // contains filtered or unexported fields
### func \(\*MigrationScript\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L140>)

<a name="MigrationScript.ClearScript"></a>
### func \(\*MigrationScript\) [ClearScript](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L135>)

<a name="MigrationScript.ClearVersion"></a>
### func \(\*MigrationScript\) [ClearVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L130>)

<a name="MigrationScript.GetDescription"></a>
### func \(\*MigrationScript\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L84>)

<a name="MigrationScript.GetScript"></a>
### func \(\*MigrationScript\) [GetScript](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L74>)

<a name="MigrationScript.GetVersion"></a>
### func \(\*MigrationScript\) [GetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L64>)

<a name="MigrationScript.HasDescription"></a>
### func \(\*MigrationScript\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L123>)

<a name="MigrationScript.HasScript"></a>
### func \(\*MigrationScript\) [HasScript](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L116>)

<a name="MigrationScript.HasVersion"></a>
### func \(\*MigrationScript\) [HasVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L109>)

<a name="MigrationScript.ProtoMessage"></a>
### func \(\*MigrationScript\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L50>)

<a name="MigrationScript.ProtoReflect"></a>
### func \(\*MigrationScript\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L52>)

<a name="MigrationScript.Reset"></a>
### func \(\*MigrationScript\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L39>)

<a name="MigrationScript.SetDescription"></a>
### func \(\*MigrationScript\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L104>)

<a name="MigrationScript.SetScript"></a>
### func \(\*MigrationScript\) [SetScript](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L99>)

<a name="MigrationScript.SetVersion"></a>
### func \(\*MigrationScript\) [SetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L94>)

<a name="MigrationScript.String"></a>
### func \(\*MigrationScript\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L46>)


## type [MigrationScript\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L145-L154>)

    // contains filtered or unexported fields
### func \(MigrationScript\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_script.pb.go#L156>)

<a name="MigrationServiceClient"></a>
## type [MigrationServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L34-L43>)
For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
### func [NewMigrationServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L49>)
<a name="MigrationServiceServer"></a>
## type [MigrationServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L99-L109>)

MigrationServiceServer is the server API for MigrationService service. All implementations must embed UnimplementedMigrationServiceServer for forward compatibility.
    // contains filtered or unexported methods
## type [MySQLConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L27-L37>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*MySQLConfig\) [ClearConnectTimeoutSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L158>)

<a name="MySQLConfig.ClearDsn"></a>
### func \(\*MySQLConfig\) [ClearDsn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L143>)

<a name="MySQLConfig.ClearMaxIdleConns"></a>
### func \(\*MySQLConfig\) [ClearMaxIdleConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L153>)

<a name="MySQLConfig.ClearMaxOpenConns"></a>
### func \(\*MySQLConfig\) [ClearMaxOpenConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L148>)

<a name="MySQLConfig.GetConnectTimeoutSeconds"></a>
### func \(\*MySQLConfig\) [GetConnectTimeoutSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L88>)

<a name="MySQLConfig.GetDsn"></a>
### func \(\*MySQLConfig\) [GetDsn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L64>)

<a name="MySQLConfig.GetMaxIdleConns"></a>
### func \(\*MySQLConfig\) [GetMaxIdleConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L81>)

<a name="MySQLConfig.GetMaxOpenConns"></a>
### func \(\*MySQLConfig\) [GetMaxOpenConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L74>)

<a name="MySQLConfig.HasConnectTimeoutSeconds"></a>
### func \(\*MySQLConfig\) [HasConnectTimeoutSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L136>)

<a name="MySQLConfig.HasDsn"></a>
### func \(\*MySQLConfig\) [HasDsn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L115>)

<a name="MySQLConfig.HasMaxIdleConns"></a>
### func \(\*MySQLConfig\) [HasMaxIdleConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L129>)

<a name="MySQLConfig.HasMaxOpenConns"></a>
### func \(\*MySQLConfig\) [HasMaxOpenConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L122>)

<a name="MySQLConfig.ProtoMessage"></a>
### func \(\*MySQLConfig\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L50>)

<a name="MySQLConfig.ProtoReflect"></a>
### func \(\*MySQLConfig\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L52>)

<a name="MySQLConfig.Reset"></a>
### func \(\*MySQLConfig\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L39>)

<a name="MySQLConfig.SetConnectTimeoutSeconds"></a>
### func \(\*MySQLConfig\) [SetConnectTimeoutSeconds](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L110>)

<a name="MySQLConfig.SetDsn"></a>
### func \(\*MySQLConfig\) [SetDsn](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L95>)

<a name="MySQLConfig.SetMaxIdleConns"></a>
### func \(\*MySQLConfig\) [SetMaxIdleConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L105>)

<a name="MySQLConfig.SetMaxOpenConns"></a>
### func \(\*MySQLConfig\) [SetMaxOpenConns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L100>)

<a name="MySQLConfig.String"></a>
### func \(\*MySQLConfig\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L46>)


## type [MySQLConfig\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L163-L174>)

    // contains filtered or unexported fields
### func \(MySQLConfig\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_config.pb.go#L176>)

<a name="MySQLStatus"></a>
## type [MySQLStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L28-L38>)
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*MySQLStatus\) [ClearOpenConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L155>)

<a name="MySQLStatus.ClearRole"></a>
### func \(\*MySQLStatus\) [ClearRole](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L160>)

<a name="MySQLStatus.ClearStartedAt"></a>
### func \(\*MySQLStatus\) [ClearStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L151>)

<a name="MySQLStatus.ClearVersion"></a>
### func \(\*MySQLStatus\) [ClearVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L146>)

<a name="MySQLStatus.GetOpenConnections"></a>
### func \(\*MySQLStatus\) [GetOpenConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L82>)

<a name="MySQLStatus.GetRole"></a>
### func \(\*MySQLStatus\) [GetRole](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L89>)

<a name="MySQLStatus.GetStartedAt"></a>
### func \(\*MySQLStatus\) [GetStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L75>)

<a name="MySQLStatus.GetVersion"></a>
### func \(\*MySQLStatus\) [GetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L65>)

<a name="MySQLStatus.HasOpenConnections"></a>
### func \(\*MySQLStatus\) [HasOpenConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L132>)

<a name="MySQLStatus.HasRole"></a>
### func \(\*MySQLStatus\) [HasRole](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L139>)

<a name="MySQLStatus.HasStartedAt"></a>
### func \(\*MySQLStatus\) [HasStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L125>)

<a name="MySQLStatus.HasVersion"></a>
### func \(\*MySQLStatus\) [HasVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L118>)

<a name="MySQLStatus.ProtoMessage"></a>
### func \(\*MySQLStatus\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L51>)

<a name="MySQLStatus.ProtoReflect"></a>
### func \(\*MySQLStatus\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L53>)

<a name="MySQLStatus.Reset"></a>
### func \(\*MySQLStatus\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L40>)

<a name="MySQLStatus.SetOpenConnections"></a>
### func \(\*MySQLStatus\) [SetOpenConnections](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L108>)

<a name="MySQLStatus.SetRole"></a>
### func \(\*MySQLStatus\) [SetRole](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L113>)

<a name="MySQLStatus.SetStartedAt"></a>
### func \(\*MySQLStatus\) [SetStartedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L104>)

<a name="MySQLStatus.SetVersion"></a>
### func \(\*MySQLStatus\) [SetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L99>)

<a name="MySQLStatus.String"></a>
### func \(\*MySQLStatus\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L47>)


## type [MySQLStatus\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L165-L176>)

    // contains filtered or unexported fields
### func \(MySQLStatus\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/my_sql_status.pb.go#L178>)


## type [NamespaceInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L26-L39>)

    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*NamespaceInfo\) [ClearCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L214>)

<a name="NamespaceInfo.ClearCurrentKeys"></a>
### func \(\*NamespaceInfo\) [ClearCurrentKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L218>)

<a name="NamespaceInfo.ClearCurrentMemoryBytes"></a>
### func \(\*NamespaceInfo\) [ClearCurrentMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L223>)

<a name="NamespaceInfo.ClearDescription"></a>
### func \(\*NamespaceInfo\) [ClearDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L209>)

<a name="NamespaceInfo.ClearName"></a>
### func \(\*NamespaceInfo\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L204>)

<a name="NamespaceInfo.ClearNamespaceId"></a>
### func \(\*NamespaceInfo\) [ClearNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L199>)

<a name="NamespaceInfo.GetConfig"></a>
### func \(\*NamespaceInfo\) [GetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L117>)

<a name="NamespaceInfo.GetCreatedAt"></a>
### func \(\*NamespaceInfo\) [GetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L96>)

<a name="NamespaceInfo.GetCurrentKeys"></a>
### func \(\*NamespaceInfo\) [GetCurrentKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L103>)

<a name="NamespaceInfo.GetCurrentMemoryBytes"></a>
### func \(\*NamespaceInfo\) [GetCurrentMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L110>)

<a name="NamespaceInfo.GetDescription"></a>
### func \(\*NamespaceInfo\) [GetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L86>)

<a name="NamespaceInfo.GetName"></a>
### func \(\*NamespaceInfo\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L76>)

<a name="NamespaceInfo.GetNamespaceId"></a>
### func \(\*NamespaceInfo\) [GetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L66>)

<a name="NamespaceInfo.HasCreatedAt"></a>
### func \(\*NamespaceInfo\) [HasCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L178>)

<a name="NamespaceInfo.HasCurrentKeys"></a>
### func \(\*NamespaceInfo\) [HasCurrentKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L185>)

<a name="NamespaceInfo.HasCurrentMemoryBytes"></a>
### func \(\*NamespaceInfo\) [HasCurrentMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L192>)

<a name="NamespaceInfo.HasDescription"></a>
### func \(\*NamespaceInfo\) [HasDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L171>)

<a name="NamespaceInfo.HasName"></a>
### func \(\*NamespaceInfo\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L164>)

<a name="NamespaceInfo.HasNamespaceId"></a>
### func \(\*NamespaceInfo\) [HasNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L157>)

<a name="NamespaceInfo.ProtoMessage"></a>
### func \(\*NamespaceInfo\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L52>)

<a name="NamespaceInfo.ProtoReflect"></a>
### func \(\*NamespaceInfo\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L54>)

<a name="NamespaceInfo.Reset"></a>
### func \(\*NamespaceInfo\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L41>)

<a name="NamespaceInfo.SetConfig"></a>
### func \(\*NamespaceInfo\) [SetConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L153>)

<a name="NamespaceInfo.SetCreatedAt"></a>
### func \(\*NamespaceInfo\) [SetCreatedAt](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L139>)

<a name="NamespaceInfo.SetCurrentKeys"></a>
### func \(\*NamespaceInfo\) [SetCurrentKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L143>)

<a name="NamespaceInfo.SetCurrentMemoryBytes"></a>
### func \(\*NamespaceInfo\) [SetCurrentMemoryBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L148>)

<a name="NamespaceInfo.SetDescription"></a>
### func \(\*NamespaceInfo\) [SetDescription](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L134>)

<a name="NamespaceInfo.SetName"></a>
### func \(\*NamespaceInfo\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L129>)

<a name="NamespaceInfo.SetNamespaceId"></a>
### func \(\*NamespaceInfo\) [SetNamespaceId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L124>)

<a name="NamespaceInfo.String"></a>
### func \(\*NamespaceInfo\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L48>)


## type [NamespaceInfo\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L228-L245>)

    // contains filtered or unexported fields
### func \(NamespaceInfo\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_info.pb.go#L247>)


## type [NamespaceStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L26-L41>)

    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*NamespaceStats\) [ClearAvgKeySizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L268>)

<a name="NamespaceStats.ClearAvgValueSizeBytes"></a>
### func \(\*NamespaceStats\) [ClearAvgValueSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L273>)

<a name="NamespaceStats.ClearCacheHits"></a>
### func \(\*NamespaceStats\) [ClearCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L253>)

<a name="NamespaceStats.ClearCacheMisses"></a>
### func \(\*NamespaceStats\) [ClearCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L258>)

<a name="NamespaceStats.ClearEvictions"></a>
### func \(\*NamespaceStats\) [ClearEvictions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L263>)

<a name="NamespaceStats.ClearHitRatePercent"></a>
### func \(\*NamespaceStats\) [ClearHitRatePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L248>)

<a name="NamespaceStats.ClearLastAccessTime"></a>
### func \(\*NamespaceStats\) [ClearLastAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L278>)

<a name="NamespaceStats.ClearMemoryUsageBytes"></a>
### func \(\*NamespaceStats\) [ClearMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L243>)

<a name="NamespaceStats.ClearTotalKeys"></a>
### func \(\*NamespaceStats\) [ClearTotalKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L238>)

<a name="NamespaceStats.GetAvgKeySizeBytes"></a>
### func \(\*NamespaceStats\) [GetAvgKeySizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L110>)

<a name="NamespaceStats.GetAvgValueSizeBytes"></a>
### func \(\*NamespaceStats\) [GetAvgValueSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L117>)

<a name="NamespaceStats.GetCacheHits"></a>
### func \(\*NamespaceStats\) [GetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L89>)

<a name="NamespaceStats.GetCacheMisses"></a>
### func \(\*NamespaceStats\) [GetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L96>)

<a name="NamespaceStats.GetEvictions"></a>
### func \(\*NamespaceStats\) [GetEvictions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L103>)

<a name="NamespaceStats.GetHitRatePercent"></a>
### func \(\*NamespaceStats\) [GetHitRatePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L82>)

<a name="NamespaceStats.GetLastAccessTime"></a>
### func \(\*NamespaceStats\) [GetLastAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L124>)

<a name="NamespaceStats.GetMemoryUsageBytes"></a>
### func \(\*NamespaceStats\) [GetMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L75>)

<a name="NamespaceStats.GetTotalKeys"></a>
### func \(\*NamespaceStats\) [GetTotalKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L68>)

<a name="NamespaceStats.HasAvgKeySizeBytes"></a>
### func \(\*NamespaceStats\) [HasAvgKeySizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L217>)

<a name="NamespaceStats.HasAvgValueSizeBytes"></a>
### func \(\*NamespaceStats\) [HasAvgValueSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L224>)

<a name="NamespaceStats.HasCacheHits"></a>
### func \(\*NamespaceStats\) [HasCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L196>)

<a name="NamespaceStats.HasCacheMisses"></a>
### func \(\*NamespaceStats\) [HasCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L203>)

<a name="NamespaceStats.HasEvictions"></a>
### func \(\*NamespaceStats\) [HasEvictions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L210>)

<a name="NamespaceStats.HasHitRatePercent"></a>
### func \(\*NamespaceStats\) [HasHitRatePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L189>)

<a name="NamespaceStats.HasLastAccessTime"></a>
### func \(\*NamespaceStats\) [HasLastAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L231>)

<a name="NamespaceStats.HasMemoryUsageBytes"></a>
### func \(\*NamespaceStats\) [HasMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L182>)

<a name="NamespaceStats.HasTotalKeys"></a>
### func \(\*NamespaceStats\) [HasTotalKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L175>)

<a name="NamespaceStats.ProtoMessage"></a>
### func \(\*NamespaceStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L54>)

<a name="NamespaceStats.ProtoReflect"></a>
### func \(\*NamespaceStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L56>)

<a name="NamespaceStats.Reset"></a>
### func \(\*NamespaceStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L43>)

<a name="NamespaceStats.SetAvgKeySizeBytes"></a>
### func \(\*NamespaceStats\) [SetAvgKeySizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L161>)

<a name="NamespaceStats.SetAvgValueSizeBytes"></a>
### func \(\*NamespaceStats\) [SetAvgValueSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L166>)

<a name="NamespaceStats.SetCacheHits"></a>
### func \(\*NamespaceStats\) [SetCacheHits](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L146>)

<a name="NamespaceStats.SetCacheMisses"></a>
### func \(\*NamespaceStats\) [SetCacheMisses](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L151>)

<a name="NamespaceStats.SetEvictions"></a>
### func \(\*NamespaceStats\) [SetEvictions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L156>)

<a name="NamespaceStats.SetHitRatePercent"></a>
### func \(\*NamespaceStats\) [SetHitRatePercent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L141>)

<a name="NamespaceStats.SetLastAccessTime"></a>
### func \(\*NamespaceStats\) [SetLastAccessTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L171>)

<a name="NamespaceStats.SetMemoryUsageBytes"></a>
### func \(\*NamespaceStats\) [SetMemoryUsageBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L136>)

<a name="NamespaceStats.SetTotalKeys"></a>
### func \(\*NamespaceStats\) [SetTotalKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L131>)

<a name="NamespaceStats.String"></a>
### func \(\*NamespaceStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L50>)


## type [NamespaceStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L282-L303>)

    // contains filtered or unexported fields
### func \(NamespaceStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/namespace_stats.pb.go#L305>)

<a name="OptimizeRequest"></a>
## type [OptimizeRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L28-L38>)

    // contains filtered or unexported fields
### func \(\*OptimizeRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L122>)

<a name="OptimizeRequest.ClearNamespace"></a>
### func \(\*OptimizeRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L117>)

<a name="OptimizeRequest.GetMetadata"></a>
### func \(\*OptimizeRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L75>)

<a name="OptimizeRequest.GetNamespace"></a>
### func \(\*OptimizeRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L65>)

<a name="OptimizeRequest.HasMetadata"></a>
### func \(\*OptimizeRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L110>)

<a name="OptimizeRequest.HasNamespace"></a>
### func \(\*OptimizeRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L103>)

<a name="OptimizeRequest.ProtoMessage"></a>
### func \(\*OptimizeRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L51>)

<a name="OptimizeRequest.ProtoReflect"></a>
### func \(\*OptimizeRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L53>)

<a name="OptimizeRequest.Reset"></a>
### func \(\*OptimizeRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L40>)

<a name="OptimizeRequest.SetMetadata"></a>
### func \(\*OptimizeRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L94>)

<a name="OptimizeRequest.SetNamespace"></a>
### func \(\*OptimizeRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L89>)

<a name="OptimizeRequest.String"></a>
### func \(\*OptimizeRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L47>)


## type [OptimizeRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L127-L134>)

    // contains filtered or unexported fields
### func \(OptimizeRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/optimize_request.pb.go#L136>)
<a name="PebbleConfig"></a>
## type [PebbleConfig](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L28-L39>)

\* PebbleConfig represents Pebble\-specific configuration options for the embedded key\-value store driver.
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
### func \(\*PebbleConfig\) [ClearCacheSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L169>)
<a name="PebbleConfig.ClearCompression"></a>
### func \(\*PebbleConfig\) [ClearCompression](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L184>)
func (x *PebbleConfig) ClearCompression()
<a name="PebbleConfig.ClearMaxOpenFiles"></a>
### func \(\*PebbleConfig\) [ClearMaxOpenFiles](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L179>)
func (x *PebbleConfig) ClearMaxOpenFiles()

<a name="PebbleConfig.ClearMemtableSize"></a>
### func \(\*PebbleConfig\) [ClearMemtableSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L174>)
func (x *PebbleConfig) ClearMemtableSize()
<a name="PebbleConfig.ClearPath"></a>
### func \(\*PebbleConfig\) [ClearPath](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L164>)
func (x *PebbleConfig) ClearPath()

<a name="PebbleConfig.GetCacheSize"></a>
### func \(\*PebbleConfig\) [GetCacheSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L76>)
func (x *PebbleConfig) GetCacheSize() int64

<a name="PebbleConfig.GetCompression"></a>
### func \(\*PebbleConfig\) [GetCompression](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L97>)
func (x *PebbleConfig) GetCompression() bool

<a name="PebbleConfig.GetMaxOpenFiles"></a>
### func \(\*PebbleConfig\) [GetMaxOpenFiles](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L90>)
func (x *PebbleConfig) GetMaxOpenFiles() int32

<a name="PebbleConfig.GetMemtableSize"></a>
### func \(\*PebbleConfig\) [GetMemtableSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L83>)
func (x *PebbleConfig) GetMemtableSize() int64

<a name="PebbleConfig.GetPath"></a>
### func \(\*PebbleConfig\) [GetPath](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L66>)
func (x *PebbleConfig) GetPath() string

<a name="PebbleConfig.HasCacheSize"></a>
### func \(\*PebbleConfig\) [HasCacheSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L136>)
func (x *PebbleConfig) HasCacheSize() bool

<a name="PebbleConfig.HasCompression"></a>
### func \(\*PebbleConfig\) [HasCompression](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L157>)
func (x *PebbleConfig) HasCompression() bool

<a name="PebbleConfig.HasMaxOpenFiles"></a>
### func \(\*PebbleConfig\) [HasMaxOpenFiles](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L150>)
func (x *PebbleConfig) HasMaxOpenFiles() bool

<a name="PebbleConfig.HasMemtableSize"></a>
### func \(\*PebbleConfig\) [HasMemtableSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L143>)
func (x *PebbleConfig) HasMemtableSize() bool

<a name="PebbleConfig.HasPath"></a>
### func \(\*PebbleConfig\) [HasPath](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L129>)
func (x *PebbleConfig) HasPath() bool

<a name="PebbleConfig.ProtoMessage"></a>
### func \(\*PebbleConfig\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L52>)
func (*PebbleConfig) ProtoMessage()

<a name="PebbleConfig.ProtoReflect"></a>
### func \(\*PebbleConfig\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L54>)
func (x *PebbleConfig) ProtoReflect() protoreflect.Message

<a name="PebbleConfig.Reset"></a>
### func \(\*PebbleConfig\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L41>)
func (x *PebbleConfig) Reset()

<a name="PebbleConfig.SetCacheSize"></a>
### func \(\*PebbleConfig\) [SetCacheSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L109>)
func (x *PebbleConfig) SetCacheSize(v int64)

<a name="PebbleConfig.SetCompression"></a>
### func \(\*PebbleConfig\) [SetCompression](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L124>)
func (x *PebbleConfig) SetCompression(v bool)

<a name="PebbleConfig.SetMaxOpenFiles"></a>
### func \(\*PebbleConfig\) [SetMaxOpenFiles](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L119>)
func (x *PebbleConfig) SetMaxOpenFiles(v int32)
<a name="PebbleConfig.SetMemtableSize"></a>
### func \(\*PebbleConfig\) [SetMemtableSize](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L114>)
```go
func (x *PebbleConfig) SetMemtableSize(v int64)

<a name="PebbleConfig.SetPath"></a>
### func \(\*PebbleConfig\) [SetPath](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L104>)
func (x *PebbleConfig) SetPath(v string)
<a name="PebbleConfig.String"></a>
### func \(\*PebbleConfig\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L48>)
func (x *PebbleConfig) String() string
<a name="PebbleConfig_builder"></a>
## type [PebbleConfig\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L189-L202>)
type PebbleConfig_builder struct {
    // Path is the directory where the database files are stored
    Path *string
    // CacheSize is the size of the block cache in bytes
    CacheSize *int64
    // MemtableSize is the memtable size in bytes
    MemtableSize *int64
    // MaxOpenFiles is the maximum number of open files
    MaxOpenFiles *int32
    // Compression enables on-disk compression when true
    Compression *bool
    // contains filtered or unexported fields
}
```
<a name="PebbleConfig_builder.Build"></a>
### func \(PebbleConfig\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pebble_config.pb.go#L204>)
func (b0 PebbleConfig_builder) Build() *PebbleConfig

<a name="PipelineRequest"></a>
## type [PipelineRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L28-L39>)

\* Request to execute a batch of cache operations atomically.
type PipelineRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="PipelineRequest.ClearMetadata"></a>
### func \(\*PipelineRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L150>)
func (x *PipelineRequest) ClearMetadata()

<a name="PipelineRequest.ClearNamespace"></a>
### func \(\*PipelineRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L145>)
func (x *PipelineRequest) ClearNamespace()

<a name="PipelineRequest.ClearOperations"></a>
### func \(\*PipelineRequest\) [ClearOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L140>)
func (x *PipelineRequest) ClearOperations()

<a name="PipelineRequest.GetMetadata"></a>
### func \(\*PipelineRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L83>)
func (x *PipelineRequest) GetMetadata() *common.RequestMetadata

<a name="PipelineRequest.GetNamespace"></a>
### func \(\*PipelineRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L73>)
func (x *PipelineRequest) GetNamespace() string

<a name="PipelineRequest.GetOperations"></a>
### func \(\*PipelineRequest\) [GetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L66>)
func (x *PipelineRequest) GetOperations() []byte

<a name="PipelineRequest.HasMetadata"></a>
### func \(\*PipelineRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L133>)
func (x *PipelineRequest) HasMetadata() bool

<a name="PipelineRequest.HasNamespace"></a>
### func \(\*PipelineRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L126>)
func (x *PipelineRequest) HasNamespace() bool

<a name="PipelineRequest.HasOperations"></a>
### func \(\*PipelineRequest\) [HasOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L119>)
func (x *PipelineRequest) HasOperations() bool

<a name="PipelineRequest.ProtoMessage"></a>
### func \(\*PipelineRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L52>)
func (*PipelineRequest) ProtoMessage()

<a name="PipelineRequest.ProtoReflect"></a>
### func \(\*PipelineRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L54>)
func (x *PipelineRequest) ProtoReflect() protoreflect.Message

<a name="PipelineRequest.Reset"></a>
### func \(\*PipelineRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L41>)
func (x *PipelineRequest) Reset()

<a name="PipelineRequest.SetMetadata"></a>
### func \(\*PipelineRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L110>)
func (x *PipelineRequest) SetMetadata(v *common.RequestMetadata)

<a name="PipelineRequest.SetNamespace"></a>
### func \(\*PipelineRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L105>)
func (x *PipelineRequest) SetNamespace(v string)

<a name="PipelineRequest.SetOperations"></a>
### func \(\*PipelineRequest\) [SetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L97>)
func (x *PipelineRequest) SetOperations(v []byte)

<a name="PipelineRequest.String"></a>
### func \(\*PipelineRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L48>)
func (x *PipelineRequest) String() string

<a name="PipelineRequest_builder"></a>
## type [PipelineRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L155-L164>)


type PipelineRequest_builder struct {
    // Encoded operations in execution order
    Operations []byte
    // Optional namespace
    Namespace *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
<a name="PipelineRequest_builder.Build"></a>
### func \(PipelineRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pipeline_request.pb.go#L166>)
func (b0 PipelineRequest_builder) Build() *PipelineRequest
<a name="PoolStats"></a>
## type [PoolStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L29-L41>)

\* PoolStats provides detailed statistics about connection pool usage. Used for monitoring pool efficiency and connection management.
type PoolStats struct {

    // contains filtered or unexported fields
<a name="PoolStats.ClearAcquisitionFailures"></a>
### func \(\*PoolStats\) [ClearAcquisitionFailures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L165>)
func (x *PoolStats) ClearAcquisitionFailures()

<a name="PoolStats.ClearAvgAcquisitionTime"></a>
### func \(\*PoolStats\) [ClearAvgAcquisitionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L170>)
func (x *PoolStats) ClearAvgAcquisitionTime()

<a name="PoolStats.ClearTotalClosed"></a>
### func \(\*PoolStats\) [ClearTotalClosed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L160>)
func (x *PoolStats) ClearTotalClosed()

<a name="PoolStats.ClearTotalCreated"></a>
### func \(\*PoolStats\) [ClearTotalCreated](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L155>)
func (x *PoolStats) ClearTotalCreated()

<a name="PoolStats.GetAcquisitionFailures"></a>
### func \(\*PoolStats\) [GetAcquisitionFailures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L82>)
func (x *PoolStats) GetAcquisitionFailures() int64

<a name="PoolStats.GetAvgAcquisitionTime"></a>
### func \(\*PoolStats\) [GetAvgAcquisitionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L89>)
func (x *PoolStats) GetAvgAcquisitionTime() *durationpb.Duration

<a name="PoolStats.GetTotalClosed"></a>
### func \(\*PoolStats\) [GetTotalClosed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L75>)
func (x *PoolStats) GetTotalClosed() int64

<a name="PoolStats.GetTotalCreated"></a>
### func \(\*PoolStats\) [GetTotalCreated](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L68>)
func (x *PoolStats) GetTotalCreated() int64

<a name="PoolStats.HasAcquisitionFailures"></a>
### func \(\*PoolStats\) [HasAcquisitionFailures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L141>)
func (x *PoolStats) HasAcquisitionFailures() bool

<a name="PoolStats.HasAvgAcquisitionTime"></a>
### func \(\*PoolStats\) [HasAvgAcquisitionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L148>)
func (x *PoolStats) HasAvgAcquisitionTime() bool

<a name="PoolStats.HasTotalClosed"></a>
### func \(\*PoolStats\) [HasTotalClosed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L134>)
func (x *PoolStats) HasTotalClosed() bool

<a name="PoolStats.HasTotalCreated"></a>
### func \(\*PoolStats\) [HasTotalCreated](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L127>)
func (x *PoolStats) HasTotalCreated() bool

<a name="PoolStats.ProtoMessage"></a>
### func \(\*PoolStats\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L54>)
func (*PoolStats) ProtoMessage()

<a name="PoolStats.ProtoReflect"></a>
### func \(\*PoolStats\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L56>)
func (x *PoolStats) ProtoReflect() protoreflect.Message

<a name="PoolStats.Reset"></a>
### func \(\*PoolStats\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L43>)
func (x *PoolStats) Reset()

<a name="PoolStats.SetAcquisitionFailures"></a>
### func \(\*PoolStats\) [SetAcquisitionFailures](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L113>)
func (x *PoolStats) SetAcquisitionFailures(v int64)

<a name="PoolStats.SetAvgAcquisitionTime"></a>
### func \(\*PoolStats\) [SetAvgAcquisitionTime](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L118>)
func (x *PoolStats) SetAvgAcquisitionTime(v *durationpb.Duration)

<a name="PoolStats.SetTotalClosed"></a>
### func \(\*PoolStats\) [SetTotalClosed](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L108>)
func (x *PoolStats) SetTotalClosed(v int64)

<a name="PoolStats.SetTotalCreated"></a>
### func \(\*PoolStats\) [SetTotalCreated](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L103>)
func (x *PoolStats) SetTotalCreated(v int64)

<a name="PoolStats.String"></a>
### func \(\*PoolStats\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L50>)
func (x *PoolStats) String() string

<a name="PoolStats_builder"></a>
## type [PoolStats\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L175-L186>)


type PoolStats_builder struct {
    // Total number of connections created since pool initialization
    TotalCreated *int64
    // Total number of connections closed since pool initialization
    TotalClosed *int64
    // Number of failed attempts to acquire connections
    AcquisitionFailures *int64
    // Average time to acquire a connection from the pool
    AvgAcquisitionTime *durationpb.Duration
    // contains filtered or unexported fields
<a name="PoolStats_builder.Build"></a>
### func \(PoolStats\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/pool_stats.pb.go#L188>)
func (b0 PoolStats_builder) Build() *PoolStats
<a name="PrependRequest"></a>
## type [PrependRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L29-L41>)

\* Request to prepend data to an existing cache entry.
type PrependRequest struct {

    // contains filtered or unexported fields
<a name="PrependRequest.ClearKey"></a>
### func \(\*PrependRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L172>)
func (x *PrependRequest) ClearKey()
<a name="PrependRequest.ClearMetadata"></a>
### func \(\*PrependRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L187>)
func (x *PrependRequest) ClearMetadata()
<a name="PrependRequest.ClearNamespace"></a>
### func \(\*PrependRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L182>)
func (x *PrependRequest) ClearNamespace()
<a name="PrependRequest.ClearValue"></a>
### func \(\*PrependRequest\) [ClearValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L177>)
func (x *PrependRequest) ClearValue()

<a name="PrependRequest.GetKey"></a>
### func \(\*PrependRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L68>)
func (x *PrependRequest) GetKey() string

<a name="PrependRequest.GetMetadata"></a>
### func \(\*PrependRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L102>)
func (x *PrependRequest) GetMetadata() *common.RequestMetadata

<a name="PrependRequest.GetNamespace"></a>
### func \(\*PrependRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L92>)
func (x *PrependRequest) GetNamespace() string

<a name="PrependRequest.GetValue"></a>
### func \(\*PrependRequest\) [GetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L78>)
func (x *PrependRequest) GetValue() *anypb.Any

<a name="PrependRequest.HasKey"></a>
### func \(\*PrependRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L144>)
func (x *PrependRequest) HasKey() bool

<a name="PrependRequest.HasMetadata"></a>
### func \(\*PrependRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L165>)
func (x *PrependRequest) HasMetadata() bool

<a name="PrependRequest.HasNamespace"></a>
### func \(\*PrependRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L158>)
func (x *PrependRequest) HasNamespace() bool

<a name="PrependRequest.HasValue"></a>
### func \(\*PrependRequest\) [HasValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L151>)
func (x *PrependRequest) HasValue() bool

<a name="PrependRequest.ProtoMessage"></a>
### func \(\*PrependRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L54>)
func (*PrependRequest) ProtoMessage()

<a name="PrependRequest.ProtoReflect"></a>
### func \(\*PrependRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L56>)
func (x *PrependRequest) ProtoReflect() protoreflect.Message

<a name="PrependRequest.Reset"></a>
### func \(\*PrependRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L43>)
func (x *PrependRequest) Reset()

<a name="PrependRequest.SetKey"></a>
### func \(\*PrependRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L116>)
func (x *PrependRequest) SetKey(v string)

<a name="PrependRequest.SetMetadata"></a>
### func \(\*PrependRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L135>)
func (x *PrependRequest) SetMetadata(v *common.RequestMetadata)

<a name="PrependRequest.SetNamespace"></a>
### func \(\*PrependRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L130>)
func (x *PrependRequest) SetNamespace(v string)

<a name="PrependRequest.SetValue"></a>
### func \(\*PrependRequest\) [SetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L121>)
func (x *PrependRequest) SetValue(v *anypb.Any)

<a name="PrependRequest.String"></a>
### func \(\*PrependRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L50>)
func (x *PrependRequest) String() string
<a name="PrependRequest_builder"></a>
## type [PrependRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L192-L203>)
type PrependRequest_builder struct {
    // Cache key to modify
    Key *string
    // Value to prepend
    Value *anypb.Any
    // Optional namespace
    Namespace *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
<a name="PrependRequest_builder.Build"></a>
### func \(PrependRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/prepend_request.pb.go#L205>)
func (b0 PrependRequest_builder) Build() *PrependRequest
<a name="QueryOptions"></a>
## type [QueryOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L30-L43>)

\* QueryOptions configures behavior for database query operations. Controls result limits, timeouts, and consistency requirements.
type QueryOptions struct {

    // contains filtered or unexported fields
<a name="QueryOptions.ClearConsistency"></a>
### func \(\*QueryOptions\) [ClearConsistency](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L198>)
func (x *QueryOptions) ClearConsistency()

<a name="QueryOptions.ClearIncludeMetadata"></a>
### func \(\*QueryOptions\) [ClearIncludeMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L193>)
func (x *QueryOptions) ClearIncludeMetadata()

<a name="QueryOptions.ClearLimit"></a>
### func \(\*QueryOptions\) [ClearLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L178>)
func (x *QueryOptions) ClearLimit()

<a name="QueryOptions.ClearOffset"></a>
### func \(\*QueryOptions\) [ClearOffset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L183>)
func (x *QueryOptions) ClearOffset()

<a name="QueryOptions.ClearTimeout"></a>
### func \(\*QueryOptions\) [ClearTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L188>)
func (x *QueryOptions) ClearTimeout()

<a name="QueryOptions.GetConsistency"></a>
### func \(\*QueryOptions\) [GetConsistency](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L105>)
func (x *QueryOptions) GetConsistency() common.DatabaseConsistencyLevel

<a name="QueryOptions.GetIncludeMetadata"></a>
### func \(\*QueryOptions\) [GetIncludeMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L98>)
func (x *QueryOptions) GetIncludeMetadata() bool

<a name="QueryOptions.GetLimit"></a>
### func \(\*QueryOptions\) [GetLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L70>)
func (x *QueryOptions) GetLimit() int32

<a name="QueryOptions.GetOffset"></a>
### func \(\*QueryOptions\) [GetOffset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L77>)
func (x *QueryOptions) GetOffset() int32

<a name="QueryOptions.GetTimeout"></a>
### func \(\*QueryOptions\) [GetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L84>)
func (x *QueryOptions) GetTimeout() *durationpb.Duration

<a name="QueryOptions.HasConsistency"></a>
### func \(\*QueryOptions\) [HasConsistency](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L171>)
func (x *QueryOptions) HasConsistency() bool

<a name="QueryOptions.HasIncludeMetadata"></a>
### func \(\*QueryOptions\) [HasIncludeMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L164>)
func (x *QueryOptions) HasIncludeMetadata() bool

<a name="QueryOptions.HasLimit"></a>
### func \(\*QueryOptions\) [HasLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L143>)
func (x *QueryOptions) HasLimit() bool

<a name="QueryOptions.HasOffset"></a>
### func \(\*QueryOptions\) [HasOffset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L150>)
func (x *QueryOptions) HasOffset() bool

<a name="QueryOptions.HasTimeout"></a>
### func \(\*QueryOptions\) [HasTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L157>)
func (x *QueryOptions) HasTimeout() bool

<a name="QueryOptions.ProtoMessage"></a>
### func \(\*QueryOptions\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L56>)
func (*QueryOptions) ProtoMessage()
<a name="QueryOptions.ProtoReflect"></a>
### func \(\*QueryOptions\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L58>)
```go
func (x *QueryOptions) ProtoReflect() protoreflect.Message

<a name="QueryOptions.Reset"></a>
### func \(\*QueryOptions\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L45>)
func (x *QueryOptions) Reset()

<a name="QueryOptions.SetConsistency"></a>
### func \(\*QueryOptions\) [SetConsistency](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L138>)
func (x *QueryOptions) SetConsistency(v common.DatabaseConsistencyLevel)

<a name="QueryOptions.SetIncludeMetadata"></a>
### func \(\*QueryOptions\) [SetIncludeMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L133>)
func (x *QueryOptions) SetIncludeMetadata(v bool)

<a name="QueryOptions.SetLimit"></a>
### func \(\*QueryOptions\) [SetLimit](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L114>)
func (x *QueryOptions) SetLimit(v int32)

<a name="QueryOptions.SetOffset"></a>
### func \(\*QueryOptions\) [SetOffset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L119>)
func (x *QueryOptions) SetOffset(v int32)

<a name="QueryOptions.SetTimeout"></a>
### func \(\*QueryOptions\) [SetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L124>)
func (x *QueryOptions) SetTimeout(v *durationpb.Duration)

<a name="QueryOptions.String"></a>
### func \(\*QueryOptions\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L52>)
func (x *QueryOptions) String() string
<a name="QueryOptions_builder"></a>
## type [QueryOptions\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L203-L216>)
type QueryOptions_builder struct {
    // Maximum number of rows to return
    Limit *int32
    // Number of rows to skip for pagination
    Offset *int32
    // Query execution timeout
    Timeout *durationpb.Duration
    // Whether to include column metadata in response
    IncludeMetadata *bool
    // Read consistency level for the query
    Consistency *common.DatabaseConsistencyLevel
    // contains filtered or unexported fields
}
```
<a name="QueryOptions_builder.Build"></a>
### func \(QueryOptions\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_options.pb.go#L218>)
func (b0 QueryOptions_builder) Build() *QueryOptions

<a name="QueryParameter"></a>
## type [QueryParameter](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L29-L40>)

\* QueryParameter represents a parameter for parameterized queries. Supports typed parameters to prevent SQL injection and improve performance.
type QueryParameter struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="QueryParameter.ClearName"></a>
### func \(\*QueryParameter\) [ClearName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L141>)
func (x *QueryParameter) ClearName()

<a name="QueryParameter.ClearType"></a>
### func \(\*QueryParameter\) [ClearType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L151>)
func (x *QueryParameter) ClearType()

<a name="QueryParameter.ClearValue"></a>
### func \(\*QueryParameter\) [ClearValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L146>)
func (x *QueryParameter) ClearValue()

<a name="QueryParameter.GetName"></a>
### func \(\*QueryParameter\) [GetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L67>)
func (x *QueryParameter) GetName() string

<a name="QueryParameter.GetType"></a>
### func \(\*QueryParameter\) [GetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L91>)
func (x *QueryParameter) GetType() string

<a name="QueryParameter.GetValue"></a>
### func \(\*QueryParameter\) [GetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L77>)
func (x *QueryParameter) GetValue() *anypb.Any

<a name="QueryParameter.HasName"></a>
### func \(\*QueryParameter\) [HasName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L120>)
func (x *QueryParameter) HasName() bool

<a name="QueryParameter.HasType"></a>
### func \(\*QueryParameter\) [HasType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L134>)
func (x *QueryParameter) HasType() bool

<a name="QueryParameter.HasValue"></a>
### func \(\*QueryParameter\) [HasValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L127>)
func (x *QueryParameter) HasValue() bool

<a name="QueryParameter.ProtoMessage"></a>
### func \(\*QueryParameter\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L53>)
func (*QueryParameter) ProtoMessage()

<a name="QueryParameter.ProtoReflect"></a>
### func \(\*QueryParameter\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L55>)
func (x *QueryParameter) ProtoReflect() protoreflect.Message

<a name="QueryParameter.Reset"></a>
### func \(\*QueryParameter\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L42>)
func (x *QueryParameter) Reset()

<a name="QueryParameter.SetName"></a>
### func \(\*QueryParameter\) [SetName](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L101>)
func (x *QueryParameter) SetName(v string)

<a name="QueryParameter.SetType"></a>
### func \(\*QueryParameter\) [SetType](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L115>)
func (x *QueryParameter) SetType(v string)

<a name="QueryParameter.SetValue"></a>
### func \(\*QueryParameter\) [SetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L106>)
func (x *QueryParameter) SetValue(v *anypb.Any)

<a name="QueryParameter.String"></a>
### func \(\*QueryParameter\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L49>)
func (x *QueryParameter) String() string
<a name="QueryParameter_builder"></a>
## type [QueryParameter\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L156-L165>)
type QueryParameter_builder struct {
    // Parameter name for named parameters
    Name *string
    // Parameter value as Any type for flexibility
    Value *anypb.Any
    // Optional type hint for better query optimization
    Type *string
    // contains filtered or unexported fields
<a name="QueryParameter_builder.Build"></a>
### func \(QueryParameter\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_parameter.pb.go#L167>)
func (b0 QueryParameter_builder) Build() *QueryParameter
<a name="QueryRequest"></a>
## type [QueryRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L26-L40>)


type QueryRequest struct {

    // contains filtered or unexported fields
<a name="QueryRequest.ClearDatabase"></a>
### func \(\*QueryRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L223>)
func (x *QueryRequest) ClearDatabase()
<a name="QueryRequest.ClearMetadata"></a>
### func \(\*QueryRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L233>)
func (x *QueryRequest) ClearMetadata()
<a name="QueryRequest.ClearOptions"></a>
### func \(\*QueryRequest\) [ClearOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L228>)
func (x *QueryRequest) ClearOptions()

<a name="QueryRequest.ClearQuery"></a>
### func \(\*QueryRequest\) [ClearQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L218>)
func (x *QueryRequest) ClearQuery()

<a name="QueryRequest.ClearTransactionId"></a>
### func \(\*QueryRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L238>)
func (x *QueryRequest) ClearTransactionId()

<a name="QueryRequest.GetDatabase"></a>
### func \(\*QueryRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L91>)
func (x *QueryRequest) GetDatabase() string

<a name="QueryRequest.GetMetadata"></a>
### func \(\*QueryRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L115>)
func (x *QueryRequest) GetMetadata() *common.RequestMetadata

<a name="QueryRequest.GetOptions"></a>
### func \(\*QueryRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L101>)
func (x *QueryRequest) GetOptions() *QueryOptions

<a name="QueryRequest.GetParameters"></a>
### func \(\*QueryRequest\) [GetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L77>)
func (x *QueryRequest) GetParameters() []*QueryParameter

<a name="QueryRequest.GetQuery"></a>
### func \(\*QueryRequest\) [GetQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L67>)
func (x *QueryRequest) GetQuery() string

<a name="QueryRequest.GetTransactionId"></a>
### func \(\*QueryRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L129>)
func (x *QueryRequest) GetTransactionId() string

<a name="QueryRequest.HasDatabase"></a>
### func \(\*QueryRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L190>)
func (x *QueryRequest) HasDatabase() bool

<a name="QueryRequest.HasMetadata"></a>
### func \(\*QueryRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L204>)
func (x *QueryRequest) HasMetadata() bool

<a name="QueryRequest.HasOptions"></a>
### func \(\*QueryRequest\) [HasOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L197>)
func (x *QueryRequest) HasOptions() bool
<a name="QueryRequest.HasQuery"></a>
### func \(\*QueryRequest\) [HasQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L183>)
```go
func (x *QueryRequest) HasQuery() bool

<a name="QueryRequest.HasTransactionId"></a>
### func \(\*QueryRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L211>)
func (x *QueryRequest) HasTransactionId() bool

<a name="QueryRequest.ProtoMessage"></a>
### func \(\*QueryRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L53>)
func (*QueryRequest) ProtoMessage()

<a name="QueryRequest.ProtoReflect"></a>
### func \(\*QueryRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L55>)
func (x *QueryRequest) ProtoReflect() protoreflect.Message

<a name="QueryRequest.Reset"></a>
### func \(\*QueryRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L42>)
func (x *QueryRequest) Reset()

<a name="QueryRequest.SetDatabase"></a>
### func \(\*QueryRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L155>)
func (x *QueryRequest) SetDatabase(v string)

<a name="QueryRequest.SetMetadata"></a>
### func \(\*QueryRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L169>)
func (x *QueryRequest) SetMetadata(v *common.RequestMetadata)

<a name="QueryRequest.SetOptions"></a>
### func \(\*QueryRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L160>)
func (x *QueryRequest) SetOptions(v *QueryOptions)

<a name="QueryRequest.SetParameters"></a>
### func \(\*QueryRequest\) [SetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L144>)
func (x *QueryRequest) SetParameters(v []*QueryParameter)

<a name="QueryRequest.SetQuery"></a>
### func \(\*QueryRequest\) [SetQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L139>)
func (x *QueryRequest) SetQuery(v string)

<a name="QueryRequest.SetTransactionId"></a>
### func \(\*QueryRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L178>)
func (x *QueryRequest) SetTransactionId(v string)

<a name="QueryRequest.String"></a>
### func \(\*QueryRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L49>)
func (x *QueryRequest) String() string
<a name="QueryRequest_builder"></a>
## type [QueryRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L243-L258>)
type QueryRequest_builder struct {

    // SQL query or NoSQL query string
    Query *string
    // Query parameters for parameterized queries
    Parameters []*QueryParameter
    // Database name (optional, uses default if not specified)
    Database *string
    // Query execution options
    Options *QueryOptions
    // Request metadata for tracing and authentication
    Metadata *common.RequestMetadata
    // Transaction ID if this query is part of a transaction
    TransactionId *string
    // contains filtered or unexported fields
}
<a name="QueryRequest_builder.Build"></a>
### func \(QueryRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_request.pb.go#L260>)
func (b0 QueryRequest_builder) Build() *QueryRequest
<a name="QueryResponse"></a>
## type [QueryResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L28-L39>)
\* QueryResponse contains the results of a database query operation. Includes result data, execution statistics, and error information.
type QueryResponse struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
<a name="QueryResponse.ClearError"></a>
### func \(\*QueryResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L166>)
func (x *QueryResponse) ClearError()
<a name="QueryResponse.ClearResultSet"></a>
### func \(\*QueryResponse\) [ClearResultSet](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L156>)
func (x *QueryResponse) ClearResultSet()

<a name="QueryResponse.ClearStats"></a>
### func \(\*QueryResponse\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L161>)
func (x *QueryResponse) ClearStats()

<a name="QueryResponse.GetError"></a>
### func \(\*QueryResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L94>)
func (x *QueryResponse) GetError() *common.Error

<a name="QueryResponse.GetResultSet"></a>
### func \(\*QueryResponse\) [GetResultSet](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L66>)
func (x *QueryResponse) GetResultSet() *ResultSet

<a name="QueryResponse.GetStats"></a>
### func \(\*QueryResponse\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L80>)
func (x *QueryResponse) GetStats() *DatabaseQueryStats

<a name="QueryResponse.HasError"></a>
### func \(\*QueryResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L149>)
func (x *QueryResponse) HasError() bool

<a name="QueryResponse.HasResultSet"></a>
### func \(\*QueryResponse\) [HasResultSet](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L135>)
func (x *QueryResponse) HasResultSet() bool

<a name="QueryResponse.HasStats"></a>
### func \(\*QueryResponse\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L142>)
func (x *QueryResponse) HasStats() bool

<a name="QueryResponse.ProtoMessage"></a>
### func \(\*QueryResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L52>)
func (*QueryResponse) ProtoMessage()
<a name="QueryResponse.ProtoReflect"></a>
### func \(\*QueryResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L54>)
```go
func (x *QueryResponse) ProtoReflect() protoreflect.Message

<a name="QueryResponse.Reset"></a>
### func \(\*QueryResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L41>)
func (x *QueryResponse) Reset()
<a name="QueryResponse.SetError"></a>
### func \(\*QueryResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L126>)
func (x *QueryResponse) SetError(v *common.Error)

<a name="QueryResponse.SetResultSet"></a>
### func \(\*QueryResponse\) [SetResultSet](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L108>)
func (x *QueryResponse) SetResultSet(v *ResultSet)

<a name="QueryResponse.SetStats"></a>
### func \(\*QueryResponse\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L117>)
func (x *QueryResponse) SetStats(v *DatabaseQueryStats)

<a name="QueryResponse.String"></a>
### func \(\*QueryResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L48>)
func (x *QueryResponse) String() string
<a name="QueryResponse_builder"></a>
## type [QueryResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L171-L180>)
type QueryResponse_builder struct {
    // Query result set with data and metadata
    ResultSet *ResultSet
    // Query execution statistics and performance metrics
    Stats *DatabaseQueryStats
    // Error information if the query failed
    Error *common.Error
    // contains filtered or unexported fields
}
```
<a name="QueryResponse_builder.Build"></a>
### func \(QueryResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_response.pb.go#L182>)
func (b0 QueryResponse_builder) Build() *QueryResponse
<a name="QueryRowRequest"></a>
## type [QueryRowRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L26-L40>)
type QueryRowRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="QueryRowRequest.ClearDatabase"></a>
### func \(\*QueryRowRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L211>)
func (x *QueryRowRequest) ClearDatabase()

<a name="QueryRowRequest.ClearMetadata"></a>
### func \(\*QueryRowRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L226>)
func (x *QueryRowRequest) ClearMetadata()

<a name="QueryRowRequest.ClearOptions"></a>
### func \(\*QueryRowRequest\) [ClearOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L216>)
func (x *QueryRowRequest) ClearOptions()

<a name="QueryRowRequest.ClearQuery"></a>
### func \(\*QueryRowRequest\) [ClearQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L206>)
func (x *QueryRowRequest) ClearQuery()

<a name="QueryRowRequest.ClearTransactionId"></a>
### func \(\*QueryRowRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L221>)
func (x *QueryRowRequest) ClearTransactionId()

<a name="QueryRowRequest.GetDatabase"></a>
### func \(\*QueryRowRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L91>)
func (x *QueryRowRequest) GetDatabase() string

<a name="QueryRowRequest.GetMetadata"></a>
### func \(\*QueryRowRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L125>)
func (x *QueryRowRequest) GetMetadata() *common.RequestMetadata

<a name="QueryRowRequest.GetOptions"></a>
### func \(\*QueryRowRequest\) [GetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L101>)
func (x *QueryRowRequest) GetOptions() *QueryOptions

<a name="QueryRowRequest.GetParameters"></a>
### func \(\*QueryRowRequest\) [GetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L77>)
func (x *QueryRowRequest) GetParameters() []*QueryParameter

<a name="QueryRowRequest.GetQuery"></a>
### func \(\*QueryRowRequest\) [GetQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L67>)
func (x *QueryRowRequest) GetQuery() string

<a name="QueryRowRequest.GetTransactionId"></a>
### func \(\*QueryRowRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L115>)
func (x *QueryRowRequest) GetTransactionId() string

<a name="QueryRowRequest.HasDatabase"></a>
### func \(\*QueryRowRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L178>)
func (x *QueryRowRequest) HasDatabase() bool
<a name="QueryRowRequest.HasMetadata"></a>
### func \(\*QueryRowRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L199>)
```go
func (x *QueryRowRequest) HasMetadata() bool

<a name="QueryRowRequest.HasOptions"></a>
### func \(\*QueryRowRequest\) [HasOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L185>)
func (x *QueryRowRequest) HasOptions() bool
<a name="QueryRowRequest.HasQuery"></a>
### func \(\*QueryRowRequest\) [HasQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L171>)
func (x *QueryRowRequest) HasQuery() bool

<a name="QueryRowRequest.HasTransactionId"></a>
### func \(\*QueryRowRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L192>)
func (x *QueryRowRequest) HasTransactionId() bool

<a name="QueryRowRequest.ProtoMessage"></a>
### func \(\*QueryRowRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L53>)
func (*QueryRowRequest) ProtoMessage()

<a name="QueryRowRequest.ProtoReflect"></a>
### func \(\*QueryRowRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L55>)
func (x *QueryRowRequest) ProtoReflect() protoreflect.Message

<a name="QueryRowRequest.Reset"></a>
### func \(\*QueryRowRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L42>)
func (x *QueryRowRequest) Reset()

<a name="QueryRowRequest.SetDatabase"></a>
### func \(\*QueryRowRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L148>)
func (x *QueryRowRequest) SetDatabase(v string)

<a name="QueryRowRequest.SetMetadata"></a>
### func \(\*QueryRowRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L167>)
func (x *QueryRowRequest) SetMetadata(v *common.RequestMetadata)

<a name="QueryRowRequest.SetOptions"></a>
### func \(\*QueryRowRequest\) [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L153>)
func (x *QueryRowRequest) SetOptions(v *QueryOptions)

<a name="QueryRowRequest.SetParameters"></a>
### func \(\*QueryRowRequest\) [SetParameters](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L137>)
func (x *QueryRowRequest) SetParameters(v []*QueryParameter)

<a name="QueryRowRequest.SetQuery"></a>
### func \(\*QueryRowRequest\) [SetQuery](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L132>)
func (x *QueryRowRequest) SetQuery(v string)

<a name="QueryRowRequest.SetTransactionId"></a>
### func \(\*QueryRowRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L162>)
func (x *QueryRowRequest) SetTransactionId(v string)

<a name="QueryRowRequest.String"></a>
### func \(\*QueryRowRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L49>)
func (x *QueryRowRequest) String() string
<a name="QueryRowRequest_builder"></a>
## type [QueryRowRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L230-L245>)
type QueryRowRequest_builder struct {
    // SQL query or NoSQL query string (should return at most one row)
    Query *string
    // Query parameters for parameterized queries
    Parameters []*QueryParameter
    // Database name (optional, uses default if not specified)
    Database *string
    // Query execution options
    Options *QueryOptions
    // Transaction ID if this query should be executed within a transaction
    TransactionId *string
    // Request metadata for tracing and correlation
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="QueryRowRequest_builder.Build"></a>
### func \(QueryRowRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_request.pb.go#L247>)
func (b0 QueryRowRequest_builder) Build() *QueryRowRequest

<a name="QueryRowResponse"></a>
## type [QueryRowResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L30-L43>)

\* QueryRowResponse contains the result of a single\-row query. If no row was found, \`found\` will be false and \`values\` will be empty.
type QueryRowResponse struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="QueryRowResponse.ClearError"></a>
### func \(\*QueryRowResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L195>)
func (x *QueryRowResponse) ClearError()
<a name="QueryRowResponse.ClearFound"></a>
### func \(\*QueryRowResponse\) [ClearFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L185>)
```go
func (x *QueryRowResponse) ClearFound()

<a name="QueryRowResponse.ClearStats"></a>
### func \(\*QueryRowResponse\) [ClearStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L190>)
func (x *QueryRowResponse) ClearStats()
<a name="QueryRowResponse.GetColumns"></a>
### func \(\*QueryRowResponse\) [GetColumns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L77>)
func (x *QueryRowResponse) GetColumns() []string

<a name="QueryRowResponse.GetError"></a>
### func \(\*QueryRowResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L112>)
func (x *QueryRowResponse) GetError() *common.Error

<a name="QueryRowResponse.GetFound"></a>
### func \(\*QueryRowResponse\) [GetFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L70>)
func (x *QueryRowResponse) GetFound() bool

<a name="QueryRowResponse.GetStats"></a>
### func \(\*QueryRowResponse\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L98>)
func (x *QueryRowResponse) GetStats() *DatabaseQueryStats

<a name="QueryRowResponse.GetValues"></a>
### func \(\*QueryRowResponse\) [GetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L84>)
func (x *QueryRowResponse) GetValues() []*anypb.Any

<a name="QueryRowResponse.HasError"></a>
### func \(\*QueryRowResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L178>)
func (x *QueryRowResponse) HasError() bool

<a name="QueryRowResponse.HasFound"></a>
### func \(\*QueryRowResponse\) [HasFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L164>)
func (x *QueryRowResponse) HasFound() bool

<a name="QueryRowResponse.HasStats"></a>
### func \(\*QueryRowResponse\) [HasStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L171>)
func (x *QueryRowResponse) HasStats() bool

<a name="QueryRowResponse.ProtoMessage"></a>
### func \(\*QueryRowResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L56>)
func (*QueryRowResponse) ProtoMessage()

<a name="QueryRowResponse.ProtoReflect"></a>
### func \(\*QueryRowResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L58>)
func (x *QueryRowResponse) ProtoReflect() protoreflect.Message

<a name="QueryRowResponse.Reset"></a>
### func \(\*QueryRowResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L45>)
func (x *QueryRowResponse) Reset()

<a name="QueryRowResponse.SetColumns"></a>
### func \(\*QueryRowResponse\) [SetColumns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L131>)
func (x *QueryRowResponse) SetColumns(v []string)

<a name="QueryRowResponse.SetError"></a>
### func \(\*QueryRowResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L155>)
func (x *QueryRowResponse) SetError(v *common.Error)

<a name="QueryRowResponse.SetFound"></a>
### func \(\*QueryRowResponse\) [SetFound](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L126>)
func (x *QueryRowResponse) SetFound(v bool)

<a name="QueryRowResponse.SetStats"></a>
### func \(\*QueryRowResponse\) [SetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L146>)
func (x *QueryRowResponse) SetStats(v *DatabaseQueryStats)

<a name="QueryRowResponse.SetValues"></a>
### func \(\*QueryRowResponse\) [SetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L135>)
func (x *QueryRowResponse) SetValues(v []*anypb.Any)

<a name="QueryRowResponse.String"></a>
### func \(\*QueryRowResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L52>)
func (x *QueryRowResponse) String() string

<a name="QueryRowResponse_builder"></a>
## type [QueryRowResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L200-L213>)


type QueryRowResponse_builder struct {
    // Indicates whether a row was found
    Found *bool
    // Column names matching the returned values
    Columns []string
    // Row values encoded as generic Any values
    Values []*anypb.Any
    // Query execution statistics
    Stats *DatabaseQueryStats
    // Error information if the query failed
    Error *common.Error
    // contains filtered or unexported fields
<a name="QueryRowResponse_builder.Build"></a>
### func \(QueryRowResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/query_row_response.pb.go#L215>)
func (b0 QueryRowResponse_builder) Build() *QueryRowResponse

<a name="RestoreRequest"></a>
## type [RestoreRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L28-L39>)

\* Request to restore cache contents from a backup.
type RestoreRequest struct {

    // contains filtered or unexported fields
<a name="RestoreRequest.ClearMetadata"></a>
### func \(\*RestoreRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L150>)

```go
func (x *RestoreRequest) ClearMetadata()
```


<a name="RestoreRequest.ClearNamespace"></a>
### func \(\*RestoreRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L145>)
func (x *RestoreRequest) ClearNamespace()

<a name="RestoreRequest.ClearSource"></a>
### func \(\*RestoreRequest\) [ClearSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L140>)
func (x *RestoreRequest) ClearSource()

<a name="RestoreRequest.GetMetadata"></a>
### func \(\*RestoreRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L86>)
func (x *RestoreRequest) GetMetadata() *common.RequestMetadata

<a name="RestoreRequest.GetNamespace"></a>
### func \(\*RestoreRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L76>)
func (x *RestoreRequest) GetNamespace() string

<a name="RestoreRequest.GetSource"></a>
### func \(\*RestoreRequest\) [GetSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L66>)
func (x *RestoreRequest) GetSource() string

<a name="RestoreRequest.HasMetadata"></a>
### func \(\*RestoreRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L133>)
func (x *RestoreRequest) HasMetadata() bool

<a name="RestoreRequest.HasNamespace"></a>
### func \(\*RestoreRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L126>)
func (x *RestoreRequest) HasNamespace() bool

<a name="RestoreRequest.HasSource"></a>
### func \(\*RestoreRequest\) [HasSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L119>)
func (x *RestoreRequest) HasSource() bool

<a name="RestoreRequest.ProtoMessage"></a>
### func \(\*RestoreRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L52>)
func (*RestoreRequest) ProtoMessage()

<a name="RestoreRequest.ProtoReflect"></a>
### func \(\*RestoreRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L54>)
func (x *RestoreRequest) ProtoReflect() protoreflect.Message

<a name="RestoreRequest.Reset"></a>
### func \(\*RestoreRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L41>)
func (x *RestoreRequest) Reset()

<a name="RestoreRequest.SetMetadata"></a>
### func \(\*RestoreRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L110>)
func (x *RestoreRequest) SetMetadata(v *common.RequestMetadata)

<a name="RestoreRequest.SetNamespace"></a>
### func \(\*RestoreRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L105>)
func (x *RestoreRequest) SetNamespace(v string)

<a name="RestoreRequest.SetSource"></a>
### func \(\*RestoreRequest\) [SetSource](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L100>)
func (x *RestoreRequest) SetSource(v string)

<a name="RestoreRequest.String"></a>
### func \(\*RestoreRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L48>)
func (x *RestoreRequest) String() string
<a name="RestoreRequest_builder"></a>
## type [RestoreRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L155-L164>)
type RestoreRequest_builder struct {
    // Source backup identifier
    Source *string
    // Optional namespace to restore
    Namespace *string
    // Request metadata
    // contains filtered or unexported fields
<a name="RestoreRequest_builder.Build"></a>
### func \(RestoreRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/restore_request.pb.go#L166>)
func (b0 RestoreRequest_builder) Build() *RestoreRequest
<a name="ResultSet"></a>
## type [ResultSet](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L28-L40>)

\* ResultSet contains the results of a database query operation. Includes column metadata, data rows, and pagination information.
type ResultSet struct {

    // contains filtered or unexported fields
<a name="ResultSet.ClearHasMore"></a>
### func \(\*ResultSet\) [ClearHasMore](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L160>)
func (x *ResultSet) ClearHasMore()

<a name="ResultSet.ClearTotalCount"></a>
### func \(\*ResultSet\) [ClearTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L155>)
func (x *ResultSet) ClearTotalCount()

<a name="ResultSet.GetColumns"></a>
### func \(\*ResultSet\) [GetColumns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L67>)
func (x *ResultSet) GetColumns() []*ColumnMetadata

<a name="ResultSet.GetHasMore"></a>
### func \(\*ResultSet\) [GetHasMore](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L102>)
func (x *ResultSet) GetHasMore() bool

<a name="ResultSet.GetRows"></a>
### func \(\*ResultSet\) [GetRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L81>)
func (x *ResultSet) GetRows() []*Row

<a name="ResultSet.GetTotalCount"></a>
### func \(\*ResultSet\) [GetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L95>)
func (x *ResultSet) GetTotalCount() int64

<a name="ResultSet.HasHasMore"></a>
### func \(\*ResultSet\) [HasHasMore](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L148>)
func (x *ResultSet) HasHasMore() bool

<a name="ResultSet.HasTotalCount"></a>
### func \(\*ResultSet\) [HasTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L141>)
func (x *ResultSet) HasTotalCount() bool

<a name="ResultSet.ProtoMessage"></a>
### func \(\*ResultSet\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L53>)
func (*ResultSet) ProtoMessage()

<a name="ResultSet.ProtoReflect"></a>
### func \(\*ResultSet\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L55>)
func (x *ResultSet) ProtoReflect() protoreflect.Message

<a name="ResultSet.Reset"></a>
### func \(\*ResultSet\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L42>)
func (x *ResultSet) Reset()

<a name="ResultSet.SetColumns"></a>
### func \(\*ResultSet\) [SetColumns](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L109>)
func (x *ResultSet) SetColumns(v []*ColumnMetadata)

<a name="ResultSet.SetHasMore"></a>
### func \(\*ResultSet\) [SetHasMore](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L136>)
func (x *ResultSet) SetHasMore(v bool)

<a name="ResultSet.SetRows"></a>
### func \(\*ResultSet\) [SetRows](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L120>)
func (x *ResultSet) SetRows(v []*Row)

<a name="ResultSet.SetTotalCount"></a>
### func \(\*ResultSet\) [SetTotalCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L131>)
func (x *ResultSet) SetTotalCount(v int64)

<a name="ResultSet.String"></a>
### func \(\*ResultSet\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L49>)
func (x *ResultSet) String() string

<a name="ResultSet_builder"></a>
## type [ResultSet\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L165-L176>)


type ResultSet_builder struct {
    // Metadata for each column in the result set
    Columns []*ColumnMetadata
    // Data rows matching the query
    Rows []*Row
    // Total row count if known (for pagination)
    TotalCount *int64
    // Whether more rows are available beyond this result set
    HasMore *bool
    // contains filtered or unexported fields
<a name="ResultSet_builder.Build"></a>
### func \(ResultSet\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/result_set.pb.go#L178>)
func (b0 ResultSet_builder) Build() *ResultSet

<a name="RevertMigrationRequest"></a>
## type [RevertMigrationRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L26-L37>)


type RevertMigrationRequest struct {

    // contains filtered or unexported fields
<a name="RevertMigrationRequest.ClearDatabase"></a>
### func \(\*RevertMigrationRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L138>)

```go
func (x *RevertMigrationRequest) ClearDatabase()
```


<a name="RevertMigrationRequest.ClearMetadata"></a>
### func \(\*RevertMigrationRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L148>)
func (x *RevertMigrationRequest) ClearMetadata()
```



<a name="RevertMigrationRequest.ClearTargetVersion"></a>
### func \(\*RevertMigrationRequest\) [ClearTargetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L143>)

```go
func (x *RevertMigrationRequest) ClearTargetVersion()
```



<a name="RevertMigrationRequest.GetDatabase"></a>
### func \(\*RevertMigrationRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L64>)

```go
func (x *RevertMigrationRequest) GetDatabase() string
```



<a name="RevertMigrationRequest.GetMetadata"></a>
### func \(\*RevertMigrationRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L84>)

```go
func (x *RevertMigrationRequest) GetMetadata() *common.RequestMetadata
```



<a name="RevertMigrationRequest.GetTargetVersion"></a>
### func \(\*RevertMigrationRequest\) [GetTargetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L74>)

```go
func (x *RevertMigrationRequest) GetTargetVersion() string

<a name="RevertMigrationRequest.HasDatabase"></a>
### func \(\*RevertMigrationRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L117>)
func (x *RevertMigrationRequest) HasDatabase() bool

<a name="RevertMigrationRequest.HasMetadata"></a>
### func \(\*RevertMigrationRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L131>)
func (x *RevertMigrationRequest) HasMetadata() bool

<a name="RevertMigrationRequest.HasTargetVersion"></a>
### func \(\*RevertMigrationRequest\) [HasTargetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L124>)
func (x *RevertMigrationRequest) HasTargetVersion() bool

<a name="RevertMigrationRequest.ProtoMessage"></a>
### func \(\*RevertMigrationRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L50>)
func (*RevertMigrationRequest) ProtoMessage()

<a name="RevertMigrationRequest.ProtoReflect"></a>
### func \(\*RevertMigrationRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L52>)
func (x *RevertMigrationRequest) ProtoReflect() protoreflect.Message

<a name="RevertMigrationRequest.Reset"></a>
### func \(\*RevertMigrationRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L39>)
func (x *RevertMigrationRequest) Reset()

<a name="RevertMigrationRequest.SetDatabase"></a>
### func \(\*RevertMigrationRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L98>)
func (x *RevertMigrationRequest) SetDatabase(v string)

<a name="RevertMigrationRequest.SetMetadata"></a>
### func \(\*RevertMigrationRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L108>)
func (x *RevertMigrationRequest) SetMetadata(v *common.RequestMetadata)

<a name="RevertMigrationRequest.SetTargetVersion"></a>
### func \(\*RevertMigrationRequest\) [SetTargetVersion](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L103>)
func (x *RevertMigrationRequest) SetTargetVersion(v string)

<a name="RevertMigrationRequest.String"></a>
### func \(\*RevertMigrationRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L46>)
func (x *RevertMigrationRequest) String() string
<a name="RevertMigrationRequest_builder"></a>
## type [RevertMigrationRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L153-L162>)
type RevertMigrationRequest_builder struct {
    // Database name to apply the reversion to
    Database *string
    // Target migration version to revert to
    TargetVersion *string
    // contains filtered or unexported fields
<a name="RevertMigrationRequest_builder.Build"></a>
### func \(RevertMigrationRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_request.pb.go#L164>)
func (b0 RevertMigrationRequest_builder) Build() *RevertMigrationRequest
<a name="RevertMigrationResponse"></a>
## type [RevertMigrationResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L28-L39>)

\* RevertMigrationResponse indicates the result of a migration reversion.
type RevertMigrationResponse struct {

    // contains filtered or unexported fields
<a name="RevertMigrationResponse.ClearError"></a>
### func \(\*RevertMigrationResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L147>)
func (x *RevertMigrationResponse) ClearError()

<a name="RevertMigrationResponse.ClearRevertedTo"></a>
### func \(\*RevertMigrationResponse\) [ClearRevertedTo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L142>)
func (x *RevertMigrationResponse) ClearRevertedTo()

<a name="RevertMigrationResponse.ClearSuccess"></a>
### func \(\*RevertMigrationResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L137>)
func (x *RevertMigrationResponse) ClearSuccess()

<a name="RevertMigrationResponse.GetError"></a>
### func \(\*RevertMigrationResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L83>)
func (x *RevertMigrationResponse) GetError() *common.Error

<a name="RevertMigrationResponse.GetRevertedTo"></a>
### func \(\*RevertMigrationResponse\) [GetRevertedTo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L73>)
func (x *RevertMigrationResponse) GetRevertedTo() string

<a name="RevertMigrationResponse.GetSuccess"></a>
### func \(\*RevertMigrationResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L66>)
func (x *RevertMigrationResponse) GetSuccess() bool
<a name="RevertMigrationResponse.HasError"></a>
### func \(\*RevertMigrationResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L130>)
```go
func (x *RevertMigrationResponse) HasError() bool

<a name="RevertMigrationResponse.HasRevertedTo"></a>
### func \(\*RevertMigrationResponse\) [HasRevertedTo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L123>)
func (x *RevertMigrationResponse) HasRevertedTo() bool

<a name="RevertMigrationResponse.HasSuccess"></a>
### func \(\*RevertMigrationResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L116>)
func (x *RevertMigrationResponse) HasSuccess() bool

<a name="RevertMigrationResponse.ProtoMessage"></a>
### func \(\*RevertMigrationResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L52>)
func (*RevertMigrationResponse) ProtoMessage()

<a name="RevertMigrationResponse.ProtoReflect"></a>
### func \(\*RevertMigrationResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L54>)
func (x *RevertMigrationResponse) ProtoReflect() protoreflect.Message

<a name="RevertMigrationResponse.Reset"></a>
### func \(\*RevertMigrationResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L41>)
func (x *RevertMigrationResponse) Reset()

<a name="RevertMigrationResponse.SetError"></a>
### func \(\*RevertMigrationResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L107>)
func (x *RevertMigrationResponse) SetError(v *common.Error)

<a name="RevertMigrationResponse.SetRevertedTo"></a>
### func \(\*RevertMigrationResponse\) [SetRevertedTo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L102>)
func (x *RevertMigrationResponse) SetRevertedTo(v string)

<a name="RevertMigrationResponse.SetSuccess"></a>
### func \(\*RevertMigrationResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L97>)
func (x *RevertMigrationResponse) SetSuccess(v bool)

<a name="RevertMigrationResponse.String"></a>
### func \(\*RevertMigrationResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L48>)
func (x *RevertMigrationResponse) String() string

<a name="RevertMigrationResponse_builder"></a>
## type [RevertMigrationResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L152-L161>)


type RevertMigrationResponse_builder struct {
    // Whether the migration was reverted successfully
    Success *bool
    // Version that the database was reverted to
    RevertedTo *string
    // Error information if the revert failed
    Error *common.Error
    // contains filtered or unexported fields
}
```
<a name="RevertMigrationResponse_builder.Build"></a>
### func \(RevertMigrationResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/revert_migration_response.pb.go#L163>)
func (b0 RevertMigrationResponse_builder) Build() *RevertMigrationResponse

<a name="RollbackTransactionRequest"></a>
## type [RollbackTransactionRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L26-L36>)


type RollbackTransactionRequest struct {

    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```

<a name="RollbackTransactionRequest.ClearMetadata"></a>
### func \(\*RollbackTransactionRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L120>)

```go
func (x *RollbackTransactionRequest) ClearMetadata()

<a name="RollbackTransactionRequest.ClearTransactionId"></a>
### func \(\*RollbackTransactionRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L115>)
func (x *RollbackTransactionRequest) ClearTransactionId()

<a name="RollbackTransactionRequest.GetMetadata"></a>
### func \(\*RollbackTransactionRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L73>)
func (x *RollbackTransactionRequest) GetMetadata() *common.RequestMetadata

<a name="RollbackTransactionRequest.GetTransactionId"></a>
### func \(\*RollbackTransactionRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L63>)
func (x *RollbackTransactionRequest) GetTransactionId() string

<a name="RollbackTransactionRequest.HasMetadata"></a>
### func \(\*RollbackTransactionRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L108>)
func (x *RollbackTransactionRequest) HasMetadata() bool

<a name="RollbackTransactionRequest.HasTransactionId"></a>
### func \(\*RollbackTransactionRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L101>)
func (x *RollbackTransactionRequest) HasTransactionId() bool
```


<a name="RollbackTransactionRequest.ProtoMessage"></a>
### func \(\*RollbackTransactionRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L49>)

```go
func (*RollbackTransactionRequest) ProtoMessage()

<a name="RollbackTransactionRequest.ProtoReflect"></a>
### func \(\*RollbackTransactionRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L51>)
func (x *RollbackTransactionRequest) ProtoReflect() protoreflect.Message
<a name="RollbackTransactionRequest.Reset"></a>
### func \(\*RollbackTransactionRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L38>)
func (x *RollbackTransactionRequest) Reset()

<a name="RollbackTransactionRequest.SetMetadata"></a>
### func \(\*RollbackTransactionRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L92>)
func (x *RollbackTransactionRequest) SetMetadata(v *common.RequestMetadata)

<a name="RollbackTransactionRequest.SetTransactionId"></a>
### func \(\*RollbackTransactionRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L87>)
func (x *RollbackTransactionRequest) SetTransactionId(v string)

<a name="RollbackTransactionRequest.String"></a>
### func \(\*RollbackTransactionRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L45>)
func (x *RollbackTransactionRequest) String() string
<a name="RollbackTransactionRequest_builder"></a>
## type [RollbackTransactionRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L125-L132>)
type RollbackTransactionRequest_builder struct {
    // Transaction ID to rollback
    TransactionId *string
    // Request metadata for tracing and authentication
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="RollbackTransactionRequest_builder.Build"></a>
### func \(RollbackTransactionRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/rollback_transaction_request.pb.go#L134>)
func (b0 RollbackTransactionRequest_builder) Build() *RollbackTransactionRequest

<a name="Row"></a>
## type [Row](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L29-L38>)

\* Row represents a single row of data from a database result set. Contains column values in the same order as defined in ColumnMetadata.
type Row struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="Row.GetValues"></a>
### func \(\*Row\) [GetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L65>)
func (x *Row) GetValues() []*anypb.Any

<a name="Row.ProtoMessage"></a>
### func \(\*Row\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L51>)
func (*Row) ProtoMessage()

<a name="Row.ProtoReflect"></a>
### func \(\*Row\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L53>)
func (x *Row) ProtoReflect() protoreflect.Message

<a name="Row.Reset"></a>
### func \(\*Row\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L40>)
func (x *Row) Reset()

<a name="Row.SetValues"></a>
### func \(\*Row\) [SetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L79>)
func (x *Row) SetValues(v []*anypb.Any)

<a name="Row.String"></a>
### func \(\*Row\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L47>)
func (x *Row) String() string
<a name="Row_builder"></a>
## type [Row\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L90-L95>)
type Row_builder struct {
    // Column values in order matching the column metadata
    Values []*anypb.Any
    // contains filtered or unexported fields
<a name="Row_builder.Build"></a>
### func \(Row\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/row.pb.go#L97>)
func (b0 Row_builder) Build() *Row
<a name="RunMigrationRequest"></a>
## type [RunMigrationRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L26-L37>)


type RunMigrationRequest struct {

    // contains filtered or unexported fields
<a name="RunMigrationRequest.ClearDatabase"></a>
### func \(\*RunMigrationRequest\) [ClearDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L141>)
func (x *RunMigrationRequest) ClearDatabase()
<a name="RunMigrationRequest.ClearMetadata"></a>
### func \(\*RunMigrationRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L146>)
func (x *RunMigrationRequest) ClearMetadata()

<a name="RunMigrationRequest.GetDatabase"></a>
### func \(\*RunMigrationRequest\) [GetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L64>)
func (x *RunMigrationRequest) GetDatabase() string

<a name="RunMigrationRequest.GetMetadata"></a>
### func \(\*RunMigrationRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L88>)
func (x *RunMigrationRequest) GetMetadata() *common.RequestMetadata

<a name="RunMigrationRequest.GetScripts"></a>
### func \(\*RunMigrationRequest\) [GetScripts](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L74>)
func (x *RunMigrationRequest) GetScripts() []*MigrationScript

<a name="RunMigrationRequest.HasDatabase"></a>
### func \(\*RunMigrationRequest\) [HasDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L127>)
func (x *RunMigrationRequest) HasDatabase() bool

<a name="RunMigrationRequest.HasMetadata"></a>
### func \(\*RunMigrationRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L134>)
func (x *RunMigrationRequest) HasMetadata() bool

<a name="RunMigrationRequest.ProtoMessage"></a>
### func \(\*RunMigrationRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L50>)
func (*RunMigrationRequest) ProtoMessage()

<a name="RunMigrationRequest.ProtoReflect"></a>
### func \(\*RunMigrationRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L52>)
func (x *RunMigrationRequest) ProtoReflect() protoreflect.Message

<a name="RunMigrationRequest.Reset"></a>
### func \(\*RunMigrationRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L39>)
func (x *RunMigrationRequest) Reset()

<a name="RunMigrationRequest.SetDatabase"></a>
### func \(\*RunMigrationRequest\) [SetDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L102>)
func (x *RunMigrationRequest) SetDatabase(v string)

<a name="RunMigrationRequest.SetMetadata"></a>
### func \(\*RunMigrationRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L118>)
func (x *RunMigrationRequest) SetMetadata(v *common.RequestMetadata)

<a name="RunMigrationRequest.SetScripts"></a>
### func \(\*RunMigrationRequest\) [SetScripts](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L107>)
func (x *RunMigrationRequest) SetScripts(v []*MigrationScript)

<a name="RunMigrationRequest.String"></a>
### func \(\*RunMigrationRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L46>)
func (x *RunMigrationRequest) String() string
<a name="RunMigrationRequest_builder"></a>
## type [RunMigrationRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L151-L160>)
type RunMigrationRequest_builder struct {
    // Database name to run migrations against
    Database *string
    // List of migration scripts to execute
    Scripts []*MigrationScript
    // Request metadata for tracing and authentication
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="RunMigrationRequest_builder.Build"></a>
### func \(RunMigrationRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_request.pb.go#L162>)
func (b0 RunMigrationRequest_builder) Build() *RunMigrationRequest
<a name="RunMigrationResponse"></a>
## type [RunMigrationResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L29-L40>)
\* RunMigrationResponse contains the result of executing database migrations. Indicates success status and lists applied migration versions.
type RunMigrationResponse struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="RunMigrationResponse.ClearError"></a>
### func \(\*RunMigrationResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L132>)
func (x *RunMigrationResponse) ClearError()

<a name="RunMigrationResponse.ClearSuccess"></a>
### func \(\*RunMigrationResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L127>)
func (x *RunMigrationResponse) ClearSuccess()
<a name="RunMigrationResponse.GetAppliedVersions"></a>
### func \(\*RunMigrationResponse\) [GetAppliedVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L74>)
func (x *RunMigrationResponse) GetAppliedVersions() []string

<a name="RunMigrationResponse.GetError"></a>
### func \(\*RunMigrationResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L81>)
func (x *RunMigrationResponse) GetError() *common.Error

<a name="RunMigrationResponse.GetSuccess"></a>
### func \(\*RunMigrationResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L67>)
func (x *RunMigrationResponse) GetSuccess() bool

<a name="RunMigrationResponse.HasError"></a>
### func \(\*RunMigrationResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L120>)
func (x *RunMigrationResponse) HasError() bool

<a name="RunMigrationResponse.HasSuccess"></a>
### func \(\*RunMigrationResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L113>)
func (x *RunMigrationResponse) HasSuccess() bool

<a name="RunMigrationResponse.ProtoMessage"></a>
### func \(\*RunMigrationResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L53>)
func (*RunMigrationResponse) ProtoMessage()

<a name="RunMigrationResponse.ProtoReflect"></a>
### func \(\*RunMigrationResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L55>)
func (x *RunMigrationResponse) ProtoReflect() protoreflect.Message

<a name="RunMigrationResponse.Reset"></a>
### func \(\*RunMigrationResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L42>)
func (x *RunMigrationResponse) Reset()

<a name="RunMigrationResponse.SetAppliedVersions"></a>
### func \(\*RunMigrationResponse\) [SetAppliedVersions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L100>)
func (x *RunMigrationResponse) SetAppliedVersions(v []string)

<a name="RunMigrationResponse.SetError"></a>
### func \(\*RunMigrationResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L104>)
func (x *RunMigrationResponse) SetError(v *common.Error)

<a name="RunMigrationResponse.SetSuccess"></a>
### func \(\*RunMigrationResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L95>)
func (x *RunMigrationResponse) SetSuccess(v bool)

<a name="RunMigrationResponse.String"></a>
### func \(\*RunMigrationResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L49>)
func (x *RunMigrationResponse) String() string
<a name="RunMigrationResponse_builder"></a>
## type [RunMigrationResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L137-L146>)
type RunMigrationResponse_builder struct {
    // Whether all migrations were applied successfully
    Success *bool
    // List of migration versions that were successfully applied
    AppliedVersions []string
    // Error information if any migration failed
    Error *common.Error
    // contains filtered or unexported fields
}
```
<a name="RunMigrationResponse_builder.Build"></a>
### func \(RunMigrationResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/run_migration_response.pb.go#L148>)
func (b0 RunMigrationResponse_builder) Build() *RunMigrationResponse

<a name="ScanRequest"></a>
## type [ScanRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L28-L40>)

\* Request to scan cache keys with a cursor.
type ScanRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
<a name="ScanRequest.ClearCursor"></a>
### func \(\*ScanRequest\) [ClearCursor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L163>)
func (x *ScanRequest) ClearCursor()
<a name="ScanRequest.ClearMetadata"></a>
### func \(\*ScanRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L178>)
func (x *ScanRequest) ClearMetadata()

<a name="ScanRequest.ClearNamespace"></a>
### func \(\*ScanRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L173>)
func (x *ScanRequest) ClearNamespace()

<a name="ScanRequest.ClearPattern"></a>
### func \(\*ScanRequest\) [ClearPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L168>)
func (x *ScanRequest) ClearPattern()

<a name="ScanRequest.GetCursor"></a>
### func \(\*ScanRequest\) [GetCursor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L67>)
func (x *ScanRequest) GetCursor() string

<a name="ScanRequest.GetMetadata"></a>
### func \(\*ScanRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L97>)
func (x *ScanRequest) GetMetadata() *common.RequestMetadata

<a name="ScanRequest.GetNamespace"></a>
### func \(\*ScanRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L87>)
func (x *ScanRequest) GetNamespace() string

<a name="ScanRequest.GetPattern"></a>
### func \(\*ScanRequest\) [GetPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L77>)
func (x *ScanRequest) GetPattern() string

<a name="ScanRequest.HasCursor"></a>
### func \(\*ScanRequest\) [HasCursor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L135>)
func (x *ScanRequest) HasCursor() bool

<a name="ScanRequest.HasMetadata"></a>
### func \(\*ScanRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L156>)
func (x *ScanRequest) HasMetadata() bool

<a name="ScanRequest.HasNamespace"></a>
### func \(\*ScanRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L149>)
func (x *ScanRequest) HasNamespace() bool

<a name="ScanRequest.HasPattern"></a>
### func \(\*ScanRequest\) [HasPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L142>)
func (x *ScanRequest) HasPattern() bool

<a name="ScanRequest.ProtoMessage"></a>
### func \(\*ScanRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L53>)
func (*ScanRequest) ProtoMessage()

<a name="ScanRequest.ProtoReflect"></a>
### func \(\*ScanRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L55>)
func (x *ScanRequest) ProtoReflect() protoreflect.Message

<a name="ScanRequest.Reset"></a>
### func \(\*ScanRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L42>)
func (x *ScanRequest) Reset()

<a name="ScanRequest.SetCursor"></a>
### func \(\*ScanRequest\) [SetCursor](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L111>)
func (x *ScanRequest) SetCursor(v string)

<a name="ScanRequest.SetMetadata"></a>
### func \(\*ScanRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L126>)
func (x *ScanRequest) SetMetadata(v *common.RequestMetadata)

<a name="ScanRequest.SetNamespace"></a>
### func \(\*ScanRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L121>)
func (x *ScanRequest) SetNamespace(v string)

<a name="ScanRequest.SetPattern"></a>
### func \(\*ScanRequest\) [SetPattern](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L116>)
func (x *ScanRequest) SetPattern(v string)

<a name="ScanRequest.String"></a>
### func \(\*ScanRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L49>)
func (x *ScanRequest) String() string

<a name="ScanRequest_builder"></a>
## type [ScanRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L183-L194>)


type ScanRequest_builder struct {
    // Scan cursor position
    Cursor *string
    // Match pattern for keys
    Pattern *string
    // Optional namespace
    Namespace *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
<a name="ScanRequest_builder.Build"></a>
### func \(ScanRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/scan_request.pb.go#L196>)
func (b0 ScanRequest_builder) Build() *ScanRequest
<a name="SetMultipleRequest"></a>
## type [SetMultipleRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L29-L40>)

\* Request to set multiple cache key\-value pairs. Supports batch operations for performance optimization.
type SetMultipleRequest struct {

    // contains filtered or unexported fields
<a name="SetMultipleRequest.ClearMetadata"></a>
### func \(\*SetMultipleRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L143>)
func (x *SetMultipleRequest) ClearMetadata()

<a name="SetMultipleRequest.ClearTtl"></a>
### func \(\*SetMultipleRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L138>)
func (x *SetMultipleRequest) ClearTtl()

<a name="SetMultipleRequest.GetMetadata"></a>
### func \(\*SetMultipleRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L88>)
func (x *SetMultipleRequest) GetMetadata() *common.RequestMetadata

<a name="SetMultipleRequest.GetTtl"></a>
### func \(\*SetMultipleRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L74>)
func (x *SetMultipleRequest) GetTtl() *durationpb.Duration

<a name="SetMultipleRequest.GetValues"></a>
### func \(\*SetMultipleRequest\) [GetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L67>)
func (x *SetMultipleRequest) GetValues() map[string][]byte

<a name="SetMultipleRequest.HasMetadata"></a>
### func \(\*SetMultipleRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L131>)
func (x *SetMultipleRequest) HasMetadata() bool

<a name="SetMultipleRequest.HasTtl"></a>
### func \(\*SetMultipleRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L124>)
func (x *SetMultipleRequest) HasTtl() bool

<a name="SetMultipleRequest.ProtoMessage"></a>
### func \(\*SetMultipleRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L53>)
func (*SetMultipleRequest) ProtoMessage()

<a name="SetMultipleRequest.ProtoReflect"></a>
### func \(\*SetMultipleRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L55>)
func (x *SetMultipleRequest) ProtoReflect() protoreflect.Message

<a name="SetMultipleRequest.Reset"></a>
### func \(\*SetMultipleRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L42>)
func (x *SetMultipleRequest) Reset()

<a name="SetMultipleRequest.SetMetadata"></a>
### func \(\*SetMultipleRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L115>)
func (x *SetMultipleRequest) SetMetadata(v *common.RequestMetadata)

<a name="SetMultipleRequest.SetTtl"></a>
### func \(\*SetMultipleRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L106>)
func (x *SetMultipleRequest) SetTtl(v *durationpb.Duration)

<a name="SetMultipleRequest.SetValues"></a>
### func \(\*SetMultipleRequest\) [SetValues](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L102>)
func (x *SetMultipleRequest) SetValues(v map[string][]byte)

<a name="SetMultipleRequest.String"></a>
### func \(\*SetMultipleRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L49>)
func (x *SetMultipleRequest) String() string
<a name="SetMultipleRequest_builder"></a>
## type [SetMultipleRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L148-L157>)
type SetMultipleRequest_builder struct {
    // Map of key-value pairs to set
    Values map[string][]byte
    // TTL for the cache entries (optional)
    Ttl *durationpb.Duration
    // Request metadata for tracing
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="SetMultipleRequest_builder.Build"></a>
### func \(SetMultipleRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_request.pb.go#L159>)
func (b0 SetMultipleRequest_builder) Build() *SetMultipleRequest

<a name="SetMultipleResponse"></a>
## type [SetMultipleResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L29-L39>)

\* Response for multiple cache value set operations. Indicates success/failure of batch set operation.
type SetMultipleResponse struct {
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
<a name="SetMultipleResponse.ClearError"></a>
### func \(\*SetMultipleResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L138>)
func (x *SetMultipleResponse) ClearError()

<a name="SetMultipleResponse.ClearSetCount"></a>
### func \(\*SetMultipleResponse\) [ClearSetCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L142>)
func (x *SetMultipleResponse) ClearSetCount()
<a name="SetMultipleResponse.ClearSuccess"></a>
### func \(\*SetMultipleResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L133>)
```go
func (x *SetMultipleResponse) ClearSuccess()

<a name="SetMultipleResponse.GetError"></a>
### func \(\*SetMultipleResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L80>)
func (x *SetMultipleResponse) GetError() *common.Error
<a name="SetMultipleResponse.GetFailedKeys"></a>
### func \(\*SetMultipleResponse\) [GetFailedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L73>)
func (x *SetMultipleResponse) GetFailedKeys() []string

<a name="SetMultipleResponse.GetSetCount"></a>
### func \(\*SetMultipleResponse\) [GetSetCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L87>)
func (x *SetMultipleResponse) GetSetCount() int32

<a name="SetMultipleResponse.GetSuccess"></a>
### func \(\*SetMultipleResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L66>)
func (x *SetMultipleResponse) GetSuccess() bool

<a name="SetMultipleResponse.HasError"></a>
### func \(\*SetMultipleResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L119>)
func (x *SetMultipleResponse) HasError() bool

<a name="SetMultipleResponse.HasSetCount"></a>
### func \(\*SetMultipleResponse\) [HasSetCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L126>)
func (x *SetMultipleResponse) HasSetCount() bool

<a name="SetMultipleResponse.HasSuccess"></a>
### func \(\*SetMultipleResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L112>)
func (x *SetMultipleResponse) HasSuccess() bool

<a name="SetMultipleResponse.ProtoMessage"></a>
### func \(\*SetMultipleResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L52>)
func (*SetMultipleResponse) ProtoMessage()

<a name="SetMultipleResponse.ProtoReflect"></a>
### func \(\*SetMultipleResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L54>)
func (x *SetMultipleResponse) ProtoReflect() protoreflect.Message

<a name="SetMultipleResponse.Reset"></a>
### func \(\*SetMultipleResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L41>)
func (x *SetMultipleResponse) Reset()

<a name="SetMultipleResponse.SetError"></a>
### func \(\*SetMultipleResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L103>)
func (x *SetMultipleResponse) SetError(v *common.Error)

<a name="SetMultipleResponse.SetFailedKeys"></a>
### func \(\*SetMultipleResponse\) [SetFailedKeys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L99>)
func (x *SetMultipleResponse) SetFailedKeys(v []string)

<a name="SetMultipleResponse.SetSetCount"></a>
### func \(\*SetMultipleResponse\) [SetSetCount](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L107>)
func (x *SetMultipleResponse) SetSetCount(v int32)

<a name="SetMultipleResponse.SetSuccess"></a>
### func \(\*SetMultipleResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L94>)
func (x *SetMultipleResponse) SetSuccess(v bool)

<a name="SetMultipleResponse.String"></a>
### func \(\*SetMultipleResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L48>)
func (x *SetMultipleResponse) String() string
<a name="SetMultipleResponse_builder"></a>
## type [SetMultipleResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L147-L158>)
type SetMultipleResponse_builder struct {
    // Whether all values were successfully set
    Success *bool
    // List of keys that failed to be set
    FailedKeys []string
    // Error details if operation failed
    Error *common.Error
    // Number of keys that were successfully set
    SetCount *int32
    // contains filtered or unexported fields
}
```
<a name="SetMultipleResponse_builder.Build"></a>
### func \(SetMultipleResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_multiple_response.pb.go#L160>)
func (b0 SetMultipleResponse_builder) Build() *SetMultipleResponse

<a name="SetOptions"></a>
## type [SetOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L28-L40>)

\* Options for advanced cache set operations. Allows conditional writes and flexible expiration policies.
type SetOptions struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="SetOptions.ClearOnlyIfAbsent"></a>
### func \(\*SetOptions\) [ClearOnlyIfAbsent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L154>)
func (x *SetOptions) ClearOnlyIfAbsent()

<a name="SetOptions.ClearOnlyIfPresent"></a>
### func \(\*SetOptions\) [ClearOnlyIfPresent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L159>)
func (x *SetOptions) ClearOnlyIfPresent()

<a name="SetOptions.ClearReturnPrevious"></a>
### func \(\*SetOptions\) [ClearReturnPrevious](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L169>)
func (x *SetOptions) ClearReturnPrevious()

<a name="SetOptions.ClearTtl"></a>
### func \(\*SetOptions\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L164>)
func (x *SetOptions) ClearTtl()

<a name="SetOptions.GetOnlyIfAbsent"></a>
### func \(\*SetOptions\) [GetOnlyIfAbsent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L67>)
func (x *SetOptions) GetOnlyIfAbsent() bool

<a name="SetOptions.GetOnlyIfPresent"></a>
### func \(\*SetOptions\) [GetOnlyIfPresent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L74>)
func (x *SetOptions) GetOnlyIfPresent() bool

<a name="SetOptions.GetReturnPrevious"></a>
### func \(\*SetOptions\) [GetReturnPrevious](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L95>)
func (x *SetOptions) GetReturnPrevious() bool

<a name="SetOptions.GetTtl"></a>
### func \(\*SetOptions\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L81>)
func (x *SetOptions) GetTtl() *durationpb.Duration

<a name="SetOptions.HasOnlyIfAbsent"></a>
### func \(\*SetOptions\) [HasOnlyIfAbsent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L126>)
func (x *SetOptions) HasOnlyIfAbsent() bool

<a name="SetOptions.HasOnlyIfPresent"></a>
### func \(\*SetOptions\) [HasOnlyIfPresent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L133>)
func (x *SetOptions) HasOnlyIfPresent() bool

<a name="SetOptions.HasReturnPrevious"></a>
### func \(\*SetOptions\) [HasReturnPrevious](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L147>)
func (x *SetOptions) HasReturnPrevious() bool

<a name="SetOptions.HasTtl"></a>
### func \(\*SetOptions\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L140>)
func (x *SetOptions) HasTtl() bool

<a name="SetOptions.ProtoMessage"></a>
### func \(\*SetOptions\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L53>)
func (*SetOptions) ProtoMessage()
<a name="SetOptions.ProtoReflect"></a>
### func \(\*SetOptions\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L55>)
```go
func (x *SetOptions) ProtoReflect() protoreflect.Message

<a name="SetOptions.Reset"></a>
### func \(\*SetOptions\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L42>)
func (x *SetOptions) Reset()
<a name="SetOptions.SetOnlyIfAbsent"></a>
### func \(\*SetOptions\) [SetOnlyIfAbsent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L102>)
func (x *SetOptions) SetOnlyIfAbsent(v bool)

<a name="SetOptions.SetOnlyIfPresent"></a>
### func \(\*SetOptions\) [SetOnlyIfPresent](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L107>)
func (x *SetOptions) SetOnlyIfPresent(v bool)

<a name="SetOptions.SetReturnPrevious"></a>
### func \(\*SetOptions\) [SetReturnPrevious](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L121>)
func (x *SetOptions) SetReturnPrevious(v bool)

<a name="SetOptions.SetTtl"></a>
### func \(\*SetOptions\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L112>)
func (x *SetOptions) SetTtl(v *durationpb.Duration)

<a name="SetOptions.String"></a>
### func \(\*SetOptions\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L49>)
func (x *SetOptions) String() string
<a name="SetOptions_builder"></a>
## type [SetOptions\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L174-L185>)
type SetOptions_builder struct {
    // Only set the value if the key does not already exist
    OnlyIfAbsent *bool
    // Only set the value if the key already exists
    OnlyIfPresent *bool
    // Time-to-live for the entry
    Ttl *durationpb.Duration
    // Return the previous value if overwritten
    ReturnPrevious *bool
    // contains filtered or unexported fields
}
```
<a name="SetOptions_builder.Build"></a>
### func \(SetOptions\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_options.pb.go#L187>)
func (b0 SetOptions_builder) Build() *SetOptions

<a name="SetRequest"></a>
## type [SetRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L31-L46>)

\* Request to store a value in the cache. Supports flexible expiration policies and namespace isolation.
type SetRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="SetRequest.ClearKey"></a>
### func \(\*SetRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L237>)
func (x *SetRequest) ClearKey()

<a name="SetRequest.ClearMetadata"></a>
### func \(\*SetRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L257>)
func (x *SetRequest) ClearMetadata()

<a name="SetRequest.ClearNamespace"></a>
### func \(\*SetRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L247>)
func (x *SetRequest) ClearNamespace()

<a name="SetRequest.ClearOverwrite"></a>
### func \(\*SetRequest\) [ClearOverwrite](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L262>)
func (x *SetRequest) ClearOverwrite()

<a name="SetRequest.ClearTtl"></a>
### func \(\*SetRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L252>)
func (x *SetRequest) ClearTtl()

<a name="SetRequest.ClearValue"></a>
### func \(\*SetRequest\) [ClearValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L242>)
func (x *SetRequest) ClearValue()

<a name="SetRequest.GetEntryMetadata"></a>
### func \(\*SetRequest\) [GetEntryMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L142>)
func (x *SetRequest) GetEntryMetadata() map[string]string

<a name="SetRequest.GetKey"></a>
### func \(\*SetRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L73>)
func (x *SetRequest) GetKey() string
<a name="SetRequest.GetMetadata"></a>
### func \(\*SetRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L121>)
```go
func (x *SetRequest) GetMetadata() *common.RequestMetadata

<a name="SetRequest.GetNamespace"></a>
### func \(\*SetRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L97>)
func (x *SetRequest) GetNamespace() string
<a name="SetRequest.GetOverwrite"></a>
### func \(\*SetRequest\) [GetOverwrite](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L135>)
func (x *SetRequest) GetOverwrite() bool

<a name="SetRequest.GetTtl"></a>
### func \(\*SetRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L107>)
func (x *SetRequest) GetTtl() *durationpb.Duration

<a name="SetRequest.GetValue"></a>
### func \(\*SetRequest\) [GetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L83>)
func (x *SetRequest) GetValue() *anypb.Any

<a name="SetRequest.HasKey"></a>
### func \(\*SetRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L195>)
func (x *SetRequest) HasKey() bool

<a name="SetRequest.HasMetadata"></a>
### func \(\*SetRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L223>)
func (x *SetRequest) HasMetadata() bool

<a name="SetRequest.HasNamespace"></a>
### func \(\*SetRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L209>)
func (x *SetRequest) HasNamespace() bool

<a name="SetRequest.HasOverwrite"></a>
### func \(\*SetRequest\) [HasOverwrite](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L230>)
func (x *SetRequest) HasOverwrite() bool

<a name="SetRequest.HasTtl"></a>
### func \(\*SetRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L216>)
func (x *SetRequest) HasTtl() bool

<a name="SetRequest.HasValue"></a>
### func \(\*SetRequest\) [HasValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L202>)
func (x *SetRequest) HasValue() bool

<a name="SetRequest.ProtoMessage"></a>
### func \(\*SetRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L59>)
func (*SetRequest) ProtoMessage()

<a name="SetRequest.ProtoReflect"></a>
### func \(\*SetRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L61>)
func (x *SetRequest) ProtoReflect() protoreflect.Message

<a name="SetRequest.Reset"></a>
### func \(\*SetRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L48>)
func (x *SetRequest) Reset()

<a name="SetRequest.SetEntryMetadata"></a>
### func \(\*SetRequest\) [SetEntryMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L191>)
func (x *SetRequest) SetEntryMetadata(v map[string]string)

<a name="SetRequest.SetKey"></a>
### func \(\*SetRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L149>)
func (x *SetRequest) SetKey(v string)

<a name="SetRequest.SetMetadata"></a>
### func \(\*SetRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L177>)
func (x *SetRequest) SetMetadata(v *common.RequestMetadata)

<a name="SetRequest.SetNamespace"></a>
### func \(\*SetRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L163>)
func (x *SetRequest) SetNamespace(v string)

<a name="SetRequest.SetOverwrite"></a>
### func \(\*SetRequest\) [SetOverwrite](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L186>)
func (x *SetRequest) SetOverwrite(v bool)

<a name="SetRequest.SetTtl"></a>
### func \(\*SetRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L168>)
func (x *SetRequest) SetTtl(v *durationpb.Duration)

<a name="SetRequest.SetValue"></a>
### func \(\*SetRequest\) [SetValue](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L154>)
func (x *SetRequest) SetValue(v *anypb.Any)

<a name="SetRequest.String"></a>
### func \(\*SetRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L55>)
func (x *SetRequest) String() string
<a name="SetRequest_builder"></a>
## type [SetRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L267-L284>)
type SetRequest_builder struct {
    // Cache key to store
    // Value to store (supports any type)
    Value *anypb.Any
    // Optional namespace for cache isolation
    // Time-to-live for the cache entry (0 for no expiration)
    Ttl *durationpb.Duration
    // Request metadata for tracing and correlation
    // Whether to overwrite existing value
    Overwrite *bool
    // Entry metadata for extensibility
    EntryMetadata map[string]string
    // contains filtered or unexported fields
<a name="SetRequest_builder.Build"></a>
### func \(SetRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_request.pb.go#L286>)
func (b0 SetRequest_builder) Build() *SetRequest
<a name="SetResponse"></a>
## type [SetResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L28-L37>)

\* Response for cache set operation. Indicates success and provides operation metadata.
type SetResponse struct {
    // contains filtered or unexported fields
<a name="SetResponse.ClearOverwritten"></a>
### func \(\*SetResponse\) [ClearOverwritten](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L126>)
func (x *SetResponse) ClearOverwritten()

<a name="SetResponse.ClearSizeBytes"></a>
### func \(\*SetResponse\) [ClearSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L131>)
func (x *SetResponse) ClearSizeBytes()

<a name="SetResponse.ClearSuccess"></a>
### func \(\*SetResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L121>)
func (x *SetResponse) ClearSuccess()

<a name="SetResponse.GetOverwritten"></a>
### func \(\*SetResponse\) [GetOverwritten](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L71>)
func (x *SetResponse) GetOverwritten() bool

<a name="SetResponse.GetSizeBytes"></a>
### func \(\*SetResponse\) [GetSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L78>)
func (x *SetResponse) GetSizeBytes() int64

<a name="SetResponse.GetSuccess"></a>
### func \(\*SetResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L64>)
func (x *SetResponse) GetSuccess() bool

<a name="SetResponse.HasOverwritten"></a>
### func \(\*SetResponse\) [HasOverwritten](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L107>)
func (x *SetResponse) HasOverwritten() bool

<a name="SetResponse.HasSizeBytes"></a>
### func \(\*SetResponse\) [HasSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L114>)
func (x *SetResponse) HasSizeBytes() bool

<a name="SetResponse.HasSuccess"></a>
### func \(\*SetResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L100>)
func (x *SetResponse) HasSuccess() bool

<a name="SetResponse.ProtoMessage"></a>
### func \(\*SetResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L50>)
func (*SetResponse) ProtoMessage()

<a name="SetResponse.ProtoReflect"></a>
### func \(\*SetResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L52>)
func (x *SetResponse) ProtoReflect() protoreflect.Message

<a name="SetResponse.Reset"></a>
### func \(\*SetResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L39>)
func (x *SetResponse) Reset()

<a name="SetResponse.SetOverwritten"></a>
### func \(\*SetResponse\) [SetOverwritten](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L90>)
func (x *SetResponse) SetOverwritten(v bool)

<a name="SetResponse.SetSizeBytes"></a>
### func \(\*SetResponse\) [SetSizeBytes](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L95>)
func (x *SetResponse) SetSizeBytes(v int64)

<a name="SetResponse.SetSuccess"></a>
### func \(\*SetResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L85>)
func (x *SetResponse) SetSuccess(v bool)

<a name="SetResponse.String"></a>
### func \(\*SetResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L46>)
func (x *SetResponse) String() string

<a name="SetResponse_builder"></a>
## type [SetResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L136-L145>)


type SetResponse_builder struct {
    // Whether the operation was successful
    // Whether an existing value was overwritten
    Overwritten *bool
    // Size of the stored value in bytes
    SizeBytes *int64
    // contains filtered or unexported fields
<a name="SetResponse_builder.Build"></a>
### func \(SetResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/set_response.pb.go#L147>)
func (b0 SetResponse_builder) Build() *SetResponse
<a name="TouchExpirationRequest"></a>
## type [TouchExpirationRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L29-L41>)

\* Request to update the TTL of an existing cache key.
type TouchExpirationRequest struct {

    // contains filtered or unexported fields
<a name="TouchExpirationRequest.ClearKey"></a>
### func \(\*TouchExpirationRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L172>)
func (x *TouchExpirationRequest) ClearKey()
<a name="TouchExpirationRequest.ClearMetadata"></a>
### func \(\*TouchExpirationRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L187>)
func (x *TouchExpirationRequest) ClearMetadata()

<a name="TouchExpirationRequest.ClearNamespace"></a>
### func \(\*TouchExpirationRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L182>)
func (x *TouchExpirationRequest) ClearNamespace()

<a name="TouchExpirationRequest.ClearTtl"></a>
### func \(\*TouchExpirationRequest\) [ClearTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L177>)
func (x *TouchExpirationRequest) ClearTtl()

<a name="TouchExpirationRequest.GetKey"></a>
### func \(\*TouchExpirationRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L68>)
func (x *TouchExpirationRequest) GetKey() string

<a name="TouchExpirationRequest.GetMetadata"></a>
### func \(\*TouchExpirationRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L102>)
func (x *TouchExpirationRequest) GetMetadata() *common.RequestMetadata

<a name="TouchExpirationRequest.GetNamespace"></a>
### func \(\*TouchExpirationRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L92>)
func (x *TouchExpirationRequest) GetNamespace() string

<a name="TouchExpirationRequest.GetTtl"></a>
### func \(\*TouchExpirationRequest\) [GetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L78>)
func (x *TouchExpirationRequest) GetTtl() *durationpb.Duration

<a name="TouchExpirationRequest.HasKey"></a>
### func \(\*TouchExpirationRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L144>)
func (x *TouchExpirationRequest) HasKey() bool

<a name="TouchExpirationRequest.HasMetadata"></a>
### func \(\*TouchExpirationRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L165>)
func (x *TouchExpirationRequest) HasMetadata() bool

<a name="TouchExpirationRequest.HasNamespace"></a>
### func \(\*TouchExpirationRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L158>)
func (x *TouchExpirationRequest) HasNamespace() bool

<a name="TouchExpirationRequest.HasTtl"></a>
### func \(\*TouchExpirationRequest\) [HasTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L151>)
func (x *TouchExpirationRequest) HasTtl() bool

<a name="TouchExpirationRequest.ProtoMessage"></a>
### func \(\*TouchExpirationRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L54>)
func (*TouchExpirationRequest) ProtoMessage()

<a name="TouchExpirationRequest.ProtoReflect"></a>
### func \(\*TouchExpirationRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L56>)
func (x *TouchExpirationRequest) ProtoReflect() protoreflect.Message

<a name="TouchExpirationRequest.Reset"></a>
### func \(\*TouchExpirationRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L43>)
func (x *TouchExpirationRequest) Reset()
<a name="TouchExpirationRequest.SetKey"></a>
### func \(\*TouchExpirationRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L116>)
```go
func (x *TouchExpirationRequest) SetKey(v string)

<a name="TouchExpirationRequest.SetMetadata"></a>
### func \(\*TouchExpirationRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L135>)
func (x *TouchExpirationRequest) SetMetadata(v *common.RequestMetadata)
<a name="TouchExpirationRequest.SetNamespace"></a>
### func \(\*TouchExpirationRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L130>)
func (x *TouchExpirationRequest) SetNamespace(v string)

<a name="TouchExpirationRequest.SetTtl"></a>
### func \(\*TouchExpirationRequest\) [SetTtl](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L121>)
func (x *TouchExpirationRequest) SetTtl(v *durationpb.Duration)

<a name="TouchExpirationRequest.String"></a>
### func \(\*TouchExpirationRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L50>)
func (x *TouchExpirationRequest) String() string
<a name="TouchExpirationRequest_builder"></a>
## type [TouchExpirationRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L192-L203>)
type TouchExpirationRequest_builder struct {
    // Key to update
    Key *string
    // New TTL duration
    Ttl *durationpb.Duration
    // Optional namespace
    Namespace *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="TouchExpirationRequest_builder.Build"></a>
### func \(TouchExpirationRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_request.pb.go#L205>)
func (b0 TouchExpirationRequest_builder) Build() *TouchExpirationRequest

<a name="TouchExpirationResponse"></a>
## type [TouchExpirationResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L28-L37>)

\* Response for cache touch expiration operations. Indicates success/failure of TTL update.
type TouchExpirationResponse struct {
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
<a name="TouchExpirationResponse.ClearError"></a>
### func \(\*TouchExpirationResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L130>)
func (x *TouchExpirationResponse) ClearError()

<a name="TouchExpirationResponse.ClearKeyExisted"></a>
### func \(\*TouchExpirationResponse\) [ClearKeyExisted](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L125>)
func (x *TouchExpirationResponse) ClearKeyExisted()

<a name="TouchExpirationResponse.ClearSuccess"></a>
### func \(\*TouchExpirationResponse\) [ClearSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L120>)
func (x *TouchExpirationResponse) ClearSuccess()

<a name="TouchExpirationResponse.GetError"></a>
### func \(\*TouchExpirationResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L78>)
func (x *TouchExpirationResponse) GetError() *common.Error

<a name="TouchExpirationResponse.GetKeyExisted"></a>
### func \(\*TouchExpirationResponse\) [GetKeyExisted](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L71>)
func (x *TouchExpirationResponse) GetKeyExisted() bool

<a name="TouchExpirationResponse.GetSuccess"></a>
### func \(\*TouchExpirationResponse\) [GetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L64>)
func (x *TouchExpirationResponse) GetSuccess() bool
<a name="TouchExpirationResponse.HasError"></a>
### func \(\*TouchExpirationResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L113>)
```go
func (x *TouchExpirationResponse) HasError() bool

<a name="TouchExpirationResponse.HasKeyExisted"></a>
### func \(\*TouchExpirationResponse\) [HasKeyExisted](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L106>)
func (x *TouchExpirationResponse) HasKeyExisted() bool
<a name="TouchExpirationResponse.HasSuccess"></a>
### func \(\*TouchExpirationResponse\) [HasSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L99>)
func (x *TouchExpirationResponse) HasSuccess() bool

<a name="TouchExpirationResponse.ProtoMessage"></a>
### func \(\*TouchExpirationResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L50>)
func (*TouchExpirationResponse) ProtoMessage()
<a name="TouchExpirationResponse.ProtoReflect"></a>
### func \(\*TouchExpirationResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L52>)
func (x *TouchExpirationResponse) ProtoReflect() protoreflect.Message

<a name="TouchExpirationResponse.Reset"></a>
### func \(\*TouchExpirationResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L39>)
func (x *TouchExpirationResponse) Reset()

<a name="TouchExpirationResponse.SetError"></a>
### func \(\*TouchExpirationResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L95>)
func (x *TouchExpirationResponse) SetError(v *common.Error)

<a name="TouchExpirationResponse.SetKeyExisted"></a>
### func \(\*TouchExpirationResponse\) [SetKeyExisted](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L90>)
func (x *TouchExpirationResponse) SetKeyExisted(v bool)

<a name="TouchExpirationResponse.SetSuccess"></a>
### func \(\*TouchExpirationResponse\) [SetSuccess](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L85>)
func (x *TouchExpirationResponse) SetSuccess(v bool)

<a name="TouchExpirationResponse.String"></a>
### func \(\*TouchExpirationResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L46>)
func (x *TouchExpirationResponse) String() string
<a name="TouchExpirationResponse_builder"></a>
## type [TouchExpirationResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L134-L143>)
type TouchExpirationResponse_builder struct {
    // Whether the key's TTL was successfully updated
    Success *bool
    // Whether the key existed before the touch operation
    KeyExisted *bool
    // Error details if touch failed
    Error *common.Error
    // contains filtered or unexported fields
}
```
<a name="TouchExpirationResponse_builder.Build"></a>
### func \(TouchExpirationResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/touch_expiration_response.pb.go#L145>)
func (b0 TouchExpirationResponse_builder) Build() *TouchExpirationResponse

<a name="TransactionOptions"></a>
## type [TransactionOptions](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L29-L40>)

\* TransactionOptions configures behavior for database transactions. Controls isolation level, timeout, and read\-only mode.
type TransactionOptions struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="TransactionOptions.ClearIsolation"></a>
### func \(\*TransactionOptions\) [ClearIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L137>)
func (x *TransactionOptions) ClearIsolation()

<a name="TransactionOptions.ClearReadOnly"></a>
### func \(\*TransactionOptions\) [ClearReadOnly](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L147>)
func (x *TransactionOptions) ClearReadOnly()

<a name="TransactionOptions.ClearTimeout"></a>
### func \(\*TransactionOptions\) [ClearTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L142>)
func (x *TransactionOptions) ClearTimeout()

<a name="TransactionOptions.GetIsolation"></a>
### func \(\*TransactionOptions\) [GetIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L67>)
func (x *TransactionOptions) GetIsolation() common.DatabaseIsolationLevel
<a name="TransactionOptions.GetReadOnly"></a>
### func \(\*TransactionOptions\) [GetReadOnly](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L90>)
```go
func (x *TransactionOptions) GetReadOnly() bool

<a name="TransactionOptions.GetTimeout"></a>
### func \(\*TransactionOptions\) [GetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L76>)
func (x *TransactionOptions) GetTimeout() *durationpb.Duration
<a name="TransactionOptions.HasIsolation"></a>
### func \(\*TransactionOptions\) [HasIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L116>)
func (x *TransactionOptions) HasIsolation() bool

<a name="TransactionOptions.HasReadOnly"></a>
### func \(\*TransactionOptions\) [HasReadOnly](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L130>)
func (x *TransactionOptions) HasReadOnly() bool

<a name="TransactionOptions.HasTimeout"></a>
### func \(\*TransactionOptions\) [HasTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L123>)
func (x *TransactionOptions) HasTimeout() bool

<a name="TransactionOptions.ProtoMessage"></a>
### func \(\*TransactionOptions\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L53>)
func (*TransactionOptions) ProtoMessage()

<a name="TransactionOptions.ProtoReflect"></a>
### func \(\*TransactionOptions\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L55>)
func (x *TransactionOptions) ProtoReflect() protoreflect.Message

<a name="TransactionOptions.Reset"></a>
### func \(\*TransactionOptions\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L42>)
func (x *TransactionOptions) Reset()

<a name="TransactionOptions.SetIsolation"></a>
### func \(\*TransactionOptions\) [SetIsolation](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L97>)
func (x *TransactionOptions) SetIsolation(v common.DatabaseIsolationLevel)

<a name="TransactionOptions.SetReadOnly"></a>
### func \(\*TransactionOptions\) [SetReadOnly](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L111>)
func (x *TransactionOptions) SetReadOnly(v bool)

<a name="TransactionOptions.SetTimeout"></a>
### func \(\*TransactionOptions\) [SetTimeout](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L102>)
func (x *TransactionOptions) SetTimeout(v *durationpb.Duration)

<a name="TransactionOptions.String"></a>
### func \(\*TransactionOptions\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L49>)
func (x *TransactionOptions) String() string
<a name="TransactionOptions_builder"></a>
## type [TransactionOptions\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L152-L161>)
type TransactionOptions_builder struct {
    // Isolation level for the transaction
    Isolation *common.DatabaseIsolationLevel
    // Transaction timeout before automatic rollback
    Timeout *durationpb.Duration
    // Whether this is a read-only transaction
    ReadOnly *bool
    // contains filtered or unexported fields
}
```
<a name="TransactionOptions_builder.Build"></a>
### func \(TransactionOptions\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_options.pb.go#L163>)
func (b0 TransactionOptions_builder) Build() *TransactionOptions

<a name="TransactionRequest"></a>
## type [TransactionRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L27-L37>)

\* Request to execute multiple cache operations in a transaction.
type TransactionRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
<a name="TransactionRequest.ClearMetadata"></a>
### func \(\*TransactionRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L121>)
func (x *TransactionRequest) ClearMetadata()
<a name="TransactionRequest.ClearOperations"></a>
### func \(\*TransactionRequest\) [ClearOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L116>)
func (x *TransactionRequest) ClearOperations()

<a name="TransactionRequest.GetMetadata"></a>
### func \(\*TransactionRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L71>)
func (x *TransactionRequest) GetMetadata() *common.RequestMetadata

<a name="TransactionRequest.GetOperations"></a>
### func \(\*TransactionRequest\) [GetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L64>)
func (x *TransactionRequest) GetOperations() []byte

<a name="TransactionRequest.HasMetadata"></a>
### func \(\*TransactionRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L109>)
func (x *TransactionRequest) HasMetadata() bool

<a name="TransactionRequest.HasOperations"></a>
### func \(\*TransactionRequest\) [HasOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L102>)
func (x *TransactionRequest) HasOperations() bool

<a name="TransactionRequest.ProtoMessage"></a>
### func \(\*TransactionRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L50>)
func (*TransactionRequest) ProtoMessage()

<a name="TransactionRequest.ProtoReflect"></a>
### func \(\*TransactionRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L52>)
func (x *TransactionRequest) ProtoReflect() protoreflect.Message

<a name="TransactionRequest.Reset"></a>
### func \(\*TransactionRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L39>)
func (x *TransactionRequest) Reset()
<a name="TransactionRequest.SetMetadata"></a>
### func \(\*TransactionRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L93>)
func (x *TransactionRequest) SetMetadata(v *common.RequestMetadata)

<a name="TransactionRequest.SetOperations"></a>
### func \(\*TransactionRequest\) [SetOperations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L85>)
func (x *TransactionRequest) SetOperations(v []byte)

<a name="TransactionRequest.String"></a>
### func \(\*TransactionRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L46>)
func (x *TransactionRequest) String() string
<a name="TransactionRequest_builder"></a>
## type [TransactionRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L126-L133>)
type TransactionRequest_builder struct {
    // Encoded operations in transaction
    Operations []byte
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="TransactionRequest_builder.Build"></a>
### func \(TransactionRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_request.pb.go#L135>)
func (b0 TransactionRequest_builder) Build() *TransactionRequest

<a name="TransactionServiceClient"></a>
## type [TransactionServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L35-L44>)

TransactionServiceClient is the client API for TransactionService service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.

\* TransactionService manages database transactions.
type TransactionServiceClient interface {
    // Begin a new transaction
    BeginTransaction(ctx context.Context, in *BeginTransactionRequest, opts ...grpc.CallOption) (*BeginTransactionResponse, error)
    // Commit the specified transaction
    CommitTransaction(ctx context.Context, in *CommitTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
    // Roll back the specified transaction
    RollbackTransaction(ctx context.Context, in *RollbackTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
    // Retrieve status information for a transaction
    GetTransactionStatus(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*TransactionStatusResponse, error)
}
<a name="NewTransactionServiceClient"></a>
### func [NewTransactionServiceClient](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L50>)
func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient

<a name="TransactionServiceServer"></a>
## type [TransactionServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L100-L110>)

TransactionServiceServer is the server API for TransactionService service. All implementations must embed UnimplementedTransactionServiceServer for forward compatibility.

\* TransactionService manages database transactions.
type TransactionServiceServer interface {
    // Begin a new transaction
    BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)
    // Commit the specified transaction
    CommitTransaction(context.Context, *CommitTransactionRequest) (*emptypb.Empty, error)
    // Roll back the specified transaction
    RollbackTransaction(context.Context, *RollbackTransactionRequest) (*emptypb.Empty, error)
    // Retrieve status information for a transaction
    GetTransactionStatus(context.Context, *TransactionStatusRequest) (*TransactionStatusResponse, error)
    // contains filtered or unexported methods
}
<a name="TransactionStatusRequest"></a>
## type [TransactionStatusRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L26-L36>)

type TransactionStatusRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="TransactionStatusRequest.ClearMetadata"></a>
### func \(\*TransactionStatusRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L120>)
func (x *TransactionStatusRequest) ClearMetadata()

<a name="TransactionStatusRequest.ClearTransactionId"></a>
### func \(\*TransactionStatusRequest\) [ClearTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L115>)
func (x *TransactionStatusRequest) ClearTransactionId()

<a name="TransactionStatusRequest.GetMetadata"></a>
### func \(\*TransactionStatusRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L73>)
func (x *TransactionStatusRequest) GetMetadata() *common.RequestMetadata

<a name="TransactionStatusRequest.GetTransactionId"></a>
### func \(\*TransactionStatusRequest\) [GetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L63>)
func (x *TransactionStatusRequest) GetTransactionId() string

<a name="TransactionStatusRequest.HasMetadata"></a>
### func \(\*TransactionStatusRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L108>)
func (x *TransactionStatusRequest) HasMetadata() bool

<a name="TransactionStatusRequest.HasTransactionId"></a>
### func \(\*TransactionStatusRequest\) [HasTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L101>)
func (x *TransactionStatusRequest) HasTransactionId() bool

<a name="TransactionStatusRequest.ProtoMessage"></a>
### func \(\*TransactionStatusRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L49>)
func (*TransactionStatusRequest) ProtoMessage()
<a name="TransactionStatusRequest.ProtoReflect"></a>
### func \(\*TransactionStatusRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L51>)
func (x *TransactionStatusRequest) ProtoReflect() protoreflect.Message

<a name="TransactionStatusRequest.Reset"></a>
### func \(\*TransactionStatusRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L38>)
func (x *TransactionStatusRequest) Reset()

<a name="TransactionStatusRequest.SetMetadata"></a>
### func \(\*TransactionStatusRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L92>)
func (x *TransactionStatusRequest) SetMetadata(v *common.RequestMetadata)

<a name="TransactionStatusRequest.SetTransactionId"></a>
### func \(\*TransactionStatusRequest\) [SetTransactionId](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L87>)
func (x *TransactionStatusRequest) SetTransactionId(v string)

<a name="TransactionStatusRequest.String"></a>
### func \(\*TransactionStatusRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L45>)
func (x *TransactionStatusRequest) String() string

<a name="TransactionStatusRequest_builder"></a>
## type [TransactionStatusRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L125-L132>)


type TransactionStatusRequest_builder struct {
    // Identifier of the transaction
    TransactionId *string
    // Request metadata for tracing and authentication
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="TransactionStatusRequest_builder.Build"></a>
### func \(TransactionStatusRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_request.pb.go#L134>)
func (b0 TransactionStatusRequest_builder) Build() *TransactionStatusRequest

<a name="TransactionStatusResponse"></a>
## type [TransactionStatusResponse](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L28-L38>)

\* TransactionStatusResponse returns the current status of a transaction.
type TransactionStatusResponse struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="TransactionStatusResponse.ClearError"></a>
### func \(\*TransactionStatusResponse\) [ClearError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L122>)
func (x *TransactionStatusResponse) ClearError()

<a name="TransactionStatusResponse.ClearStatus"></a>
### func \(\*TransactionStatusResponse\) [ClearStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L117>)
func (x *TransactionStatusResponse) ClearStatus()
<a name="TransactionStatusResponse.GetError"></a>
### func \(\*TransactionStatusResponse\) [GetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L75>)
func (x *TransactionStatusResponse) GetError() *common.Error

<a name="TransactionStatusResponse.GetStatus"></a>
### func \(\*TransactionStatusResponse\) [GetStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L65>)
func (x *TransactionStatusResponse) GetStatus() string

<a name="TransactionStatusResponse.HasError"></a>
### func \(\*TransactionStatusResponse\) [HasError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L110>)
func (x *TransactionStatusResponse) HasError() bool

<a name="TransactionStatusResponse.HasStatus"></a>
### func \(\*TransactionStatusResponse\) [HasStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L103>)
func (x *TransactionStatusResponse) HasStatus() bool

<a name="TransactionStatusResponse.ProtoMessage"></a>
### func \(\*TransactionStatusResponse\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L51>)
func (*TransactionStatusResponse) ProtoMessage()

<a name="TransactionStatusResponse.ProtoReflect"></a>
### func \(\*TransactionStatusResponse\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L53>)
func (x *TransactionStatusResponse) ProtoReflect() protoreflect.Message

<a name="TransactionStatusResponse.Reset"></a>
### func \(\*TransactionStatusResponse\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L40>)
func (x *TransactionStatusResponse) Reset()

<a name="TransactionStatusResponse.SetError"></a>
### func \(\*TransactionStatusResponse\) [SetError](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L94>)
func (x *TransactionStatusResponse) SetError(v *common.Error)

<a name="TransactionStatusResponse.SetStatus"></a>
### func \(\*TransactionStatusResponse\) [SetStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L89>)
func (x *TransactionStatusResponse) SetStatus(v string)
<a name="TransactionStatusResponse.String"></a>
### func \(\*TransactionStatusResponse\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L47>)
func (x *TransactionStatusResponse) String() string
<a name="TransactionStatusResponse_builder"></a>
## type [TransactionStatusResponse\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L127-L134>)
type TransactionStatusResponse_builder struct {
    // Current status of the transaction (e.g., ACTIVE, COMMITTED, ROLLED_BACK)
    Status *string
    // Error information if the transaction encountered an issue
    Error *common.Error
    // contains filtered or unexported fields
}
```
<a name="TransactionStatusResponse_builder.Build"></a>
### func \(TransactionStatusResponse\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_status_response.pb.go#L136>)
func (b0 TransactionStatusResponse_builder) Build() *TransactionStatusResponse
<a name="UnimplementedCacheAdminServiceServer"></a>
## type [UnimplementedCacheAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L132>)
UnimplementedCacheAdminServiceServer must be embedded to have forward compatible implementations.
NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedCacheAdminServiceServer struct{}
<a name="UnimplementedCacheAdminServiceServer.ConfigurePolicy"></a>
### func \(UnimplementedCacheAdminServiceServer\) [ConfigurePolicy](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L146>)
func (UnimplementedCacheAdminServiceServer) ConfigurePolicy(context.Context, *ConfigurePolicyRequest) (*ConfigurePolicyResponse, error)
<a name="UnimplementedCacheAdminServiceServer.CreateNamespace"></a>
### func \(UnimplementedCacheAdminServiceServer\) [CreateNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L134>)
func (UnimplementedCacheAdminServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)

<a name="UnimplementedCacheAdminServiceServer.DeleteNamespace"></a>
### func \(UnimplementedCacheAdminServiceServer\) [DeleteNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L137>)
func (UnimplementedCacheAdminServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error)

<a name="UnimplementedCacheAdminServiceServer.GetNamespaceStats"></a>
### func \(UnimplementedCacheAdminServiceServer\) [GetNamespaceStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L143>)
func (UnimplementedCacheAdminServiceServer) GetNamespaceStats(context.Context, *GetNamespaceStatsRequest) (*GetNamespaceStatsResponse, error)

<a name="UnimplementedCacheAdminServiceServer.ListNamespaces"></a>
### func \(UnimplementedCacheAdminServiceServer\) [ListNamespaces](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L140>)
func (UnimplementedCacheAdminServiceServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
<a name="UnimplementedCacheServiceServer"></a>
## type [UnimplementedCacheServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L270>)
UnimplementedCacheServiceServer must be embedded to have forward compatible implementations.
NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedCacheServiceServer struct{}
<a name="UnimplementedCacheServiceServer.Clear"></a>
### func \(UnimplementedCacheServiceServer\) [Clear](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L299>)
func (UnimplementedCacheServiceServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error)
<a name="UnimplementedCacheServiceServer.Decrement"></a>
### func \(UnimplementedCacheServiceServer\) [Decrement](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L296>)
func (UnimplementedCacheServiceServer) Decrement(context.Context, *DecrementRequest) (*DecrementResponse, error)
<a name="UnimplementedCacheServiceServer.Delete"></a>
### func \(UnimplementedCacheServiceServer\) [Delete](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L278>)
func (UnimplementedCacheServiceServer) Delete(context.Context, *CacheDeleteRequest) (*CacheDeleteResponse, error)

<a name="UnimplementedCacheServiceServer.DeleteMultiple"></a>
### func \(UnimplementedCacheServiceServer\) [DeleteMultiple](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L290>)
func (UnimplementedCacheServiceServer) DeleteMultiple(context.Context, *DeleteMultipleRequest) (*DeleteMultipleResponse, error)

<a name="UnimplementedCacheServiceServer.Exists"></a>
### func \(UnimplementedCacheServiceServer\) [Exists](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L281>)
func (UnimplementedCacheServiceServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)

<a name="UnimplementedCacheServiceServer.Flush"></a>
### func \(UnimplementedCacheServiceServer\) [Flush](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L308>)
func (UnimplementedCacheServiceServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error)

<a name="UnimplementedCacheServiceServer.Get"></a>
### func \(UnimplementedCacheServiceServer\) [Get](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L272>)
func (UnimplementedCacheServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error)

<a name="UnimplementedCacheServiceServer.GetMultiple"></a>
### func \(UnimplementedCacheServiceServer\) [GetMultiple](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L284>)
func (UnimplementedCacheServiceServer) GetMultiple(context.Context, *GetMultipleRequest) (*GetMultipleResponse, error)

<a name="UnimplementedCacheServiceServer.GetStats"></a>
### func \(UnimplementedCacheServiceServer\) [GetStats](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L305>)
func (UnimplementedCacheServiceServer) GetStats(context.Context, *CacheGetStatsRequest) (*CacheGetStatsResponse, error)

<a name="UnimplementedCacheServiceServer.Increment"></a>
### func \(UnimplementedCacheServiceServer\) [Increment](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L293>)
func (UnimplementedCacheServiceServer) Increment(context.Context, *IncrementRequest) (*IncrementResponse, error)

<a name="UnimplementedCacheServiceServer.Keys"></a>
### func \(UnimplementedCacheServiceServer\) [Keys](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L302>)
func (UnimplementedCacheServiceServer) Keys(context.Context, *KeysRequest) (*KeysResponse, error)

<a name="UnimplementedCacheServiceServer.Set"></a>
### func \(UnimplementedCacheServiceServer\) [Set](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L275>)
func (UnimplementedCacheServiceServer) Set(context.Context, *SetRequest) (*SetResponse, error)

<a name="UnimplementedCacheServiceServer.SetMultiple"></a>
### func \(UnimplementedCacheServiceServer\) [SetMultiple](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L287>)
func (UnimplementedCacheServiceServer) SetMultiple(context.Context, *SetMultipleRequest) (*SetMultipleResponse, error)

<a name="UnimplementedCacheServiceServer.TouchExpiration"></a>
### func \(UnimplementedCacheServiceServer\) [TouchExpiration](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L311>)
func (UnimplementedCacheServiceServer) TouchExpiration(context.Context, *TouchExpirationRequest) (*TouchExpirationResponse, error)
<a name="UnimplementedDatabaseAdminServiceServer"></a>
## type [UnimplementedDatabaseAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L164>)
UnimplementedDatabaseAdminServiceServer must be embedded to have forward compatible implementations.
NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedDatabaseAdminServiceServer struct{}
<a name="UnimplementedDatabaseAdminServiceServer.CreateDatabase"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [CreateDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L166>)
func (UnimplementedDatabaseAdminServiceServer) CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error)
<a name="UnimplementedDatabaseAdminServiceServer.CreateSchema"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [CreateSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L178>)
```go
func (UnimplementedDatabaseAdminServiceServer) CreateSchema(context.Context, *CreateSchemaRequest) (*CreateSchemaResponse, error)

<a name="UnimplementedDatabaseAdminServiceServer.DropDatabase"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [DropDatabase](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L169>)
func (UnimplementedDatabaseAdminServiceServer) DropDatabase(context.Context, *DropDatabaseRequest) (*emptypb.Empty, error)
<a name="UnimplementedDatabaseAdminServiceServer.DropSchema"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [DropSchema](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L181>)
func (UnimplementedDatabaseAdminServiceServer) DropSchema(context.Context, *DropSchemaRequest) (*emptypb.Empty, error)
<a name="UnimplementedDatabaseAdminServiceServer.GetDatabaseInfo"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [GetDatabaseInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L175>)
func (UnimplementedDatabaseAdminServiceServer) GetDatabaseInfo(context.Context, *GetDatabaseInfoRequest) (*GetDatabaseInfoResponse, error)
<a name="UnimplementedDatabaseAdminServiceServer.ListDatabases"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [ListDatabases](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L172>)
func (UnimplementedDatabaseAdminServiceServer) ListDatabases(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
<a name="UnimplementedDatabaseAdminServiceServer.ListSchemas"></a>
### func \(UnimplementedDatabaseAdminServiceServer\) [ListSchemas](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L184>)
func (UnimplementedDatabaseAdminServiceServer) ListSchemas(context.Context, *ListSchemasRequest) (*ListSchemasResponse, error)
<a name="UnimplementedDatabaseServiceServer"></a>
## type [UnimplementedDatabaseServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L148>)

UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.

NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedDatabaseServiceServer struct{}
<a name="UnimplementedDatabaseServiceServer.Execute"></a>
### func \(UnimplementedDatabaseServiceServer\) [Execute](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L156>)
func (UnimplementedDatabaseServiceServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
<a name="UnimplementedDatabaseServiceServer.ExecuteBatch"></a>
### func \(UnimplementedDatabaseServiceServer\) [ExecuteBatch](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L159>)
func (UnimplementedDatabaseServiceServer) ExecuteBatch(context.Context, *ExecuteBatchRequest) (*ExecuteBatchResponse, error)

<a name="UnimplementedDatabaseServiceServer.GetConnectionInfo"></a>
### func \(UnimplementedDatabaseServiceServer\) [GetConnectionInfo](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L162>)
func (UnimplementedDatabaseServiceServer) GetConnectionInfo(context.Context, *GetConnectionInfoRequest) (*GetConnectionInfoResponse, error)

<a name="UnimplementedDatabaseServiceServer.HealthCheck"></a>
### func \(UnimplementedDatabaseServiceServer\) [HealthCheck](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L165>)
func (UnimplementedDatabaseServiceServer) HealthCheck(context.Context, *DatabaseHealthCheckRequest) (*DatabaseHealthCheckResponse, error)

<a name="UnimplementedDatabaseServiceServer.Query"></a>
### func \(UnimplementedDatabaseServiceServer\) [Query](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L150>)
func (UnimplementedDatabaseServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error)

<a name="UnimplementedDatabaseServiceServer.QueryRow"></a>
### func \(UnimplementedDatabaseServiceServer\) [QueryRow](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L153>)
func (UnimplementedDatabaseServiceServer) QueryRow(context.Context, *QueryRowRequest) (*QueryRowResponse, error)
<a name="UnimplementedMigrationServiceServer"></a>
## type [UnimplementedMigrationServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L116>)
UnimplementedMigrationServiceServer must be embedded to have forward compatible implementations.
NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedMigrationServiceServer struct{}
<a name="UnimplementedMigrationServiceServer.ApplyMigration"></a>
### func \(UnimplementedMigrationServiceServer\) [ApplyMigration](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L118>)
func (UnimplementedMigrationServiceServer) ApplyMigration(context.Context, *RunMigrationRequest) (*RunMigrationResponse, error)

<a name="UnimplementedMigrationServiceServer.GetMigrationStatus"></a>
### func \(UnimplementedMigrationServiceServer\) [GetMigrationStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L124>)
func (UnimplementedMigrationServiceServer) GetMigrationStatus(context.Context, *GetMigrationStatusRequest) (*GetMigrationStatusResponse, error)

<a name="UnimplementedMigrationServiceServer.ListMigrations"></a>
### func \(UnimplementedMigrationServiceServer\) [ListMigrations](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L127>)
func (UnimplementedMigrationServiceServer) ListMigrations(context.Context, *ListMigrationsRequest) (*ListMigrationsResponse, error)

<a name="UnimplementedMigrationServiceServer.RevertMigration"></a>
### func \(UnimplementedMigrationServiceServer\) [RevertMigration](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L121>)
func (UnimplementedMigrationServiceServer) RevertMigration(context.Context, *RevertMigrationRequest) (*RevertMigrationResponse, error)
<a name="UnimplementedTransactionServiceServer"></a>
## type [UnimplementedTransactionServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L117>)
UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
NOTE: this should be embedded by value instead of pointer to avoid a nil pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}
<a name="UnimplementedTransactionServiceServer.BeginTransaction"></a>
### func \(UnimplementedTransactionServiceServer\) [BeginTransaction](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L119>)
func (UnimplementedTransactionServiceServer) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)

<a name="UnimplementedTransactionServiceServer.CommitTransaction"></a>
### func \(UnimplementedTransactionServiceServer\) [CommitTransaction](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L122>)
func (UnimplementedTransactionServiceServer) CommitTransaction(context.Context, *CommitTransactionRequest) (*emptypb.Empty, error)

<a name="UnimplementedTransactionServiceServer.GetTransactionStatus"></a>
### func \(UnimplementedTransactionServiceServer\) [GetTransactionStatus](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L128>)
func (UnimplementedTransactionServiceServer) GetTransactionStatus(context.Context, *TransactionStatusRequest) (*TransactionStatusResponse, error)

<a name="UnimplementedTransactionServiceServer.RollbackTransaction"></a>
### func \(UnimplementedTransactionServiceServer\) [RollbackTransaction](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L125>)
func (UnimplementedTransactionServiceServer) RollbackTransaction(context.Context, *RollbackTransactionRequest) (*emptypb.Empty, error)
<a name="UnlockRequest"></a>
## type [UnlockRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L28-L39>)
\* Request to release a previously acquired cache lock.
type UnlockRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="UnlockRequest.ClearKey"></a>
### func \(\*UnlockRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L140>)
func (x *UnlockRequest) ClearKey()

<a name="UnlockRequest.ClearMetadata"></a>
### func \(\*UnlockRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L150>)
func (x *UnlockRequest) ClearMetadata()

<a name="UnlockRequest.ClearNamespace"></a>
### func \(\*UnlockRequest\) [ClearNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L145>)
func (x *UnlockRequest) ClearNamespace()

<a name="UnlockRequest.GetKey"></a>
### func \(\*UnlockRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L66>)
func (x *UnlockRequest) GetKey() string

<a name="UnlockRequest.GetMetadata"></a>
### func \(\*UnlockRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L86>)
func (x *UnlockRequest) GetMetadata() *common.RequestMetadata

<a name="UnlockRequest.GetNamespace"></a>
### func \(\*UnlockRequest\) [GetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L76>)
func (x *UnlockRequest) GetNamespace() string

<a name="UnlockRequest.HasKey"></a>
### func \(\*UnlockRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L119>)
func (x *UnlockRequest) HasKey() bool

<a name="UnlockRequest.HasMetadata"></a>
### func \(\*UnlockRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L133>)
func (x *UnlockRequest) HasMetadata() bool

<a name="UnlockRequest.HasNamespace"></a>
### func \(\*UnlockRequest\) [HasNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L126>)
func (x *UnlockRequest) HasNamespace() bool

<a name="UnlockRequest.ProtoMessage"></a>
### func \(\*UnlockRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L52>)
func (*UnlockRequest) ProtoMessage()

<a name="UnlockRequest.ProtoReflect"></a>
### func \(\*UnlockRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L54>)
func (x *UnlockRequest) ProtoReflect() protoreflect.Message

<a name="UnlockRequest.Reset"></a>
### func \(\*UnlockRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L41>)
func (x *UnlockRequest) Reset()

<a name="UnlockRequest.SetKey"></a>
### func \(\*UnlockRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L100>)
func (x *UnlockRequest) SetKey(v string)

<a name="UnlockRequest.SetMetadata"></a>
### func \(\*UnlockRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L110>)
func (x *UnlockRequest) SetMetadata(v *common.RequestMetadata)

<a name="UnlockRequest.SetNamespace"></a>
### func \(\*UnlockRequest\) [SetNamespace](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L105>)
func (x *UnlockRequest) SetNamespace(v string)

<a name="UnlockRequest.String"></a>
### func \(\*UnlockRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L48>)
func (x *UnlockRequest) String() string
<a name="UnlockRequest_builder"></a>
## type [UnlockRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L155-L164>)
type UnlockRequest_builder struct {

    // Lock key
    Key *string
    // Optional namespace
    Namespace *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
<a name="UnlockRequest_builder.Build"></a>
### func \(UnlockRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unlock_request.pb.go#L166>)
func (b0 UnlockRequest_builder) Build() *UnlockRequest
<a name="UnsafeCacheAdminServiceServer"></a>
## type [UnsafeCacheAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_admin_service_grpc.pb.go#L155-L157>)
UnsafeCacheAdminServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to CacheAdminServiceServer will result in compilation errors.
type UnsafeCacheAdminServiceServer interface {
    // contains filtered or unexported methods
}
<a name="UnsafeCacheServiceServer"></a>
## type [UnsafeCacheServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/cache_service_grpc.pb.go#L320-L322>)
UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to CacheServiceServer will result in compilation errors.
type UnsafeCacheServiceServer interface {
    // contains filtered or unexported methods
}
<a name="UnsafeDatabaseAdminServiceServer"></a>
## type [UnsafeDatabaseAdminServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_admin_service_grpc.pb.go#L193-L195>)
UnsafeDatabaseAdminServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to DatabaseAdminServiceServer will result in compilation errors.
type UnsafeDatabaseAdminServiceServer interface {
    // contains filtered or unexported methods
}
<a name="UnsafeDatabaseServiceServer"></a>
## type [UnsafeDatabaseServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/database_service_grpc.pb.go#L174-L176>)
UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to DatabaseServiceServer will result in compilation errors.
type UnsafeDatabaseServiceServer interface {
    // contains filtered or unexported methods
}
<a name="UnsafeMigrationServiceServer"></a>
## type [UnsafeMigrationServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/migration_service_grpc.pb.go#L136-L138>)
UnsafeMigrationServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to MigrationServiceServer will result in compilation errors.
type UnsafeMigrationServiceServer interface {
    // contains filtered or unexported methods
}
<a name="UnsafeTransactionServiceServer"></a>
## type [UnsafeTransactionServiceServer](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/transaction_service_grpc.pb.go#L137-L139>)
UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service. Use of this interface is not recommended, as added methods to TransactionServiceServer will result in compilation errors.
type UnsafeTransactionServiceServer interface {
    // contains filtered or unexported methods
<a name="UnwatchRequest"></a>
## type [UnwatchRequest](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L28-L38>)
\* Request to stop watching a cache key for changes.
type UnwatchRequest struct {
    // Deprecated: Do not use. This will be deleted in the near future.
    XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
    XXX_raceDetectHookData protoimpl.RaceDetectHookData
    XXX_presence           [1]uint32
    // contains filtered or unexported fields
}
```
<a name="UnwatchRequest.ClearKey"></a>
### func \(\*UnwatchRequest\) [ClearKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L117>)
func (x *UnwatchRequest) ClearKey()

<a name="UnwatchRequest.ClearMetadata"></a>
### func \(\*UnwatchRequest\) [ClearMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L122>)
func (x *UnwatchRequest) ClearMetadata()

<a name="UnwatchRequest.GetKey"></a>
### func \(\*UnwatchRequest\) [GetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L65>)
func (x *UnwatchRequest) GetKey() string

<a name="UnwatchRequest.GetMetadata"></a>
### func \(\*UnwatchRequest\) [GetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L75>)
func (x *UnwatchRequest) GetMetadata() *common.RequestMetadata

<a name="UnwatchRequest.HasKey"></a>
### func \(\*UnwatchRequest\) [HasKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L103>)
func (x *UnwatchRequest) HasKey() bool

<a name="UnwatchRequest.HasMetadata"></a>
### func \(\*UnwatchRequest\) [HasMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L110>)
func (x *UnwatchRequest) HasMetadata() bool

<a name="UnwatchRequest.ProtoMessage"></a>
### func \(\*UnwatchRequest\) [ProtoMessage](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L51>)
func (*UnwatchRequest) ProtoMessage()

<a name="UnwatchRequest.ProtoReflect"></a>
### func \(\*UnwatchRequest\) [ProtoReflect](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L53>)
func (x *UnwatchRequest) ProtoReflect() protoreflect.Message

<a name="UnwatchRequest.Reset"></a>
### func \(\*UnwatchRequest\) [Reset](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L40>)
func (x *UnwatchRequest) Reset()

<a name="UnwatchRequest.SetKey"></a>
### func \(\*UnwatchRequest\) [SetKey](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L89>)
func (x *UnwatchRequest) SetKey(v string)

<a name="UnwatchRequest.SetMetadata"></a>
### func \(\*UnwatchRequest\) [SetMetadata](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L94>)
func (x *UnwatchRequest) SetMetadata(v *common.RequestMetadata)

<a name="UnwatchRequest.String"></a>
### func \(\*UnwatchRequest\) [String](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L47>)
func (x *UnwatchRequest) String() string
<a name="UnwatchRequest_builder"></a>
## type [UnwatchRequest\\\_builder](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L127-L134>)
type UnwatchRequest_builder struct {
    // Key being watched
    Key *string
    // Request metadata
    Metadata *common.RequestMetadata
    // contains filtered or unexported fields
}
```
<a name="UnwatchRequest_builder.Build"></a>
### func \(UnwatchRequest\_builder\) [Build](<https://github.com/jdfalk/gcommon/blob/main/gcommon/sdks/go/v1/database/unwatch_request.pb.go#L136>)
func (b0 UnwatchRequest_builder) Build() *UnwatchRequest


Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
