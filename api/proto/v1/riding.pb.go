// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: v1/riding.proto

package routiq

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	date "google.golang.org/genproto/googleapis/type/date"
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

// Riding リソース定義
type Riding struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // XID型の一意のID
	RouteId       string                 `protobuf:"bytes,2,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	Tags          []*Tag                 `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Date          *date.Date             `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Title         string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"` // ルート名
	Note          string                 `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
	Score         int32                  `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	UserId        string                 `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 作成日時 (ISO8601形式の文字列)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Riding) Reset() {
	*x = Riding{}
	mi := &file_v1_riding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Riding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Riding) ProtoMessage() {}

func (x *Riding) ProtoReflect() protoreflect.Message {
	mi := &file_v1_riding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Riding.ProtoReflect.Descriptor instead.
func (*Riding) Descriptor() ([]byte, []int) {
	return file_v1_riding_proto_rawDescGZIP(), []int{0}
}

func (x *Riding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Riding) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *Riding) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Riding) GetDate() *date.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Riding) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Riding) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Riding) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Riding) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Riding) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// CreateRiding RPC
type CreateRidingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Riding        *Riding                `protobuf:"bytes,1,opt,name=riding,proto3" json:"riding,omitempty"` // ライド名
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRidingRequest) Reset() {
	*x = CreateRidingRequest{}
	mi := &file_v1_riding_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRidingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRidingRequest) ProtoMessage() {}

func (x *CreateRidingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_riding_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRidingRequest.ProtoReflect.Descriptor instead.
func (*CreateRidingRequest) Descriptor() ([]byte, []int) {
	return file_v1_riding_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRidingRequest) GetRiding() *Riding {
	if x != nil {
		return x.Riding
	}
	return nil
}

// ListRidings RPC
type ListRidingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filter        string                 `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`   // 一度に取得するルートの数
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"` // 次のページを指定するトークン
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRidingsRequest) Reset() {
	*x = ListRidingsRequest{}
	mi := &file_v1_riding_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRidingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRidingsRequest) ProtoMessage() {}

func (x *ListRidingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_riding_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRidingsRequest.ProtoReflect.Descriptor instead.
func (*ListRidingsRequest) Descriptor() ([]byte, []int) {
	return file_v1_riding_proto_rawDescGZIP(), []int{2}
}

func (x *ListRidingsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListRidingsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRidingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRidingsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ridings       []*Riding              `protobuf:"bytes,1,rep,name=ridings,proto3" json:"ridings,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` // 次のページのトークン
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRidingsResponse) Reset() {
	*x = ListRidingsResponse{}
	mi := &file_v1_riding_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRidingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRidingsResponse) ProtoMessage() {}

func (x *ListRidingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_riding_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRidingsResponse.ProtoReflect.Descriptor instead.
func (*ListRidingsResponse) Descriptor() ([]byte, []int) {
	return file_v1_riding_proto_rawDescGZIP(), []int{3}
}

func (x *ListRidingsResponse) GetRidings() []*Riding {
	if x != nil {
		return x.Ridings
	}
	return nil
}

func (x *ListRidingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// GetRiding RPC
type GetRidingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 取得するルートのID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRidingRequest) Reset() {
	*x = GetRidingRequest{}
	mi := &file_v1_riding_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRidingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRidingRequest) ProtoMessage() {}

func (x *GetRidingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_riding_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRidingRequest.ProtoReflect.Descriptor instead.
func (*GetRidingRequest) Descriptor() ([]byte, []int) {
	return file_v1_riding_proto_rawDescGZIP(), []int{4}
}

func (x *GetRidingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_riding_proto protoreflect.FileDescriptor

var file_v1_riding_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x71, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x06, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x3a, 0x4e, 0xea, 0x41, 0x4b, 0x0a, 0x17, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x1f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x7d, 0x2f,
	0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x7d,
	0x2a, 0x07, 0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x32, 0x06, 0x72, 0x69, 0x64, 0x69, 0x6e,
	0x67, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x69, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x68, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x72, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07,
	0x72, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8d, 0x02,
	0x0a, 0x0d, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x26, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e,
	0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x69, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x61, 0x79,
	0x63, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_riding_proto_rawDescOnce sync.Once
	file_v1_riding_proto_rawDescData = file_v1_riding_proto_rawDesc
)

func file_v1_riding_proto_rawDescGZIP() []byte {
	file_v1_riding_proto_rawDescOnce.Do(func() {
		file_v1_riding_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_riding_proto_rawDescData)
	})
	return file_v1_riding_proto_rawDescData
}

var file_v1_riding_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_riding_proto_goTypes = []any{
	(*Riding)(nil),                // 0: asaycle.routiq.v1.Riding
	(*CreateRidingRequest)(nil),   // 1: asaycle.routiq.v1.CreateRidingRequest
	(*ListRidingsRequest)(nil),    // 2: asaycle.routiq.v1.ListRidingsRequest
	(*ListRidingsResponse)(nil),   // 3: asaycle.routiq.v1.ListRidingsResponse
	(*GetRidingRequest)(nil),      // 4: asaycle.routiq.v1.GetRidingRequest
	(*Tag)(nil),                   // 5: asaycle.routiq.v1.Tag
	(*date.Date)(nil),             // 6: google.type.Date
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_v1_riding_proto_depIdxs = []int32{
	5, // 0: asaycle.routiq.v1.Riding.tags:type_name -> asaycle.routiq.v1.Tag
	6, // 1: asaycle.routiq.v1.Riding.date:type_name -> google.type.Date
	7, // 2: asaycle.routiq.v1.Riding.created_at:type_name -> google.protobuf.Timestamp
	0, // 3: asaycle.routiq.v1.CreateRidingRequest.riding:type_name -> asaycle.routiq.v1.Riding
	0, // 4: asaycle.routiq.v1.ListRidingsResponse.ridings:type_name -> asaycle.routiq.v1.Riding
	1, // 5: asaycle.routiq.v1.RidingService.CreateRiding:input_type -> asaycle.routiq.v1.CreateRidingRequest
	2, // 6: asaycle.routiq.v1.RidingService.ListRidings:input_type -> asaycle.routiq.v1.ListRidingsRequest
	4, // 7: asaycle.routiq.v1.RidingService.GetRiding:input_type -> asaycle.routiq.v1.GetRidingRequest
	0, // 8: asaycle.routiq.v1.RidingService.CreateRiding:output_type -> asaycle.routiq.v1.Riding
	3, // 9: asaycle.routiq.v1.RidingService.ListRidings:output_type -> asaycle.routiq.v1.ListRidingsResponse
	0, // 10: asaycle.routiq.v1.RidingService.GetRiding:output_type -> asaycle.routiq.v1.Riding
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_riding_proto_init() }
func file_v1_riding_proto_init() {
	if File_v1_riding_proto != nil {
		return
	}
	file_v1_route_proto_init()
	file_v1_tag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_riding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_riding_proto_goTypes,
		DependencyIndexes: file_v1_riding_proto_depIdxs,
		MessageInfos:      file_v1_riding_proto_msgTypes,
	}.Build()
	File_v1_riding_proto = out.File
	file_v1_riding_proto_rawDesc = nil
	file_v1_riding_proto_goTypes = nil
	file_v1_riding_proto_depIdxs = nil
}
