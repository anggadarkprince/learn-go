package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Angga")
	data.PushBack("Ari")
	data.PushBack("Wijaya")
	data.PushFront("Mr.")
	
	var head *list.Element = data.Front()
	//head.Next().Next().Next()
	fmt.Println(head.Value)
	
	next := head.Next()
	fmt.Println(next.Value)
	
	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}