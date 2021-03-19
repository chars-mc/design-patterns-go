package decorator

import "fmt"

// Route contains a map with Decorators
type Route struct {
	decorators map[string]Decorator
}

// NewRoute creates a new Route
func NewRoute() Route {
	return Route{make(map[string]Decorator)}
}

// Add push a new Decorator on the decorators map with the path as key
func (r *Route) Add(decorator Decorator, path string) {
	r.decorators[path] = decorator
}

// Exec executes the
func (r *Route) Exec(path string) {
	d, ok := r.decorators[path]
	if !ok {
		fmt.Println("the route does not exist")
	}

	err := d.Process()
	if err != nil {
		fmt.Println("EROR: ", err)
	}
}
