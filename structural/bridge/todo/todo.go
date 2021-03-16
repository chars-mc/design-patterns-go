package todo

import "github.com/chars-mc/design-patterns-go/structural/bridge/list"

// ToDo contains the methods of a list
// implemetations will implement this interface
type ToDo interface {
	SetList(l list.List) // l is the representation
	Add(name string)
	Print()
}
