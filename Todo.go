package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
