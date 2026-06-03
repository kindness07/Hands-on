package main

import "fmt"

type dog struct {
	first string
}

// type grownup struct {
// 	age int
// 	dog
// }

// func (g grownup) grown() {
// 	fmt.Println("My name is", g.first, "and I'm", g.age, "years old.")

// }

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d dog) treats() {
	fmt.Println("My name is", d.first, "and I love treats.")
}

type youngin interface {
	walk()
	run()
	treats()
	// grown()
}

func youngRun(y youngin) {
	y.run()
	// y.grown()
	y.treats()
	y.walk()
}

func main() {
	d1 := dog{"Henry"}

	youngRun(d1)

	d2 := dog{"Padget"}

	youngRun(d2)

	// d3 := grownup{age: 42, dog: dog{first: "Rover"}}

	// youngRun(d3)

}
