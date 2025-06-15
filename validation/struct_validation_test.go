package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestStructValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	type User struct {
		Name     string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8"`
	}

	user := User{
		Name:     "John Doe",
		Email:    "john@mail.com",
		Password: "password123",
	}
	err := validate.Struct(user)
	if err != nil {
		t.Fatalf("Expected no validation error, but got: %v", err)
	}

	userInvalid := User{
		Name:     "",
		Email:    "invalid-email",
		Password: "short",
	}
	errInvalid := validate.Struct(userInvalid)
	if errInvalid == nil {
		t.Fatal("Expected validation error, but got none")
	} else {
		// Get validation errors
		validationErrors := errInvalid.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			t.Logf("Field: %s, Tag: %s, Error: %s", fieldError.Field(), fieldError.Tag(), fieldError.Error())
		}
	}
}

type RegisterUser struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

func TestRegisterUserValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	user := RegisterUser{
		Username: "testuser",
		Email: "test@mail.com",
		Password: "password123",
		ConfirmPassword: "password123",
	}
	err := validate.Struct(user)
	if err != nil {
		t.Fatalf("Expected no validation error, but got: %v", err)
	}
}


type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
}

type UserProfile struct {
	Name    string  `validate:"required"`
	Address Address `validate:"required"`
}

type UserAccount struct {
	Name    string  `validate:"required"`
	Address []Address `validate:"required,min=3,dive"` // add `dive` to validate each item in the slice + address is required and minimum 3 items
}

func TestNestedStruct(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	profile := UserProfile{
		Name: "Jane Doe",
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}
	err := validate.Struct(profile)
	if err != nil {
		t.Fatalf("Expected no validation error, but got: %v", err)
	}
}

func TestCollectionStruct(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	account := UserAccount{
		Name: "John Smith",
		Address: []Address{
			{Street: "456 Elm St", City: ""}, // tag dive will validate each item in the slice
			{Street: "", City: "Anothertown"},
		},
	}
	err := validate.Struct(account)
	if err != nil {
		t.Logf("Expected no validation error, but got: %v", err)
	}
}
