// Package repository provides database access for the API's resources.
package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/badrchoubai/api/internal/data"
)

var _ Repository = (*ItemsRepository)(nil)

// Repository reads items from storage.
type Repository interface {
	GetAll() ([]*data.Item, error)
	GetByID(id uuid.UUID) (*data.Item, error)
	Create(item *data.Item) error
}

// ItemsRepository is a Repository backed by a SQL database.
type ItemsRepository struct {
	db *sql.DB
}

// NewItemsRepository returns an ItemsRepository using db.
func NewItemsRepository(db *sql.DB) *ItemsRepository {
	return &ItemsRepository{
		db: db,
	}
}

// GetAll returns every item.
func (r *ItemsRepository) GetAll() ([]*data.Item, error) {
	query := `
SELECT id, name, description
FROM items;
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	items := []*data.Item{}

	for rows.Next() {
		var item data.Item

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// GetByID returns sql.ErrNoRows when no item matches id.
func (r *ItemsRepository) GetByID(id uuid.UUID) (*data.Item, error) {
	query := `
SELECT id, name, description
FROM items
WHERE id = ?;
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var item data.Item

	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// Create inserts item.
func (r *ItemsRepository) Create(item *data.Item) error {
	query := `
INSERT INTO items (id, name, description)
VALUES (?, ?, ?);
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, item.ID.String(), item.Name, item.Description)
	return err
}
