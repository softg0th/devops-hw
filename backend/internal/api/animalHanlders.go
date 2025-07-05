package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"restservice/internal/domain/entities"

	"github.com/gin-gonic/gin"

	"restservice/internal/infra"
)

func (h *Handler) CreateAnimal(c *gin.Context) {
	var animal entities.NewAnimal
	ctx := c.Request.Context()
	infra.RequestsTotal.Inc()

	if err := c.BindJSON(&animal); err != nil {
		log.Printf("[CreateAnimal] BindJSON failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	insertedID, err := h.AnimalRepo.CreateAnimal(ctx, animal)
	if err != nil {
		log.Printf("[CreateAnimal] DB insert failed: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	data := map[string]string{
		"updation_type": "animal",
		"updation_name": animal.Name,
	}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://bot-service:8080/new_item/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[CreateAnimal] Failed to notify bot-service: %v", err)
	} else {
		defer resp.Body.Close()
	}

	c.JSON(http.StatusCreated, gin.H{"id": insertedID})
}

func (h *Handler) GetAnimalsByFilter(c *gin.Context) {
	filter := c.Query("filter")
	value := c.Query("value")
	ctx := c.Request.Context()
	infra.RequestsTotal.Inc()

	if filter == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing filter or value"})
		return
	}

	animals, err := h.AnimalRepo.GetAnimalsByFilter(ctx, filter, value)
	if err != nil {
		log.Printf("[GetAnimalsByFilter] DB query failed: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animals)
}

func (h *Handler) GetAllAnimals(c *gin.Context) {
	ctx := c.Request.Context()
	infra.RequestsTotal.Inc()

	animals, err := h.AnimalRepo.GetAllAnimals(ctx)
	if err != nil {
		log.Printf("[GetAllAnimals] DB query failed: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animals)
}

func (h *Handler) UpdateAnimal(c *gin.Context) {
	var animal entities.UpdatedAnimal
	ctx := c.Request.Context()
	infra.RequestsTotal.Inc()

	if err := c.BindJSON(&animal); err != nil {
		log.Printf("[UpdateAnimal] BindJSON failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	updatedID, err := h.AnimalRepo.UpdateAnimal(ctx, animal)
	if err != nil {
		log.Printf("[UpdateAnimal] DB update failed: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": updatedID})
}

func (h *Handler) DeleteAnimal(c *gin.Context) {
	id := c.Query("id")
	ctx := c.Request.Context()
	infra.RequestsTotal.Inc()

	deletedID, err := h.AnimalRepo.DeleteAnimal(ctx, id)
	if err != nil {
		log.Printf("[DeleteAnimal] Delete failed: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": deletedID})
}
