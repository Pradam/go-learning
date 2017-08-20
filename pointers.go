package main

import (
	"fmt"
)

func main() {
	var i int = 100
	p := &i
	*p = 1000
	fmt.Println(i)
}
