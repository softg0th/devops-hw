package repository

import (
	"restservice/internal/infra"
)

type Repository struct {
	DB *infra.PostgresConnect
}

func NewRepository(db *infra.PostgresConnect) *Repository {
	return &Repository{DB: db}
}

type Handler struct {
	Repo *Repository
}
