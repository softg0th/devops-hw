package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"restservice/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStores(c *gin.Context) {
	ctx := c.Request.Context()
	stores, err := h.StoreRepo.GetAllStores(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stores"})
		return
	}
	c.JSON(http.StatusOK, stores)
}

func (h *Handler) CreateStore(c *gin.Context) {
	var store entities.NewStore
	ctx := c.Request.Context()

	if err := c.BindJSON(&store); err != nil {
		log.Printf("[CreateStore] BindJSON failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	insertedID, err := h.StoreRepo.CreateStore(ctx, store)
	if err != nil {
		log.Printf("[CreateStore] DB insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]string{
		"updation_type": "store",
		"updation_name": store.Name,
	}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://bot-service:8080/new_item/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[CreateStore] Failed to notify bot-service: %v", err)
	} else {
		defer resp.Body.Close()
	}

	c.JSON(http.StatusCreated, gin.H{"id": insertedID})
}

func (h *Handler) DeleteStore(c *gin.Context) {
	id := c.Query("id")
	ctx := c.Request.Context()

	deletedId := h.StoreRepo.DeleteStore(ctx, id)
	if deletedId == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": deletedId})
}
