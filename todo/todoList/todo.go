package todolist

import (
	"errors"
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

func New(id int64, task, description string) (*ToDo, error) {
	if id == 0 || task == "" || description == "" {
         return nil, errors.New("id, task and description is required")
	}

	todo := &ToDo{
		ID:          id,
		Task:        task,
		Description: description,
		CreatedAt:   time.Now(),
	}

	return todo, nil
}

func Delete(id int64, todo []ToDo) ([]ToDo, error){
	var newTodos []ToDo

	for i, v := range todo {
		if v.ID == id {
          newTodos = append(newTodos[:i], newTodos[i+1:]... )
		} else if v.ID != id {
            return nil, errors.New("Todo does don't exist")
		}
	}
	return newTodos, nil
}
