package server

import (
	auth_http "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/handlers"
	auth_middleware "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/middleware"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/config"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/storages"
	cookie_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/cookie"
	user_http "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/handlers"
	user_repo "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var AllowOrigins = []string{
	"localhost:8080",
	"zenehu:8080",
}

type Server struct {
	Router *mux.Router

	Logger *logger.Logger
	Config *config.Config

	Storage    *storage.Storage
	UserRepo   *user_repo.Repo
	CookieRepo *cookie_repo.Repo

	AuthHTTPHandler *auth_http.AuthHTTPHandler
	UserHTTPHandler *user_http.UserHTTPHandler
}

func (s *Server) Start() error {
	s.Logger.Logrus.Info("server started")
	log.SetOutput(s.Logger.Logrus.Writer())
	s.makeStorage()
	s.makeHandlers()
	s.makeRouter()

	if err := http.ListenAndServe(s.Config.Port, s.Router); err != nil {
		return err
	}
	return nil
}

func (s *Server) makeStorage() {
	s.CookieRepo = cookie_repo.New()
	s.UserRepo = user_repo.New(s.Storage)
}

func (s *Server) makeHandlers() {
	s.AuthHTTPHandler = auth_http.New(s.UserRepo, s.CookieRepo)
	s.UserHTTPHandler = user_http.New(s.UserRepo, s.CookieRepo)
}

func (s *Server) makeRouter() {
	authPostRouter := s.Router.Methods("POST").Subrouter()
	authPostRouter.HandleFunc("/api/v1/login", s.AuthHTTPHandler.Login)
	authPostRouter.HandleFunc("/api/v1/users", s.AuthHTTPHandler.SignUp)

	authGetRouter := s.Router.Methods("GET").Subrouter()
	authGetRouter.HandleFunc("api/v1/auth", s.AuthHTTPHandler.Auth)

	authDeleteRouter := s.Router.Methods("DELETE").Subrouter()
	authDeleteRouter.HandleFunc("/api/v1/logout", s.AuthHTTPHandler.Logout).Methods("DELETE")
	authDeleteRouter.Use(auth_middleware.New(s.UserRepo, s.CookieRepo).LoginRequired)

	usersGetRouter := s.Router.Methods("GET").Subrouter()
	usersGetRouter.HandleFunc("/api/v1/users/{id}", s.UserHTTPHandler.GetUser)
	usersGetRouter.Use(auth_middleware.New(s.UserRepo, s.CookieRepo).LoginRequired)

	s.Router.Use(auth_middleware.CORS(AllowOrigins))
}

func New(s *storage.Storage, l *logger.Logger, c *config.Config) *Server {
	return &Server{
		Router:  mux.NewRouter(),
		Logger:  l,
		Config:  c,
		Storage: s,
	}
}
