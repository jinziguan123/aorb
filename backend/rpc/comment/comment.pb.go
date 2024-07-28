// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: comment.proto

package comment

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

type ActionCommentType int32

const (
	ActionCommentType_ACTION_COMMENT_TYPE_UNSPECIFIED ActionCommentType = 0
	ActionCommentType_ACTION_COMMENT_TYPE_ADD         ActionCommentType = 1
	ActionCommentType_ACTION_COMMENT_TYPE_DELETE      ActionCommentType = 2
)

// Enum value maps for ActionCommentType.
var (
	ActionCommentType_name = map[int32]string{
		0: "ACTION_COMMENT_TYPE_UNSPECIFIED",
		1: "ACTION_COMMENT_TYPE_ADD",
		2: "ACTION_COMMENT_TYPE_DELETE",
	}
	ActionCommentType_value = map[string]int32{
		"ACTION_COMMENT_TYPE_UNSPECIFIED": 0,
		"ACTION_COMMENT_TYPE_ADD":         1,
		"ACTION_COMMENT_TYPE_DELETE":      2,
	}
)

func (x ActionCommentType) Enum() *ActionCommentType {
	p := new(ActionCommentType)
	*p = x
	return p
}

func (x ActionCommentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionCommentType) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_proto_enumTypes[0].Descriptor()
}

func (ActionCommentType) Type() protoreflect.EnumType {
	return &file_comment_proto_enumTypes[0]
}

func (x ActionCommentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionCommentType.Descriptor instead.
func (ActionCommentType) EnumDescriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentUsername string                 `protobuf:"bytes,1,opt,name=comment_username,json=commentUsername,proto3" json:"comment_username,omitempty"`
	Content         string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 评论内容
	CommentUuid     string                 `protobuf:"bytes,3,opt,name=comment_uuid,json=commentUuid,proto3" json:"comment_uuid,omitempty"`
	CreateAt        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"` // 创建时间
	DeleteAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=delete_at,json=deleteAt,proto3" json:"delete_at,omitempty"` // 删除时间
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetCommentUsername() string {
	if x != nil {
		return x.CommentUsername
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCommentUuid() string {
	if x != nil {
		return x.CommentUuid
	}
	return ""
}

func (x *Comment) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Comment) GetDeleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteAt
	}
	return nil
}

type ActionCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PollUuid   string            `protobuf:"bytes,2,opt,name=poll_uuid,json=pollUuid,proto3" json:"poll_uuid,omitempty"`
	ActionType ActionCommentType `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=rpc.comment.ActionCommentType" json:"action_type,omitempty"`
	// Types that are assignable to Action:
	//
	//	*ActionCommentRequest_CommentText
	//	*ActionCommentRequest_CommentUuid
	Action isActionCommentRequest_Action `protobuf_oneof:"action"`
}

func (x *ActionCommentRequest) Reset() {
	*x = ActionCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionCommentRequest) ProtoMessage() {}

func (x *ActionCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionCommentRequest.ProtoReflect.Descriptor instead.
func (*ActionCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *ActionCommentRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ActionCommentRequest) GetPollUuid() string {
	if x != nil {
		return x.PollUuid
	}
	return ""
}

func (x *ActionCommentRequest) GetActionType() ActionCommentType {
	if x != nil {
		return x.ActionType
	}
	return ActionCommentType_ACTION_COMMENT_TYPE_UNSPECIFIED
}

func (m *ActionCommentRequest) GetAction() isActionCommentRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ActionCommentRequest) GetCommentText() string {
	if x, ok := x.GetAction().(*ActionCommentRequest_CommentText); ok {
		return x.CommentText
	}
	return ""
}

func (x *ActionCommentRequest) GetCommentUuid() string {
	if x, ok := x.GetAction().(*ActionCommentRequest_CommentUuid); ok {
		return x.CommentUuid
	}
	return ""
}

type isActionCommentRequest_Action interface {
	isActionCommentRequest_Action()
}

type ActionCommentRequest_CommentText struct {
	CommentText string `protobuf:"bytes,4,opt,name=comment_text,json=commentText,proto3,oneof"`
}

type ActionCommentRequest_CommentUuid struct {
	CommentUuid string `protobuf:"bytes,5,opt,name=comment_uuid,json=commentUuid,proto3,oneof"`
}

func (*ActionCommentRequest_CommentText) isActionCommentRequest_Action() {}

func (*ActionCommentRequest_CommentUuid) isActionCommentRequest_Action() {}

type ActionCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Comment    *Comment `protobuf:"bytes,3,opt,name=comment,proto3,oneof" json:"comment,omitempty"`
}

func (x *ActionCommentResponse) Reset() {
	*x = ActionCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionCommentResponse) ProtoMessage() {}

