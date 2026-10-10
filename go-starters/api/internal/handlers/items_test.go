package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/badrchoubai/api/internal/data"
)

// fakeRepository is an in-memory repository.Repository for handler tests.
type fakeRepository struct {
	items []*data.Item
	err   error
}

func (f *fakeRepository) GetAll() ([]*data.Item, error) {
	return f.items, f.err
}

func (f *fakeRepository) GetByID(id uuid.UUID) (*data.Item, error) {
	if f.err != nil {
		return nil, f.err
	}
	for _, item := range f.items {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, sql.ErrNoRows
}

var testItem = &data.Item{
	ID:          uuid.MustParse("c04ab9a0-e398-471d-8cfe-95317ccf0a20"),
	Name:        "First item",
	Description: "An example item",
}

// serve routes the request through a ServeMux so {id} path values are populated.
func serve(t *testing.T, repo *fakeRepository, target string) *httptest.ResponseRecorder {
	t.Helper()

	mux := http.NewServeMux()
	mux.Handle("GET /items", GetItems(repo))
	mux.Handle("GET /items/{id}", GetItem(repo))

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, target, http.NoBody))
	return rec
}

func TestGetItems(t *testing.T) {
	rec := serve(t, &fakeRepository{items: []*data.Item{testItem}}, "/items")

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var body struct {
		Items []data.Item `json:"items"`
		Count int         `json:"count"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if body.Count != 1 || len(body.Items) != 1 || body.Items[0] != *testItem {
		t.Errorf("body = %+v, want one item %+v", body, *testItem)
	}
}

func TestGetItemsRepositoryError(t *testing.T) {
	rec := serve(t, &fakeRepository{err: errors.New("database unavailable")}, "/items")

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}

func TestGetItem(t *testing.T) {
	repo := &fakeRepository{items: []*data.Item{testItem}}

	tests := []struct {
		name       string
		id         string
		wantStatus int
	}{
		{"existing item", testItem.ID.String(), http.StatusOK},
		{"missing item", "00000000-0000-0000-0000-000000000000", http.StatusNotFound},
		{"invalid id", "not-a-uuid", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := serve(t, repo, "/items/"+tt.id)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantStatus != http.StatusOK {
				return
			}

			var got data.Item
			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decoding response: %v", err)
			}
			if got != *testItem {
				t.Errorf("item = %+v, want %+v", got, *testItem)
			}
		})
	}
}
