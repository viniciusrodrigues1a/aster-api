package main

import (
	"expense-service/cmd/expense-service/application/consumer"
	"expense-service/cmd/expense-service/external/messaging"
	"expense-service/cmd/expense-service/server"
	"expense-service/cmd/expense-service/server/factory"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	srv := server.NewServer()

	go func() {
		srv.Start()
	}()

	m := messaging.New()

	go func() {
		productEventStateConsumer := consumer.NewProductEventStateConsumer(m, factory.MakeProductRedisStateStoreRepository())
		productEventStateConsumer.Consume()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	var _ = <-ch

	fmt.Printf("\nGracefully shutting down...\n")
	srv.Stop()
}
