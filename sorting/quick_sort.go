package dsa

func quick_sort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	pivot := data[len(data)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(data)-1; i++ {
		if data[i] < pivot {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	return append(append(quick_sort(left), pivot), quick_sort(right)...)
}
