package main

import (
	"fmt"
	"os"
	"os/signal"
	"transaction-service/cmd/transaction-service/application/consumer"
	"transaction-service/cmd/transaction-service/external/messaging"
	"transaction-service/cmd/transaction-service/server"
	"transaction-service/cmd/transaction-service/server/factory"
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
