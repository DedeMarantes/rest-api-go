package routes

import (
	"log"
	"net/http"

	"github.com/DedeMarantes/rest-api-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8500", r))
}
