// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto/hotel.proto

package hotel

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

type AddHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ContactNo   string `protobuf:"bytes,3,opt,name=contact_no,json=contactNo,proto3" json:"contact_no,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Website     string `protobuf:"bytes,5,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AddHotelRequest) Reset() {
	*x = AddHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHotelRequest) ProtoMessage() {}

func (x *AddHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHotelRequest.ProtoReflect.Descriptor instead.
func (*AddHotelRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *AddHotelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddHotelRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddHotelRequest) GetContactNo() string {
	if x != nil {
		return x.ContactNo
	}
	return ""
}

func (x *AddHotelRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddHotelRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *AddHotelRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddHotelRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AddHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddHotelResponse) Reset() {
	*x = AddHotelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHotelResponse) ProtoMessage() {}

func (x *AddHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHotelResponse.ProtoReflect.Descriptor instead.
func (*AddHotelResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *AddHotelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddHotelResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHotelRequest) Reset() {
	*x = GetHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelRequest) ProtoMessage() {}

func (x *GetHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelRequest.ProtoReflect.Descriptor instead.
func (*GetHotelRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *GetHotelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ContactNo   string `protobuf:"bytes,4,opt,name=contact_no,json=contactNo,proto3" json:"contact_no,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Website     string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetHotelResponse) Reset() {
	*x = GetHotelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelResponse) ProtoMessage() {}

func (x *GetHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelResponse.ProtoReflect.Descriptor instead.
func (*GetHotelResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{3}
}

func (x *GetHotelResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetHotelResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetHotelResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetHotelResponse) GetContactNo() string {
	if x != nil {
		return x.ContactNo
	}
	return ""
}

func (x *GetHotelResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetHotelResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetHotelResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetHotelResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ListHotelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ListHotelsRequest) Reset() {
	*x = ListHotelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHotelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHotelsRequest) ProtoMessage() {}

func (x *ListHotelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHotelsRequest.ProtoReflect.Descriptor instead.
func (*ListHotelsRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{4}
}

func (x *ListHotelsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ListHotelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels,proto3" json:"hotels,omitempty"`
}

func (x *ListHotelsResponse) Reset() {
	*x = ListHotelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHotelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHotelsResponse) ProtoMessage() {}

func (x *ListHotelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHotelsResponse.ProtoReflect.Descriptor instead.
func (*ListHotelsResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{5}
}

func (x *ListHotelsResponse) GetHotels() []*Hotel {
	if x != nil {
		return x.Hotels
	}
	return nil
}

type Hotel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ContactNo   string `protobuf:"bytes,4,opt,name=contact_no,json=contactNo,proto3" json:"contact_no,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Website     string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Hotel) Reset() {
	*x = Hotel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hotel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotel) ProtoMessage() {}

func (x *Hotel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotel.ProtoReflect.Descriptor instead.
func (*Hotel) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{6}
}

func (x *Hotel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Hotel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hotel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Hotel) GetContactNo() string {
	if x != nil {
		return x.ContactNo
	}
	return ""
}

func (x *Hotel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Hotel) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Hotel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Hotel) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UpdateHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ContactNo   string `protobuf:"bytes,4,opt,name=contact_no,json=contactNo,proto3" json:"contact_no,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Website     string `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Username    string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UpdateHotelRequest) Reset() {
	*x = UpdateHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHotelRequest) ProtoMessage() {}

func (x *UpdateHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHotelRequest.ProtoReflect.Descriptor instead.
func (*UpdateHotelRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateHotelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateHotelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHotelRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateHotelRequest) GetContactNo() string {
	if x != nil {
		return x.ContactNo
	}
	return ""
}

func (x *UpdateHotelRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateHotelRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *UpdateHotelRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateHotelRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UpdateHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateHotelResponse) Reset() {
	*x = UpdateHotelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHotelResponse) ProtoMessage() {}

func (x *UpdateHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHotelResponse.ProtoReflect.Descriptor instead.
func (*UpdateHotelResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateHotelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateHotelResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DeleteHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHotelRequest) Reset() {
	*x = DeleteHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHotelRequest) ProtoMessage() {}

func (x *DeleteHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHotelRequest.ProtoReflect.Descriptor instead.
func (*DeleteHotelRequest) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteHotelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteHotelResponse) Reset() {
	*x = DeleteHotelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hotel_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHotelResponse) ProtoMessage() {}

func (x *DeleteHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hotel_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHotelResponse.ProtoReflect.Descriptor instead.
func (*DeleteHotelResponse) Descriptor() ([]byte, []int) {
	return file_proto_hotel_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteHotelResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteHotelResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_proto_hotel_proto protoreflect.FileDescriptor

var file_proto_hotel_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0xcc, 0x01, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x06, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x22,
	0xd2, 0x01, 0x0a, 0x05, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd6,
	0x02, 0x0a, 0x0c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x16, 0x2e, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x16, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x19, 0x2e, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x12, 0x19, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_hotel_proto_rawDescOnce sync.Once
	file_proto_hotel_proto_rawDescData = file_proto_hotel_proto_rawDesc
)

func file_proto_hotel_proto_rawDescGZIP() []byte {
	file_proto_hotel_proto_rawDescOnce.Do(func() {
		file_proto_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hotel_proto_rawDescData)
	})
	return file_proto_hotel_proto_rawDescData
}

var file_proto_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_hotel_proto_goTypes = []any{
	(*AddHotelRequest)(nil),     // 0: hotel.AddHotelRequest
	(*AddHotelResponse)(nil),    // 1: hotel.AddHotelResponse
	(*GetHotelRequest)(nil),     // 2: hotel.GetHotelRequest
	(*GetHotelResponse)(nil),    // 3: hotel.GetHotelResponse
	(*ListHotelsRequest)(nil),   // 4: hotel.ListHotelsRequest
	(*ListHotelsResponse)(nil),  // 5: hotel.ListHotelsResponse
	(*Hotel)(nil),               // 6: hotel.Hotel
	(*UpdateHotelRequest)(nil),  // 7: hotel.UpdateHotelRequest
	(*UpdateHotelResponse)(nil), // 8: hotel.UpdateHotelResponse
	(*DeleteHotelRequest)(nil),  // 9: hotel.DeleteHotelRequest
	(*DeleteHotelResponse)(nil), // 10: hotel.DeleteHotelResponse
}
var file_proto_hotel_proto_depIdxs = []int32{
	6,  // 0: hotel.ListHotelsResponse.hotels:type_name -> hotel.Hotel
	0,  // 1: hotel.HotelService.AddHotel:input_type -> hotel.AddHotelRequest
	2,  // 2: hotel.HotelService.GetHotel:input_type -> hotel.GetHotelRequest
	4,  // 3: hotel.HotelService.GetHotels:input_type -> hotel.ListHotelsRequest
	7,  // 4: hotel.HotelService.UpdateHotel:input_type -> hotel.UpdateHotelRequest
	9,  // 5: hotel.HotelService.DeleteHotel:input_type -> hotel.DeleteHotelRequest
	1,  // 6: hotel.HotelService.AddHotel:output_type -> hotel.AddHotelResponse
	3,  // 7: hotel.HotelService.GetHotel:output_type -> hotel.GetHotelResponse
	5,  // 8: hotel.HotelService.GetHotels:output_type -> hotel.ListHotelsResponse
	8,  // 9: hotel.HotelService.UpdateHotel:output_type -> hotel.UpdateHotelResponse
	10, // 10: hotel.HotelService.DeleteHotel:output_type -> hotel.DeleteHotelResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_hotel_proto_init() }
func file_proto_hotel_proto_init() {
	if File_proto_hotel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hotel_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddHotelRequest); i {
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
		file_proto_hotel_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddHotelResponse); i {
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
		file_proto_hotel_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetHotelRequest); i {
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
		file_proto_hotel_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetHotelResponse); i {
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
		file_proto_hotel_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListHotelsRequest); i {
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
		file_proto_hotel_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListHotelsResponse); i {
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
		file_proto_hotel_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Hotel); i {
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
		file_proto_hotel_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateHotelRequest); i {
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
		file_proto_hotel_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateHotelResponse); i {
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
		file_proto_hotel_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHotelRequest); i {
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
		file_proto_hotel_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHotelResponse); i {
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
			RawDescriptor: file_proto_hotel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hotel_proto_goTypes,
		DependencyIndexes: file_proto_hotel_proto_depIdxs,
		MessageInfos:      file_proto_hotel_proto_msgTypes,
	}.Build()
	File_proto_hotel_proto = out.File
	file_proto_hotel_proto_rawDesc = nil
	file_proto_hotel_proto_goTypes = nil
	file_proto_hotel_proto_depIdxs = nil
}
