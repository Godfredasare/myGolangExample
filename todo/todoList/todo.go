package todolist

import (
	"fmt"
	"time"
)

type ToDo struct {
	ID          int64     `json:"id"`
	Task        string    `json:"task"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (todo ToDo) Print() {
	fmt.Println("New Task Added")
	fmt.Print("Task: ", todo.Task, "\n")
	fmt.Print("Description: ", todo.Description, "\n")
	fmt.Print("Created At: ", todo.CreatedAt, "\n")
}

func New(id int64, task, description string) *ToDo {

	todo := &ToDo{
		ID:          id,
		Task:        task,
		Description: description,
		CreatedAt:   time.Now(),
	}

	return todo
}
