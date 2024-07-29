package graph

type GridNode struct {
	Position Vector2
}

func (node *GridNode) GetNeighbours(matrix *[][]byte) []*Edge {
	var nodes []*Edge

	x := node.Position.X
	y := node.Position.Y

	orthogonalDirectionWeight := 1.0

	// TOP
	if y-1 >= 0 {
		nodes = append(nodes, &Edge{node, &GridNode{Position: Vector2{X: x, Y: y - 1}}, orthogonalDirectionWeight})
	}

	// DOWN
	if y+1 <= len(*matrix)-1 {
		nodes = append(nodes, &Edge{node, &GridNode{Position: Vector2{X: x, Y: y + 1}}, orthogonalDirectionWeight})
	}

	// RIGHT
	if x+1 <= len((*matrix)[0])-1 {
		nodes = append(nodes, &Edge{node, &GridNode{Position: Vector2{X: x + 1, Y: y}}, orthogonalDirectionWeight})
	}

	// LEFT
	if x-1 >= 0 {
		nodes = append(nodes, &Edge{node, &GridNode{Position: Vector2{X: x - 1, Y: y}}, orthogonalDirectionWeight})
	}

	return nodes
}
