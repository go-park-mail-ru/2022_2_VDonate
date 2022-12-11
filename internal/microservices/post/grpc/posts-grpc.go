package grpcPosts

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"

	usersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

func ConvertToProto(p models.Post) *protobuf.Post {
	return &protobuf.Post{
		ID:          p.ID,
		UserID:      p.UserID,
		Content:     p.Content,
		Tier:        p.Tier,
		IsAllowed:   p.IsAllowed,
		DateCreated: timestamppb.New(p.DateCreated),
		Tags:        p.Tags,
		Author: &usersProto.LessUser{
			Id:       p.Author.UserID,
			Username: p.Author.Username,
			Avatar:   p.Author.ImgPath,
		},
		LikesNum: p.LikesNum,
		IsLiked:  p.IsLiked,
	}
}

func ConvertToModel(p *protobuf.Post) models.Post {
	return models.Post{
		ID:          p.ID,
		UserID:      p.UserID,
		Content:     p.Content,
		Tier:        p.Tier,
		IsAllowed:   p.IsAllowed,
		DateCreated: p.DateCreated.AsTime(),
		Tags:        p.Tags,
		Author: models.ResponseImageUsers{
			UserID:   p.Author.GetId(),
			Username: p.Author.GetUsername(),
			ImgPath:  p.Author.GetAvatar(),
		},
		LikesNum: p.LikesNum,
		IsLiked:  p.IsLiked,
	}
}

type PostsService struct {
	postsRepo domain.PostsRepository
	protobuf.UnimplementedPostsServer
}

func New(p domain.PostsRepository) protobuf.PostsServer {
	return &PostsService{
		postsRepo: p,
	}
}

func (s PostsService) GetAllByUserID(_ context.Context, userID *usersProto.UserID) (*protobuf.PostArray, error) {
	posts, err := s.postsRepo.GetAllByUserID(userID.GetUserId())
	if err != nil {
		return nil, err
	}

	result := make([]*protobuf.Post, 0)
	for _, post := range posts {
		result = append(result, ConvertToProto(post))
	}

	return &protobuf.PostArray{Posts: result}, nil
}

func (s PostsService) GetPostByID(_ context.Context, postID *protobuf.PostID) (*protobuf.Post, error) {
	post, err := s.postsRepo.GetPostByID(postID.GetPostID())
	if err != nil {
		return nil, err
	}

	return ConvertToProto(post), nil
}

func (s PostsService) Create(_ context.Context, post *protobuf.Post) (*protobuf.PostID, error) {
	id, err := s.postsRepo.Create(ConvertToModel(post))
	if err != nil {
		return nil, err
	}

	return &protobuf.PostID{PostID: id}, nil
}

func (s PostsService) Update(_ context.Context, post *protobuf.Post) (*emptypb.Empty, error) {
	err := s.postsRepo.Update(ConvertToModel(post))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s PostsService) DeleteByID(_ context.Context, postID *protobuf.PostID) (*emptypb.Empty, error) {
	err := s.postsRepo.DeleteByID(postID.GetPostID())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s PostsService) GetPostsBySubscriptions(_ context.Context, userID *usersProto.UserID) (*protobuf.PostArray, error) {
	posts, err := s.postsRepo.GetPostsBySubscriptions(userID.GetUserId())
	if err != nil {
		return nil, err
	}
	result := make([]*protobuf.Post, 0)
	for _, post := range posts {
		result = append(result, ConvertToProto(post))
	}

	return &protobuf.PostArray{Posts: result}, nil
}

func (s PostsService) GetLikeByUserAndPostID(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*protobuf.Like, error) {
	like, err := s.postsRepo.GetLikeByUserAndPostID(postUserIDs.UserID, postUserIDs.PostID)
	if err != nil {
		return nil, err
	}

	return &protobuf.Like{
		UserID: like.UserID,
		PostID: like.PostID,
	}, nil
}

func (s PostsService) GetAllLikesByPostID(_ context.Context, postID *protobuf.PostID) (*protobuf.Likes, error) {
	likes, err := s.postsRepo.GetAllLikesByPostID(postID.GetPostID())
	if err != nil {
		return nil, err
	}

	result := make([]*protobuf.Like, 0)
	for _, like := range likes {
		result = append(result, &protobuf.Like{
			UserID: like.UserID,
			PostID: like.PostID,
		})
	}

	return &protobuf.Likes{Likes: result}, nil
}

func (s PostsService) CreateLike(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*emptypb.Empty, error) {
	err := s.postsRepo.CreateLike(postUserIDs.UserID, postUserIDs.PostID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s PostsService) DeleteLikeByID(_ context.Context, postUserIDs *protobuf.PostUserIDs) (*emptypb.Empty, error) {
	err := s.postsRepo.DeleteLikeByID(postUserIDs.UserID, postUserIDs.PostID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s PostsService) CreateTag(_ context.Context, tagName *protobuf.TagName) (*protobuf.TagID, error) {
	id, err := s.postsRepo.CreateTag(tagName.GetTagName())
	if err != nil {
		return nil, err
	}

	return &protobuf.TagID{TagID: id}, nil
}

func (s PostsService) GetTagById(_ context.Context, tagID *protobuf.TagID) (*protobuf.Tag, error) {
	tag, err := s.postsRepo.GetTagById(tagID.GetTagID())
	if err != nil {
		return nil, err
	}

	return &protobuf.Tag{
		Id:      tag.ID,
		TagName: tag.TagName,
	}, nil
}

func (s PostsService) GetTagByName(_ context.Context, tagName *protobuf.TagName) (*protobuf.Tag, error) {
	tag, err := s.postsRepo.GetTagByName(tagName.GetTagName())
	if err != nil {
		return nil, err
	}

	return &protobuf.Tag{
		Id:      tag.ID,
		TagName: tag.TagName,
	}, nil
}

func (s PostsService) CreateDepTag(_ context.Context, tagDep *protobuf.TagDep) (*emptypb.Empty, error) {
	err := s.postsRepo.CreateDepTag(tagDep.PostID, tagDep.TagID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s PostsService) GetTagDepsByPostId(_ context.Context, postID *protobuf.PostID) (*protobuf.TagDeps, error) {
	tagDeps, err := s.postsRepo.GetTagDepsByPostId(postID.GetPostID())
	if err != nil {
		return nil, err
	}

	result := make([]*protobuf.TagDep, 0)

	for _, tagDep := range tagDeps {
		result = append(result, &protobuf.TagDep{
			PostID: tagDep.PostID,
			TagID:  tagDep.TagID,
		})
	}

	return &protobuf.TagDeps{TagDeps: result}, nil
}

func (s PostsService) DeleteDepTag(_ context.Context, tagDep *protobuf.TagDep) (*emptypb.Empty, error) {
	err := s.postsRepo.DeleteDepTag(models.TagDep{
		PostID: tagDep.GetPostID(),
		TagID:  tagDep.GetTagID(),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
