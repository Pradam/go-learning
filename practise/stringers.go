package main

import "fmt"

type Abser interface {
	M() string
}

type Ruckus string

func (r Ruckus) M() string {
	return string(r)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v '(%#v)'", p.Name, p.Age)
}

func main() {

	person := Person{"Abhil", 23}
	var value string = person.String()
	fmt.Printf("%v is added with town is %v", value, "Doddaballapur")
	var k Abser
	p := Ruckus("Pradam")
	k = p
	fmt.Println(k.M())

}
