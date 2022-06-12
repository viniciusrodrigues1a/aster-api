package server

import (
	"context"
	"log"
	"net/http"
	"time"
	"transaction-service/cmd/transaction-service/server/database"
	"transaction-service/cmd/transaction-service/server/factory"

	"github.com/gorilla/mux"
	"github.com/viniciusrodrigues1a/aster-api/pkg/server/middleware"
)

type server struct {
	server *http.Server
	router *mux.Router
}

func NewServer() *server {
	router := mux.NewRouter()
	httpServer := &http.Server{
		Handler: router,
		Addr:    "localhost:8083",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	s.router.Use(middleware.AuthorizationMiddleware)
	s.router.HandleFunc("/transactions", factory.MakeCreateTransactionController().HandleRequest).Methods("POST")
	s.router.HandleFunc("/transactions/{id}", factory.MakeUpdateTransactionController().HandleRequest).Methods("PUT")
	s.router.HandleFunc("/transactions/{id}", factory.MakeDeleteTransactionController().HandleRequest).Methods("DELETE")
	s.router.HandleFunc("/transactions/{id}/debit-money", factory.MakeDebitMoneyToTransactionController().HandleRequest).Methods("PATCH")

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
	database.StopMongo()
	database.StopRedis()
	database.StopProductRedis()
}
