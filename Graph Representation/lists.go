package main

import (
	"container/list"
	"fmt"
)

func main() {
	// create a new List

	var intList list.List
	intList.PushBack(11)
	intList.PushBack(12)
	intList.PushBack(13)

	// Front is the first value
	for e := intList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
