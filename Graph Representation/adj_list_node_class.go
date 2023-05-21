package main

/*
 * @Author: Newt Tan
 * @Date: 2023-04-15 16:04:02
 * @Last Modified by:   Newt Tan
 * @Last Modified time: 2023-04-15 16:04:02
 */

import (
	"fmt"
)

type Node struct {
	// define the node name
	name string
	// define the node neighbors
	nbrs []string
}

func NewNode(name string) *Node {
	return &Node{
		name: name,
	}
}

type Graph struct {
	// define all nodes and hashmap (string, *node)
	m map[string]*Node
}

func NewGraph(cities []string) *Graph {
	// initialize a Graph instance
	g := &Graph{
		m: make(map[string]*Node),
	}

	// add cities inside graph
	for _, city := range cities {
		g.m[city] = NewNode(city)
	}

	return g
}

// add edges inside graph
func (g *Graph) addEdge(x, y string, undir bool) {
	g.m[x].nbrs = append(g.m[x].nbrs, y)

	if undir {
		g.m[y].nbrs = append(g.m[y].nbrs, x)
	}
}

// print neighbors of a city
func (g *Graph) printAdjList() {
	for city, node := range g.m {
		fmt.Printf("%s -->", city)
		for _, nbr := range node.nbrs {
			fmt.Printf("%s", nbr)
		}
		fmt.Println()
	}
}

func main() {
	cities := []string{"Delhi", "London", "Paris", "New York"}
	g := NewGraph(cities)

	g.addEdge("Delhi", "London", false)
	g.addEdge("New York", "London", false)
	g.addEdge("Delhi", "Paris", false)
	g.addEdge("Paris", "New York", false)

	g.printAdjList()
}
