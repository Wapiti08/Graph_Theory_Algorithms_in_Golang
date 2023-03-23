package main

import (
	"fmt"
)

// Method 1 -- implemented with list package

// type Graph struct {
// 	// define the number of node --- Size
// 	V int
// 	// define the edge
// 	List *[][]list
// }

// func AddEdge(g Graph i,j int, undir bool) {
// l = new g.

// }

// Method 2 --- Pure implementation with array

type SocialGraph struct {
	// define the size
	V int
	// define the type of array of int list, row is the list order, column is the list of neighbors
	// whether to create array from make or just define?
	Arr [][]int
}

func AddEdge(g SocialGraph, i, j int, undir bool) {
	/*
		param g: instance of SocialGraph
		param i: edge start
		param j: edge end
		param undir: whether directed
	*/

	g.Arr = append(g.Arr[i][], j)
	if undir {
		g.Arr = append(g.Arr[j][], i)
	}

	fmt.Println(g.Arr)

}

func PrintEdgeList(g SocialGraph) {
	for i := range g.V {
		fmt.Printf("the neighbor of %d includes %d",i, g.Arr[i])
	}
}

func main() {
	graph := SocialGraph()

	AddEdge(graph, 1, 3, true)
	AddEdge(graph, 2, 3, true)
	AddEdge(graph, 2, 4, true)
	AddEdge(graph, 1, 5, true)
	AddEdge(graph, 1, 6, true)
}
