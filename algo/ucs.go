package algo

import (
	"container/list"
	"fmt"

	"sw/visualizer/graph"
	"sw/visualizer/heap"
	"sw/visualizer/utils"
)

type UcsNode struct {
	CurrentNode    *graph.GridNode
	LastEdge       *graph.Edge
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
	Visited      []*graph.GridNode
}

func (result *SearchResult) ShowCompletePath() {
	for e := result.CompletePath.Front(); e != nil; e = e.Next() {
		val := e.Value.(*graph.Edge)
		fmt.Printf("{%d, %d} ----> {%d, %d}\n", val.From.Position.X, val.From.Position.Y, val.To.Position.X, val.To.Position.Y)
	}
}

func RunUcs(matrix *[][]byte, start *graph.GridNode, end *graph.GridNode) *SearchResult {
	pqueue := heap.MinHeap{}
	visited := []*graph.GridNode{}

	startUcsNode := &UcsNode{start, nil, nil, 0}

	pqueue.Push(startUcsNode)

	result := &SearchResult{false, nil, visited}

	for pqueue.Len() > 0 {
		ucsNode := pqueue.Pop().(*UcsNode)
		currentNode := ucsNode.CurrentNode

		if currentNode.Position.X == end.Position.X && currentNode.Position.Y == end.Position.Y {
			result = &SearchResult{true, ucsNode.GetCompletePath(), visited}
			break
		}

		if !utils.ContainsNode(visited, currentNode) {
			visited = append(visited, currentNode)

			for _, edge := range currentNode.GetNeighbours(matrix) {
				newCost := ucsNode.TravelCost + edge.Weight
				pqueue.Push(&UcsNode{edge.To, edge, ucsNode, newCost})
			}
		}
	}

	return result
}
