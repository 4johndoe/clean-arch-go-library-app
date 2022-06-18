package main

import (
	"ca-library-app/internal/composites"
	"ca-library-app/internal/config"
	"ca-library-app/pkg/logging"
	"context"
	"fmt"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("mongo composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), "", "", "", "", "", "")
	if err != nil {
		logger.Fatal("mongodb composit failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}

	logger.Info("book  composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}

	// register in router
	fmt.Println(authorComposite, bookComposite, cfg)
}
