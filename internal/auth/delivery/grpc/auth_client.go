package authClient

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

type authServiceClient struct {
	authClient protobuf.AuthServiceClient
}

func NewAuthClient(authClient protobuf.AuthServiceClient) domain.AuthServiceManager {
	return authServiceClient{
		authClient: authClient,
	}
}

func (ac authServiceClient) CreateSession(userID uint64) (string, error) {
	sessionID, err := ac.authClient.CreateSession(context.Background(), &protobuf.Session{
		UserId: userID,
	})
	if err != nil {
		return string(""), err
	}

	return sessionID.GetSessionId(), nil
}

func (ac authServiceClient) DeleteBySessionID(sessionID string) error {
	_, err := ac.authClient.DeleteBySessionID(context.Background(), &protobuf.SessionID{
		SessionId: sessionID,
	})
	return err
}

func (ac authServiceClient) GetBySessionID(sessionID string) (*protobuf.Session, error) {
	session, err := ac.authClient.GetBySessionID(context.Background(), &protobuf.Session{
		SessionId: sessionID,
	})
	if err != nil {
		return nil, err
	}

	return &protobuf.Session{
		SessionId: session.SessionId,
	}, nil
}
