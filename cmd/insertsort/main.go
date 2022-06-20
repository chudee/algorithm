package main

import (
	"algorithm/internal/basicsort"
	"fmt"
	"math/rand"
	"time"
)

/**
# 삽입 정렬

- 삽입 정렬은 두 번째 인덱스 부터 시작
- 해당 인덱스(key 값) 앞에 있는 데이터(B)부터 비교해서 key 값이 더 작으면, B값을 뒤 인덱스로 복사
- 이를 key 값이 더 큰 데이터를 만날때까지 반복, 그리고 큰 데이터를 만난 위치 바로 뒤에 key 값을 이동

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
	result := basicsort.InsertSort(data)
	fmt.Println("after : ", result)
}
