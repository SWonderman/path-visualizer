package main

import (
	"sw/visualizer/graph"
	"testing"
)

func TestGetNeighboursTopLeft(t *testing.T) {
	matrix := [][]byte{
		{'S', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
	}

	node := &graph.GridNode{Row: 0, Column: 0}
	neighbours := node.GetNeighbours(&matrix, make(map[graph.GridNode]bool), &[]byte{'x'})

	expected := []struct {
		row    int
		column int
	}{
		{1, 0},
		{0, 1},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.row != neighbours[idx].To.Row {
			t.Fatalf("Invalid neighbour row coordinate, expected %d but got %d", coords.row, neighbours[idx].To.Row)
		}

		if coords.column != neighbours[idx].To.Column {
			t.Fatalf("Invalid neighbour column coordinate, expected %d but got %d", coords.column, neighbours[idx].To.Column)
		}
	}
}

func TestGetNeighboursTopRight(t *testing.T) {
	matrix := [][]byte{
		{'-', '-', '-', 'S'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
	}

	node := &graph.GridNode{Row: 0, Column: 3}
	neighbours := node.GetNeighbours(&matrix, make(map[graph.GridNode]bool), &[]byte{'x'})

	expected := []struct {
		row    int
		column int
	}{
		{1, 3},
		{0, 2},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.row != neighbours[idx].To.Row {
			t.Fatalf("Invalid neighbour row coordinate, expected %d but got %d", coords.row, neighbours[idx].To.Row)
		}

		if coords.column != neighbours[idx].To.Column {
			t.Fatalf("Invalid neighbour column coordinate, expected %d but got %d", coords.column, neighbours[idx].To.Column)
		}
	}
}

func TestGetNeighboursBottomLeft(t *testing.T) {
	matrix := [][]byte{
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'S', '-', '-', '-'},
	}

	node := &graph.GridNode{Row: 4, Column: 0}
	neighbours := node.GetNeighbours(&matrix, make(map[graph.GridNode]bool), &[]byte{'x'})

	expected := []struct {
		row    int
		column int
	}{
		{3, 0},
		{4, 1},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.row != neighbours[idx].To.Row {
			t.Fatalf("Invalid neighbour row coordinate, expected %d but got %d", coords.row, neighbours[idx].To.Row)
		}

		if coords.column != neighbours[idx].To.Column {
			t.Fatalf("Invalid neighbour column coordinate, expected %d but got %d", coords.column, neighbours[idx].To.Column)
		}
	}
}

func TestGetNeighboursBottomRight(t *testing.T) {
	matrix := [][]byte{
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', 'S'},
	}

	node := &graph.GridNode{Row: 4, Column: 3}
	neighbours := node.GetNeighbours(&matrix, make(map[graph.GridNode]bool), &[]byte{'x'})

	expected := []struct {
		row    int
		column int
	}{
		{3, 3},
		{4, 2},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.row != neighbours[idx].To.Row {
			t.Fatalf("Invalid neighbour row coordinate, expected %d but got %d", coords.row, neighbours[idx].To.Row)
		}

		if coords.column != neighbours[idx].To.Column {
			t.Fatalf("Invalid neighbour column coordinate, expected %d but got %d", coords.column, neighbours[idx].To.Column)
		}
	}
}

func TestGetNeighboursMiddle(t *testing.T) {
	matrix := [][]byte{
		{'-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-'},
		{'-', '-', 'S', '-', '-'},
		{'-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-'},
	}

	node := &graph.GridNode{Row: 2, Column: 2}
	neighbours := node.GetNeighbours(&matrix, make(map[graph.GridNode]bool), &[]byte{'x'})

	expected := []struct {
		row    int
		column int
	}{
		{1, 2},
		{3, 2},
		{2, 3},
		{2, 1},
	}

	if len(neighbours) != 4 {
		t.Fatalf("Invalid number of neighbours found. Expected 4 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.row != neighbours[idx].To.Row {
			t.Fatalf("Invalid neighbour row coordinate, expected %d but got %d", coords.row, neighbours[idx].To.Row)
		}

		if coords.column != neighbours[idx].To.Column {
			t.Fatalf("Invalid neighbour column coordinate, expected %d but got %d", coords.column, neighbours[idx].To.Column)
		}
	}
}
