package main

import (
	"container/list"
	"fmt"
	"sw/visualizer/heap"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vector2 struct {
	X int
	Y int
}

type Node struct {
	Position Vector2
}

func (node *Node) GetNeighbours(matrix *[][]byte) []*Edge {
	var nodes []*Edge

	x := node.Position.X
	y := node.Position.Y

	orthogonalDirectionWeight := 1.0

	// TOP
	if y-1 >= 0 {
		nodes = append(nodes, &Edge{node, &Node{Position: Vector2{X: x, Y: y - 1}}, orthogonalDirectionWeight})
	}

	// DOWN
	if y+1 <= len(*matrix)-1 {
		nodes = append(nodes, &Edge{node, &Node{Position: Vector2{X: x, Y: y + 1}}, orthogonalDirectionWeight})
	}

	// RIGHT
	if x+1 <= len((*matrix)[0])-1 {
		nodes = append(nodes, &Edge{node, &Node{Position: Vector2{X: x + 1, Y: y}}, orthogonalDirectionWeight})
	}

	// LEFT
	if x-1 >= 0 {
		nodes = append(nodes, &Edge{node, &Node{Position: Vector2{X: x - 1, Y: y}}, orthogonalDirectionWeight})
	}

	return nodes
}

type Edge struct {
	From   *Node
	To     *Node
	Weight float64
}

type UcsNode struct {
	CurrentNode    *Node
	LastEdge       *Edge
	BackUcsPointer *UcsNode
	TravelCost     float64
}

func (node UcsNode) GetCompletePath() *list.List {
	path := list.New()

	currentNode := node.BackUcsPointer
	for currentNode != nil && currentNode.LastEdge != nil {
		path.PushFront(currentNode.LastEdge)
		currentNode = currentNode.BackUcsPointer
	}

	return path
}

func (node *UcsNode) GetCost() float64 {
	return node.TravelCost
}

type SearchResult struct {
	Success      bool
	CompletePath *list.List
}

func (result *SearchResult) ShowCompletePath() {
	for e := result.CompletePath.Front(); e != nil; e = e.Next() {
		val := e.Value.(*Edge)
		fmt.Printf("{%d, %d} ----> {%d, %d}\n", val.From.Position.X, val.From.Position.Y, val.To.Position.X, val.To.Position.Y)
	}
}

func containsNode(haystack []*Node, needle *Node) bool {
	// Not the most efficient method...
	contains := false
	for _, n := range haystack {
		if n.Position.X == needle.Position.X && n.Position.Y == needle.Position.Y {
			contains = true
		}
	}

	return contains
}

func findInLinkedListByIndex(linkedlist *list.List, index int) *Edge {
	if index < 0 || index > linkedlist.Len() {
		return nil
	}

	currentElement := linkedlist.Front()
	for i := 0; i < index; i++ {
		currentElement = currentElement.Next()
	}

	return currentElement.Value.(*Edge)
}

func win() {
	const WINDOW_WIDTH int32 = 450
	const WINDOW_HEIGHT int32 = 450

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Hello Window!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	const ROWS int32 = 9
	const BLOCK_SIZE int32 = 50

	matrix := [][]byte{
		{'S', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', 'F'},
	}

	start := Node{Position: Vector2{X: 0, Y: 0}}
	end := Node{Position: Vector2{X: 8, Y: 8}}

	pqueue := heap.MinHeap{}
	visited := []*Node{}

	startUcsNode := &UcsNode{&start, nil, nil, 0}

	pqueue.Push(startUcsNode)

	result := &SearchResult{false, nil}

	for pqueue.Len() > 0 {
		ucsNode := pqueue.Pop().(*UcsNode)
		currentNode := ucsNode.CurrentNode

		if currentNode.Position.X == end.Position.X && currentNode.Position.Y == end.Position.Y {
			result = &SearchResult{true, ucsNode.GetCompletePath()}
			break
		}

		if !containsNode(visited, currentNode) {
			visited = append(visited, currentNode)

			for _, edge := range currentNode.GetNeighbours(&matrix) {
				newCost := ucsNode.TravelCost + edge.Weight
				pqueue.Push(&UcsNode{edge.To, edge, ucsNode, newCost})
			}
		}
	}

	visitedIndex := 0
	pathIndex := 0
	fillInterval := float32(0.05)
	fillPathInterval := float32(0.3)

	intervalAccumulator := float32(0.0)
	pathIntervalAccumulator := float32(0.0)

	readyToDrawVisitedNodes := []*Node{}
	readyToDrawPathEdges := []*Edge{}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		intervalAccumulator += dt
		pathIntervalAccumulator += dt

		// Push nodes that should be drawn on interval basis
		if visitedIndex < len(visited) && intervalAccumulator >= fillInterval {
			readyToDrawVisitedNodes = append(readyToDrawVisitedNodes, visited[visitedIndex])
			visitedIndex += 1
			intervalAccumulator = 0.0
		}

		if result.CompletePath != nil && pathIndex < result.CompletePath.Len() && pathIntervalAccumulator >= fillPathInterval {
			readyToDrawPathEdges = append(readyToDrawPathEdges, findInLinkedListByIndex(result.CompletePath, pathIndex))
			pathIndex += 1
			pathIntervalAccumulator = 0.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := range ROWS {
			for j := range ROWS {
				// Draw start and end nodes
				if i == int32(start.Position.X) && j == int32(start.Position.Y) {
					rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Orange)
				}

				if i == int32(end.Position.X) && j == int32(end.Position.Y) {
					rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Green)
				}

				// Draw visited
				for _, v := range readyToDrawVisitedNodes {
					// Do not draw on top of the start node
					if i == int32(start.Position.X) && j == int32(start.Position.Y) {
						continue
					}

					if i == int32(v.Position.X) && j == int32(v.Position.Y) {
						rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.LightGray)
					}
				}

				// Draw path
				for _, e := range readyToDrawPathEdges {
					if i == int32(e.To.Position.X) && j == int32(e.To.Position.Y) {
						rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Beige)
					}
				}

				// Draw grid
				rl.DrawLine(i*BLOCK_SIZE, 0, i*BLOCK_SIZE, WINDOW_HEIGHT, rl.Black)
				rl.DrawLine(0, j*BLOCK_SIZE, WINDOW_WIDTH, j*BLOCK_SIZE, rl.Black)
			}
		}

		rl.EndDrawing()
	}
}

func main() {
	win()
}
