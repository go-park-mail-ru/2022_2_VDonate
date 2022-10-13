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
	"io"
)

type Server struct {
	Echo *echo.Echo

	Config *config.Config

	UserService         users.UseCase
	PostsService        posts.UseCase
	AuthService         auth.UseCase
	SubscriptionService subscriptions.UseCase
	SubscribersService  subscribers.UseCase

	authHandler          *httpAuth.Handler
	userHandler          *httpUsers.Handler
	postsHandler         *httpPosts.Handler
	subscriptionsHandler *httpSubscriptions.Handler
	subscribersHandler   *httpsubscribers.Handler

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
	s.postsHandler = httpPosts.NewHandler(s.PostsService)
	s.userHandler = httpUsers.NewHandler(s.UserService, s.AuthService)
	s.subscriptionsHandler = httpSubscriptions.New(s.SubscriptionService, s.UserService)
	s.subscribersHandler = httpsubscribers.New(s.SubscribersService, s.UserService)
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

	user := v1.Group("/users/:id")
	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)
	user.Use(s.authMiddleware.LoginRequired)

	post := v1.Group("/posts")
	post.GET("/users/:id", s.postsHandler.GetPosts)
	post.POST("/users/:id", s.postsHandler.CreatePosts, s.authMiddleware.SameSession)
	post.Use(s.authMiddleware.LoginRequired)

	subscription := v1.Group("/subscriptions")
	subscription.GET("/:author_id", s.subscriptionsHandler.GetSubscriptions)
	subscription.POST("/", s.subscriptionsHandler.CreateSubscription)
	subscription.PUT("/", s.subscriptionsHandler.UpdateSubscription)
	subscription.DELETE("/:subscription_id", s.subscriptionsHandler.DeleteSubscription)

	subscriber := v1.Group("/subscribers")
	subscriber.GET("/:author_id", s.subscribersHandler.GetSubscribers)
	subscriber.POST("/", s.subscribersHandler.CreateSubscriber)
	subscriber.DELETE("/", s.subscribersHandler.DeleteSubscriber)
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
