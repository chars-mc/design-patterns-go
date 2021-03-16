package list

import "fmt"

// Plain prints the plain list
// is a representation
type Plain struct{}

// NewPlain creates a NewPlain
func NewPlain() *Plain {
	return &Plain{}
}

// Print shows the content of the list
func (p *Plain) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", v)
	}
}
