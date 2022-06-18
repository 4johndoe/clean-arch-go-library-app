package author

type CreateAuthorDto struct {
	Name string `json:"name"`
	Age  int    `json:"year"`
}

type UpdateAuthorDto struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"year"`
}
