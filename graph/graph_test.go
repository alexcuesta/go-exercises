package graph

import "testing"

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
	t.Logf("Graph after first edge: %+v", g)
	g.AddEdge(1, 3)
	t.Logf("Graph after second edge: %+v", g)

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
