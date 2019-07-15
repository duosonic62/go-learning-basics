package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Country struct {
	Point
	name string
}

func main() {
	basicStructure()
	complexStructure()
}

func basicStructure() {
	var pt Point
	pt.X = 10
	pt.Y = 12
	fmt.Println(pt)

	ptLiteral := Point{X: 1, Y: 2}
	fmt.Println(ptLiteral)
}

func complexStructure() {
	country := Country{
		name: "Japan",
		Point: Point {
			X: 100,
			Y: 200,
		},
	}
	fmt.Println(country)
}