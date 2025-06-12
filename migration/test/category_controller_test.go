package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"migration/app"
	"migration/controller"
	"migration/helper"
	"migration/middleware"
	"migration/model/domain"
	"migration/repository"
	"migration/service"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	// Assuming you have a function to initialize your database connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test_sandbox")
	helper.PanicIfError(err)
	
	// Set the maximum number of open connections to the database
	db.SetMaxOpenConns(25)
	
	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(5)
	
	// Set the maximum lifetime of a connection
	db.SetConnMaxLifetime(time.Minute * 10)

	// Set the maximum idle time for connections
	db.SetConnMaxIdleTime(time.Minute * 5)
	
	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE categories")
	helper.PanicIfError(err)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "Test Category"
	}`)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "Test Category", responseBody["data"].(map[string]interface{})["name"])
}


func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": ""
	}`)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "Bad Request", responseBody["status"])
	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Create(context.Background(), tx, domain.Category{Name: "Initial Category"})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "Updated Category"
	}`)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Updated Category", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Create(context.Background(), tx, domain.Category{Name: "Initial Category"})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": ""
	}`)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "Bad Request", responseBody["status"])
	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Create(context.Background(), tx, domain.Category{Name: "Existing Category"})
	tx.Commit()

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Create(context.Background(), tx, domain.Category{Name: "Existing Category"})
	tx.Commit()

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/" + strconv.Itoa(category.Id), nil)
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
}
func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
}

func TestListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Create(context.Background(), tx, domain.Category{Name: "Existing Category"})
	tx.Commit()

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	req.Header.Set("Authorization", "secret-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))

	var categories = responseBody["data"].([]interface{})
	categoryResponse := categories[0].(map[string]interface{})
	assert.Equal(t, category.Id, int(categoryResponse["id"].(float64)))
	assert.Equal(t, category.Name, categoryResponse["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	req.Header.Set("Authorization", "no-token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, "Unauthorized", responseBody["status"])
	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
}