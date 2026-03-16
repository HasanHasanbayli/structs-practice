package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs-practice/note"
	"structs-practice/todo"
)

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

	userTodo.Display()
	err = userTodo.Save()
	if err != nil {
		fmt.Println("Saving todo failed: ", err)
	}
	fmt.Println("Saving the todo successfully")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed: ", err)
	}

	fmt.Println("Saving the note successfully")
}

func getNotData() (string, string) {
	title := getUserInput("Note title: ")

	description := getUserInput("Note description: ")

	return title, description
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
