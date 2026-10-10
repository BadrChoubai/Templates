// Package data defines the API's domain types.
package data

import "github.com/google/uuid"

// Item is the example resource served by the API.
type Item struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
