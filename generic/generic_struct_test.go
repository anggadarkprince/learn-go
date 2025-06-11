package generic

import "testing"

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello" + name
}
func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	dataInt := Data[int]{First: 1, Second: 2}
	if dataInt.First != 1 || dataInt.Second != 2 {
		t.Errorf("Expected Data[int] First=1, Second=2, got First=%d, Second=%d", dataInt.First, dataInt.Second)
	}

	dataString := Data[string]{First: "Hello", Second: "World"}
	if dataString.First != "Hello" || dataString.Second != "World" {
		t.Errorf("Expected Data[string] First='Hello', Second='World', got First='%s', Second='%s'", dataString.First, dataString.Second)
	}
}

func TestDataMethod(t *testing.T) {
	dataInt := Data[int]{First: 1, Second: 2}
	result := dataInt.SayHello(" World")
	if result != "Hello World" {
		t.Errorf("Expected SayHello to return 'Hello World', got '%s'", result)
	}

	dataInt.ChangeFirst(10)
	if dataInt.First != 10 {
		t.Errorf("Expected ChangeFirst to change First to 10, got %d", dataInt.First)
	}
}