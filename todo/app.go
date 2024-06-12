package main

import (		
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Godfredasare/myGolangExample.git/todo/file"
	todolist "github.com/Godfredasare/myGolangExample.git/todo/todoList"
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

	readTodo, err := file.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Simple TODO App.")
	fmt.Println("________________")
	fmt.Println("")

	fmt.Println("1: Add ToDo.")
	fmt.Println("2: View All ToDo.")
	fmt.Println("3: Delete ToDo By Id.")
	fmt.Println("0: Exit.")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)
	fmt.Println("")

	if choice == 1 {
		id := input("Enter id: ")
		ID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		task := input("Enter Task: ")
		description := input("Enter Description: ")

		todo, err := todolist.New(ID, task, description)
		if err != nil {
			fmt.Println(err)
			return
		}
		todo.Print()

		readTodo = append(readTodo, *todo)

		err = file.Save(readTodo)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else if choice == 2 {
		for i, todo := range readTodo {
			fmt.Printf("Task %v: %v\n", i+1, todo.Task)
			fmt.Println("Description: ", todo.Description)
			fmt.Print("Created_At: ", todo.CreatedAt, "\n\n")
		}

	} else if choice == 3 {
		var id int64

		fmt.Print("Enter id to delete ToDo: ")
		fmt.Scan(&id)

		newTodo, err := todolist.Delete(id, readTodo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Item with id: %v deleted successfully", id)

		err = file.Save(newTodo)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else if choice == 0 {
		return
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
