package main

import ("fmt"
        "math")

type Abser interface{
    M() float64
}

 type Pradam float64

 func (b Pradam) M() float64{
    return math.Sqrt(float64(b))
 }

 func main(){
    var a Abser
    mapping := map[string]int{}
    mapping["pradma"] = 1
    kill := Pradam(45)
    a = kill
    fmt.Println(a.M(),mapping)
 }