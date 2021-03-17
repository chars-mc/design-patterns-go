package main

import (
	"fmt"
	"strconv"
	"strings"
)

// SubTask defines the task
type SubTask struct {
	name  string
	price int
}

// Add push a task
func (s *SubTask) Add(i Item) {
	fmt.Println("subtask are full")
}

// String ...
func (s *SubTask) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t\t|-- ")
	sb.WriteString(s.name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(s.Price()))
	sb.WriteString("\n")

	return sb.String()
}

// Price return the price of the task
func (s *SubTask) Price() int {
	return s.price
}
