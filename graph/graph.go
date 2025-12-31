package graph

import (
	"fmt"
	"slices"
)

type Graph struct {
	adj map[int][]int // int to slice of ints
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(from, to int) {
	neighbours := g.adj[from]
	neighbours = append(neighbours, to)
	g.adj[from] = neighbours
}

func (g *Graph) Neighbours(node int) ([]int, error) {
	neighbours, exists := g.adj[node]
	if !exists {
		return nil, fmt.Errorf("node %d does not exist", node)
	}
	return neighbours, nil
}

func (g *Graph) BFS(start int) ([]int, error) {
	// enqueue first item
	var queue []int
	var visited []int
	queue = append(queue, start)

	// while queue not empty
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if !slices.Contains(visited, current) {
			visited = append(visited, current)
			if neighbours, err := g.Neighbours(current); err == nil {
				queue = append(queue, neighbours...)
			}
		}

	}

	return visited, nil
}
