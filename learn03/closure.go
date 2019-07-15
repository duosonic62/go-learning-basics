package main

import (
	"fmt"
)

func closureMain() {
	a := closure()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := closure()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func closure() func() int {
	i := 0
	return func() int { 
		i = i + 1 
		return i 
	}
}