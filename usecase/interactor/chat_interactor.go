package interactor

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatInteractor interface {
	CreateRoom(ctx context.Context, r *domain.Room) error
	GetRoom(ctx context.Context, id int64, postID int64) (*domain.Room, error)
	ListMembers(ctx context.Context, roomID int64) ([]*domain.Member, error)
	GetMember(ctx context.Context, roomID int64, userID int64) (*domain.Member, error)
	CreateMember(ctx context.Context, m *domain.Member) error
	DeleteMember(ctx context.Context, roomID int64, userID int64) error
	ListMessages(ctx context.Context, roomID int64) ([]*domain.Message, error)
	CreateMessage(ctx context.Context, m *domain.Message) error
	StreamMessage(ctx context.Context, roomID int64, userID int64, msgChan chan *domain.Message) error
}

type chatInteractor struct {
	db         *gorm.DB
	rdb        *redis.Client
	ctxTimeout time.Duration
}

func NewChatInteractor(db *gorm.DB, redis *redis.Client, t time.Duration) ChatInteractor {
	return &chatInteractor{db, redis, t}
}

func (i *chatInteractor) CreateRoom(ctx context.Context, r *domain.Room) error {
	if err := i.db.Create(r).Error; err != nil {
		return err
	}
	return nil
}

func (i *chatInteractor) GetRoom(ctx context.Context, id int64, pID int64) (*domain.Room, error) {
	r := &domain.Room{}
	if id != 0 {
		if err := i.db.Take(id, &r).Error; err != nil {
			return nil, err
		}
	}
	if pID != 0 {
		if err := i.db.Where("post_id = ?", pID).Take(&r).Error; err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (i *chatInteractor) GetMember(ctx context.Context, rID int64, uID int64) (*domain.Member, error) {
	m := &domain.Member{}
	if err := i.db.Where("room_id = ? AND user_id = ?", rID, uID).Take(m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "member with user_id='%d', room_id='%d' is not found", uID, rID)
		}
		return nil, err
	}

	return m, nil
}

func (i *chatInteractor) ListMembers(ctx context.Context, rID int64) ([]*domain.Member, error) {
	r := &domain.Room{ID: rID}
	if err := i.db.Model(r).Related(&r.Members).Error; err != nil {
		return nil, err
	}
	return r.Members, nil
}

func (i *chatInteractor) CreateMember(ctx context.Context, m *domain.Member) error {
	if err := i.db.Create(m).Error; err != nil {
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}
	return nil
}

func (i *chatInteractor) DeleteMember(ctx context.Context, rID int64, uID int64) error {
	if err := i.db.Where("room_id = ? AND user_id = ?", rID, uID).Delete(&domain.Member{}).Error; err != nil {
		return err
	}
	return nil
}

func (i *chatInteractor) ListMessages(ctx context.Context, rID int64) ([]*domain.Message, error) {
	r := &domain.Room{ID: rID}
	if err := i.db.Model(r).Related(&r.Messages).Error; err != nil {
		return nil, err
	}
	return r.Messages, nil
}

func (i *chatInteractor) CreateMessage(ctx context.Context, m *domain.Message) error {
	if err := i.db.Create(m).Error; err != nil {
		return err
	}

	mb, err := json.Marshal(*m)
	if err != nil {
		return err
	}
	i.rdb.Publish(strconv.FormatInt(m.RoomID, 10), mb)
	return nil
}

func (i *chatInteractor) StreamMessage(ctx context.Context, rID int64, uID int64, msgChan chan *domain.Message) error {
	pubsub := i.rdb.WithContext(ctx).Subscribe(strconv.FormatInt(rID, 10))
	go func() {
		<-ctx.Done()
		pubsub.Close()
	}()

	for m := range pubsub.Channel() {
		msg := &domain.Message{}
		if err := json.Unmarshal([]byte(m.Payload), &msg); err != nil {
			return err
		}
		if msg.UserID != uID {
			msgChan <- msg
		}
	}
	return nil
}