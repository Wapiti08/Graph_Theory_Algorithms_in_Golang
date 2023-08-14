import "containers/list"

type Graph struct {
    V: n
    Nbrs: []*list.List
}

func newGraph(n int) *Graph {
    g := &Graph{
        V: n,
        Nbrs: make([]*list.List,n),
    }
    for i:=0; i<n; i++ {
        g.Nbrs[i] = list.New()
    }

    return g
}

func addEdge(g *Graph, graph [][]int) {
    
}

func bfshelper(g *Graph, node int, visited []int, parent []int) {
    visited[node] = 1

    for e := g.Nbrs[node].Front(); e!=nil; e=e.Next() {
        nbr := e.Value.(int)
        if visited[nbr]!=0 {
            visited[nbr] = 2
            parent[nbr] = node
            bfshelper(g, nbr, visited, parent)   
        } 
        if visited[nbr] == visisted[node] && nbr !=parent[node] {
            return false
        }    
    
    }
    return true

}


// func dfshelper() {

// }


func isBipartite(graph [][]int) bool {
    parent := make([]int, g.V)
    visited := make([]int, g.V)
    g := NewGraph(len(graph))


}