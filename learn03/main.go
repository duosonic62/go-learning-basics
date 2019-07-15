package main

import (
	"fmt"
)

func main() {
	var (
		boolValue = true
		s = "go string"
		array1 = [5]int{1, 2, 3, 4, 5}
		array2 = [5]int{5, 4, 3, 2, 1}
	)
	fmt.Println(boolValue)
	fmt.Printf("%v\n", s)

	fmt.Printf("%v\n", array1)
	array1 = array2
	array2[0] = 1000
	fmt.Printf("%v\n", array1)
	
	q, r := div(20, 4)
	fmt.Printf("%v , %v\n", q, r)

	f := func(x, y int) int { return x * y }
	fmt.Printf("%v\n", f(2, 3))

	fmt.Printf("%v\n", returnFunc()(2, 3))

	closureMain()
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func returnFunc() func(a, b int) int {
	return func(a, b int) int { return a * b }
}