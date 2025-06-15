package validation

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestCustomValidation(t *testing.T) {
	// Create a new validator instance
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register a custom validation function
	validate.RegisterValidation("valid", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "valid"
	})

	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	validate.RegisterValidation("username", func(field validator.FieldLevel) bool {
		return usernameRegex.MatchString(field.Field().String())
	})

	type User struct {
		Username string `validate:"required,username"`
	}
	testUsers := []User{
		{"valid_user-123.name"}, // ✅
		{"Invalid*user!"},       // ❌
		{"another.valid-User"},  // ✅
		{"with space"},          // ❌
	}

	for _, user := range testUsers {
		err := validate.Struct(user)
		if err != nil {
			t.Logf("Username '%s' is INVALID ❌: %s\n", user.Username, err)
		} else {
			t.Logf("Username '%s' is valid ✅\n", user.Username)
		}
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")
func MusValidPin(field validator.FieldLevel) bool {
	length, _ := strconv.Atoi(field.Param())
	valud := field.Field().String()
	return regexNumber.MatchString(valud) && len(valud) == length
}

func TestCustomValidationWithParam(t *testing.T) {
	// Create a new validator instance
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register a custom validation function with a parameter
	validate.RegisterValidation("pin", MusValidPin)

	type User struct {
		Pin string `validate:"required,pin=5"`
	}

	testPins := []User{
		{"12345"},  // ✅
		{"1234"},   // ❌
	}
	for _, user := range testPins {
		err := validate.Struct(user)
		if err != nil {
			t.Logf("Pin '%s' is INVALID ❌: %s\n", user.Pin, err)
		} else {
			t.Logf("Pin '%s' is valid ✅\n", user.Pin)
		}
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2();
	if !ok {
		panic("Field not found")
	}
	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())
	return firstValue == secondValue;
}

func TestCrossFieldValidation(t *testing.T) {
	// Create a new validator instance
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register a custom validation function for cross-field validation
	validate.RegisterValidation("equalsIgnoreCase", MustEqualsIgnoreCase)

	type User struct {
		Password string `validate:"required"`
		ConfirmPassword string `validate:"required,equalsIgnoreCase=Password"`
	}

	testUsers := []User{
		{"password123", "PASSWORD123"}, // ✅
		{"password123", "wrongpassword"}, // ❌
	}

	for _, user := range testUsers {
		err := validate.Struct(user)
		if err != nil {
			t.Logf("User with Password '%s' and ConfirmPassword '%s' is INVALID ❌: %s\n", user.Password, user.ConfirmPassword, err)
		} else {
			t.Logf("User with Password '%s' and ConfirmPassword '%s' is valid ✅\n", user.Password, user.ConfirmPassword)
		}
	}
}


type RegisterRequest struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Phone	string `validate:"required,numeric"`
	Password string `validate:"required,min=8"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	request := level.Current().Interface().(RegisterRequest)
	if request.Username == request.Email || request.Username == request.Phone {
		// ok
	} else {
		level.ReportError(request.Username, "Username", "Username", "usernameValidation", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	// Create a new validator instance
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Register a custom struct-level validation function
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	testRequests := RegisterRequest{
		Username: "1234567890",
		Email:    "testuser@mail.com",
		Phone:    "1234567890",
		Password: "password123",
	}
	err := validate.Struct(testRequests)
	if err != nil {
		t.Fatalf("Expected no validation error, but got: %v", err)
	} else {
		t.Logf("RegisterRequest is valid ✅\n")
	}
}