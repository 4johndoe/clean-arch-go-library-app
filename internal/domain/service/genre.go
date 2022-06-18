package service

import (
	"ca-library-app/internal/domain/entity"
	"context"
)

type GenreStorage interface {
	GetByUUID(ctx context.Context, id string) *entity.Genre
	GetAll(ctx context.Context, limit, offset int) []*entity.Genre
	//Create(ctx context.Context, author2 *entity.Author) *entity.Author
}

type genreService struct {
	storage GenreStorage
}

func NewGenreService(storage GenreStorage) *genreService {
	return &genreService{storage: storage}
}

//func (s *genreService) Create(ctx context.Context) *entity.Genre {
//	//author := s.authorService.GetByUUID(ctx, dto.AuthorUUID) check if exists
//	return nil
//}

func (s *genreService) GetByUUID(ctx context.Context, id string) *entity.Genre {
	return s.storage.GetByUUID(ctx, id)
}

func (s *genreService) GetAll(ctx context.Context, limit, offset int) []*entity.Genre {
	return s.storage.GetAll(ctx, limit, offset)
}
