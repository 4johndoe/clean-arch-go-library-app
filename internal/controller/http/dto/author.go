package dto

type CreateAuthorDto struct {
	Name string `json:"name"`
	Age  int    `json:"year"`
}

type UpdateAuthorDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"year"`
}
