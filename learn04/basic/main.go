package main

import (
	"fmt"
	"errors"
)

func main()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Printf("%v\n", i)
		}
	}

	x, y := 3,5
	if n := x * y; n % 2 == 0 {
		fmt.Printf("even %v\n", n)
	} else {
		fmt.Printf("odd  %v\n", n)
	}

	if _, err := throwException(); err != nil {
		fmt.Printf("%v", err)
	}

	forSample()
	forArrays()
}

func throwException() (int, error) {
	return 1, errors.New("erorr")
}

func forSample() {
	i := 0
	for {
		fmt.Printf("%v\n", i)
		if i == 3 { 
			break
		}
		i++
	}

	for i < 10 {
		i++
		if i % 2 ==0 {
			continue
		}
		fmt.Printf("%v\n", i)
	}
}

func forArrays() {
	fruits := [3]string{"apple", "banana", "orange"}
	for i, s := range fruits {
		fmt.Printf("%d : %s\n", i, s)
	}
}