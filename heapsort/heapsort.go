package heapsort

import (
	"github.com/dorin131/go-data-structures/maxheap"
	"github.com/dorin131/go-data-structures/minheap"
)

// SortAscending : sort ascending a slice of integers using Heap Sort
func SortAscending(input []int) []int {
	result := []int{}

	mh := minheap.New(input)

	for range input {
		result = append(result, mh.ExtractMin())
	}

	return result
}

// SortDescending : sort descending a slice of integers using Heap Sort
func SortDescending(input []int) []int {
	result := []int{}

	mh := maxheap.New(input)

	for range input {
		result = append(result, mh.ExtractMax())
	}

	return result
}
