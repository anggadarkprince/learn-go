package web

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T) {
	// Initialize the server
	server := http.Server{
		Addr:   "localhost:8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Hello, World!"))
		fmt.Fprint(w, "Hello, World!")
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	// Initialize the server
	server := http.Server{
		Addr:   "localhost:8080",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})

	mux.HandleFunc("/about/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Contact Page")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) { // add slash in the bacl
		fmt.Fprint(w, "Images") // This will match /images/ and /images/some-image.png
	})
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail Images")
	})

	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandleWithParams(t *testing.T) {
	// Initialize the server
	server := http.Server{
		Addr:   "localhost:8080",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		postId := r.PathValue("id")
		fmt.Fprintf(w, "Post: %s", postId)
	})

	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	// Initialize the server
	server := http.Server{
		Addr:   "localhost:8080",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.RemoteAddr)
		fmt.Fprintln(w, r.Host)
	})

	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}