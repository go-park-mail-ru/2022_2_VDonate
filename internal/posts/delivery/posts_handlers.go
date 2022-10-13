package httpPosts

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	postsUseCase posts.UseCase
}

func NewHandler(p posts.UseCase) *Handler {
	return &Handler{postsUseCase: p}
}

func (h *Handler) GetPosts(c echo.Context) error {
	id, err := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	allPosts, err := h.postsUseCase.GetPostsByUserID(id)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, allPosts)
}

func (h *Handler) DeletePost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("post_id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	if err = h.postsUseCase.DeleteByID(id); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) PutPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("post_id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}

	var prevPost *models.PostDB
	if err = c.Bind(&prevPost); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}

	if prevPost, err = h.postsUseCase.Update(id, prevPost); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, prevPost)
}

func (h *Handler) CreatePosts(c echo.Context) error {
	id, err := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	var post *models.PostDB
	if err := c.Bind(&post); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}
	post.UserID = id
	if post, err = h.postsUseCase.Create(post); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, post)
}
