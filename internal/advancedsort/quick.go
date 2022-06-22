package advancedsort

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	pivot := data[0]
	var left []int
	var right []int

	for i := 1; i < len(data); i++ {
		if pivot > data[i] {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
