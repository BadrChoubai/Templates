package repository

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"

	"github.com/badrchoubai/api/pkg/database"
)

// newTestDB returns a database in a temporary directory with every up
// migration applied, so tests run against the real schema and seed data.
func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := database.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("opening database: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	migrations, err := filepath.Glob("../../migrations/migrations/*.up.sql")
	if err != nil || len(migrations) == 0 {
		t.Fatalf("finding migrations: %v (found %d)", err, len(migrations))
	}
	for _, path := range migrations {
		query, err := os.ReadFile(path) //nolint:gosec // paths come from the repository's migrations directory
		if err != nil {
			t.Fatalf("reading %s: %v", path, err)
		}
		if _, err := db.Exec(string(query)); err != nil {
			t.Fatalf("applying %s: %v", path, err)
		}
	}

	return db
}

func TestItemsRepositoryGetAll(t *testing.T) {
	repo := NewItemsRepository(newTestDB(t))

	items, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("len(items) = %d, want 2 seeded items", len(items))
	}
}

func TestItemsRepositoryGetByID(t *testing.T) {
	repo := NewItemsRepository(newTestDB(t))
	id := uuid.MustParse("c04ab9a0-e398-471d-8cfe-95317ccf0a20")

	item, err := repo.GetByID(id)
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if item.ID != id || item.Name != "First item" {
		t.Errorf("item = %+v, want seeded item %s", item, id)
	}
}

func TestItemsRepositoryGetByIDNotFound(t *testing.T) {
	repo := NewItemsRepository(newTestDB(t))

	_, err := repo.GetByID(uuid.Nil)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("err = %v, want sql.ErrNoRows", err)
	}
}
