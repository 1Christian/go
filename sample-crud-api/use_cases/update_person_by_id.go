package use_cases

import (
	"crud-app/client"
	"crud-app/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entities.Person
	json.Unmarshal(requestBody, &person)
	client.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}
