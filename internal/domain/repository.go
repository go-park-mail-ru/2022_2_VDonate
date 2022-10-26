package domain

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"mime/multipart"
	"net/url"
)

type AuthRepository interface {
	GetBySessionID(sessionID string) (*models.Cookie, error)
	GetByUserID(id uint64) (*models.Cookie, error)
	GetByUsername(username string) (*models.Cookie, error)
	CreateSession(cookie models.Cookie) (*models.Cookie, error)
	DeleteBySessionID(sessionID string) error
	DeleteByUserID(id uint64) error
	Close() error
}

type PostsRepository interface {
	GetAllByUserID(userID uint64) ([]*models.Post, error)
	GetPostByID(postID uint64) (*models.Post, error)
	Create(post models.Post) error
	Update(post models.Post) error
	DeleteByID(postID uint64) error
	Close() error
}

type SubscribersRepository interface {
	GetSubscribers(authorID uint64) ([]uint64, error)
	Subscribe(subscription models.Subscription) error
	Unsubscribe(userID, authorID uint64) error
}

type SubscriptionsRepository interface {
	GetSubscriptionsByAuthorID(authorID uint64) ([]*models.AuthorSubscription, error)
	GetSubscriptionsByID(ID uint64) (*models.AuthorSubscription, error)
	AddSubscription(sub models.AuthorSubscription) error
	UpdateSubscription(sub *models.AuthorSubscription) error
	DeleteSubscription(subID uint64) error
}

type UsersRepository interface {
	Create(user *models.User) error
	GetByUsername(username string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetBySessionID(sessionID string) (*models.User, error)
	GetUserByPostID(postID uint64) (*models.User, error)
	Update(user *models.User) error
	DeleteByID(id uint64) error
	Close() error
}

type ImagesRepository interface {
	CreateImage(image *multipart.FileHeader, bucket string) (string, error)
	GetImage(bucket string, filename string) (*url.URL, error)
	GetPermanentImage(bucket string, filename string) (string, error)
}
