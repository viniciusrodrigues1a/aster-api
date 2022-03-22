package main

import (
	"account-service/cmd/account-service/server"
	"fmt"
	"os"
	"os/signal"
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
