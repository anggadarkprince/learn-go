package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	data.Value = "Val 1"
	data = data.Next()
	data.Value = "Val 2"
	data = data.Next()
	data.Value = "Val 3"

	data.Do(func(value any) {
		fmt.Println(value)
	})

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}