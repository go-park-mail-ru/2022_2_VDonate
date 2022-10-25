package app

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/middlewares"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	httpPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	httpsubscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/delivery"
	subscribersRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/repository"
	subscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/usecase"
	httpSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/delivery"
	subscriptionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/repository"
	subscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/usecase"
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

	UserService         domain.UsersUseCase
	PostsService        domain.PostsUseCase
	AuthService         domain.AuthUseCase
	SubscriptionService domain.SubscriptionsUseCase
	SubscribersService  domain.SubscribersUseCase

	authHandler          *httpAuth.Handler
	userHandler          *httpUsers.Handler
	postsHandler         *httpPosts.Handler
	subscriptionsHandler *httpSubscriptions.Handler
	subscribersHandler   *httpsubscribers.Handler

	authMiddleware *authMiddlewares.Middlewares
}

func (s *Server) init() {
	s.makeEchoLogger()
	s.makeUseCase(s.Config.DB.URL)
	s.makeMiddlewares()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
	s.makeCSRF()
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
	subscriptionsRepo, err := subscriptionsRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}
	subscribersRepo, err := subscribersRepository.NewPostgres(URL)
	if err != nil {
		s.Echo.Logger.Error(err)
	}

	//-----------------------sessions-----------------------//
	s.AuthService = auth.New(sessionRepo, userRepo)

	//-------------------------user-------------------------//
	s.UserService = users.New(userRepo)

	//-------------------------post-------------------------//
	s.PostsService = posts.New(postsRepo)

	//----------------------subscriber----------------------//
	s.SubscribersService = subscribers.New(subscribersRepo, userRepo)

	//---------------------subscription---------------------//
	s.SubscriptionService = subscriptions.New(subscriptionsRepo)
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.AuthService, s.UserService)
	s.postsHandler = httpPosts.NewHandler(s.PostsService, s.UserService)
	s.userHandler = httpUsers.NewHandler(s.UserService, s.AuthService)
	s.subscriptionsHandler = httpSubscriptions.New(s.SubscriptionService, s.UserService)
	s.subscribersHandler = httpsubscribers.New(s.SubscribersService, s.UserService)
}

func (s *Server) makeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.Info("server started")
}

func (s *Server) makeRouter() {
	v1 := s.Echo.Group("/api/v1")
	if s.Config.Debug.Request {
		v1.Use(logger.Middleware())
	}
	v1.Use(middleware.Secure())

	v1.GET("/get_csrf", s.authHandler.GetCSRF)

	v1.POST("/login", s.authHandler.Login, s.authMiddleware.CSRFRequired)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired, s.authMiddleware.CSRFRequired)
	v1.POST("/users", s.authHandler.SignUp, s.authMiddleware.CSRFRequired)

	user := v1.Group("/users/:id")
	user.Use(s.authMiddleware.LoginRequired)

	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)

	post := v1.Group("/posts")
	post.Use(s.authMiddleware.LoginRequired)

	post.GET("", s.postsHandler.GetPosts)
	post.POST("", s.postsHandler.CreatePosts)
	post.GET("/:id", s.postsHandler.GetPost)
	post.DELETE("/:id", s.postsHandler.DeletePost, s.authMiddleware.PostSameSessionByID)
	post.PUT("/:id", s.postsHandler.PutPost, s.authMiddleware.PostSameSessionByID)

	subscription := v1.Group("/subscriptions")
	subscription.Use(s.authMiddleware.LoginRequired)

	subscription.GET("/:id", s.subscriptionsHandler.GetSubscription)
	subscription.GET("", s.subscriptionsHandler.GetSubscriptions)
	subscription.POST("", s.subscriptionsHandler.CreateSubscription)
	subscription.PUT("", s.subscriptionsHandler.UpdateSubscription)
	subscription.DELETE("/:id", s.subscriptionsHandler.DeleteSubscription)

	subscriber := v1.Group("/subscribers")
	subscription.Use(s.authMiddleware.LoginRequired)

	subscriber.GET("/:author_id", s.subscribersHandler.GetSubscribers)
	subscriber.POST("", s.subscribersHandler.CreateSubscriber)
	subscriber.DELETE("", s.subscribersHandler.DeleteSubscriber)
}

func (s *Server) makeCORS() {
	s.Echo.Use(NewCORS())
}

func (s *Server) makeCSRF() {
	s.Echo.Use(NewCSRF())
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.AuthService, s.UserService, s.PostsService)
}

func New(echo *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   echo,
		Config: c,
	}
}
