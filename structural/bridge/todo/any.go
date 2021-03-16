package todo

import "github.com/chars-mc/design-patterns-go/structural/bridge/list"

// Any allows duplicate tasks
// is an implementation
type Any struct {
	rendering list.List // rendering is the representation
	todos     []string
}

// NewAny creates a new Any
func NewAny() *Any {
	return &Any{}
}

// SetList sets the list (the representation)
func (a *Any) SetList(l list.List) {
	a.rendering = l
}

// Add push a task to the list
func (a *Any) Add(name string) {
	a.todos = append(a.todos, name)
}

// Print shows the list (the representation)
func (a *Any) Print() {
	a.rendering.Print(a.todos)
}
