package authHandlers

import (
	"encoding/json"
	"errors"
	authErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/auth/errors"
	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/session/repository"
	userRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/repository"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type HTTPHandler struct {
	userRepo    *userRepository.Repo
	sessionRepo *sessionRepository.Repo
}

func NewHTTPHandler(userRepo *userRepository.Repo, sessionRepo *sessionRepository.Repo) *HTTPHandler {
	return &HTTPHandler{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (h *HTTPHandler) Auth(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		authErrors.Wrap(w, authErrors.ErrNoSession, err)
		return
	}

	id, ok := h.sessionRepo.Storage[session.Value]
	if !ok {
		authErrors.Wrap(w, authErrors.ErrNoSession, errors.New("failed to get session"))
		return
	}
	user, err := h.userRepo.FindByID(id)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrUserNotFound, err)
		return
	}
	info, err := json.Marshal(user)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrJSONMarshal, err)
		return
	}
	if _, err = w.Write(info); err != nil {
		authErrors.Wrap(w, authErrors.ErrResponse, err)
		return
	}
}

func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&authData)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrJSONUnmarshal, err)
		return
	}

	user, err := h.userRepo.FindByUsername(authData.Username)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrUserNotFound, err)
		return
	}

	if authData.Password != user.Password {
		authErrors.Wrap(w, authErrors.ErrInvalidLoginOrPassword, errors.New("passwords not the same"))
		return
	}

	c := h.sessionRepo.Create(user.ID)

	info, err := json.Marshal(user)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrJSONMarshal, err)
		return
	}
	http.SetCookie(w, c)
	if _, err = w.Write(info); err != nil {
		authErrors.Wrap(w, authErrors.ErrResponse, err)
		return
	}
}

func (h *HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		authErrors.Wrap(w, authErrors.ErrNoSession, err)
		return
	}

	if _, ok := h.sessionRepo.Storage[session.Value]; !ok {
		authErrors.Wrap(w, authErrors.ErrNoSession, errors.New("failed to get session"))
		return
	}

	h.sessionRepo.Remove(session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (h *HTTPHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser := model.UserDB{}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		authErrors.Wrap(w, authErrors.ErrJSONUnmarshal, err)
		return
	}

	if _, err := h.userRepo.FindByUsername(newUser.Username); err == nil {
		authErrors.Wrap(w, authErrors.ErrUserAlreadyExist, authErrors.ErrUserAlreadyExist)
		return
	}
	if _, err := h.userRepo.FindByEmail(newUser.Email); err == nil {
		authErrors.Wrap(w, authErrors.ErrEmailAlreadyExist, authErrors.ErrEmailAlreadyExist)
		return
	}

	if err = h.userRepo.Create(&newUser); err != nil {
		authErrors.Wrap(w, authErrors.ErrCreateUser, err)
		return
	}

	c := h.sessionRepo.Create(newUser.ID)

	var info []byte

	switch newUser.IsAuthor {
	case true:
		var user model.Author
		if err = copier.Copy(&user, &newUser); err != nil {
			authErrors.Wrap(w, authErrors.ErrCopy, err)
			return
		}
		if info, err = json.Marshal(user); err != nil {
			authErrors.Wrap(w, authErrors.ErrJSONMarshal, err)
			return
		}
	default:
		var user model.NotAuthor
		if err = copier.Copy(&user, &newUser); err != nil {
			authErrors.Wrap(w, authErrors.ErrCopy, err)
			return
		}
		if info, err = json.Marshal(user); err != nil {
			authErrors.Wrap(w, authErrors.ErrJSONMarshal, err)
			return
		}
	}

	http.SetCookie(w, c)
	if _, err = w.Write(info); err != nil {
		authErrors.Wrap(w, authErrors.ErrResponse, err)
		return
	}
}
