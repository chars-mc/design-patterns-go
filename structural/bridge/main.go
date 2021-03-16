package main

import (
	"github.com/chars-mc/design-patterns-go/structural/bridge/list"
	"github.com/chars-mc/design-patterns-go/structural/bridge/todo"
)

func main() {
	uniqueToDos := factoryTodo("any")
	rendering := factoryList("bullet")

	uniqueToDos.SetList(rendering)
	uniqueToDos.Add("first task")
	uniqueToDos.Add("second task")
	uniqueToDos.Add("second task")
	uniqueToDos.Print()

}

// factoryTodo returns a ToDo
func factoryTodo(s string) todo.ToDo {
	if s == "unique" {
		return todo.NewUnique()
	}
	return todo.NewAny()
}

// factoryList
func factoryList(s string) list.List {
	if s == "plain" {
		return list.NewPlain()
	}
	return list.NewBullet('*')
}
