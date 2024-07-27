package heap

import "testing"


func TestPushToEmptyHeap(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)

    if minHeap.Len() != 1 {
        t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 7 {
        t.Fatalf("Unexpected root of the heap. Expected 7 but got %d", (*minHeap)[0])
    }
}

func TestPushTwoElements(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)
    minHeap.Push(5)

    if minHeap.Len() != 2 {
        t.Fatalf("Unexpected length of the heap. Expected 2 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 5 {
        t.Fatalf("Unexpected root of the heap. Expected 5 but got %d", (*minHeap)[0])
    }
    
    expectedList := []int{5, 7}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }
}

func TestPushThreeElements(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)
    minHeap.Push(5)
    minHeap.Push(9)

    if minHeap.Len() != 3 {
        t.Fatalf("Unexpected length of the heap. Expected 3 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 5 {
        t.Fatalf("Unexpected root of the heap. Expected 5 but got %d", (*minHeap)[0])
    }
    
    expectedList := []int{5, 7, 9}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }
}

func TestPopHeapOneElement(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)

    if minHeap.Len() != 1 {
        t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 7 {
        t.Fatalf("Unexpected root of the heap. Expected 7 but got %d", (*minHeap)[0])
    }
    
    expectedList := []int{7}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }

    popedValue := minHeap.Pop()
    if *popedValue != 7 {
        t.Fatalf("Unexpected poped heap value. Expected 7 but got %d", *popedValue)
    }

    if minHeap.Len() != 0 {
        t.Fatalf("Unexpected length of the heap. Expected 0 but got %d", minHeap.Len())
    }
}

func TestPopHeapTwoElements(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)
    minHeap.Push(4)

    if minHeap.Len() != 2 {
        t.Fatalf("Unexpected length of the heap. Expected 2 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 4 {
        t.Fatalf("Unexpected root of the heap. Expected 4 but got %d", (*minHeap)[0])
    }
    
    expectedList := []int{4, 7}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }

    popedValue := minHeap.Pop()
    if *popedValue != 4 {
        t.Fatalf("Unexpected poped heap value. Expected 4 but got %d", *popedValue)
    }

    if minHeap.Len() != 1 {
        t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
    }

    expectedList = []int{7}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }
}

func TestPopHeapThreeElements(t *testing.T) {
    minHeap := &MinHeap{}
    minHeap.Push(7)
    minHeap.Push(4)
    minHeap.Push(8)

    if minHeap.Len() != 3 {
        t.Fatalf("Unexpected length of the heap. Expected 3 but got %d", minHeap.Len())
    }

    if (*minHeap)[0] != 4 {
        t.Fatalf("Unexpected root of the heap. Expected 4 but got %d", (*minHeap)[0])
    }
    
    expectedList := []int{4, 7, 8}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }

    popedValue := minHeap.Pop()
    if *popedValue != 4 {
        t.Fatalf("Unexpected poped heap value. Expected 4 but got %d", *popedValue)
    }

    if minHeap.Len() != 2 {
        t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
    }

    expectedList = []int{7, 8}
    for idx := range expectedList {
        if (*minHeap)[idx] != expectedList[idx] {
            t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
        }
    }
}
