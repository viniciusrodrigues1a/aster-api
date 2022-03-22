package main

import (
	"fmt"
	"inventory-service/cmd/inventory-service/application/consumer"
	"inventory-service/cmd/inventory-service/external/factory"
	"inventory-service/cmd/inventory-service/external/messaging"
	"os"
	"os/signal"
)

func main() {
	go func() {
		m := messaging.New()
		consumer := consumer.NewAccountCreationConsumer(m, factory.MakeCreateInventoryUseCase())
		consumer.Consume()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	var _ = <-ch

	fmt.Printf("\nGracefully shutting down...\n")
}
