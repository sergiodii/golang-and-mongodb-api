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

type MeterWithId struct {
	ID     primitive.ObjectID     `json:"_id,omitempty" bson:"_id"`
	Brand  string                 `json:"brand" bson:"brand"`
	Client string                 `json:"client" bson:"client"`
	Test   map[string]interface{} `json:"test" bson:"test,omitempty"`
}
type Meter struct {
	Brand  string                 `json:"brand" bson:"brand"`
	Client string                 `json:"client" bson:"client"`
	Test   map[string]interface{} `json:"test" bson:"test"`
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

func All() []MeterWithId {
	var meters []MeterWithId
	cursor, err := CollectionInfo.Find(ctx(), bson.M{})
	if err != nil {
		log.Fatal("Erro on load all data: ", err)
	}
	defer cursor.Close(ctx())
	for cursor.Next(ctx()) {
		var atualMeter MeterWithId
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
	result, err := CollectionInfo.InsertOne(ctx(), meter)
	if err != nil {
		log.Fatal("Erro on isert data in database: ", err)
	}
	return result
}

func GetById(id string) MeterWithId {
	var meter MeterWithId
	_ID, _ := primitive.ObjectIDFromHex(id)
	err := CollectionInfo.FindOne(ctx(), bson.M{"_id": _ID}).Decode(&meter)
	if err != nil {
		log.Fatal("Erro on featch data in database: ", err)
	}
	return meter
}

func FindBy(list bson.M) []MeterWithId {
	var meters []MeterWithId
	cursor, err := CollectionInfo.Find(ctx(), list) // .FindOne(ctx(), bson.Marshal(list)).Decode(&meter)
	if err != nil {
		log.Fatal("Erro on featch data in database: ", err)
	}
	defer cursor.Close(ctx())
	for cursor.Next(ctx()) {
		var atualMeter MeterWithId
		cursor.Decode(&atualMeter)
		meters = append(meters, atualMeter)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal("Error: ", err)
	}
	return meters
}
