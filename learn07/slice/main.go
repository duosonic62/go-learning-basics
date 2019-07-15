package main

import (
	"fmt"
)

func main() {
	makeSlice()
	easySlice()
	appendSlice()
	copySlice()
	fullSlice()
	forSlice()
	fmt.Printf("sum = %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("sum = %d\n", sum([]int{1, 2, 3}...))
}

func makeSlice() {
	// var s = []int
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func easySlice() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println (b)

	s := "abcdefg"[0:3]
	fmt.Println(s)
}

func appendSlice() {
	s := make([]int, 3, 5)
	fmt.Println(len(s), cap(s))
	
	s = append(s, 4)
	fmt.Println(len(s), cap(s))

	s = append(s, 5)
	fmt.Println(len(s), cap(s))

	s = append(s, 6, 7, 8)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
}

func copySlice() {
	a := []int {1, 2, 3, 4, 5}
	b := []int {6, 7}
	n := copy(a, b)
	fmt.Println(a, n)
}

func fullSlice() {
	fmt.Println("=== Full Slice ===")
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a=%v\n", a)

	s1 := a[2:4]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := a[2:4:4]
	fmt.Println(s2, len(s2), cap(s2))

	s3 := a[2:4:6]
	fmt.Println(s3, len(s3), cap(s3))

}

func forSlice() {
	fmt.Println("=== For Slice ===")
	s := []string{"apple", "banana", "orange"}
	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}

func sum(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}