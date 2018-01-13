package main

import (
	"fmt"
)

func main() {
	// Decalre Variable
	var a int // Assign 0 as default value
	a = 10
	fmt.Println("a =", a)

	// Declare with value
	var b int = 20
	fmt.Printf("b = %d\n", b)

	// Omitting type
	var c = 7
	fmt.Println("c =", c)

	// Shorthand style using := to declare variable
	d := 8
	fmt.Println("d =", d)
	e := getE()
	fmt.Println("e =", e)

	f, g := 8, 9
	// Can also use =
	f, g = 10, 11
	fmt.Println("f + g =", f+g)

	// If any variable is new := can be used
	g, h := 20, 30
	fmt.Println("g * h =", g*h)

	/* conclusion
	1. use var [name] [type] to decalre to default value
	2. use [name] := value to declare with value
	3. use [name] = value to assign value to declared variable
	*/
}

func getE() int {
	return 100
}
