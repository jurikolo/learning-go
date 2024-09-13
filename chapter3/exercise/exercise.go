package main

import "fmt"

func main() {
	type employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := employee{
		"bob", "davidov", 1,
	}

	fmt.Printf("e1: %v\n", e1)

	e2 := employee{
		firstName: "bob",
		lastName:  "davidov",
		id:        2,
	}

	fmt.Printf("e2: %v\n", e2)

	var e3 employee
	e3.firstName = "bob"
	e3.lastName = "davidov"
	e3.id = 3

	fmt.Printf("e3: %v\n", e3)
}
