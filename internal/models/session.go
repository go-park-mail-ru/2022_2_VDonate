package models

import "time"

type Cookie struct {
	Value   string    `json:"value" db:"value"`
	UserID  uint      `json:"user_id" db:"user_id"`
	Expires time.Time `json:"expire_date" db:"expire_date"`
}
