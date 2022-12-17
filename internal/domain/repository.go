package domain

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthRepository interface {
	GetBySessionID(sessionID string) (models.Cookie, error)
	GetByUserID(id uint64) (models.Cookie, error)
	GetByUsername(username string) (models.Cookie, error)
	CreateSession(cookie models.Cookie) (models.Cookie, error)
	DeleteBySessionID(sessionID string) error
	DeleteByUserID(id uint64) error
	Close() error
}

type PostsRepository interface {
	GetAllByUserID(authorID uint64) ([]models.Post, error)
	GetPostByID(postID uint64) (models.Post, error)
	Create(post models.Post) (models.Post, error)
	Update(post models.Post) error
	DeleteByID(postID uint64) error
	Close() error
	GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error)
	GetAllLikesByPostID(postID uint64) ([]models.Like, error)
	CreateLike(userID, postID uint64) error
	DeleteLikeByID(userID, postID uint64) error
	GetPostsBySubscriptions(userID uint64) ([]models.Post, error)
	CreateTag(tagName string) (uint64, error)
	CreateDepTag(postID, tagID uint64) error
	DeleteDepTag(tagDep models.TagDep) error
	GetTagById(tagID uint64) (models.Tag, error)
	GetTagDepsByPostId(postID uint64) ([]models.TagDep, error)
	GetTagByName(tagName string) (models.Tag, error)
	CreateComment(comment models.Comment) (uint64, *timestamppb.Timestamp, error)
	GetCommentByID(commentID uint64) (models.Comment, error)
	GetCommentsByPostId(postID uint64) ([]models.Comment, error)
	UpdateComment(comment models.Comment) error
	DeleteCommentByID(commentID uint64) error
}

type SubscribersRepository interface {
	GetSubscribers(authorID uint64) ([]uint64, error)
	Subscribe(subscription models.Subscription) error
	Unsubscribe(userID, authorID uint64) error
}

type SubscriptionsRepository interface {
	GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error)
	GetSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error)
	GetSubscriptionByUserAndAuthorID(userID, authorID uint64) (models.AuthorSubscription, error)
	GetSubscriptionByID(ID uint64) (models.AuthorSubscription, error)
	AddSubscription(sub models.AuthorSubscription) (uint64, error)
	UpdateSubscription(sub models.AuthorSubscription) error
	DeleteSubscription(subID uint64) error
}

type UsersRepository interface {
	Create(user models.User) (uint64, error)
	GetByUsername(username string) (models.User, error)
	GetByID(id uint64) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetBySessionID(sessionID string) (models.User, error)
	GetUserByPostID(postID uint64) (models.User, error)
	Update(user models.User) error
	DeleteByID(id uint64) error
	GetAuthorByUsername(username string) ([]models.User, error)
	GetAllAuthors() ([]models.User, error)
	Close() error
}

type DonatesRepository interface {
	SendDonate(donate models.Donate) (models.Donate, error)
	GetDonatesByUserID(userID uint64) ([]models.Donate, error)
	GetDonateByID(donateID uint64) (models.Donate, error)
}

type ImagesRepository interface {
	CreateOrUpdateImage(filename string, file []byte, size int64, oldFilename string) (string, error)
	GetPermanentImage(filename string) (string, error)
}

type NotificationsRepository interface {
	GetNotificationsByUserID(userID uint64) ([]models.Notification, error)
	DeleteNotificationByUserID(userID uint64) error
}
