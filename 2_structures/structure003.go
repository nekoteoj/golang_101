package main

import (
	"fmt"
)

// Creating function on structure
type Saiyan struct {
	Name  string
	Power int
}

// Creating Super method with the reciever s *Saiyan
func (s *Saiyan) Super() {
	s.Power *= 100
}

// Struct factory <YAY, it's c style programming>
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{name, power}
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	goku.Super()
	fmt.Println(goku)

	vegeta := NewSaiyan("Vegeta", 10000)
	vegeta.Super()
	fmt.Println(vegeta)
}
