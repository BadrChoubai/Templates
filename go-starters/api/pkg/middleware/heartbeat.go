// Package middleware provides reusable HTTP middleware.
package middleware

import (
	"net/http"
	"strings"
)

// Heartbeat responds 200 OK to GET and HEAD requests for endpoint and passes
// every other request to handler.
func Heartbeat(handler http.Handler, endpoint string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "GET" || r.Method == "HEAD") && strings.EqualFold(r.URL.Path, endpoint) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
