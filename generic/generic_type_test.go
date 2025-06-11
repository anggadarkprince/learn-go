package generic

import "testing"

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, item := range bag {
		println(item)
	}
}

func TestBag(t *testing.T) {
	bag := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(bag)

	bagStr := Bag[string]{"apple", "banana", "cherry"}
	PrintBag(bagStr)

	bagFloat := Bag[float64]{1.1, 2.2, 3.3}
	PrintBag(bagFloat)
}