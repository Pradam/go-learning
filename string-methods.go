package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	s = "Pradam"
	s1 := "Abhilash"
	s2 := "pradam"
	s3 := "abhilash"
	_ = s3
	_ = s1
	fmt.Println(strings.Compare(s, s2))
	fmt.Println(strings.Contains(s, s2))
	fmt.Println(strings.Count(s))
}
