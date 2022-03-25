package server

import (
	"context"
	"gateway/cmd/gateway/proxy"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	server *http.Server
	router *mux.Router
}

func New() *server {
	router := mux.NewRouter()
	httpServer := &http.Server{
		Handler: router,
		Addr:    "localhost:8081",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	productProxy := proxy.New("http://localhost:8080")
	accountProxy := proxy.New("http://localhost:8081")
	expenseProxy := proxy.New("http://localhost:8082")

	s.router.HandleFunc("/accounts/{rest:.*}", accountProxy.HandleRequest)
	s.router.HandleFunc("/products/{rest:.*}", productProxy.HandleRequest)
	s.router.HandleFunc("/expenses/{rest:.*}", expenseProxy.HandleRequest)

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
}
