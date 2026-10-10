package routes

import (
	"net/http"
)

func RootHandler() http.Handler {
	return http.NotFoundHandler()
}
