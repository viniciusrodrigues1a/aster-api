package main

import (
	"fmt"
	"os"
	"os/signal"
	"transaction-service/cmd/transaction-service/server"
)

func main() {
	srv := server.NewServer()

	go func() {
		srv.Start()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	var _ = <-ch

	fmt.Printf("\nGracefully shutting down...\n")
	srv.Stop()
}
