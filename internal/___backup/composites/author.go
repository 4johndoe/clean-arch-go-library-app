package composites

import (
	"ca-library-app/internal/___backup/adapters/api"
	"ca-library-app/internal/___backup/adapters/api/author"
	author2 "ca-library-app/internal/adaptors/db"
	"ca-library-app/internal/adaptors/db/mongodb"
	service2 "ca-library-app/internal/domain/service"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := mongodb.NewStorage(mongoComposite.db)
	service := service2.NewService(storage)
	handler := author.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
