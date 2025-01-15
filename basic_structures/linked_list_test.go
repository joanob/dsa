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
		desc        string
		list        *linkedList[int]
		pos         int
		expected    int
		expectedErr error
	}{
		{
			desc:        "with empty list",
			list:        &linkedList[int]{},
			pos:         0,
			expected:    0,
			expectedErr: errors.New(""),
		},
		{
			desc: "first position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
				},
			},
			pos:         0,
			expected:    5,
			expectedErr: nil,
		},
		{
			desc: "middle position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 8,
						},
					},
				},
			},
			pos:         1,
			expected:    7,
			expectedErr: nil,
		},
		{
			desc: "last position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 8,
						},
					},
				},
			},
			pos:         2,
			expected:    8,
			expectedErr: nil,
		},
		{
			desc: "position not in list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 8,
						},
					},
				},
			},
			pos:         3,
			expected:    0,
			expectedErr: errors.New(""),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			val, err := tC.list.Get(tC.pos)

			if tC.expectedErr != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tC.expected, val)
		})
	}
}

func Test_Length(t *testing.T) {
	testCases := []struct {
		desc     string
		list     *linkedList[int]
		expected int
	}{
		{
			desc:     "with empty list",
			list:     &linkedList[int]{},
			expected: 0,
		},
		{
			desc: "with non empty list",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 5,
					next: &linkedListNode[int]{
						data: 7,
						next: &linkedListNode[int]{
							data: 8,
						},
					},
				},
			},
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			len := tC.list.Length()

			assert.Equal(t, tC.expected, len)
		})
	}
}

func Test_Search(t *testing.T) {
	testCases := []struct {
		desc        string
		list        *linkedList[int]
		data        int
		expected    int
		expectedErr error
	}{
		{
			desc:        "with empty list",
			list:        &linkedList[int]{},
			data:        5,
			expected:    0,
			expectedErr: errors.New(""),
		},
		{
			desc: "first position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 6,
					next: &linkedListNode[int]{
						data: 8,
					},
				},
			},
			data:        6,
			expected:    0,
			expectedErr: nil,
		},
		{
			desc: "middle position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 6,
					next: &linkedListNode[int]{
						data: 8,
						next: &linkedListNode[int]{
							data: 2,
						},
					},
				},
			},
			data:        8,
			expected:    1,
			expectedErr: nil,
		},
		{
			desc: "last position",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 6,
					next: &linkedListNode[int]{
						data: 8,
						next: &linkedListNode[int]{
							data: 2,
						},
					},
				},
			},
			data:        2,
			expected:    2,
			expectedErr: nil,
		},
		{
			desc: "not found",
			list: &linkedList[int]{
				head: &linkedListNode[int]{
					data: 6,
					next: &linkedListNode[int]{
						data: 8,
						next: &linkedListNode[int]{
							data: 2,
						},
					},
				},
			},
			data:        9,
			expected:    0,
			expectedErr: errors.New(""),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			pos, err := tC.list.Search(tC.data)

			if tC.expectedErr != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tC.expected, pos)
		})
	}
}
