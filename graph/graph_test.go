package graph

import (
	"testing"
)

func TestNew(t *testing.T) {
	// when
	g := NewGraph()

	// then
	if g == nil {
		t.Fatal("NewGraph() returned nil")
	}
	if g.adj == nil {
		t.Error("adj map should be init")
	}
}

func TestAddEdge(t *testing.T) {

	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	neighbours, err := g.Neighbours(1)
	t.Logf("neightbours of 1: %v", neighbours)
	if err != nil {
		t.Fatalf("expected 1 node but got %v", err)
	}

	if len(neighbours) != 2 {
		t.Errorf("expected 2 neighbours, got %d", len(neighbours))
	}

}

func TestNeighbours_UnknownNode(t *testing.T) {
	g := NewGraph()
	_, err := g.Neighbours(1)
	if err == nil {
		t.Error("expected error for non-existent node, got nil")
	}
}

func TestNeighbours_Success(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	neighbours, err := g.Neighbours(1)

	if err != nil {
		t.Fatal("Node not found")
	}

	if len(neighbours) != 2 {
		t.Errorf("expected neighbours len 2 but got %d", len(neighbours))
	}

}

func TestBFS(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 22)
	g.AddEdge(22, 222)

	visited, err := g.BFS(1)

	if err != nil {
		t.Fatalf("error BFS: %v", err)
	}

	if len(visited) != 6 {
		t.Errorf("Expected all nodes visited but got %d", len(visited))
	}

	if !checkVisited(visited, 1, 2, 3, 4, 22, 222) {
		t.Errorf("Wrong visited: %v", visited)
	}
}

func checkVisited(visited []int, items ...int) bool {

	for index, item := range items {
		if visited[index] != item {
			return false
		}
	}

	return true
}
