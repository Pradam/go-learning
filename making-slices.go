package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(len(a), cap(a), a)
}
