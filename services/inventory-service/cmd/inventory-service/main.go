package main

import (
	"fmt"
	"inventory-service/cmd/inventory-service/application/consumer"
	"inventory-service/cmd/inventory-service/external/factory"
	"inventory-service/cmd/inventory-service/external/messaging"
	"inventory-service/cmd/inventory-service/server"
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
		createInventoryCommandConsumer := consumer.NewCreateInventoryCommandConsumer(m, factory.MakeCreateInventoryUseCase())
		createInventoryCommandConsumer.Consume()
	}()

	go func() {
		expenseEventStateConsumer := consumer.NewExpenseEventStateConsumer(m, factory.MakeAddExpenseToInventoryUseCase())
		expenseEventStateConsumer.Consume()
	}()

	go func() {
		transactionEventStateConsumer := consumer.NewTransactionEventStateConsumer(m, factory.MakeAddTransactionToInventoryUseCase())
		transactionEventStateConsumer.Consume()
	}()

	go func() {
		productEventStateConsumer := consumer.NewProductEventStateConsumer(m, factory.MakeAddProductToInventoryUseCase())
		productEventStateConsumer.Consume()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	var _ = <-ch

	fmt.Printf("\nGracefully shutting down...\n")
}
