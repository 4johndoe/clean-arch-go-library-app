package main

import (
	author3 "ca-library-app/internal/adapters/api/author"
	book3 "ca-library-app/internal/adapters/api/book"
	"ca-library-app/internal/adapters/db/author"
	"ca-library-app/internal/adapters/db/book"
	author2 "ca-library-app/internal/domain/author"
	book2 "ca-library-app/internal/domain/book"
)

func main() {
	// entry point
	bookStorage := book.NewStorage()
	bookService := book2.NewService(bookStorage)
	book3.NewHandler(bookService)

	authorStorage := author.NewStorage()
	authorService := author2.NewService(authorStorage)
	author3.NewHandler(authorService)

	// register in router
}
