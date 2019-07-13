package main

import (
	"fmt"
)

func main() {
	makeMap()
	literalMap()
	getMapValue()
	forMap()
	deleteMap()
}

func makeMap() {
	m := make(map[int]string)
	m[1] 	= "US"
	m[81]	= "Japan"
	m[86]	= "China"
	fmt.Println(m)
}

func literalMap() {
	m := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m)

	ms := map[int][]string {
		1: {"a", "b"},
		2: {"c"},
		3: {"d", "e"},
	}
	fmt.Println(ms)
}

func getMapValue() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	s, ok := m[1]
	println(s, ok)
	s, ok = m[9]
	println(s, ok)

	if _, ok := m[1]; ok {
		fmt.Println("Exists[1]")
	}
}

func forMap() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
	}
}

func deleteMap() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	delete(m, 2)
	fmt.Println(m)
}