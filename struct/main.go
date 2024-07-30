package main

import "fmt"

type Pesron struct {
	Fname  string
	lname  string
	gender string
	age    int64
}

func main() {
	person1 := Pesron{Fname: "Medina", lname: "Nesro", age: 22, gender: "female"}
	fmt.Println(person1)

}