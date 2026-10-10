package routes

import (
	"net/http"
)

// RootHandler handles requests that match no other route.
func RootHandler() http.Handler {
	return http.NotFoundHandler()
}
