package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Hello, World!"))
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	HelloHandler(recorder, request)
	response := recorder.Result()
	if response.StatusCode != http.StatusOK {	
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	//body, _ := io.ReadAll(response.Body)
	//bodyString := string(body)

	body := make([]byte, 1024)
	n, err := response.Body.Read(body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(body[:n]) != "Hello, World!" {
		t.Errorf("Expected body 'Hello, World!', got '%s'", string(body[:n]))
	}
	if response.Header.Get("Content-Type") != "text/html" {
		t.Errorf("Expected Content-Type 'text/html', got '%s'", response.Header.Get("Content-Type"))
	}
}	
