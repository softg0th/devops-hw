package repository

import (
	"context"
	"log"
	"restservice/internal/domain/entities"
	"restservice/internal/domain/entities/queries"
	"strconv"
)

func (r *Repository) GetStoreById(id int) error {
	row := r.DB.Pool.QueryRow(context.Background(), queries.SelectStoreById, id)

	var store entities.ExistingStore
	if err := row.Scan(&store.Id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllStores(ctx context.Context) ([]entities.ExistingStoreWithAddress, error) {
	rows, err := r.DB.Pool.Query(ctx, queries.SelectAllStoreQuery)
	if err != nil {
		log.Printf("DB Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var stores []entities.ExistingStoreWithAddress
	for rows.Next() {
		var store entities.ExistingStoreWithAddress
		if err := rows.Scan(&store.Id, &store.Address); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *Repository) CreateStore(ctx context.Context, newStore entities.NewStore) (int, error) {
	var insertedID int
	err := r.DB.Pool.QueryRow(ctx, queries.InsertStoreQurey, newStore.Name, newStore.Address).
		Scan(&insertedID)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return insertedID, err
}

func (r *Repository) DeleteStore(ctx context.Context, stringId string) int {
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Println(err)
		return -1
	}

	err = r.GetStoreById(id)
	if err != nil {
		log.Println("Store not found")
		return -1
	}

	var deletedID int
	err = r.DB.Pool.QueryRow(ctx, queries.DeleteStoreQurey, id).
		Scan(&deletedID)
	if err != nil {
		log.Println(err)
		return -1
	}
	return deletedID
}
