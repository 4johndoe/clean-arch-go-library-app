package book

import (
	"ca-library-app/internal/domain/author"
)

type authorStorage struct {
}

func (bs *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}

func (bs *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}

func (bs *authorStorage) Create(author *author.Author) *author.Author {
	return nil
}

func (bs *authorStorage) Delete(author *author.Author) error {
	return nil
}
