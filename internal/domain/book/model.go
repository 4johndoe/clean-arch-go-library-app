package book

import (
	"ca-library-app/internal/domain/author"
	"fmt"
)

type Book struct {
	UUID   string        `json:"uuid"`
	Name   string        `json:"name"`
	Year   int           `json:"year"`
	Author author.Author `json:"author"`
	Busy   bool          `json:"busy"`
	Owner  string        `json:"owner"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
