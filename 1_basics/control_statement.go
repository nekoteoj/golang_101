package main

import (
	"fmt"
)

func main() {
	// in go while is for
	a := 1024
	for a > 0 {
		fmt.Print(a, " ")
		a /= 2
	}
	fmt.Println()

	// Normal for loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i*2, " ")
	}
	fmt.Println()

	// if else example
	k := 500
	if k > 600 {
		fmt.Println("YAY")
	} else if k > 400 {
		fmt.Println("foo")
	} else {
		fmt.Println("NANI!!!")
	}

	// no need for break in switch case
	name := "Joe"
	switch name {
	case "Joe":
		fmt.Println("You are")
	case "Jack":
		fmt.Println("YAY")
	default:
		fmt.Println("Annonymous")
	}
}
