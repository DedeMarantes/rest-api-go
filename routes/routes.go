package routes

import (
	"log"
	"net/http"

	"github.com/DedeMarantes/rest-api-go/controllers"
	"github.com/DedeMarantes/rest-api-go/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.SetHeader)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriarPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8500", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
