package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	// Simplest way to create value of structure
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	// we don't have to set all field
	vegeta := Saiyan{}
	vegeta.Name = "Vegeta"
	vegeta.Power = 15000
	// or
	gohan := Saiyan{Power: 500}
	gohan.Name = "Gohan"
	// create with omitted field name
	trunk := Saiyan{"Trunk", 300}

	fmt.Println(goku, vegeta, gohan, trunk)
}
