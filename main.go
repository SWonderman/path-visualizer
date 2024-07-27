package main

import (
	"fmt"
    "container/list"
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

func win() {
	const WINDOW_WIDTH int32 = 450
	const WINDOW_HEIGHT int32 = 450

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Hello Window!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	const ROWS int32 = 9
	const BLOCK_SIZE int32 = 50

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := range ROWS {
			rl.DrawLine(i*BLOCK_SIZE, 0, i*BLOCK_SIZE, WINDOW_HEIGHT, rl.Black)
			rl.DrawLine(0, i*BLOCK_SIZE, WINDOW_WIDTH, i*BLOCK_SIZE, rl.Black)
		}

		rl.EndDrawing()
	}
}

type SearchResult struct {
    Success      bool
    CompletePath *list.List
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

func ucs() *SearchResult {
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

	for pqueue.Len() > 0 {
		ucsNode := pqueue.Pop().(*UcsNode)
		currentNode := ucsNode.CurrentNode

		if currentNode.Position.X == end.Position.X && currentNode.Position.Y == end.Position.Y {
			return &SearchResult{true, ucsNode.GetCompletePath()}
		}

		if !containsNode(visited, currentNode) {
			visited = append(visited, currentNode)

			for _, edge := range currentNode.GetNeighbours(&matrix) {
				newCost := ucsNode.TravelCost + edge.Weight
				pqueue.Push(&UcsNode{edge.To, edge, ucsNode, newCost})
			}
		}
	}

    return &SearchResult{false, nil}
}

func main() {
    searchResult := ucs()

    if searchResult.Success {
        for e := searchResult.CompletePath.Front(); e != nil; e = e.Next() {
            val := e.Value.(*Edge)
            fmt.Printf("{%d, %d} ----> {%d, %d}\n", val.From.Position.X, val.From.Position.Y, val.To.Position.X, val.To.Position.Y)
        }
    }
}
