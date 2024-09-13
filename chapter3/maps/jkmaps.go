package main

import "fmt"

func main() {
	// MAP
	fmt.Println("=====\nMAP\n=====")
	var nilMap map[string]int //always return 0, doesn't allow to write
	fmt.Println(nilMap)       //map[]

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams["Orcas"])    //[Fred Ralph Bijou]
	fmt.Println(teams["Orcas"][1]) //Ralph

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	mapOfInts := map[string]int{}
	mapOfInts["a"] = 2
	mapOfInts["b"] = 3
	fmt.Println(mapOfInts)
	mapOfStrings := map[string]string{}
	mapOfStrings["a"] = "a"
	mapOfStrings["b"] = "b"
	fmt.Println(mapOfStrings)

	// ,ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true
	v, ok = m["world"]
	fmt.Println(v, ok) // 0 true
	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 false

	// delete
	delete(m, "world")
	v, ok = m["world"]
	fmt.Println(v, ok) // 0 false

	fmt.Println(m, len(m)) // map[hello:5] 1
	clear(m)
	fmt.Println(m, len(m)) // map[] 0

	// use map as a set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// STRUCT
	fmt.Println("=====\nSTRUCT\n=====")
	type person struct {
		name string
		age  int
		pet  string
	}
	var oleg person
	fmt.Println(oleg)

	johnny := person{
		"Johnny",
		40,
		"dog",
	}
	fmt.Println(johnny)

	bobby := person{
		name: "Bobby",
		pet:  "frog",
	}
	fmt.Println(bobby)
}
