package main

import (
	"fmt"
)

func main() {
	a, b := f4(5, 4)
	fmt.Println("a =", a, "b =", b)
	// Some returned value can be discard by using _
	_, c := f4(4, 5)
	fmt.Println("c =", c)
}

// function can be return many value

// simple function with 1 parameter
func f1(s string) {
	fmt.Println("You input :", s)
}

//function with return type
func f2(a int, b int) int {
	return a + b
}

// same type argument can be shorten
func f3(a, b int) int {
	return a + b
}

// return multiple values
func f4(a, b int) (int, bool) {
	return a + b, a > b
}
