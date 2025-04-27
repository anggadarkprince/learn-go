package web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler home executed")
		fmt.Fprint(writer, "Hello home")
	})
	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler about executed")
		fmt.Fprint(writer, "Hello about")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler panic executed")
		panic("Something went wrong")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}
	errorHandlerMiddleware := ErrorHandler{
		Handler: &logMiddleware,
	}
	server := http.Server{
		Addr: "localhost:8080",
		Handler: &errorHandlerMiddleware,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error occured")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}