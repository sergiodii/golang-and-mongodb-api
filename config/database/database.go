package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func gererateContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetMongoClient() (*mongo.Client, func() context.Context) {
	ctx, _ := context.WithTimeout(gererateContext(), 10*time.Second)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro on load .env file")
	}
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")

	Url := "mongodb://" + host + ":" + port + "/" + name

	if user != "" {
		Url += "?authSource=" + pass + " --username " + user
	}
	clientOptions := options.Client().ApplyURI(Url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, gererateContext
}
