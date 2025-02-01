package todo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var ErrTodoNotFound = fmt.Errorf("todo not found")

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(ctx context.Context, description string, dueDate time.Time) (Todo, error) {
	query := "INSERT INTO todo_items (description, due_date) VALUES (?, ?)"

	res, err := r.db.ExecContext(ctx, query, description, dueDate)
	if err != nil {
		return Todo{}, fmt.Errorf("storing todo: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, fmt.Errorf("returning todo id: %w", err)
	}

	return Todo{
		ID:          id,
		Description: description,
		DueDate:     dueDate,
	}, nil
}

func (r *Repository) FindLast(ctx context.Context) (Todo, error) {
	query := "SELECT id, description, due_date FROM todo_items ORDER BY id DESC LIMIT 1"

	var t Todo

	err := r.db.GetContext(ctx, &t, query)

	switch {
	case err == nil:
		return t, nil
	case errors.Is(err, sql.ErrNoRows):
		return Todo{}, fmt.Errorf("%w: %w", ErrTodoNotFound, err)
	default:
		return Todo{}, fmt.Errorf("fetching last todo: %w", err)
	}
}
