package main

/*
	Implementation of Maximal Network Rank
*/

import (
	// "container/list"
	"fmt"
)

// // define the graph instance
// type SocialGraph struct {
// 	// define the size
// 	V int
// 	// define the nbrs with list type
// 	nbrs      []*list.List
// 	edges     []Tuple
// 	Road_Rank map[Tuple]int
// }

// type Tuple struct {
// 	item1 int
// 	item2 int
// }

// // define the function to add egdge for undirected graph
// func (g *SocialGraph) AddEdges(i, j int) {
// 	if g.nbrs[i] == nil {
// 		g.nbrs[i] = list.New()
// 	}
// 	g.nbrs[i].PushBack(j)

// 	if g.nbrs[j] == nil {
// 		g.nbrs[j] = list.New()
// 	}
// 	g.nbrs[j].PushBack(i)

// 	g.edges = append(g.edges, Tuple{
// 		item1: i,
// 		item2: j,
// 	})

// }

// func (g *SocialGraph) PrintEdgeList() {
// 	for i := 0; i < g.V; i++ {
// 		fmt.Printf("the neighbor of %d includes: \n", i)
// 		// check whether Arr[i] is null
// 		if g.nbrs[i].Len() != 0 {
// 			for e := g.nbrs[i].Front(); e != nil; e = e.Next() {
// 				fmt.Println(e.Value)
// 			}
// 		}

// 	}
// }

// // count the edges of two nodes without repetitions
// func (g *SocialGraph) EdgeCount(i, j int) {
// 	Count := 0
// 	// get the nbrs of each node
// 	for e1 := g.nbrs[i].Front(); e1 != nil; e1.Next() {
// 		nbr1 := e1.Value.(int)
// 		if nbr1 != j {
// 			Count += 1
// 		}
// 	}

// 	fmt.Println(Count)
// 	// choose one side computation to add the overlap road
// 	for e2 := g.nbrs[j].Front(); e2 != nil; e2.Next() {
// 		Count += 1
// 	}
// 	fmt.Println(Count)

// 	if i < j {
// 		g.Road_Rank[Tuple{item1: i, item2: j}] = Count
// 	} else {
// 		g.Road_Rank[Tuple{item1: j, item2: i}] = Count
// 	}

// }

// func main() {

// 	g := SocialGraph{
// 		V:         4,
// 		nbrs:      make([]*list.List, 4),
// 		edges:     make([]Tuple, 0),
// 		Road_Rank: make(map[Tuple]int),
// 	}

// 	g.AddEdges(0, 1)
// 	g.AddEdges(0, 3)
// 	g.AddEdges(1, 2)
// 	g.AddEdges(1, 3)

// 	for _, Value := range g.edges {
// 		g.EdgeCount(Value.item1, Value.item2)

// 	}

// 	maxScore := 0

// 	for _, count := range g.Road_Rank {
// 		if count > maxScore {
// 			maxScore = count
// 		}
// 	}
// 	fmt.Println(g.Road_Rank)
// 	fmt.Println(g.edges)

// 	g.PrintEdgeList()

// 	fmt.Printf("The Maximal Network Rank is %d ", maxScore)

// }

// without graph to implement

// input roads
func MaxDegree(n int, roads [][]int) int {
	// create degree list and edges arr
	deg := make([]int, n)
	edges := make([][]int, n)

	for i := 0; i < n; i++ {
		// define rows
		edges[i] = make([]int, n)
	}

	for i := 0; i < len(roads); i++ {
		// get the seperate elements in roads
		x := roads[i][0]
		y := roads[i][1]

		deg[x]++
		deg[y]++

		edges[x][y]++
		edges[y][x]++
	}
	// temp value
	ans := 0
	// traverse with two loops to find the largest
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// minus the repeated edge
			ans = max(deg[i]+deg[j]-edges[i][j], ans)
		}
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 4
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}
	result := maximalNetworkRank(n, roads)
	fmt.Println(result)
}
