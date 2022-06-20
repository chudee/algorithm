package main

import (
	"algorithm/internal/basicsort"
	"fmt"
	"math/rand"
	"time"
)

/**
# 버블 정렬

- 두 인접한 데이터를 비교해서, 앞에 있는 데이터가 뒤에 있는 데이터보다 크면, 자리를 바꾸는 정렬 알고리즘

## 링크
https://visualgo.net/en/sorting
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	var data []int
	for range [50]int{} {
		data = append(data, rand.Intn(100))
	}

	fmt.Println("before: ", data)
	result := basicsort.BubbleSort(data)
	fmt.Println("after : ", result)
}
