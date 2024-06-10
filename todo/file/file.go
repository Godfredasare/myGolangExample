package file

import (
	"encoding/json"
	"os"

	"example.com/max_go_course/todo/todoList"
)

const path = "./todo/todo.json"

func  Save(todo []todolist.ToDo) error{
	file, err := os.Create(path)
	if err !=nil {
       return err
	}

	defer file.Close()

	json, err := json.MarshalIndent(todo, "", "")
	if err !=nil {
		return err
	 }

	_, err =file.WriteString(string(json))
	if err !=nil {
		return err
	 }
	 return nil
}

func Read() ( []todolist.ToDo, error){
    file, err := os.ReadFile(path)
	if err !=nil {
		if os.IsNotExist(err){
			return []todolist.ToDo{}, nil
		}
		return  []todolist.ToDo{}, err
	 }

	 var todo []todolist.ToDo

	err = json.Unmarshal(file, &todo)
	if err !=nil {
		return []todolist.ToDo{}, err
	 }

	 return todo, nil
}