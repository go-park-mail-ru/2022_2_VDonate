package authMicroservice

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

type AuthMicroservice struct {
	authClient protobuf.AuthServiceClient
}

func New(authClient protobuf.AuthServiceClient) domain.AuthMicroservice {
	return &AuthMicroservice{
		authClient: authClient,
	}
}

func (m AuthMicroservice) CreateSession(userID uint64) (string, error) {
	sessionID, err := m.authClient.CreateSession(context.Background(), &protobuf.Session{
		UserId: userID,
	})
	if err != nil {
		return string(""), err
	}

	return sessionID.GetSessionId(), nil
}

func (m AuthMicroservice) DeleteBySessionID(sessionID string) error {
	_, err := m.authClient.DeleteBySessionID(context.Background(), &protobuf.SessionID{
		SessionId: sessionID,
	})
	return err
}

func (m AuthMicroservice) GetBySessionID(sessionID string) (*protobuf.Session, error) {
	session, err := m.authClient.GetBySessionID(context.Background(), &protobuf.Session{
		SessionId: sessionID,
	})
	if err != nil {
		return nil, err
	}

	return &protobuf.Session{
		SessionId: session.SessionId,
	}, nil
}
