package app

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"
	authMiddlewares "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http/middlewares"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	httpPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery/http"
	postsPostgres "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository/postgres"
	postsAPI "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	sessionPostgres "github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository/postgres"
	httpUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery/http"
	userPostgres "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository/postgres"
	userAPI "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
)

type Server struct {
	Echo *echo.Echo

	Config *config.Config

	UserAPI     userAPI.UseCase
	PostsAPI    postsAPI.UseCase
	SessionRepo sessionRepository.API

	authHandler  *httpAuth.Handler
	userHandler  *httpUsers.Handler
	postsHandler *httpPosts.Handler

	authMiddleware *authMiddlewares.Middlewares
}

func (s *Server) init() {
	s.makeLogger(logger.NewLogrus().Logrus.Writer())
	s.Echo.Logger.Info("server started")
	s.makeStorages(s.Config.DB.URL)
	s.makeMiddlewares()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
}

func (s *Server) Start() error {
	s.init()
	return s.Echo.Start(s.Config.Server.Host + ":" + s.Config.Server.Port)
}

func (s *Server) makeStorages(URL string) {
	var err error

	s.SessionRepo, err = sessionPostgres.New(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	userRepo, err := userPostgres.New(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	s.UserAPI = userAPI.New(userRepo, s.SessionRepo)
	postsRepo, err := postsPostgres.New(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	s.PostsAPI = postsAPI.New(postsRepo)
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.UserAPI, s.SessionRepo)
	s.postsHandler = httpPosts.NewHandler(s.PostsAPI)
	s.userHandler = httpUsers.NewHandler(s.UserAPI, s.SessionRepo)
}

func (s *Server) makeLogger(l *io.PipeWriter) {
	s.Echo.Logger.SetOutput(l)
}

func (s *Server) makeRouter() {
	v1 := s.Echo.Group("/api/v1", middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
	v1.POST("/users", s.authHandler.SignUp)

	v1.GET("/users/:id", s.userHandler.GetUser, s.authMiddleware.LoginRequired)
	v1.PUT("/update/users/:id", s.userHandler.PutUser, s.authMiddleware.LoginRequired)

	v1.GET("/users/:id/posts", s.postsHandler.GetPosts, s.authMiddleware.LoginRequired)
}

func (s *Server) makeCORS() {
	s.Echo.Use(NewCORS())
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.SessionRepo)
}

func New(echo *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   echo,
		Config: c,
	}
}
