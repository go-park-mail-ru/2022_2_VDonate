package usersErrors

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrConvertID     = errors.New("unable to convert id")
	ErrUserNotFound  = errors.New("user not found")
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
)

func Wrap(w http.ResponseWriter, errHTTP, errLog error) {
	log.Printf("errors: %s", errLog)
	switch errHTTP {
	case ErrUserNotFound:
		http.Error(w, errHTTP.Error(), http.StatusNotFound)
	case ErrConvertID, ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse:
		http.Error(w, errHTTP.Error(), http.StatusInternalServerError)
	}
}
