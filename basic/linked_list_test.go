package dsa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLinkedList(t *testing.T) {
	testCases := []struct {
		desc     string
		expected *linkedList[int]
	}{
		{
			desc: "head is nil",
			expected: &linkedList[int]{
				head: nil,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			list := CreateLinkedList[int]()

			assert.Equal(t, tC.expected, list)
		})
	}
}

func Test_InsertLast(t *testing.T) {
	testCases := []struct {
		desc     string
		list     *linkedList[int]
		data     int
		expected *linkedList[int]
	}{
		{
			desc: "with empty list",
			list: CreateLinkedList[int](),
			data: 5,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: nil,
				},
			},
		},
		{
			desc: "with non empty list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: nil,
				},
			},
			data: 2,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 2,
						next: nil,
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.list.InsertLast(tC.data)

			assert.Equal(t, tC.expected, tC.list)
		})
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct {
		desc          string
		list          *linkedList[int]
		data          int
		pos           int
		expected      *linkedList[int]
		expectedError error
	}{
		{
			desc: "head on empty list",
			list: &linkedList[int]{},
			data: 8,
			pos:  0,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 8,
					next: nil,
				},
			},
			expectedError: nil,
		},
		{
			desc: "head on non empty list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 2,
						next: nil,
					},
				},
			},
			data: 8,
			pos:  0,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 8,
					next: &linkedListNode[int]{
						data: 5,
						next: &linkedListNode[int]{
							data: 2,
							next: nil,
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			desc: "middle node",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 2,
						next: nil,
					},
				},
			},
			data: 8,
			pos:  1,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 8,
						next: &linkedListNode[int]{
							data: 2,
							next: nil,
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			desc: "last node",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 2,
						next: nil,
					},
				},
			},
			data: 8,
			pos:  2,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 2,
						next: &linkedListNode[int]{
							data: 8,
							next: nil,
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			desc: "position not in list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
			data: 3,
			pos:  3, // 2 would be next available position
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
			expectedError: errors.New(""),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.list.Insert(tC.pos, tC.data)

			assert.Equal(t, tC.expected, tC.list)

			if tC.expectedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	testCases := []struct {
		desc     string
		list     *linkedList[int]
		pos      int
		expected *linkedList[int]
	}{
		{
			desc: "with empty list",
			list: &linkedList[int]{
				head: nil,
			},
			pos: 0,
			expected: &linkedList[int]{
				head: nil,
			},
		},
		{
			desc: "first position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
			pos: 0,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 7,
					next: nil,
				},
			},
		},
		{
			desc: "middle position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 4,
							next: nil,
						},
					},
				},
			},
			pos: 1,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 4,
						next: nil,
					},
				},
			},
		},
		{
			desc: "last position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 12,
							next: nil,
						},
					},
				},
			},
			pos: 2,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
		},
		{
			desc: "position not in list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
			pos: 2,
			expected: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: nil,
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.list.Delete(tC.pos)

			assert.Equal(t, tC.expected, tC.list)
		})
	}
}

func Test_Get(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "with empty list",
		},
		{
			desc: "first position",
		},
		{
			desc: "middle position",
		},
		{
			desc: "last position",
		},
		{
			desc: "position not in list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func Test_Length(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "with empty list",
		},
		{
			desc: "with non empty list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func Test_Search(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "with empty list",
		},
		{
			desc: "first position",
		},
		{
			desc: "middle position",
		},
		{
			desc: "last position",
		},
		{
			desc: "not found",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
