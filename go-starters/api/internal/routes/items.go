// Package routes registers the API's routes on a ServeMux.
package routes

import (
	"database/sql"
	"net/http"

	"github.com/badrchoubai/api/internal/handlers"
	"github.com/badrchoubai/api/internal/repository"
)

// NewItemsRouter registers the items routes on mux.
func NewItemsRouter(mux *http.ServeMux, db *sql.DB) {
	itemsRepository := repository.NewItemsRepository(db)

	addRoutes(mux, itemsRepository)
}

func addRoutes(mux *http.ServeMux, itemsRepository *repository.ItemsRepository) {
	mux.Handle("GET /items", handlers.GetItems(itemsRepository))
	mux.Handle("GET /items/{id}", handlers.GetItem(itemsRepository))
}
