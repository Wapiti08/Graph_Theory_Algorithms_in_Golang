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

func (g *Graph) addEdge(i, j int, undir bool) {
	g.Nbrs[i].PushBack(j)

	if undir {
		g.Nbrs[j].PushBack(i)
	}
}

func (g *Graph) dfs(node int, visited []bool, parent int) bool {
	visited[node] = true
	for e := g.Nbrs[node].Front(); e != nil; e = e.Next() {
		nbr := e.Value.(int)
		if !visited[nbr] {
			nbrFoundCycle := g.dfs(nbr, visited, node)
			if nbrFoundCycle {
				return true
			}
		}
		if nbr != parent {
			return true
		}
	}
	return false
}

func (g *Graph) contains_cycle() bool {
	visited := make([]bool, g.V)
	for i := 0; i < g.V; i++ {
		visited[i] = false
	}
	return g.dfs(0, visited, -1)
}

func main() {
	g := NewGraph(3)
	g.addEdge(0, 1, true)
	g.addEdge(1, 2, true)
	g.addEdge(2, 0, true)

	fmt.Println(g.contains_cycle())
}
