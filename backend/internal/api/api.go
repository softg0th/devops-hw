package api

import (
	"restservice/internal/repository"
)

type Handler struct {
	AnimalRepo repository.AnimalRepository
	StoreRepo  repository.StoreRepository
}

func NewHandler(animalRepo repository.AnimalRepository, storeRepo repository.StoreRepository) *Handler {
	return &Handler{
		AnimalRepo: animalRepo,
		StoreRepo:  storeRepo,
	}
}
