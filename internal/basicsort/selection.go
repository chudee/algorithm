package basicsort

func SelectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		lowest := i

		for j := i + 1; j < len(data); j++ {
			if data[lowest] > data[j] {
				lowest = j
			}
		}

		if lowest != i {
			data[lowest], data[i] = data[i], data[lowest]
		}
	}

	return data
}
