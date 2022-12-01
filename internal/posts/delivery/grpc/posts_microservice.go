package postsMicroservice

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	grpcPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/grpc"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type PostsMicroservice struct {
	client protobuf.PostsClient
}

func New(c protobuf.PostsClient) domain.PostsMicroservice {
	return &PostsMicroservice{
		client: c,
	}
}

func (m PostsMicroservice) GetAllByUserID(userID uint64) ([]models.Post, error) {
	posts, err := m.client.GetAllByUserID(context.Background(), &userProto.UserID{UserId: userID})
	if err != nil {
		return nil, err
	}

	result := make([]models.Post, 0)

	for _, p := range posts.GetPosts() {
		result = append(result, grpcPosts.ConvertToModel(p))
	}

	return result, nil
}

func (m PostsMicroservice) GetPostByID(postID uint64) (models.Post, error) {
	post, err := m.client.GetPostByID(context.Background(), &protobuf.PostID{
		PostID: postID,
	})
	if err != nil {
		return models.Post{}, err
	}

	return grpcPosts.ConvertToModel(post), nil
}

func (m PostsMicroservice) Create(post models.Post) (uint64, error) {
	id, err := m.client.Create(context.Background(), grpcPosts.ConvertToProto(post))
	if err != nil {
		return 0, err
	}

	return id.GetPostID(), nil
}

func (m PostsMicroservice) Update(post models.Post) error {
	_, err := m.client.Update(context.Background(), grpcPosts.ConvertToProto(post))

	return err
}

func (m PostsMicroservice) DeleteByID(postID uint64) error {
	_, err := m.client.DeleteByID(context.Background(), &protobuf.PostID{PostID: postID})

	return err
}

func (m PostsMicroservice) GetPostsBySubscriptions(userID uint64) ([]models.Post, error) {
	posts, err := m.client.GetPostsBySubscriptions(context.Background(), &userProto.UserID{UserId: userID})
	if err != nil {
		return nil, err
	}

	result := make([]models.Post, 0)

	for _, p := range posts.GetPosts() {
		result = append(result, grpcPosts.ConvertToModel(p))
	}

	return result, nil
}

func (m PostsMicroservice) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	like, err := m.client.GetLikeByUserAndPostID(context.Background(), &protobuf.PostUserIDs{
		PostID: postID,
		UserID: userID,
	})
	if err != nil {
		return models.Like{}, err
	}

	return models.Like{
		UserID: like.GetUserID(),
		PostID: like.GetPostID(),
	}, nil
}

func (m PostsMicroservice) GetAllLikesByPostID(postID uint64) ([]models.Like, error) {
	likes, err := m.client.GetAllLikesByPostID(context.Background(), &protobuf.PostID{
		PostID: postID,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.Like, 0)

	for _, like := range likes.GetLikes() {
		result = append(result, models.Like{
			UserID: like.GetUserID(),
			PostID: like.GetPostID(),
		})
	}

	return result, nil
}

func (m PostsMicroservice) CreateLike(userID, postID uint64) error {
	_, err := m.client.CreateLike(context.Background(), &protobuf.PostUserIDs{
		PostID: postID,
		UserID: userID,
	})

	return err
}

func (m PostsMicroservice) DeleteLikeByID(userID, postID uint64) error {
	_, err := m.client.DeleteLikeByID(context.Background(), &protobuf.PostUserIDs{
		PostID: postID,
		UserID: userID,
	})

	return err
}

func (m PostsMicroservice) CreateTag(tagName string) (uint64, error) {
	id, err := m.client.CreateTag(context.Background(), &protobuf.TagName{
		TagName: tagName,
	})
	if err != nil {
		return 0, err
	}

	return id.GetTagID(), nil
}

func (m PostsMicroservice) GetTagById(tagID uint64) (models.Tag, error) {
	tag, err := m.client.GetTagById(context.Background(), &protobuf.TagID{
		TagID: tagID,
	})
	if err != nil {
		return models.Tag{}, err
	}

	return models.Tag{
		ID:      tag.GetId(),
		TagName: tag.GetTagName(),
	}, nil
}

func (m PostsMicroservice) GetTagByName(tagName string) (models.Tag, error) {
	tag, err := m.client.GetTagByName(context.Background(), &protobuf.TagName{
		TagName: tagName,
	})
	if err != nil {
		return models.Tag{}, err
	}

	return models.Tag{
		ID:      tag.GetId(),
		TagName: tag.GetTagName(),
	}, nil
}

func (m PostsMicroservice) CreateDepTag(postID, tagID uint64) error {
	_, err := m.client.CreateDepTag(context.Background(), &protobuf.TagDep{
		PostID: postID,
		TagID:  tagID,
	})

	return err
}

func (m PostsMicroservice) GetTagDepsByPostId(postID uint64) ([]models.TagDep, error) {
	tagDeps, err := m.client.GetTagDepsByPostId(context.Background(), &protobuf.PostID{
		PostID: postID,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.TagDep, 0)

	for _, tagDep := range tagDeps.GetTagDeps() {
		result = append(result, models.TagDep{
			PostID: tagDep.GetPostID(),
			TagID:  tagDep.GetTagID(),
		})
	}

	return result, nil
}

func (m PostsMicroservice) DeleteDepTag(tagDep models.TagDep) error {
	_, err := m.client.DeleteDepTag(context.Background(), &protobuf.TagDep{
		PostID: tagDep.PostID,
		TagID:  tagDep.TagID,
	})

	return err
}