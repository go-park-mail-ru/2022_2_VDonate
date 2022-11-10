package models

import (
	"time"
)

type Cookie struct {
	Value   string    `json:"value" db:"value"`
	UserID  uint64    `json:"userID" db:"user_id"`
	Expires time.Time `json:"expire_date" db:"expire_date"`
}
