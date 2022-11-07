package imagesMiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestBucketManager_Images(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("api/v1/posts")

	h := BucketManager(echo.NotFoundHandler)

	err := h(c)
	assert.ErrorIs(t, err, echo.ErrNotFound)

	assert.Equal(t, c.Get("bucket"), "image")
}

func TestBucketManager_Avatar(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.URL.Path = "api/v1/users"
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetPath("api/v1/users")

	h := BucketManager(echo.NotFoundHandler)

	err := h(c)
	assert.ErrorIs(t, err, echo.ErrNotFound)

	assert.Equal(t, c.Get("bucket"), "avatar")
}
