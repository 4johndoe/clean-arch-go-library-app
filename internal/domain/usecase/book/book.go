package book

import (
	"ca-library-app/internal/domain/entity"
	"context"
)

type Service interface {
	GetAllForList(ctx context.Context) []entity.BookView
	GetByID(ctx context.Context, id string) entity.Book
}

type AuthorService interface {
	GetByID(ctx context.Context, id string) entity.Author
}

type GenreService interface {
	GetByID(ctx context.Context, id string) entity.Genre
}

type bookUseCase struct {
	bookService   Service
	authorService AuthorService
	genreService  GenreService
}

func (u bookUseCase) ListAllBooks(ctx context.Context) []entity.BookView {
	// список книг с именем жанра и именем автора
	return u.bookService.GetAllForList(ctx)
}

func (u bookUseCase) GetFullBook(ctx context.Context, id string) entity.FullBook {
	book := u.bookService.GetByID(ctx, id)
	genre := u.genreService.GetByID(ctx, book.GenreID)
	author := u.authorService.GetByID(ctx, book.AuthorID)

	return entity.FullBook{
		Book:   book,
		Genre:  genre,
		Author: author,
	}
}
