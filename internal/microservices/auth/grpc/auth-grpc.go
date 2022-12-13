package grpcAuth

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type cookieCreator func(id uint64) models.Cookie

func createCookie(id uint64) models.Cookie {
	return models.Cookie{
		UserID:  id,
		Value:   uuid.New().String(),
		Expires: time.Now().AddDate(0, 1, 0),
	}
}

type Auth struct {
	authRepo      domain.AuthRepository
	cookieCreator cookieCreator
	protobuf.UnimplementedAuthServer
}

func New(authRepo domain.AuthRepository) protobuf.AuthServer {
	return &Auth{
		authRepo:      authRepo,
		cookieCreator: createCookie,
	}
}

func NewWithCookieCreator(authRepo domain.AuthRepository, c cookieCreator) protobuf.AuthServer {
	return &Auth{
		authRepo:      authRepo,
		cookieCreator: c,
	}
}

func (m *Auth) CreateSession(_ context.Context, in *protobuf.Session) (*protobuf.SessionID, error) {
	session, err := m.authRepo.CreateSession(m.cookieCreator(in.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protobuf.SessionID{
		SessionId: session.Value,
	}, nil
}

func (m *Auth) DeleteBySessionID(_ context.Context, in *protobuf.SessionID) (*emptypb.Empty, error) {
	err := m.authRepo.DeleteBySessionID(in.GetSessionId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (m *Auth) GetBySessionID(_ context.Context, in *protobuf.SessionID) (*protobuf.Session, error) {
	cookie, err := m.authRepo.GetBySessionID(in.GetSessionId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &protobuf.Session{
		SessionId: cookie.Value,
		UserId:    cookie.UserID,
		Expires:   timestamppb.New(cookie.Expires),
	}, nil
}
