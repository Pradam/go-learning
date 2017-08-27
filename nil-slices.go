package main

import "fmt"

func main() {
	var s []int
	s1 := []int{}
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%T %T", s, s1)
	if s1 == nil {
		fmt.Println("This List is Empty.!!")
	}
}
