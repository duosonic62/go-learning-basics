package main

import (
	"fmt"
)

func main() {
	arrayPointer()
}

func basicPointer() {
	i := 10  
	p := &i
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	pp := &p
	fmt.Printf("%T\n", pp)
	fmt.Println(**pp)
}

func arrayPointer() {
	p := &[3]int {1, 2, 3}
	fmt.Printf("%p\n", p)

	pow(p)
	for i, v := range *p {
		println(i, v)
	} 
}

func pow(p *[3]int) {
	for i:=0 ; i < len(p); i++ {
		(*p)[i] = (*p)[i] * (*p)[i]
		p[i] = p[i] - 1
	}
}