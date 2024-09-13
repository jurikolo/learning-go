package main

import "fmt"

func main() {
	var s string = "Hello, Олег"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs) //[72 101 108 108 111 44 32 208 158 208 187 208 181 208 179]
	fmt.Println(rs) //[72 101 108 108 111 44 32 1054 1083 1077 1075]
}
