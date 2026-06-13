package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("My name is", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Kind"}
	saySomething(&p1) // you can pass type person into saySomething
	p1.speak()        // you can pass type person into speak
	// saySomething(p1)  // you cannot pass type person into saySomething because it does not satisfy the human interface, but *person does
}
