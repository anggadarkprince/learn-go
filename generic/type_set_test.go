package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
} // adding ~ allows for type approximation, so Age can be used as well

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, 1, Min[int](1, 2))
	assert.Equal(t, int8(1), Min[int8](1, 2))
	assert.Equal(t, int16(1), Min[int16](1, 2))
	assert.Equal(t, int32(1), Min[int32](1, 2))
	assert.Equal(t, int64(1), Min[int64](1, 2))
	assert.Equal(t, float32(1.0), Min[float32](1.0, 2.0))
	assert.Equal(t, float64(1.0), Min[float64](1.0, 2.0))
}

func TestTypeApproximation(t *testing.T) {
	assert.Equal(t, Age(1), Min[Age](1, 2)) // type approximation for Age
}
