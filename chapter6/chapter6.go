package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15

	y := 10
	var pointerToY *int
	pointerToY = &y
	fmt.Println(pointerToY) // prints a memory address

	fmt.Println("\nModify variable by pointer")
	x2 := 10
	failedUpdate(&x2)
	fmt.Println(x2) // prints 10
	update(&x2)
	fmt.Println(x2) // prints 20
}

func failedUpdate(px *int) {
	x3 := 20
	px = &x3
}
func update(px *int) {
	*px = 20
}
