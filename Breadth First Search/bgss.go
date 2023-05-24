package main

/*
	shortest path undirected graph
*/

import (
	"container/list"
	"fmt"
)

type SocialGraph struct {
	// the size of Graph
	V int
	// define the type of array of int list, row is the list order, column is the list of neighbors
	// list.List if pointer type, *list.List is the type of List
	Nbrs []*list.List
}

func NewGraph(v int) *SocialGraph {
	// initialize a Graph instance
	g := &SocialGraph{
		V:    v,
		Nbrs: make([]*list.List, v),
	}
	// initialize the list for element in list
	for i := 0; i < v; i++ {
		g.Nbrs[i] = list.New()
	}

	return g
}

func (g *SocialGraph) addEdge(i, j int, undir bool) {
	/*
		param g: instance of SocialGraph
		param i: edge start
		param j: edge end
		param undir: whether undirected
	*/
	// if g.Nbrs[i] == nil {
	// g.Nbrs[i] = list.New()
	// }
	g.Nbrs[i].PushBack(j)

	if undir {
		// if g.Nbrs[j] == nil {
		// g.Nbrs[j] = list.New()
		// }
		g.Nbrs[j].PushBack(i)
	}

}

func (g *SocialGraph) PrintEdgeList() {
	for i := 0; i < g.V; i++ {
		fmt.Printf("the neighbor of %d includes: \n", i)
		// check whether Arr[i] is null
		if g.Nbrs[i].Len() != 0 {
			for e := g.Nbrs[i].Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}

	}
}

func (g *SocialGraph) bfs(source int, dest int) {
	// func (g *SocialGraph) bfs(source int) {
	// build queue to save source
	q := list.New()
	// build list to save visted node with map <source, bool> format
	visited := make([]bool, g.V)

	q.PushBack(source)
	visited[source] = true

	dist := make([]int, g.V)
	// initialize the distance map
	// for back trace
	parent := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		parent[i] = -1
	}

	dist[source] = 0
	parent[source] = source
	// print out the element from queue
	// check whether the q is empty ---- have to use for instead of if
	for q.Len() > 0 {
		f := q.Front().Value.(int)
		fmt.Printf("\n")
		fmt.Println(f)

		q.Remove(q.Front())
		// need to give int type to list element
		for e := g.Nbrs[f].Front(); e != nil; e = e.Next() {
			nbr := e.Value.(int)
			if !visited[nbr] {
				q.PushBack(nbr)
				dist[nbr] = dist[f] + 1
				parent[nbr] = f
				visited[nbr] = true
			}
		}

	}
	// print out the distance
	for i := 0; i < g.V; i++ {
		fmt.Printf("the shorted distance of %d from source %d is %d \n", i, source, dist[i])
	}

	// print out the shorted distance for any source
	// print parent one by one
	if dest != -1 {
		temp := dest

		for temp != source {
			fmt.Printf("%d -- ", temp)
			temp = parent[temp]
		}
		fmt.Printf("%d", source)
	}

}

func main() {
	g := NewGraph(7)
	g.addEdge(0, 1, true)
	g.addEdge(1, 2, true)
	g.addEdge(2, 3, true)
	g.addEdge(3, 5, true)
	g.addEdge(5, 6, true)
	g.addEdge(4, 5, true)
	g.addEdge(0, 4, true)
	g.addEdge(3, 4, true)
	// g.PrintEdgeList()
	g.bfs(1, 6)
}
