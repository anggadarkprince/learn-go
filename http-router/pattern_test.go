package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPatternNamedParameter(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		fmt.Fprint(w, "Product " + id + " Item " + itemId)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 2", string(bytes))
}


func TestPatternCatchAll(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprint(w, "Image " + image)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/avatar.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image /small/avatar.jpg", string(bytes))
}