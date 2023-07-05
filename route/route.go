package route

import (
	"github.com/gorilla/mux"
	"github.com/laioncorcino/personality/controller"
	"github.com/laioncorcino/personality/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMid)
	r.HandleFunc("/api/v1/personality", controller.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/personality/{personalityId}", controller.GetById).Methods("GET")
	r.HandleFunc("/api/v1/personality", controller.Create).Methods("POST")
	r.HandleFunc("/api/v1/personality/{personalityId}", controller.Update).Methods("PUT")
	r.HandleFunc("/api/v1/personality/{personalityId}", controller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
