package userHandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/session/repository"
	usersErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/errors"
	userRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/repository"
	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	userRepo    userRepository.RepoI
	sessionRepo sessionRepository.RepoI
}

func NewHTTPHandler(userRepo userRepository.RepoI, sessionRepo sessionRepository.RepoI) *HTTPHandler {
	return &HTTPHandler{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (h *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		usersErrors.Wrap(w, usersErrors.ErrConvertID, err)
		return
	}
	user, err := h.userRepo.FindByID(uint(id))
	if err != nil {
		usersErrors.Wrap(w, usersErrors.ErrUserNotFound, err)
		return
	}

	info, err := json.Marshal(user)
	if err != nil {
		usersErrors.Wrap(w, usersErrors.ErrJSONMarshal, err)
		return
	}
	if _, err = w.Write(info); err != nil {
		usersErrors.Wrap(w, usersErrors.ErrResponse, err)
		return
	}
}
