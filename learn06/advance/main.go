package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Yeah!")
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())
	fmt.Printf("version: %s\n", runtime.Version())
}