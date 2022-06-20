package basicsort

func InsertSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			} else {
				break
			}
		}
	}

	return data
}
