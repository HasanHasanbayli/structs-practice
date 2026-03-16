package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs-practice/note"
)

func main() {

	title, content := getNotData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

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
