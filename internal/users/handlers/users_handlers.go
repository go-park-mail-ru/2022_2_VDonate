package user_http

import (
	"encoding/json"
	cookie_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/cookie"
	user_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHTTPHandler struct {
	userRepo   *user_repo.Repo
	cookieRepo *cookie_repo.Repo
}

func New(userRepo *user_repo.Repo, cookieRepo *cookie_repo.Repo) *UserHTTPHandler {
	return &UserHTTPHandler{userRepo: userRepo, cookieRepo: cookieRepo}
}

func (h *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusInternalServerError)
		return
	}
	user, err := h.userRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	info, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(info); err != nil {
		http.Error(w, "failed to write user info", http.StatusInternalServerError)
		return
	}
}
