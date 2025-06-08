package test

import (
	"dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	service, err := simple.InitializedService(false)
	if err != nil {
		t.Fatalf("Failed to initialize service: %v", err)
	}

	if service == nil {
		t.Fatal("Expected service to be initialized, but got nil")
	}
}

func TestSimpleServiceError(t *testing.T) {
	service, err := simple.InitializedService(true)
	assert.NotNil(t, err, "Expected an error when initializing service with isError=true")
	assert.Nil(t, service, "Expected service to be nil when an error occurs")
}