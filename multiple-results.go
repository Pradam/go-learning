package main

import "fmt"

func swap(x string, y int, z string) (int, string, string) {
	return y, z, x
}

func main() {
	a, b, c := swap("hello", 23, "world")
	fmt.Println(a, b, c)
}
