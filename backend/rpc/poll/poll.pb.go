// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: poll.proto

package poll

import (
	comment "github.com/BigNoseCattyHome/aorb/backend/rpc/comment"
	vote "github.com/BigNoseCattyHome/aorb/backend/rpc/vote"
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

type Poll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 投票id
	Title        string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                           // 投票标题
	Option       []string               `protobuf:"bytes,3,rep,name=option,proto3" json:"option,omitempty"`                                         // 选项
	OptionsCount []int32                `protobuf:"varint,4,rep,packed,name=options_count,json=optionsCount,proto3" json:"options_count,omitempty"` // 投票计数
	PollType     string                 `protobuf:"bytes,5,opt,name=poll_type,json=pollType,proto3" json:"poll_type,omitempty"`                     // 投票类型
	Username     string                 `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`                                     // 投票发起人
	CreateAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`                     // 创建时间
	UpdateAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`                     // 删除时间
	DeleteAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=delete_at,json=deleteAt,proto3" json:"delete_at,omitempty"`                     // 删除时间
	CommentList  []*comment.Comment     `protobuf:"bytes,10,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"`           // 评论列表
	VoteList     []*vote.Vote           `protobuf:"bytes,11,rep,name=vote_list,json=voteList,proto3" json:"vote_list,omitempty"`                    // 投票列表
}

func (x *Poll) Reset() {
	*x = Poll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poll) ProtoMessage() {}

func (x *Poll) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Poll.ProtoReflect.Descriptor instead.
func (*Poll) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{0}
}

func (x *Poll) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Poll) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Poll) GetOption() []string {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *Poll) GetOptionsCount() []int32 {
	if x != nil {
		return x.OptionsCount
	}
	return nil
}

func (x *Poll) GetPollType() string {
	if x != nil {
		return x.PollType
	}
	return ""
}

func (x *Poll) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Poll) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Poll) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

func (x *Poll) GetDeleteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteAt
	}
	return nil
}

func (x *Poll) GetCommentList() []*comment.Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

func (x *Poll) GetVoteList() []*vote.Vote {
	if x != nil {
		return x.VoteList
	}
	return nil
}

type CreatePollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Poll *Poll `protobuf:"bytes,1,opt,name=poll,proto3" json:"poll,omitempty"`
}

func (x *CreatePollRequest) Reset() {
	*x = CreatePollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePollRequest) ProtoMessage() {}

func (x *CreatePollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePollRequest.ProtoReflect.Descriptor instead.
func (*CreatePollRequest) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePollRequest) GetPoll() *Poll {
	if x != nil {
		return x.Poll
	}
	return nil
}

type CreatePollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	PollId     string `protobuf:"bytes,3,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`              // 返回创建的投票id
}

func (x *CreatePollResponse) Reset() {
	*x = CreatePollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePollResponse) ProtoMessage() {}

func (x *CreatePollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePollResponse.ProtoReflect.Descriptor instead.
func (*CreatePollResponse) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePollResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreatePollResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CreatePollResponse) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type GetPollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId string `protobuf:"bytes,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *GetPollRequest) Reset() {
	*x = GetPollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPollRequest) ProtoMessage() {}

func (x *GetPollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPollRequest.ProtoReflect.Descriptor instead.
func (*GetPollRequest) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{3}
}

func (x *GetPollRequest) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type GetPollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	Poll       *Poll  `protobuf:"bytes,3,opt,name=poll,proto3" json:"poll,omitempty"`
}

func (x *GetPollResponse) Reset() {
	*x = GetPollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPollResponse) ProtoMessage() {}

func (x *GetPollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPollResponse.ProtoReflect.Descriptor instead.
func (*GetPollResponse) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{4}
}

func (x *GetPollResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetPollResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetPollResponse) GetPoll() *Poll {
	if x != nil {
		return x.Poll
	}
	return nil
}

type ListPollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 第几页
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`   // 一页多少
}

func (x *ListPollRequest) Reset() {
	*x = ListPollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPollRequest) ProtoMessage() {}

func (x *ListPollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPollRequest.ProtoReflect.Descriptor instead.
func (*ListPollRequest) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{5}
}

func (x *ListPollRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPollRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListPollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	PollList   []*Poll `protobuf:"bytes,3,rep,name=poll_list,json=pollList,proto3" json:"poll_list,omitempty"`        // 投票列表
}

func (x *ListPollResponse) Reset() {
	*x = ListPollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poll_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPollResponse) ProtoMessage() {}

