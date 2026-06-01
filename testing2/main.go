package main

import "fmt"

func main() {
	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Subtract)
	fmt.Printf("%T\n", DoMath)
	x := DoMath(42, 16, Add)
	fmt.Println(x)
	y := DoMath(42, 16, Subtract)
	fmt.Println(y)
	fmt.Println(paradise("Hawaii"))
}

func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}

func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
