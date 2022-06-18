package dto

type BookEvent struct {
	EventType string
	Data      struct {
		Name     string `bson:"name"`
		Year     int    `bson:"year"`
		AuthorID string `bson:"author_id"`
		GenreID  string `bson:"genre_id"`
	}
}
