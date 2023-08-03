package main

import (
	"fmt"
)

func dfs(x int, v [][]int, visit []int) {
	// choose one starting node and set it as visited
	visit[x] = 1
	// check its adjacent neighbors
	for _, itr := range v[x] {
		// if not visited
		if visit[itr] == 0 {
			// choose one as the starting node again
			dfs(itr, v, visit)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visit := make([]int, n)
	// 0 must be the root node to start
	dfs(0, rooms, visit)
	// find all visited rooms 0-1-2-..n-1
	for i := 0; i < n; i++ {
		if visit[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	// there is no need to visited in room order, but have to make sure all rooms can be visisted
	rooms := [][]int{{2}, {}, {1}}
	fmt.Println(canVisitAllRooms(rooms))
}
