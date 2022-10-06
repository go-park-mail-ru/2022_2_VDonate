package middleware

import (
	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/session/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/repository"
	"net/http"
)

type Middleware struct {
	userRepo    *userRepository.Repo
	sessionRepo *sessionRepository.Repo
}

func NewAuth(u *userRepository.Repo, c *sessionRepository.Repo) *Middleware {
	return &Middleware{userRepo: u, sessionRepo: c}
}

func (m *Middleware) LoginRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			return
		}

		_, ok := m.sessionRepo.Storage[session.Value]
		if !ok {
			http.Error(w, `no existing session`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
