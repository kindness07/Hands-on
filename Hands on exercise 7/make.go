package main

import (
	"fmt"
)

func main() {
	xi := make([]string, 0, 50)

	xi = append(xi, "alabama", "alaska", "arizona", "arkansas", "california", "colorado", "connecticut", "delaware", "florida", "georgia", "hawaii", "idaho", "illinois", "indiana", "iowa", "kansas", "kentucky", "louisiana", "maine", "maryland", "massachusetts", "michigan", "minnesota", "mississippi", "missouri", "montana", "nebraska", "nevada", "new hampshire", "new jersey", "new mexico", "new york", "north carolina", "north dakota", "ohio", "oklahoma", "oregon", "pennsylvania", "rhode island", "south carolina", "south dakota", "tennessee", "texas", "utah", "vermont", "virginia", "washington", "west virginia", "wisconsin", "wyoming")
	for i := 0; i < len(xi); i++ {
		fmt.Printf(" %d - %s\n", i, xi[i])
	}

	fmt.Println(len(xi))
	fmt.Println(cap(xi))

}
