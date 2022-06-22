package advancedsort

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	m := len(data) / 2
	left := MergeSort(data[:m])
	right := MergeSort(data[m:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var list []int
	lp, rp := 0, 0

	for {
		if lp >= len(left) {
			list = append(list, right[rp:]...)
			break
		}
		if rp >= len(right) {
			list = append(list, left[lp:]...)
			break
		}

		if left[lp] > right[rp] {
			list = append(list, right[rp])
			rp++
		} else {
			list = append(list, left[lp])
			lp++
		}
	}

	return list
}
