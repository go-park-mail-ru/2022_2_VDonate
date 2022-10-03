package authErrors

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrUserAlreadyExist       = errors.New("username is already exist")
	ErrEmailAlreadyExist      = errors.New("email is already exist")
	ErrCreateUser             = errors.New("failed to create user")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrJSONMarshal            = errors.New("failed to marshal json")
	ErrJSONUnmarshal          = errors.New("failed to unmarshal json")
	ErrCopy                   = errors.New("failed to copy")
	ErrResponse               = errors.New("failed to response")
	ErrNoSession              = errors.New("no existing session")
)

func Wrap(w http.ResponseWriter, errHTTP, errLog error) {
	log.Printf("error: %s", errLog)
	switch errHTTP {
	case ErrNoSession:
		http.Error(w, errHTTP.Error(), http.StatusUnauthorized)
	case ErrUserNotFound:
		http.Error(w, errHTTP.Error(), http.StatusNotFound)
	case ErrInvalidLoginOrPassword:
		http.Error(w, errHTTP.Error(), http.StatusBadRequest)
	case ErrUserAlreadyExist, ErrEmailAlreadyExist:
		http.Error(w, errHTTP.Error(), http.StatusConflict)
	case ErrJSONMarshal, ErrResponse, ErrJSONUnmarshal, ErrCreateUser, ErrCopy:
		http.Error(w, errHTTP.Error(), http.StatusInternalServerError)
	}
}
