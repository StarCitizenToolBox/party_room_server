// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: index.proto

package protos

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

type RoomStatus int32

const (
	RoomStatus_All         RoomStatus = 0
	RoomStatus_Open        RoomStatus = 1
	RoomStatus_Private     RoomStatus = 2
	RoomStatus_Full        RoomStatus = 3
	RoomStatus_Closed      RoomStatus = 4
	RoomStatus_WillOffline RoomStatus = 5
	RoomStatus_Offline     RoomStatus = 6
)

// Enum value maps for RoomStatus.
var (
	RoomStatus_name = map[int32]string{
		0: "All",
		1: "Open",
		2: "Private",
		3: "Full",
		4: "Closed",
		5: "WillOffline",
		6: "Offline",
	}
	RoomStatus_value = map[string]int32{
		"All":         0,
		"Open":        1,
		"Private":     2,
		"Full":        3,
		"Closed":      4,
		"WillOffline": 5,
		"Offline":     6,
	}
)

func (x RoomStatus) Enum() *RoomStatus {
	p := new(RoomStatus)
	*p = x
	return p
}

func (x RoomStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_index_proto_enumTypes[0].Descriptor()
}

func (RoomStatus) Type() protoreflect.EnumType {
	return &file_index_proto_enumTypes[0]
}

func (x RoomStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomStatus.Descriptor instead.
func (RoomStatus) EnumDescriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{0}
}

type RoomSortType int32

const (
	RoomSortType_Default             RoomSortType = 0
	RoomSortType_MostPlayerNumber    RoomSortType = 1
	RoomSortType_MinimumPlayerNumber RoomSortType = 2
	RoomSortType_RecentlyCreated     RoomSortType = 3
	RoomSortType_OldestCreated       RoomSortType = 4
)

// Enum value maps for RoomSortType.
var (
	RoomSortType_name = map[int32]string{
		0: "Default",
		1: "MostPlayerNumber",
		2: "MinimumPlayerNumber",
		3: "RecentlyCreated",
		4: "OldestCreated",
	}
	RoomSortType_value = map[string]int32{
		"Default":             0,
		"MostPlayerNumber":    1,
		"MinimumPlayerNumber": 2,
		"RecentlyCreated":     3,
		"OldestCreated":       4,
	}
)

func (x RoomSortType) Enum() *RoomSortType {
	p := new(RoomSortType)
	*p = x
	return p
}

func (x RoomSortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomSortType) Descriptor() protoreflect.EnumDescriptor {
	return file_index_proto_enumTypes[1].Descriptor()
}

func (RoomSortType) Type() protoreflect.EnumType {
	return &file_index_proto_enumTypes[1]
}

func (x RoomSortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomSortType.Descriptor instead.
func (RoomSortType) EnumDescriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{0}
}

type BaseRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BaseRespData) Reset() {
	*x = BaseRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseRespData) ProtoMessage() {}

func (x *BaseRespData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseRespData.ProtoReflect.Descriptor instead.
func (*BaseRespData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{1}
}

func (x *BaseRespData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseRespData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BasePageRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	HasNext    bool   `protobuf:"varint,3,opt,name=hasNext,proto3" json:"hasNext,omitempty"`
	CurPageNum uint64 `protobuf:"varint,4,opt,name=curPageNum,proto3" json:"curPageNum,omitempty"`
	PageSize   int64  `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *BasePageRespData) Reset() {
	*x = BasePageRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasePageRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasePageRespData) ProtoMessage() {}

func (x *BasePageRespData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasePageRespData.ProtoReflect.Descriptor instead.
func (*BasePageRespData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{2}
}

func (x *BasePageRespData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BasePageRespData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BasePageRespData) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

func (x *BasePageRespData) GetCurPageNum() uint64 {
	if x != nil {
		return x.CurPageNum
	}
	return 0
}

func (x *BasePageRespData) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data          string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ClientVersion int64  `protobuf:"zigzag64,2,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	ServerVersion int64  `protobuf:"zigzag64,3,opt,name=serverVersion,proto3" json:"serverVersion,omitempty"`
}

func (x *PingData) Reset() {
	*x = PingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingData) ProtoMessage() {}

func (x *PingData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingData.ProtoReflect.Descriptor instead.
func (*PingData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{3}
}

func (x *PingData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *PingData) GetClientVersion() int64 {
	if x != nil {
		return x.ClientVersion
	}
	return 0
}

func (x *PingData) GetServerVersion() int64 {
	if x != nil {
		return x.ServerVersion
	}
	return 0
}

type RoomTypesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTypes []*RoomType `protobuf:"bytes,1,rep,name=roomTypes,proto3" json:"roomTypes,omitempty"`
}

func (x *RoomTypesData) Reset() {
	*x = RoomTypesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomTypesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTypesData) ProtoMessage() {}

func (x *RoomTypesData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTypesData.ProtoReflect.Descriptor instead.
func (*RoomTypesData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{4}
}

func (x *RoomTypesData) GetRoomTypes() []*RoomType {
	if x != nil {
		return x.RoomTypes
	}
	return nil
}

type RoomType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string         `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Desc     string         `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	SubTypes []*RoomSubtype `protobuf:"bytes,5,rep,name=subTypes,proto3" json:"subTypes,omitempty"`
}

func (x *RoomType) Reset() {
	*x = RoomType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomType) ProtoMessage() {}

func (x *RoomType) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomType.ProtoReflect.Descriptor instead.
func (*RoomType) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{5}
}

