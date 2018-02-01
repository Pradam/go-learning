package main 

import "fmt"
import "math"
type Vertex struct{
    X, Y float64
}

func (a *Vertex) Abs() float64{
    return math.Sqrt(a.X+a.Y)
}


type MyArray []int

func (array MyArray) Getsum() []int{
    new_array := []int{} 
    for i:=0;i<len(array);i++{
        new_array = append(new_array, array[i]*2)
    }
    return new_array
}

func main(){
    v := &Vertex{3,4}
    fmt.Println(v.Abs())
    var array  =  []int{1,2,3,4,5,6,7,8,9}
    fmt.Println(MyArray.Getsum(array))

}