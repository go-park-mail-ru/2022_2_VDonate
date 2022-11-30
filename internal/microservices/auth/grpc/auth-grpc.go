package grpcAuth

import (
	"context"
	"time"

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

type AuthService struct {
	authRepo      domain.AuthRepository
	cookieCreator cookieCreator
	protobuf.UnimplementedAuthServiceServer
}

func New(authRepo domain.AuthRepository) *AuthService {
	return &AuthService{
		authRepo:      authRepo,
		cookieCreator: createCookie,
	}
}

func (as *AuthService) CreateSession(c context.Context, in *protobuf.Session) (*protobuf.SessionID, error) {
	session, err := as.authRepo.CreateSession(as.cookieCreator(in.UserId))
	if err != nil {
		return nil, err
	}
	return &protobuf.SessionID{
		SessionId: session.Value,
	}, nil
}

func (as *AuthService) DeleteBySessionID(c context.Context, in *protobuf.Session) (*emptypb.Empty, error) {
	err := as.authRepo.DeleteBySessionID(in.SessionId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (as *AuthService) GetBySessionID(c context.Context, in *protobuf.Session) (*protobuf.SessionID, error) {
	cookie, err := as.authRepo.GetBySessionID(in.SessionId)
	if err != nil {
		return nil, err
	}
	return &protobuf.SessionID{
		SessionId: cookie.Value,
	}, nil
}
