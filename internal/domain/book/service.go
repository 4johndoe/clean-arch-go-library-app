package book

import "context"

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{storage: storage}
}

func (s *service) GetBookByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAllBooks(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
