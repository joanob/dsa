package dsa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateQueue(t *testing.T) {
	testCases := []struct {
		desc     string
		expected *queue[int]
	}{
		{
			desc: "empty list with capacity 8",
			expected: &queue[int]{
				list: make([]int, 0, 8),
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			queue := CreateQueue[int]()

			assert.Equal(t, tC.expected, queue)
		})
	}
}

func TestEnqueue(t *testing.T) {
	testCases := []struct {
		desc     string
		queue    *queue[int]
		data     int
		expected *queue[int]
	}{
		{
			desc: "on empty queue",
			queue: &queue[int]{
				list: []int{},
			},
			data: 7,
			expected: &queue[int]{
				list: []int{7},
			},
		},
		{
			desc: "on non empty queue",
			queue: &queue[int]{
				list: []int{5, 8},
			},
			data: 7,
			expected: &queue[int]{
				list: []int{5, 8, 7},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.queue.Enqueue(tC.data)

			assert.Equal(t, tC.expected, tC.queue)
		})
	}
}

func TestDequeue(t *testing.T) {
	testCases := []struct {
		desc         string
		queue        *queue[int]
		expected     *queue[int]
		expectedData int
		expectedErr  error
	}{
		{
			desc: "on empty queue",
			queue: &queue[int]{
				list: []int{},
			},
			expected: &queue[int]{
				list: []int{},
			},
			expectedData: 0,
			expectedErr:  errors.New(""),
		},
		{
			desc: "on non empty queue",
			queue: &queue[int]{
				list: []int{5, 8, 7},
			},
			expected: &queue[int]{
				list: []int{8, 7},
			},
			expectedData: 5,
			expectedErr:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			val, err := tC.queue.Dequeue()

			if tC.expectedErr != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tC.expectedData, val)

			assert.Equal(t, tC.expected, tC.queue)
		})
	}
}

func TestPeed(t *testing.T) {
	testCases := []struct {
		desc         string
		queue        *queue[int]
		expected     *queue[int]
		expectedData int
		expectedErr  error
	}{
		{
			desc: "on empty queue",
			queue: &queue[int]{
				list: []int{},
			},
			expected: &queue[int]{
				list: []int{},
			},
			expectedData: 0,
			expectedErr:  errors.New(""),
		},
		{
			desc: "on non empty queue",
			queue: &queue[int]{
				list: []int{5, 8, 7},
			},
			expected: &queue[int]{
				list: []int{5, 8, 7},
			},
			expectedData: 5,
			expectedErr:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			val, err := tC.queue.Peek()

			if tC.expectedErr != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tC.expectedData, val)

			assert.Equal(t, tC.expected, tC.queue)
		})
	}
}
