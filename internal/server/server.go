package server

import (
	auth_http "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/handlers"
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
	s.Router.HandleFunc("/api/v1/login", s.AuthHTTPHandler.Login).Methods("POST")
	s.Router.HandleFunc("/api/v1/auth", s.AuthHTTPHandler.Auth).Methods("GET")
	s.Router.HandleFunc("/api/v1/users", s.AuthHTTPHandler.SignUp).Methods("POST")
	s.Router.HandleFunc("/api/v1/logout", s.AuthHTTPHandler.Logout).Methods("DELETE")
	s.Router.HandleFunc("/api/v1/users/{id}", s.UserHTTPHandler.GetUser).Methods("GET")
}

func New(s *storage.Storage, l *logger.Logger, c *config.Config) *Server {
	return &Server{
		Router:  mux.NewRouter(),
		Logger:  l,
		Config:  c,
		Storage: s,
	}
}
