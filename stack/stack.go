package stack

import "container/list"

type Stack struct {
	dll *list.List
}

func NewStack() *Stack {
	return &Stack{dll: list.New()}
}

func (stack *Stack) Push(element interface{}) {
	stack.dll.PushBack(element)
}

func (stack *Stack) Pop() interface{} {
	if stack.dll.Len() == 0 {
		return nil
	}

	lastElement := stack.dll.Back()
	lastElementValue := lastElement.Value

	stack.dll.Remove(lastElement)

	return lastElementValue
}

func (stack *Stack) Len() int {
	return stack.dll.Len()
}
