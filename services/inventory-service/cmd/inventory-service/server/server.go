package server

import (
	"context"
	"inventory-service/cmd/inventory-service/external/database"
	"inventory-service/cmd/inventory-service/external/factory"
	"log"
	"net/http"
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
		Addr:    "localhost:8085",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	s.router.Use(middleware.AuthorizationMiddleware)
	s.router.HandleFunc("/inventories/{id}", factory.MakeListInventoryController().HandleRequest).Methods("GET")

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
	database.StopMongo()
	database.StopRedis()
}
