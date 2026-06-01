package main

import "fmt"

type person struct {
	first          string
	friends        map[string]int
	favoriteDrinks []string
}

func main() {
	p1 := struct {
		first          string
		friends        map[string]int
		favoriteDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny": 5,
			"John":  3,
		},
		favoriteDrinks: []string{"Martini", "Whiskey"},
	}

	p2 := person{
		first: "Jenny",
		friends: map[string]int{
			"James": 5,
			"John":  3,
		},
		favoriteDrinks: []string{"Martini", "Whiskey"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends -", k, v)
	}

	for _, v := range p2.favoriteDrinks {
		fmt.Println(p2.first, "- favorite drinks -", v)
	}
}
