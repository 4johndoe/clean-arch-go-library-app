package mongodb

import (
	"ca-library-app/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewAuthorStorage(db *mongo.Database) *authorStorage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(id string) *entity.Author {
	return nil
}

func (bs *authorStorage) GetAll(limit, offset int) []*entity.Author {
	return nil
}

func (bs *authorStorage) Create(author *entity.Author) *entity.Author {
	return nil
}

func (bs *authorStorage) Delete(author *entity.Author) error {
	return nil
}
