package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text   string `json:"content"`
}

func (t Todo) Display() {
	fmt.Printf("Your todo: %v", t.Text)
}

func New(text string) (Todo, error) {
	if text == ""  {
		return Todo{}, errors.New("title and content is required")
	}

	todo := Todo{
		Text: text,
	}
	return todo, nil
}

func (t Todo) Save() error{

   json, err := json.Marshal(t)
   if  err != nil {
	return err
   }

   return os.WriteFile("todo.json", json, 0644)
}
