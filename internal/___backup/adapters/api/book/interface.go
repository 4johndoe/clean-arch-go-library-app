package book

import (
	"ca-library-app/internal/___backup/domain/book"
	"ca-library-app/internal/domain/entity"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Book
	GetAll(ctx context.Context, limit, offset int) []*entity.Book
	Create(ctx context.Context, dto *book.CreateBookDto) *entity.Book
}
