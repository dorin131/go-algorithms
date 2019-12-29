package insertionsort

// Sort : sorts a slice of integers using the Insertion Sort algorithm
func Sort(a []int) []int {
	result := make([]int, len(a))
	copy(result, a)
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
