package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	// Create a new validator instance
	var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

	if validate == nil {
		t.Fatal("Validator instance is nil")
	}
}

func TestValidationField(t *testing.T) {
	name := ""

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Var(name, "required")

	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	} else {
		t.Fatal("Expected validation error, but got none")
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	password := "password123"
	confirmPassword := "password1234"
	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	} else {
		t.Fatal("Expected validation error, but got none")
	}
}

func TestMultipleTagValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	name := "John Doe"
	err := validate.Var(name, "required,alpha")
	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	} else {
		t.Fatal("Expected validation error, but got none")
	}
}

func TestBakedInvalidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	// https://pkg.go.dev/github.com/go-playground/validator/v10#readme-baked-in-validations
	number := 10
	email := "invalid-email"
	url := "htp://invalid-url"
	date := "2023-10-01"
	errNumber := validate.Var(number, "required ,numeric")
	errNumberGt := validate.Var(number, "required ,numeric,gt=100")
	errEmail := validate.Var(email, "required,email")
	errURL := validate.Var(url, "required,url")
	errDate := validate.Var(date, "required,date")

	if errNumber != nil {
		t.Logf("Validation error for number: %v", errNumber.Error())
	} else {
		t.Fatal("Expected validation error for number, but got none")
	}
	if errNumberGt != nil {
		t.Logf("Validation error for number greater than 100: %v", errNumberGt.Error())
	} else {
		t.Fatal("Expected validation error for number greater than 100, but got none")
	}
	if errEmail != nil {
		t.Logf("Validation error for email: %v", errEmail.Error())
	} else {
		t.Fatal("Expected validation error for email, but got none")
	}
	if errURL != nil {
		t.Logf("Validation error for URL: %v", errURL.Error())
	} else {
		t.Fatal("Expected validation error for URL, but got none")
	}
	if errDate != nil {
		t.Logf("Validation error for date: %v", errDate.Error())
	} else {
		t.Fatal("Expected validation error for date, but got none")
	}
}

func TestTagParameterValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	number := 10
	err := validate.Var(number, "required,numeric,gt=5,lt=20")
	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	}

	number = 25
	err = validate.Var(number, "required,numeric,gt=5,lt=20")
	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	}
}

func TestOrRuleValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	username := "085655444444" // or email
	err := validate.Var(username, "required,email|numeric")
	if err != nil {
		t.Logf("Validation error: %v", err.Error())
	}
}