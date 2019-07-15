package main

import (
	"fmt"
)

func sub() {
	for {
		fmt.Println(sub)
	}
}

func main() {
	go sub()
	for {
		fmt.Println("main")
	}
}