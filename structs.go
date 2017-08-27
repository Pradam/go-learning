package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	var name = "Pradam"
	fmt.Println(Vertex{1, 2})
	fmt.Println(len(name))
	v := Vertex{1, 7}
	v.X = 50
	fmt.Println(v.X*v.Y, v.X, v.Y)
	// Struct Pointer
	p := &v
	p.X = 1e9
	fmt.Println(v.X, p.Y)
}
