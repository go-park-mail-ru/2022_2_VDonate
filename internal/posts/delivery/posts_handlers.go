package httpPosts

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	postsUseCase domain.PostsUseCase
	usersUseCase domain.UsersUseCase
}

func NewHandler(p domain.PostsUseCase, u domain.UsersUseCase) *Handler {
	return &Handler{
		postsUseCase: p,
		usersUseCase: u,
	}
}

func (h *Handler) GetPosts(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}
	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadSession, err)
	}

	allPosts, err := h.postsUseCase.GetPostsByUserID(user.ID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, allPosts)
}

func (h *Handler) GetPost(c echo.Context) error {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := h.postsUseCase.GetPostByID(postID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}
	return c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	if err = h.postsUseCase.DeleteByID(postID); err != nil {
		return utils.WrapEchoError(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) PutPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	var prevPost models.Post
	if err := c.Bind(&prevPost); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	prevPost.ID = postID
	post, err := h.postsUseCase.Update(prevPost)
	if err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePosts(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}
	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	post.UserID = user.ID
	newPost, err := h.postsUseCase.Create(post)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}
	return c.JSON(http.StatusOK, newPost)
}
