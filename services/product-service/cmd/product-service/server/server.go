package server

import (
	"context"
	"log"
	"net/http"
	"product-service/cmd/product-service/server/database"
	"product-service/cmd/product-service/server/factory"
	"time"

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
		Addr:    "localhost:8084",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	s.router.Use(middleware.AuthorizationMiddleware)
	s.router.HandleFunc("/products", factory.MakeCreateProductController().HandleRequest).Methods("POST")
	s.router.HandleFunc("/products/{id}", factory.MakeUpdateProductController().HandleRequest).Methods("PUT")
	s.router.HandleFunc("/products/{id}", factory.MakeDeleteProductController().HandleRequest).Methods("DELETE")

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
	database.StopMongo()
	database.StopRedis()
}
