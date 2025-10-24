package main

import (
"net/http"
 "log"
 "time"
 "github.com/patricetekeda/library-course-api/internal/httpapi"
)


func main () {
	srv := &http.Server {
		Addr: ":8080",
		Handler: httpapi.Routes(),
		ReadTimeout: 5 * time.Second, 
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Println("Server is running at http://localhost:8080")
	log.Println("Press Ctrl+C to stop the server")
	log.Println("Health check endpoint available at http://localhost:8080/healthz")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
