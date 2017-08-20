package main

import (
	"fmt"
)

func main() {
	i := 1991
	switch i {
	case 1991:
		fmt.Println("Pradam is here")
		fmt.Printf("%T %v", i, i)
	case 1992:
		fmt.Println("Pradam is not yet ready.")
	default:
		fmt.Printf("Fucked up my life.")
	}
}
