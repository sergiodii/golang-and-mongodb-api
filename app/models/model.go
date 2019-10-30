package model

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Model struct {
	TableName string
}

func (c Model) _GetMongoCollection() *mongo.Collection {
	client := _GetMongoClient()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Erro on connect to database:", err)
	}
	collection, err := client.Database(os.Getenv("DATABASE_NAME")).Collection(c.Collection)
	if err != nil {
		log.Fatal("Erro on get collection:", err)
	}
	return collection
}
func (c DataBase) InsertOne()
