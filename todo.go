package main

import "time"

type Todo struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

type TodoData struct {
	Name      string
	Completed bool
	DueSec    int64
}

func (t TodoData) toTodo(id uint32) Todo {
	return Todo{
		Id:        id,
		Name:      t.Name,
		Completed: t.Completed,
		Due:       time.Unix(t.DueSec, 0),
	}
}

func (t *TodoData) readFrom(from Todo) {
	t.Name = from.Name
	t.Completed = from.Completed
	t.DueSec = from.Due.Unix()
}