func (x *RoomType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomType) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *RoomType) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *RoomType) GetSubTypes() []*RoomSubtype {
	if x != nil {
		return x.SubTypes
	}
	return nil
}

type RoomSubtype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoomSubtype) Reset() {
	*x = RoomSubtype{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomSubtype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomSubtype) ProtoMessage() {}

func (x *RoomSubtype) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomSubtype.ProtoReflect.Descriptor instead.
func (*RoomSubtype) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{6}
}

func (x *RoomSubtype) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomSubtype) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoomData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomTypeID     string     `protobuf:"bytes,2,opt,name=roomTypeID,proto3" json:"roomTypeID,omitempty"`
	RoomSubTypeIds []string   `protobuf:"bytes,3,rep,name=roomSubTypeIds,proto3" json:"roomSubTypeIds,omitempty"`
	Owner          string     `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	MaxPlayer      int32      `protobuf:"varint,5,opt,name=maxPlayer,proto3" json:"maxPlayer,omitempty"`
	CreateTime     int64      `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CurPlayer      int32      `protobuf:"varint,7,opt,name=curPlayer,proto3" json:"curPlayer,omitempty"`
	Status         RoomStatus `protobuf:"varint,8,opt,name=status,proto3,enum=RoomStatus" json:"status,omitempty"`
	DeviceUUID     string     `protobuf:"bytes,9,opt,name=deviceUUID,proto3" json:"deviceUUID,omitempty"`
}

func (x *RoomData) Reset() {
	*x = RoomData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomData) ProtoMessage() {}

func (x *RoomData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomData.ProtoReflect.Descriptor instead.
func (*RoomData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{7}
}

func (x *RoomData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomData) GetRoomTypeID() string {
	if x != nil {
		return x.RoomTypeID
	}
	return ""
}

func (x *RoomData) GetRoomSubTypeIds() []string {
	if x != nil {
		return x.RoomSubTypeIds
	}
	return nil
}

func (x *RoomData) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RoomData) GetMaxPlayer() int32 {
	if x != nil {
		return x.MaxPlayer
	}
	return 0
}

func (x *RoomData) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *RoomData) GetCurPlayer() int32 {
	if x != nil {
		return x.CurPlayer
	}
	return 0
}

func (x *RoomData) GetStatus() RoomStatus {
	if x != nil {
		return x.Status
	}
	return RoomStatus_All
}

func (x *RoomData) GetDeviceUUID() string {
	if x != nil {
		return x.DeviceUUID
	}
	return ""
}

type RoomListPageReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeID    string       `protobuf:"bytes,1,opt,name=typeID,proto3" json:"typeID,omitempty"`
	SubTypeID string       `protobuf:"bytes,2,opt,name=subTypeID,proto3" json:"subTypeID,omitempty"`
	Status    RoomStatus   `protobuf:"varint,3,opt,name=status,proto3,enum=RoomStatus" json:"status,omitempty"`
	Sort      RoomSortType `protobuf:"varint,4,opt,name=sort,proto3,enum=RoomSortType" json:"sort,omitempty"`
	PageNum   uint64       `protobuf:"varint,5,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
}

func (x *RoomListPageReqData) Reset() {
	*x = RoomListPageReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListPageReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListPageReqData) ProtoMessage() {}

func (x *RoomListPageReqData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListPageReqData.ProtoReflect.Descriptor instead.
func (*RoomListPageReqData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{8}
}

func (x *RoomListPageReqData) GetTypeID() string {
	if x != nil {
		return x.TypeID
	}
	return ""
}

func (x *RoomListPageReqData) GetSubTypeID() string {
	if x != nil {
		return x.SubTypeID
	}
	return ""
}

func (x *RoomListPageReqData) GetStatus() RoomStatus {
	if x != nil {
		return x.Status
	}
	return RoomStatus_All
}

func (x *RoomListPageReqData) GetSort() RoomSortType {
	if x != nil {
		return x.Sort
	}
	return RoomSortType_Default
}

func (x *RoomListPageReqData) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type RoomListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageData *BasePageRespData `protobuf:"bytes,1,opt,name=pageData,proto3" json:"pageData,omitempty"`
	Rooms    []*RoomData       `protobuf:"bytes,2,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *RoomListData) Reset() {
	*x = RoomListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListData) ProtoMessage() {}

