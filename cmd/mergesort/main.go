package main

import (
	"algorithm/internal/advancedsort"
	"fmt"
	"math/rand"
	"time"
)

/**
# 병합 정렬

- 재귀용법을 활용한 정렬 알고리즘
   	1. 리스트를 절반으로 잘라 비슷한 크기의 두 부분 리스트로 나눈다.
   	2. 각 부분 리스트를 재귀적으로 합병 정렬을 이용해 정렬한다.
	3. 두 부분 리스트를 다시 하나의 정렬된 리스트로 합병한다.

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
	result := advancedsort.MergeSort(data)
	fmt.Println("after : ", result)
}
