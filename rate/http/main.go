// http 限速 NewLimiter
package main

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// Wrap the serve mux with the limit middleware.
	log.Println("Listening on :4000...")
	err := http.ListenAndServe(":4000", limit(mux))
	if err != nil {
		return
	}
}
