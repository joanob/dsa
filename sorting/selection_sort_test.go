package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_selection_sort(t *testing.T) {
	testCases := []struct {
		desc   string
		length int
	}{
		{
			desc:   "Empty",
			length: 0,
		},
		{
			desc:   "One item",
			length: 1,
		},
		{
			desc:   "Two items ",
			length: 2,
		},
		{
			desc:   "50 items",
			length: 50,
		},
		{
			desc:   "100 items",
			length: 100,
		},
		{
			desc:   "1000 items",
			length: 1000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data := random_array(tC.length)

			sorted := bubble_sort(data)

			selectionSorted := selection_sort(data)

			assert.EqualValues(t, sorted, selectionSorted)
		})
	}
}
