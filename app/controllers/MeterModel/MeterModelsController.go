package metermodel

import (
	"encoding/json"
	"net/http"
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
	teste := []Lista{
		{Name: "sergio", List: "asdasdas"},
		{Name: "sergio", List: "asdasdas"},
	}
	respondWithJson(res, http.StatusOK, teste)
}
