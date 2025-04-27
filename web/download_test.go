package web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writter http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writter, "Bad request")
		return
	}

	writter.Header().Add("Content-Disposition", "attachment; filename=\"" + file + "\"")
	http.ServeFile(writter, request, "./resources/" + file)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}