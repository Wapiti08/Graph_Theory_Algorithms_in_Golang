import "container/list"

type Graph struct {
	V    int
	Nbrs []*list.List
}

func newGraph(n int) *Graph {
	g := &Graph{
		V:    n,
		Nbrs: make([]*list.List, n),
	}
	for i := 0; i < n; i++ {
		g.Nbrs[i] = list.New()
	}

	return g
}

func addEdge(g *Graph, graph [][]int) *Graph {
	for row := range graph {
		for column := range graph[row] {
			g.Nbrs[row].PushBack(graph[row][column])
		}
	}
	return g
}

// func bfshelper(g *Graph, node int, visited []int, parent []int) {
func bfshelper(g *Graph, node int, visited []int) bool {
	visited[node] = 1

	for e := g.Nbrs[node].Front(); e != nil; e = e.Next() {
		nbr := e.Value.(int)
		if visited[nbr] != 0 {
			visited[nbr] = 2
			// parent[nbr] = node
			// bfshelper(g, nbr, visited, parent)
			bfshelper(g, nbr, visited)
		}
		// if visited[nbr] == visisted[node] && nbr !=parent[node] {
		if visited[nbr] == visited[node] {
			return false
		}

	}
	return true

}

// func dfshelper() {

// }

func isBipartite(graph [][]int) bool {

	g := newGraph(len(graph))
	g = addEdge(g, graph)
	// parent := make([]int, g.V)
	visited := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		// bfshelper(g, n, visited, parent)
		if bfshelper(g, i, visited) {
			return true
		}
	}
	return false
}