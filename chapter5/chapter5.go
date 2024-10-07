package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("\n\nInvoke div function")
	fmt.Println(myDiv(10, 2))

	fmt.Println("\n\nNamed and optional parameters")
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	fmt.Println("\n\nVariadic input parameters")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	fmt.Println("\n\nMultiple returns")
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
	result2, _, _ := divAndRemainder(5, 2)
	fmt.Println(result2)

	fmt.Println("\n\nNamed returned values")
	result3, remainder3, _ := divAndRemainder2(5, 2)
	fmt.Println(result3, remainder3)
	result3, remainder3, _ = divAndRemainder3(5, 2)
	fmt.Println(result3, remainder3)

	fmt.Println("\n\nFunctions are values")
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result = myFuncVariable("Hello")
	fmt.Println(result)
	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)

	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	fmt.Println("\n\nAnonymous functions")
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

	fmt.Println("Same as above, but another style")
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}

	fmt.Println("\n\nClosures")
	a4 := 20
	f4 := func() {
		fmt.Println(a4)
		a4 = 30
	}
	f4()
	fmt.Println(a4)

	fmt.Println("\n\nPassing functions as parameters")
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	fmt.Println("\n\nReturning Functions from Functions")
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	defer fmt.Println("defer")
	fmt.Println("Exiting")
}

func myDiv(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	fmt.Println(opts)
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func divAndRemainder2(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func divAndRemainder3(num, denom int) (result int, remainder int, err error) {
	// assign some values
	result, remainder = 20, 30 // BOOM!!! These are never used, Go compiler uses values in a return statement
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func divAndRemainderBlankReturn(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

// functions as values
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

// end of functions as values
// passing functions as parameters
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// end passing functions as parameters
// returning functions from functions
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
