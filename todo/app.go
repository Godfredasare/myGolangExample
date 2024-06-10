package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/max_go_course/todo/file"
	todolist "example.com/max_go_course/todo/todoList"
)

func input(prompt string) string {
	value, err := inputTodo(prompt)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

func main() {
	id := input("Enter id: ")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	task := input("Enter Task: ")
	description := input("Enter Description: ")

	todo := todolist.New(ID, task, description)
	todo.Print()

	readTodo, err := file.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	readTodo = append(readTodo, *todo)

	err = file.Save(readTodo)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, todo := range readTodo {
		fmt.Println(todo)
	}

}

func inputTodo(prompt string) (string, error) {
	fmt.Print(prompt)

	// var input string
	value := bufio.NewReader(os.Stdin)
	v, err := value.ReadString('\n')
	if err != nil {
		return "", err
	}
	v = strings.TrimSpace(v)
	return v, nil
}
