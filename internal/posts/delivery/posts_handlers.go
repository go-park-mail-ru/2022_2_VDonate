package httpPosts

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
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

// GetPosts godoc
// @Summary     Get posts
// @Description Get posts with provided filters
// @ID          get_posts
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       user_id query    integer        true "User ID"
// @Success     200     {object} []models.Post  "Posts were successfully received"
// @Failure     400     {object} echo.HTTPError "Bad request"
// @Failure     401     {object} echo.HTTPError "No session provided"
// @Failure     404     {object} echo.HTTPError "User not found"
// @Failure     500     {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts [get]
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
			return utils.WrapEchoError(domain.ErrInternal, errGetImage)
		}
		post.Img = url.String()
	}

	return c.JSON(http.StatusOK, allPosts)
}

// GetPost godoc
// @Summary     Get single post
// @Description Get single post by post id
// @ID          get_post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id  path     integer        true "Post ID"
// @Success     200 {object} models.Post    "Post was successfully received"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id} [get]
func (h *Handler) GetPost(c echo.Context) error {
	postID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := h.postsUseCase.GetPostByID(postID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	url, err := h.imageUseCase.GetImage(h.bucket, post.Img)
	if err != nil {
		return utils.WrapEchoError(domain.ErrInternal, err)
	}
	post.Img = url.String()

	return c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary     Delete post
// @Description Delete post by post id
// @ID          delete_post
// @Tags        posts
// @Accept      json
// @Produce     json
// @Param       id  path     integer            true "Post ID"
// @Success     200 {object} models.EmptyStruct "Post was successfully deleted"
// @Failure     400 {object} echo.HTTPError     "Bad request"
// @Failure     401 {object} echo.HTTPError     "No session provided"
// @Failure     403 {object} echo.HTTPError     "Not a creator of post"
// @Failure     404 {object} echo.HTTPError     "Post not found"
// @Failure     500 {object} echo.HTTPError     "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id} [delete]
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

// PutPost godoc
// @Summary     Update post
// @Description Update post by post id
// @ID          update_post
// @Tags        posts
// @Accept      mpfd
// @Produce     json
// @Param       id   path     integer              true  "Post ID"
// @Param       post formData models.PostMpfd      true  "New Post"
// @Param       file formData file                 false "Uploaded file"
// @Success     200  {object} models.ResponseImage "Post was successfully updated"
// @Failure     400  {object} echo.HTTPError       "Bad request"
// @Failure     401  {object} echo.HTTPError       "No session provided"
// @Failure     403  {object} echo.HTTPError       "Not a creator of post"
// @Failure     404  {object} echo.HTTPError       "Post not found"
// @Failure     500  {object} echo.HTTPError       "Internal error / failed to create image"
// @Security    ApiKeyAuth
// @Router      /posts/{id} [put]
func (h *Handler) PutPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var prevPost models.Post

	if err = c.Bind(&prevPost); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	newFile, err := h.imageUseCase.CreateImage(file, h.bucket)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	prevPost.Img = newFile
	prevPost.ID = postID
	if err = h.postsUseCase.Update(prevPost); err != nil {
		return utils.WrapEchoError(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImage{
		ImgPath: prevPost.Img,
	})
}

// CreatePost godoc
// @Summary     Create post
// @Description Create post by user cookie
// @ID          create_post
// @Tags        posts
// @Accept      mpfd
// @Param       post formData models.PostMpfd true  "New Post"
// @Param       file formData file            false "Uploaded file"
// @Produce     json
// @Success     200 {object} models.ResponseImage "Post was successfully created"
// @Failure     400 {object} echo.HTTPError       "Bad request"
// @Failure     401 {object} echo.HTTPError       "No session provided"
// @Failure     403 {object} echo.HTTPError       "Not a creator of post"
// @Failure     404 {object} echo.HTTPError       "Post not found"
// @Failure     500 {object} echo.HTTPError       "Internal error / failed to create post"
// @Security    ApiKeyAuth
// @Router      /posts [post]
func (h *Handler) CreatePost(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	file, err := images.GetFileFromContext(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	var post models.Post

	if err = c.Bind(&post); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	newImage, err := h.imageUseCase.CreateImage(file, h.bucket)
	if err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	post.Img = newImage
	post.UserID = user.ID
	if err = h.postsUseCase.Create(post); err != nil {
		return utils.WrapEchoError(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImage{
		ImgPath: post.Img,
	})
}
