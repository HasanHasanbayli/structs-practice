package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs-practice/note"
	"structs-practice/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	todoText := getUserInput("Todo text: ")
	title, content := getNotData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func getNotData() (string, string) {
	title := getUserInput("Note title: ")

	description := getUserInput("Note description: ")

	return title, description
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note successfully.")
	return nil
}

func outputData(data outputtable) error {

	data.Display()
	return saveData(data)
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r ")

	return text
}
