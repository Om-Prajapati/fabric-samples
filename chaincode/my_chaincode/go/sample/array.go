package main

import "fmt"

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}

	defer fmt.Println(intArray[2])
	fmt.Println("------start-----")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	var x [5]int
	x[0] = 10
	x[4] = 20
	x[1] = 30
	x[3] = 40
	x[2] = 50
	fmt.Println("Values of Array X: ", x)

	y := [5]int{0: 100, 1: 200, 3: 500}
	fmt.Println("Values of Array Y: ", y)
	fmt.Println("-----end------")
}
