package handlers

import (
	_ "avito/docs"
	"encoding/json"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type UserCase interface {
	UserService
	SegmentService
}

func NewHandler(router *mux.Router, userCase UserCase) {
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	newUserHandler(router, userCase)
	newSegmentHandler(router, userCase)
}

func makeHTTPHandleFunc(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err = writeJson(w, http.StatusBadRequest, err.Error())
			if err != nil {
				return
			}
		}

	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}

type ApiAnswer struct {
	Ans string
}
