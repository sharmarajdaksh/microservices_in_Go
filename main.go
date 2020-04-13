package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/sharmarajdaksh/microservices_in_Go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// Custom ServeMux using gorilla/mux
	sm := mux.NewRouter()

	// .Methods("GET") filters requests to give us only GET requests
	// The Subrouter converts this filtered routes as a router
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	// Extract id from URL
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

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
