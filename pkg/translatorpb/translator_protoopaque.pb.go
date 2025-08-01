// file: proto/translator.proto
// version: 1.2.0
// guid: 70eca88d-31fe-4044-8b76-a22bd3c9e0bd

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.27.1
// source: proto/translator.proto

//go:build protoopaque

package translatorpb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
	configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TranslateRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Meta        *proto.RequestMetadata `protobuf:"bytes,1,opt,name=meta"`
	xxx_hidden_Text        *string                `protobuf:"bytes,2,opt,name=text"`
	xxx_hidden_Language    *string                `protobuf:"bytes,3,opt,name=language"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *TranslateRequest) Reset() {
	*x = TranslateRequest{}
	mi := &file_proto_translator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranslateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRequest) ProtoMessage() {}

func (x *TranslateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TranslateRequest) GetMeta() *proto.RequestMetadata {
	if x != nil {
		return x.xxx_hidden_Meta
	}
	return nil
}

func (x *TranslateRequest) GetText() string {
	if x != nil {
		if x.xxx_hidden_Text != nil {
			return *x.xxx_hidden_Text
		}
		return ""
	}
	return ""
}

func (x *TranslateRequest) GetLanguage() string {
	if x != nil {
		if x.xxx_hidden_Language != nil {
			return *x.xxx_hidden_Language
		}
		return ""
	}
	return ""
}

func (x *TranslateRequest) SetMeta(v *proto.RequestMetadata) {
	x.xxx_hidden_Meta = v
}

func (x *TranslateRequest) SetText(v string) {
	x.xxx_hidden_Text = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *TranslateRequest) SetLanguage(v string) {
	x.xxx_hidden_Language = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *TranslateRequest) HasMeta() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Meta != nil
}

func (x *TranslateRequest) HasText() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *TranslateRequest) HasLanguage() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *TranslateRequest) ClearMeta() {
	x.xxx_hidden_Meta = nil
}

func (x *TranslateRequest) ClearText() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Text = nil
}

func (x *TranslateRequest) ClearLanguage() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Language = nil
}

type TranslateRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Meta     *proto.RequestMetadata
	Text     *string
	Language *string
}

func (b0 TranslateRequest_builder) Build() *TranslateRequest {
	m0 := &TranslateRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Meta = b.Meta
	if b.Text != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Text = b.Text
	}
	if b.Language != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Language = b.Language
	}
	return m0
}

type TranslateResponse struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TranslatedText *string                `protobuf:"bytes,1,opt,name=translated_text,json=translatedText"`
	xxx_hidden_Errors         *[]*proto.Error        `protobuf:"bytes,2,rep,name=errors"`
	XXX_raceDetectHookData    protoimpl.RaceDetectHookData
	XXX_presence              [1]uint32
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *TranslateResponse) Reset() {
	*x = TranslateResponse{}
	mi := &file_proto_translator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranslateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateResponse) ProtoMessage() {}

func (x *TranslateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TranslateResponse) GetTranslatedText() string {
	if x != nil {
		if x.xxx_hidden_TranslatedText != nil {
			return *x.xxx_hidden_TranslatedText
		}
		return ""
	}
	return ""
}

func (x *TranslateResponse) GetErrors() []*proto.Error {
	if x != nil {
		if x.xxx_hidden_Errors != nil {
			return *x.xxx_hidden_Errors
		}
	}
	return nil
}

func (x *TranslateResponse) SetTranslatedText(v string) {
	x.xxx_hidden_TranslatedText = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *TranslateResponse) SetErrors(v []*proto.Error) {
	x.xxx_hidden_Errors = &v
}

func (x *TranslateResponse) HasTranslatedText() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *TranslateResponse) ClearTranslatedText() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_TranslatedText = nil
}

type TranslateResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TranslatedText *string
	Errors         []*proto.Error
}

func (b0 TranslateResponse_builder) Build() *TranslateResponse {
	m0 := &TranslateResponse{}
	b, x := &b0, m0
	_, _ = b, x
	if b.TranslatedText != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_TranslatedText = b.TranslatedText
	}
	x.xxx_hidden_Errors = &b.Errors
	return m0
}

// RateLimit defines basic provider rate limiting parameters.
type RateLimit struct {
	state                        protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_RequestsPerMinute uint32                 `protobuf:"varint,1,opt,name=requests_per_minute,json=requestsPerMinute"`
	xxx_hidden_Burst             uint32                 `protobuf:"varint,2,opt,name=burst"`
	XXX_raceDetectHookData       protoimpl.RaceDetectHookData
	XXX_presence                 [1]uint32
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *RateLimit) Reset() {
	*x = RateLimit{}
	mi := &file_proto_translator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimit) ProtoMessage() {}

func (x *RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RateLimit) GetRequestsPerMinute() uint32 {
	if x != nil {
		return x.xxx_hidden_RequestsPerMinute
	}
	return 0
}

func (x *RateLimit) GetBurst() uint32 {
	if x != nil {
		return x.xxx_hidden_Burst
	}
	return 0
}

func (x *RateLimit) SetRequestsPerMinute(v uint32) {
	x.xxx_hidden_RequestsPerMinute = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *RateLimit) SetBurst(v uint32) {
	x.xxx_hidden_Burst = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *RateLimit) HasRequestsPerMinute() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *RateLimit) HasBurst() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *RateLimit) ClearRequestsPerMinute() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_RequestsPerMinute = 0
}

func (x *RateLimit) ClearBurst() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Burst = 0
}

type RateLimit_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	RequestsPerMinute *uint32
	Burst             *uint32
}

func (b0 RateLimit_builder) Build() *RateLimit {
	m0 := &RateLimit{}
	b, x := &b0, m0
	_, _ = b, x
	if b.RequestsPerMinute != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_RequestsPerMinute = *b.RequestsPerMinute
	}
	if b.Burst != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_Burst = *b.Burst
	}
	return m0
}

// ProviderInfo describes metadata about a provider service.
type ProviderInfo struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name         *string                `protobuf:"bytes,1,opt,name=name"`
	xxx_hidden_RateLimit    *RateLimit             `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit"`
	xxx_hidden_Capabilities []string               `protobuf:"bytes,3,rep,name=capabilities"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *ProviderInfo) Reset() {
	*x = ProviderInfo{}
	mi := &file_proto_translator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderInfo) ProtoMessage() {}

func (x *ProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_translator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ProviderInfo) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *ProviderInfo) GetRateLimit() *RateLimit {
	if x != nil {
		return x.xxx_hidden_RateLimit
	}
	return nil
}

func (x *ProviderInfo) GetCapabilities() []string {
	if x != nil {
		return x.xxx_hidden_Capabilities
	}
	return nil
}

func (x *ProviderInfo) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *ProviderInfo) SetRateLimit(v *RateLimit) {
	x.xxx_hidden_RateLimit = v
}

func (x *ProviderInfo) SetCapabilities(v []string) {
	x.xxx_hidden_Capabilities = v
}

func (x *ProviderInfo) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ProviderInfo) HasRateLimit() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_RateLimit != nil
}

func (x *ProviderInfo) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Name = nil
}

func (x *ProviderInfo) ClearRateLimit() {
	x.xxx_hidden_RateLimit = nil
}

type ProviderInfo_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Name         *string
	RateLimit    *RateLimit
	Capabilities []string
}

func (b0 ProviderInfo_builder) Build() *ProviderInfo {
	m0 := &ProviderInfo{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Name = b.Name
	}
	x.xxx_hidden_RateLimit = b.RateLimit
	x.xxx_hidden_Capabilities = b.Capabilities
	return m0
}

var File_proto_translator_proto protoreflect.FileDescriptor

const file_proto_translator_proto_rawDesc = "" +
	"\n" +
	"\x16proto/translator.proto\x12\x15gcommon.v1.translator\x1a!google/protobuf/go_features.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a0pkg/common/proto/messages/request_metadata.proto\x1a%pkg/common/proto/messages/error.proto\x1a\x12proto/config.proto\"z\n" +
	"\x10TranslateRequest\x126\n" +
	"\x04meta\x18\x01 \x01(\v2\".gcommon.v1.common.RequestMetadataR\x04meta\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x12\x1a\n" +
	"\blanguage\x18\x03 \x01(\tR\blanguage\"n\n" +
	"\x11TranslateResponse\x12'\n" +
	"\x0ftranslated_text\x18\x01 \x01(\tR\x0etranslatedText\x120\n" +
	"\x06errors\x18\x02 \x03(\v2\x18.gcommon.v1.common.ErrorR\x06errors\"Q\n" +
	"\tRateLimit\x12.\n" +
	"\x13requests_per_minute\x18\x01 \x01(\rR\x11requestsPerMinute\x12\x14\n" +
	"\x05burst\x18\x02 \x01(\rR\x05burst\"\x87\x01\n" +
	"\fProviderInfo\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12?\n" +
	"\n" +
	"rate_limit\x18\x02 \x01(\v2 .gcommon.v1.translator.RateLimitR\trateLimit\x12\"\n" +
	"\fcapabilities\x18\x03 \x03(\tR\fcapabilities2\x91\x02\n" +
	"\x11TranslatorService\x12^\n" +
	"\tTranslate\x12'.gcommon.v1.translator.TranslateRequest\x1a(.gcommon.v1.translator.TranslateResponse\x12M\n" +
	"\tGetConfig\x12\x16.google.protobuf.Empty\x1a(.gcommon.v1.config.SubtitleManagerConfig\x12M\n" +
	"\tSetConfig\x12(.gcommon.v1.config.SubtitleManagerConfig\x1a\x16.google.protobuf.EmptyBJZ@github.com/jdfalk/subtitle-manager/pkg/translatorpb;translatorpb\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_proto_translator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_translator_proto_goTypes = []any{
	(*TranslateRequest)(nil),               // 0: gcommon.v1.translator.TranslateRequest
	(*TranslateResponse)(nil),              // 1: gcommon.v1.translator.TranslateResponse
	(*RateLimit)(nil),                      // 2: gcommon.v1.translator.RateLimit
	(*ProviderInfo)(nil),                   // 3: gcommon.v1.translator.ProviderInfo
	(*proto.RequestMetadata)(nil),          // 4: gcommon.v1.common.RequestMetadata
	(*proto.Error)(nil),                    // 5: gcommon.v1.common.Error
	(*emptypb.Empty)(nil),                  // 6: google.protobuf.Empty
	(*configpb.SubtitleManagerConfig)(nil), // 7: gcommon.v1.config.SubtitleManagerConfig
}
var file_proto_translator_proto_depIdxs = []int32{
	4, // 0: gcommon.v1.translator.TranslateRequest.meta:type_name -> gcommon.v1.common.RequestMetadata
	5, // 1: gcommon.v1.translator.TranslateResponse.errors:type_name -> gcommon.v1.common.Error
	2, // 2: gcommon.v1.translator.ProviderInfo.rate_limit:type_name -> gcommon.v1.translator.RateLimit
	0, // 3: gcommon.v1.translator.TranslatorService.Translate:input_type -> gcommon.v1.translator.TranslateRequest
	6, // 4: gcommon.v1.translator.TranslatorService.GetConfig:input_type -> google.protobuf.Empty
	7, // 5: gcommon.v1.translator.TranslatorService.SetConfig:input_type -> gcommon.v1.config.SubtitleManagerConfig
	1, // 6: gcommon.v1.translator.TranslatorService.Translate:output_type -> gcommon.v1.translator.TranslateResponse
	7, // 7: gcommon.v1.translator.TranslatorService.GetConfig:output_type -> gcommon.v1.config.SubtitleManagerConfig
	6, // 8: gcommon.v1.translator.TranslatorService.SetConfig:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_translator_proto_init() }
func file_proto_translator_proto_init() {
	if File_proto_translator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_translator_proto_rawDesc), len(file_proto_translator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_translator_proto_goTypes,
		DependencyIndexes: file_proto_translator_proto_depIdxs,
		MessageInfos:      file_proto_translator_proto_msgTypes,
	}.Build()
	File_proto_translator_proto = out.File
	file_proto_translator_proto_goTypes = nil
	file_proto_translator_proto_depIdxs = nil
}
