package graph

import "fmt"

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
