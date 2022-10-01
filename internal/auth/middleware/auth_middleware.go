package auth_middleware

import (
	cookie_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/cookie"
	user_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type MiddleWare struct {
	userRepo   *user_repo.Repo
	cookieRepo *cookie_repo.Repo
}

func New(u *user_repo.Repo, c *cookie_repo.Repo) *MiddleWare {
	return &MiddleWare{userRepo: u, cookieRepo: c}
}

func (m *MiddleWare) LoginRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			return
		}

		_, ok := m.cookieRepo.Storage[session.Value]
		if !ok {
			http.Error(w, `no session`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}

func CORS(allowOrigin []string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			for _, o := range allowOrigin {
				if o == origin {
					r.Header.Set("Access-Control-Allow-Origin", o)
					r.Header.Set("Access-Control-Allow-Credentials", "true")

					next.ServeHTTP(w, r)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
