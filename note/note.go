package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title       string
	description string
	createdAt   time.Time
}

func (note Note) Display() {

	fmt.Printf("Your note titled %v has the following content: \n\n%v", note.title, note.description)
}
func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title or content is empty")
	}

	return Note{
		title:       title,
		description: content,
		createdAt:   time.Now(),
	}, nil
}
