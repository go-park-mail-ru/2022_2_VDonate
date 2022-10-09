package app

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http/middlewares"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery/http"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery/http"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
)

type Server struct {
	Echo *echo.Echo

	Config *config.Config

	UserService  users.UseCase
	PostsService posts.UseCase
	AuthService  auth.UseCase

	authHandler  *httpAuth.Handler
	userHandler  *httpUsers.Handler
	postsHandler *httpPosts.Handler

	authMiddleware *authMiddlewares.Middlewares
}

func (s *Server) init() {
	s.makeLogger(logger.NewLogrus().Logrus.Writer())
	s.makeUseCase(s.Config.DB.URL)
	s.makeMiddlewares()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
}

func (s *Server) Start() error {
	s.init()
	return s.Echo.Start(s.Config.Server.Host + ":" + s.Config.Server.Port)
}

func (s *Server) makeUseCase(URL string) {
	//-----------------------sessions-----------------------//
	sessionRepo, err := sessionsRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	s.AuthService = auth.New(sessionRepo)

	//-------------------------user-------------------------//
	userRepo, err := userRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	s.UserService = users.New(userRepo)

	//-------------------------post-------------------------//
	postsRepo, err := postsRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	s.PostsService = posts.New(postsRepo)
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.AuthService, s.UserService)
	s.postsHandler = httpPosts.NewHandler(s.PostsService)
	s.userHandler = httpUsers.NewHandler(s.UserService, s.AuthService)
}

func (s *Server) makeLogger(l *io.PipeWriter) {
	s.Echo.Logger.SetOutput(l)
	s.Echo.Logger.Info("server started")
}

func (s *Server) makeRouter() {
	v1 := s.Echo.Group("/api/v1")
	if s.Config.Debug.Request {
		v1.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	}

	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
	v1.POST("/users", s.authHandler.SignUp)

	user := v1.Group("/users/:id", s.authMiddleware.LoginRequired)
	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)

	post := v1.Group("/posts", s.authMiddleware.LoginRequired)
	post.GET("/users/:id", s.postsHandler.GetPosts)
	post.POST("/users/:id", s.postsHandler.CreatePosts, s.authMiddleware.SameSession)
}

func (s *Server) makeCORS() {
	s.Echo.Use(NewCORS())
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.AuthService)
}

func New(echo *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   echo,
		Config: c,
	}
}
