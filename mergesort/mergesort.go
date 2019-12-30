package mergesort

// Sort : sorts a slice using the Merge Sort algorithm
func Sort(input []int) []int {
	sorted := make([]int, len(input))
	copy(sorted, input)

	subSort(sorted, 0, len(input)-1)
	return sorted
}

func subSort(sorted []int, leftStart int, rightEnd int) {
	if leftStart >= rightEnd {
		return
	}
	middle := (leftStart + rightEnd) / 2
	subSort(sorted, leftStart, middle)
	subSort(sorted, middle+1, rightEnd)
	merge(sorted, leftStart, rightEnd)
}

func merge(sorted []int, leftStart int, rightEnd int) {
	temp := make([]int, len(sorted))
	copy(temp, sorted)

	leftEnd := (leftStart + rightEnd) / 2
	rightStart := leftEnd + 1

	left := leftStart
	right := rightStart
	index := leftStart

	for left <= leftEnd && right <= rightEnd {
		if sorted[left] < sorted[right] {
			temp[index] = sorted[left]
			left++
		} else {
			temp[index] = sorted[right]
			right++
		}
		index++
	}

	copy(temp[index:rightEnd+1], sorted[right:rightEnd+1])
	copy(temp[index:rightEnd+1], sorted[left:leftEnd+1])
	copy(sorted, temp)
}
