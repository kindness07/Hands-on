package main

import (
	"fmt"
)

type person struct {
	first string
}

func Changename(p person, newname string) person {
	p.first = newname
	return p
}
func Changenamep(p *person, newname string) {
	p.first = newname

}

func main() {
	p1 := person{first: "kind"}
	fmt.Println(p1.first)
	p1 = Changename(p1, "ijeoma")
	fmt.Println(p1.first)

	Changenamep(&p1, "okezie")
	fmt.Println(p1.first)
}
