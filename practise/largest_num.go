package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 2, 4, 8, 10, 11, 10, 54, 56, 2, 3, 5, 7, 43, 24, 54, 56, 3, 21, 2}

	var index []int

	var max int = 0

	for i := 0; i < len(a); i++ {
		if a[i] >= max {
			max = a[i]
		}
	}
	fmt.Println(max)

	for val := range a {
		if a[val] == max {
			index = append(index, val)
		}
	}
	fmt.Println(index)
}
