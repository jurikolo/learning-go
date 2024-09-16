package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func main() {
	// Shadowing variables
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5         // creates shadow variable
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	fmt.Println(true)
	true := "xyz"
	fmt.Println(true) // xyz

	// if
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// for
	fmt.Println("\n\nfor complete loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n\nfor condition-only statemet")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i * 2
	}

	fmt.Println("\n\nfor infinite statement")
	// commented out to avoid program to print endless Hello message
	// for {
	// fmt.Println("Hello")
	// }

	fmt.Println("\n\nfor range")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	fmt.Println("\n\nContinue statement")
	// fizzbuzz
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n\nIterating over map")
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	fmt.Println("\n\nIterating over strings")
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	fmt.Print("\n\nfor range copy")
	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		v *= 2
	}
	fmt.Println(evenVals2)

	fmt.Println("\n\nfor loop exit to outer loop")
	samples2 := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples2 {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

	fmt.Println("\n\nSwitch statement")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	fmt.Println("\n\nGoTo")
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

	fmt.Println("\n\nExercise")
	var mySlice = make([]int, 100)
	for i := 0; i < 100; i++ {
		mySlice[i] = rand.Intn(1000)
		switch {
		case mySlice[i]%2 == 0 && mySlice[i]%3 == 0:
			fmt.Println("Six!")
		case mySlice[i]%2 == 0:
			fmt.Println("Two!")
		case mySlice[i]%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
	fmt.Println(cap(mySlice))
}
