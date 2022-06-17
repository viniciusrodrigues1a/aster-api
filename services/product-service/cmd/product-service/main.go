package main

import (
	"fmt"
	"os"
	"os/signal"
	"product-service/cmd/product-service/application/consumer"
	"product-service/cmd/product-service/external/messaging"
	"product-service/cmd/product-service/server"
	"product-service/cmd/product-service/server/factory"
)

func main() {
	srv := server.NewServer()

	go func() {
		srv.Start()
	}()

	m := messaging.New()

	go func() {
		subtractProductQuantityCommandConsumer := consumer.NewSubtractProductQuantityCommandConsumer(m, factory.MakeSubtractProductQuantityUseCase())
		subtractProductQuantityCommandConsumer.Consume()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	var _ = <-ch

	fmt.Printf("\nGracefully shutting down...\n")
	srv.Stop()
}
