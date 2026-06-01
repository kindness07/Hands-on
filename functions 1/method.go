package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("I am", p.name, "and I am", p.age, "years old")
}

func main() {
	p1 := person{
		name: "James",
		age:  32,
	}
	p1.speak()
}
