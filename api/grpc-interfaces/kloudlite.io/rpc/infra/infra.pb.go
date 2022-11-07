// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: infra.proto

package infra

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

type GetInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagedResName string `protobuf:"bytes,1,opt,name=managedResName,proto3" json:"managedResName,omitempty"`
	ClusterId      string `protobuf:"bytes,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Namespace      string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetInput) Reset() {
	*x = GetInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInput) ProtoMessage() {}

func (x *GetInput) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInput.ProtoReflect.Descriptor instead.
func (*GetInput) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{0}
}

func (x *GetInput) GetManagedResName() string {
	if x != nil {
		return x.ManagedResName
	}
	return ""
}

func (x *GetInput) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetInput) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output map[string]string `protobuf:"bytes,1,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetOutput() map[string]string {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_infra_proto protoreflect.FileDescriptor

var file_infra_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x70, 0x0a,
	0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x30, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x09, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x42, 0x18, 0x5a, 0x16, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_infra_proto_rawDescOnce sync.Once
	file_infra_proto_rawDescData = file_infra_proto_rawDesc
)

func file_infra_proto_rawDescGZIP() []byte {
	file_infra_proto_rawDescOnce.Do(func() {
		file_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_proto_rawDescData)
	})
	return file_infra_proto_rawDescData
}

var file_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_proto_goTypes = []interface{}{
	(*GetInput)(nil), // 0: GetInput
	(*Output)(nil),   // 1: Output
	nil,              // 2: Output.OutputEntry
}
var file_infra_proto_depIdxs = []int32{
	2, // 0: Output.output:type_name -> Output.OutputEntry
	0, // 1: Infra.GetResourceOutput:input_type -> GetInput
	1, // 2: Infra.GetResourceOutput:output_type -> Output
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_proto_init() }
func file_infra_proto_init() {
	if File_infra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInput); i {
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
		file_infra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_infra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_proto_goTypes,
		DependencyIndexes: file_infra_proto_depIdxs,
		MessageInfos:      file_infra_proto_msgTypes,
	}.Build()
	File_infra_proto = out.File
	file_infra_proto_rawDesc = nil
	file_infra_proto_goTypes = nil
	file_infra_proto_depIdxs = nil
}
