package main

import (
	"fmt"
	"os"
)

func main() {
	runDeffer()
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
	runPanic(5)
}

func runDeffer() {
	fmt.Println("open")

	file, err := os.Open("path/file")

	defer func() {
		file.Close()
		fmt.Println("closed file")
	}()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func runPanic(x int) {
	fmt.Println("start")
	defer fmt.Println("defer process")
	if x % 2 == 1 {
		panic("runtime error")
	}	
	fmt.Println("end")
}

func testRecover(src interface{}) {
	defer func () {
		if x:= recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Printf("panic: Unknown")
			}
		}
	}()
	panic(src)
	return
}