package main

import ("fmt")

type Vertex struct {
	X int
	Y int
}


type Pradam struct{
	name string
    age int
    qualities  []string
}

func main(){
	var k = 100
	v := Vertex{1, 2}
	v.X= 900
	fmt.Println(&k, &v)
	w := Vertex{1,2}
	x := *&Vertex{1,2}
	y := Vertex{X: 1}
	z := Vertex{}
	all_vertex := []Vertex{w, x, y, z}
	fmt.Println(all_vertex)
	for i:=0; i < len(all_vertex);i++{
		val := all_vertex[i]
		val.X = 700
		fmt.Println(val.X, val.Y)
	}
	skills := []string{"coding", "learning"}
	pradam := Pradam{"abhilash", 24, skills}
	fmt.Println(pradam.name)
	fmt.Println(pradam.age)
	fmt.Println(pradam.qualities[0])
	fmt.Println(v.X, w.X)
}