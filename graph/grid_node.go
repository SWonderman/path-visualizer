package graph

type GridNode struct {
	Row    int
	Column int
}

func (node *GridNode) GetNeighbours(matrix *[][]byte, customObstaclePositions map[GridNode]bool) []*Edge {
	var nodes []*Edge

	row := node.Row
	column := node.Column

	orthogonalDirectionWeight := 1.0

	// TOP
	if row-1 >= 0 && !customObstaclePositions[GridNode{row - 1, column}] {
		nodes = append(nodes, &Edge{node, &GridNode{Row: row - 1, Column: column}, orthogonalDirectionWeight})
	}

	// DOWN
	if row+1 <= len(*matrix)-1 && !customObstaclePositions[GridNode{row + 1, column}] {
		nodes = append(nodes, &Edge{node, &GridNode{Row: row + 1, Column: column}, orthogonalDirectionWeight})
	}

	// RIGHT
	if column+1 <= len((*matrix)[0])-1 && !customObstaclePositions[GridNode{row, column + 1}] {
		nodes = append(nodes, &Edge{node, &GridNode{Row: row, Column: column + 1}, orthogonalDirectionWeight})
	}

	// LEFT
	if column-1 >= 0 && !customObstaclePositions[GridNode{row, column - 1}] {
		nodes = append(nodes, &Edge{node, &GridNode{Row: row, Column: column - 1}, orthogonalDirectionWeight})
	}

	return nodes
}
