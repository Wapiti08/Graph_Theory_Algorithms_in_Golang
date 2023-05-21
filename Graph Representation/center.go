package main

import (
	"container/list"
	"fmt"
)

// import the center of a star graph --- neighbors

type SocialGraph struct {
	// define the size
	V int
	// define the type of array of int list, row is the list order, column is the list of neighbors
	// list.List if pointer type, *list.List is the type of List
	Arr map[int]*list.List
}

func AddEdge(g SocialGraph, i, j int, undir bool) {
	/*
		param g: instance of SocialGraph
		param i: edge start
		param j: edge end
		param undir: whether directed
	*/
	if g.Arr[i] == nil {
		g.Arr[i] = list.New()
	}
	g.Arr[i].PushBack(j)
	if undir {
		if g.Arr[j] == nil {
			g.Arr[j] = list.New()
		}
		g.Arr[j].PushBack(i)
	}

}

func PrintEdgeList(g SocialGraph) {
	for i := 1; i <= g.V; i++ {
		fmt.Printf("the neighbor of %d includes: \n", i)
		// check whether Arr[i] is null
		if g.Arr[i].Len() != 0 {
			for e := g.Arr[i].Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}

	}
}

func findCenter(adj SocialGraph) {
	var center int
	// return the length of edge list
	n := len(adj.Arr)
	for i := 1; i < n-1; i++ {
		if adj.Arr[i].Len() < adj.Arr[i+1].Len() {
			// node number is lager than order number for 1
			center = i + 1
		} else {
			center = i
		}

	}
	fmt.Println("The center of star grpah is: ", center)
}

func main() {
	var n int
	// receive the input
	fmt.Println("Please input the largest node number: ")
	// in case read the \n as a parameter
	fmt.Scanf("%d\n", &n)

	// define the 2d array --- edges
	var arr [][]int
	// 读取输入的二维数组, edge number is n-1, equals to row number
	for i := 0; i < n-1; i++ {
		// define row
		var row []int
		for j := 0; j < 2; j++ {
			fmt.Printf("Please input the element of %d row and %d column: \n", i+1, j+1)
			var num int
			fmt.Scanf("%d\n", &num)
			row = append(row, num)
		}
		arr = append(arr, row)

	}

	fmt.Println("The 2D array is: ", arr)
	// build the graph instance
	graph := SocialGraph{V: n, Arr: make(map[int]*list.List)}
	// add edges n-1 times
	for order := 0; order < n-1; order++ {
		AddEdge(graph, arr[order][0], arr[order][1], true)

	}

	PrintEdgeList(graph)

	// find the center of node with max neighbor
	findCenter(graph)
}
