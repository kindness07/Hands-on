package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "Model 3",
		doors: 4,
		color: "red",
	}
	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Ford",
		model: "Mustang",
		doors: 2,
		color: "blue",
	}

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%#v\n", v2)

	fmt.Println(v1.electric)
	fmt.Println(v2.electric)

}
