package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/laioncorcino/personality/database"
	"github.com/laioncorcino/personality/entity"
	"net/http"
)

func AllPersonalities(resp http.ResponseWriter, req *http.Request) {
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
