package models

type Subscription struct {
	AuthorID             uint64 `json:"author_id" db:"author_id"`
	SubscriberID         uint64 `json:"subscriber_id" db:"subscriber_id"`
	AuthorSubscriptionID uint64 `json:"author_subscription_id" db:"author_subscription_id"`
}

type AuthorSubscription struct {
	ID       uint64 `json:"id" form:"id" db:"id"`
	AuthorID uint64 `json:"author_id" form:"author_id" db:"author_id"`
	Img      string `json:"img" db:"img"`
	Tier     uint64 `json:"tier" form:"tier" db:"tier"`
	Title    string `json:"title" form:"title" db:"title"`
	Text     string `json:"text" form:"text" db:"text"`
	Price    uint64 `json:"price" form:"price" db:"price"`
}
