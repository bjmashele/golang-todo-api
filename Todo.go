package main

import "time"

type Todo struct {
	Name      string    `json:"name`
	completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
