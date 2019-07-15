package main

import (
	"fmt"
)

type Base struct {
	Id int
	Owner string 
}

type A struct {
	Base
	Name string
	Area string
}

type B struct {
	Base
	Title string
	Bodies []string
}

func main() {
	a := A {
		Base: Base {17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}	

	b := B{
		Base: Base{81, "Hanako"},
		Title: "no title",
		Bodies: [] string{"a", "b"},
	}

	fmt.Println(a)
	fmt.Println(b)

	pa := new(A)
	swap(pa)
	fmt.Println(*pa)
}

func swap(a *A) {
	a.Name = "new Name"
	a.Area = "New Area"
	a.Base = Base {
		Id: 10,
		Owner: "NewOwner",
	}
}