package composites

import (
	"ca-library-app/internal/___backup/adapters/api"
	"ca-library-app/internal/___backup/adapters/api/author"
	author2 "ca-library-app/internal/adaptors/db"
	"ca-library-app/internal/adaptors/db/mongodb"
	"ca-library-app/internal/controller/http/v1"
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
	handler := v1.NewHandler(service)

	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
