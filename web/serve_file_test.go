package web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		http.ServeFile(writer, request, "./resources/404.html")
	}
}

func TestServeFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeFile)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/index.html
var resourceOk string

//go:embed resources/404.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileWithEmbed(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeFileEmbed)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}