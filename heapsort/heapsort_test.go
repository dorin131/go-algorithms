package heapsort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortAscending(t *testing.T) {
	tests := []struct {
		Given []int
	}{
		{[]int{4, 9, 10, 0, -4, 7}},
		{[]int{1000, 100, 10, 0, -1000}},
		{[]int{1, 1, 1, 1}},
		{[]int{}},
		{[]int{777}},
	}
	for n, test := range tests {
		correctlySorted := make([]int, len(test.Given))
		copy(correctlySorted, test.Given)
		result := SortAscending(test.Given)

		if len(result) != len(test.Given) {
			t.Errorf("[%d] Lengths don't match", n)
			return
		}

		sort.Slice(correctlySorted, func(i, j int) bool {
			return correctlySorted[i] < correctlySorted[j]
		})

		for k := range correctlySorted {
			if correctlySorted[k] != result[k] {
				t.Errorf("[%v] Expected: %v, Got: %v\n", n, correctlySorted, result)
				return
			}
		}
		fmt.Printf("[%v] Correctly sorted: %v\n", n, result)
	}
}

func TestSortDescending(t *testing.T) {
	tests := []struct {
		Given []int
	}{
		{[]int{4, 9, 10, 0, -4, 7}},
		{[]int{1000, 100, 10, 0, -1000}},
		{[]int{1, 1, 1, 1}},
		{[]int{}},
		{[]int{777}},
	}
	for n, test := range tests {
		correctlySorted := make([]int, len(test.Given))
		copy(correctlySorted, test.Given)
		result := SortDescending(test.Given)

		if len(result) != len(test.Given) {
			t.Errorf("[%d] Lengths don't match", n)
			return
		}

		sort.Slice(correctlySorted, func(i, j int) bool {
			return correctlySorted[i] > correctlySorted[j]
		})

		for k := range correctlySorted {
			if correctlySorted[k] != result[k] {
				t.Errorf("[%v] Expected: %v, Got: %v\n", n, correctlySorted, result)
				return
			}
		}
		fmt.Printf("[%v] Correctly sorted: %v\n", n, result)
	}
}
