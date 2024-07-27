package main

import "testing"

func TestGetNeighboursTopLeft(t *testing.T) {
	matrix := [][]byte{
		{'S', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
	}

	node := &Node{Position: Vector2{X: 0, Y: 0}}
	neighbours := node.GetNeighbours(&matrix)

	expected := []struct {
		x int
		y int
	}{
		{0, 1},
		{1, 0},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.x != neighbours[idx].To.Position.X {
			t.Fatalf("Invalid neighbour x coordinate, expected %d but got %d", coords.x, neighbours[idx].To.Position.X)
		}

		if coords.y != neighbours[idx].To.Position.Y {
			t.Fatalf("Invalid neighbour y coordinate, expected %d but got %d", coords.y, neighbours[idx].To.Position.Y)
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

	node := &Node{Position: Vector2{X: 3, Y: 0}}
	neighbours := node.GetNeighbours(&matrix)

	expected := []struct {
		x int
		y int
	}{
		{3, 1},
		{2, 0},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.x != neighbours[idx].To.Position.X {
			t.Fatalf("Invalid neighbour x coordinate, expected %d but got %d", coords.x, neighbours[idx].To.Position.X)
		}

		if coords.y != neighbours[idx].To.Position.Y {
			t.Fatalf("Invalid neighbour y coordinate, expected %d but got %d", coords.y, neighbours[idx].To.Position.Y)
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

	node := &Node{Position: Vector2{X: 0, Y: 4}}
	neighbours := node.GetNeighbours(&matrix)

	expected := []struct {
		x int
		y int
	}{
		{0, 3},
		{1, 4},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.x != neighbours[idx].To.Position.X {
			t.Fatalf("Invalid neighbour x coordinate, expected %d but got %d", coords.x, neighbours[idx].To.Position.X)
		}

		if coords.y != neighbours[idx].To.Position.Y {
			t.Fatalf("Invalid neighbour y coordinate, expected %d but got %d", coords.y, neighbours[idx].To.Position.Y)
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

	node := &Node{Position: Vector2{X: 3, Y: 4}}
	neighbours := node.GetNeighbours(&matrix)

	expected := []struct {
		x int
		y int
	}{
		{3, 3},
		{2, 4},
	}

	if len(neighbours) != 2 {
		t.Fatalf("Invalid number of neighbours found. Expected 2 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.x != neighbours[idx].To.Position.X {
			t.Fatalf("Invalid neighbour x coordinate, expected %d but got %d", coords.x, neighbours[idx].To.Position.X)
		}

		if coords.y != neighbours[idx].To.Position.Y {
			t.Fatalf("Invalid neighbour y coordinate, expected %d but got %d", coords.y, neighbours[idx].To.Position.Y)
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

	node := &Node{Position: Vector2{X: 2, Y: 2}}
	neighbours := node.GetNeighbours(&matrix)

	expected := []struct {
		x int
		y int
	}{
		{2, 1},
		{2, 3},
		{3, 2},
		{1, 2},
	}

	if len(neighbours) != 4 {
		t.Fatalf("Invalid number of neighbours found. Expected 4 but got %d", len(neighbours))
	}

	for idx, coords := range expected {
		if coords.x != neighbours[idx].To.Position.X {
			t.Fatalf("Invalid neighbour x coordinate, expected %d but got %d", coords.x, neighbours[idx].To.Position.X)
		}

		if coords.y != neighbours[idx].To.Position.Y {
			t.Fatalf("Invalid neighbour y coordinate, expected %d but got %d", coords.y, neighbours[idx].To.Position.Y)
		}
	}
}
