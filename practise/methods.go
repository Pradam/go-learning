package main 

import "fmt"
import "math"
type Vertex struct{
    X, Y float64
}

func (a *Vertex) Abs() float64{
    return math.Sqrt(a.X+a.Y)
}


func (array []int) Getsum() []int{
    for i:=0;i<len(array);i++{
        array = append(array, array[i]*2)
    }
    return array
}

func main(){
    v := &Vertex{3,4}
    fmt.Println(v.Abs())
    array := make([]int,5)
    fmt.Println(array.Getsum())

}