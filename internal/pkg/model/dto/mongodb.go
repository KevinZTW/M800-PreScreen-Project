package dto

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"m800/internal/pkg/config"
	"sync"
)

var (
	once          sync.Once
	mongoDatabase *mongo.Database
)

func MongoDB() *mongo.Database {
	once.Do(func() {
		conf := config.New()
		if cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.MongoDB.Url)); err != nil {
			log.Fatal(err)
		} else {
			mongoDatabase = cli.Database(conf.MongoDB.Name)
		}
	})
	return mongoDatabase
}
