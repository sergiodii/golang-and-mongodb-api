package main

import (
	"fmt"
	"log"
	"net/http"

	metermodel "github.com/sergiodii/golang-and-mongodb-api/app/controllers/MeterModel"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro on load .env file")
	}
	server := mux.NewRouter()
	server.HandleFunc("/", metermodel.Index).Methods("GET")

	var port = ":8000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, server))
}
