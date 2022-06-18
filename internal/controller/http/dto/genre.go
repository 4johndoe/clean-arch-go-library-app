package dto

type CreateGenreDto struct {
	Name string `json:"name"`
}

type UpdateGenreDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
