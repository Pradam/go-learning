package main

import "fmt"

func add(sum int) (x, y int) {
	x = sum * 2
	y = x / 2
	return
}

func main() {
	fmt.Println(add(200))
}
