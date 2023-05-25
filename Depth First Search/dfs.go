package main

/*
 * @Author: Newt Tan
 * @Date: 2023-04-15 16:04:02
 * @Last Modified by:   Newt Tan
 * @Last Modified time: 2023-04-15 16:04:02
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

// achieve DFS with recursive traverse
func (g *SocialGraph) dfsHelper(node int, visited []bool) {
	// define visited list
	visited[node] = true
	fmt.Println(node)
	for e := g.Nbrs[node].Front(); e != nil; e = e.Next() {
		// for _, e := range g.Nbrs[node] {
		nbr := e.Value.(int)
		if visited[nbr] != true {
			g.dfsHelper(nbr, visited)
		}
	}

}

func (g *SocialGraph) dfs(source int) {
	visited := make([]bool, g.V)
	g.dfsHelper(source, visited)
}

func main() {
	g := NewGraph(7)
	// g.addEdge(0, 1, true)
	// g.addEdge(1, 2, true)
	// g.addEdge(2, 3, true)
	// g.addEdge(3, 5, true)
	// g.addEdge(5, 6, true)
	// g.addEdge(4, 5, true)
	// g.addEdge(0, 4, true)
	// g.addEdge(3, 4, true)
	g.addEdge(0, 1, true)
	g.addEdge(0, 2, true)
	g.addEdge(1, 3, true)
	g.addEdge(1, 4, true)
	g.addEdge(2, 5, true)
	g.addEdge(2, 6, true)
	// g.PrintEdgeList()
	g.dfs(0)
}
