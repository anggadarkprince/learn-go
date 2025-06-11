package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Inline function to test type set
func FindMin[T interface {int | int64 | float32 | float64}](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func TestInlineTypeSet(t *testing.T) {
	assert.Equal(t, 1, FindMin(1, 2))
	assert.Equal(t, int64(1), FindMin[int64](1, 2))
	assert.Equal(t, float32(1.0), FindMin[float32](1.0, 2.0))
	assert.Equal(t, float64(1.0), FindMin(1.0, 2.0))
}

func GetFirst[T []E, E any](data T) E {
	if len(data) == 0 {
		var zero E
		return zero
	}
	return data[0]
}

func TestGetFirst(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, 1, GetFirst(intSlice))

	stringSlice := []string{"a", "b", "c"}
	assert.Equal(t, "a", GetFirst(stringSlice))

	floatSlice := []float64{1.1, 2.2, 3.3}
	assert.Equal(t, 1.1, GetFirst(floatSlice))

	emptySlice := []int{}
	assert.Equal(t, 0, GetFirst(emptySlice)) // zero value for int
}