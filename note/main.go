package main

import (
	"bufio"
	"errors"
	"fmt"
	"golearn/note/note"
	"golearn/note/todo"
	"os"
	"strings"
)

type saver interface {
  Display()
  Save() error
}


func main() {

  // note
	title, err := getUserInput("Note title: ")

	if err != nil {
		return
	}

	content, err := getUserInput("Note content: ")

	if err != nil {
		return
	}

	note := note.New(title, content)
	note.Display()
  err = saveData(note)

  if err != nil {
    return
  }

  // todo
  input, err := getUserInput("Todo: ")

  if err != nil {
    return
  }

  todo := todo.New(input)
  todo.Display()
  err = saveData(todo)

  if err != nil {
    return
  }
}

func createTodo() {
}

func createNote() {
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
    fmt.Println(err)
		return "", err
	}

	if input == "" {
    err = errors.New("You must submit a non-empty input.")
    fmt.Println(err)
		return "", err
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input, nil
}

func saveData(data saver) error {
  err := data.Save()

  if err != nil {
    fmt.Printf("An error occurred while saving your data: %v\n", err)
    return err
  }

  fmt.Println("Your data was saved successfully.")
  return nil
}
