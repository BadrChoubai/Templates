package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/badrchoubai/api/internal/data"
	"github.com/badrchoubai/api/pkg/problem"
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

func (f *fakeRepository) Create(item *data.Item) error {
	if f.err != nil {
		return f.err
	}
	f.items = append(f.items, item)
	return nil
}

var testItem = &data.Item{
	ID:          uuid.MustParse("c04ab9a0-e398-471d-8cfe-95317ccf0a20"),
	Name:        "First item",
	Description: "An example item",
}

// serve routes the request through a ServeMux so {id} path values are populated.
func serve(t *testing.T, repo *fakeRepository, method, target string, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()

	mux := http.NewServeMux()
	mux.Handle("GET /items", GetItems(repo))
	mux.Handle("GET /items/{id}", GetItem(repo))
	mux.Handle("POST /items", CreateItem(repo))

	if body == nil {
		body = http.NoBody
	}

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(method, target, body))
	return rec
}

// decodeProblem asserts the response is a Problem Details object with status.
func decodeProblem(t *testing.T, rec *httptest.ResponseRecorder, status int) problem.Details {
	t.Helper()

	if rec.Code != status {
		t.Fatalf("status = %d, want %d", rec.Code, status)
	}
	if got := rec.Header().Get("Content-Type"); got != problem.ContentType {
		t.Errorf("Content-Type = %q, want %q", got, problem.ContentType)
	}

	var details problem.Details
	if err := json.NewDecoder(rec.Body).Decode(&details); err != nil {
		t.Fatalf("decoding problem: %v", err)
	}
	if details.Status != status {
		t.Errorf("problem status = %d, want %d", details.Status, status)
	}
	return details
}

func TestGetItems(t *testing.T) {
	rec := serve(t, &fakeRepository{items: []*data.Item{testItem}}, http.MethodGet, "/items", nil)

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
	rec := serve(t, &fakeRepository{err: errors.New("database unavailable")}, http.MethodGet, "/items", nil)

	details := decodeProblem(t, rec, http.StatusInternalServerError)
	if strings.Contains(details.Detail, "database unavailable") {
		t.Errorf("detail %q leaks the internal error", details.Detail)
	}
}

func TestGetItem(t *testing.T) {
	repo := &fakeRepository{items: []*data.Item{testItem}}

	rec := serve(t, repo, http.MethodGet, "/items/"+testItem.ID.String(), nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got data.Item
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if got != *testItem {
		t.Errorf("item = %+v, want %+v", got, *testItem)
	}
}

func TestGetItemNotFound(t *testing.T) {
	repo := &fakeRepository{items: []*data.Item{testItem}}

	for _, id := range []string{"00000000-0000-0000-0000-000000000000", "not-a-uuid"} {
		t.Run(id, func(t *testing.T) {
			rec := serve(t, repo, http.MethodGet, "/items/"+id, nil)

			decodeProblem(t, rec, http.StatusNotFound)
		})
	}
}

func TestCreateItem(t *testing.T) {
	repo := &fakeRepository{}

	rec := serve(t, repo, http.MethodPost, "/items",
		strings.NewReader(`{"name": "New item", "description": "Created in a test"}`))

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	var got data.Item
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decoding response: %v", err)
	}
	if got.ID == uuid.Nil || got.Name != "New item" || got.Description != "Created in a test" {
		t.Errorf("item = %+v, want the created item with an ID", got)
	}
	if want := "/items/" + got.ID.String(); rec.Header().Get("Location") != want {
		t.Errorf("Location = %q, want %q", rec.Header().Get("Location"), want)
	}
	if len(repo.items) != 1 || repo.items[0].ID != got.ID {
		t.Errorf("repository items = %+v, want the created item", repo.items)
	}
}

func TestCreateItemValidation(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantFields []string
	}{
		{"missing name", `{"description": "No name"}`, []string{"name"}},
		{"blank name", `{"name": "   "}`, []string{"name"}},
		{"name too long", `{"name": "` + strings.Repeat("a", 256) + `"}`, []string{"name"}},
		{
			"name and description invalid",
			`{"name": "", "description": "` + strings.Repeat("a", 1025) + `"}`,
			[]string{"name", "description"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRepository{}

			rec := serve(t, repo, http.MethodPost, "/items", strings.NewReader(tt.body))

			details := decodeProblem(t, rec, http.StatusBadRequest)
			if len(details.Errors) != len(tt.wantFields) {
				t.Errorf("errors = %v, want errors for %v", details.Errors, tt.wantFields)
			}
			for _, field := range tt.wantFields {
				if len(details.Errors[field]) == 0 {
					t.Errorf("errors = %v, missing %q", details.Errors, field)
				}
			}
			if len(repo.items) != 0 {
				t.Errorf("invalid request stored %d items", len(repo.items))
			}
		})
	}
}

func TestCreateItemInvalidJSON(t *testing.T) {
	rec := serve(t, &fakeRepository{}, http.MethodPost, "/items", strings.NewReader(`{"name":`))

	decodeProblem(t, rec, http.StatusBadRequest)
}

func TestCreateItemRepositoryError(t *testing.T) {
	repo := &fakeRepository{err: errors.New("database unavailable")}

	rec := serve(t, repo, http.MethodPost, "/items", strings.NewReader(`{"name": "New item"}`))

	decodeProblem(t, rec, http.StatusInternalServerError)
}
