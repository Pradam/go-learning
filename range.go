package main

import (
	"fmt"
)

func main() {
	var pow = []int{1, 2, 3, 4, 5, 6, 67, 7, 8, 8}

	for i, v := range pow {
		fmt.Printf("Index: %V | Values: %V\n", i, v)
	}
}
