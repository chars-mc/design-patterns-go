package main

import (
	"strconv"
	"strings"
)

// Task represents a task with subtasks
type Task struct {
	name        string
	responsable string
	price       int
	subTask     []Item
}

// Add push a new task
func (t *Task) Add(i Item) {
	t.subTask = append(t.subTask, i)
}

// String prints the task information
func (t *Task) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t|--")
	sb.WriteString(t.name)
	sb.WriteString(" - ")
	sb.WriteString(t.responsable)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(t.Price()))
	sb.WriteString("\n")

	for _, v := range t.subTask {
		sb.WriteString(v.String())
	}

	return sb.String()
}

// Price returns the sum of the subtasks prices
func (t *Task) Price() int {
	price := t.price
	for _, v := range t.subTask {
		price += v.Price()
	}

	return price
}
