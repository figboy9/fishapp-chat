// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat_grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	RoomId               int64                `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId               int64                `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Message) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Message) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Message) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Room struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId               int64                `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Room) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *Room) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Room) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type SendMessageReq struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	RoomId               int64    `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReq.Unmarshal(m, b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return xxx_messageInfo_SendMessageReq.Size(m)
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

func (m *SendMessageReq) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendMessageReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *SendMessageReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ReceiveMessageReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMessageReq) Reset()         { *m = ReceiveMessageReq{} }
func (m *ReceiveMessageReq) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessageReq) ProtoMessage()    {}
func (*ReceiveMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *ReceiveMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMessageReq.Unmarshal(m, b)
}
func (m *ReceiveMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMessageReq.Marshal(b, m, deterministic)
}
func (m *ReceiveMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMessageReq.Merge(m, src)
}
func (m *ReceiveMessageReq) XXX_Size() int {
	return xxx_messageInfo_ReceiveMessageReq.Size(m)
}
func (m *ReceiveMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMessageReq proto.InternalMessageInfo

func (m *ReceiveMessageReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ReceiveMessageReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AddMemberReq struct {
	RoomId               int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMemberReq) Reset()         { *m = AddMemberReq{} }
func (m *AddMemberReq) String() string { return proto.CompactTextString(m) }
func (*AddMemberReq) ProtoMessage()    {}
func (*AddMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
}

func (m *AddMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMemberReq.Unmarshal(m, b)
}
func (m *AddMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMemberReq.Marshal(b, m, deterministic)
}
func (m *AddMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMemberReq.Merge(m, src)
}
func (m *AddMemberReq) XXX_Size() int {
	return xxx_messageInfo_AddMemberReq.Size(m)
}
func (m *AddMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddMemberReq proto.InternalMessageInfo

func (m *AddMemberReq) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *AddMemberReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateRoomReq struct {
	PostId               int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomReq) Reset()         { *m = CreateRoomReq{} }
func (m *CreateRoomReq) String() string { return proto.CompactTextString(m) }
func (*CreateRoomReq) ProtoMessage()    {}
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{5}
}

func (m *CreateRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomReq.Unmarshal(m, b)
}
func (m *CreateRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomReq.Marshal(b, m, deterministic)
}
func (m *CreateRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomReq.Merge(m, src)
}
func (m *CreateRoomReq) XXX_Size() int {
	return xxx_messageInfo_CreateRoomReq.Size(m)
}
func (m *CreateRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomReq proto.InternalMessageInfo

func (m *CreateRoomReq) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "chat_grpc.Message")
	proto.RegisterType((*Room)(nil), "chat_grpc.Room")
	proto.RegisterType((*SendMessageReq)(nil), "chat_grpc.SendMessageReq")
	proto.RegisterType((*ReceiveMessageReq)(nil), "chat_grpc.ReceiveMessageReq")
	proto.RegisterType((*AddMemberReq)(nil), "chat_grpc.AddMemberReq")
	proto.RegisterType((*CreateRoomReq)(nil), "chat_grpc.CreateRoomReq")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x14, 0x94, 0xd3, 0xd0, 0xaa, 0xaf, 0x50, 0x84, 0x39, 0x34, 0x44, 0x08, 0xa2, 0x9e, 0x72, 0x4a,
	0x61, 0x39, 0x21, 0x24, 0xa4, 0x76, 0xb9, 0x2c, 0xd2, 0x5e, 0xb2, 0x2b, 0x38, 0x56, 0x6e, 0xfc,
	0x68, 0x23, 0x35, 0x38, 0x38, 0x4e, 0x11, 0xff, 0xc2, 0x95, 0x4f, 0xe2, 0x67, 0x38, 0x21, 0xdb,
	0x49, 0xd6, 0x4d, 0x57, 0x5a, 0x81, 0xf6, 0xe6, 0xe7, 0x19, 0x4f, 0xe7, 0xcd, 0x34, 0x00, 0xd9,
	0x8e, 0xa9, 0xa4, 0x94, 0x42, 0x09, 0x3a, 0xd6, 0xe7, 0xf5, 0x56, 0x96, 0x59, 0x38, 0x3b, 0xb0,
	0x7d, 0xce, 0x99, 0xc2, 0x45, 0x7b, 0xb0, 0x9c, 0xf0, 0xe5, 0x56, 0x88, 0xed, 0x1e, 0x17, 0x66,
	0xda, 0xd4, 0x5f, 0x16, 0x2a, 0x2f, 0xb0, 0x52, 0xac, 0x28, 0x1b, 0xc2, 0x8b, 0x3e, 0xe1, 0xbb,
	0x64, 0x65, 0x89, 0xb2, 0xb2, 0xf8, 0xfc, 0x37, 0x81, 0xd1, 0x25, 0x56, 0x15, 0xdb, 0x22, 0x9d,
	0x82, 0x97, 0xf3, 0x80, 0x44, 0x24, 0x1e, 0xa4, 0x5e, 0xce, 0x29, 0x05, 0x7f, 0x23, 0xf8, 0x8f,
	0xc0, 0x8b, 0x48, 0x3c, 0x4e, 0xcd, 0x99, 0xce, 0x60, 0x24, 0x85, 0x28, 0xd6, 0x39, 0x0f, 0x06,
	0x86, 0x38, 0xd4, 0xe3, 0x05, 0xd7, 0x40, 0x5d, 0xa1, 0xd4, 0x80, 0x6f, 0x01, 0x3d, 0x5e, 0x70,
	0xfa, 0x16, 0x20, 0x93, 0xc8, 0x14, 0xf2, 0x35, 0x53, 0xc1, 0x83, 0x88, 0xc4, 0x93, 0xb3, 0x30,
	0xb1, 0xb6, 0x92, 0xd6, 0x56, 0x72, 0xdd, 0xfa, 0x4e, 0xc7, 0x0d, 0x7b, 0xa9, 0xf4, 0xd3, 0xba,
	0xe4, 0xed, 0xd3, 0xe1, 0xdd, 0x4f, 0x1b, 0xf6, 0x52, 0xcd, 0x7f, 0x11, 0xf0, 0x53, 0x21, 0x8a,
	0x93, 0xa5, 0x66, 0x30, 0x2a, 0x45, 0xa5, 0xb4, 0x4f, 0xcf, 0xfa, 0xd4, 0xe3, 0x89, 0xcf, 0xc1,
	0xff, 0xfb, 0xf4, 0xff, 0xc5, 0xe7, 0x0e, 0xa6, 0x57, 0xf8, 0x95, 0x37, 0x15, 0xa4, 0xf8, 0xad,
	0x4b, 0x9d, 0x38, 0xa9, 0x47, 0x37, 0xa9, 0x1b, 0xd3, 0xab, 0xd1, 0x9f, 0x95, 0x3f, 0xf7, 0x62,
	0xd2, 0xc5, 0x1f, 0xdd, 0xc4, 0x3f, 0xe8, 0x31, 0x6c, 0x0f, 0xf3, 0xcf, 0xf0, 0x24, 0xc5, 0x0c,
	0xf3, 0x03, 0x3a, 0x3f, 0xe6, 0x08, 0x93, 0x3b, 0x85, 0xbd, 0xdb, 0x85, 0x53, 0x78, 0xb8, 0xe4,
	0xfc, 0x12, 0x8b, 0x0d, 0xca, 0xfb, 0xd2, 0x7c, 0x0d, 0x8f, 0xce, 0x4d, 0xbc, 0xba, 0xc3, 0x46,
	0xb4, 0xad, 0xad, 0x2f, 0x6a, 0xfb, 0x3b, 0xfb, 0xe9, 0xc1, 0xe4, 0x7c, 0xc7, 0xd4, 0x15, 0xca,
	0x43, 0x9e, 0x21, 0x7d, 0x07, 0x53, 0x2b, 0xa1, 0x2f, 0xcd, 0x5f, 0x21, 0x48, 0xba, 0x2f, 0x2a,
	0x39, 0x52, 0x0f, 0x1f, 0x3b, 0x88, 0xa1, 0x7e, 0x84, 0xa7, 0xdd, 0x4e, 0xd7, 0xa2, 0x53, 0x98,
	0x39, 0x3c, 0x77, 0xe7, 0xf0, 0xb4, 0xed, 0x95, 0x10, 0xfb, 0x4f, 0x6c, 0x5f, 0x23, 0x7d, 0x0f,
	0x13, 0xa7, 0x62, 0xfa, 0xcc, 0xd1, 0x38, 0xae, 0x3e, 0xa4, 0x0e, 0xd4, 0x5c, 0xc7, 0x84, 0x7e,
	0x80, 0xe9, 0x71, 0x71, 0xf4, 0xb9, 0x6b, 0xb7, 0xdf, 0xe9, 0x6d, 0x2a, 0xaf, 0xc8, 0x66, 0x68,
	0x9c, 0xbd, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x05, 0xf9, 0x5e, 0xd6, 0x62, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChatRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*Room, error)
	AddMemberToChatRoom(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
	ReceiveMessage(ctx context.Context, in *ReceiveMessageReq, opts ...grpc.CallOption) (ChatService_ReceiveMessageClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChatRoom(ctx context.Context, in *CreateRoomReq, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/chat_grpc.ChatService/CreateChatRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddMemberToChatRoom(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/chat_grpc.ChatService/AddMemberToChatRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat_grpc.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*SendMessageReq) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *SendMessageReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ReceiveMessage(ctx context.Context, in *ReceiveMessageReq, opts ...grpc.CallOption) (ChatService_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[1], "/chat_grpc.ChatService/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ReceiveMessageClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceReceiveMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	CreateChatRoom(context.Context, *CreateRoomReq) (*Room, error)
	AddMemberToChatRoom(context.Context, *AddMemberReq) (*wrappers.BoolValue, error)
	SendMessage(ChatService_SendMessageServer) error
	ReceiveMessage(*ReceiveMessageReq, ChatService_ReceiveMessageServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) CreateChatRoom(ctx context.Context, req *CreateRoomReq) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatRoom not implemented")
}
func (*UnimplementedChatServiceServer) AddMemberToChatRoom(ctx context.Context, req *AddMemberReq) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberToChatRoom not implemented")
}
func (*UnimplementedChatServiceServer) SendMessage(srv ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedChatServiceServer) ReceiveMessage(req *ReceiveMessageReq, srv ChatService_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_CreateChatRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChatRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_grpc.ChatService/CreateChatRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChatRoom(ctx, req.(*CreateRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddMemberToChatRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMemberToChatRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_grpc.ChatService/AddMemberToChatRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMemberToChatRoom(ctx, req.(*AddMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	SendAndClose(*Message) error
	Recv() (*SendMessageReq, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*SendMessageReq, error) {
	m := new(SendMessageReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveMessageReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ReceiveMessage(m, &chatServiceReceiveMessageServer{stream})
}

type ChatService_ReceiveMessageServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceReceiveMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat_grpc.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChatRoom",
			Handler:    _ChatService_CreateChatRoom_Handler,
		},
		{
			MethodName: "AddMemberToChatRoom",
			Handler:    _ChatService_AddMemberToChatRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveMessage",
			Handler:       _ChatService_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
