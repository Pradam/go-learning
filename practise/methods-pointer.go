package main

import "fmt"

type sun struct{
    X int
    Y string
    Z map[int]string
}

func (s *sun) display(){
    s.X = 100
    s.Y = "oil"
    s.Z = map[int]string{1:"killing",2:"jogging",3:"opponent"}
}

func pradam(v sun){
    v.X = v.X * 100
    v.Y = v.Y + " " + "world"
    for key := range v.Z{
        v.Z[key] = "abhilash"
    }
}

func main(){
    var tiger  = map[int]string{1:"pradam",2:"abhilash"}
    r := &sun{1,"abhi",tiger}
    r.display()
    pradam(*r)
    fmt.Println(r)
}