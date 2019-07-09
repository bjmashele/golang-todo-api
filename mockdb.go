package main

import "fmt"

var currentId int

var todos Todos

// Seed data
func init() {
	MockCreateTodo(Todo{Name: "Go shopping"})
	MockCreateTodo(Todo{Name: "Run 5 miles"})
}

func MockFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func MockCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func MockDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
