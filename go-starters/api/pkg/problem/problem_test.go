package problem

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrite(t *testing.T) {
	rec := httptest.NewRecorder()

	Write(rec, http.StatusNotFound, "item not found")

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
	if got := rec.Header().Get("Content-Type"); got != ContentType {
		t.Errorf("Content-Type = %q, want %q", got, ContentType)
	}

	var body map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decoding body: %v", err)
	}
	want := map[string]any{
		"type":   "about:blank",
		"title":  "Not Found",
		"status": float64(http.StatusNotFound),
		"detail": "item not found",
	}
	if len(body) != len(want) {
		t.Errorf("body = %v, want %v", body, want)
	}
	for key, value := range want {
		if body[key] != value {
			t.Errorf("body[%q] = %v, want %v", key, body[key], value)
		}
	}
}

func TestWriteValidation(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteValidation(rec, map[string][]string{"name": {"name is required"}})

	var details Details
	if err := json.NewDecoder(rec.Body).Decode(&details); err != nil {
		t.Fatalf("decoding body: %v", err)
	}
	if details.Status != http.StatusBadRequest || rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d (body %d), want %d", rec.Code, details.Status, http.StatusBadRequest)
	}
	if got := details.Errors["name"]; len(got) != 1 || got[0] != "name is required" {
		t.Errorf("errors = %v, want name is required", details.Errors)
	}
}
