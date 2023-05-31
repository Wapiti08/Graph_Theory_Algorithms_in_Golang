package main

import "fmt"

func minTrioDegree(n int, roads [][]int) int {

	deg := make([]int, n+1)
	edges := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		edges[i] = make([]int, n+1)
	}
	for i := 0; i < len(roads); i++ {
		x := roads[i][0]
		y := roads[i][1]

		deg[x]++
		deg[y]++
		edges[x][y]++
		edges[y][x]++
	}

	ans := 0

	for i := 1; i < n+1; i++ {
		for j := i + 1; j <= n; j++ {
			for k := j + 1; k <= n; k++ {
				// the minimum degree of a connected trio: minux 6 for edges in cycle
				if edges[i][j] != 0 && edges[j][k] != 0 && edges[i][k] != 0 {
					ans = min(deg[i]+deg[j]+deg[k]-6, ans)

				}

			}
		}
	}
	if ans == 0 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 6
	roads := [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}
	result := minTrioDegree(n, roads)
	fmt.Println(result)
}
