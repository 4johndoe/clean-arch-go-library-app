package user

type User struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
}
