package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/patricetekeda/library-course-api/internal/courses"
)

// Routes returns an http.Handler wired with the provided courses store.
// This makes the routes easy to test and allows callers to control the store
// lifecycle (in-memory, mocked, or backed by a DB).
func Routes(store *courses.Store) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// Return a friendly, descriptive message so callers know the service is healthy.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("library-course-api: service is healthy\n"))
	})

	// List courses
	mux.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		courseList := store.GetAllCourses()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(courseList); err != nil {
			http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
			return
		}
	})

	return mux
}
