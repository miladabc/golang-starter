package todo

import "time"

type Todo struct {
	ID          int64     `db:"id"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
}
