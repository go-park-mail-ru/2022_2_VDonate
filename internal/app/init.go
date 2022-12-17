package app

import (
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc/credentials/insecure"

	imagesMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/delivery/grpc"

	donatesMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/delivery/grpc"

	subscribersMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/delivery/grpc"

	subscriptionsMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/delivery/grpc"

	authMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/grpc"
	postsMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery/grpc"

	usersMicroservice "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery/grpc"

	authProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	donatesProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	imagesProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/images/protobuf"
	postProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	subscribersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"
	subscriptionsProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	usersProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"

	"google.golang.org/grpc"

	httpSubscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/delivery/http"

	httpDonates "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/delivery/http"

	httpImages "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/delivery/http"

	httpPosts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/delivery/http"

	httpSubscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/delivery/http"

	httpUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery/http"

	authMiddlewares "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http/middlewares"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	donates "github.com/go-park-mail-ru/2022_2_VDonate/internal/donates/usecase"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	posts "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts/usecase"
	subscribers "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers/usecase"
	subscriptions "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscriptions/usecase"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	Echo *echo.Echo

	Config *config.Config

	UserMicroservice         domain.UsersMicroservice
	PostsMicroservice        domain.PostsMicroservice
	AuthMicroservice         domain.AuthMicroservice
	SubscriptionMicroservice domain.SubscriptionMicroservice
	SubscribersMicroservice  domain.SubscribersMicroservice
	DonatesMicroservice      domain.DonatesMicroservice
	ImagesMicroservice       domain.ImageMicroservice

	UserUseCase         domain.UsersUseCase
	PostsUseCase        domain.PostsUseCase
	AuthUseCase         domain.AuthUseCase
	SubscriptionUseCase domain.SubscriptionsUseCase
	SubscribersUseCase  domain.SubscribersUseCase
	DonatesUseCase      domain.DonatesUseCase
	ImagesUseCase       domain.ImageUseCase

	authHandler          *httpAuth.Handler
	userHandler          *httpUsers.Handler
	postsHandler         *httpPosts.Handler
	subscriptionsHandler *httpSubscriptions.Handler
	subscribersHandler   *httpSubscribers.Handler
	donatesHandler       *httpDonates.Handler
	imagesHandler        *httpImages.Handler

	authMiddleware *authMiddlewares.Middlewares
}

