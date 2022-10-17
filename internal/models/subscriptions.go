package models

type Subscription struct {
	AuthorID             uint64 `json:"author_id" db:"author_id"`
	SubscriberID         uint64 `json:"subscriber_id" db:"subscriber_id"`
	AuthorSubscriptionID uint64 `json:"author_subscription_id" db:"author_subscription_id"`
}

type AuthorSubscription struct {
	ID       uint64 `json:"id" db:"id"`
	AuthorID uint64 `json:"author_id" db:"author_id"`
	Tier     uint64 `json:"tier" db:"tier"`
	Text     string `json:"text" db:"text"`
	Price    uint64 `json:"price" db:"price"`
}
