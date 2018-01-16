package main 

import ("fmt"
"reflect"
"time")

func main(){

    var a int
    a = 20
    today := time.Now().Weekday() + 1
    fmt.Println(time.Saturday, today)
    switch b:=200; b {
    case b:
        fmt.Println("Number Matched.")
        fmt.Println(reflect.TypeOf(b))
    default:
        fmt.Println("Number didn't Matched.")
    }
}