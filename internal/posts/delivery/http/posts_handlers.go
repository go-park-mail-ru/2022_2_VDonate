package httpPosts

import (
	"net/http"
	"strconv"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
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
// @Param       filter query    string         true "filter to use to get posts. Filters: subscriptions, user_id(as digit)"
// @Success     200    {object} []models.Post  "Posts were successfully received"
// @Failure     400    {object} echo.HTTPError "Bad request"
// @Failure     401    {object} echo.HTTPError "No session provided"
// @Failure     404    {object} echo.HTTPError "User not found"
// @Failure     500    {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts [get]
func (h Handler) GetPosts(c echo.Context) error {
	var allPosts []models.Post
	var user models.User
	var authorID uint64

	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	if user, err = h.usersUseCase.GetBySessionID(cookie.Value); err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	filter := c.QueryParam("filter")
	switch filter {
	case "subscriptions":
	default:
		if authorID, err = strconv.ParseUint(filter, 10, 64); err != nil {
			return errorHandling.WrapEcho(domain.ErrBadRequest, err)
		}
	}

	if allPosts, err = h.postsUseCase.GetPostsByFilter(user.ID, authorID); err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
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
func (h Handler) GetPost(c echo.Context) error {
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

	post, err := h.postsUseCase.GetPostByID(postID, user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

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
func (h Handler) DeletePost(c echo.Context) error {
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
func (h Handler) PutPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	var prevPost models.Post

	if err = c.Bind(&prevPost); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if prevPost, err = h.postsUseCase.Update(prevPost, postID); err != nil {
		return errorHandling.WrapEcho(domain.ErrUpdate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImagePosts{
		PostID:  postID,
		Content: prevPost.Content,
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
func (h Handler) CreatePost(c echo.Context) error {
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

	if post, err = h.postsUseCase.Create(post, user.ID); err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, models.ResponseImagePosts{
		PostID:      post.ID,
		Content:     post.Content,
		DateCreated: post.DateCreated,
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
func (h Handler) GetLikes(c echo.Context) error {
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
// @Param       id path integer true "Post id"
// @Produce     json
// @Success     200 {object} integer        "Likes were successfully create"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/likes [post]
func (h Handler) CreateLike(c echo.Context) error {
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
// @Description Delete like on post
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
func (h Handler) DeleteLike(c echo.Context) error {
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

// GetComments godoc
// @Summary     Get comments
// @Description Get all comments by post id
// @ID          get_posts_comments
// @Tags        posts
// @Param       id path integer true "Post id"
// @Produce     json
// @Success     200 {object} []models.Comment "Comments were successfully recieved"
// @Failure     400 {object} echo.HTTPError   "Bad request"
// @Failure     401 {object} echo.HTTPError   "No session provided"
// @Failure     404 {object} echo.HTTPError   "Post not found"
// @Failure     500 {object} echo.HTTPError   "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/comments [get]
func (h Handler) GetComments(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	allComments, err := h.postsUseCase.GetCommentsByPostID(postID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}
	if len(allComments) == 0 {
		return c.JSON(http.StatusOK, []models.Comment{})
	}
	return c.JSON(http.StatusOK, allComments)
}

// PostComment godoc
// @Summary     Post comment
// @Description Post comment on post
// @ID          post_comment
// @Tags        posts
// @Param       id      path integer        true "Post id"
// @Param       comment body models.Comment true "Comment"
// @Produce     json
// @Success     200 {object} integer        "Comment was successfully put"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/{id}/comments [post]
func (h Handler) PostComment(c echo.Context) error {
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

	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	comment.UserID = user.ID
	comment.PostID = postID

	if comment, err = h.postsUseCase.CreateComment(comment); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, comment)
}

// PutComment godoc
// @Summary     Put comment
// @Description Put comment on post
// @ID          put_comment
// @Tags        posts
// @Param       id      path integer        true "Comment id"
// @Param       comment body models.Comment true "Comment"
// @Produce     json
// @Success     200 {object} integer        "Comment was successfully put"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/comments/:id [put]
func (h Handler) PutComment(c echo.Context) error {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}
	comment.ID = commentID

	if comment, err = h.postsUseCase.UpdateComment(comment.ID, comment.Content); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary     Delete comment
// @Description Delete comment on post
// @ID          delete_comment
// @Tags        posts
// @Param       id path integer true "Comment id"
// @Produce     json
// @Success     200 {object} integer        "Comment was successfully deleted"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /posts/comments/{id} [delete]
func (h Handler) DeleteComment(c echo.Context) error {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	if err = h.postsUseCase.DeleteComment(commentID); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}
