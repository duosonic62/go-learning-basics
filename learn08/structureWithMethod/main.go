package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x * x + y * y))
}

func main() {
	p := &Point{X:5, Y:10}
	p.Render()
	dp := &Point{X: 22, Y:-3}
	fmt.Println(p.Distance(dp))

}