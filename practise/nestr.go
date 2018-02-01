package main

import "fmt"

type Registration struct {
	Name     string
	Age      string
	Email    string
	Password string
}

type User struct {
	Registration
	Email string
}

func main() {
	user := User{}
	user.Age = "29"

	user.Name = "pradam"
	fmt.Println(user)

}
