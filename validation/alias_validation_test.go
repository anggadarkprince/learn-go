package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAliasValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterAlias("username", "required,alphanum,min=3,max=20")
	validate.RegisterAlias("password", "required,min=8,max=20")

	// Define a struct with an alias
	type User struct {
		Username string `validate:"username"`
		Password string `validate:"password"`
	}

	user := User{Username: "user123", Password: "securepassword"}

	// Validate the struct
	err := validate.Struct(user)
	if err != nil {
		t.Fatalf("Expected no validation error, but got: %v", err)
	}
}