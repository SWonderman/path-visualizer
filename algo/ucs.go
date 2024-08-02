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
		fmt.Printf("{%d, %d} ----> {%d, %d}\n", val.From.Row, val.From.Column, val.To.Row, val.To.Column)
	}
}

func RunUcs(matrix *[][]byte, start *graph.GridNode, end *graph.GridNode, customObstaclePositions map[graph.GridNode]bool, obstacles *[]byte) *SearchResult {
	pqueue := heap.MinHeap{}
	visited := []*graph.GridNode{}

	startUcsNode := &UcsNode{start, nil, nil, 0}

	pqueue.Push(startUcsNode)

	result := &SearchResult{false, nil, visited}

	for pqueue.Len() > 0 {
		ucsNode := pqueue.Pop().(*UcsNode)
		currentNode := ucsNode.CurrentNode

		if currentNode.Row == end.Row && currentNode.Column == end.Column {
			result = &SearchResult{true, ucsNode.GetCompletePath(), visited}
			break
		}

		if !utils.ContainsNode(visited, currentNode) {
			visited = append(visited, currentNode)

			for _, edge := range currentNode.GetNeighbours(matrix, customObstaclePositions, obstacles) {
				newCost := ucsNode.TravelCost + edge.Weight
				pqueue.Push(&UcsNode{edge.To, edge, ucsNode, newCost})
			}
		}
	}

	// Include in the search result which nodes were visited by the algorithm
	// when a path to the end/goal could not be found
	if !result.Success {
		result.Visited = visited
	}

	return result
}
