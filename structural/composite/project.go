package main

import (
	"strconv"
	"strings"
)

// Project represents a project with it's tasks
type Project struct {
	name  string
	tasks []Item
}

// Add push a new task
func (p *Project) Add(i Item) {
	p.tasks = append(p.tasks, i)
}

// String prints the name of the project and it's total price
// and the tasks with their prices
func (p *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(p.name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(p.Price()))
	sb.WriteString("\n")
	for _, v := range p.tasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

// Price returns the sum of the tasks prices
func (p *Project) Price() int {
	price := 0
	for _, v := range p.tasks {
		price += v.Price()
	}

	return price
}
