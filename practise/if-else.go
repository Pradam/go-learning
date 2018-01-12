package main


import (
      "fmt"
      "math"
)


func pow(x int) string {
	if v:=x; v < 100 {
		return "less than 100"
	} else if v == 100{
    return "equals to 100"
  }
	return "not less than 100"
}


func main(){
  var get_sqrt = func(x float64) float64{
       return math.Sqrt(x)
  }
  fmt.Println(pow(100))
  for i:=0; i < 100; i++{
  	fmt.Printf("%d === %g",i,get_sqrt(float64(i)))
  	fmt.Println()
  }
}