package auth_http

import (
	"encoding/json"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models/user"
	cookie_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/cookie"
	user_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"net/http"
	"time"
)

type AuthHTTPHandler struct {
	userRepo   *user_repo.Repo
	cookieRepo *cookie_repo.Repo
}

func New(userRepo *user_repo.Repo, cookieRepo *cookie_repo.Repo) *AuthHTTPHandler {
	return &AuthHTTPHandler{userRepo: userRepo, cookieRepo: cookieRepo}
}

func (h *AuthHTTPHandler) Auth(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no session`, http.StatusUnauthorized)
		return
	}

	id, ok := h.cookieRepo.Storage[session.Value]
	if !ok {
		http.Error(w, `no session`, http.StatusUnauthorized)
		return
	}
	user, err := h.userRepo.FindByID(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
	}
	info, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "json marshal error", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(info); err != nil {
		http.Error(w, "response error", http.StatusInternalServerError)
		return
	}
}

func (h *AuthHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var authData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&authData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.userRepo.FindByUsername(authData.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if authData.Password != user.Password {
		http.Error(w, `bad request`, http.StatusBadRequest)
		return
	}

	c := h.cookieRepo.Create(user.ID)

	info, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, c)
	if _, err = w.Write(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *AuthHTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no session`, http.StatusUnauthorized)
		return
	}

	if _, ok := h.cookieRepo.Storage[session.Value]; !ok {
		http.Error(w, `no session`, http.StatusUnauthorized)
		return
	}

	h.cookieRepo.Remove(session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (h *AuthHTTPHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
		Email     string `json:"email"`
		Phone     uint   `json:"phone"`
		IsAuthor  bool   `json:"is_author"`
		About     string `json:"about"`
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := h.userRepo.FindByUsername(newUser.Username); err == nil {
		http.Error(w, "username is already exists", http.StatusConflict)
		return
	}
	if _, err := h.userRepo.FindByEmail(newUser.Email); err == nil {
		http.Error(w, "email is already exists", http.StatusConflict)
		return
	}

	var user *model.User
	if newUser.IsAuthor {
		user = &model.User{
			Username:  newUser.Username,
			Password:  newUser.Password,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Avatar:    newUser.Avatar,
			Email:     newUser.Email,
			Phone:     newUser.Phone,
			IsAuthor:  newUser.IsAuthor,
			About:     newUser.About,
		}
	} else {
		user = &model.User{
			Username:  newUser.Username,
			Password:  newUser.Password,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Avatar:    newUser.Avatar,
			Email:     newUser.Email,
			Phone:     newUser.Phone,
			IsAuthor:  newUser.IsAuthor,
		}
	}

	if err = h.userRepo.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := h.cookieRepo.Create(user.ID)

	info, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, c)
	if _, err = w.Write(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
