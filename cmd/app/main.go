package main

import (
	"ca-library-app/pkg/logging"
)

func main() {
	logging.Init()
	//logger := logging.GetLogger()

	// TODO composition
	//logger.Info("config initializing")
	//cfg := config.GetConfig()
	//
	//logger.Info("router initializing")
	//router := httprouter.New()
	//
	//logger.Info("mongo composite initializing")
	//mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, "")
	//if err != nil {
	//	logger.Fatal("mongodb composit failed")
	//}
	//
	//logger.Info("author composite initializing")
	//authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	//if err != nil {
	//	logger.Fatal("author composite failed")
	//}
	//authorComposite.Handler.Register(router)
	//
	//logger.Info("book  composite initializing")
	//bookComposite, err := composites.NewBookComposite(mongoDBC)
	//if err != nil {
	//	logger.Fatal("book composite failed")
	//}
	//bookComposite.Handler.Register(router)
	//
	//// start app
	//fmt.Println(authorComposite, bookComposite, cfg)
}
