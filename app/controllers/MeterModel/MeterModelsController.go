package metermodel

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	metermodels "github.com/sergiodii/golang-and-mongodb-api/app/models/MeterModels"
	"go.mongodb.org/mongo-driver/bson"
)

type Lista struct {
	Name string
	List string
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }
func Index(res http.ResponseWriter, req *http.Request) {
	var tmpVar = make(map[string]interface{})
	brand, ok := req.URL.Query()["brand"]
	if ok {
		tmpVar["brand"] = brand[0]
	}
	if len(tmpVar) == 0 {
		tmpVar["message"] = "Vazio"
		fmt.Print("bom-dia")
	}
	respondWithJson(res, http.StatusOK, metermodels.All())
}

func Show(res http.ResponseWriter, req *http.Request) {
	parans := mux.Vars(req)
	respondWithJson(res, http.StatusOK, metermodels.GetById(parans["id"]))
}

func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	result := metermodels.InsertOne(req.Body)
	json.NewEncoder(res).Encode(result)
}
func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	// parans := mux.Vars(req)
	fmt.Println("params", req.URL.Query())
	var tmpVar bson.M // make(map[string]interface{})
	brand, ok := req.URL.Query()["brand"]
	if ok {
		tmpVar["brand"] = brand[0]
	}
	if len(tmpVar) == 0 {
		tmpVar["brand"] = "sergio"
		fmt.Print("bom-dia")
	}
	metermodels.FindBy(tmpVar)
	respondWithJson(res, http.StatusOK, metermodels.FindBy(tmpVar))
	// _ID, _ := primitive.ObjectIDFromHex(parans["id"])
	// response, _ := json.Marshal(tmpVar)
	// res.Write(response)

}