func (x *ActionCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionCommentResponse.ProtoReflect.Descriptor instead.
func (*ActionCommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *ActionCommentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ActionCommentResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ActionCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type ListCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollUuid string `protobuf:"bytes,1,opt,name=poll_uuid,json=pollUuid,proto3" json:"poll_uuid,omitempty"`
}

func (x *ListCommentRequest) Reset() {
	*x = ListCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentRequest) ProtoMessage() {}

func (x *ListCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentRequest.ProtoReflect.Descriptor instead.
func (*ListCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *ListCommentRequest) GetPollUuid() string {
	if x != nil {
		return x.PollUuid
	}
	return ""
}

type ListCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg   string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	CommentList []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"`
}

func (x *ListCommentResponse) Reset() {
	*x = ListCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentResponse) ProtoMessage() {}

func (x *ListCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentResponse.ProtoReflect.Descriptor instead.
func (*ListCommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *ListCommentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListCommentResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListCommentResponse) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

type CountCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollUuid string `protobuf:"bytes,1,opt,name=poll_uuid,json=pollUuid,proto3" json:"poll_uuid,omitempty"`
}

func (x *CountCommentRequest) Reset() {
	*x = CountCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCommentRequest) ProtoMessage() {}

func (x *CountCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCommentRequest.ProtoReflect.Descriptor instead.
func (*CountCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5}
}

func (x *CountCommentRequest) GetPollUuid() string {
	if x != nil {
		return x.PollUuid
	}
	return ""
}

type CountCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg    string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	CommentCount uint32 `protobuf:"varint,3,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
}

func (x *CountCommentResponse) Reset() {
	*x = CountCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCommentResponse) ProtoMessage() {}

func (x *CountCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCommentResponse.ProtoReflect.Descriptor instead.
func (*CountCommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{6}
}

func (x *CountCommentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CountCommentResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CountCommentResponse) GetCommentCount() uint32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x74, 0x22, 0xe4, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c,
	0x6c, 0x55, 0x75, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x6c, 0x6c, 0x55, 0x75, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x55, 0x75, 0x69, 0x64, 0x22, 0x7b, 0x0a,
	0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x75, 0x0a, 0x11, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x23, 0x0a, 0x1f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x10,
	0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x02, 0x32, 0x8f, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x69, 0x67, 0x4e, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x74, 0x74, 0x79, 0x48, 0x6f,
	0x6d, 0x65, 0x2f, 0x61, 0x6f, 0x72, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_comment_proto_goTypes = []any{
	(ActionCommentType)(0),        // 0: rpc.comment.ActionCommentType
	(*Comment)(nil),               // 1: rpc.comment.Comment
	(*ActionCommentRequest)(nil),  // 2: rpc.comment.ActionCommentRequest
	(*ActionCommentResponse)(nil), // 3: rpc.comment.ActionCommentResponse
	(*ListCommentRequest)(nil),    // 4: rpc.comment.ListCommentRequest
	(*ListCommentResponse)(nil),   // 5: rpc.comment.ListCommentResponse
	(*CountCommentRequest)(nil),   // 6: rpc.comment.CountCommentRequest
	(*CountCommentResponse)(nil),  // 7: rpc.comment.CountCommentResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_comment_proto_depIdxs = []int32{
	8, // 0: rpc.comment.Comment.create_at:type_name -> google.protobuf.Timestamp
	8, // 1: rpc.comment.Comment.delete_at:type_name -> google.protobuf.Timestamp
	0, // 2: rpc.comment.ActionCommentRequest.action_type:type_name -> rpc.comment.ActionCommentType
	1, // 3: rpc.comment.ActionCommentResponse.comment:type_name -> rpc.comment.Comment
	1, // 4: rpc.comment.ListCommentResponse.comment_list:type_name -> rpc.comment.Comment
	2, // 5: rpc.comment.CommentService.ActionComment:input_type -> rpc.comment.ActionCommentRequest
	4, // 6: rpc.comment.CommentService.ListComment:input_type -> rpc.comment.ListCommentRequest
	6, // 7: rpc.comment.CommentService.CountComment:input_type -> rpc.comment.CountCommentRequest
	3, // 8: rpc.comment.CommentService.ActionComment:output_type -> rpc.comment.ActionCommentResponse
	5, // 9: rpc.comment.CommentService.ListComment:output_type -> rpc.comment.ListCommentResponse
	7, // 10: rpc.comment.CommentService.CountComment:output_type -> rpc.comment.CountCommentResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Comment); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActionCommentRequest); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ActionCommentResponse); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListCommentRequest); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListCommentResponse); i {
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
		file_comment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CountCommentRequest); i {
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
		file_comment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CountCommentResponse); i {
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
	file_comment_proto_msgTypes[1].OneofWrappers = []any{
		(*ActionCommentRequest_CommentText)(nil),
		(*ActionCommentRequest_CommentUuid)(nil),
	}
	file_comment_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		EnumInfos:         file_comment_proto_enumTypes,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}
