// Package handlers implements the API's HTTP handlers.
package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"

	"github.com/badrchoubai/api/internal/data"
	"github.com/badrchoubai/api/internal/repository"
	"github.com/badrchoubai/api/pkg/encoding"
	"github.com/badrchoubai/api/pkg/problem"
)

const (
	maxNameLength        = 255
	maxDescriptionLength = 1024
	maxRequestBodyBytes  = 1 << 20
)

// GetItems responds with every item and the item count.
func GetItems(itemsRepository repository.Repository) http.Handler {
	type ListResponse struct {
		Items []*data.Item `json:"items"`
		Count int          `json:"count"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		items, err := itemsRepository.GetAll()
		if err != nil {
			problem.Write(w, http.StatusInternalServerError, "failed to load items")
			return
		}
		response := &ListResponse{
			Items: items,
			Count: len(items),
		}

		ec := encoding.NewEncoderDecoder()

		_ = ec.EncodeResponse(w, http.StatusOK, response)
	})
}

// GetItem responds with the item identified by the {id} path value. An id
// that is not a UUID cannot match an item, so it is reported as not found.
func GetItem(itemsRepository repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			problem.Write(w, http.StatusNotFound, "item not found")
			return
		}

		item, err := itemsRepository.GetByID(id)
		if errors.Is(err, sql.ErrNoRows) {
			problem.Write(w, http.StatusNotFound, "item not found")
			return
		}
		if err != nil {
			problem.Write(w, http.StatusInternalServerError, "failed to load item")
			return
		}

		ec := encoding.NewEncoderDecoder()

		_ = ec.EncodeResponse(w, http.StatusOK, item)
	})
}

// CreateItemRequest is the request body accepted by CreateItem.
type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate returns the request's field errors, keyed by JSON field name.
func (req *CreateItemRequest) Validate() map[string][]string {
	errs := map[string][]string{}

	switch {
	case strings.TrimSpace(req.Name) == "":
		errs["name"] = append(errs["name"], "name is required")
	case utf8.RuneCountInString(req.Name) > maxNameLength:
		errs["name"] = append(errs["name"], "name must be at most 255 characters")
	}
	if utf8.RuneCountInString(req.Description) > maxDescriptionLength {
		errs["description"] = append(errs["description"], "description must be at most 1024 characters")
	}

	return errs
}

// CreateItem validates the request body, stores a new item and responds 201
// Created with the item and its location.
func CreateItem(itemsRepository repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

		var req CreateItemRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			problem.Write(w, http.StatusBadRequest, "request body must be a valid JSON object")
			return
		}
		if errs := req.Validate(); len(errs) > 0 {
			problem.WriteValidation(w, errs)
			return
		}

		item := &data.Item{
			ID:          uuid.New(),
			Name:        req.Name,
			Description: req.Description,
		}
		if err := itemsRepository.Create(item); err != nil {
			problem.Write(w, http.StatusInternalServerError, "failed to create item")
			return
		}

		w.Header().Set("Location", "/items/"+item.ID.String())

		ec := encoding.NewEncoderDecoder()

		_ = ec.EncodeResponse(w, http.StatusCreated, item)
	})
}
