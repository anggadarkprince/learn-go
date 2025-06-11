package generic

import "testing"

// type generic
type GetterSetter[T any] interface {
	Get() T
	Set(value T)
}

// function generic
func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.Set(value)
	return param.Get()
}

// implementation of the generic interface
type MyData[T any] struct {
	Value T
}
func (d *MyData[T]) Get() T {
	return d.Value
}
func (d *MyData[T]) Set(value T) {
	d.Value = value
}

func TestInterface(t *testing.T) {
	data := &MyData[int]{Value: 10}
	newValue := ChangeValue(data, 20)
	if newValue != 20 || data.Get() != 20 {
		t.Errorf("Expected value to be changed to 20, got %d", newValue)
	}

	dataStr := &MyData[string]{Value: "Hello"}
	newStrValue := ChangeValue(dataStr, "World")
	if newStrValue != "World" || dataStr.Get() != "World" {
		t.Errorf("Expected value to be changed to 'World', got '%s'", newStrValue)
	}
}