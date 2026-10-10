// Package problem writes RFC 9457 Problem Details error responses.
package problem

import (
	"encoding/json"
	"net/http"
)

// ContentType is the media type of a Problem Details response.
const ContentType = "application/problem+json"

// Details is an RFC 9457 Problem Details object. Errors maps a request field
// to its validation messages and is only set on validation problems.
type Details struct {
	Type   string              `json:"type"`
	Title  string              `json:"title"`
	Status int                 `json:"status"`
	Detail string              `json:"detail,omitempty"`
	Errors map[string][]string `json:"errors,omitempty"`
}

// Write responds with a problem for status, explained by detail.
func Write(w http.ResponseWriter, status int, detail string) {
	write(w, &Details{
		Type:   "about:blank",
		Title:  http.StatusText(status),
		Status: status,
		Detail: detail,
	})
}

// WriteValidation responds 400 Bad Request with the given field errors.
func WriteValidation(w http.ResponseWriter, errors map[string][]string) {
	write(w, &Details{
		Type:   "about:blank",
		Title:  "One or more validation errors occurred.",
		Status: http.StatusBadRequest,
		Errors: errors,
	})
}

func write(w http.ResponseWriter, details *Details) {
	w.Header().Set("Content-Type", ContentType)
	w.WriteHeader(details.Status)
	_ = json.NewEncoder(w).Encode(details)
}

// NotFoundHandler responds 404 Not Found to every request.
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		Write(w, http.StatusNotFound, "the requested resource does not exist")
	})
}
