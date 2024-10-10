// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: user.proto

package api

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

// 空返回结构
type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

// 地理位置结构
type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                        // 类型
	Coordinates []float64 `protobuf:"fixed64,2,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"` // 坐标点
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPoint.ProtoReflect.Descriptor instead.
func (*GeoPoint) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GeoPoint) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GeoPoint) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                     // 姓名
	Age         int32     `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`                                      // 年龄
	Gender      int32     `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`                                // 性别
	Location    *GeoPoint `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`                             // GPS位置
	MatchGender int32     `protobuf:"varint,5,opt,name=match_gender,json=matchGender,proto3" json:"match_gender,omitempty"`   // 性取向
	MatchMinAge int32     `protobuf:"varint,6,opt,name=match_min_age,json=matchMinAge,proto3" json:"match_min_age,omitempty"` // 年龄偏好范围(min)
	MatchMaxAge int32     `protobuf:"varint,7,opt,name=match_max_age,json=matchMaxAge,proto3" json:"match_max_age,omitempty"` // 年龄偏好范围(max)
}

func (x *UserData) Reset() {
	*x = UserData{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserData) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserData) GetLocation() *GeoPoint {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *UserData) GetMatchGender() int32 {
	if x != nil {
		return x.MatchGender
	}
	return 0
}

func (x *UserData) GetMatchMinAge() int32 {
	if x != nil {
		return x.MatchMinAge
	}
	return 0
}

func (x *UserData) GetMatchMaxAge() int32 {
	if x != nil {
		return x.MatchMaxAge
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // 用户数据结构
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinAge      int32 `protobuf:"varint,1,opt,name=min_age,json=minAge,proto3" json:"min_age,omitempty"`                  // 年龄范围(min)
	MaxAge      int32 `protobuf:"varint,2,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`                  // 年龄范围(max)
	Gender      int32 `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`                                // 性别
	MatchGender int32 `protobuf:"varint,4,opt,name=match_gender,json=matchGender,proto3" json:"match_gender,omitempty"`   // 性取向
	MatchMinAge int32 `protobuf:"varint,5,opt,name=match_min_age,json=matchMinAge,proto3" json:"match_min_age,omitempty"` // 年龄偏好范围(min)
	MatchMaxAge int32 `protobuf:"varint,6,opt,name=match_max_age,json=matchMaxAge,proto3" json:"match_max_age,omitempty"` // 年龄偏好范围(max)
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *SearchRequest) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *SearchRequest) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *SearchRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SearchRequest) GetMatchGender() int32 {
	if x != nil {
		return x.MatchGender
	}
	return 0
}

func (x *SearchRequest) GetMatchMinAge() int32 {
	if x != nil {
		return x.MatchMinAge
	}
	return 0
}

func (x *SearchRequest) GetMatchMaxAge() int32 {
	if x != nil {
		return x.MatchMaxAge
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserData `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // 查询匹配到的用户列表
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetUsers() []*UserData {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x69, 0x6e, 0x41, 0x67,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d,
	0x61, 0x78, 0x41, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x69, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x69,
	0x6e, 0x41, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65,
	0x22, 0x35, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x6f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []any{
	(*EmptyResponse)(nil),  // 0: api.EmptyResponse
	(*GeoPoint)(nil),       // 1: api.GeoPoint
	(*UserData)(nil),       // 2: api.UserData
	(*CreateRequest)(nil),  // 3: api.CreateRequest
	(*SearchRequest)(nil),  // 4: api.SearchRequest
	(*SearchResponse)(nil), // 5: api.SearchResponse
}
var file_user_proto_depIdxs = []int32{
	1, // 0: api.UserData.location:type_name -> api.GeoPoint
	2, // 1: api.CreateRequest.user:type_name -> api.UserData
	2, // 2: api.SearchResponse.users:type_name -> api.UserData
	3, // 3: api.User.Create:input_type -> api.CreateRequest
	4, // 4: api.User.Search:input_type -> api.SearchRequest
	0, // 5: api.User.Create:output_type -> api.EmptyResponse
	5, // 6: api.User.Search:output_type -> api.SearchResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
