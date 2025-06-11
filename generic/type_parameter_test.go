package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println("Length called with parameter:", param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Hello, World!")
	assert.Equal(t, "Hello, World!", result)
	var intResult int = Length[int](42)
	assert.Equal(t, 42, intResult)
	var floatResult float64 = Length[float64](3.14)
	assert.Equal(t, 3.14, floatResult)
	var boolResult bool = Length[bool](true)
	assert.Equal(t, true, boolResult)
}

func MultipleParameters[T any, U any](param1 T, param2 U) (T, U) {
	fmt.Println("MultipleParameters called with parameters:", param1, param2)
	return param1, param2
}
func TestMultipleParameters(t *testing.T) {
	var strResult, intResult = MultipleParameters[string, int]("Hello", 42)
	assert.Equal(t, "Hello", strResult)
	assert.Equal(t, 42, intResult)

	var floatResult, boolResult = MultipleParameters[float64, bool](3.14, true)
	assert.Equal(t, 3.14, floatResult)
	assert.Equal(t, true, boolResult)
}

func TestInference(t *testing.T) {
	var strResult = Length("Hello, World!") // type inference (no need add type parameter Length[string])
	assert.Equal(t, "Hello, World!", strResult)

	var intResult = Length(42) // type inference
	assert.Equal(t, 42, intResult)

	var floatResult = Length(3.14) // type inference
	assert.Equal(t, 3.14, floatResult)

	var boolResult = Length(true) // type inference
	assert.Equal(t, true, boolResult)
}