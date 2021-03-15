package main

import (
	"fmt"
)

func main() {
	color := "White"
	phones := map[string]string{"home": "98765", "work": "56789"}
	u1 := user{
		"Carlos",
		21,
		[]string{"Friend 1", "Friend 2", "Friend 3"},
		&color,
		phones,
	}
	// copy of the original object
	u2 := u1.clone() // client

	// changing some values
	u2.age = 20
	u2.name = "Manuel"
	u2.friends[1] = "Friend 5"
	color = "Blue"
	phones["home"] = "12345"

	// showing the information
	fmt.Println(u1.string())
	fmt.Println(u2.string())
}
