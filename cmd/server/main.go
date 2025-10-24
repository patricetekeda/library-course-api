package main

import (
	"log"
	"net/http"
	"time"

	"github.com/patricetekeda/library-course-api/internal/courses"
	"github.com/patricetekeda/library-course-api/internal/httpapi"
)

func main() {
	// Initialize in-memory courses store and inject into routes.
	store := courses.NewStore()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      httpapi.Routes(store),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server is running at http://localhost:8080")
	log.Println("Press Ctrl+C to stop the server")
	log.Println("Health check endpoint available at http://localhost:8080/healthz")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
