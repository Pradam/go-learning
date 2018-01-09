package main

// Declare the strings within double quotes.

import (
        "fmt"
        "reflect")

const global = "I am global thing"


var (
   Tobe bool = true
   num int64 = 89
   dynamic string = "Dynamic Values declaration" 
)


const (
   big  = 1 << 100
   small  =  big >> 99
)

func main(){
	var a int = 10
	var b int
	b = 20
	c := 30
	var first_name, last_name string = "pradam", "abhilash"
	a = 77
	fmt.Printf("Standard Variable declaration %d", a)
	fmt.Println()
	fmt.Printf("After declaration %d", b)
	fmt.Println()
	fmt.Printf("Dynamic declaration %d", c)
	fmt.Println()
	fmt.Println(first_name + " " + last_name)
	fmt.Println(Tobe, num, dynamic)
	fmt.Printf(global)
	fmt.Println(reflect.TypeOf(dynamic))
}