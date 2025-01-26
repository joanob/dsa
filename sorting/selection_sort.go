package dsa

func selection_sort(data []int) []int {
	sorted := make([]int, len(data))
	copy(sorted, data)

	for i := 0; i < len(sorted); i++ {
		min := sorted[i]
		minIndex := i
		for j := i; j < len(sorted); j++ {
			if sorted[j] < min {
				min = sorted[j]
				minIndex = j
			}
		}
		sorted[i], sorted[minIndex] = min, sorted[i]
	}

	return sorted
}
