package dsa

import (
	"math/rand"
)

func random_array(length int) []int {
	data := make([]int, length)

	for i := 0; i < length; i++ {
		data[i] = rand.Int()
	}

	return data
}
