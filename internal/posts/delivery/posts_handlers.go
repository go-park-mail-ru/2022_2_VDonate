package httpPosts

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	postsDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/errors"
	usersDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/users"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	postsUseCase postsDomain.UseCase
	usersUseCase usersDomain.UseCase
}

func NewHandler(p postsDomain.UseCase, u usersDomain.UseCase) *Handler {
	return &Handler{
		postsUseCase: p,
		usersUseCase: u,
	}
}

func (h *Handler) GetPosts(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrNoSession, err)
	}
	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrNoSession, err)
	}

	allPosts, err := h.postsUseCase.GetPostsByUserID(user.ID)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	return c.JSON(http.StatusOK, allPosts)
}

func (h *Handler) GetPost(c echo.Context) error {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := h.postsUseCase.GetPostByID(postID)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	return c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	if err = h.postsUseCase.DeleteByID(postID); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) PutPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	var prevPost models.PostDB
	if err := c.Bind(&prevPost); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}

	prevPost.ID = postID
	post, err := h.postsUseCase.Update(prevPost)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePosts(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrNoSession, err)
	}
	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrNoSession, err)
	}
	var post models.PostDB
	if err := c.Bind(&post); err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrBadRequest, err)
	}
	post.UserID = user.ID
	newPost, err := h.postsUseCase.Create(post)
	if err != nil {
		return postsErrors.Wrap(c, postsErrors.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, newPost)
}
