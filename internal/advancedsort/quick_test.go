package advancedsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	assert.Equal(t, []int{1}, QuickSort([]int{1}))

	assert.Equal(t, []int{1, 2}, QuickSort([]int{2, 1}))

	assert.Equal(t, []int{0, 1, 2, 4, 5}, QuickSort([]int{5, 4, 0, 2, 1}))
}
