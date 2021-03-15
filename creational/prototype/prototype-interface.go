package main

// user interface that the clonable object must implement
type userPrototype interface {
	string() string
	clone() *user
}
