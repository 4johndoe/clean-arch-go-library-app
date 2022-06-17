package book

type CreateBookDto struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author string `json:"author"`
}

type UpdateBookDto struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author string `json:"author"`
	Busy   bool   `json:"busy"`
	Owner  string `json:"owner"`
}
