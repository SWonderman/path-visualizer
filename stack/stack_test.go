package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expectedFromPop := [3]int{3, 2, 1}

	currentValue := stack.Pop()
	currentIdx := 0
	for currentValue != nil {
		if currentValue != expectedFromPop[currentIdx] {
			t.Fatalf("Unexpected value received from stack pop. Expected %d, but got %d.", expectedFromPop[currentIdx], currentValue)
		}
		currentIdx += 1
		currentValue = stack.Pop()
	}
}
