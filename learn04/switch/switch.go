package main

import (
	"fmt"
)

func main() {
	normalSwitch()
	boolSwitch()
	optionalSwitch()
}

func normalSwitch() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}
}

func boolSwitch() {
	n := 4
	switch {
		case n % 2 == 1:
			fmt.Println("odd")
		case n % 2 == 0:
			fmt.Println("even")
		default:
			fmt.Print("Unknown")
	}
}

func optionalSwitch() {
	n := 3
	switch x := n * n; x {
	case 10:
		fmt.Println("odd")
	case 9:
		fmt.Println("even")

	}
}

func typeAssertion(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("number")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unknown")
	}
}