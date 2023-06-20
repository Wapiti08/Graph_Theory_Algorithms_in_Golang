package main

import (
	"container/list"
	"fmt"
)

// implement snakes and ladder game

/*

shorted path with breadth first search

*/

type Pair struct {
	First  int
	Second int
}

type SocialGraph struct {
	V    int
	Nbrs []*list.List
}

func NewGraph(v int) *SocialGraph {
	g := &SocialGraph{
		V:    v,
		Nbrs: make([]*list.List, v+1),
	}

	// initialize the list for element in list
	for i := 1; i <= v; i++ {
		g.Nbrs[i] = list.New()
	}

	return g
}

func (g *SocialGraph) addedge(i, j int, undir bool) {
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

func (g *SocialGraph) min_dict_throws(snakes []Pair, ladders []Pair) {
	board := make(map[int]int, g.V+1)
	// define the position change for snake
	for _, snake := range snakes {
		s := snake.First
		e := snake.Second
		board[s] = e - s
	}
	// define the position change for ladder
	for _, ladder := range ladders {
		s := ladder.First
		e := ladder.Second
		board[s] = e - s
	}

	// build the graph based on snake and ladder positions
	for i := 1; i <= g.V; i++ {
		// the max value of mice is 6, min value is 1
		for mice := 1; mice <= 6; mice++ {
			// check the snake or ladder and whether dest
			v := i + mice
			v += board[v]
			if v <= g.V {
				g.addedge(i, v, false)
			}

		}
	}

}

func (g *SocialGraph) PrintEdgeList() {
	for i := 0; i < g.V; i++ {
		fmt.Printf("the neighbor of %d includes: \n", i)
		// check whether g.Nbrs[i] is null
		if g.Nbrs[i].Len() != 0 {
			for e := g.Nbrs[i].Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}
	}
}

func (g *SocialGraph) bfss(source int, dest int) {
	// build queue to save source
	q := list.New()
	// build list to save visted node with map <source, bool> format
	visited := make([]bool, g.V+1)

	q.PushBack(source)
	visited[source] = true

	dist := make([]int, g.V+1)
	dist[source] = 0

	// check whether the q is empty ---- have to use for instead of if
	for q.Len() > 0 {
		f := q.Front().Value.(int)
		q.Remove(q.Front())
		// need to give int type to list element
		for e := g.Nbrs[f].Front(); e != nil; e = e.Next() {
			nbr := e.Value.(int)
			if !visited[nbr] {
				q.PushBack(nbr)
				dist[nbr] = dist[f] + 1
				visited[nbr] = true
			}
		}

	}
	fmt.Println(dist[dest])

}

func main() {
	g := NewGraph(36)
	ladders := []Pair{{2, 15}, {5, 7}, {9, 27}, {18, 29}, {25, 35}}
	snakers := []Pair{{17, 4}, {20, 6}, {34, 12}, {24, 16}, {32, 30}}
	g.min_dict_throws(snakers, ladders)
	g.bfss(1, 36)

}
