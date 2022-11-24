package httpPosts

import (
	"errors"
	"net/http"
	"strconv"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	postsUseCase domain.PostsUseCase
	usersUseCase domain.UsersUseCase
	imageUseCase domain.ImageUseCase
}

func NewHandler(p domain.PostsUseCase, u domain.UsersUseCase, i domain.ImageUseCase) *Handler {
	return &Handler{
		postsUseCase: p,
		usersUseCase: u,
		imageUseCase: i,
	}
}

// GetPosts godoc
// @Summary     Get posts
// @Description Get posts with provided filters
// @ID          get_posts
// @Tags        posts
// @Produce     json
// @Param       user_id query    integer        true "User ID"
// @Param       filter  query    string         true "filter to use to get posts. Filters: subscriptions"
// @Success     200     {object} []models.Post  "Posts were successfully received"
// @Failure     400     {object} echo.HTTPError "Bad request"
// @Failure     401     {object} echo.HTTPError "No session provided"
// @Failure     404     {object} echo.HTTPError "User not found"
// @Failure     500     {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts [get]
func (h *Handler) GetPosts(c echo.Context) error {
	var allPosts []models.Post
	var user models.User
	var userID uint64
	var err error

	id := c.QueryParam("user_id")
	filter := c.QueryParam("filter")

	switch {
	case len(id) != 0:
		if userID, err = strconv.ParseUint(id, 10, 64); err != nil {
			return errorHandling.WrapEcho(domain.ErrBadRequest, err)
		}
		if user, err = h.usersUseCase.GetByID(userID); err != nil {
			return errorHandling.WrapEcho(domain.ErrNotFound, err)
		}
		if allPosts, err = h.postsUseCase.GetPostsByUserID(userID); err != nil {
			return errorHandling.WrapEcho(domain.ErrNotFound, err)
		}
	case len(filter) != 0:
		cookie, errU := httpAuth.GetCookie(c)
		if errU != nil {
			return errorHandling.WrapEcho(domain.ErrNoSession, err)
		}

		if user, errU = h.usersUseCase.GetBySessionID(cookie.Value); errU != nil {
			return errorHandling.WrapEcho(domain.ErrNoSession, errU)
		}
		if allPosts, err = h.postsUseCase.GetPostsByFilter(filter, user.ID); err != nil {
			return errorHandling.WrapEcho(domain.ErrNotFound, err)
		}
	default:
		return errorHandling.WrapEcho(domain.ErrBadRequest, errors.New("bad request"))
	}

	if len(allPosts) == 0 {
		return c.JSON(http.StatusOK, make([]models.Post, 0))
	}

	for i, post := range allPosts {
		if allPosts[i].Img, err = h.imageUseCase.GetImage(post.Img); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
		if allPosts[i].Author.ImgPath, err = h.imageUseCase.GetImage(post.Author.ImgPath); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
		if allPosts[i].LikesNum, err = h.postsUseCase.GetLikesNum(post.ID); err != nil {
			return errorHandling.WrapEcho(domain.ErrInternal, err)
		}
		allPosts[i].IsLiked = h.postsUseCase.IsPostLiked(user.ID, post.ID)
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
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	post, err := h.postsUseCase.GetPostByID(postID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	if post.Img, err = h.imageUseCase.GetImage(post.Img); err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	if post.Author.ImgPath, err = h.imageUseCase.GetImage(post.Author.ImgPath); err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	post.LikesNum, err = h.postsUseCase.GetLikesNum(postID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	post.IsLiked = h.postsUseCase.IsPostLiked(user.ID, postID)

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
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if err = h.postsUseCase.DeleteByID(postID); err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}

// PutPost godoc
// @Summary     Update post
// @Description Update post by post id
// @ID          update_post
// @Tags        posts
// @Accept      mpfd
// @Produce     json
// @Param       id   path     integer                   true  "Post ID"
// @Param       post formData models.PostMpfd           true  "New Post"
// @Param       file formData file                      false "Uploaded file"
// @Success     200  {object} models.ResponseImagePosts "Post was successfully updated"
// @Failure     400  {object} echo.HTTPError            "Bad request"
// @Failure     401  {object} echo.HTTPError            "No session provided"
// @Failure     403  {object} echo.HTTPError            "Not a creator of post"
// @Failure     404  {object} echo.HTTPError            "Post not found"
// @Failure     500  {object} echo.HTTPError            "Internal error / failed to create image"
// @Security    ApiKeyAuth
// @Router      /posts/{id} [put]
func (h *Handler) PutPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	var prevPost models.Post

	if err = c.Bind(&prevPost); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)

	if file != nil && !errors.Is(err, http.ErrMissingFile) {
		if prevPost.Img, err = h.imageUseCase.CreateImage(file); err != nil {
			return errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	if err = h.postsUseCase.Update(prevPost, postID); err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}

	tmpURL, err := h.imageUseCase.GetImage(prevPost.Img)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImagePosts{
		PostID:  postID,
		ImgPath: tmpURL,
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
// @Success     200 {object} models.ResponseImagePosts "Post was successfully created"
// @Failure     400 {object} echo.HTTPError            "Bad request"
// @Failure     401 {object} echo.HTTPError            "No session provided"
// @Failure     403 {object} echo.HTTPError            "Not a creator of post"
// @Failure     404 {object} echo.HTTPError            "Post not found"
// @Failure     500 {object} echo.HTTPError            "Internal error / failed to create post"
// @Security    ApiKeyAuth
// @Router      /posts [post]
func (h *Handler) CreatePost(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var post models.Post

	if err = c.Bind(&post); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	file, err := images.GetFileFromContext(c)
	if file != nil && !errors.Is(err, http.ErrMissingFile) {
		if post.Img, err = h.imageUseCase.CreateImage(file); err != nil {
			return errorHandling.WrapEcho(domain.ErrCreate, err)
		}
	}

	id, err := h.postsUseCase.Create(post, user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	tmpURL, err := h.imageUseCase.GetImage(post.Img)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImagePosts{
		PostID:  id,
		ImgPath: tmpURL,
	})
}

// GetLikes godoc
// @Summary     Get likes
// @Description Get all likes by post id
// @ID          get_posts_likes
// @Tags        posts
// @Param       id path integer true "Post id"
// @Produce     json
// @Success     200 {object} []models.Like  "Likes were successfully recieved"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/likes [get]
func (h *Handler) GetLikes(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	allLikes, err := h.postsUseCase.GetLikesByPostID(postID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	if len(allLikes) == 0 {
		return c.JSON(http.StatusOK, []models.Like{})
	}
	return c.JSON(http.StatusOK, allLikes)
}

// CreateLike godoc
// @Summary     Create like
// @Description Create like on post
// @ID          create_like
// @Tags        posts
// @Produce     json
// @Success     200 {object} integer        "Likes were successfully create"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/likes [post]
func (h *Handler) CreateLike(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	err = h.postsUseCase.LikePost(user.ID, postID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrLikeExist, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}

// DeleteLike godoc
// @Summary     Delete like
// @Description Create like on post
// @ID          delete_like
// @Tags        posts
// @Param       id path integer true "Post id"
// @Produce     json
// @Success     200 {object} integer        "Likes were successfully deleted"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/likes [delete]
func (h *Handler) DeleteLike(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	if err = h.postsUseCase.UnlikePost(user.ID, postID); err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}
