package service

import (
	"ca-library-app/internal/domain/entity"
	"context"
)

type AuthorStorage interface {
	GetByUUID(ctx context.Context, id string) *entity.Author
	GetAll(ctx context.Context, limit, offset int) []*entity.Author
	//Create(ctx context.Context, author2 *entity.Author) *entity.Author
}

type authorService struct {
	storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) *authorService {
	return &authorService{storage: storage}
}

//func (s *authorService) Create(ctx context.Context) *entity.Author {
//	return nil
//}

func (s *authorService) GetAll(ctx context.Context, limit, offset int) []*entity.Author {
	return s.storage.GetAll(ctx, limit, offset)
}

func (s *authorService) GetByUUID(ctx context.Context, id string) *entity.Author {
	return s.storage.GetByUUID(ctx, id)
}
