package main

import (
	"fmt"
)

type A struct {
	k int
}

func f1(a A) {
	a.k++
}

func f2(a *A) {
	a.k++
}

func main() {
	n := A{5}
	// Pass by value
	f1(n)
	fmt.Println(n)
	// Passing pointer
	f2(&n)
	fmt.Println(n)

	// Another way to get pointer
	m := &A{7}
	fmt.Println(m)
}
