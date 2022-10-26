package httpPosts

import (
	"net/http"
	"strconv"
	"strings"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	postsUseCase domain.PostsUseCase
	usersUseCase domain.UsersUseCase
	imageUseCase domain.ImageUseCase

	bucket string
}

func NewHandler(p domain.PostsUseCase, u domain.UsersUseCase, i domain.ImageUseCase, bucket string) *Handler {
	return &Handler{
		postsUseCase: p,
		usersUseCase: u,
		imageUseCase: i,

		bucket: bucket,
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

	for _, post := range allPosts {
		url, errGetImage := h.imageUseCase.GetImage(h.bucket, post.Img)
		if errGetImage != nil {
			return errGetImage
		}
		post.Img = url.String()
	}

	return c.JSON(http.StatusOK, allPosts)
}

func (h *Handler) GetPost(c echo.Context) error {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := h.postsUseCase.GetPostByID(postID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	url, err := h.imageUseCase.GetImage(h.bucket, post.Img)
	if err != nil {
		return err
	}
	post.Img = url.String()

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
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	file.Filename = uuid.New().String() + file.Filename[strings.IndexByte(file.Filename, '.'):]

	var prevPost models.Post

	if err = c.Bind(&prevPost); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.imageUseCase.CreateImage(file, h.bucket); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	prevPost.ID = postID
	if err = h.postsUseCase.Update(prevPost); err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
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

	file, err := c.FormFile("file")
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}
	file.Filename = uuid.New().String() + file.Filename[strings.IndexByte(file.Filename, '.'):]

	var post models.Post

	if err = c.Bind(&post); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	if err = h.imageUseCase.CreateImage(file, h.bucket); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	post.Img = file.Filename
	post.UserID = user.ID
	if err = h.postsUseCase.Create(post); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
