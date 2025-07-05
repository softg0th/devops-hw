package repository

import (
	"context"
	"fmt"
	"log"
	"restservice/internal/domain/entities"
	"restservice/internal/domain/entities/queries"
	"restservice/internal/exceptions"
	"strconv"
	"strings"
)

func (r *Repository) GetAnimalByID(id int) (string, error) {
	var animalID string
	err := r.DB.Pool.QueryRow(context.Background(), queries.SelectAnimalById, id).Scan(&animalID)
	if err != nil {
		return "", err
	}
	return animalID, nil
}

func (r *Repository) CreateAnimal(ctx context.Context, newAnimal entities.NewAnimal) (int, error) {
	var insertedID int
	err := r.DB.Pool.QueryRow(ctx, queries.InsertAnimalQuery,
		newAnimal.Name, newAnimal.Type, newAnimal.Color, newAnimal.StoreID, newAnimal.Age, newAnimal.Price).Scan(&insertedID)

	if err != nil {
		log.Printf("[CreateAnimal] Insert failed: %v", err)
		return -1, err
	}
	return insertedID, nil
}

func (r *Repository) GetAllAnimals(ctx context.Context) ([]entities.Animal, error) {
	rows, err := r.DB.Pool.Query(ctx, queries.SelectAllAnimalQuery)
	if err != nil {
		log.Printf("[GetAllAnimals] Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var animals []entities.Animal
	for rows.Next() {
		var animal entities.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Type, &animal.Color, &animal.StoreID, &animal.StoreAddress, &animal.Age, &animal.Price); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return animals, nil
}

func (r *Repository) GetAnimalsByFilter(ctx context.Context, filter string, value interface{}) ([]entities.Animal, error) {
	query := fmt.Sprintf(queries.SelectAnimalsByFilterQuery, filter)
	rows, err := r.DB.Pool.Query(ctx, query, value)
	if err != nil {
		log.Printf("[GetAnimalsByFilter] Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var animals []entities.Animal
	for rows.Next() {
		var animal entities.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Type, &animal.Color, &animal.StoreID, &animal.Age, &animal.Price); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return animals, nil
}

func (r *Repository) UpdateAnimal(ctx context.Context, animal entities.UpdatedAnimal) (int, error) {
	_, err := r.GetAnimalByID(animal.ID)
	if err != nil {
		return -1, exceptions.HandleRunTimeError("Animal ID not found")
	}

	updates := []string{}
	args := []interface{}{animal.ID}
	argPos := 2

	appendField := func(name string, val interface{}) {
		updates = append(updates, fmt.Sprintf("%s = $%d", name, argPos))
		args = append(args, val)
		argPos++
	}

	if animal.Name != "" {
		appendField("name", animal.Name)
	}
	if animal.Type != "" {
		appendField("type", animal.Type)
	}
	if animal.Color != "" {
		appendField("color", animal.Color)
	}
	if animal.StoreID != 0 {
		appendField("store_id", animal.StoreID)
	}
	if animal.Age != 0 {
		appendField("age", animal.Age)
	}
	if animal.Price != 0 {
		appendField("price", animal.Price)
	}

	if len(updates) == 0 {
		return animal.ID, nil
	}

	query := fmt.Sprintf("UPDATE animals SET %s WHERE id = $1 RETURNING id", strings.Join(updates, ", "))

	var updatedID int
	err = r.DB.Pool.QueryRow(ctx, query, args...).Scan(&updatedID)
	if err != nil {
		log.Printf("[UpdateAnimal] Query failed: %v", err)
		return -1, exceptions.HandleRunTimeError("Failed to update animal")
	}
	return updatedID, nil
}

func (r *Repository) DeleteAnimal(ctx context.Context, stringId string) (int, error) {
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return -1, exceptions.HandleRunTimeError("Invalid animal ID format")
	}

	_, err = r.GetAnimalByID(id)
	if err != nil {
		return -1, exceptions.HandleRunTimeError("Animal ID not found")
	}

	var deletedID int
	err = r.DB.Pool.QueryRow(ctx, queries.DeleteAnimalQuery, id).Scan(&deletedID)
	if err != nil {
		log.Printf("[DeleteAnimal] Query failed: %v", err)
		return -1, exceptions.HandleRunTimeError("Failed to delete animal")
	}
	return deletedID, nil
}
