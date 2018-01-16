package main

import (
         "fmt"
         "reflect"
       )

var (
   
   sum int = 0
)

func main(){
	var local int = 10
	for i:=0; i < local; i++ {
		sum += i
	}
	fmt.Println(sum)
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
    fmt.Println(reflect.TypeOf(sum))
    for sum < 3000{
    	sum += sum
    }
    fmt.Println(sum)
    list := make([]int, 5, 5)
    fmt.Println(len(list), list)
    for i:=0;i < len(list);i++{
    	fmt.Println(list[i])
    	if (i == 4){
    		list[i] = 90
    	}
    }
     fmt.Println(len(list), list)

    boil := true

    for ;boil; {
        fmt.Println("Hello world")
        boil = false
    }
}