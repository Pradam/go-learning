package main

import "fmt"

type Pradam struct {
	i int
	b bool
}

func main() {
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q, len(q))

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r, len(r))

	s := []struct {
		i int
		b bool
	}{}
	for i := 0; i < len(q); i++ {
		for j := 0; j < len(r); j++ {
			l := Pradam{q[i], r[j]}
			s = append(s, l)
		}
	}
	fmt.Printf("%T %T %T\n", s, r, q)
	fmt.Println(s, len(s))
}
