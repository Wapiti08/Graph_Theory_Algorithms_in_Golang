package main

import (
	"fmt"
)

type Pair struct {
	First  int
	Second int
}

type SocialGraph struct {
	V int
	l []Pair
}

func (g *SocialGraph) addEdges(u, v int) {
	g.l = append(g.l, Pair{u, v})
}

func findSet(i int, parent map[int]int) int {
	if parent[i] == -1 {
		return i
	}
	return findSet(parent[i], parent)
}

func UnionSet(i, j int, parent map[int]int) bool {
	s1 := findSet(i, parent)
	s2 := findSet(j, parent)
	if s1 != s2 {
		parent[s1] = s2
		return false
	}
	return true
}

func (g *SocialGraph) contain_cycle(V int) bool {
	parent := make(map[int]int, V)
	for i := 0; i < V; i++ {
		// initialize the parent
		parent[i] = -1
	}

	for _, ed := range g.l {
		i := ed.First
		j := ed.Second

		s1 := findSet(i, parent)
		s2 := findSet(j, parent)

		if s1 != s2 {
			UnionSet(s1, s2, parent)
		} else {
			return true
		}
	}
	return false
}

func main() {
	var g SocialGraph
	g.addEdges(0, 1)
	g.addEdges(1, 2)
	g.addEdges(2, 3)
	g.addEdges(3, 0)
	fmt.Println(g.contain_cycle(4))
}
