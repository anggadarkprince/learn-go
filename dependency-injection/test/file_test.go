package test

import (
	"dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	// Initialize the connection
	conn, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, conn)

	cleanup()
}