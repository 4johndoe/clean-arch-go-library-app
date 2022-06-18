package service

import (
	"ca-library-app/internal/domain/entity"
	"context"
)

type BookStorage interface {
	GetByUUID(ctx context.Context, id string) *entity.Book
	GetAll(ctx context.Context, limit, offset int) []*entity.Book
	//Create(ctx context.Context, author2 *entity.Author) *entity.Author
}

type bookService struct {
	storage BookStorage
}

func NewBookService(storage BookStorage) *bookService {
	return &bookService{storage: storage}
}

//func (s *bookService) Create(ctx context.Context) *entity.Book {
//	//author := s.authorService.GetByUUID(ctx, dto.AuthorUUID) check if exists
//	return nil
//}

func (s *bookService) GetByUUID(ctx context.Context, id string) *entity.Book {
	return s.storage.GetByUUID(ctx, id)
}

func (s *bookService) GetAll(ctx context.Context, limit, offset int) []*entity.Book {
	return s.storage.GetAll(ctx, limit, offset)
}
