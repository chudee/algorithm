package basicsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, []int{1}, SelectionSort([]int{1}))

	assert.Equal(t, []int{1, 2}, SelectionSort([]int{2, 1}))

	assert.Equal(t, []int{0, 1, 2, 4, 5}, SelectionSort([]int{5, 4, 0, 2, 1}))
}
