package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/laioncorcino/personality/database"
	"github.com/laioncorcino/personality/entity"
	"net/http"
)

func GetAll(resp http.ResponseWriter, req *http.Request) {
	var p []entity.Personality
	database.DB.Find(&p)
	_ = json.NewEncoder(resp).Encode(p)
}

func GetById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["personalityId"]
	var p entity.Personality
	database.DB.First(&p, id)
	_ = json.NewEncoder(resp).Encode(p)
}

func Create(resp http.ResponseWriter, req *http.Request) {
	var p entity.Personality
	_ = json.NewDecoder(req.Body).Decode(&p)
	database.DB.Create(&p)
	_ = json.NewEncoder(resp).Encode(p)
}

func Delete(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["personalityId"]
	var p entity.Personality
	database.DB.Delete(&p, id)
	_ = json.NewEncoder(resp).Encode(p)
}

func Update(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["personalityId"]
	var p entity.Personality
	database.DB.First(&p, id)
	_ = json.NewDecoder(req.Body).Decode(&p)
	database.DB.Save(&p)
	_ = json.NewEncoder(resp).Encode(p)
}
