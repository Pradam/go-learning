package main

import "fmt"

func main() {
	list := []int{2, 4, 5, 6, 7, 87, 8, 9, 9, 90}

	for value := range list {
		fmt.Println(value)
	}
}
