package utils

import (
	"container/list"
	"sw/visualizer/graph"
)

func ContainsNode(haystack []*graph.GridNode, needle *graph.GridNode) bool {
	// Not the most efficient method...
	contains := false
	for _, n := range haystack {
		if n.Row == needle.Row && n.Column == needle.Column {
			contains = true
		}
	}

	return contains
}

func FindInLinkedListByIndex(linkedlist *list.List, index int) any {
	if index < 0 || index > linkedlist.Len() {
		return nil
	}

	currentElement := linkedlist.Front()
	for i := 0; i < index; i++ {
		currentElement = currentElement.Next()
	}

	return currentElement.Value
}
