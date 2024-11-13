package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
  Text string `json:"text"`
}

func New(text string) Todo {
  return Todo {
    Text: text,
  }
}

func (todo Todo) Save() error {
  fileName := "todo.json"
  data, err := json.Marshal(todo)

  if err != nil {
    return err
  }

  return os.WriteFile(fileName, data, 0644)
}

func (todo Todo) Display() {
  fmt.Println(todo.Text)
}
