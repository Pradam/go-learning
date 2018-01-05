package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Taking First Line.")
	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}
}

func main() {
	defer hello()
	fmt.Println("Hello")
	if i := 90; i > 80 {
		fmt.Println("Pradam!!")
	}
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		defer fmt.Println("Taking from bottom up.")
	}
	fmt.Printf("done")
}
