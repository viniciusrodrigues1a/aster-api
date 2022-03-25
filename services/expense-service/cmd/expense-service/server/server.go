package server

import (
	"context"
	"expense-service/cmd/expense-service/server/database"
	"expense-service/cmd/expense-service/server/factory"
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
		Addr:    "localhost:8082",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	s.router.HandleFunc("/expenses", factory.MakeCreateExpenseController().HandleRequest).Methods("POST")
	s.router.HandleFunc("/expenses/{id}", factory.MakeUpdateExpenseController().HandleRequest).Methods("PUT")

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
	database.StopMongo()
	database.StopRedis()
}
