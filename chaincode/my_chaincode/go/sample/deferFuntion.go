package main

import "fmt"

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func main() {
	defer second()
	fmt.Println("Welcome to bridgelabz start")
	first()
	fmt.Println("Welcome to bridgelabz end")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	xyz := 1
	for xyz < 5 {
		fmt.Println("xyz inside loop-->", xyz)
		xyz++
	}
	fmt.Println("xyz-->", xyz)

	variadicExample("abc", "om", "bridge", "it")
}

func variadicExample(s ...string) {
	defer fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[3])
}
