package postsErrors

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
	ErrBadRequest    = errors.New("bar request")
)

func Wrap(w http.ResponseWriter, errHTTP, errLog error) {
	log.Printf("errors: %s", errLog)
	switch errHTTP {
	case ErrJSONMarshal, ErrResponse, ErrJSONUnmarshal, ErrBadRequest:
		http.Error(w, errHTTP.Error(), http.StatusInternalServerError)
	}
}
