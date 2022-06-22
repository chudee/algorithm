package advancedsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{1}, QuickSort([]int{1}))

	assert.Equal(t, []int{1, 2}, QuickSort([]int{2, 1}))

	assert.Equal(t, []int{0, 1, 2, 4, 5}, QuickSort([]int{5, 4, 0, 2, 1}))
}

func TestMerge(t *testing.T) {
	assert.Equal(t, []int{1, 3}, merge([]int{}, []int{1, 3}))

	assert.Equal(t, []int{1, 2, 3}, merge([]int{3}, []int{1, 2}))

	assert.Equal(t, []int{1, 2, 3}, merge([]int{1, 2}, []int{3}))

	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8}, merge([]int{0, 2}, []int{1, 3, 5, 6, 7, 8}))
}
