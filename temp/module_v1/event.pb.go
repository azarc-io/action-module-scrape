// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: internal/domain/module/api/v1/event.proto

package module_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *Module `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *CreatedEvent) Reset() {
	*x = CreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedEvent) ProtoMessage() {}

func (x *CreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedEvent.ProtoReflect.Descriptor instead.
func (*CreatedEvent) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedEvent) GetModule() *Module {
	if x != nil {
		return x.Module
	}
	return nil
}

type SparkAddedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo    string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Spark   *Spark `protobuf:"bytes,3,opt,name=spark,proto3" json:"spark,omitempty"`
}

func (x *SparkAddedEvent) Reset() {
	*x = SparkAddedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_module_api_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparkAddedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparkAddedEvent) ProtoMessage() {}

func (x *SparkAddedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_module_api_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparkAddedEvent.ProtoReflect.Descriptor instead.
func (*SparkAddedEvent) Descriptor() ([]byte, []int) {
	return file_internal_domain_module_api_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *SparkAddedEvent) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *SparkAddedEvent) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SparkAddedEvent) GetSpark() *Spark {
	if x != nil {
		return x.Spark
	}
	return nil
}

var File_internal_domain_module_api_v1_event_proto protoreflect.FileDescriptor

var file_internal_domain_module_api_v1_event_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x39, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x67, 0x0a, 0x0f,
	0x53, 0x70, 0x61, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x05, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x52, 0x05,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x7a, 0x61, 0x72, 0x63, 0x2d, 0x69, 0x6f, 0x2f, 0x76, 0x65, 0x72,
	0x61, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_domain_module_api_v1_event_proto_rawDescOnce sync.Once
	file_internal_domain_module_api_v1_event_proto_rawDescData = file_internal_domain_module_api_v1_event_proto_rawDesc
)

func file_internal_domain_module_api_v1_event_proto_rawDescGZIP() []byte {
	file_internal_domain_module_api_v1_event_proto_rawDescOnce.Do(func() {
		file_internal_domain_module_api_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_domain_module_api_v1_event_proto_rawDescData)
	})
	return file_internal_domain_module_api_v1_event_proto_rawDescData
}

var file_internal_domain_module_api_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_domain_module_api_v1_event_proto_goTypes = []interface{}{
	(*CreatedEvent)(nil),    // 0: module_v1.CreatedEvent
	(*SparkAddedEvent)(nil), // 1: module_v1.SparkAddedEvent
	(*Module)(nil),          // 2: module_v1.Module
	(*Spark)(nil),           // 3: module_v1.Spark
}
var file_internal_domain_module_api_v1_event_proto_depIdxs = []int32{
	2, // 0: module_v1.CreatedEvent.module:type_name -> module_v1.Module
	3, // 1: module_v1.SparkAddedEvent.spark:type_name -> module_v1.Spark
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_domain_module_api_v1_event_proto_init() }
func file_internal_domain_module_api_v1_event_proto_init() {
	if File_internal_domain_module_api_v1_event_proto != nil {
		return
	}
	file_internal_domain_module_api_v1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_domain_module_api_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_domain_module_api_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparkAddedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_domain_module_api_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_domain_module_api_v1_event_proto_goTypes,
		DependencyIndexes: file_internal_domain_module_api_v1_event_proto_depIdxs,
		MessageInfos:      file_internal_domain_module_api_v1_event_proto_msgTypes,
	}.Build()
	File_internal_domain_module_api_v1_event_proto = out.File
	file_internal_domain_module_api_v1_event_proto_rawDesc = nil
	file_internal_domain_module_api_v1_event_proto_goTypes = nil
	file_internal_domain_module_api_v1_event_proto_depIdxs = nil
}
