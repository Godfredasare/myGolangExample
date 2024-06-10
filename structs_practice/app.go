package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/max_go_course/structs_practice/note"
	"example.com/max_go_course/structs_practice/todo"
)

type saver interface {
	Save() error
}

type outPuttable interface {
	saver
	Display()
}

func outPuttData(data outPuttable) error {
	data.Display()

	return saveDate(data)
}

func saveDate(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving the file")
		return err
	}
	fmt.Println(" file saved")
	return nil
}

func main() {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")
	text := getUserInput("Enter todo: ")

	note, err := note.New(title, content)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	todo, err := todo.New(text)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	err = outPuttData(note)
	if err != nil {
		return
	}

	err = outPuttData(todo)
	if err != nil {
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
