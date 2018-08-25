package main

import "fmt"

var xyz int 
var abc string

// // we can not declear variable like this a:= 10 in outside of function
// // a:=10
// // b:=20
// func main(){
// 	var x int
// 	x=10
// 	fmt.Println(x)
	
// 	var y int = 20
// 	fmt.Println(y)
	
// 	z:=30
// 	fmt.Println(z)
	
// 	var a,b,c,d=10,20,30,40 
// 	fmt.Println(a,b,c,d)

// 	//we can declearl variable like this
// 	var _abc string 
// 	_abc = "Welcome to bridgelabz"
// 	fmt.Println(_abc)
// }


type rect struct {
    width int
    height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
}