package models

import "time"

type Subscription struct {
	AuthorID             uint64 `json:"authorID" db:"author_id" validate:"required" example:"1"`
	SubscriberID         uint64 `json:"subscriberID" db:"subscriber_id" validate:"required" example:"2"`
	AuthorSubscriptionID uint64 `json:"authorSubscriptionID" db:"subscription_id" validate:"required" example:"1"`
}

type Follower struct {
	FollowerID  uint64    `json:"followerID" db:"follower_id" validate:"required" example:"1"`
	AuthorID    uint64    `json:"authorID" db:"author_id" validate:"required" example:"2"`
	DateCreated time.Time `json:"dateCreated" db:"date_created" validate:"required" example:"2021-01-01T00:00:00Z"`
}

type AuthorSubscription struct {
	ID       uint64 `json:"id" form:"id" db:"id" example:"1"`
	AuthorID uint64 `json:"authorID" form:"authorID" db:"author_id" example:"1"`
	Img      string `json:"img" db:"img" example:"filename.jpeg"`
	Tier     uint64 `json:"tier" form:"tier" db:"tier" validate:"required" example:"15"`
	Title    string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text     string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	Price    uint64 `json:"price" form:"price" db:"price" validate:"required" example:"2999"`

	AuthorName   string `json:"authorName,omitempty" example:"leo"`
	AuthorAvatar string `json:"authorAvatar,omitempty" example:"path/to/img"`
}
