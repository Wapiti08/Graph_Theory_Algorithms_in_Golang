package main

import (
	"container/list"
	"fmt"
)

type AdjGraph struct {
	// define the length of single word
	V int
	// define the neighbors of every word
	nbrs map[string]*list.List
}

func isOneCharDiff(s1, s2 string) bool {
	// check whether the length if different
	if len(s1) != len(s2) {
		return false
	}
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
		}
		if diffCount > 1 {
			return false
		}
	}
	return diffCount == 1
}

func stringInSlice(str string, strlist []string) bool {
	for _, v := range strlist {
		if v == str {
			return true
		}
	}
	return false
}

// return a adjgraph instance
func NewAdjGraph(v int, wordlist []string) *AdjGraph {
	g := &AdjGraph{
		V:    v,
		nbrs: make(map[string]*list.List, v),
	}
	// // check whether beginword is in wordlist
	// if !stringInSlice(beginword, wordlist) {
	//     adjgraph.nbrs[beginword]=list.New()
	// }

	for _, word := range wordlist {
		g.nbrs[word] = list.New()
	}

	return g
}

func FindNbr(g *AdjGraph, word string, wordlist []string) *AdjGraph {
	// check whether word is in the nbrs key
	if _, ok := g.nbrs[word]; !ok {
		g.nbrs[word] = list.New()
	}
	// traverse the key
	for key := range g.nbrs {
		// traverse the word in list
		for _, w := range wordlist {
			// define directed graph
			if isOneCharDiff(key, w) {
				g.nbrs[key].PushBack(w)
			}
		}
	}
	return g
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	// define the distance array
	distance := make(map[string]int)
	// parent := make(map[string]string)
	// record the access order
	queue := list.New()
	queue.PushBack(beginWord)
	// check whether has been visisted
	adjgraph := NewAdjGraph(len(wordList), wordList)

	// create the neighbors
	adjgraph = FindNbr(adjgraph, beginWord, wordList)

	visited := make(map[string]bool, len(adjgraph.nbrs))
	// initialize the first word
	visited[beginWord] = true
	distance[beginWord] = 1
	// parent[beginWord] = beginWord

	for queue.Len() > 0 {
		// extract the front node in queue
		f := queue.Front().Value.(string)
		queue.Remove(queue.Front())
		visited[f] = true
		// traverse the nbrs of picked word
		for e := adjgraph.nbrs[f].Front(); e != nil; e = e.Next() {
			nbr := e.Value.(string)
			if !visited[nbr] {
				queue.PushBack(nbr)
				distance[nbr] = distance[f] + 1
				// parent[nbr] = f
				visited[nbr] = true
			}

		}
	}
	// print out the shorted distance
	return distance[endWord]
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	short_word_number := ladderLength(beginWord, endWord, wordList)
	fmt.Println(short_word_number)
}
