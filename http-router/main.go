package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello, World!"))
	})
	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {			
		name := ps.ByName("name")
		w.Write([]byte("Hello, " + name + "!"))
	})

	server := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	server.ListenAndServe();
}