package book

import (
	"ca-library-app/internal/domain/entity"
)

type bookStorage interface {
	GetOne(uuid string) *entity.Book
	GetAll(limit, offset int) []*entity.Book
	Create(book *entity.Book) *entity.Book
	Delete(book *entity.Book) error
}
