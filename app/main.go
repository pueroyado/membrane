package main

import (
	"context"
	"demo/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// @title Product API
// @version 1.0
// @description This is a service product
// @host localhost
// @BasePath /
func main() {
	fmt.Println("Start fn main")

	serverApi := server.Create()
	wg := sync.WaitGroup{}
	ctx := context.Background()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := serverApi.Start()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	err := serverApi.Shutdown(ctx)
	if err == nil {
		log.Println("Shutdown: halted active connections", err)
	}

	wg.Wait()
}
