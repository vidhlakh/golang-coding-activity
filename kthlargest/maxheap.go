package kthlargest

import "fmt"

// Max heap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// HEap is complete tree
// Complete tree means all levels are full but  lower level can be empty except lower left
// Binary heap -> All nodes have ax of 2 childern
//Parent index* 2 +1 = left child index
// parent index * 2 +2 = right child index

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

//maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	// when parent smaller than current index
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Extract return the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	N := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Empty heap")
		return -1
	}
	h.array[0] = h.array[N]
	h.array = h.array[:N]
	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has atleast one child
	for l <= lastIndex {
		// when index is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else {
			// when right child is larger
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get parent index
// left child odd number. right child even number
func parent(i int) int {
	return (i - 1) / 2
}

// left child index
func left(i int) int {
	return i*2 + 1
}

// right child index
func right(i int) int {
	return i*2 + 2
}

// swap keys of index 1 and index 2  in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
	// temp := h.array[i2]
	// h.array[i2] = h.array[i1]
	// h.array[i1] = temp
}
