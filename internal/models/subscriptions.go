package models

type Subscription struct {
	AuthorID             uint64 `json:"author_id" db:"author_id" validate:"required" example:"1"`
	SubscriberID         uint64 `json:"subscriber_id" db:"subscriber_id" validate:"required" example:"2"`
	AuthorSubscriptionID uint64 `json:"author_subscription_id" db:"author_subscription_id" validate:"required" example:"1"`
}

type AuthorSubscription struct {
	ID       uint64 `json:"id" form:"id" db:"id" example:"1"`
	AuthorID uint64 `json:"author_id" form:"author_id" db:"author_id" example:"1"`
	Img      string `json:"img" db:"img" example:"filename.jpeg"`
	Tier     uint64 `json:"tier" form:"tier" db:"tier" validate:"required" example:"15"`
	Title    string `json:"title" form:"title" db:"title" validate:"required" example:"some title"`
	Text     string `json:"text" form:"text" db:"text" validate:"required" example:"some text"`
	Price    uint64 `json:"price" form:"price" db:"price" validate:"required" example:"2999"`
}
