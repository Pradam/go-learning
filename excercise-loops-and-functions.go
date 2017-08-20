package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	sum := x
	for sum < 10 {
		sum += 1
		newton := float64(sum - (sum*sum-x)/(2*sum))
		fmt.Println(newton)
	}
	return sum
}

func main() {
	fmt.Println(Sqrt(2))
}
