package author

import (
	"ca-library-app/internal/domain/entity"
)

type authorStorage interface {
	GetOne(uuid string) *entity.Author
	GetAll(limit, offset int) []*entity.Author
	Create(author *entity.Author) *entity.Author
	Delete(author *entity.Author) error
}
