package author

import (
	"ca-library-app/internal/___backup/domain/author"
	"ca-library-app/internal/domain/entity"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Author
	GetAll(ctx context.Context, limit, offset int) []*entity.Author
	Create(ctx context.Context, dto *author.CreateAuthorDto) *entity.Author
}
