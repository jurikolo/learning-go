package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	fmt.Println("x:", x)
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	y = append(y, "r")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

}
