package main

import (
	// "container/list"
	"fmt"
)

// type SocialGraph struct {
// 	// define the size
// 	V int
// 	// define the type of array of int list, row is the list order, column is the list of neighbors
// 	// whether to create array from make or just define?
// 	Arr map[int]*list.List
// }

// func AddEdge(g SocialGraph, i, j int, undir bool) {
// 	/*
// 		param g: instance of SocialGraph
// 		param i: edge start
// 		param j: edge end
// 		param undir: whether directed
// 	*/

// 	g.Arr[i].PushBack(j)
// 	if undir {
// 		g.Arr[j].PushBack(i)
// 	}

// }

// func PrintEdgeList(g SocialGraph) {
// 	for i := 0; i <= g.V; i++ {
// 		fmt.Printf("the neighbor of %d includes %d \n", i, g.Arr[i])
// 	}
// }

// func main() {
// 	graph := SocialGraph{V: 6, Arr: make(map[int]*list.List)}
// 	AddEdge(graph, 1, 3, true)
// 	AddEdge(graph, 2, 3, true)
// 	AddEdge(graph, 2, 4, true)
// 	AddEdge(graph, 1, 5, true)
// 	AddEdge(graph, 1, 6, true)

// 	PrintEdgeList(graph)
// }

type SocialGraph struct {
	// define the size
	V int
	// define the type of array of int list, row is the list order, column is the list of neighbors
	// whether to create array from make or just define?
	Arr map[int][]int
}

func AddEdge(g SocialGraph, i, j int, undir bool) {
	/*
		param g: instance of SocialGraph
		param i: edge start
		param j: edge end
		param undir: whether directed
	*/

	g.Arr[i] = append(g.Arr[i], j)
	if undir {
		g.Arr[j] = append(g.Arr[j], i)
	}

}

func PrintEdgeList(g SocialGraph) {
	for i := 0; i <= g.V; i++ {
		fmt.Printf("the neighbor of %d includes %d \n", i, g.Arr[i])
	}
}

func main() {
	graph := SocialGraph{V: 6, Arr: make(map[int][]int, 6)}
	AddEdge(graph, 1, 3, true)
	AddEdge(graph, 2, 3, true)
	AddEdge(graph, 2, 4, true)
	AddEdge(graph, 1, 5, true)
	AddEdge(graph, 1, 6, true)

	PrintEdgeList(graph)
}
