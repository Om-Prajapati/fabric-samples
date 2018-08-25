package main

import ("fmt"
		"log")

type Employee struct {
	name string
	age int
	salary int 
}

func main() {
	log.Println("main started")
	values := Employee{"Om",20 ,10000}
	fmt.Println(values)
	
    log.Fatalln("fatal message")
 
    log.Panicln("panic message")
}


