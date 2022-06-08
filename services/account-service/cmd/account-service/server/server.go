package server

import (
	"account-service/cmd/account-service/server/database"
	"account-service/cmd/account-service/server/factory"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	server *http.Server
	router *mux.Router
}

func NewServer() *server {
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
	s.router.HandleFunc("/accounts", factory.MakeCreateAccountController().HandleRequest).Methods("POST")
	s.router.HandleFunc("/sessions", factory.MakeLoginController().HandleRequest).Methods("POST")
	s.router.HandleFunc("/sessions", factory.MakeValidateTokenController().HandleRequest).Methods("GET")

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
	database.StopMongo()
	database.StopRedis()
}
