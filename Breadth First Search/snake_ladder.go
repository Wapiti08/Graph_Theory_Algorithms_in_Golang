package main

import "testing/quick"

// implement snakes and ladder game

/*

shorted path with breadth first search

*/

type Pair struct {
	First  int
	Second int
}

type SocialGraph struct {
	V int
	Nbrs []*list.List
}

func NewGraph(v int) *SocialGraph {
	g := &SocialGraph{
		V: v,
		Nbrs: make([]*list.List, v),
	}

	// initialize the list for element in list
	for i:=1; i<=v; i++ {
		g.Nbrs[i] = list.New()
	}

	return g
}

func (g *SocialGraph) addedge(N int, snake []Pair, ladder []Pair) {
	// build the graph based on snake and ladder positions
	for i:=1; i<=N-6;i++{
		// the max value of mice is 6, min value is 1
		for j:=i+1; j<=i+7;j++ {
			// check the snake or ladder and whether dest
			resp_sna, res_sna := IsElementExist(j, snake) 
			if resp_sna {
				j = res_sna
			} else {
				resp_lad, res_lad := IsElementExist(j, ladder)
				if resp_lad {
					j = res_lad
				}
			g.Nbrs[i].PushBack(j)
		}
	}
}
func (g *SocialGraph) printedge(N int, snake []Pair, ladder []Pair) {
	for i := 0; i < g.V; i++ {
		fmt.Printf("the neighbor of %d includes: \n", i)
		// check whether Arr[i] is null
		if g.Nbrs[i].Len() != 0 {
			for e := g.Nbrs[i].Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}

	}

func (g *SocialGraph) min_dict_throws(source, dest int) {
	// find the shorted path using breadth first search	
	q := list.New()
	visited := make([]bool, g.V)

	q.PushBack(source)
	visited[source] = true
	// define the dist from the source
	dist := make([]int, g.V)
	parent := make([]int, g.V)

	for i:=0; i <g.V; i++ {
		parent[i] = -1
	}

	dist[source] = 0
	// initialize the first element
	parent[source] = source

	for q.Len() > 0 {
		f := q.Front().Value.(int)
		q.Remove(q.Front())
		// traverse the neighbors of f
		for e:= g.Nbrs[f].Front(); e!=nil; e=e.Next {
			nbr = e.Value.(int)
			if !visited[nbr] {
				q.PushBack(nbr)
				parent[nbr] = f
				dist[nbr] = dist[f] + 1
				visited[nbr] = true
			}
		}
	}

	// print out the shorted path through parent for given node
	if dest != -1 {
		temp := dest
		if temp != source {
			fmt.Println("%d -- ",temp)
			temp = parent[temp]
		}
		fmt.Print("%d", source)
	}


	}
}

func IsElementExist(ele int, tup []Pair) (bool, int){
	for element := range tup {
		if ele == element.First {
			return true, element.Second
		}

		return false, ele
	}

}

func main() {
	g := NewNewGraph(36)
	ladders := make([][]Pair){{2,15},{5,7},{9,27},{18,29},{25,35}}
	snakers := make([][]Pair){{17,4},{20,6},{34,12},{24,16},{32,30}}
	g.addedge(36, snakes, ladders)
	g.printedge()
	g.min_dict_throws(1, 36)

}