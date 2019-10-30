package metermodels

import (
	"github.com/sergiodii/golang-and-mongodb-api/app/models/Model"
)

tableName := "meter_models"

func InsertOne(Obj) {
	m := model.Model {
		TableName: tableName
	}
	return m.InsertOne(Obj)
}