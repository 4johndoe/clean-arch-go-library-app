package entity

import "fmt"

type BookView struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorName string `json:"author"`
	GenreName  string `json:"genre"`
	Busy       bool   `json:"busy"`
}

type FullBook struct {
	Book
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID string `json:"author_id"`
	GenreID  string `json:"genre_id"`
	Busy     bool   `json:"busy"`
	Owner    string `json:"owner"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
