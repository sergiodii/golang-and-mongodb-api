package main

import (
	"fmt"
	"log"
	"net/http"

	metermodelController "github.com/sergiodii/golang-and-mongodb-api/app/controllers/MeterModel"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro on load .env file")
	}
	route := mux.NewRouter()
	route.HandleFunc("/", metermodelController.Index).Methods("GET")
	route.HandleFunc("/{id}", metermodelController.Show).Methods("GET")
	route.HandleFunc("/", metermodelController.Store).Methods("POST")
	route.HandleFunc("/{id}", metermodelController.Update).Methods("PUT")

	var port = ":8000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, route))
}
