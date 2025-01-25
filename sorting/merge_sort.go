package dsa

func merge_sort(data []int) []int {
	length := len(data)

	if length < 2 {
		return data
	}

	left := merge_sort(data[:length/2])
	right := merge_sort(data[length/2:])

	leftIndex := 0
	rightIndex := 0

	for i := 0; i < length; i++ {
		if leftIndex >= len(left) {
			// Only add rights
			data[i] = right[rightIndex]
			rightIndex++
		} else if rightIndex >= len(right) {
			// Only add lefts
			data[i] = left[leftIndex]
			leftIndex++
		} else {
			if left[leftIndex] < right[rightIndex] {
				data[i] = left[leftIndex]
				leftIndex++
			} else {
				data[i] = right[rightIndex]
				rightIndex++
			}
		}
	}

	return data
}
