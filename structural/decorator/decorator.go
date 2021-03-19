package decorator

// Decorator interface to decorate structs
type Decorator interface {
	Process() error
}
