package main

import "fmt"

var (
	b rune = 10
)

func main() {
	a := 0
	for a < 1000 {
		a += 1
		fmt.Println(a)
	}

	fmt.Println(string(a) + "i")
}
