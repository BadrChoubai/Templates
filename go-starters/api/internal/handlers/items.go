package handlers

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"net/http"

	"github.com/badrchoubai/api/internal/data"
	"github.com/badrchoubai/api/internal/repository"
	"github.com/badrchoubai/api/pkg/encoding"
)

func GetItems(itemsRepository repository.Repository) http.Handler {
	type ListResponse struct {
		Items []*data.Item `json:"items"`
		Count int          `json:"count"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		items, err := itemsRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := &ListResponse{
			Items: items,
			Count: len(items),
		}

		ec := encoding.NewEncoderDecoder()

		err = ec.EncodeResponse(w, http.StatusOK, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func GetItem(itemsRepository repository.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid item id", http.StatusBadRequest)
			return
		}

		item, err := itemsRepository.GetByID(id)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "item not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ec := encoding.NewEncoderDecoder()

		err = ec.EncodeResponse(w, http.StatusOK, item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
