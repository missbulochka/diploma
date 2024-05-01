// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: piglet-bills/api/proto/bills.proto

package billsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BillType       bool                   `protobuf:"varint,2,opt,name=billType,proto3" json:"billType,omitempty"`
	BillStatus     bool                   `protobuf:"varint,3,opt,name=billStatus,proto3" json:"billStatus,omitempty"`
	BillName       string                 `protobuf:"bytes,4,opt,name=billName,proto3" json:"billName,omitempty"`
	CurrentSum     float32                `protobuf:"fixed32,5,opt,name=currentSum,proto3" json:"currentSum,omitempty"`
	Date           *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	MonthlyPayment float32                `protobuf:"fixed32,8,opt,name=monthlyPayment,proto3" json:"monthlyPayment,omitempty"`
}

func (x *Bill) Reset() {
	*x = Bill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bill) ProtoMessage() {}

func (x *Bill) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bill.ProtoReflect.Descriptor instead.
func (*Bill) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{0}
}

func (x *Bill) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bill) GetBillType() bool {
	if x != nil {
		return x.BillType
	}
	return false
}

func (x *Bill) GetBillStatus() bool {
	if x != nil {
		return x.BillStatus
	}
	return false
}

func (x *Bill) GetBillName() string {
	if x != nil {
		return x.BillName
	}
	return ""
}

func (x *Bill) GetCurrentSum() float32 {
	if x != nil {
		return x.CurrentSum
	}
	return 0
}

func (x *Bill) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Bill) GetMonthlyPayment() float32 {
	if x != nil {
		return x.MonthlyPayment
	}
	return 0
}

type CreateBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bill *Bill `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
}

func (x *CreateBillRequest) Reset() {
	*x = CreateBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBillRequest) ProtoMessage() {}

func (x *CreateBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBillRequest.ProtoReflect.Descriptor instead.
func (*CreateBillRequest) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBillRequest) GetBill() *Bill {
	if x != nil {
		return x.Bill
	}
	return nil
}

type CreateBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bill *Bill `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
}

func (x *CreateBillResponse) Reset() {
	*x = CreateBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBillResponse) ProtoMessage() {}

func (x *CreateBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBillResponse.ProtoReflect.Descriptor instead.
func (*CreateBillResponse) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBillResponse) GetBill() *Bill {
	if x != nil {
		return x.Bill
	}
	return nil
}

type GetSomeBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSomeBillsRequest) Reset() {
	*x = GetSomeBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSomeBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSomeBillsRequest) ProtoMessage() {}

func (x *GetSomeBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSomeBillsRequest.ProtoReflect.Descriptor instead.
func (*GetSomeBillsRequest) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{3}
}

type GetSomeBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bills []*Bill `protobuf:"bytes,1,rep,name=bills,proto3" json:"bills,omitempty"`
}

func (x *GetSomeBillsResponse) Reset() {
	*x = GetSomeBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSomeBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSomeBillsResponse) ProtoMessage() {}

func (x *GetSomeBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSomeBillsResponse.ProtoReflect.Descriptor instead.
func (*GetSomeBillsResponse) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{4}
}

func (x *GetSomeBillsResponse) GetBills() []*Bill {
	if x != nil {
		return x.Bills
	}
	return nil
}

type GetBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBillRequest) Reset() {
	*x = GetBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillRequest) ProtoMessage() {}

func (x *GetBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillRequest.ProtoReflect.Descriptor instead.
func (*GetBillRequest) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{5}
}

func (x *GetBillRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bill *Bill `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
}

func (x *GetBillResponse) Reset() {
	*x = GetBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillResponse) ProtoMessage() {}

func (x *GetBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillResponse.ProtoReflect.Descriptor instead.
func (*GetBillResponse) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{6}
}

func (x *GetBillResponse) GetBill() *Bill {
	if x != nil {
		return x.Bill
	}
	return nil
}

type UpdateBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bill *Bill `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
}

func (x *UpdateBillRequest) Reset() {
	*x = UpdateBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillRequest) ProtoMessage() {}

func (x *UpdateBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillRequest.ProtoReflect.Descriptor instead.
func (*UpdateBillRequest) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBillRequest) GetBill() *Bill {
	if x != nil {
		return x.Bill
	}
	return nil
}

type UpdateBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bill *Bill `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
}

func (x *UpdateBillResponse) Reset() {
	*x = UpdateBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBillResponse) ProtoMessage() {}

func (x *UpdateBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBillResponse.ProtoReflect.Descriptor instead.
func (*UpdateBillResponse) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateBillResponse) GetBill() *Bill {
	if x != nil {
		return x.Bill
	}
	return nil
}

type DeleteBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBillRequest) Reset() {
	*x = DeleteBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillRequest) ProtoMessage() {}

func (x *DeleteBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBillRequest.ProtoReflect.Descriptor instead.
func (*DeleteBillRequest) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteBillRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteBillResponse) Reset() {
	*x = DeleteBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillResponse) ProtoMessage() {}

func (x *DeleteBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_piglet_bills_api_proto_bills_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBillResponse.ProtoReflect.Descriptor instead.
func (*DeleteBillResponse) Descriptor() ([]byte, []int) {
	return file_piglet_bills_api_proto_bills_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBillResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_piglet_bills_api_proto_bills_proto protoreflect.FileDescriptor

var file_piglet_bills_api_proto_bills_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x69, 0x67, 0x6c, 0x65, 0x74, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a,
	0x04, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x6d, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x34, 0x0a, 0x11, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x62, 0x69, 0x6c,
	0x6c, 0x22, 0x35, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x53,
	0x6f, 0x6d, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x39, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x69, 0x6c, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x52, 0x05, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x67, 0x65,
	0x74, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x62, 0x69, 0x6c, 0x6c,
	0x22, 0x34, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x52, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x22, 0x35, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04,
	0x62, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x62, 0x69, 0x6c, 0x6c, 0x22, 0x23, 0x0a,
	0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0xd9, 0x02, 0x0a, 0x0b, 0x70, 0x69, 0x67, 0x6c, 0x65, 0x74, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c,
	0x12, 0x18, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x53, 0x6f, 0x6d, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1a, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x67, 0x65,
	0x74, 0x53, 0x6f, 0x6d, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x53, 0x6f, 0x6d,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24,
	0x5a, 0x22, 0x70, 0x69, 0x67, 0x6c, 0x65, 0x74, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x62, 0x69, 0x6c,
	0x6c, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_piglet_bills_api_proto_bills_proto_rawDescOnce sync.Once
	file_piglet_bills_api_proto_bills_proto_rawDescData = file_piglet_bills_api_proto_bills_proto_rawDesc
)

func file_piglet_bills_api_proto_bills_proto_rawDescGZIP() []byte {
	file_piglet_bills_api_proto_bills_proto_rawDescOnce.Do(func() {
		file_piglet_bills_api_proto_bills_proto_rawDescData = protoimpl.X.CompressGZIP(file_piglet_bills_api_proto_bills_proto_rawDescData)
	})
	return file_piglet_bills_api_proto_bills_proto_rawDescData
}

var file_piglet_bills_api_proto_bills_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_piglet_bills_api_proto_bills_proto_goTypes = []interface{}{
	(*Bill)(nil),                  // 0: bills.Bill
	(*CreateBillRequest)(nil),     // 1: bills.createBillRequest
	(*CreateBillResponse)(nil),    // 2: bills.createBillResponse
	(*GetSomeBillsRequest)(nil),   // 3: bills.getSomeBillsRequest
	(*GetSomeBillsResponse)(nil),  // 4: bills.getSomeBillsResponse
	(*GetBillRequest)(nil),        // 5: bills.getBillRequest
	(*GetBillResponse)(nil),       // 6: bills.getBillResponse
	(*UpdateBillRequest)(nil),     // 7: bills.updateBillRequest
	(*UpdateBillResponse)(nil),    // 8: bills.updateBillResponse
	(*DeleteBillRequest)(nil),     // 9: bills.deleteBillRequest
	(*DeleteBillResponse)(nil),    // 10: bills.deleteBillResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_piglet_bills_api_proto_bills_proto_depIdxs = []int32{
	11, // 0: bills.Bill.date:type_name -> google.protobuf.Timestamp
	0,  // 1: bills.createBillRequest.bill:type_name -> bills.Bill
	0,  // 2: bills.createBillResponse.bill:type_name -> bills.Bill
	0,  // 3: bills.getSomeBillsResponse.bills:type_name -> bills.Bill
	0,  // 4: bills.getBillResponse.bill:type_name -> bills.Bill
	0,  // 5: bills.updateBillRequest.bill:type_name -> bills.Bill
	0,  // 6: bills.updateBillResponse.bill:type_name -> bills.Bill
	1,  // 7: bills.pigletBills.createBill:input_type -> bills.createBillRequest
	3,  // 8: bills.pigletBills.getSomeBills:input_type -> bills.getSomeBillsRequest
	5,  // 9: bills.pigletBills.getBill:input_type -> bills.getBillRequest
	7,  // 10: bills.pigletBills.updateBill:input_type -> bills.updateBillRequest
	9,  // 11: bills.pigletBills.deleteBill:input_type -> bills.deleteBillRequest
	2,  // 12: bills.pigletBills.createBill:output_type -> bills.createBillResponse
	4,  // 13: bills.pigletBills.getSomeBills:output_type -> bills.getSomeBillsResponse
	6,  // 14: bills.pigletBills.getBill:output_type -> bills.getBillResponse
	8,  // 15: bills.pigletBills.updateBill:output_type -> bills.updateBillResponse
	10, // 16: bills.pigletBills.deleteBill:output_type -> bills.deleteBillResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_piglet_bills_api_proto_bills_proto_init() }
func file_piglet_bills_api_proto_bills_proto_init() {
	if File_piglet_bills_api_proto_bills_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_piglet_bills_api_proto_bills_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bill); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBillRequest); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBillResponse); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSomeBillsRequest); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSomeBillsResponse); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBillRequest); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBillResponse); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillRequest); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBillResponse); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBillRequest); i {
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
		file_piglet_bills_api_proto_bills_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBillResponse); i {
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
			RawDescriptor: file_piglet_bills_api_proto_bills_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_piglet_bills_api_proto_bills_proto_goTypes,
		DependencyIndexes: file_piglet_bills_api_proto_bills_proto_depIdxs,
		MessageInfos:      file_piglet_bills_api_proto_bills_proto_msgTypes,
	}.Build()
	File_piglet_bills_api_proto_bills_proto = out.File
	file_piglet_bills_api_proto_bills_proto_rawDesc = nil
	file_piglet_bills_api_proto_bills_proto_goTypes = nil
	file_piglet_bills_api_proto_bills_proto_depIdxs = nil
}