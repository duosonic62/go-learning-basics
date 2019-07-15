package main

import (
	"fmt"
)

type (
	IntPair [2]int
	AreaMap map[string][2]float64
	CallBack func(i int) int
)

func main() {

	pair := IntPair{1, 2}
	amap := AreaMap{"Tokyo": {35.68488, 139.691706}}

	fmt.Println(pair)
	fmt.Println(amap)

	nums := []int{1, 2, 3}
	fmt.Println(sum(nums, func(i int) int {return i * i}))
}

func sum(ints []int, callback CallBack) int {
	var sum int 
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}