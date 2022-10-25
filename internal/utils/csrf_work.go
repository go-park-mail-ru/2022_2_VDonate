package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type HashToken struct {
	Secret []byte
}

func NewHMACHashToken(secret string) *HashToken {
	return &HashToken{Secret: []byte(secret)}
}

func (tk *HashToken) CreateCSRF(s *models.Cookie, tokenExpTime int64) (string, error) {
	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d:%d", s.Value, s.UserID, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token, nil
}

func (tk *HashToken) CheckCSRF(s *models.Cookie, inputToken string) (bool, error) {
	tokenData := strings.Split(inputToken, ":")
	if len(tokenData) != 2 {
		return false, errors.New("bad token data")
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false, errors.New("bad token time")
	}

	if tokenExp < time.Now().Unix() {
		return false, errors.New("token expired")
	}

	h := hmac.New(sha256.New, []byte(tk.Secret))
	data := fmt.Sprintf("%s:%d:%d", s.Value, s.UserID, tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false, errors.New("cand hex decode token")
	}

	return hmac.Equal(messageMAC, expectedMAC), nil
}
