package book

type CreateBookDto struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
}

type UpdateBookDto struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
	Busy       bool   `json:"busy"`
	Owner      string `json:"owner"`
}
