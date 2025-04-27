package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func StoreProduct(writter http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Invalid form data"))
		return
	}
	name := request.FormValue("name")
	price := request.FormValue("price")
	//price := request.PostFormValue("price") // without ParseForm
	if name == "" || price == "" {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("Name and price are required"))
		return
	}
		
	writter.WriteHeader(http.StatusCreated)
	writter.Write([]byte("Product Created\n"))
	writter.Write([]byte(name + "\n"))
	writter.Write([]byte(price + "\n"))
}

func TestStoreProduct(t *testing.T) {
	requestBody := strings.NewReader("name=iphone&price=1000")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/products", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	recorder := httptest.NewRecorder()

	StoreProduct(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}