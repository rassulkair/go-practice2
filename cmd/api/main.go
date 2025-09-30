package main

import (
	"log"
	"net/http"

	"go-practice2/internal/handlers"
	"go-practice2/internal/middleware"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/user", middleware.APIKeyMiddleware(http.HandlerFunc(handlers.UserHandler)))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our Golang application!"))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)

	}
}
