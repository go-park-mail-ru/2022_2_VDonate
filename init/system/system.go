package system

import (
	authHandlers "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/auth/handlers"
	postsHandlers "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/posts/handlers"
	postsRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/posts/repository"
	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/session/repository"
	userHandlers "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/handlers"
	userRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/app/users/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/middleware"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router

	Logger *logger.Logger
	Config *config.Config

	Storage     *storage.Storage
	UserRepo    *userRepository.Repo
	SessionRepo *sessionRepository.Repo
	PostsRepo   *postsRepository.Repo

	AuthHTTPHandler  *authHandlers.HTTPHandler
	UserHTTPHandler  *userHandlers.HTTPHandler
	PostsHTTPHandler *postsHandlers.HTTPHandler
}

func (s *Server) init() {
	s.Logger.Logrus.Info("server started")
	log.SetOutput(s.Logger.Logrus.Writer())
	s.makeStorage()
	s.makeHandlers()
	s.makeRouter()
	s.makeCORS()
}

func (s *Server) Start() error {
	s.init()
	if err := http.ListenAndServe(":"+s.Config.Server.Port, s.Router); err != nil {
		return err
	}
	return nil
}

func (s *Server) makeStorage() {
	s.SessionRepo = sessionRepository.New()
	s.UserRepo = userRepository.New(s.Storage)
	s.PostsRepo = postsRepository.New(s.Storage)
}

func (s *Server) makeHandlers() {
	s.AuthHTTPHandler = authHandlers.NewHTTPHandler(s.UserRepo, s.SessionRepo)
	s.UserHTTPHandler = userHandlers.NewHTTPHandler(s.UserRepo, s.SessionRepo)
	s.PostsHTTPHandler = postsHandlers.NewHTPPHandler(s.PostsRepo)
}

func (s *Server) makeRouter() {
	authPostRouter := s.Router.Methods("POST", "OPTIONS").Subrouter()
	authPostRouter.HandleFunc("/api/v1/login", s.AuthHTTPHandler.Login)
	authPostRouter.HandleFunc("/api/v1/users", s.AuthHTTPHandler.SignUp)

	authGetRouter := s.Router.Methods("GET", "OPTIONS").Subrouter()
	authGetRouter.HandleFunc("/api/v1/auth", s.AuthHTTPHandler.Auth)

	authDeleteRouter := s.Router.Methods("DELETE", "OPTIONS").Subrouter()
	authDeleteRouter.HandleFunc("/api/v1/logout", s.AuthHTTPHandler.Logout)
	authDeleteRouter.Use(middleware.NewAuth(s.UserRepo, s.SessionRepo).LoginRequired)

	usersGetRouter := s.Router.Methods("GET", "OPTIONS").Subrouter()
	usersGetRouter.HandleFunc("/api/v1/users/{id}", s.UserHTTPHandler.GetUser)
	usersGetRouter.Use(middleware.NewAuth(s.UserRepo, s.SessionRepo).LoginRequired)

	postsGetRouter := s.Router.Methods("GET", "OPTIONS").Subrouter()
	postsGetRouter.HandleFunc("api/v1/users/{id}/posts", s.PostsHTTPHandler.Posts)
	postsGetRouter.Use(middleware.NewAuth(s.UserRepo, s.SessionRepo).LoginRequired)

}

func (s *Server) makeCORS() {
	c := middleware.NewCORS(s.Config.Debug.CORS)

	// due to strange logic of rs/cors
	if s.Config.Debug.CORS {
		c.Log = s.Logger.Logrus
	}

	s.Router.Use(c.Handler)
}

func New(s *storage.Storage, l *logger.Logger, c *config.Config) *Server {
	return &Server{
		Router:  mux.NewRouter(),
		Logger:  l,
		Config:  c,
		Storage: s,
	}
}
