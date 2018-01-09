package main

import (
		"fmt"
		"strings"
	)



func convert_into_strings_func(x string) string{
	return strings.ToUpper("Gopher")
}


func main(){
	fmt.Println(convert_into_strings_func("Pradam"))
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
}