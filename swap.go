package main

import "fmt"

func swap(x string, y int) (int, string) {
	return y, x
}

func main() {
	a, b := swap("Hello ", 2)
	fmt.Print(a, b)
	fmt.Println()
}
