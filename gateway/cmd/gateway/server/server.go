package server

import (
	"context"
	"gateway/cmd/gateway/proxy"
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

func New() *server {
	router := mux.NewRouter()

	httpServer := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	return &server{
		server: httpServer,
		router: router,
	}
}

func (s *server) Start() {
	accountsProxy := proxy.New("http://127.0.0.1:8081", "accounts")
	sessionsProxy := proxy.New("http://127.0.0.1:8081", "sessions")
	expensesProxy := proxy.New("http://127.0.0.1:8082", "expenses")
	transactionsProxy := proxy.New("http://127.0.0.1:8083", "transactions")
	productsProxy := proxy.New("http://127.0.0.1:8084", "products")
	inventoriesProxy := proxy.New("http://127.0.0.1:8085", "inventories")

	s.router.Use(middleware.CORSMiddleware)

	s.router.HandleFunc("/accounts/{rest:.*}", accountsProxy.HandleRequest)
	s.router.HandleFunc("/accounts", accountsProxy.HandleRequest)

	s.router.HandleFunc("/sessions/{rest:.*}", sessionsProxy.HandleRequest)
	s.router.HandleFunc("/sessions", sessionsProxy.HandleRequest)

	s.router.HandleFunc("/expenses/{rest:.*}", expensesProxy.HandleRequest)
	s.router.HandleFunc("/expenses", expensesProxy.HandleRequest)

	s.router.HandleFunc("/transactions/{rest:.*}", transactionsProxy.HandleRequest)
	s.router.HandleFunc("/transactions", transactionsProxy.HandleRequest)

	s.router.HandleFunc("/products/{rest:.*}", productsProxy.HandleRequest)
	s.router.HandleFunc("/products", productsProxy.HandleRequest)

	s.router.HandleFunc("/inventories/{rest:.*}", inventoriesProxy.HandleRequest)
	s.router.HandleFunc("/inventories", inventoriesProxy.HandleRequest)

	log.Fatal(s.server.ListenAndServe())
}

func (s *server) Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.server.Shutdown(ctx)
}
