package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DedeMarantes/rest-api-go/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