func (x *ListPollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_poll_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPollResponse.ProtoReflect.Descriptor instead.
func (*ListPollResponse) Descriptor() ([]byte, []int) {
	return file_poll_proto_rawDescGZIP(), []int{6}
}

func (x *ListPollResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListPollResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListPollResponse) GetPollList() []*Poll {
	if x != nil {
		return x.PollList
	}
	return nil
}

var File_poll_proto protoreflect.FileDescriptor

var file_poll_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x03, 0x0a, 0x04, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x08,
	0x76, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6c,
	0x6c, 0x22, 0x6d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x22,
	0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x04, 0x70, 0x6f,
	0x6c, 0x6c, 0x22, 0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xd9, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x6f, 0x6c, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42,
	0x69, 0x67, 0x4e, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x74, 0x74, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x2f,
	0x61, 0x6f, 0x72, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x3b, 0x70, 0x6f, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_poll_proto_rawDescOnce sync.Once
	file_poll_proto_rawDescData = file_poll_proto_rawDesc
)

func file_poll_proto_rawDescGZIP() []byte {
	file_poll_proto_rawDescOnce.Do(func() {
		file_poll_proto_rawDescData = protoimpl.X.CompressGZIP(file_poll_proto_rawDescData)
	})
	return file_poll_proto_rawDescData
}

var file_poll_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_poll_proto_goTypes = []any{
	(*Poll)(nil),                  // 0: rpc.poll.Poll
	(*CreatePollRequest)(nil),     // 1: rpc.poll.CreatePollRequest
	(*CreatePollResponse)(nil),    // 2: rpc.poll.CreatePollResponse
	(*GetPollRequest)(nil),        // 3: rpc.poll.GetPollRequest
	(*GetPollResponse)(nil),       // 4: rpc.poll.GetPollResponse
	(*ListPollRequest)(nil),       // 5: rpc.poll.ListPollRequest
	(*ListPollResponse)(nil),      // 6: rpc.poll.ListPollResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*comment.Comment)(nil),       // 8: rpc.comment.Comment
	(*vote.Vote)(nil),             // 9: rpc.vote.Vote
}
var file_poll_proto_depIdxs = []int32{
	7,  // 0: rpc.poll.Poll.create_at:type_name -> google.protobuf.Timestamp
	7,  // 1: rpc.poll.Poll.update_at:type_name -> google.protobuf.Timestamp
	7,  // 2: rpc.poll.Poll.delete_at:type_name -> google.protobuf.Timestamp
	8,  // 3: rpc.poll.Poll.comment_list:type_name -> rpc.comment.Comment
	9,  // 4: rpc.poll.Poll.vote_list:type_name -> rpc.vote.Vote
	0,  // 5: rpc.poll.CreatePollRequest.poll:type_name -> rpc.poll.Poll
	0,  // 6: rpc.poll.GetPollResponse.poll:type_name -> rpc.poll.Poll
	0,  // 7: rpc.poll.ListPollResponse.poll_list:type_name -> rpc.poll.Poll
	1,  // 8: rpc.poll.PollService.CreatePoll:input_type -> rpc.poll.CreatePollRequest
	3,  // 9: rpc.poll.PollService.GetPoll:input_type -> rpc.poll.GetPollRequest
	5,  // 10: rpc.poll.PollService.ListPoll:input_type -> rpc.poll.ListPollRequest
	2,  // 11: rpc.poll.PollService.CreatePoll:output_type -> rpc.poll.CreatePollResponse
	4,  // 12: rpc.poll.PollService.GetPoll:output_type -> rpc.poll.GetPollResponse
	6,  // 13: rpc.poll.PollService.ListPoll:output_type -> rpc.poll.ListPollResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_poll_proto_init() }
func file_poll_proto_init() {
	if File_poll_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_poll_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Poll); i {
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
		file_poll_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePollRequest); i {
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
		file_poll_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePollResponse); i {
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
		file_poll_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPollRequest); i {
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
		file_poll_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPollResponse); i {
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
		file_poll_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListPollRequest); i {
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
		file_poll_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListPollResponse); i {
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
			RawDescriptor: file_poll_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_poll_proto_goTypes,
		DependencyIndexes: file_poll_proto_depIdxs,
		MessageInfos:      file_poll_proto_msgTypes,
	}.Build()
	File_poll_proto = out.File
	file_poll_proto_rawDesc = nil
	file_poll_proto_goTypes = nil
	file_poll_proto_depIdxs = nil
}
