package main

import (
	"dependency-injection/helper"
	"dependency-injection/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
    return validator.New()
}


func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}