// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: finance-infra.proto

package finance

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

type ComputeStartIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *ComputeStartIn) Reset() {
	*x = ComputeStartIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeStartIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeStartIn) ProtoMessage() {}

func (x *ComputeStartIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeStartIn.ProtoReflect.Descriptor instead.
func (*ComputeStartIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeStartIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *ComputeStartIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ComputeEndIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *ComputeEndIn) Reset() {
	*x = ComputeEndIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeEndIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeEndIn) ProtoMessage() {}

func (x *ComputeEndIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeEndIn.ProtoReflect.Descriptor instead.
func (*ComputeEndIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeEndIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *ComputeEndIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type LambdaStartIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *LambdaStartIn) Reset() {
	*x = LambdaStartIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaStartIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaStartIn) ProtoMessage() {}

func (x *LambdaStartIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaStartIn.ProtoReflect.Descriptor instead.
func (*LambdaStartIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{2}
}

func (x *LambdaStartIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *LambdaStartIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type LambdaEndIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *LambdaEndIn) Reset() {
	*x = LambdaEndIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaEndIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaEndIn) ProtoMessage() {}

func (x *LambdaEndIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaEndIn.ProtoReflect.Descriptor instead.
func (*LambdaEndIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{3}
}

func (x *LambdaEndIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *LambdaEndIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type BlockStorageStartIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *BlockStorageStartIn) Reset() {
	*x = BlockStorageStartIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStorageStartIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStorageStartIn) ProtoMessage() {}

func (x *BlockStorageStartIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStorageStartIn.ProtoReflect.Descriptor instead.
func (*BlockStorageStartIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{4}
}

func (x *BlockStorageStartIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *BlockStorageStartIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type BlockStorageEndIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *BlockStorageEndIn) Reset() {
	*x = BlockStorageEndIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStorageEndIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStorageEndIn) ProtoMessage() {}

func (x *BlockStorageEndIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStorageEndIn.ProtoReflect.Descriptor instead.
func (*BlockStorageEndIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{5}
}

func (x *BlockStorageEndIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *BlockStorageEndIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ObjectStorageStartIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *ObjectStorageStartIn) Reset() {
	*x = ObjectStorageStartIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectStorageStartIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStorageStartIn) ProtoMessage() {}

func (x *ObjectStorageStartIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStorageStartIn.ProtoReflect.Descriptor instead.
func (*ObjectStorageStartIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectStorageStartIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *ObjectStorageStartIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ObjectStorageEndIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *ObjectStorageEndIn) Reset() {
	*x = ObjectStorageEndIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectStorageEndIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStorageEndIn) ProtoMessage() {}

func (x *ObjectStorageEndIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStorageEndIn.ProtoReflect.Descriptor instead.
func (*ObjectStorageEndIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{7}
}

func (x *ObjectStorageEndIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *ObjectStorageEndIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CIStartIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *CIStartIn) Reset() {
	*x = CIStartIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIStartIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIStartIn) ProtoMessage() {}

func (x *CIStartIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIStartIn.ProtoReflect.Descriptor instead.
func (*CIStartIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{8}
}

func (x *CIStartIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *CIStartIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CIEndIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceRef string `protobuf:"bytes,1,opt,name=ResourceRef,proto3" json:"ResourceRef,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *CIEndIn) Reset() {
	*x = CIEndIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIEndIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIEndIn) ProtoMessage() {}

func (x *CIEndIn) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIEndIn.ProtoReflect.Descriptor instead.
func (*CIEndIn) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{9}
}

func (x *CIEndIn) GetResourceRef() string {
	if x != nil {
		return x.ResourceRef
	}
	return ""
}

func (x *CIEndIn) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type FinanceInfraVoid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinanceInfraVoid) Reset() {
	*x = FinanceInfraVoid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_infra_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinanceInfraVoid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinanceInfraVoid) ProtoMessage() {}

func (x *FinanceInfraVoid) ProtoReflect() protoreflect.Message {
	mi := &file_finance_infra_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinanceInfraVoid.ProtoReflect.Descriptor instead.
func (*FinanceInfraVoid) Descriptor() ([]byte, []int) {
	return file_finance_infra_proto_rawDescGZIP(), []int{10}
}

var File_finance_infra_proto protoreflect.FileDescriptor

var file_finance_infra_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x0d, 0x4c, 0x61, 0x6d, 0x62, 0x64,
	0x61, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0b, 0x4c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x53,
	0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x12, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x49,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x4b, 0x0a, 0x09, 0x43, 0x49, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x49,
	0x0a, 0x07, 0x43, 0x49, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x46, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x96, 0x04,
	0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x32,
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0f,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x1a,
	0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x12, 0x0d, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x1a,
	0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x0e, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x45, 0x6e,
	0x64, 0x12, 0x0c, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x1a,
	0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x3c, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e,
	0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f, 0x69, 0x64,
	0x12, 0x38, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x45, 0x6e, 0x64, 0x12, 0x12, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x12, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x15, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x10, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x13,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x64, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x49, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x0a, 0x2e, 0x43, 0x49, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e,
	0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x6f, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x05, 0x43, 0x49, 0x45, 0x6e, 0x64, 0x12, 0x08, 0x2e, 0x43, 0x49, 0x45, 0x6e,
	0x64, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x6b, 0x6c, 0x6f, 0x75, 0x64, 0x6c,
	0x69, 0x74, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_finance_infra_proto_rawDescOnce sync.Once
	file_finance_infra_proto_rawDescData = file_finance_infra_proto_rawDesc
)

func file_finance_infra_proto_rawDescGZIP() []byte {
	file_finance_infra_proto_rawDescOnce.Do(func() {
		file_finance_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_finance_infra_proto_rawDescData)
	})
	return file_finance_infra_proto_rawDescData
}

var file_finance_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_finance_infra_proto_goTypes = []interface{}{
	(*ComputeStartIn)(nil),       // 0: ComputeStartIn
	(*ComputeEndIn)(nil),         // 1: ComputeEndIn
	(*LambdaStartIn)(nil),        // 2: LambdaStartIn
	(*LambdaEndIn)(nil),          // 3: LambdaEndIn
	(*BlockStorageStartIn)(nil),  // 4: BlockStorageStartIn
	(*BlockStorageEndIn)(nil),    // 5: BlockStorageEndIn
	(*ObjectStorageStartIn)(nil), // 6: ObjectStorageStartIn
	(*ObjectStorageEndIn)(nil),   // 7: ObjectStorageEndIn
	(*CIStartIn)(nil),            // 8: CIStartIn
	(*CIEndIn)(nil),              // 9: CIEndIn
	(*FinanceInfraVoid)(nil),     // 10: FinanceInfraVoid
}
var file_finance_infra_proto_depIdxs = []int32{
	0,  // 0: FinanceInfra.ComputeStart:input_type -> ComputeStartIn
	1,  // 1: FinanceInfra.ComputeEnd:input_type -> ComputeEndIn
	2,  // 2: FinanceInfra.LambdaStart:input_type -> LambdaStartIn
	3,  // 3: FinanceInfra.LambdaEnd:input_type -> LambdaEndIn
	4,  // 4: FinanceInfra.BlockStorageStart:input_type -> BlockStorageStartIn
	5,  // 5: FinanceInfra.BlockStorageEnd:input_type -> BlockStorageEndIn
	6,  // 6: FinanceInfra.ObjectStorageStart:input_type -> ObjectStorageStartIn
	7,  // 7: FinanceInfra.ObjectStorageEnd:input_type -> ObjectStorageEndIn
	8,  // 8: FinanceInfra.CIStart:input_type -> CIStartIn
	9,  // 9: FinanceInfra.CIEnd:input_type -> CIEndIn
	10, // 10: FinanceInfra.ComputeStart:output_type -> FinanceInfraVoid
	10, // 11: FinanceInfra.ComputeEnd:output_type -> FinanceInfraVoid
	10, // 12: FinanceInfra.LambdaStart:output_type -> FinanceInfraVoid
	10, // 13: FinanceInfra.LambdaEnd:output_type -> FinanceInfraVoid
	10, // 14: FinanceInfra.BlockStorageStart:output_type -> FinanceInfraVoid
	10, // 15: FinanceInfra.BlockStorageEnd:output_type -> FinanceInfraVoid
	10, // 16: FinanceInfra.ObjectStorageStart:output_type -> FinanceInfraVoid
	10, // 17: FinanceInfra.ObjectStorageEnd:output_type -> FinanceInfraVoid
	10, // 18: FinanceInfra.CIStart:output_type -> FinanceInfraVoid
	10, // 19: FinanceInfra.CIEnd:output_type -> FinanceInfraVoid
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_finance_infra_proto_init() }
func file_finance_infra_proto_init() {
	if File_finance_infra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_finance_infra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeStartIn); i {
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
		file_finance_infra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeEndIn); i {
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
		file_finance_infra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaStartIn); i {
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
		file_finance_infra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaEndIn); i {
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
		file_finance_infra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStorageStartIn); i {
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
		file_finance_infra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStorageEndIn); i {
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
		file_finance_infra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectStorageStartIn); i {
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
		file_finance_infra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectStorageEndIn); i {
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
		file_finance_infra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIStartIn); i {
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
		file_finance_infra_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIEndIn); i {
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
		file_finance_infra_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinanceInfraVoid); i {
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
			RawDescriptor: file_finance_infra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_finance_infra_proto_goTypes,
		DependencyIndexes: file_finance_infra_proto_depIdxs,
		MessageInfos:      file_finance_infra_proto_msgTypes,
	}.Build()
	File_finance_infra_proto = out.File
	file_finance_infra_proto_rawDesc = nil
	file_finance_infra_proto_goTypes = nil
	file_finance_infra_proto_depIdxs = nil
}