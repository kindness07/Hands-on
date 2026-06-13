package main

import (
	"fmt"
)

type customErr struct {
	Errmsg string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("MAYDAY: %v", ce.Errmsg)

}

func main() {
	c1 := customErr{
		Errmsg: "There's an error",
	}
	foo(c1)
}

func foo(e error) {

	fmt.Println(e)
}
