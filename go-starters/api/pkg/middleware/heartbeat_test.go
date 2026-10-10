package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeartbeat(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})
	handler := Heartbeat(next, "/health")

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{"GET health", http.MethodGet, "/health", http.StatusOK},
		{"HEAD health", http.MethodHead, "/health", http.StatusOK},
		{"POST health passes through", http.MethodPost, "/health", http.StatusTeapot},
		{"other path passes through", http.MethodGet, "/items", http.StatusTeapot},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, http.NoBody))

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
