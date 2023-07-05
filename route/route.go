package route

import (
	"github.com/gorilla/mux"
	"github.com/laioncorcino/personality/controller"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/personality", controller.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/v1/personality/{personalityId}", controller.GetById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
