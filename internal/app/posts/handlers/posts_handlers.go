package postsHandlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	postsErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/posts/errors"
	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/posts/repository"
	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	postsRepo postsRepository.RepoI
}

func NewHTPPHandler(p postsRepository.RepoI) *HTTPHandler {
	return &HTTPHandler{postsRepo: p}
}

func (h *HTTPHandler) Posts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		postsErrors.Wrap(w, postsErrors.ErrBadRequest, postsErrors.ErrBadRequest)
		return
	}
	posts, err := h.postsRepo.GetPostsByID(uint(id))

	postsInfo, err := json.Marshal(posts)
	if err != nil {
		postsErrors.Wrap(w, postsErrors.ErrJSONMarshal, errors.New("failed to marshal posts"))
		return
	}

	if _, err := w.Write(postsInfo); err != nil {
		postsErrors.Wrap(w, postsErrors.ErrResponse, errors.New("failed to write json"))
		return
	}
}