func (s *Server) init() error {
	s.makeEchoLogger()
	if err := s.makeGRPCClients(); err != nil {
		return err
	}
	if err := s.makeUseCase(); err != nil {
		return err
	}
	s.makeMiddlewares()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
	if s.Config.CSRF.Enabled {
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

func makeAddress(host, port string) string {
	return host + ":" + port
}

func (s *Server) makeGRPCClients() error {
	//----------------------connection----------------------//
	userConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Users.Host, s.Config.Services.Users.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	postsConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Posts.Host, s.Config.Services.Posts.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	authConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Auth.Host, s.Config.Services.Auth.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	subscriptionConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Subscriptions.Host, s.Config.Services.Subscriptions.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	subscribersConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Subscribers.Host, s.Config.Services.Subscribers.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	donatesConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Donates.Host, s.Config.Services.Donates.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	imagesConnection, err := grpc.Dial(
		makeAddress(s.Config.Services.Images.Host, s.Config.Services.Images.Port),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	//---------------------microservice---------------------//
	s.UserMicroservice = usersMicroservice.New(usersProto.NewUsersClient(userConnection))

	s.PostsMicroservice = postsMicroservice.New(postProto.NewPostsClient(postsConnection))

	s.AuthMicroservice = authMicroservice.New(authProto.NewAuthClient(authConnection))

	s.SubscriptionMicroservice = subscriptionsMicroservice.New(subscriptionsProto.NewSubscriptionsClient(subscriptionConnection))

	s.SubscribersMicroservice = subscribersMicroservice.New(subscribersProto.NewSubscribersClient(subscribersConnection))

	s.DonatesMicroservice = donatesMicroservice.New(donatesProto.NewDonatesClient(donatesConnection))

	s.ImagesMicroservice = imagesMicroservice.New(imagesProto.NewImagesClient(imagesConnection))

	return nil
}

func (s *Server) makeUseCase() error {
	//------------------------images------------------------//
	s.ImagesUseCase = images.New(s.ImagesMicroservice)

	//-----------------------sessions-----------------------//
	s.AuthUseCase = auth.New(s.AuthMicroservice, s.UserMicroservice)

	//-------------------------user-------------------------//
	s.UserUseCase = users.New(s.UserMicroservice, s.ImagesUseCase)

	//-------------------------post-------------------------//
	s.PostsUseCase = posts.New(s.PostsMicroservice, s.UserMicroservice, s.ImagesUseCase, s.SubscriptionMicroservice)

	//----------------------subscriber----------------------//
	s.SubscribersUseCase = subscribers.New(s.SubscribersMicroservice, s.UserMicroservice)

	//---------------------subscription---------------------//
	s.SubscriptionUseCase = subscriptions.New(s.SubscriptionMicroservice, s.UserMicroservice, s.ImagesUseCase)

	//-----------------------donates------------------------//
	s.DonatesUseCase = donates.New(s.DonatesMicroservice, s.UserMicroservice)

	return nil
}

func (s *Server) makeHandlers() {
	s.authHandler = httpAuth.NewHandler(s.AuthUseCase, s.UserUseCase)

	s.imagesHandler = httpImages.NewHandler(s.ImagesUseCase)
	s.donatesHandler = httpDonates.NewHandler(s.DonatesUseCase, s.UserUseCase)
	s.postsHandler = httpPosts.NewHandler(s.PostsUseCase, s.UserUseCase, s.ImagesUseCase)
	s.userHandler = httpUsers.NewHandler(s.UserUseCase, s.AuthUseCase, s.ImagesUseCase, s.SubscriptionUseCase, s.SubscribersUseCase)
	s.subscriptionsHandler = httpSubscriptions.NewHandler(s.SubscriptionUseCase, s.UserUseCase, s.ImagesUseCase)
	s.subscribersHandler = httpSubscribers.NewHandler(s.SubscribersUseCase, s.UserUseCase)
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

	v1.POST("/login", s.authHandler.Login)
	v1.GET("/auth", s.authHandler.Auth)
	v1.DELETE("/logout", s.authHandler.Logout, s.authMiddleware.LoginRequired)
	v1.POST("/users", s.authHandler.SignUp)

	user := v1.Group("/users/:id")
	user.Use(s.authMiddleware.LoginRequired)

	user.GET("", s.userHandler.GetUser)
	user.PUT("", s.userHandler.PutUser)

	search := v1.Group("/search")
	search.Use(s.authMiddleware.LoginRequired)

	search.GET("", s.userHandler.GetAuthors)

	post := v1.Group("/posts")
	post.Use(s.authMiddleware.LoginRequired)

	post.GET("/:id/comments", s.postsHandler.GetComments)
	post.POST("/:id/comments", s.postsHandler.PutComment)
	post.PUT("/:id/comments", s.postsHandler.PutComment)
	post.DELETE("/:id/comments", s.postsHandler.DeleteComment)

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

	image := v1.Group("/image")
	image.POST("", s.imagesHandler.CreateOrUpdateImage)
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
		TokenLength:    s.Config.CSRF.TokenLength,
		TokenLookup:    "header:" + echo.HeaderXCSRFToken,
		ContextKey:     s.Config.CSRF.ContextKey,
		CookieName:     s.Config.CSRF.ContextName,
		CookieMaxAge:   s.Config.CSRF.MaxAge,
		CookiePath:     "/",
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure:   true,
	}))
}

func (s *Server) makeMiddlewares() {
	s.authMiddleware = authMiddlewares.New(s.AuthUseCase, s.UserUseCase, s.PostsUseCase)
}

func New(echo *echo.Echo, c *config.Config) *Server {
	return &Server{
		Echo:   echo,
		Config: c,
	}
}
