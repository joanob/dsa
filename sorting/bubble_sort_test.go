package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubble_sort(t *testing.T) {
	testCases := []struct {
		desc     string
		unsorted []int
		sorted   []int
	}{
		{
			desc:     "Empty",
			unsorted: []int{},
			sorted:   []int{},
		},
		{
			desc:     "One item",
			unsorted: []int{1},
			sorted:   []int{1},
		},
		{
			desc:     "Two items sorted",
			unsorted: []int{1, 2},
			sorted:   []int{1, 2},
		},
		{
			desc:     "Two items unsorted",
			unsorted: []int{3, 1},
			sorted:   []int{1, 3},
		},
		{
			desc:     "First item lower",
			unsorted: []int{2, 5, 4},
			sorted:   []int{2, 4, 5},
		},
		{
			desc:     "Last item greater",
			unsorted: []int{5, 4, 8},
			sorted:   []int{4, 5, 8},
		},
		{
			desc:     "Completely unsorted",
			unsorted: []int{7, 5, 8, 2, 4},
			sorted:   []int{2, 4, 5, 7, 8},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sorted := bubble_sort(tC.unsorted)

			assert.EqualValues(t, tC.sorted, sorted)
		})
	}
}
