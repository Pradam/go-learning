package main

import (
         "fmt"
         "math"
        )

func pradam_add_up(x,y int) (wonder int) {
   wonder = x + y
   return 
}

func list_square_root(x float64) []float64 {
	var pow float64 = 2
	mod_ := int(x)
	square_list := make([]float64, 6)
	for i := 0 ; i < mod_; i++ {
		val := math.Pow(float64(i), pow)
		square_list[i] = val
	}
	return square_list
}

func main(){
	a := 20000
	fmt.Printf("Hello World!!! %d", a)
	var b float64 = 30
	fmt.Println(math.Sqrt(b))
	fmt.Println(pradam_add_up(89,100))
	fmt.Println(list_square_root(6))
}
