package routes

import (
	"net/http"

	"github.com/badrchoubai/api/pkg/problem"
)

// RootHandler handles requests that match no other route.
func RootHandler() http.Handler {
	return problem.NotFoundHandler()
}
