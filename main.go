package main

import (
	"context"
	"fmt"
	"go-assesment/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func setupContext() context.Context {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
	return ctx
}
func main() {
	ctx := setupContext()
	http.HandleFunc("/find", handlers.FindHandler)
	http.HandleFunc("/compare", handlers.CompareHandler)
	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/remove", handlers.RemoveHandler)
	http.HandleFunc("/find-longest", handlers.FindLongestHandler)

	go (func() {
		err := http.ListenAndServe(":3001", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	})()

	<-ctx.Done()
	fmt.Println("context canceled, terminating")
}
