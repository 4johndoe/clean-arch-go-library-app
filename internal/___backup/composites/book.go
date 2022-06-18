package composites

import (
	"ca-library-app/internal/___backup/adapters/api"
	"ca-library-app/internal/___backup/adapters/api/book"
	"ca-library-app/internal/adaptors/db"
	"ca-library-app/internal/adaptors/db/mongodb"
	book4 "ca-library-app/internal/domain/service"
)

type BookComposite struct {
	Storage db.Storage
	Service book.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := mongodb.NewStorage(mongoComposite.db)
	service := book4.NewService(storage)
	handler := book.NewHandler(service)

	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
