package basicsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	assert.Equal(t, []int{1}, InsertSort([]int{1}))

	assert.Equal(t, []int{1, 2}, InsertSort([]int{2, 1}))
	
	assert.Equal(t, []int{0, 1, 2, 4, 5}, InsertSort([]int{5, 4, 0, 2, 1}))
}
