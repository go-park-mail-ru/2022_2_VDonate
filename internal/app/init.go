package app

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	authMiddlewares "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/middlewares"
	sessionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/repository"
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	httpdonates "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/delivery"
	donatesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/repository"
	donates "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/usecase"
	imagesMiddleware "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/middlewares"
	imagesRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/repository"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	httpPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery"
	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/repository"
	posts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	httpsubscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/delivery"
	subscribersRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/repository"
	subscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/usecase"
	httpSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/delivery"
	subscriptionsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/repository"
	subscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/usecase"
	httpUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery"
	userRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	Echo *echo.Echo

	Config *config.Config

	UserService         domain.UsersUseCase
	PostsService        domain.PostsUseCase
	AuthService         domain.AuthUseCase
	SubscriptionService domain.SubscriptionsUseCase
	SubscribersService  domain.SubscribersUseCase
	DonatesService      domain.DonatesUseCase
	ImagesService       domain.ImageUseCase

	authHandler          *httpAuth.Handler
	userHandler          *httpUsers.Handler
	postsHandler         *httpPosts.Handler
	subscriptionsHandler *httpSubscriptions.Handler
	subscribersHandler   *httpsubscribers.Handler
	donatesHandler       *httpdonates.Handler

	authMiddleware *authMiddlewares.Middlewares
}

func (s *Server) init() error {
	s.makeEchoLogger()
	if err := s.makeUseCase(s.Config.DB.URL); err != nil {
		return err
	}
	s.makeMiddlewares()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
	if s.Config.CSRF.Status {
		s.makeCSRF()
	}

	return nil
}

func (s *Server) Start() error {
	if err := s.init(); err != nil {
		return err
	}

	return s.Echo.Start(s.Config.Server.Host + ":" + s.Config.Server.Port)
}

func (s *Server) StartTLS() error {
	if err := s.init(); err != nil {
		return err
	}
	return s.Echo.StartTLS(
		s.Config.Server.Host+":"+s.Config.Server.Port,
		s.Config.Server.CertPath,
		s.Config.Server.KeyPath,
	)
}

func (s *Server) makeUseCase(url string) error {
	//-------------------------repo-------------------------//
	sessionRepo, err := sessionsRepository.NewPostgres(url)
	if err != nil {
		return err
	}
	userRepo, err := userRepository.NewPostgres(url)
	if err != nil {
		return err
	}
	postsRepo, err := postsRepository.NewPostgres(url)
	if err != nil {
		return err
	}
	subscriptionsRepo, err := subscriptionsRepository.NewPostgres(url)
	if err != nil {
		return err
	}
	subscribersRepo, err := subscribersRepository.NewPostgres(url)
	if err != nil {
		return err
	}
	imagesRepo, err := imagesRepository.New(
		s.Config.S3.Endpoint,
		s.Config.S3.AccessKeyID,
		s.Config.S3.SecretAccessKey,
		s.Config.S3.UseSSL,
		s.Config.S3.Buckets,
	)
	if err != nil {
		return err
	}
	donatesRepo, err := donatesRepository.NewPostgres(url)
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

	//-----------------------donates------------------------//
	s.DonatesService = donates.New(donatesRepo, userRepo)

	//------------------------images------------------------//
	s.ImagesService = images.New(imagesRepo)

	return nil
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.AuthService, s.UserService)

	s.donatesHandler = httpdonates.NewHandler(s.DonatesService, s.UserService)
	s.postsHandler = httpPosts.NewHandler(s.PostsService, s.UserService, s.ImagesService)
	s.userHandler = httpUsers.NewHandler(s.UserService, s.AuthService, s.ImagesService, s.SubscriptionService, s.SubscribersService)
	s.subscriptionsHandler = httpSubscriptions.NewHandler(s.SubscriptionService, s.UserService, s.ImagesService)
	s.subscribersHandler = httpsubscribers.NewHandler(s.SubscribersService, s.UserService)
}

func (s *Server) makeEchoLogger() {
	s.Echo.Logger = logger.GetInstance()
	s.Echo.Logger.SetLevel(logger.ToLevel(s.Config.Logger.Level))
	s.Echo.Logger.Info("server started")
}

func (s *Server) makeRouter() {
	s.Echo.Pre(middleware.RemoveTrailingSlash())

	s.Echo.GET("docs/*", echoSwagger.WrapHandler)

	s.Echo.Use(logger.Middleware())
	s.Echo.Use(middleware.Secure())
	v1 := s.Echo.Group("/api/v1")

	v1.Use(imagesMiddleware.BucketManager)

	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
	v1.POST("/users", s.authHandler.SignUp)

	user := v1.Group("/users/:id")
	user.Use(s.authMiddleware.LoginRequired)

	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)

	post := v1.Group("/posts")
	post.Use(s.authMiddleware.LoginRequired)

	post.GET("/:id/likes", s.postsHandler.GetLikes)
	post.POST("/:id/likes", s.postsHandler.CreateLike)
	post.DELETE("/:id/likes", s.postsHandler.DeleteLike)

	post.GET("", s.postsHandler.GetPosts)
	post.POST("", s.postsHandler.CreatePost)
	post.GET("/:id", s.postsHandler.GetPost)
	post.DELETE("/:id", s.postsHandler.DeletePost, s.authMiddleware.PostSameSessionByID)
	post.PUT("/:id", s.postsHandler.PutPost, s.authMiddleware.PostSameSessionByID)

	subscription := v1.Group("/subscriptions")
	subscription.Use(s.authMiddleware.LoginRequired)

	subscription.GET("", s.subscriptionsHandler.GetSubscriptions)

	authorSubscription := subscription.Group("/author")

	authorSubscription.GET("", s.subscriptionsHandler.GetAuthorSubscriptions)
	authorSubscription.POST("", s.subscriptionsHandler.CreateAuthorSubscription)
	authorSubscription.PUT("/:id", s.subscriptionsHandler.UpdateAuthorSubscription)
	authorSubscription.GET("/:id", s.subscriptionsHandler.GetAuthorSubscription)
	authorSubscription.DELETE("/:id", s.subscriptionsHandler.DeleteAuthorSubscription)

	subscriber := v1.Group("/subscribers")
	subscriber.Use(s.authMiddleware.LoginRequired)

	subscriber.GET("/:author_id", s.subscribersHandler.GetSubscribers)
	subscriber.POST("", s.subscribersHandler.CreateSubscriber)
	subscriber.DELETE("", s.subscribersHandler.DeleteSubscriber)

	donate := v1.Group("/donates")
	donate.Use(s.authMiddleware.LoginRequired)

	donate.GET("/:id", s.donatesHandler.GetDonate)
	donate.GET("", s.donatesHandler.GetDonates)
	donate.POST("", s.donatesHandler.CreateDonate)
}

func (s *Server) makeCORS() {
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     s.Config.CORS.AllowOrigins,
		AllowMethods:     s.Config.CORS.AllowMethods,
		AllowCredentials: s.Config.CORS.AllowCredentials,
		AllowHeaders:     s.Config.CORS.AllowHeaders,
	}))
}

func (s *Server) makeCSRF() {
	s.Echo.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:  s.Config.CSRF.TokenLength,
		TokenLookup:  "header:" + echo.HeaderXCSRFToken,
		ContextKey:   s.Config.CSRF.ContextKey,
		CookieName:   s.Config.CSRF.ContextName,
		CookieMaxAge: s.Config.CSRF.MaxAge,
		CookiePath:   "/",
	}))
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
