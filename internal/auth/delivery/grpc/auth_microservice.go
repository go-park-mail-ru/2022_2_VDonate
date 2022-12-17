package authMicroservice

import (
	"context"

	"github.com/ztrue/tracerr"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

type AuthMicroservice struct {
	authClient protobuf.AuthClient
}

func New(authClient protobuf.AuthClient) domain.AuthMicroservice {
	return &AuthMicroservice{
		authClient: authClient,
	}
}

func (m AuthMicroservice) CreateSession(userID uint64) (string, error) {
	sessionID, err := m.authClient.CreateSession(context.Background(), &protobuf.Session{
		UserId: userID,
	})
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	return sessionID.GetSessionId(), nil
}

func (m AuthMicroservice) DeleteBySessionID(sessionID string) error {
	_, err := m.authClient.DeleteBySessionID(context.Background(), &protobuf.SessionID{
		SessionId: sessionID,
	})
	return tracerr.Wrap(err)
}

func (m AuthMicroservice) GetBySessionID(sessionID string) (models.Cookie, error) {
	session, err := m.authClient.GetBySessionID(context.Background(), &protobuf.SessionID{
		SessionId: sessionID,
	})
	if err != nil {
		return models.Cookie{}, tracerr.Wrap(err)
	}
	if len(session.SessionId) == 0 {
		return models.Cookie{}, tracerr.Wrap(domain.ErrNoSession)
	}

	return models.Cookie{
		UserID:  session.GetUserId(),
		Value:   session.GetSessionId(),
		Expires: session.Expires.AsTime(),
	}, nil
}
