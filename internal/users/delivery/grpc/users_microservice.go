package usersMicroservice

import (
	"context"

	"github.com/ztrue/tracerr"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	authProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"

	grpcUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type UsersMicroservice struct {
	client protobuf.UsersClient
}

func New(c protobuf.UsersClient) domain.UsersMicroservice {
	return &UsersMicroservice{
		client: c,
	}
}

func (m UsersMicroservice) Update(user models.User) error {
	_, err := m.client.Update(context.Background(), grpcUsers.ConvertToProto(user))

	return tracerr.Wrap(err)
}

func (m UsersMicroservice) Create(user models.User) (uint64, error) {
	id, err := m.client.Create(context.Background(), grpcUsers.ConvertToProto(user))
	if err != nil {
		return 0, tracerr.Wrap(err)
	}

	return id.GetUserId(), nil
}

func (m UsersMicroservice) GetAuthorByUsername(keyword string) ([]models.User, error) {
	authors, err := m.client.GetAuthorByUsername(context.Background(), &protobuf.Keyword{
		Keyword: keyword,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	result := make([]models.User, 0)

	for _, a := range authors.GetUsers() {
		result = append(result, grpcUsers.ConvertToModel(a))
	}

	return result, nil
}

func (m UsersMicroservice) GetAllAuthors() ([]models.User, error) {
	authors, err := m.client.GetAllAuthors(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	result := make([]models.User, 0)

	for _, author := range authors.GetUsers() {
		result = append(result, grpcUsers.ConvertToModel(author))
	}

	return result, nil
}

func (m UsersMicroservice) GetByID(id uint64) (models.User, error) {
	user, err := m.client.GetByID(context.Background(), &protobuf.UserID{UserId: id})
	if err != nil {
		return models.User{}, tracerr.Wrap(err)
	}

	return grpcUsers.ConvertToModel(user), nil
}

func (m UsersMicroservice) GetBySessionID(sessionID string) (models.User, error) {
	user, err := m.client.GetBySessionID(context.Background(), &authProto.SessionID{
		SessionId: sessionID,
	})
	if err != nil {
		return models.User{}, tracerr.Wrap(err)
	}

	return grpcUsers.ConvertToModel(user), nil
}

func (m UsersMicroservice) GetByEmail(email string) (models.User, error) {
	user, err := m.client.GetByEmail(context.Background(), &protobuf.Email{
		Email: email,
	})
	if err != nil {
		return models.User{}, tracerr.Wrap(err)
	}

	return grpcUsers.ConvertToModel(user), nil
}

func (m UsersMicroservice) GetByUsername(username string) (models.User, error) {
	user, err := m.client.GetByUsername(context.Background(), &protobuf.Username{
		Username: username,
	})
	if err != nil {
		return models.User{}, tracerr.Wrap(err)
	}

	return grpcUsers.ConvertToModel(user), nil
}

func (m UsersMicroservice) GetUserByPostID(postID uint64) (models.User, error) {
	user, err := m.client.GetUserByPostID(context.Background(), &protobuf.PostID{
		PostID: postID,
	})
	if err != nil {
		return models.User{}, tracerr.Wrap(err)
	}

	return grpcUsers.ConvertToModel(user), nil
}

func (m UsersMicroservice) GetPostsNum(userID uint64) (uint64, error) {
	num, err := m.client.GetPostsNum(context.Background(), &protobuf.UserID{UserId: userID})
	if err != nil {
		return 0, tracerr.Wrap(err)
	}

	return num.GetCountPosts(), nil
}

func (m UsersMicroservice) GetSubscribersNum(userID uint64) (uint64, error) {
	num, err := m.client.GetSubscribersNumForMounth(context.Background(), &protobuf.UserID{UserId: userID})
	if err != nil {
		return 0, tracerr.Wrap(err)
	}

	return num.GetCountSubscribers(), nil
}

func (m UsersMicroservice) GetProfitForMounth(userID uint64) (uint64, error) {
	num, err := m.client.GetProfitForMounth(context.Background(), &protobuf.UserID{UserId: userID})
	if err != nil {
		return 0, tracerr.Wrap(err)
	}

	return num.GetCountMounthProfit(), nil
}
