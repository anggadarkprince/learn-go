package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](a, b T) bool {
	return a == b
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[int](1, 1))
	assert.Equal(t, false, IsSame[int](1, 2))
	assert.Equal(t, true, IsSame[string]("hello", "hello"))
	assert.Equal(t, false, IsSame[string]("hello", "world"))
	assert.Equal(t, true, IsSame[float64](3.14, 3.14))
	assert.Equal(t, false, IsSame[float64](3.14, 2.71))
}