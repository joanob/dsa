package dsa

func bubble_sort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}
