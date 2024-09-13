package main

import (
	"fmt"
	"math"
)

const value = 0

func main() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)

	i = value
	f = value

	fmt.Println(i)
	fmt.Println(f)

	var b byte = ^byte(0)
	var smallI int32 = math.MaxInt32
	var bigI uint64 = ^uint64(0)

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)	
}
