package main

// Item interface that objects must implement to proccess the information
type Item interface {
	Add(Item)
	String() string
	Price() int
}
