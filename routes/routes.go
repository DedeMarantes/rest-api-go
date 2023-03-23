package routes

import (
	"log"
	"net/http"

	"github.com/DedeMarantes/rest-api-go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe("localhost:8500", nil))
}
