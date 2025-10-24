package httpapi

import "net/http"

func Routes() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// Return a friendly, descriptive message so callers know the service is healthy.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("library-course-api: service is healthy\n"))
	})

	return mux
}
