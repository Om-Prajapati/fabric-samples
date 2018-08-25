package main

import "fmt"

func main() {
	x := 5
	y := 6
	sum := x + y

	fmt.Println("Welcome to bridgelabz----> Start")
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println(sum)
		fmt.Println("abc")
	}
	c := abc(20, 10)
	fmt.Println(c)

	disString := str("hello", " Bridgelabz")
	fmt.Println(disString)
	fmt.Println("Welcome to bridgelabz----> End")
}
func abc(a int, b int) int {
	return a + b
}
func str(x string, y string) string {
	return x + y
}
