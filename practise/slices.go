package main

import (
    "fmt"
    "unsafe"
    "reflect"
    )

func main(){
   list_of_int := []int{2,3,4,5,6,7,8,9,10}
   var check_nil []int
   fmt.Println(list_of_int)
   fmt.Println(unsafe.Sizeof(list_of_int))
   fmt.Println(reflect.TypeOf(list_of_int).Size())
   fmt.Println(cap(list_of_int))
   for i := 0; i < len(list_of_int); i++{
        fmt.Println(list_of_int[i])
   }
   fmt.Println(list_of_int[:4])
   fmt.Println(list_of_int[:])
   a := make([]int, 5, 5)
   for i:=0; i < 5; i++{
      a[i] = i

      }
   fmt.Println(a)
   if check_nil == nil{
    check_nil := append(check_nil, []int{1,2,3,4}...)
    fmt.Println("It's a nil value.",check_nil)
   }
   slice := append([]byte("hello "), "world"...)
   fmt.Println(slice)

   // var pow = []int{1,2,4,8,16,32,64,128}

   for i, v := range slice{
       fmt.Println(i,v)
   }
   check_nil = append(check_nil, 1,3,4,5,6,7,8)
   fmt.Println(check_nil)

   var new_lis = []int{1,2,3,4,5,56,6}
   
   fmt.Println(new_lis)
}