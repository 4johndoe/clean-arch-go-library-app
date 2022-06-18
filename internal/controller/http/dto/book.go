package dto

type CreateBookDto struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID string `json:"author_id"`
	GenreID  string `json:"genre_id"`
}

type UpdateBookDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorID string `json:"author_id"`
	Busy     bool   `json:"busy"`
	Owner    string `json:"owner"`
}
