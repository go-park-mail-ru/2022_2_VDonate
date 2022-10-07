package middlewares

import (
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserResponse(u *model.UserDB, c echo.Context) error {
	if u.IsAuthor {
		return c.JSON(http.StatusOK, model.ToAuthor(u))
	}
	return c.JSON(http.StatusOK, model.ToNonAuthor(u))
}

//
//import (
//	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/session/repository/local"
//	"net/http"
//)
//
//type Middleware struct {
//	userRepo    *postgres.Repo
//	sessionRepo *sessionRepository.Repo
//}
//
//func NewAuth(u *postgres.Repo, c *sessionRepository.Repo) *Middleware {
//	return &Middleware{userRepo: u, sessionRepo: c}
//}
//
//func (m *Middleware) LoginRequired(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		session, err := r.Cookie("session_id")
//		if err == http.ErrNoCookie {
//			return
//		}
//
//		_, ok := m.sessionRepo.Storage[session.Value]
//		if !ok {
//			http.Error(w, `no existing session`, http.StatusUnauthorized)
//			return
//		}
//
//		next.ServeHTTP(w, r)
//	})
//
//}
