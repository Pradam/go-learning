package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var d data
	d.Num = 10
	d.Count = 20
	d.Max = 30
	d.Test = 100

	js, err := json.Marshal(d)
	fmt.Println("Hello")
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println(string(js))
}

type Num int
type Count int
type Max int
type Test int

type data struct {
	Num
	Count
	Max
	Test
}
