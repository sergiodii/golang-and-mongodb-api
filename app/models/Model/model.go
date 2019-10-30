// package model

// import (
// 	"context"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"github.com/sergiodii/golang-and-mongodb-api/config/database"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// type Model struct {
// 	tablename string
// }

// var clientConection, ctx = database.GetMongoClient()

// func (m Model) _GetMongoCollection() *mongo.Collection {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Erro on load .env file")
// 	}
// 	err = clientConection.Ping(context.Background(), readpref.Primary())
// 	if err != nil {
// 		log.Fatal("Erro on connect to database:", err)
// 	}
// 	collection := clientConection.Database(os.Getenv("DATABASE_NAME")).Collection(m.tablename)
// 	if collection == nil {
// 		log.Fatal("Erro on get collection:", err)
// 	}
// 	return collection
// }

// func (m Model) InsertOne(Obj interface{}) *mongo.InsertOneResult {
// 	collection := m._GetMongoCollection()
// 	result, err := collection.InsertOne(ctx, Obj)
// 	if err != nil {
// 		log.Fatal("Erro on isert data in database:", err)
// 	}
// 	return result
// }
