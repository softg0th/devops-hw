package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restservice/internal/domain/entities"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct{}

type erroringRepo struct {
	mockRepo
}

func (m *mockRepo) CreateAnimal(ctx context.Context, animal entities.NewAnimal) (int, error) {
	return 1, nil
}

func (m *mockRepo) GetAnimalsByFilter(ctx context.Context, filter string, value interface{}) ([]entities.Animal, error) {
	return []entities.Animal{{ID: 1, Name: "Test"}}, nil
}

func (m *mockRepo) GetAllAnimals(ctx context.Context) ([]entities.Animal, error) {
	return []entities.Animal{{ID: 1, Name: "Test"}}, nil
}

func (m *mockRepo) UpdateAnimal(ctx context.Context, animal entities.UpdatedAnimal) (int, error) {
	return animal.ID, nil
}

func (m *mockRepo) DeleteAnimal(ctx context.Context, id string) (int, error) {
	return 1, nil
}

func (m *mockRepo) CreateStore(ctx context.Context, store entities.NewStore) (int, error) {
	return 42, nil
}

func (m *mockRepo) GetAllStores(ctx context.Context) ([]entities.ExistingStoreWithAddress, error) {
	return []entities.ExistingStoreWithAddress{{Id: 1, Address: "Test Address"}}, nil
}

func (m *mockRepo) DeleteStore(ctx context.Context, id string) int {
	return 42
}

func newHandlerWithMocks() *Handler {
	return &Handler{
		AnimalRepo: &mockRepo{},
		StoreRepo:  &mockRepo{},
	}
}

func newTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

// --- Animal Handlers ---

func TestCreateAnimal_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()

	body, _ := json.Marshal(entities.NewAnimal{Name: "Leo", Type: "Cat", Color: "White", StoreID: 1, Age: 2, Price: 100})
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	h.CreateAnimal(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateAnimal_BadJSON(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()

	req, _ := http.NewRequest("POST", "/animals", bytes.NewBufferString("bad-json"))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	h.CreateAnimal(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllAnimals(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/animals", nil)
	c.Request = req
	h.GetAllAnimals(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAnimalsByFilter_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/animals/filter?filter=type&value=Cat", nil)
	c.Request = req
	h.GetAnimalsByFilter(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAnimalsByFilter_MissingParam(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/animals/filter", nil)
	c.Request = req
	h.GetAnimalsByFilter(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateAnimal_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	body, _ := json.Marshal(entities.UpdatedAnimal{ID: 1, Name: "Updated"})
	req, _ := http.NewRequest("PUT", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.UpdateAnimal(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateAnimal_BadJSON(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("PUT", "/animals", bytes.NewBufferString("bad-json"))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.UpdateAnimal(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteAnimal_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("DELETE", "/animals?id=123", nil)
	c.Request = req
	h.DeleteAnimal(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

// --- Store Handlers ---

func TestGetStores_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/stores", nil)
	c.Request = req
	h.GetStores(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateStore_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	body, _ := json.Marshal(entities.NewStore{Name: "Test", Address: "123 St"})
	req, _ := http.NewRequest("POST", "/stores", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.CreateStore(c)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateStore_BadJSON(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("POST", "/stores", bytes.NewBufferString("bad-json"))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.CreateStore(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteStore_Success(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()
	req, _ := http.NewRequest("DELETE", "/stores?id=123", nil)
	c.Request = req
	h.DeleteStore(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

// --- Constructor Coverage ---

func TestNewHandler(t *testing.T) {
	h := NewHandler(&mockRepo{}, &mockRepo{})
	assert.NotNil(t, h)
}

func (e *erroringRepo) GetAllAnimals(ctx context.Context) ([]entities.Animal, error) {
	return nil, assert.AnError
}

func TestGetAllAnimals_Error(t *testing.T) {
	h := &Handler{AnimalRepo: &erroringRepo{}}
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/animals", nil)
	c.Request = req
	h.GetAllAnimals(c)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func (e *erroringRepo) GetAllStores(ctx context.Context) ([]entities.ExistingStoreWithAddress, error) {
	return nil, assert.AnError
}

func TestGetStores_Error(t *testing.T) {
	h := &Handler{StoreRepo: &erroringRepo{}}
	c, w := newTestContext()
	req, _ := http.NewRequest("GET", "/stores", nil)
	c.Request = req
	h.GetStores(c)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

type zeroDeleteRepo struct {
	mockRepo
}

func (z *zeroDeleteRepo) DeleteStore(ctx context.Context, id string) int {
	return 0
}

func TestDeleteStore(t *testing.T) {
	h := &Handler{StoreRepo: &zeroDeleteRepo{}}
	c, w := newTestContext()
	req, _ := http.NewRequest("DELETE", "/stores?id=42", nil)
	c.Request = req
	h.DeleteStore(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAnimal_EmptyID(t *testing.T) {
	h := newHandlerWithMocks()
	c, w := newTestContext()

	req, _ := http.NewRequest("DELETE", "/animals?id=", nil)
	c.Request = req

	h.DeleteAnimal(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

type failingRepo struct {
	mockRepo
}

func (f *failingRepo) CreateAnimal(ctx context.Context, animal entities.NewAnimal) (int, error) {
	return 0, assert.AnError
}

func TestCreateAnimal_RepoError(t *testing.T) {
	h := &Handler{AnimalRepo: &failingRepo{}}
	c, w := newTestContext()
	body, _ := json.Marshal(entities.NewAnimal{Name: "Fail", Type: "Cat", Color: "Black", StoreID: 1, Age: 1, Price: 50})
	req, _ := http.NewRequest("POST", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.CreateAnimal(c)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func (f *failingRepo) UpdateAnimal(ctx context.Context, animal entities.UpdatedAnimal) (int, error) {
	return 0, assert.AnError
}

func TestUpdateAnimal_RepoError(t *testing.T) {
	h := &Handler{AnimalRepo: &failingRepo{}}
	c, w := newTestContext()
	body, _ := json.Marshal(entities.UpdatedAnimal{ID: 1, Name: "Bad"})
	req, _ := http.NewRequest("PUT", "/animals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req
	h.UpdateAnimal(c)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func (f *failingRepo) DeleteAnimal(ctx context.Context, id string) (int, error) {
	return 0, assert.AnError
}

func TestDeleteAnimal_RepoError(t *testing.T) {
	h := &Handler{AnimalRepo: &failingRepo{}}
	c, w := newTestContext()
	req, _ := http.NewRequest("DELETE", "/animals?id=123", nil)
	c.Request = req
	h.DeleteAnimal(c)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
