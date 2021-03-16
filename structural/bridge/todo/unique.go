package todo

import "github.com/chars-mc/design-patterns-go/structural/bridge/list"

// Unique allows unique tasks
// is an implementation
type Unique struct {
	rendering list.List
	todos     []string
}

// NewUnique creates an new Unique
func NewUnique() *Unique {
	return &Unique{}
}

// SetList sets the list (the representation)
func (u *Unique) SetList(l list.List) {
	u.rendering = l
}

// Add push a new task to the list
func (u *Unique) Add(name string) {
	for _, v := range u.todos {
		if name == v {
			return
		}
	}
	u.todos = append(u.todos, name)
}

// Print shows the list (the representation)
func (u *Unique) Print() {
	u.rendering.Print(u.todos)
}
