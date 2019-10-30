package metermodels

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sergiodii/golang-and-mongodb-api/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var tablename = "meter_models"

type Meter struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand  string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Client string             `json:"client,omitempty" bson:"client,omitempty"`
}

var CollectionInfo, ctx = _GetMongoCollection()

func _GetMongoCollection() (*mongo.Collection, func() context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro on load .env file")
	}
	// clientConection := database.GetMongoClient()
	clientConection, ctxx := database.GetMongoClient()
	err = clientConection.Ping(ctxx(), readpref.Primary())
	if err != nil {
		log.Fatal("Erro on connect to database:", err)
	}
	collection := clientConection.Database(os.Getenv("DATABASE_NAME")).Collection(tablename)
	if collection == nil {
		log.Fatal("Erro on get collection:", err)
	}
	return collection, ctxx
}

func All() []Meter {
	var meters []Meter
	cursor, err := CollectionInfo.Find(ctx(), bson.M{})
	if err != nil {
		log.Fatal("Erro on load all data: ", err)
	}
	defer cursor.Close(ctx())
	for cursor.Next(ctx()) {
		var atualMeter Meter
		cursor.Decode(&atualMeter)
		meters = append(meters, atualMeter)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal("Error: ", err)
	}
	return meters
}

func InsertOne(ObjJson io.Reader) *mongo.InsertOneResult {
	var meter Meter
	json.NewDecoder(ObjJson).Decode(&meter)
	result, err := CollectionInfo.InsertOne(ctx(), ObjJson)
	if err != nil {
		log.Fatal("Erro on isert data in database: ", err)
	}
	return result
}
