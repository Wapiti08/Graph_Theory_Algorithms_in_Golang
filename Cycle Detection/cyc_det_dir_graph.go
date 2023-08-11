package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	V    int
	Nbrs []*list.List
}

func NewGraph(v int) *Graph {

	g := &Graph{
		V:    v,
		Nbrs: make([]*list.List, v),
	}

	for i := 0; i < v; i++ {
		g.Nbrs[i] = list.New()
	}

	return g

}

func (g *Graph) addEdge(i, j int) {
	g.Nbrs[i].PushBack(j)
}

func (g *Graph) dfs(node int, visited []bool, stacked []bool) bool {
	visited[node] = true
	// make the current visisted node
	stacked[node] = true

	for e := g.Nbrs[node].Front(); e != nil; e = e.Next() {
		nbr := e.Value.(int)
		if stacked[nbr] {
			fmt.Println("Found a cycle for current processing node")
			return true
		}
		if !visited[nbr] {
			nbrFoundCycle := g.dfs(nbr, visited, stacked)
			if nbrFoundCycle {
				return true
			}
		}
	}

	stacked[node] = false
	return false
}

func (g *Graph) contains_cycle() bool {
	visited := make([]bool, g.V)
	stacked := make([]bool, g.V)
	for i := 0; i < g.V; i++ {
		visited[i] = false
		stacked[i] = false
	}

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			if g.dfs(i, visited, stacked) {
				return true
			}
		}
	}
	return false

}

func main() {
	g := NewGraph(3)
	g.addEdge(0, 1)
	g.addEdge(1, 2)
	// g.addEdge(2, 0)

	fmt.Println(g.contains_cycle())
}
