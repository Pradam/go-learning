package main

import (
	"fmt"
)

func pradam(x, y int, z float32) (h uint8, k float32) {
	h = uint8((x * y) / 21)
	k = float32(h + 7)
	return
}

func main() {
	fmt.Println(pradam(100, 1000, 1991.0))
}
