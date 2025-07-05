package repository

import (
	"context"
	"restservice/internal/domain/entities"
)

type AnimalRepository interface {
	CreateAnimal(ctx context.Context, animal entities.NewAnimal) (int, error)
	GetAnimalsByFilter(ctx context.Context, filter string, value interface{}) ([]entities.Animal, error)
	GetAllAnimals(ctx context.Context) ([]entities.Animal, error)
	UpdateAnimal(ctx context.Context, animal entities.UpdatedAnimal) (int, error)
	DeleteAnimal(ctx context.Context, id string) (int, error)
}

type StoreRepository interface {
	CreateStore(ctx context.Context, store entities.NewStore) (int, error)
	GetAllStores(ctx context.Context) ([]entities.ExistingStoreWithAddress, error)
	DeleteStore(ctx context.Context, id string) int
}
