package main

import (
		"fmt"
        "math"
		"strings"
        "reflect"
	)


func convert_into_strings_func(x string) string{
	return strings.ToUpper("Gopher")
}


func adder() func(int) int{
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func show_name() func(string) []int{
    names := make([]int,0)
    return func(name string) []int{
        for i:=0 ; i< len(name); i++{
            names = append(names, int(name[i]))
        return names
        }
    }
}

func main(){
    hypot := func(x,y float64) float64{
        return math.Sqrt(x*x + y*y)
    }
	fmt.Println(convert_into_strings_func("Pradam"))
    fmt.Println(hypot(3, 4))
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
    pos, neg := adder(), adder()
    fmt.Println()
    fmt.Println(reflect.TypeOf(pos), reflect.TypeOf(neg))
    for i:= 0; i < 10; i++{
        fmt.Println(pos(i), neg(-2*i))
    }
    get_val := show_name()
    fmt.Println(reflect.TypeOf(get_val))
}