func (x *RoomListData) ProtoReflect() protoreflect.Message {
	mi := &file_index_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListData.ProtoReflect.Descriptor instead.
func (*RoomListData) Descriptor() ([]byte, []int) {
	return file_index_proto_rawDescGZIP(), []int{9}
}

func (x *RoomListData) GetPageData() *BasePageRespData {
	if x != nil {
		return x.PageData
	}
	return nil
}

func (x *RoomListData) GetRooms() []*RoomData {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_index_proto protoreflect.FileDescriptor

var file_index_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6a, 0x0a,
	0x08, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0d, 0x52, 0x6f, 0x6f,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x28, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x75,
	0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x08, 0x52, 0x6f,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x75, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x55, 0x55, 0x49, 0x44, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x5e, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2a, 0x60, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4f, 0x70, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x75, 0x6c, 0x6c, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x69, 0x6c,
	0x6c, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x06, 0x2a, 0x72, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x6d, 0x53,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x6f, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x6c, 0x64, 0x65,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x04, 0x32, 0xba, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0a,
	0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x09, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x09, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x09, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_proto_rawDescOnce sync.Once
	file_index_proto_rawDescData = file_index_proto_rawDesc
)

func file_index_proto_rawDescGZIP() []byte {
	file_index_proto_rawDescOnce.Do(func() {
		file_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_proto_rawDescData)
	})
	return file_index_proto_rawDescData
}

var file_index_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_index_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_index_proto_goTypes = []interface{}{
	(RoomStatus)(0),             // 0: RoomStatus
	(RoomSortType)(0),           // 1: RoomSortType
	(*Empty)(nil),               // 2: Empty
	(*BaseRespData)(nil),        // 3: BaseRespData
	(*BasePageRespData)(nil),    // 4: BasePageRespData
	(*PingData)(nil),            // 5: PingData
	(*RoomTypesData)(nil),       // 6: RoomTypesData
	(*RoomType)(nil),            // 7: RoomType
	(*RoomSubtype)(nil),         // 8: RoomSubtype
	(*RoomData)(nil),            // 9: RoomData
	(*RoomListPageReqData)(nil), // 10: RoomListPageReqData
	(*RoomListData)(nil),        // 11: RoomListData
}
var file_index_proto_depIdxs = []int32{
	7,  // 0: RoomTypesData.roomTypes:type_name -> RoomType
	8,  // 1: RoomType.subTypes:type_name -> RoomSubtype
	0,  // 2: RoomData.status:type_name -> RoomStatus
	0,  // 3: RoomListPageReqData.status:type_name -> RoomStatus
	1,  // 4: RoomListPageReqData.sort:type_name -> RoomSortType
	4,  // 5: RoomListData.pageData:type_name -> BasePageRespData
	9,  // 6: RoomListData.rooms:type_name -> RoomData
	5,  // 7: IndexService.PingServer:input_type -> PingData
	2,  // 8: IndexService.GetRoomTypes:input_type -> Empty
	9,  // 9: IndexService.CreateRoom:input_type -> RoomData
	10, // 10: IndexService.GetRoomList:input_type -> RoomListPageReqData
	5,  // 11: IndexService.PingServer:output_type -> PingData
	6,  // 12: IndexService.GetRoomTypes:output_type -> RoomTypesData
	9,  // 13: IndexService.CreateRoom:output_type -> RoomData
	11, // 14: IndexService.GetRoomList:output_type -> RoomListData
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_index_proto_init() }
func file_index_proto_init() {
	if File_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseRespData); i {
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
		file_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasePageRespData); i {
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
		file_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingData); i {
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
		file_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomTypesData); i {
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
		file_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomType); i {
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
		file_index_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomSubtype); i {
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
		file_index_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomData); i {
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
		file_index_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListPageReqData); i {
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
		file_index_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListData); i {
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
			RawDescriptor: file_index_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_proto_goTypes,
		DependencyIndexes: file_index_proto_depIdxs,
		EnumInfos:         file_index_proto_enumTypes,
		MessageInfos:      file_index_proto_msgTypes,
	}.Build()
	File_index_proto = out.File
	file_index_proto_rawDesc = nil
	file_index_proto_goTypes = nil
	file_index_proto_depIdxs = nil
}
