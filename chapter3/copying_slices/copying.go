package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	z := make([]int, 2)
	num = copy(z, x)
	fmt.Println(z, num)

	a := make([]int, 2)
	copy(a, x[1:])
	fmt.Println(a)

	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	xSlice[1] = 42
	fmt.Println("xSlice:", xSlice)
	fmt.Println("xArray:", xArray)
	xSlice = append(xSlice, 9)
	xSlice[1] = 6
	fmt.Println("xSlice:", xSlice)
	fmt.Println("xArray:", xArray)
}
