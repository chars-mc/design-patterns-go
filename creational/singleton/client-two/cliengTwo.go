package client_two

import "github.com/chars-mc/design-patterns-go/creational/singleton"

// SetAge set the age of person
func SetAge(a int) {
	p := singleton.GetInstance()
	p.SetAge(a)
}

// IncrementAge add 1 to age of person
func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
