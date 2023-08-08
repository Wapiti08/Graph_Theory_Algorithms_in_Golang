package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'journeyToMoon' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY astronaut
 */

type Graph struct {
	V    int32
	nbrs []*list.List
}

func NewGraph(v int32) *Graph {
	g := &Graph{
		V:    v,
		nbrs: make([]*list.List, v),
	}
	for i := int32(0); i < v; i++ {
		g.nbrs[i] = list.New()
	}

	return g
}
func addEdges(g *Graph, i int32, j int32, undir bool) *Graph {
	g.nbrs[i].PushBack(j)
	if undir {
		g.nbrs[j].PushBack(i)
	}
	return g
}

func dfshelper(g *Graph, node int32, visited []bool) int32 {
	var size int32
	size = 1
	visited[node] = true

	for e := g.nbrs[node].Front(); e != nil; e = e.Next() {
		nbr := e.Value.(int32)
		if !visited[nbr] {
			size += dfshelper(g, nbr, visited)
		}
	}
	return size
}

func journeyToMoon(n int32, astronaut [][]int32) int32 {
	// Write your code here
	visited := make([]bool, n)
	var output int32
	// calculate the total edges
	output = n * (n - 1) / 2
	g := NewGraph(n)
	for row := range astronaut {
		g = addEdges(g, astronaut[row][0], astronaut[row][1], true)
	}
	var cns int32
	for i := int32(0); i < n; i++ {
		// bypass the visited node
		if !visited[i] {
			// minus the all possible edges from the same country
			cns = dfshelper(g, i, visited)
			output -= cns * (cns - 1) / 2
		}
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	p := int32(pTemp)

	var astronaut [][]int32
	for i := 0; i < int(p); i++ {
		astronautRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var astronautRow []int32
		for _, astronautRowItem := range astronautRowTemp {
			astronautItemTemp, err := strconv.ParseInt(astronautRowItem, 10, 64)
			checkError(err)
			astronautItem := int32(astronautItemTemp)
			astronautRow = append(astronautRow, astronautItem)
		}

		if len(astronautRow) != 2 {
			panic("Bad input")
		}

		astronaut = append(astronaut, astronautRow)
	}

	result := journeyToMoon(n, astronaut)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
