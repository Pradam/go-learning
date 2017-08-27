package main

import "fmt"

type Pradam struct {
	X, Y int
}

func main() {
	names := [4]string{
		"Blue",
		"White",
		"Red",
		"Orange",
	}
	x := Pradam{1, 2}
	y := &x
	y.X = 7
	fmt.Println(names)
	fmt.Println(x.X, x.Y)
	fmt.Println(names[:2+1])
	names[0] = "Idoit Boy"
	fmt.Println(names)
}
