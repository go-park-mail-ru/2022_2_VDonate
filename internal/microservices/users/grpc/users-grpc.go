package grpcUsers

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	authProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
)

func ConvertToModel(u *protobuf.User) models.User {
	return models.User{
		ID:                 u.GetId(),
		Username:           u.GetUsername(),
		Email:              u.GetEmail(),
		Avatar:             u.GetAvatar(),
		Password:           u.GetPassword(),
		IsAuthor:           u.GetIsAuthor(),
		About:              u.GetAbout(),
		CountSubscriptions: u.GetCountSubscriptions(),
		CountSubscribers:   u.GetCountSubscribers(),
	}
}

func ConvertToProto(u models.User) *protobuf.User {
	return &protobuf.User{
		Id:                 u.ID,
		Username:           u.Username,
		Email:              u.Email,
		Password:           u.Password,
		Avatar:             u.Avatar,
		IsAuthor:           u.IsAuthor,
		About:              u.About,
		CountSubscriptions: u.CountSubscriptions,
		CountSubscribers:   u.CountSubscribers,
	}
}

type UserService struct {
	userRepo domain.UsersRepository
	protobuf.UnimplementedUsersServer
}

func New(s domain.UsersRepository) protobuf.UsersServer {
	return &UserService{
		userRepo: s,
	}
}

func (s UserService) GetAllAuthors(_ context.Context, _ *emptypb.Empty) (*protobuf.UsersArray, error) {
	authors, err := s.userRepo.GetAllAuthors()
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*protobuf.User, 0)

	for _, author := range authors {
		result = append(result, ConvertToProto(author))
	}

	return &protobuf.UsersArray{
		Users: result,
	}, nil
}

func (s UserService) Create(_ context.Context, user *protobuf.User) (*protobuf.UserID, error) {
	id, err := s.userRepo.Create(ConvertToModel(user))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &protobuf.UserID{UserId: id}, nil
}

func (s UserService) Update(_ context.Context, user *protobuf.User) (*emptypb.Empty, error) {
	err := s.userRepo.Update(ConvertToModel(user))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s UserService) GetAuthorByUsername(_ context.Context, key *protobuf.Keyword) (*protobuf.UsersArray, error) {
	users, err := s.userRepo.GetAuthorByUsername(key.GetKeyword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*protobuf.User, 0)
	for _, u := range users {
		result = append(result, ConvertToProto(u))
	}

	return &protobuf.UsersArray{Users: result}, nil
}

func (s UserService) GetUserByPostID(_ context.Context, postID *protobuf.PostID) (*protobuf.User, error) {
	user, err := s.userRepo.GetUserByPostID(postID.GetPostID())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(user), nil
}

func (s UserService) GetByID(_ context.Context, id *protobuf.UserID) (*protobuf.User, error) {
	user, err := s.userRepo.GetByID(id.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(user), nil
}

func (s UserService) GetBySessionID(_ context.Context, sessionID *authProto.SessionID) (*protobuf.User, error) {
	user, err := s.userRepo.GetBySessionID(sessionID.GetSessionId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(user), nil
}

func (s UserService) GetByEmail(_ context.Context, email *protobuf.Email) (*protobuf.User, error) {
	user, err := s.userRepo.GetByEmail(email.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(user), nil
}

func (s UserService) GetByUsername(_ context.Context, username *protobuf.Username) (*protobuf.User, error) {
	user, err := s.userRepo.GetByUsername(username.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(user), nil
}
