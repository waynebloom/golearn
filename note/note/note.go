package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
  Title string `json:"title"`
  Content string `json:"content"`
  CreatedAt time.Time `json:"createdAt"`
}

func New(title string, content string) Note {
  return Note {
    Title: title,
    Content: content,
    CreatedAt: time.Now(),
  }
}

func (note Note) Save() error {
  fileName := strings.ReplaceAll(note.Title, " ", "_")
  fileName = strings.ToLower(fileName)
  fileName = fileName + ".json"

  data, err := json.Marshal(note)

  if err != nil {
    return err
  }

  return os.WriteFile(fileName, data, 0644)
}

func (note Note) Display() {
  fmt.Printf("%v\n--------\n%v\n", note.Title, note.Content)
}
