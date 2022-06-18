package service

import (
	"ca-library-app/internal/domain/entity"
	"context"
)

type BookStorage interface {
	GetByID(ctx context.Context, id string) entity.Book
	GetAll(ctx context.Context, limit, offset int) []entity.Book
	//Create(ctx context.Context, author2 *entity.Author) *entity.Author
}

type bookService struct {
	storage BookStorage
}

func NewBookService(storage BookStorage) *bookService {
	return &bookService{storage: storage}
}

//func (s *bookService) Create(ctx context.Context) *entity.Book {
//	//author := s.authorService.GetByID(ctx, dto.AuthorID) check if exists
//	return nil
//}

func (s bookService) GetByID(ctx context.Context, id string) entity.Book {
	return s.storage.GetByID(ctx, id)
}

func (s bookService) GetAll(ctx context.Context, limit, offset int) []entity.Book {
	return s.storage.GetAll(ctx, limit, offset)
}

func (s bookService) GetAllForList(ctx context.Context, limit, offset int) []entity.BookView {
	// TODO implement
	return nil
}
