package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	//Incluir endpoint
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}