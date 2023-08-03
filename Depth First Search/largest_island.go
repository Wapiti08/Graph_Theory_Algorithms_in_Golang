// type Node struct {
//     First int
//     Second int
// }

// func findnbr(ininode Node, matrix [][]int, visited [][]int, island map[Node]int) {
//     /*
//     initnode: the value is 1

//     */
//     x := ininode.First
//     y := ininode.Second
//     visited[x][y] = 1
//     island[ininode]++
//     // check three nbrs
//     if x-1>=0 && matrix[x-1][y] == 1 &&visited[x-1][y] ==0 {
//         visited[x-1][y] = 1
//         findnbr(Node{First: x-1, Second: y}, matrix, visited, island)
//     }
//     // check four nbrs
//     if y-1>=0 && matrix[x][y-1] == 1 && visited[x][y-1] == 0{
//         visited[x][y-1] = 1
//             findnbr(Node{First: x, Second: y-1}, matrix, visited, island)
//     }
//     if y+1 < len(matrix[0]) && matrix[x][y+1] == 1 && visited[x][y+1]==0{
//         visited[x][y+1] = 1
//             findnbr(Node{First: x, Second: y+1}, matrix, visited, island)
//     }
//     if x+1 < len(matrix) && matrix[x+1][y] == 1 && visited[x+1][y] == 0{
//         visited[x+1][y] = 1
//             findnbr(Node{First: x+1, Second: y}, matrix, visited, island)
//     }
// }

// func maxAreaOfIsland(grid [][]int) int {
//     visited := make([][]int, len(grid))
//     for i:= range visited{
//         visited[i] = make([]int, len(grid[0]))
//     }
//     island := make(map[Node]int)
//     for row := range grid {
//         for column := range grid[row]{
//             if grid[row][column] == 1 {
//                 node := Node{First:row, Second:column}
//                 island[node] = 0
//                 findnbr(Node{First: row, Second: column}, grid, visited, island)
//             }
//         }
//     }
//     var maxVal int
//     for _, num := range island {
//         if num > maxVal {
//             maxVal = num
//         }
//     }

//     return maxVal
// }

func findnbr(row int, column int, area int, grid [][]int) int {
	if grid[row][column] != 1 {
		return area
	}
	// visited
	grid[row][column] = 2
	area++
	if row > 0 {
		area = findnbr(row-1, column, area, grid)
	}
	if column > 0 {
		area = findnbr(row, column-1, area, grid)
	}
	if row+1 < len(grid) {
		area = findnbr(row+1, column, area, grid)
	}
	if column+1 < len(grid[row]) {
		area = findnbr(row, column+1, area, grid)
	}

	return area
}

func maxAreaOfIsland(grid [][]int) int {
	/*
	   mark not visisted ocean as 0, visisted land as 1, visited area as 2
	*/
	max := 0
	for row := range grid {
		for column := range grid[row] {
			if a := findnbr(row, column, 0, grid); a > max {
				max = a
			}
		}
	}
	return max
}