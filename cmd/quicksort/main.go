package main

import (
	"algorithm/internal/advancedsort"
	"fmt"
	"math/rand"
	"time"
)

/**
# 퀵 정렬

- 다음과 같은 순서를 반복하며 정렬하는 알고리즘
   	1. 기준점(pivot)을 정해서, 기준점보다 작은 데이터는 왼쪽(left), 큰 데이터는 오른쪽(right)으로 모으는 함수를 작성 함
   	2. 각 왼쪽(left), 오른쪽(right)은 재귀용법을 사용해서 다시 동일 함수를 호출하여 위 작업을 반복함
	3. 함수는 왼쪽(left) + 기준점(pivot) + 오른쪽(right)을 리턴함

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
	result := advancedsort.QuickSort(data)
	fmt.Println("after : ", result)
}
