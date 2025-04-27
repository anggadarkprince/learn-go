package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// Initialize the server
	server := http.Server{
		Addr:   "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
