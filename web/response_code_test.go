package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getProductItem(writter http.ResponseWriter, request *http.Request) {
	search := request.URL.Query().Get("search")
	if search == "" {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Search parameter is required"))
		return
	}
	writter.WriteHeader(http.StatusOK)
	writter.Write([]byte("Search: " + search))
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products?search=", nil)
	recorder := httptest.NewRecorder()

	getProductItem(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Should bad request, got %d", response.StatusCode)
	}
}