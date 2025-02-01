package todo

import (
	"github.com/jmoiron/sqlx"
)

type App struct {
	controller *Controller
	Repo       *Repository
}

func New(db *sqlx.DB) *App {
	repo := NewRepository(db)

	return &App{
		controller: NewController(repo),
		Repo:       repo,
	}
}

func (a *App) Shutdown() {}
