package mergesort

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		Given    []int
		Expected []int
	}{
		{[]int{2, 1}, []int{1, 2}},
		{[]int{4, 9, 10, 0, -4, 7}, []int{-4, 0, 4, 7, 9, 10}},
		{[]int{1000, 100, 10, 0, -1000}, []int{-1000, 0, 10, 100, 1000}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
		{[]int{777}, []int{777}},
	}
	for n, test := range tests {
		result := Sort(test.Given)
		if len(result) != len(test.Expected) {
			t.Errorf("[%d] Lengths don't match", n)
		}
		for i, v := range result {
			if v != test.Expected[i] {
				t.Errorf("[%d] Expected: %v Got: %v", n, test.Expected, result)
			}
		}
	}
}
