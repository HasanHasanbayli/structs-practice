package note

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (note Note) Display() {

	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.Title, note.Description)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json2.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title or content is empty")
	}

	return Note{
		Title:       title,
		Description: content,
		CreatedAt:   time.Now(),
	}, nil
}
