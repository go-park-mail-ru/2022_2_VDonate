package sessionRepository

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
)

type RepoI interface {
	Create(id uint) *http.Cookie
	Remove(CID string)
	GetId(CID string) (uint, bool)
}

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

func (r *Repo) GetId(CID string) (uint, bool) {
	r.m.Lock()
	id, err := r.Storage[CID]
	r.m.Unlock()
	return id, err
}
