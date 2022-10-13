package app

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/middlewares"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	httpPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	httpUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	s.makeEchoLogger()
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
	//-------------------------repo-------------------------//
	sessionRepo, err := sessionsRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	userRepo, err := userRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	postsRepo, err := postsRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	//-----------------------sessions-----------------------//
	s.AuthService = auth.New(sessionRepo, userRepo)

	//-------------------------user-------------------------//
	s.UserService = users.New(userRepo)

	//-------------------------post-------------------------//
	s.PostsService = posts.New(postsRepo)
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.AuthService, s.UserService)
	s.postsHandler = httpPosts.NewHandler(s.PostsService)
	s.userHandler = httpUsers.NewHandler(s.UserService, s.AuthService)
}

func (s *Server) makeEchoLogger() {
	s.Echo.Logger = logger.GlobalLogger
	s.Echo.Logger.Info("server started")
}

func (s *Server) makeRouter() {
	v1 := s.Echo.Group("/api/v1")
	if s.Config.Debug.Request {
		v1.Use(logger.Middleware())
	}

	v1.Use(middleware.Secure())

	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
	v1.POST("/users", s.authHandler.SignUp)

	user := v1.Group("/users/:id")
	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)
	user.Use(s.authMiddleware.LoginRequired)

	post := v1.Group("/posts")
	post.GET("/", s.postsHandler.GetPosts)
	post.POST("/", s.postsHandler.CreatePosts, s.authMiddleware.SameSession)
	post.DELETE("/:post_id/", s.postsHandler.DeletePost, s.authMiddleware.SameSession)
	post.PUT("/:post_id/", s.postsHandler.PutPost, s.authMiddleware.SameSession)
	post.Use(s.authMiddleware.LoginRequired)
}

func (s *Server) makeCORS() {
	s.Echo.Use(NewCORS())
	middleware.Logger()
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.AuthService, s.UserService)
}

func New(echo *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   echo,
		Config: c,
	}
}
