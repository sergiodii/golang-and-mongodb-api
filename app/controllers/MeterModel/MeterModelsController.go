package metermodel

import (
	"encoding/json"
	"net/http"

	metermodels "github.com/sergiodii/golang-and-mongodb-api/app/models/MeterModels"
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

func Delete(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func Index(res http.ResponseWriter, req *http.Request) {
	respondWithJson(res, http.StatusOK, metermodels.All())
}

func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	result := metermodels.InsertOne(req.Body)
	json.NewEncoder(res).Encode(result)
}
