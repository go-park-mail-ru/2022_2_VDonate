package models

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"net/http"
	"time"
)

const (
	cookieName = "session_id"
)

type Cookie struct {
	Value   string    `json:"value" db:"value"`
	UserID  uint64    `json:"user_id" db:"user_id"`
	Expires time.Time `json:"expire_date" db:"expire_date"`
}

func CreateCookie(id uint64) *Cookie {
	return &Cookie{
		UserID:  id,
		Value:   utils.RandStringRunes(32),
		Expires: time.Now().AddDate(0, 1, 0),
	}
}

func MakeHTTPCookieFromValue(value string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    value,
		Expires:  time.Now().AddDate(0, 1, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
