package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"

	"github.com/badrchoubai/api/internal/data"
)

var _ Repository = (*ItemsRepository)(nil)

type Repository interface {
	GetAll() ([]*data.Item, error)
	GetByID(id uuid.UUID) (*data.Item, error)
}
type ItemsRepository struct {
	db *sql.DB
}

func NewItemsRepository(db *sql.DB) *ItemsRepository {
	return &ItemsRepository{
		db: db,
	}
}

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

	defer rows.Close()

	items := []*data.Item{}

	for rows.Next() {
		var item data.Item

		err := rows.Scan(
			&item.Id,
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
		&item.Id,
		&item.Name,
		&item.Description,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
