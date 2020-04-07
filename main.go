package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-microservice-basic/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// Custom ServeMux
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Customer server using the custom servemux
	// With timeouts specified
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// The sigChan channel listens to os iterrupts
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate; graceful shutdown", sig)

	// Ensure a graceful timeout: If processes are running when an interrupt is received,
	// Wait until those processes end or until the end of 30 seconds.
	ctx, c := context.WithTimeout(context.Background(), 30*time.Second)
	defer c()
	s.Shutdown(ctx)
}
