package dsa

func insertion__sort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	sorted := make([]int, 0, len(data))

	sorted = append(sorted, data[0])

	for i := 1; i < len(data); i++ {
		element := data[i]
		for j := 0; j < len(sorted); j++ {
			if element < sorted[j] {
				sorted = append(sorted, element)
				break
			}
		}
		if len(sorted) == i {
			sorted = append(sorted, element)
		}
	}

	return sorted
}
