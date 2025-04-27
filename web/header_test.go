package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func showProducts(writter http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	writter.Header().Add("X-Powered-By", "Go Web")
	writter.WriteHeader(http.StatusOK)
	fmt.Fprintf(writter, "Content-Type: %s\n", contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	
	showProducts(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by")) // in-case sensitive
}