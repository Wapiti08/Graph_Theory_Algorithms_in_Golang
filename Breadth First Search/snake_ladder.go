package main

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
	for i:=1; i<=v*v; i++ {
		g.Nbrs[i] = list.New()
	}

	return g
}

func addedge(N int, snake []Pair, ladder []Pair) {
	for i:=1; i<=v*v;i++{
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


func min_dict_throws(N int, snake []Pair, ladder []Pair) int {
	// define the positions of snakers and ladders
	arr := make([][]int, N)
	// assign the value
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < N; j++ {
			arr[i][j] -= N * N
		}
	}

	throw := 0
	source = 1
	// traverse the nbr
	for i := source; i <= source+6; i++ {
		// check the nbr of i
		for j := i+1; j<i+6; j++ {
			
			// check the snake or ladder and whether dest
			resp_sna, res_sna := IsElementExist(j, snake) 
			if resp_sna {
				j = res_sna
			} else {
				resp_lad, res_lad := IsElementExist(j, ladder)
				if resp_lad {
					j = res_lad
				}
			}
			
			}

		}
		
		// calculate the total throw
		

		// if snake --- next node is snake[1]，break current loop and restart from snake[1]

		// if ladder --- next node is ladder[1], break current loop and restart from ladder[1]

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