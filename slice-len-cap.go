package main

import "fmt"

func main() {
	s := make([]int, 0, 10)
	fmt.Println(cap(s), len(s))
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Println(cap(s), len(s))
}
