package performance

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arraySearchTesting(linear_search, t)
}

func BenchmarkLinearSearch(b *testing.B) {
	data := fill_array(1000000)
	for i := 0; i < b.N; i++ {
		linear_search(data, 654321)
	}
}

func TestBinarySearch(t *testing.T) {
	arraySearchTesting(binary_search, t)
}

func BenchmarkBinarySearch(b *testing.B) {
	data := fill_array(1000000)
	for i := 0; i < b.N; i++ {
		binary_search(data, 3)
	}
}

func linear_search(data []int, v int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == v {
			return i
		}
	}
	return -1
}

func binary_search(data []int, v int) int {
	left := 0
	middle := 0
	right := len(data) - 1

	if len(data) == 0 || v < data[left] || v > data[right] {
		return -1
	}

	if data[right] == v {
		return right
	}

	for left < right {
		middle = left + (right-left)/2
		if data[middle] == v {
			return middle
		} else if data[middle] < v {
			left = middle
		} else {
			right = middle
		}
	}

	return -1
}

func arraySearchTesting(search func([]int, int) int, t *testing.T) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, size := range sizes {
		t.Run("Size_"+fmt.Sprint(size), func(t *testing.T) {
			data := fill_array(size)

			tests := []struct {
				name     string
				value    int
				expected int
			}{
				{"Elemento en el medio", size / 2, size / 2},
				{"Elemento en el inicio", 0, 0},
				{"Elemento en el final", size - 1, size - 1},
				{"Elemento menor que el mínimo", -1, -1},
				{"Elemento mayor que el máximo", size, -1},
				{"Elemento intermedio inexistente", size/2 + 1, size/2 + 1},
				{"Elemento en índice 9", 9, 9},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result := search(data, tt.value)
					if result == tt.expected {
						fmt.Printf("%s completado\n\n", tt.name)
					} else {
						t.Errorf("Para %s en tamaño %d, se esperaba %d pero se obtuvo %d", tt.name, size, tt.expected, result)
					}
				})
			}
		})
	}
}

func fill_array(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}
	return array
}
