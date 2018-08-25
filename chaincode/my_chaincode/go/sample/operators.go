package main

import "fmt"

func main() {
	var x int = 45
	var y int = 20

	// Arithmetic Operators
	var m, n int = 100, 25
	var a, b string = "John", "Carry"
	var sumstring = a + " " + b
	fmt.Println(sumstring)
	c, d, e, f, g := m+n, m-n, m*n, m/n, m%n
	fmt.Println(c, d, e, f, g)

	//Bitwise Operators
	z := x & y
	fmt.Println(z)
	p := x | y
	fmt.Println(p)
	q := x ^ y
	fmt.Println(q)
	r := x &^ y
	fmt.Println(r)

	//Logical Operators
	if x != y && x <= y {
		fmt.Println("True")
	} else {
		fmt.Println("flase")
	}
	if x != y || x <= y {
		fmt.Println("True")
	} else {
		fmt.Println("flase")
	}
	if !(x == y) {
		fmt.Println("True")
	} else {
		fmt.Println("flase")
	}

}
