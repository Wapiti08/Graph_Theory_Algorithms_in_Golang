// import "container/list"

// type Graph struct {
//     V int
//     Nbrs []*list.List
// }

// func newGraph(n int) *Graph {
//     g := &Graph{
//         V: n,
//         Nbrs: make([]*list.List, n),
//     }
//     for i:=0; i<n; i++ {
//         g.Nbrs[i] = list.New()
//     }

//     return g
// }

// func addEdge(g *Graph, graph [][]int) *Graph {
//     for row := range graph {
//         for column := range graph[row] {
//             g.Nbrs[row].PushBack(graph[row][column])
//         }
//     }
//     return g
// }

// func dfshelper(g *Graph, node int, visited []int, parent int, color int) bool {
//     visited[node] = color
//     if g.Nbrs[node].Len() != 0 {
//         for e := g.Nbrs[node].Front(); e!=nil; e=e.Next() {
//             parent = node
//             nbr := e.Value.(int)
//             if visited[nbr]==0 {
//                 subprob := dfshelper(g, nbr, visited, parent, 3-color)
//                 if !subprob {
//                     return false
//                 }   
//             } else if nbr != parent && color == visited[nbr] {
//                 return false
//             } 
//         }
//         return true
//     } else {
//         node += 1
//         if visited[node]==0 {
//                 subprob := dfshelper(g, node, visited, -1, 3-color)
//                 if !subprob {
//                     return false
//                 }   
//         }
//         return true
//     }


// }


// func bfshelper() {

// }


func isBipartite(graph [][]int) bool {
    n := len(graph)
    colors := make([]int, n)
    q := make([]int, n)

    for i:=0 ; i<n; i++ {
        if colors[i] != 0 {
            continue
        }
        colors[i]=1
        q = append(q, i)
        // implement bfs
        for len(q)>0 {
            temp := q[0]
            // pop out the first element
            q = q[1:]
            for _, nbr := range graph[temp] {
                if colors[nbr] == 0 {
                    colors[nbr]=-colors[temp]
                    q = append(q, nbr)
                } else if colors[nbr] == colors[temp] {
                    return false
                }
            }
        }
    }
    return true

}