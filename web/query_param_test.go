package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func getProducts(writter http.ResponseWriter, request *http.Request) {
	q := request.URL.Query().Get("q")
	category := request.URL.Query().Get("category")
	merchants := request.URL.Query()["merchant"]

	if (category == "") {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Category is required"))
		return
	}

	writter.WriteHeader(http.StatusOK)
	fmt.Fprintf(writter, "Search: %s\n", q)
	fmt.Fprintf(writter, "Category: %s\n", category)
	fmt.Fprintf(writter, "Merchant: %s\n", strings.Join(merchants, ","))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products?category=electronics&q=iphone&merchant=star&merchant=official", nil)
	recorder := httptest.NewRecorder()
	
	getProducts(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}