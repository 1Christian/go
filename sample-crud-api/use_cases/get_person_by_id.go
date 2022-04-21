package use_cases

import (
	"crud-app/client"
	"crud-app/entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var person entities.Person
	client.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
