// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: v1/route.proto

package routiq

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TagCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           *Tag                   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Count         int32                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TagCount) Reset() {
	*x = TagCount{}
	mi := &file_v1_route_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCount) ProtoMessage() {}

func (x *TagCount) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCount.ProtoReflect.Descriptor instead.
func (*TagCount) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{0}
}

func (x *TagCount) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *TagCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Route リソース定義
type Route struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // XID型の一意のID
	DisplayName   string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // ルート名
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`                    // 説明
	Distance      float32                `protobuf:"fixed32,4,opt,name=distance,proto3" json:"distance,omitempty"`
	GeoJson       string                 `protobuf:"bytes,5,opt,name=geo_json,json=geoJson,proto3" json:"geo_json,omitempty"` // GeoJSON形式で表現されるfeature_collection
	GoogleMapUrl  string                 `protobuf:"bytes,6,opt,name=google_map_url,json=googleMapUrl,proto3" json:"google_map_url,omitempty"`
	TagCounts     []*TagCount            `protobuf:"bytes,7,rep,name=tag_counts,json=tagCounts,proto3" json:"tag_counts,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 作成日時 (ISO8601形式の文字列)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Route) Reset() {
	*x = Route{}
	mi := &file_v1_route_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Route) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Route) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Route) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Route) GetGeoJson() string {
	if x != nil {
		return x.GeoJson
	}
	return ""
}

func (x *Route) GetGoogleMapUrl() string {
	if x != nil {
		return x.GoogleMapUrl
	}
	return ""
}

func (x *Route) GetTagCounts() []*TagCount {
	if x != nil {
		return x.TagCounts
	}
	return nil
}

func (x *Route) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// CreateRoute RPC
type CreateRouteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Route         *Route                 `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"` // ルート名
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRouteRequest) Reset() {
	*x = CreateRouteRequest{}
	mi := &file_v1_route_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteRequest) ProtoMessage() {}

func (x *CreateRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteRequest) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRouteRequest) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

// ListRoutes RPC
type ListRoutesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int32                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`   // 一度に取得するルートの数
	PageToken     string                 `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"` // 次のページを指定するトークン
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoutesRequest) Reset() {
	*x = ListRoutesRequest{}
	mi := &file_v1_route_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesRequest) ProtoMessage() {}

func (x *ListRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{3}
}

func (x *ListRoutesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRoutesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRoutesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Routes        []*Route               `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`                                      // ルートの一覧
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` // 次のページのトークン
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRoutesResponse) Reset() {
	*x = ListRoutesResponse{}
	mi := &file_v1_route_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesResponse) ProtoMessage() {}

func (x *ListRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{4}
}

func (x *ListRoutesResponse) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ListRoutesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// GetRoute RPC
type GetRouteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 取得するルートのID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRouteRequest) Reset() {
	*x = GetRouteRequest{}
	mi := &file_v1_route_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteRequest) ProtoMessage() {}

func (x *GetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_route_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteRequest.ProtoReflect.Descriptor instead.
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return file_v1_route_proto_rawDescGZIP(), []int{5}
}

func (x *GetRouteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_route_proto protoreflect.FileDescriptor

var file_v1_route_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x54,
	0x61, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xed, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6f, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x61, 0x70,
	0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x3b, 0xea, 0x41, 0x38, 0x0a,
	0x17, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2f, 0x7b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x7d, 0x2a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x32, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x4f, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6e,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x83, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x24, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x71, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x73, 0x61, 0x79,
	0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x61, 0x79, 0x63, 0x6c, 0x65, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x71, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_route_proto_rawDescOnce sync.Once
	file_v1_route_proto_rawDescData = file_v1_route_proto_rawDesc
)

func file_v1_route_proto_rawDescGZIP() []byte {
	file_v1_route_proto_rawDescOnce.Do(func() {
		file_v1_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_route_proto_rawDescData)
	})
	return file_v1_route_proto_rawDescData
}

var file_v1_route_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_route_proto_goTypes = []any{
	(*TagCount)(nil),              // 0: asaycle.routiq.v1.TagCount
	(*Route)(nil),                 // 1: asaycle.routiq.v1.Route
	(*CreateRouteRequest)(nil),    // 2: asaycle.routiq.v1.CreateRouteRequest
	(*ListRoutesRequest)(nil),     // 3: asaycle.routiq.v1.ListRoutesRequest
	(*ListRoutesResponse)(nil),    // 4: asaycle.routiq.v1.ListRoutesResponse
	(*GetRouteRequest)(nil),       // 5: asaycle.routiq.v1.GetRouteRequest
	(*Tag)(nil),                   // 6: asaycle.routiq.v1.Tag
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_v1_route_proto_depIdxs = []int32{
	6, // 0: asaycle.routiq.v1.TagCount.tag:type_name -> asaycle.routiq.v1.Tag
	0, // 1: asaycle.routiq.v1.Route.tag_counts:type_name -> asaycle.routiq.v1.TagCount
	7, // 2: asaycle.routiq.v1.Route.created_at:type_name -> google.protobuf.Timestamp
	1, // 3: asaycle.routiq.v1.CreateRouteRequest.route:type_name -> asaycle.routiq.v1.Route
	1, // 4: asaycle.routiq.v1.ListRoutesResponse.routes:type_name -> asaycle.routiq.v1.Route
	2, // 5: asaycle.routiq.v1.RouteService.CreateRoute:input_type -> asaycle.routiq.v1.CreateRouteRequest
	3, // 6: asaycle.routiq.v1.RouteService.ListRoutes:input_type -> asaycle.routiq.v1.ListRoutesRequest
	5, // 7: asaycle.routiq.v1.RouteService.GetRoute:input_type -> asaycle.routiq.v1.GetRouteRequest
	1, // 8: asaycle.routiq.v1.RouteService.CreateRoute:output_type -> asaycle.routiq.v1.Route
	4, // 9: asaycle.routiq.v1.RouteService.ListRoutes:output_type -> asaycle.routiq.v1.ListRoutesResponse
	1, // 10: asaycle.routiq.v1.RouteService.GetRoute:output_type -> asaycle.routiq.v1.Route
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_route_proto_init() }
func file_v1_route_proto_init() {
	if File_v1_route_proto != nil {
		return
	}
	file_v1_tag_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_route_proto_goTypes,
		DependencyIndexes: file_v1_route_proto_depIdxs,
		MessageInfos:      file_v1_route_proto_msgTypes,
	}.Build()
	File_v1_route_proto = out.File
	file_v1_route_proto_rawDesc = nil
	file_v1_route_proto_goTypes = nil
	file_v1_route_proto_depIdxs = nil
}
