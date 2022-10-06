package sessionRepository

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"net/http"
	"sync"
	"time"
)

type Repo struct {
	Storage map[string]uint
	m       sync.Mutex
}

func New() *Repo {
	return &Repo{Storage: make(map[string]uint)}
}

func (r *Repo) Create(id uint) *http.Cookie {
	CID := utils.RandStringRunes(32)

	r.m.Lock()
	r.Storage[CID] = id
	r.m.Unlock()

	return &http.Cookie{
		Name:    "session_id",
		Value:   CID,
		Expires: time.Now().Add(30 * 24 * time.Hour),
	}
}

func (r *Repo) Remove(CID string) {
	delete(r.Storage, CID)
}
