package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {
	p1 := person{
		first:            "James",
		last:             "Bond",
		favoriteIceCream: []string{"Chocolate", "Vanilla"},
	}

	p2 := person{
		first:            "Jenny",
		last:             "Moneypenny",
		favoriteIceCream: []string{"Strawberry", "Mint"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favoriteIceCream {
		fmt.Printf("%d - %s\n", i, v)
	}
	for i, v := range p2.favoriteIceCream {
		fmt.Printf("%d - %s\n", i, v)

	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	// for k, v := range m {
	// 	fmt.Printf("key: %s value: %v\n", k, v)
	// 	for i, val := range v.favoriteIceCream {
	// 		fmt.Printf("\tindex position: %d \t value: %s\n", i, val)
	// 	}
	// }

	for _, v := range m {
		fmt.Println(v)
		for _, v1 := range v.favoriteIceCream {
			fmt.Println(v.first, v.last, v1)
		}
	}
}
