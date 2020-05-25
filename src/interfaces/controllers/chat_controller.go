package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb/chat"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type chatController struct {
	chatInteractor interactor.ChatInteractor
}

func NewChatController(ci interactor.ChatInteractor) chat.ChatServiceServer {
	return &chatController{ci}
}

func (c *chatController) GetRoom(ctx context.Context, in *chat.GetRoomReq) (*chat.Room, error) {
	if (in.Id == 0 && in.PostId == 0) || (in.Id != 0 && in.PostId != 0) {
		return nil, status.Error(codes.InvalidArgument, "invalid GetRoomReq.Id, GetRoomReq.PostId: value must be set either id or post_id")
	}
	r, err := c.chatInteractor.GetRoom(ctx, in.Id, in.PostId)
	if err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) CreateRoom(ctx context.Context, in *chat.CreateRoomReq) (*chat.Room, error) {
	r := &domain.Room{
		PostID: in.PostId,
		Members: []*domain.Member{
			{UserID: in.UserId},
		},
	}
	if err := c.chatInteractor.CreateRoom(ctx, r); err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) GetMember(ctx context.Context, in *chat.GetMemberReq) (*chat.Member, error) {
	m, err := c.chatInteractor.GetMember(ctx, in.RoomId, in.UserId)
	if err != nil {
		return nil, err
	}
	return convMemberProto(m)
}

func (c *chatController) ListMembers(ctx context.Context, in *chat.ListMembersReq) (*chat.ListMembersRes, error) {
	list, err := c.chatInteractor.ListMembers(ctx, in.RoomId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMembersProto(list)
	if err != nil {
		return nil, err
	}
	return &chat.ListMembersRes{Members: listMProto}, nil
}

func (c *chatController) CreateMember(ctx context.Context, in *chat.CreateMemberReq) (*chat.Member, error) {
	m := &domain.Member{RoomID: in.RoomId, UserID: in.UserId}
	if err := c.chatInteractor.CreateMember(ctx, m); err != nil {
		return nil, err
	}
	return convMemberProto(m)
}

func (c *chatController) DeleteMember(ctx context.Context, in *chat.DeleteMemberReq) (*empty.Empty, error) {
	if err := c.chatInteractor.DeleteMember(ctx, in.RoomId, in.UserId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *chatController) ListMessages(ctx context.Context, in *chat.ListMessagesReq) (*chat.ListMessagesRes, error) {
	list, err := c.chatInteractor.ListMessages(ctx, in.RoomId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMessagesProto(list)
	if err != nil {
		return nil, err
	}
	return &chat.ListMessagesRes{Messages: listMProto}, nil
}

func (c *chatController) CreateMessage(ctx context.Context, in *chat.CreateMessageReq) (*chat.Message, error) {
	m := &domain.Message{
		Body:   in.Body,
		RoomID: in.RoomId,
		UserID: in.UserId,
	}
	if err := c.chatInteractor.CreateMessage(ctx, m); err != nil {
		return nil, err
	}
	return convMessageProto(m)
}

func (c *chatController) StreamMessage(in *chat.StreamMessageReq, stream chat.ChatService_StreamMessageServer) error {
	eg, ctx := errgroup.WithContext(stream.Context())
	msgChan := make(chan *domain.Message)
	go func() {
		eg.Wait()
		close(msgChan)
	}()

	eg.Go(func() error {
		if err := c.chatInteractor.StreamMessage(ctx, in.RoomId, in.UserId, msgChan); err != nil {
			return err
		}
		return nil
	})
	for m := range msgChan {
		mProto, err := convMessageProto(m)
		if err != nil {
			return err
		}
		if err := stream.Send(mProto); err != nil {
			return err
		}
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
