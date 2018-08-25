package main

import "fmt"

func main() {
	fmt.Println(add(20, 30))
	multi(20, 30)
}

// function with return type
func add(x int, y int) int {
	return x + y
}

// funtion without return type
func multi(x, y int) {
	var z = x * y
	fmt.Println(z)
}
