package basicsort

func BubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		isSwap := false

		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isSwap = true
			}
		}
		
		if !isSwap {
			break
		}
	}

	return data
}
