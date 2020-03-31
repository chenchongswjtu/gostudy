// http 限速 NewLimiter
package http

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// Wrap the servemux with the limit middleware.
	log.Println("Listening on :4000...")
	http.ListenAndServe(":4000", limit(mux))
}
